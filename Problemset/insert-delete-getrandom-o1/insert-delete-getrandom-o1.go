
// @Title: 常数时间插入、删除和获取随机元素 (Insert Delete GetRandom O(1))
// @Author: 846188037@qq.com
// @Date: 2021-01-27 16:17:35
// @Runtime: 36 ms
// @Memory: 7.9 MB

type RandomizedSet struct {
    rMap map[int]int
    rSlice []int

}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        rMap:make(map[int]int),
        rSlice:make([]int,0),
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    _,ok:=this.rMap[val]
    if !ok{
        this.rSlice=append(this.rSlice,val)
        this.rMap[val]=len(this.rSlice)-1
        return true
    }
    return false
    //return this.rSlice

}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    _,ok:=this.rMap[val]
    if ok{
        //将最后一个元素的值存放在被删除的那个人的位置上
        this.rSlice[this.rMap[val]]=this.rSlice[len(this.rSlice)-1]
        
        this.rMap[this.rSlice[len(this.rSlice)-1]]=this.rMap[val]
        this.rSlice=this.rSlice[:len(this.rSlice)-1]
        delete(this.rMap,val)

        return true
    }
    return false
    //return this.rSlice

}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    num:=rand.Intn(len(this.rSlice))//取一个随机数
    return this.rSlice[num]
    //return this.rSlice
    


}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
