package mylist

import (
	"container/list"
)

type MyList list.List

func New(elements ...any) MyList {
	l := list.New()
	for _, e := range elements {
		l.PushBack(e)
	}
	return MyList(*l)
}
func (l *MyList) Filter(predicate func(any) bool) MyList {
	result := list.New()
	li := list.List(*l)
	for e := li.Front(); e != nil; e = e.Next() {
		v := e.Value
		if predicate(v) {
			result.PushBack(v)
		}
	}
	return MyList(*result)
}

func (l *MyList) Map(f func(item any) any) MyList {
	result := list.New()
	li := list.List(*l)
	for e := li.Front(); e != nil; e = e.Next() {
		v := f(e.Value)
		result.PushBack(v)

	}
	return MyList(*result)
}
