public class BuyandSell {


    public static void main(String[] args) {
        int prices[] = {7,1,5,3,6,4,5};

    // 7 = 0th day

    int buyPrice = prices[0];

    int maxProfit = 0 ;

    for ( int i = 0 ; i < prices.length; i++){

        if(buyPrice < prices[i]){
            int currProfit = prices[i] - buyPrice;
            maxProfit = Math.max(currProfit,maxProfit);
        }else {
            buyPrice = prices[i];
        }

    }

       System.out.println(maxProfit);

    }




    
}
