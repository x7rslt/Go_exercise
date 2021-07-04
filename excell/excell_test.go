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