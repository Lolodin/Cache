package Cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Cache struct {
	sync.RWMutex
	item []Item
}

type Item struct {
	Data []byte
}
func NewCache() *Cache {
	item := make([]Item,1)
	c:= Cache{item:item}

	return &c

}
func (c *Cache) GetItem() (Item, error) {

	count:=len(c.item)
	fmt.Println(count)
	if count == 1 {
		return Item{nil},fmt.Errorf("Cache is Empty")

	} else {
		c.Lock()
		defer c.Unlock()
		el:=c.item[count-1]
		c.item = c.item[:count-1]

		return el, nil
	}




}
func (c *Cache) PushItem(s []byte) {

item:= Item{Data:s}
	c.Lock()
	defer c.Unlock()
c.item = append(c.item, item)
}
func (c *Cache) GenerateItem() {
	randomValue:= rand.Int()
	b:=[]byte(strconv.Itoa(randomValue))
	c.PushItem(b)

}