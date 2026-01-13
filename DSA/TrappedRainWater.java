public class TrappedRainWater {


    static int trappedWater(int[] arr){

        int tw =0;
        int n = arr.length;
        
        // finding left Max
        int[] lMax = new int[n];
        lMax[0] = arr[0];
        for(int i =1; i<n; i++){

            lMax[i] = Math.max(lMax[i-1], arr[i]);
        }
        //finding right max
         int[] rMax = new int[n];
         rMax[n-1] = arr[n-1];
         for(int j = n-2; j>=0; j--){

            rMax[j] = Math.max(rMax[j+1], arr[j]);
        }
        // finding total trapped water

        for(int k=0; k<n; k++){
           int wl = Math.min(lMax[k], rMax[k]);
           tw += wl- arr[k];
        }

        return tw;

    }

        static int totalTrappedWater(int[] height) {

        int l = 0;
        int r = height.length -1;
        int lMax = 0, rMax = 0;
        int tw =0;

        while(l<r){
            if(height[l] <= height[r]){
                if (lMax > height[l]){
                    tw += lMax - height[l];
                }else{
                    lMax = height[l];
                }
                l++;
            }else{
                if (rMax> height[r]){
                    tw += rMax - height[r];
                }else{
                    rMax = height[r];
                }
                r--;
            }
        }
        return tw;
        
    }
    public static void main(String[] args) {
     
       int[] height = {0,1,0,2,1,0,1,3,2,1,2,1};
        // tc = 0(n) and sc = 0(n) 
        int total_trapped_Water = trappedWater(height);

        System.out.println("Total Trapped Water : "+total_trapped_Water);

        // Optimal approch - tc = 0(n) and sc = 0(1) 

         int trapped_Water = totalTrappedWater(height);
         System.out.println("Total Trapped Water : "+trapped_Water);
    }
    
}
