import java.util.ArrayList;

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