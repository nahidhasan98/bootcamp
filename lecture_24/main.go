package main

import (
	"fmt"
	"os"
)

func main() {
	// //getting working directory
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(dir)

	//creating a file
	isErr := createFile("master_academy3.txt", "Testing file creation")
	fmt.Println(isErr)

	// //status of a file
	// fi, err := os.Stat("master_academy3.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(fi.IsDir())
	// fmt.Println(fi.ModTime().Date())
	// fmt.Println(fi.Name())
	// fmt.Println(fi.Size())

	// //creating a folder
	// err := os.Mkdir("test_folder", 0777)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// //rename a file
	// os.Rename(`D:\TEST`, `D:\TEST_01`)

	// base := filepath.Base(dir)
	// relativePath := filepath.Join("master_academy")
	// absolutePath, _ := filepath.Abs("master_academy")

	// newPath := filepath.Join(absolutePath, "..", "..", "lecture_25")
	// fmt.Println(base)
	// fmt.Println(relativePath)
	// fmt.Println(absolutePath)
	// fmt.Println(newPath)

	// //working with excel file
	//  f := excelize.NewFile()
	// // Create a new sheet.
	// index := f.NewSheet("Sheet2")
	// // Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// // Save spreadsheet by the given path.
	// err := f.SaveAs("test.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// f, err := excelize.OpenFile("test.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Get value from cell by given worksheet name and axis.
	// cell := f.GetCellValue("Sheet2", "A2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)

	// // Get all the rows in the Sheet1.
	// rows, err := f.GetRows("Sheet1")
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }
}

func createFile(fileName, content string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer posf.Close()
	_, err = posf.Write([]byte(content))
	//fmt.Println(n, err)

	return err == nil
}
