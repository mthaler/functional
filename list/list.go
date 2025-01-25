package mylist

import "container/list"

type MyList list.List

func New(elements ...any) MyList {
	l := list.New()
	for e := range elements {
		l.PushBack(e)
	}
	return *l
}
func (l *MyList) Filter(predicate func(any) bool) MyList {
	result := list.New()
	e := l.Front()
	for e.Next() != nil {
		if predicate(e.Value) {
			result.PushFront(e)
		}
		e = e.Next()
	}
	return MyList(list.New(result))
}
