public class KadansAlgo{
    public static void main(String[] args) {
        
        int arr[] = {1,2,-2,-1,7,-3,5};
        
        int maxSum = Integer.MIN_VALUE;
        int curr = 0;

        
        for(int i = 0; i<arr.length; i++){

            curr = curr + arr[i];
            if(curr < 0){
                curr = 0;
            }
            maxSum = Math.max(maxSum, curr);
        }
        System.out.println("maximum sum : "+maxSum);
    }
}