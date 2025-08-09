## Basic Implementation on UseCase/Service layer

Before Goroutines:

```go
func (s *UserServiceImpl) Create(ctx context.Context, request request.CreateUserRequest) (*response.CreateUserResponse, error) {
	// step 1: begin transaction
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	// step 2: rollback
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// step 3: convert request ke model User
	user := entity.User{
		Username:   strings.TrimSpace(strings.ToLower(request.Username)),
		Fullname:   request.Fullname,
		Email:      strings.TrimSpace(strings.ToLower(request.Email)),
		Password:   request.Password,
		ProfileUrl: sql.NullString{String: "https://upload.wikimedia.org/wikipedia/commons/a/ac/Default_pfp.jpg", Valid: true}, // Default URL for new users
		CreatedAt:  time.Now(),
	}

	// Appendix: validate request
	if err := utils.ValidateUserInput(&user); err != nil {
		return nil, err
	}

	// Appendix: validate if username already exists
	existingUser, err := s.UserRepository.Find(ctx, s.DB, user.Username)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("username %s already exists", user.Username)
	}

	// Appendix: validate if email already exists
	existingEmail, err := s.UserRepository.FindByEmail(ctx, s.DB, user.Email)
	if err == nil && existingEmail != nil {
		return nil, fmt.Errorf("email %s already exists", user.Email)
	}

	// step 4: call repository to create user
	createdUser, err := s.UserRepository.Create(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	// step 5: commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// step 6: Find the created user
	result, err := s.Find(ctx, createdUser.Username)
	if err != nil {
		return nil, err
	}
	return result, nil
}
```
Bagian-bagian ini bersifat independen dan tidak bergantungan satu sama lain sehingga bisa dijalankan secara parallel.
```go
	// Appendix: validate request
	if err := utils.ValidateUserInput(&user); err != nil {
		return nil, err
	}

	// Appendix: validate if username already exists
	existingUser, err := s.UserRepository.Find(ctx, s.DB, user.Username)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("username %s already exists", user.Username)
	}

	// Appendix: validate if email already exists
	existingEmail, err := s.UserRepository.FindByEmail(ctx, s.DB, user.Email)
	if err == nil && existingEmail != nil {
		return nil, fmt.Errorf("email %s already exists", user.Email)
	}
```

After Goroutines:

```go
func (s *UserServiceImpl) Create(ctx context.Context, request request.CreateUserRequest) (*response.CreateUserResponse, error) {
	// same step 1 - 3...

	// create the necessary setup
	var wg sync.WaitGroup
	var mu sync.Mutex
	var existingUser, existingEmail *entity.User
	var errUsername, errEmail, errRequest error

	// create a wait group and add 3 value since we are going to run 3 goroutines (the value that were added here are how many goroutines we want to run)
	wg.Add(3)
	
	go func() {
		defer wg.Done() // after the function finish, decrement the value on Add() by 1.
		err := utils.ValidateUserInput(&user)
		if err != nil {
			mu.Lock()	// pokoknya tiap mau assign sebuah value ke dalam variable wajib di lock (build habit)
			errRequest = err
			mu.Unlock()	// unlock kalau udah di assign
		}
	}()

	go func() {
		defer wg.Done()
		user, err := s.UserRepository.Find(ctx, s.DB, user.Username)
		mu.Lock()
		existingUser, errUsername = u, err
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		email, err := s.UserRepository.FindByEmail(ctx, s.DB, user.Email)
		mu.Lock()
		existingEmail, errEmail = email, err
		mu.Unlock()
	}

	wg.Wait() // KODE DIBAWAH TIDAK AKAN DIJALANKAN SEBELUM VALUE PADA ADD() == 0

	if errRequest != nil {
		return nil, fmt.Errorf("error in request body")
	}
	if errUsername == nil && existingUser != nil {
		return nil, fmt.Errorf("username %s already exists!", user.Username)
	}
	if errEmail == nil && existingEmail != nil {
		return nil, fmt.Errorf("email %s already taken!", user.Email)
	}

	// same step 5 - 6...
}
```
