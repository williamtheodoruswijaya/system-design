public class OrderService {
    public void save(String orderId) {
        Connection conn = DatabasePool.getConnection();
        conn.sql("INSERT INTO ORDER ...");
    }
}
