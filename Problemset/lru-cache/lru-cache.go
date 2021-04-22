
// @Title: LRU 缓存机制 (LRU Cache)
// @Author: 846188037@qq.com
// @Date: 2021-04-18 20:29:24
// @Runtime: 156 ms
// @Memory: 12.1 MB

type LRUCache struct {
    size int
    capacity int 
    cache map[int]*DLinkedList
    head *DLinkedList
    tail *DLinkedList

}

type DLinkedList struct{
    key int 
    value int 
    prev *DLinkedList
    next *DLinkedList
}

func initDLinkedList(key,value int)*DLinkedList{
    return &DLinkedList{
        key:key,
        value:value,
    }
}


func Constructor(capacity int) LRUCache {
    l:=LRUCache{
        capacity:capacity,
        cache:make(map[int]*DLinkedList),
        head:initDLinkedList(0,0),
        tail:initDLinkedList(0,0),

    }
    l.head.next=l.tail
    l.tail.prev=l.head
    return l

}


func (this *LRUCache) Get(key int) int {
    if _,ok:=this.cache[key];!ok{
        return -1
    }
    node:=this.cache[key]


    this.moveToTail(node)
    //fmt.Println(this.cache)
    return node.value

}

func (this *LRUCache) moveToTail(node *DLinkedList){
    node.prev.next=node.next
    node.next.prev=node.prev

    node.next=this.tail
    node.prev=this.tail.prev
    node.prev.next=node
    this.tail.prev=node
}


func (this *LRUCache) Put(key int, value int)  {
    _,ok:=this.cache[key]
    if ok{
        this.cache[key].value=value
        node:=this.cache[key]

    this.moveToTail(node)

    }else{
        node:=initDLinkedList(key,value)
        this.cache[key]=node
        this.tail.prev.next=node
        node.prev=this.tail.prev
        node.next=this.tail
        this.tail.prev=node
        this.size++
        for this.size>this.capacity{
            delete(this.cache,this.head.next.key)
            this.size--
            this.head.next=this.head.next.next
            this.head.next.prev=this.head
        }
    }
    //fmt.Println(this.cache)

}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
