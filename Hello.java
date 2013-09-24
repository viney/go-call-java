public class Hello {
    public static void main(String[] args) {
        if (args.length == 0) {
            System.out.println("hello, world!");
        } else {
            System.out.printf("hello %s!", args[0]);
        }
    }
}
