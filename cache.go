package lru_cache

import (
	"container/list"
	"errors"
	"fmt"
)

/* LRUCache
data is a map to hold the elements
list to order the data in access order
capacity length of LRU cache
*/

type LRUCache struct {
	capacity int
	list     *list.List
	data     map[string]string
}

// Pair Data structure to hold the data in list
type Pair struct {
	key   string
	value string
}

// ILRUCache ...
type ILRUCache interface {
	Put(key string, value string)
	Get(key string) (string, error)
	Print()
}

func NewLRUCache(capacity int) ILRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		data:     make(map[string]string, capacity),
	}
}

// Put Puts a element in map as well to the list
func (c *LRUCache) Put(key string, value string) {
	if c.list.Len() == c.capacity {
		// remove the last element
		item := c.list.Back()
		c.list.Remove(item)
		delete(c.data, item.Value.(Pair).key)
	}

	c.data[key] = value
	c.list.PushFront(Pair{key: key, value: value})
}

// moveToTop Iterates the list and moves the associated key to the top of the list
func (c *LRUCache) moveToTop(key string) {
	for item := c.list.Front(); item != nil; item = item.Next() {
		if item.Value.(Pair).key == key {
			c.list.MoveToFront(item)
			break
		}
	}

}

// Get Gets a key from the map and moves the item to the top of list
func (c *LRUCache) Get(key string) (string, error) {
	if val, exists := c.data[key]; exists {
		c.moveToTop(val)

		return val, nil
	}

	return "", errors.New("key not found")
}

// Print prints the list items
func (c *LRUCache) Print() {
	for item := c.list.Front(); item != nil; item.Next() {
		fmt.Println("key", item.Value.(Pair).key)
		fmt.Println("value", item.Value.(Pair).value)
	}
}
