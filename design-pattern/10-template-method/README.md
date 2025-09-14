# Template Method

Method yang digunakan sebagai template agar proses pembuatan method serupa dengan logic yang sama dan sedikit perbedaan bisa dilakukan dengan mudah. Kalau dari [sini](https://refactoring.guru/design-patterns/template-method/go/example#example-0), katanya "Template Method is a behavioral design pattern that allows you to define a skeleton of an algorithm in a base class and let subclasses override the steps without changing the overall algorithm’s structure.".

## Contoh Kasus:

Dalam suatu aplikasi, untuk proses login-nya ada proses pengiriman OTP untuk verifikasi, namun, ada sedikit perbedaan untuk pengiriman OTP via Email dan pengiriman OTP via SMS. Sebenarnya, kita bisa aja buat salah satu method dan copas aja ke method kedua yang udah diubah. Tapi hal kek gitu justru ga scalable karena kita harus ngelakuin copas secara terus-menerus, dan kalau seandainya codenya ga disentuh dalam waktu lama, kita akan lupa lagi dengan logicnya dan perlu waktu tambahan buat memahami ulang codenya. Template Method mempermudah hal ini dengan membuat class yang ingin mengadaptasi method serupa tinggal melakukan beberapa perubahan dengan cara mengisi abstract method yang ada dengan logic sendiri (biasanya perbedaan return value doang).
<br/>
Pengiriman OTP terdiri dari step-by-step seperti ini:
- Generate a random n digit number.
- Save this number in the cache for later verification.
- Prepare the content.
- Send the notification.

## Contoh Code:

1. Template Method:
```go
type IOtp interface {
    genRandomOTP(int) string
    saveOTPCache(string)
    getMessage(string) string
    sendNotification(string) error
}
type Otp struct {
    iOtp IOtp
}

func NewOTP(iOtp IOtp) *Otp {    // kenapa return structnya bukan interfacenya? karena yang mau digunakan adalah method `genAndSendOTP()`, kalau return interface tar malah method-method yang di define di interface yang bisa dijalankan sedangkan `genAndSendOTP()`-nya gabisa.
    return &Otp{iOtp}
}

func (o *Otp) genAndSendOTP(otpLength int) error {    // template methodnya
    otp := o.iOtp.genRandomOTP(otpLength)
    o.iOtp.saveOTPCache(otp)
    message := o.iOtp.getMessage(otp)
    err := o.iOtp.sendNotification(message)
    if err != nil {
        return err
    }
    return nil
}
```
Seperti yang terlihat pada file diatas, kita ga define sama sekali bagaimana function `genRandomOTP()`, `getMessage()`, dan `sendNotification()` sama sekali. Kalau di Java, anggap saja mereka" ini sebagai abstract method yang harus diimplementasikan di child class-nya. Nah, jadinya nanti tinggal gini:

2. Implementation (Sms & Email class must implements IOTP interface which means they have to have that 4 methods with its own logic)

sms.go:
```go
type Sms struct {}

func (s *Sms) genRandomOTP(len int) string {
    randomOTP := "1234"
    fmt.Printf("SMS: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
    fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
    return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
    fmt.Printf("SMS: sending sms: %s\n", message)
    return nil
}
```

email.go:
```go
type Email struct {}

func (s *Email) genRandomOTP(len int) string {
    randomOTP := "1234"
    fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
    fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
    return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
    fmt.Printf("EMAIL: sending email: %s\n", message)
    return nil
}
```

Nah, jadi kurang lebih tar tiap class yang mau implementasi Template Method tersebut cukup implementasikan aja abstract method atau singkatnya method-method yang ada di interface `IOtp`.

3. main.go implementation
```go
package main

import "fmt"

func main() {
    // to have otp object method `genAndSendOTP()` to work, OTP object must have iOTP implemented on other classes (in this cases are SMS and Email Classes)

    // 1. SMS Implementation
    smsOTP := &Sms{}
    o := NewOTP(smsOTP)
    o.genAndSendOTP(4)

    // Email Implementation
    emailOTP := &Email{}
    o = NewOTP(emailOTP)
    o.genAndSendOTP(4)
}
```
