package html_test

import (
	"fmt"
	"html"
	"testing"
)

func Test(t *testing.T) {
	rawString := `<script>alert(1)</script>`
	safeString := html.EscapeString(rawString)
	fmt.Println(safeString)
}
