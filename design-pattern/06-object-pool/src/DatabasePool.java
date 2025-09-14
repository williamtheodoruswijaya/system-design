import java.util.ArrayList;

public class DatabasePool {
    // Simple Implementation

    // 1. create Pool of Connection in an ArrayList
    private static ArrayList<Connection> pool = new ArrayList<>();

    // 2. insert Connection to Pool
    static {
        for (int i = 0; i < 100; i++) {
            Connection conn = new Connection("localhost", "root", "");
            pool.add(conn);
        }
    }

    public static Connection getConnection() {
        if (pool.isEmpty()) {
            // Blocking logic or set default value
            throw new RuntimeException("Database Pool is Empty");
        } else {
            return pool.get(0);
        }
    }

    public static void closeConnection(Connection conn) {
        pool.add(conn);
    }
}
