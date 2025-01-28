package mylist

import "container/list"

type MyList list.List

func New(elements ...any) MyList {
	l := list.New()
	for e := range elements {
		l.PushBack(e)
	}
	return MyList(*l)
}
func (l *MyList) Filter(predicate func(any) bool) MyList {
	result := list.New()
	li := list.List(*l)
	e := li.Front()
	for e.Next() != nil {
		if predicate(e.Value) {

			result.PushFront(e)
		}
		e = e.Next()
	}
	return MyList(*result)
}
