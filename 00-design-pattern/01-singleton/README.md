# Singleton

### Sample Cases

Misal kita punya file seperti [ini](src/Main.java)
```java
public class Main {
    public static void main(String[] args) {
        OrderService orderService = new OrderService();
        orderService.save("0001");
        
        OrderDetailService orderDetailService = new OrderDetailService();
        orderDetailService.save("0001", "Indomie");
        orderDetailService.save("0001", "Sabun");
        orderDetailService.save("0001", "Pepsodent");
    }
}
```

Nah, jadi ini program untuk menyimpan data order, dimana setiap order pasti mempunyai sebuah detail dimana masing-masing detail adalah detail produk-produk yang dibeli di order tersebut.
Untuk menyimpan data order-nya, pertama kita buat dulu ID-nya, dan kita simpan produk-produknya ke dalam ID tersebut. Artinya anggep aja dalam order 0001, ada produk Indomie, Sabun, dan Pepsodent.
<br/>
Pertanyaannya, gimana cara kita masukin ke dalam databasenya? ya, biasanya kita bakal buat SQL script disetiap class-nya (di bagian method save).
```java
public class OrderService {
    public void save(String orderId) {
        // 1. make connection to database
        Connection connection = new Connection("localhost", "root", "");

        // 2. establish sql connection
        connection.sql("INSERT INTO ORDER ...");
    }
}

public class OrderDetailService {
    public void save(String orderId, String product) {
        Connection connection = new Connection("localhost", "root", "");
        connection.sql("INSERT INTO ORDER_DETAILS...");
    }
}
```
Nah, jadi di setiap class, di method save-nya kita harus implementasi logic SQL yang sama. Sampai ada kejadian dimana aplikasi kita punya traffic yang tinggi dan aplikasi kita jadi lambat.
Kenapa jadi lambat? karena dalam setiap method save, kita membuat koneksi baru ke database. Artinya, bahkan dalam code seperti diatas:
```java
OrderService orderService = new OrderService();
orderService.save("0001");
        
OrderDetailService orderDetailService = new OrderDetailService();
orderDetailService.save("0001", "Indomie");
orderDetailService.save("0001", "Sabun");
orderDetailService.save("0001", "Pepsodent");
```
Terdapat 4 jenis connection yang terbentuk. Nah, kalau traffic lagi besar dan ada 400 orang yang memakai software kita, maka dapat dikatakan ada kemungkinan terdapat 400 connection yang terjadi.
Hal ini jelas membuat aplikasi kita berjalan lambat.<br/>
Nah, padahal sebenarnya kita ga butuh membuat koneksi baru setiap kali kita mau insert. Cukup 1 kali buat connection ke database dan pakai berulang kali. Nah, inilah konsep dari Singleton Design Pattern.

### Singleton Design Pattern

Singleton Design Pattern sederhananya adalah kita membuat 1 objek dalam 1 aplikasi, dimana di tempat manapun dalam aplikasi kita, ketika kita membutuhkan objek tersebut, alih-alih membuat objek baru, kita menggunakan objek yang sama.
Contohnya dalam koneksi ke database, instead kita buat object connection secara berulang kali, kita reuse aja object connectionnya.

##### Gimana cara kita membuat object Connectionnya menjadi 1 Object? 

Nah, stylenya ada banyak, tapi biasanya kita buat 1 class baru sebagai class helper, dimana helper tersebut harus memastikan kalau dipanggil 1 kali, dia mengembalikan sebuah Object Connection, dan ketika dipanggil berulang kali, dia mengembalikan Object Connection yang sudah dibuat.
```java
public class DatabaseHelper {
    private static Connection connection;

    public static Connection getConnection() {
//          kalau connection belum pernah dibuat, buat connection baru
     if(connection == null) {
         connection = new Connection("localhost", "root", "root");
     }
//          kalau udah, return connection yang dulu udah dibuat
     return connection;
    }
}

```