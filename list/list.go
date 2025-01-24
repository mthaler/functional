package mylist

import "container/list"

type MyList list.List

func (l *MyList) Filter(predicate func(any) bool) MyList {
	result := list.New()
	e := l.Front()
	for e.Next() != nil {
		if predicate(e.Value) {
			result.PushFront(e)
		}
		e = e.Next()
	}
	return New[T](result)
}
