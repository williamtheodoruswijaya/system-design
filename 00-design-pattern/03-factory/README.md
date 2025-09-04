# Factory

### Apa itu Factory Method?
Factory Method intinya adalah sebuah method yang digunakan ketika kita ingin membuat object yang ada jenisnya dan jenis tersebut terbatas (3 doang misal). Nah, agar kita bisa membatasinya, best practicesnya kita menggunakan Factory Design Pattern ini.

### Contoh kasus 1:
Misal, kita punya Employee Class yang punya 2 jenis role yaitu Manager dan Staff.
```java
        Employee manager1 = new Employee("Rudi", "Manager", 10000000);
        Employee manager2 = new Employee("Budi", "Manager", 10000000);

        Employee staff1 = new Employee("Eko", "Staff", 5000000);
        Employee staff2 = new Employee("Roni", "Staff", 5000000);
```
Nah, instead kita set mereka secara manual, kita pakai Factory method untuk membuat masing-masing class:
```java
        EmployeeFactory employeeFactory = new EmployeeFactory();
        Employee manager1 = employeeFactory.createManager("Rudi");
        Employee manager2 = employeeFactory.createManager("Budi");

        Employee staff1 = employeeFactory.createStaff("Eko");
        Employee staff2 = employeeFactory.createStaff("Roni");
```
Nanti, isi dari Factory Method literally:
```java
public class EmployeeFactory {
    public Employee createManager(String name) {
        return new Employee(name, "Manager", 10000000);
    }

    public Employee createStaff(String name) {
        return new Employee(name, "Staff", 5000000);
    }
}
```
Hal ini jelas lebih bagus ketimbang kita buat secara satu per satu.

### Contoh kasus 2:
Misal, kita punya Animal interface yang diimplementasikan oleh class-class ini:
```java
interface Animal {
    public void speak();
}

class Tiger implements Animal {
    public void speak() {
        System.out.println("Rawr!");
    }
}

class Dog implements Animal {
    public void speak() {
        System.out.println("Woof!");
    }
}

class Cat implements Animal {
    public void speak() {
        System.out.println("Meow!");
    }
}
```
Nah, biasanya kita akan initiate seperti ini:
```java
        Animal tiger = new Tiger();
        Animal cat = new Cat();
        Animal dog = new Dog();
```
Tapi, jika seandainya ada kesalahan implementasi dalam class Tiger yang mengharuskan Class Tiger diubah namanya jadi Tiger2. Maka secara otomatis, kita harus ubah semua object Tiger yang diinitiate disemua file. Solusinya, kita buat Factory Method aja.
```java
public class AnimalFactory {
    public Animal create(String type) {
        if (type.equalsIgnoreCase("Tiger")) {
            return new Tiger();
        }else if (type.equalsIgnoreCase("Dog")) {
            return new Dog();
        }else {
            return new Cat();
        }
    }
}
```
Nah, sehingga kita tinggal ganti `return new Object`nya aja.
```java
AnimalFactory animalFactory = new AnimalFactory();
Animal tiger = animalFactory.create("tiger");
Animal dog = animalFactory.create("dog");
Animal cat = animalFactory.create("cat");
```

## Conclusions

Gunakan Factory Method Design Pattern jika object yang kita ingin buat memiliki jenis-jenis yang sudah ditentukan dan jenis tersebut cenderung sedikit. Seperti kasus pada class Employee, dimana yang membedakan kan literally nama yah sementara salary dan role itu sifatnya tergantung dengan role dan role hanya ada 2 jenis. Nah, itu bagusnya menggunakan Factory Method Design Pattern ini.