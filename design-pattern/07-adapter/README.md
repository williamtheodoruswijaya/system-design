# Adapter

Design Pattern untuk mengubah berbagai jenis class menjadi 1 jenis class yang lebih general.

### Contoh Kasus 1: E-Commerce Problem Catalog that sells more than one kinds of item.

Misal, kita ada 2 jenis item (ya kalau dalam E-Commerce bisa lebih dari 2 tapi anggep aja 2 dulu).

- Book:
```java
public class Book {
    private String title;
    private String author;

    public Book(String title, String author) {
        this.title = title;
        this.author = author;
    }
    
    // GETTER AND SETTER ...
}
```

- Screencast:
```java
public class Screencast {
    private String title;
    private String author;
    private String duration;

    public Screencast(String title, String author, String duration) {
        this.title = title;
        this.author = author;
        this.duration = duration;
    }
    
    // GETTER AND SETTER ...
}
```

Nah, gimana kalau misal kita mau tampilin dalam 1 halaman? Nah, kalau manual kek gini:
```java
import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        ArrayList<Object> list = new ArrayList<>();

        list.add(new Book("Percy Jackson", "Rick Riordan"));
        list.add(new Book("Red Queen", "Victoria Aveyard"));
        list.add(new Book("Weathering with You", "Makoto Shinkai"));

        list.add(new Screencast("Shadow of Bone", "Netflix", "02:00:00"));
        list.add(new Screencast("Alice in Borderland", "Netflix", "02:30:00"));
        list.add(new Screencast("Sandman", "Netflix", "03:30:00"));

        list.forEach(item -> {
            if (item instanceof Book) {
                Book book = (Book) item;
                System.out.println(book.getTitle() + " by " + book.getAuthor());
            } else if (item instanceof Screencast) {
                Screencast screencast = (Screencast) item;
                System.out.println(screencast.getTitle() + " by " + screencast.getAuthor() + " duration -> " + screencast.getDuration());
            }
            // nambah lagi else if kalau ada jenis baru
        });
    }
}
```
Nah, kebayang ga kalau ada 1000 jenis item. Bakal ada 1000 else if, belum lagi jenis menampilkan nya itu beda-beda jadi harus diatur lagi.
Nah, Adapter Design Pattern itu mirip dengan logika yang muncul ketika melihat problem ini yaitu tinggal pakai interface aja. Nah, caranya kek gini:

#### 1. Buat Interface CatalogAdapter (method yang wajib ada adalah function untuk menampilkan itemnya)
```java
public interface CatalogAdapter {
    public String getCatalogTitle();
}
```

#### 2. Lalu buat class adapter buat masing-masing Item.
fungsinya adalah untuk menciptakan logic khusus untuk `getCatalogTitle()` dan juga mengubah berbagai jenis class menjadi 1 jenis class (`classAdapter` itu sendiri).
```java
public class BookAdapter implements CatalogAdapter{
    private Book book;
    
    @Override
    public String getCatalogTitle() {
        return book.getTitle() + " by " + book.getAuthor();
    }
}

public class ScreencastAdapter implements  CatalogAdapter {
    private Screencast screencast;

    @Override
    public String getCatalogTitle() {
        return screencast.getTitle() + " by " + screencast.getAuthor() + " -> " + screencast.getDuration();
    }
}
```

#### 3. Terus kita tinggal ubah main-nya
```java
import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        ArrayList<CatalogAdapter> list = new ArrayList<>();

        list.add(new BookAdapter(new Book("Percy Jackson", "Rick Riordan")));
        list.add(new BookAdapter(new Book("Red Queen", "Victoria Aveyard")));
        list.add(new BookAdapter(new Book("Weathering with You", "Makoto Shinkai")));

        list.add(new ScreencastAdapter(new Screencast("Shadow of Bone", "Netflix", "02:00:00")));
        list.add(new ScreencastAdapter(new Screencast("Alice in Borderland", "Netflix", "02:30:00")));
        list.add(new ScreencastAdapter(new Screencast("Sandman", "Netflix", "03:30:00")));

        list.forEach(item -> {
            System.out.println(item.getCatalogTitle());
        });
    }
}
```

Nanti kalua ada item baru (misal Movie), tinggal buat CatalogAdapter (Movie Adapter) nya aja terus input deh.