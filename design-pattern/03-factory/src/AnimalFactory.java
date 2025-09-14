public class AnimalFactory {
    public Animal create(String type) {
        if (type.equalsIgnoreCase("Tiger")) {
            return new Tiger();
        }else if (type.equalsIgnoreCase("Dog")) {
            return new Dog();
        }else {
            return new Cat();
        }
    }
}
