//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
//        Cara yang salah
//        Employee manager1 = new Employee("Rudi", "Manager", 10000000);
//        Employee manager2 = new Employee("Budi", "Manager", 10000000);
//
//        Employee staff1 = new Employee("Eko", "Staff", 5000000);
//        Employee staff2 = new Employee("Roni", "Staff", 5000000);

//        Cara yang benar (buat jadi 1 jenis function aja yang membuat Manager dan Staff instead kita initiate manual)
        EmployeeFactory employeeFactory = new EmployeeFactory();
        Employee manager1 = employeeFactory.createManager("Rudi");
        Employee manager2 = employeeFactory.createManager("Budi");

        Employee staff1 = employeeFactory.createStaff("Eko");
        Employee staff2 = employeeFactory.createStaff("Roni");

//        Cara yang salah
//        Animal tiger = new Tiger();
//        Animal cat = new Cat();
//        Animal dog = new Dog();

//        Cara yang benar (pakai Factory method untuk auto create object based on limited types yang sudah diatur)
        AnimalFactory animalFactory = new AnimalFactory();
        Animal tiger = animalFactory.create("tiger");
        Animal dog = animalFactory.create("dog");
        Animal cat = animalFactory.create("cat");
    }
}