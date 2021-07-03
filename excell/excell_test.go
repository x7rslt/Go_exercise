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