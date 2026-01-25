public class MissingNumber {
    
    public static void main(String[] args) {
         int nums[] = {1,2,3,5};

        int n = nums.length;
      
        int xOr = 0;
        for (int i =1; i<= n; i++){
            xOr ^= i;
        }

        for (int i =0; i< n; i++){
            xOr ^= nums[i];
        }

        System.out.println(xOr);
        
    

    }
}
