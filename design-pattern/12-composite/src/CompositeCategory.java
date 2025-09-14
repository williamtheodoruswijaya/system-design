import java.util.ArrayList;

public class CompositeCategory extends Category {
    private ArrayList<Category> list = new ArrayList<>();

    public CompositeCategory(String name) {
        super(name);
    }

    // Karena kita pakai ArrayList, wajib punya method buat add and remove ke list
    public void add(Category item) {
        list.add(item);
    }

    public void remove(Category item) {
        list.remove(item);
    }

    public ArrayList<Category> getList() {
        return list;
    }

    public void setList(ArrayList<Category> list) {
        this.list = list;
    }

    @Override
    public void display(String indent) {
        System.out.println(indent + "+ " + this.getName());
        for (Category child : list) {
            child.display(indent + "  ");
        }
    }
}
