# Bridge

Bridge Design Pattern adalah Design Pattern yang ibaratnya membuat kategori dari sebuah Parent Class Interface.

## Contoh Kasus:

Misal, kita ada interface ini:
```java
package parent;

public interface Binatang {
    public String getNama();
    public boolean hidupDiDarat();
    public boolean hidupDiLaut();
}
```
Nah, class-class yang implements interface ini adalah:
- Kucing:
```java
public class Kucing implements Binatang {
    @Override
    public String getNama() {
        return "Kucing";
    }
    
    @Override
    public boolean hidupDiDarat() {
        return true;
    }
    
    @Override
    public boolean hidupDiLaut() {
        return false;
    }
}
```
- Anjing:
```java
public class Anjing implements Binatang {
    @Override
    public String getNama() {
        return "Anjing";
    }
    
    @Override
    public boolean hidupDiDarat() {
        return true;
    }
    
    @Override
    public boolean hidupDiLaut() {
        return false;
    }
}
```
- Kambing:
```java
public class Kambing implements Binatang {
    @Override
    public String getNama() {
        return "Kambing";
    }
    
    @Override
    public boolean hidupDiDarat() {
        return true;
    }
    
    @Override
    public boolean hidupDiLaut() {
        return false;
    }
}
```
- Hiu:
```java
public class Hiu implements Binatang {
    @Override
    public String getNama() {
        return "Hiu";
    }
    
    @Override
    public boolean hidupDiDarat() {
        return false;
    }
    
    @Override
    public boolean hidupDiLaut() {
        return true;
    }
}
```
- Koi:
```java
public class Koi implements Binatang {
    @Override
    public String getNama() {
        return "Koi";
    }

    @Override
    public boolean hidupDiDarat() {
        return false;
    }

    @Override
    public boolean hidupDiLaut() {
        return true;
    }
} 
```
- Lele:
```java
public class Lele implements Binatang {
    @Override
    public String getNama() {
        return "Lele";
    }

    @Override
    public boolean hidupDiDarat() {
        return false;
    }

    @Override
    public boolean hidupDiLaut() {
        return true;
    }
} 
```
Dan Mainnya nanti seperti ini:
```java
public class Main {
    public static void main(String[] args) {
        ArrayList<Binatang> list = new ArrayList<>();
        list.add(new Anjing());
        list.add(new Kucing());
        list.add(new Monyet());
        list.add(new Hiu());
        list.add(new Koi());
        list.add(new Lele());

        for(Binatang binatang : list) {
            if (binatang.hidupDiDarat()) {
                System.out.println(binatang.getNama() + " hidup di darat");
            } else if (binatang.hidupDiLaut()) {
                System.out.println(binatang.getNama() + " hidup di laut.");
            }
        }
    }
}
```
Sekilas memang tidak ada masalah, tapi gimana jadinya kalau nanti kita mau nambahin method khusus binatang yang hidup di Darat, dia punya atribut jumlah kaki dan jumlahnya tersebut bisa di Get. Otomatis, cara-nya adalah, kita modifikasi object Anjing, Kucing, dan Monyet agar punya method jumlah kaki, kemudian di Main, kita buat jadi seperti ini:

```java
import darat.Anjing;
import darat.Kucing;
import darat.Monyet;

public class Main {
    public static void main(String[] args) {
        ArrayList<Binatang> list = new ArrayList<>();
        list.add(new Anjing());
        list.add(new Kucing());
        list.add(new Monyet());
        list.add(new Hiu());
        list.add(new Koi());
        list.add(new Lele());

        for (Binatang binatang : list) {
            if (binatang.hidupDiDarat()) {
                // wajib buat if-else satu-per-satu
                if (binatang instanceof Anjing) {
                    System.out.println(binatang.getNama() + " hidup di darat dengan kaki " + binatang.getJumlahKaki());
                } else if (binatang instanceof Kucing) {
                    System.out.println(binatang.getNama() + " hidup di darat dengan kaki " + binatang.getJumlahKaki());
                } else if (binatang instanceof Monyet) {
                    System.out.println(binatang.getNama() + " hidup di darat dengan kaki " + binatang.getJumlahKaki());
                }
            } else if (binatang.hidupDiLaut()) {
                System.out.println(binatang.getNama() + " hidup di laut.");
            }
        }
    }
}
```
Nah, intinya ini akan jadi problem kalau kita mau menambahkan method ke beberapa class tertentu dimana di Main, nantinya kita harus buat if-else staircase berdasarkan banyaknya object dengan jenis tersebut. 
Solusinya, kita bisa buat abstract class yang merupakan subkategori dari class Binatang. Metode ini yang disebut sebagai Bridge Design Pattern.

<img width="1335" height="791" alt="image" src="https://github.com/user-attachments/assets/1cdbdf41-85b8-4bd6-aa8f-305f53d3c0e2" />

Nah, tar jadinya begini ges:
- BinatangDarat:
```java
package bridges;

import parent.Binatang;

public abstract class BinatangDarat implements Binatang {
    @Override
    public boolean hidupDiDarat() {
        return true;
    }

    @Override
    public boolean hidupDiLaut() {
        return false;
    }

    public abstract int getJumlahKaki(); // method additional bisa ditambahkan disini dalam bentuk abstract function
}
```
- BinatangLaut:
```java
package bridges;

import parent.Binatang;

public abstract class BinatangLaut implements Binatang {
    @Override
    public boolean hidupDiLaut() {
        return true;
    }

    @Override
    public boolean hidupDiDarat() {
        return false;
    }
}
```
Nah, setelah itu, kita di Main, jadinya lebih simple dan bisa langsung kek begini:
```java
public class Main {
    public static void main(String[] args) {
        ArrayList<Binatang> list = new ArrayList<>();
        list.add(new Anjing());
        list.add(new Kucing());
        list.add(new Monyet());
        list.add(new Hiu());
        list.add(new Koi());
        list.add(new Lele());

        for(Binatang binatang : list) {
            if (binatang instanceof BinatangDarat) {
                BinatangDarat binatangDarat = (BinatangDarat) binatang;
                System.out.println(binatangDarat.getNama() + " hidup di darat dengan " + binatangDarat.getJumlahKaki() + " kaki.");
            } else if (binatang instanceof BinatangLaut) {
                BinatangLaut binatangLaut = (BinatangLaut) binatang;
                System.out.println(binatangLaut.getNama() + " hidup di laut.");
            }
        }
    }
}
```