# Composite

## Contoh Kasus

Misal, kita punya class seperti ini:
```java
public class Category {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Category(String name) {
        this.name = name;
    }
}
```
Terus di main ceritanya kek begini:
```java
import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        ArrayList<Category> list = new ArrayList<>();
        list.add(new Category("Handphone"));
        list.add(new Category("Operating System"));
        list.add(new Category("Computer"));
        list.add(new Category("Fashion"));

        list.forEach(item -> {
            System.out.println(item.getName());
        });
    }
}
```
Sekarang, gimana caranya kalau kita ingin memberi subcategory dari masing-masing category yang ada, seperti:
```md
Handphone:
- Android
- iOS

Operating System:
- Windows
- Linux
- iOS

Computer:
- Laptop
- PC

Fashion:
- Woman
- Man
```
Oh, ywd kita bisa aja pakai Bridge Design Pattern kek sebelumnya. Oke, bisa" aja, tapi gimana kalau mau ada subcategory lagi? Bridge Design Pattern memang cocok untuk memisahkan sebuah class menjadi subcategory, tapi beda kasusnya kalau kek gini:
```java
Handphone:
- Android
    - Oppo
    - Samsung
    - Xiaomi
- iOS
    - Iphone

Operating System:
- Windows
    - ASUS
    - Acer
    - HP
- Linux
    - Ubuntu
- iOS
    - Macbook

Computer:
- Laptop
- PC

Fashion:
- Woman
- Man
```
Dan seterusnya. Kalau kek begini, jelas Bridge Design Pattern tidak akan optimal/cocok untuk pembuatan subcategory dari subcategory dstnya. Nah, Composite Design Pattern lah yang harus digunakan disini.<br/>
Kita bisa membayangkan ini sebagai sebuah tree dimana 1 parent class, bisa memiliki lebih dari 1 child class. Caranya, sesimple buat Child Class dari class Category, kemudian buat Child Class tersebut, memiliki sebuah List of Item yang bertipekan Parent dari Class tersebut.<br/>
Dalam context ini anggep aja kek begini:
```java
import java.util.ArrayList;

public class CompositeCategory extends Category {
    // ArrayList berisikan parent dari class ini
    private ArrayList<Category> list = new ArrayList<>();

    public CompositeCategory(String name) {
        super(name);
    }

    // Karena kita pakai ArrayList, wajib punya method buat add and remove ke list
    public void add(Category item) {
        list.add(item);
    }

    public void remove(Category item) {
        list.remove(item);
    }

    public ArrayList<Category> getList() {
        return list;
    }

    public void setList(ArrayList<Category> list) {
        this.list = list;
    }
}
```
Jadi, sebenarnya `CompositeCategory` yang sebenarnya adalah **Child** dari class `Category`, memiliki sebuah list yang justru berisikan parent dari `CompositeCategory` yaitu `Category`. Dengan begini, kita bisa merealisasikan subcategory diatas.
```java
public class Main {
    public static void main(String[] args) {
        ArrayList<Category> list = new ArrayList<>();

        // Handphone
        CompositeCategory handphone = new CompositeCategory("Handphone");
        CompositeCategory android = new CompositeCategory("Android");
        android.add(new CompositeCategory("Oppo"));
        android.add(new CompositeCategory("Samsung"));
        android.add(new CompositeCategory("Xiaomi"));
        handphone.add(android);

        CompositeCategory ios = new CompositeCategory("iOS");
        ios.add(new CompositeCategory("iPhone"));
        handphone.add(ios);

        list.add(handphone);


        // Operating System
        CompositeCategory operatingSystem = new CompositeCategory("Operating System");
        CompositeCategory windows = new CompositeCategory("Windows");
        windows.add(new CompositeCategory("ASUS"));
        windows.add(new CompositeCategory("HP"));
        windows.add(new CompositeCategory("Acer"));
        operatingSystem.add(windows);

        CompositeCategory linux = new CompositeCategory("Linux");
        linux.add(new CompositeCategory("Ubuntu"));
        operatingSystem.add(linux);

        CompositeCategory iOSPC = new CompositeCategory("iOS");
        iOSPC.add(new CompositeCategory("Macbook"));
        iOSPC.add(new CompositeCategory("iMac"));
        operatingSystem.add(iOSPC);

        list.add(operatingSystem);


        // Computer
        CompositeCategory computer = new CompositeCategory("Computer");
        computer.add(new CompositeCategory("PC"));
        computer.add(new CompositeCategory("Laptop"));

        list.add(computer);


        // Fashion
        CompositeCategory fashion = new CompositeCategory("Fashion");
        fashion.add(new CompositeCategory("Men"));
        fashion.add(new CompositeCategory("Women"));

        list.add(fashion);

        list.forEach(item -> {
            item.display("");
        });
    }
}
```
Output:
```md
+ Handphone
    + Android
        + Oppo
        + Samsung
        + Xiaomi
    + iOS
        + iPhone
+ Operating System
    + Windows
        + ASUS
        + HP
        + Acer
    + Linux
        + Ubuntu
    + iOS
        + Macbook
        + iMac
+ Computer
    + PC
    + Laptop
+ Fashion
    + Men
    + Women
```
Nah, inti dari code diatas emang complex banget sih, tapi intinya, mau ada berapa subcategory pun itu bisa. Taruh aja `CompositeClass` dalam `CompositeClass`.