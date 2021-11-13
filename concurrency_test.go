package l6hw

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	// 7 functions in array that will return an error
	var testArray = []func(ctx context.Context) error{
		createFunc(1, true), createFunc(1, false), createFunc(6, false), createFunc(11, false),
		createFunc(3, true), createFunc(4, false), createFunc(7, false), createFunc(12, false),
		createFunc(2, true), createFunc(8, false), createFunc(8, false), createFunc(13, false),
		createFunc(6, true), createFunc(5, true), createFunc(4, true), createFunc(6, true),
	}

	if err := Execute(testArray, 4); err == nil {
		t.Errorf("Execute funct should have returned an error, but it didn't")
	}
	if err := Execute(testArray, 7); err == nil {
		t.Errorf("Execute funct should have returned an error, but it didn't")
	}
	if err := Execute(testArray, 10); err != nil {
		t.Errorf("Execute funct shouldn't have returned an error, but it did")
	}


}

func createFunc(sec int, is bool) func(ctx context.Context) error {

	return func(ctx context.Context) error {
		time.Sleep(time.Millisecond * time.Duration(sec)*100)
		select {
		case <-ctx.Done():
			fmt.Println(sec, " this task canceled")
			return nil
		default:
			fmt.Println(sec, " this func end, returns error == ", is)
			if is {
				return errors.New("error occurred")
			}
			return nil
		}
	}
}


func main() {
	var arr = []func(ctx context.Context) error{
		createFunc(1, true), createFunc(1, false), createFunc(6, false), createFunc(11, false),
		createFunc(3, true), createFunc(4, false), createFunc(7, false), createFunc(12, false),
		createFunc(2, true), createFunc(8, false), createFunc(8, false), createFunc(13, false),
		createFunc(6, true), createFunc(10, true), createFunc(9, true), createFunc(14, true),
	}
	fmt.Println(Execute(arr, 4))
}