package main
import (
	"fmt"
)

type Node struct{
	data int
	next *Node
}

type sll struct{
	head *Node
}

func (l *sll) add(dt int) {
	n := &Node{data: dt, next: nil}
	if l.head == nil {
		l.head = n
	}else{
		ptr := l.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = n
	}
}

func (l *sll) remb(){
	if l.head == nil {
		fmt.Println("nothing to delete")
	}else{
		if l.head.next == nil {
			l.head = nil
		}else{
			ptr := l.head
			for ptr.next.next != nil {
				ptr = ptr.next
			}
			ptr.next = nil
		}
	}
}

func (l sll) showall(){
	ptr := l.head
	for ptr != nil {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}

}

func main(){
	l := sll{head: nil}
	l.add(46)
	l.add(35)
	l.remb()
	l.remb()
	l.remb()
	l.showall()
}
