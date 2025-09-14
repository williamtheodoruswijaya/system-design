import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        ArrayList<CatalogAdapter> list = new ArrayList<>();

        list.add(new BookAdapter(new Book("Percy Jackson", "Rick Riordan")));
        list.add(new BookAdapter(new Book("Red Queen", "Victoria Aveyard")));
        list.add(new BookAdapter(new Book("Weathering with You", "Makoto Shinkai")));

        list.add(new ScreencastAdapter(new Screencast("Shadow of Bone", "Netflix", "02:00:00")));
        list.add(new ScreencastAdapter(new Screencast("Alice in Borderland", "Netflix", "02:30:00")));
        list.add(new ScreencastAdapter(new Screencast("Sandman", "Netflix", "03:30:00")));

        list.forEach(item -> {
            System.out.println(item.getCatalogTitle());
        });
    }
}