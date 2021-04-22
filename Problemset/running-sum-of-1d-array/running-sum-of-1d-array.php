
// @Title: 一维数组的动态和 (Running Sum of 1d Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-14 21:27:58
// @Runtime: 44 ms
// @Memory: 14.7 MB

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function runningSum($nums) {
        $sum=0;
        $runningSum = array();
        for($i=0;$i<count($nums);$i++){
            for($j=0;$j<$i+1;$j++){
                $sum+=$nums[$j];
            }
            array_push($runningSum,$sum);
            $sum=0;
        }
        return $runningSum;
        

    }
}
