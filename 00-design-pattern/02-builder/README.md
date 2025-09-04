# Builder

### Apa itu Builder Design Pattern?

Pattern dimana kita memisahkan cara pembuatan Object, dari class Objectnya sendiri.

### Contoh Kasus:

```java
public class Customer {
    private int id;
    private String firstName;
    private String lastName;
    private String email;
    private String phone;

    // this two are the additional variables
    private String address;
    private int age;

    /*
    - to prevent error on constructor across our apps, we'll use overload function where one was the previous constructor
    - and one as the new constructor with new attributes as the parameter
    - the drawback of this pattern is that we need to create a new overload constructor for any combination of assignable attributes where there are some nullable attributes
    */
    public Customer(int id, String firstName, String lastName, String email, String phone) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
        this.phone = phone;
        this.address = "";
        this.age = 0;
    }

    public Customer(int id, String firstName, String lastName, String email, String phone, String address, int age) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
        this.phone = phone;
        this.address = address;
        this.age = age;
    }
}

```
Contoh-nya:
<ol>
    <li>Misal, kita punya class Customer beserta Attribute-attributenya.</li>
    <li>Biasanya, kita akan buat sebuah Constructor untuk membuat object Customer tersebut.</li>
    <li>Dan Constructor tersebut akan digunakan untuk membuat Object tersebut.</li>
    <li>Tapi, ternyata ada kasus dimana kita ingin memodifikasi atribut dalam class kita.</li>
    <li>Otomatis, kita akan membuat atribut baru dan Constructor terkait juga harus kita modifikasi (ditambahin atributnya).</li>
    <li>Nah, lalu apa yang terjadi kalau app sudah besar dan Constructornya ada banyak, maka akan sulit untuk kita maintain untuk error-errornya.</li>
    <li>Solusinya, bisa kita tambahin default value pada atribut yang ada di Constructor.</li>
    <li>Jadi setiap atribut baru yang kita buat pada sebuah class, kita bisa akalin constructornya dengan menambahkan default value pada constructor tersebut.</li>
    <li>Tapi ini bakal jadi masalah, dimana kalau kita ingin assign skip seperti kasus dibawah, misal mau assign age tapi ga mau assign address.</li>
    <li>Otomatis, harus buat overload function yang baru lagi dan terus-menerus kek begitu.</li>
</ol>

### Solusi dari Builder Pattern

Menurut Builder Pattern, untuk membuat object Customer, jangan menggunakan Constructor secara Direct. Tapi, kita bisa menggunakan Builder function dari class terpisah, dan memanggil Constructor Customer didalam Builder Function tersebut.

#### Step 1: Copy paste semua atribut yang ada dalam class Customer (khusus yang tambahan, kasih default value langsung di atributenya)
```java
public class CustomerBuilder {
    // step 1: create the same attributes on Customer Classes
    private int id;
    private String firstName;
    private String lastName;
    private String email;
    private String phone;

    // atribut tambahan (default value-nya kita set disini)
    private String address = "";
    private int age = 0;
}
```

#### Step 2: Buat Setter Function untuk semua atribut, dan makesure return objek baru yang sudah dibuat.
```java
// step 2: create a setter function for each attributes and changes void to return CustomerBuilder class (this must in order for the builder function to work)
public CustomerBuilder setId(int id) {
    this.id = id;
    return this;
}

public CustomerBuilder setFirstName(String firstName) {
    this.firstName = firstName;
    return this;
}

public CustomerBuilder setLastName(String lastName) {
    this.lastName = lastName;
    return this;
}

public CustomerBuilder setEmail(String email) {
    this.email = email;
    return this;
}

public CustomerBuilder setPhone(String phone) {
    this.phone = phone;
    return this;
}

public CustomerBuilder setAddress(String address) {
    this.address = address;
    return this;
}

public CustomerBuilder setAge(int age) {
    this.age = age;
    return this;
}
```

### Step 3: Buat build function (ini yang akan kita pakai instead constructor function dari object yang diinginkan)
```java
// step 3: create a build function (we will use this function instead of using Customer constructor)
public Customer build() {
    return new Customer(
            this.id,
            this.firstName,
            this.lastName,
            this.email,
            this.phone,
            this.address,
            this.age
    );
}
```

### Step 4: Cara menggunakannya, semudah pakai setter yang ada dan assign setiap attribute yang diinginkan, baru panggil Build()
```java
        Customer customer = new CustomerBuilder()
                .setFirstName("Budi")
                .setLastName("Kurniawan")
                .setEmail("budi@gmail.com")
                .setPhone("123")
                .build();
```
Seperti yang terlihat, kolom id tidak kita assign pun tidak akan menghasilkan error sekarang. Nah, jadi setiap kali kita mau membuat object Customer, kita tidak perlu memanggil Constructor dari Customer tapi panggil Builder functionnya.