package regexp_test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexBool(t *testing.T) {
	re := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`) //regex equation

	fmt.Println(re.MatchString("adam[23]")) //判断是否匹配
}

func TestMatch(t *testing.T) {
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	//matched, err = regexp.Match(`foo.*`, "seafood") //error:cannot use "seafood" (type string) as type []byte in argument to regexp.Match
	//fmt.Println(matched, err)
	matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`a\(b`, []byte(`seafood`))
	fmt.Println(matched, err)

}

func TestMatchstring(t *testing.T) {
	matched, err := regexp.MatchString(`hello`, "hello world")
	fmt.Println(matched, err)

}

func TestQuoteMeta(t *testing.T) {
	matched := regexp.QuoteMeta(`Escape regex expression:.+()[0-9]`) //提取正则表达式，并转义输出
	fmt.Println(matched)
}

func TestReagexpFind(t *testing.T) {
	re := regexp.MustCompile(`hello.*`)
	fmt.Printf("%q\n", re.Find([]byte(`hello world`)))

}

func TestReagexpFindAll(t *testing.T) {
	re := regexp.MustCompile(`h.?`)
	fmt.Printf("%q\n", re.FindAll([]byte(`hello world happy new year`), -1)) //第二个参数是指定所需的匹配计数。 -1表示不指定计数。-1全部输出，0，1只输出指定位

}

func TestRegexFindAllSubmatch(t *testing.T) {
	re := regexp.MustCompile(`h(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`hello world happy holiday`), -1))
}

func TestFindString(t *testing.T) {
	re := regexp.MustCompile(`h(.?)`)
	fmt.Println(re.FindString(`hello happy holiday`))
}

func TestFindStringIndex(t *testing.T) {
	re := regexp.MustCompile(`g([a-z]*)g`)
	fmt.Println(re.FindStringIndex(`I like code golang program.`)) //output[12 18] is mean find start 12 end 18.
}

func TestRegexReplace(t *testing.T) {
	re := regexp.MustCompile(`g([a-z]+)ng`)
	altered := re.ReplaceAllString(`Can golang be the prefect lanuage.`, "python")
	fmt.Println(altered)
}
