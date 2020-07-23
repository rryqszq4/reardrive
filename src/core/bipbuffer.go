package core

import (
	"fmt"
)

type bipbuf_t struct {

	size 		uint32

	/* region A */
	a_start 	uint32

	a_end 		uint32

	/* region B */
	b_end 		uint32

	/* is B inuse? */
	b_inuse 	bool

	data 		[]byte

}

func (self *bipbuf_t) __check_for_switch_to_b() {
	if self.size - self.a_end < self.a_start - self.b_end {
		self.b_inuse = true
	}
}

func (self *bipbuf_t) push() {

}

func (self *bipbuf_t) used() uint32 {
	return (self.a_end - self.a_start) + self.b_end
}

func (self *bipbuf_t) unused() uint32 {
	if self.b_inuse {
		return self.a_start - self.b_end
	}else {
		return self.size - self.a_end
	}
}

/*
====================
NewBipBuffer


====================
 */
func NewBipBuffer(size uint32) *bipbuf_t {
	return &bipbuf_t{
		size:size,
		a_start:0,
		a_end:0,
		b_end:0,
		b_inuse:false,
		data:make([]byte, size, size),
	}
}

func (self *bipbuf_t) free() {
	self.data = self.data[:0]
	self.size = 0
	self.a_start = 0
	self.a_end = 0
	self.b_end = 0
	self.b_inuse = false
}

/*
====================
Size


====================
 */
func (self *bipbuf_t) Size() uint32{
	return self.size
}

/*
====================
Offer


====================
 */
func (self *bipbuf_t) Offer(data []byte) uint32{
	var size = uint32(len(data))

	if self.unused() < size {
		return 0
	}

	if self.b_inuse {
		self.data = append(
				append(self.data[0:self.b_end], data...),
				self.data[size:]...
			)
		self.b_end += uint32(size)
	}else {
		self.data = append(
				append(self.data[0:self.a_end], data...),
				self.data[size:]...
			)
		self.a_end += uint32(size)
	}

	self.__check_for_switch_to_b()

	return size
}

/*
====================
Peek


====================
 */
func (self *bipbuf_t) Peek(size uint32) []byte{
	if self.size < self.a_start + size {
		return nil
	}

	if self.IsEmpty() {
		return nil
	}

	return self.data[0:self.a_start]
}

/*
====================
Poll


====================
 */
func (self *bipbuf_t) Poll(size uint32) []byte {
	if self.IsEmpty() {
		return nil
	}

	if self.size < self.a_start + size {
		return nil
	}

	end := self.data[self.a_start:self.a_start+size]

	self.a_start += size

	if self.a_start == self.a_end {
		if self.b_inuse {
			self.a_start = 0
			self.a_end = self.b_end
			self.b_end = 0
			self.b_inuse = false
		}else {
			self.a_start = 0
			self.a_end = 0
		}
	}

	self.__check_for_switch_to_b()

	return end
}

/*
====================
IsEmpty


====================
 */
func (self *bipbuf_t) IsEmpty() bool {
	return self.a_start >= self.a_end
}

/*
====================
Print


====================
 */
func (self *bipbuf_t) Print() {
	fmt.Println(string(self.data), self.size, self.a_start, self.a_end, self.b_end, self.b_inuse)
}