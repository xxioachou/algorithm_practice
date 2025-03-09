package main

import (
	"container/list"
)

type TextEditor struct {
	text 		*list.List		// 文本内容，头尾都有一个哨兵
	head, tail 	*list.Element	// 头尾哨兵
	curr		*list.Element	// cursor 在 curr 的右边
}

func Constructor() TextEditor {
	t := TextEditor{}
	t.text = list.New()
	t.head = t.text.PushFront(byte('$'))
	t.tail = t.text.PushBack(byte('#'))
	t.curr = t.head
	return t
}

func Reverse(a []byte) {
	for i := 0; i < len(a) / 2; i ++ {
		a[i], a[len(a) - 1 - i] = a[len(a) - 1 - i], a[i]
	}
}

func (t *TextEditor) LeftContent() string {
	res := make([]byte, 0)
	curr := t.curr
	for i := 0; i < 10 && curr != t.head; i ++ {
		res = append(res, curr.Value.(byte))
		curr = curr.Prev()
	}
	Reverse(res)

	return string(res)
}

func (t *TextEditor) AddText(text string)  {
    for _, ch := range text {
		ne := t.text.InsertAfter(byte(ch), t.curr)
		t.curr = ne
	}
}


func (t *TextEditor) DeleteText(k int) int {
	cnt := 0
    for i := 0; i < k && t.curr != t.head; i ++ {
		tmp := t.curr
		t.curr = t.curr.Prev()
		t.text.Remove(tmp)
		cnt ++
	}
	return cnt
}


func (t *TextEditor) CursorLeft(k int) string {
    for i := 0; i < k && t.curr != t.head; i ++ {
		t.curr = t.curr.Prev()
	}

	return t.LeftContent()
}


func (t *TextEditor) CursorRight(k int) string {
    for i := 0; i < k && t.curr.Next() != t.tail; i ++ {
		t.curr = t.curr.Next()
	}

	return t.LeftContent()
}