import java.util.Scanner;
    
class Main{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int a = Integer.parseInt(sc.next());
        int b = Integer.parseInt(sc.next());
        int mod = a * b % 2;
        String result = mod == 0? "Even" : "Odd";
        System.out.println(result);
    }
}