package goto_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoto(t *testing.T) {
Loop:
	for i := 0; i < 10; i++ {
		if i == 7 {
			time.Sleep(1 * time.Second)
			goto Loop
		}
		fmt.Println(i)

	}
}
