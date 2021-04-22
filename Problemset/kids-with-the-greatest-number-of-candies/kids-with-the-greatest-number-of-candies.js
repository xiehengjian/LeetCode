
// @Title: 拥有最多糖果的孩子 (Kids With the Greatest Number of Candies)
// @Author: 846188037@qq.com
// @Date: 2020-07-16 20:43:28
// @Runtime: 88 ms
// @Memory: 37 MB

const kidsWithCandies = (candies, extraCandies) => {
  let max = 0
  for (const candie of candies) { // 找到数组中最大值
    max = Math.max(max, candie)
  }
  for (let i = 0; i < candies.length; i++) { // 再循环一遍数组
    candies[i] = candies[i] + extraCandies >= max
  }
  return candies
};

