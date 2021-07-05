package excell_test

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"testing"
)

//创建xlsx
func TestExcell(t *testing.T){
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet2")
	xlsx.SetCellValue("Sheet2","A2","Hello world")
	xlsx.SetCellValue("Sheet1","B2",100)
	xlsx.SetActiveSheet(index)
	err := xlsx.SaveAs("./Excell.xlsx")
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

//读xlsx
func TestExcellRead(t *testing.T){
	f,err := excelize.OpenFile("Excell.xlsx")
	if err != nil{
		fmt.Println(err)
		return

	}
	cell:= f.GetCellValue("Sheet1","B2")
	fmt.Println(cell)
	rows := f.GetRows("Sheet1")
	for _,row := range rows{
		for _,colCell := range row{
			fmt.Println(colCell,"\t")

		}
		fmt.Println()
	}
}

func TestExcellAddChart(t *testing.T){
	categories := map[string]string{
		"A2":"Small","A3":"Normal","A4":"Large",
		"B1":"Apple","C1":"Orange","D1":"Pear",
	}
	values:= map[string]int{
		"B2":2,"C2":3,"D2":3,"B3":5,"C3":2,"D3":4,"B4":6,"C4":7,"D4":8}
	f:= excelize.NewFile()
	for k,v := range categories{
		f.SetCellValue("Sheet1",k,v)

	}
	for k,v := range values{
		f.SetCellValue("Sheet1",k,v)


	}
	if err := f.AddChart("Sheet1", "E1", `{
		"type":"col3DClustered",
		"series":[
		{
		"name": "Sheet1!$A$2",
		"categories": "Sheet1!$B$1:$D$1",
		"values": "Sheet1!$B$2:$D$2"
		},
		{
		"name": "Sheet1!$A$3",
		"categories": "Sheet1!$B$1:$D$1",
		"values": "Sheet1!$B$3:$D$3"
		},
		{
		"name": "Sheet1!$A$4",
		"categories": "Sheet1!$B$1:$D$1",
		"values": "Sheet1!$B$4:$D$4"
		}],
		"title":{
			"name":"Fruit 3D Clustered Column Chart"
	}}`);err!=nil{
		fmt.Println(err)
		return
	}
	if err := f.SaveAs("ChartExcell.xlsx");err!=nil{
		fmt.Println(err)
	}
}