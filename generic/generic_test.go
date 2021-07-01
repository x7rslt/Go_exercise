package generic_test

import (
	"fmt"
	"reflect"
	"testing"
)
//泛型示例1：
type GeneralKey interface {
	GetKeyString() string
}

type Akey struct{
	Id int `default:0`
	Type int `default:0`
	Date int `default:0`
}

type BKey struct {
	Id int `default:0`
	Hash int `default:0`
}

func (key Akey)GetKeyString()string{
	return fmt.Sprintf("%d-%d-%d",key.Id,key.Type,key.Date)

}
func (key BKey)GetKeyString()string{
	return fmt.Sprintf("%d",key.Id&key.Hash)
}
func work(key GeneralKey){
	keyStr := key.GetKeyString()
	fmt.Println(keyStr)
}

func TestGeneric1(t *testing.T){
	ak := BKey{12,34}
	work(ak)
}

//泛型示例2：
func TInterface(v interface{})interface{}{
	switch v.(type){
	case int:
		return v.(int)+10
	case float64:
		return v.(float64)+22.3
	}
	return v
}

func TestGeneric2(t *testing.T){
	t1 := TInterface(10)
	t2 := TInterface(10.0)
	fmt.Println(t1,reflect.TypeOf(t1).String())
	fmt.Println(t2,reflect.TypeOf(t2).String())


}