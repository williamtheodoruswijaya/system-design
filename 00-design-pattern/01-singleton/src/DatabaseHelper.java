public class DatabaseHelper {
    private static Connection connection;

    public static Connection getConnection() {
//          kalau connection belum pernah dibuat, buat connection baru
     if(connection == null) {
         connection = new Connection("localhost", "root", "root");
     }
//          kalau udah, return connection yang dulu udah dibuat
     return connection;
    }
}
