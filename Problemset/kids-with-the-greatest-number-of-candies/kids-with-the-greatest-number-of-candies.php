
// @Title: 拥有最多糖果的孩子 (Kids With the Greatest Number of Candies)
// @Author: 846188037@qq.com
// @Date: 2020-07-16 20:45:21
// @Runtime: 12 ms
// @Memory: 14.7 MB

class Solution {

    /**
     * @param Integer[] $candies
     * @param Integer $extraCandies
     * @return Boolean[]
     */
    function kidsWithCandies($candies, $extraCandies) {
        $max = $candies[0];

        for ($i = 1, $count = count($candies); $i < $count; $i++) {
            $max > $candies[$i] ?: $max = $candies[$i];
        }

        for ($i = 0; $i < $count; $i++) {
            $candies[$i] = $candies[$i] + $extraCandies >= $max ? true : false;
        }

        return $candies;
    }
}

