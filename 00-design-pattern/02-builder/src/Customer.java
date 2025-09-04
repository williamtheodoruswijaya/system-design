public class Customer {
    private int id;
    private String firstName;
    private String lastName;
    private String email;
    private String phone;

    // this two are the additional variables
    private String address;
    private int age;

    /*
    - to prevent error on constructor across our apps, we'll use overload function where one was the previous constructor
    - and one as the new constructor with new attributes as the parameter
    - the drawback of this pattern is that we need to create a new overload constructor for any combination of assignable attributes where there are some nullable attributes
    */
    public Customer(int id, String firstName, String lastName, String email, String phone) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
        this.phone = phone;
        this.address = "";
        this.age = 0;
    }

    public Customer(int id, String firstName, String lastName, String email, String phone, String address, int age) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
        this.phone = phone;
        this.address = address;
        this.age = age;
    }
}
