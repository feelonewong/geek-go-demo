package fib

import (
	"testing"
)

// 1 1 2 3 5 8 13
func TestFib(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)
	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		temp := a
		a = b
		b = temp + a
	}
	//	常量
	const (
		Monday    = iota + 1 // 默认为0， +1以后为1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)
	const (
		Open = 1
		Close
		Pending
	)
	t.Log("Monday:", Monday, "Tuesday:", Tuesday,
		"Wednesday:", Wednesday, "Thursday:", Thursday,
		"Friday:", Friday, "Saturday:", Saturday,
		"Sunday:", Sunday)
}
