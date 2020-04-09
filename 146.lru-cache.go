/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (30.53%)
 * Likes:    4882
 * Dislikes: 216
 * Total Accepted:    460.1K
 * Total Submissions: 1.5M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 *
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 *
 * The cache is initialized with a positive capacity.
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 * Example:
 *
 *
 * LRUCache cache = new LRUCache( 2 capacity)
 */

// @lc code=start
type Node struct {
	Val  int
	Key  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	Cap  int
	Map  map[int]*Node
	Head *Node
	Tail *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Cap:  capacity,
		Map:  make(map[int]*Node, capacity),
		Head: &Node{},
		Tail: &Node{},
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Pre = cache.Head
	return cache

}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.putHead(node)

	return node.Val

}
func (this *LRUCache) remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}
func (this *LRUCache) putHead(node *Node) {
	originNext := this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head
	node.Next = originNext
	originNext.Pre = node

}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]
	if ok {
		node.Val = value
		this.remove(node)
		this.putHead(node)
	} else {
		if len(this.Map) >= this.Cap {
			deleteKey := this.Tail.Pre.Key
			this.remove(this.Tail.Pre)
			delete(this.Map, deleteKey)
		}
		newNode := &Node{Key: key, Val: value}
		this.putHead(newNode)
		this.Map[key] = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

