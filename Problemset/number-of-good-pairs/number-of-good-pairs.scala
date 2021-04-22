
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-13 01:59:02
// @Runtime: 532 ms
// @Memory: 49.3 MB

object Solution {
    def numIdenticalPairs(nums: Array[Int]): Int = {
        var count=0;
        var i=0;
        var j=0;
        for(i <- 0 to nums.length-2){
            for(j <- i+1 to nums.length-1){
                if(nums(i)==nums(j)){
                    count=count+1;
                }
            }
        }
        return count;

    }
}
