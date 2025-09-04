public class OrderDetailService {
    public void save(String orderId, String product) {
//        Contoh implementasi yang salah:
//
//        Connection connection = new Connection("localhost", "root", "");
//        connection.sql("INSERT INTO ORDER_DETAIL...");

//        Contoh yang benar
        DatabaseHelper.getConnection().sql("INSERT INTO ORDER_DETAIL ...");
    }
}
