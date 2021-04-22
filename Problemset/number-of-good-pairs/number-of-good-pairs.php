
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-13 00:06:00
// @Runtime: 12 ms
// @Memory: 14.3 MB

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function numIdenticalPairs($nums) {
        $count=0;
        for($i=0;$i<count($nums)-1;$i++){
            for($j=$i+1;$j<count($nums);$j++){
                if($nums[$i]==$nums[$j]){
                    $count++;
                }
            }
        }
        return $count;

    }
}
