
// @Title: 好数对的数目 (Number of Good Pairs)
// @Author: 846188037@qq.com
// @Date: 2020-07-12 21:15:36
// @Runtime: 0 ms
// @Memory: 5.5 MB

int numIdenticalPairs(int* nums, int numsSize){
    int count=0;
    for(int i=0;i<numsSize-1;i++){
        for(int j=i+1;j<numsSize;j++){
            if(nums[i]==nums[j]){
                count++;
            }
        }
    }
    return count;

}
