interface Animal {
    public void speak();
}

class Tiger implements Animal {
    public void speak() {
        System.out.println("Rawr!");
    }
}

class Dog implements Animal {
    public void speak() {
        System.out.println("Woof!");
    }
}

class Cat implements Animal {
    public void speak() {
        System.out.println("Meow!");
    }
}
