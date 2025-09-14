public class Store {
    private String name;
    private String city;
    private String country;
    private String category;

    public Store(String name, String city, String country, String category) {
        this.name = name;
        this.city = city;
        this.country = country;
        this.category = category;
    }

    public String getName() {
        return name;
    }

    public String getCity() {
        return city;
    }

    public String getCountry() {
        return country;
    }

    public String getCategory() {
        return category;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setCity(String city) {
        this.city = city;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    public Store clone() {
        return new Store(
                this.name,
                this.city,
                this.country,
                this.category
        );
    }
}
