//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
//        Cara yang salah (pake constructor dari class langsung)
//        Customer customer = new Customer(1, "Budi", "Kurniawan", "budi@gmail.com", "123", "Jakarta", 30);

//        Cara yang benar (pakai builder function dan cukup masukkan atribut yang diinginkan)
        Customer customer = new CustomerBuilder()
                .setFirstName("Budi")
                .setLastName("Kurniawan")
                .setEmail("budi@gmail.com")
                .setPhone("123")
                .build();
//        Lihat id aja tak kosongin ini masih jalan ges
//        Terus nanti kalau ada atribut baru, kita cukup tambahin di Build function dan gaush takut akan ngaruh ke file-file yang lain
    }
}