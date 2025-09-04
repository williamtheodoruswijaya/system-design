public class CustomerBuilder {
    // step 1: create the same attributes on Customer Classes
    private int id;
    private String firstName;
    private String lastName;
    private String email;
    private String phone;

    // atribut tambahan (default value-nya kita set disini)
    private String address = "";
    private int age = 0;

    // step 2: create a setter function for each attributes and changes void to return CustomerBuilder class (this must in order for the builder function to work)
    public CustomerBuilder setId(int id) {
        this.id = id;
        return this;
    }

    public CustomerBuilder setFirstName(String firstName) {
        this.firstName = firstName;
        return this;
    }

    public CustomerBuilder setLastName(String lastName) {
        this.lastName = lastName;
        return this;
    }

    public CustomerBuilder setEmail(String email) {
        this.email = email;
        return this;
    }

    public CustomerBuilder setPhone(String phone) {
        this.phone = phone;
        return this;
    }

    public CustomerBuilder setAddress(String address) {
        this.address = address;
        return this;
    }

    public CustomerBuilder setAge(int age) {
        this.age = age;
        return this;
    }

    // step 3: create a build function (we will use this function instead of using Customer constructor)
    public Customer build() {
        return new Customer(
                this.id,
                this.firstName,
                this.lastName,
                this.email,
                this.phone,
                this.address,
                this.age
        );
    }
}
