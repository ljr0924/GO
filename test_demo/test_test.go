package test_demo

import (
	"context"
	"fmt"
	"testing"
)

func Add(i, j int) int {
	return i + j
}

func TestAdd(t *testing.T) {
	cases := []struct{
		Name string
		A, B, Answer int
	}{
		{"Add(1, 2)", 1, 2, 3},
		{"Add(3, 4)", 3, 4, 5},
		{"Add(5, 5)", 5, 5, 10},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d + %d", c.A, c.B), func(t *testing.T) {
			if answer := Add(c.A, c.B); answer != c.Answer {
				t.Fatalf("%d + %d got %d expect %d", c.A, c.B, answer, c.Answer)
			}
		})
	}
}

type Mgr struct {
	cancel context.CancelFunc
}

var GoroutineMgr map[string]*Mgr

func RunGoroutine(ctx context.Context, taskChan chan int) {
	for {
		select {
		case task := <-taskChan:
			fmt.Println("get task ", task)
		case <-ctx.Done():
			fmt.Println("goroutine exit")
			return
		}
	}
}

func TestGoroutineMgr(t *testing.T) {




}
