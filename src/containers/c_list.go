package containers

import (
	"fmt"
)

type CListItem struct {

	data 	interface{}

	next 	*CListItem

}

type CList struct {

	size 		int

	head 		*CListItem

	tail 		*CListItem

}

func NewCListItem(v interface{}) *CListItem {
	return &CListItem{
		data : v,
		next : nil,
	}
}

func NewCList() *CList {
	return new(CList).Init()
}

func (self *CList) Init() *CList{
	self.size = 0

	self.head = nil

	self.tail = nil

	return self
}

func (self *CList) Size() int  {
	return self.size
}

func (self *CList) Head() *CListItem {
	return self.head
}

func (self *CList) Tail() *CListItem {
	return self.tail
}

func (self *CList) IsHead(item *CListItem) bool {
	if item == self.head {
		return true
	}

	return false
}

func (self *CList) IsTail(item *CListItem) bool {
	if self.tail == nil {
		return true
	}

	return false
}

func (self *CList) Data(item *CListItem) interface{} {
	return item.data
}

func (self *CList) Next(item *CListItem) *CListItem {
	return item.next
}

func (self *CList) NextInsert(item *CListItem, v interface{}) int {
	var newItem *CListItem

	newItem = NewCListItem(v)

	if item == nil {
		if self.size == 0 {
			self.tail = newItem
		}
		newItem.next = self.head
		self.head = newItem
	}else {
		if item.next == nil {
			self.tail = newItem
		}
		newItem.next = item.next
		item.next = newItem
	}

	self.size++

	return 0
}

func (self *CList) NextRemove(item *CListItem, data *interface{}) (int, *interface{}) {

	if self.size == 0 {
		return -1,nil
	}

	if item == nil {
		*data = self.head.data
		//reflect.ValueOf(data).Elem().Set(reflect.ValueOf(self.head.data))
		self.head = self.head.next

		if self.size == 1 {
			self.tail = nil
		}
	}else {
		if item.next == nil {
			return -1, nil
		}

		*data = item.next.data
		//reflect.ValueOf(data).Elem().Set(reflect.ValueOf(item.next.data))
		item.next = item.next.next

		if item.next == nil {
			self.tail = item
		}
	}

	self.size--

	return 0, data
}

func (self *CList) Print() {
	var item *CListItem
	var i int

	fmt.Printf("List size is %d, head is %d, tail is %d\n", self.Size(), self.Data(self.Head()), self.Data(self.Head()))

	i = 0
	item = self.Head()

	for {
		if item == nil {
			break
		}


		data := self.Data(item)
		fmt.Printf("list[%d]=%d\n", i, data)

		i++

		if self.IsTail(item) {
			break
		} else {
			item = self.Next(item)
		}
	}

	return
}