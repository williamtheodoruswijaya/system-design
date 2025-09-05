public class OrderDetailService {
    public void save(String orderId, String product) {
        Connection conn = DatabasePool.getConnection();
        conn.sql("INSERT INTO ORDER_DETAIL ...");
    }
}
