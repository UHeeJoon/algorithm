package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (q *Stack) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Stack) Pop() interface{} {
	back := q.v.Back()
	if back == nil {
		return nil
	}

	return q.v.Remove(back)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for i := 0; i < tc; i++ {
		var (
			str    string
			result string
			stack  *Stack = NewStack()
		)
		fmt.Fscan(reader, &str)
		result = "YES"
		for _, data := range []rune(str) {
			if data == '(' {
				stack.Push(data)
			} else {
				if stack.v.Len() == 0 {
					result = "NO"
					break
				}
				stack.Pop()
			}
		}
		if stack.v.Len() != 0 {
			result = "NO"
		}
		fmt.Fprintln(writer, result)
	}

}
