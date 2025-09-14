public class Connection {
    private String host;
    private String username;
    private String password;

    public Connection(String host, String username, String password){
        this.host = host;
        this.username = username;
        this.password = password;
    }

    public static void sql(String query) {
        // logic ceritanya buat execute SQL queries
    }
}
