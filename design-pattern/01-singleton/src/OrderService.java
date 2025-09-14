public class OrderService {
    public void save(String orderId) {
//        contoh yang salah
//
//        // 1. make connection to database
//        Connection connection = new Connection("localhost", "root", "");
//
//        // 2. establish sql connection
//        connection.sql("INSERT INTO ORDER ...");

//        Contoh yang benar
        DatabaseHelper.getConnection().sql("INSERT INTO ORDER ...");
    }
}
