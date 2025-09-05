//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        // to see how IntelliJ IDEA suggests fixing it.
        OrderService orderService = new OrderService();
        orderService.save("0001");

        OrderDetailService orderDetailService = new OrderDetailService();
        orderDetailService.save("0001", "Indomie");
        orderDetailService.save("0001", "Sabun");
        orderDetailService.save("0001", "Pepsodent");
    }
}