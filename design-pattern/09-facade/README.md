# Facade

Nah, kalau untuk Facade, ini lebih ke arah logic dimana kita kadang suka memikirkan untuk memisahkan logic-logic yang ribet ke sebuah function tersendiri agar code lebih readable.
<br/>
Anggap kasus dari code repository sebelumnya.
```php
class ArticleController extends Controller {
  public function createArticle(Request $req){
    $id = $req->get('id');
    $repo = new ArticleRepository();
    $article = $repo->findById($id);
    return response()->json(['article'->...]);
  }
}

class ArticleRepository {
  public function findById($id){
    $query = "SELECT * FROM Article WHERE id = {$id}";
    $article = $query->getRow(0, \App\Entities\User::class);
    return $article;
  }
}
```

Nah sebenarnya dari 2 class ini, pada bagian Controller, sebenarnya kita menyatukan Business Logic kita dan Repository kita dalam class Controller. Yang padahal sebenarnya, karena Query Logic sudah dikeluarkan ke dalam class Repository, ada baiknya Business Logic juga dipisahkan ke dalam class tersendiri. 
<br/>
Nah, metode ini kita kenal dengan nama **Facade Design Pattern**.

<img width="740" height="384" alt="image" src="https://github.com/user-attachments/assets/9baaab10-edc8-4249-9d79-bb47396d81c4" />

Ini juga yang sebenarnya dilakukan dalam Layered Architecture.

Jadi, ketimbang kita taruh business logic dalam controller kita, kita buat jadi 3 class (Service/Usecase):
```php
class ArticleController extends Controller {
  public function createArticle(Request $req){
    $usecase = new ArticleUsecase();
    return $usecase->createArticle($req);
  }
}

class ArticleUsecase {
  public function createArticle(Request $req){
    $id = $req->get('id');
    $repo = new ArticleRepository();
    $article = $repo->findById($id);
    return response()->json(['article'->...]);
  }
}

class ArticleRepository {
  public function findById($id){
    $query = "SELECT * FROM Article WHERE id = {$id}";
    $article = $query->getRow(0, \App\Entities\User::class);
    return $article;
  }
}
```

Terkadang juga hal seperti ini bisa dibilang sebagai Facade Design Pattern juga:
```go
func (s *FriendServiceImpl) AddFriend(ctx context.Context, req request.FriendRequest) (*response.FriendResponse, error) {

  // ...

	// step 3: validate request
	if err := utils.ValidateFriendInput(req.UserID, req.FriendUserID, 0); err != nil {
		return nil, err
	}

  // ...
}
```
```go
package utils

import "fmt"

func ValidateFriendInput(userID, friendUserID, friendStatus int) error {
	if userID <= 0 {
		return fmt.Errorf("user_id must be greater than 0")
	}
	if friendUserID <= 0 {
		return fmt.Errorf("friend_user_id must be greater than 0")
	}
	if friendStatus < 0 || friendStatus > 2 {
		return fmt.Errorf("friend_status must be between 0 and 2")
	}
	if userID == friendUserID {
		return fmt.Errorf("user_id and friend_user_id cannot be the same")
	}
	return nil
}
```
Nah, ketimbang kita buat banyak `if` dalam `AddFriend()` hanya untuk validasi, kita bisa pisahkan jadi function di package lain. Ini juga bisa disebut sebagai Facade Design Pattern.
