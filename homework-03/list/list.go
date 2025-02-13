// Реализация двусвязного списка вместе с базовыми операциями.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

/* 
	Elem - узел списка.
	Val - значение узла. (interface{} чтобы можно было хранить любые данные)
	next - указатель на следующий узел.
	prev - указатель на предыдущий узел.

	+-------+
	| root  |----<
	+-------+
*/ 
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

/* 
	New создаёт список и возвращает указатель на него. 
	l.root - корневой элемент списка.
	l.root.next - указатель на следующий элемент списка.
	l.root.prev - указатель на предыдуший элемент списка.
*/
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет узел 'e' в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
// Условие проверяет, что список не пустой
func (l *List) Pop() *List {
	if l.root.next == l.root {
		return l
	}

	oldHead := l.root.next
	l.root.next = oldHead.next
	oldHead.next.prev = l.root
	
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	if l.root.next == l.root {
		return l
	}

	current := l.root.next
	var prev *Elem 
	var next *Elem

	for current != l.root {
		next = current.next
		prev = current.prev
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}

	l.root.next = prev
	l.root.prev = prev.prev
	return l
}