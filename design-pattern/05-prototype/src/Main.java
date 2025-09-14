//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
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