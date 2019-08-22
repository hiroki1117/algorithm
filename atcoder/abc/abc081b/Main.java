import java.util.Scanner;

class Main {
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int n = Integer.parseInt(sc.next());
        sc.nextLine();
        String[] input = sc.nextLine().split(" ");
        int[] arr = new int[n];
        for (int i=0; i<n; i++) {
            arr[i] = Integer.parseInt(input[i]);
        }
        int count = 0;
        boolean flg = true;
        while(flg){
            for(int i=0;i<n; i++){
                if(arr[i]%2!=0) {
                    flg = false;
                    break;
                }
                arr[i] = arr[i]/2;
            }
            
            if(flg) count ++;
        }
        System.out.print(count);
    }
}