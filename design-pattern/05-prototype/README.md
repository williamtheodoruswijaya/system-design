# Prototype

Prototype Design Pattern adalah Design Pattern yang digunakan ketika kita ingin menggunakan value dari object A ke berbagai object tanpa mengassign value dari object A secara langsung.<br/>
Maksudnya gimana?

### Contoh Kasus 1:
Misal, kita punya class Store sebagai berikut:
```java
public class Store {
    private String name;
    private String city;
    private String country;
    private String category;
    
    public Store(String name, String city, String country, String category) {
        this.name = name;
        this.city = city;
        this.country = country;
        this.category = category;
    }
    
    // Getter and Setter ...
}
```
Dan kita ada kebutuhan sehingga diinisialisasi seperti ini:
```java
public class Main {
    public static void main(String[] args) {
        Store store1 = new Store("Toko X", "Jakarta", "Indonesia", "Gadget");
        Store store2 = new Store("Toko Z", "Jakarta", "Indonesia", "Gadget");
        Store store3 = new Store("Toko Y", "Bandung", "Indonesia", "Gadget");
        Store store4 = new Store("Toko W", "Bandung", "Indonesia", "Fashion");
    }
}
```
Nah, ini seperti yang terlihat, ketimbang kita initialize value secara manual seperti ini, biasanya kita menggunakan value dari store1 karena emang tujuannya reuse value of an object.
```java
public class Main {
    public static void main(String[] args) {
        Store store1 = new Store("Toko X", "Jakarta", "Indonesia", "Gadget");
        Store store2 = new Store("Toko Z", store1.getCity(), store1.getCountry(), store1.getCategory());
        Store store3 = new Store("Toko Y", "Bandung", store1.getCountry(), store1.getCategory());
        Store store4 = new Store("Toko W", store3.getCity(), store1.getCountry(), "Fashion");
    }
}
```
Nah, kalau kayak gini kan annoying banget ya dan sangat tidak teratur. Nah, jadi prototype design function ini bertujuan untuk membuat sebuah clone dari object yang sering dipakai valuenya, dan value-value yang unik tinggal diubah menggunakan Setter. Nah, kalau ga kebayang, kurang lebih kita harus buat method Clonenya dulu di Object-nya.
```java
public class Store {
    private String name;
    private String city;
    private String country;
    private String category;
    
    public Store(String name, String city, String country, String category) {
        this.name = name;
        this.city = city;
        this.country = country;
        this.category = category;
    }
    
    // Getter and Setter ...
    
    // Prototype clone function
    public Store clone() {
        return new Store(
                this.name,
                this.city,
                this.country,
                this.category
        );
    }
}
```
Nah, sekarang alih-alih kita buat secara manual, kita tinggal clone aja terus pake setter buat object yang valuenya beda.
```java
public class Main {
    public static void main(String[] args) {
//        Cara yang salah
//        Store store1 = new Store("Toko X", "Jakarta", "Indonesia", "Gadget");
//        Store store2 = new Store("Toko Z", store1.getCity(), store1.getCountry(), store1.getCategory());
//        Store store3 = new Store("Toko Y", "Bandung", store1.getCountry(), store1.getCategory());
//        Store store4 = new Store("Toko W", store3.getCity(), store1.getCountry(), "Fashion");

//        Cara yang benar
        Store store1 = new Store("Toko X", "Jakarta", "Indonesia", "Gadget");

        Store store2 = store1.clone();
        store2.setName("Toko Z");

        Store store3 = store1.clone();
        store3.setName("Toko Y");
        store3.setCity("Bandung");
        
        Store store4 = store3.clone();
        store4.setName("Toko W");
        store4.setCategory("Fashion");
    }
}
```
Dengan begini akan lebih terbaca, apalagi kalau misal nanti ada puluhan atribut + ada nested data kan bakal pusing kalau manual tapi ada reuse value dari object yang satu ke satu lainnya. Nah, ini kegunaan dari cloning object. Anggep aja kita menjadikan 1 object sebagai prototype buat data lainnya.