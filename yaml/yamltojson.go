package main

import (
	"bytes"
	"fmt"
	"github.com/ghodss/yaml"
	"unicode"
)

func ToJSON(data []byte)([]byte,error){
	if hasJSONPrefix(data){
		return data,nil
	}
	return yaml.YAMLToJSON(data)
}


var jsonPrefix= []byte("{")

func hasJSONPrefix(buf []byte)bool{
	return hasPrefix(buf,jsonPrefix)
}

func hasPrefix(buf []byte,prefix []byte)bool{
	trim := bytes.TrimLeftFunc(buf,unicode.IsSpace)
	return bytes.HasPrefix(trim,prefix)

}

func main(){
	j := []byte(`{"name": "John", "age": 30}`)
	jsstr,_ := ToJSON(j)
	fmt.Println(jsstr)

	y, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("yaml:\n",string(y))
	/* Output:
	name: John
	age: 30
	*/
	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("JSON：\t",string(j2))
	/* Output:
	{"age":30,"name":"John"}
	*/
}