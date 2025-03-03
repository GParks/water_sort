package user.gparks;

/**
 * Hello world!
 */
public final class App {
    private App() {
    }

    public static void println(String s) {
        System.out.println(s);
    } 

    /**
     * Says hello to the world.
     * @param args The arguments of the program.
     */
    public static void main(String[] args) {
        // System.out.println("Hello World!");
        println("\t Green = " + ColorEnum.GREEN.toString());
        println("\t Green (by name) = " + ColorEnum.GREEN.name());
        println("\t Green (by ordinal) = " + ColorEnum.GREEN.ordinal());
        println("\t Green (by valueOf) = " + ColorEnum.valueOf("GREEN")); 
    }
}
