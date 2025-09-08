public class Category {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Category(String name) {
        this.name = name;
    }

    public void display(String indent) {
        System.out.println(indent + "- " + name);
    }
}
