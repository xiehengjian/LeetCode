
// @Title: 重新排列数组 (Shuffle the Array)
// @Author: 846188037@qq.com
// @Date: 2020-07-15 20:13:52
// @Runtime: 16 ms
// @Memory: 14.4 MB

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $n
     * @return Integer[]
     */
    function shuffle($nums, $n) {
        $newNums=array();
        for($i=0;$i<$n;$i++){
            array_push($newNums,$nums[$i],$nums[$n+$i]);
        }
        return $newNums;

    }
}
