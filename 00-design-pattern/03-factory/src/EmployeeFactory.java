public class EmployeeFactory {
    public Employee createManager(String name) {
        return new Employee(name, "Manager", 10000000);
    }

    public Employee createStaff(String name) {
        return new Employee(name, "Staff", 5000000);
    }
}
