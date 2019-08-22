import java.util.Scanner;

class Main {
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        String input = sc.next();
        input.split("");
        int answer = 0;
        for(String s : input.split("")) {
            answer += Integer.parseInt(s);
        }
        System.out.print(answer);
    }
}