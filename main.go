package main

import (
	"fmt"
	"os"
	"strconv"
)

func dirMaker(number int, folderType string){
	dirNumber := strconv.Itoa(number)
		switch length:= len(dirNumber); length{
		case 1:
			dirNumber = "00000"+dirNumber
		case 2:
			dirNumber = "0000"+dirNumber
		case 3:
			dirNumber = "000"+dirNumber
		case 4:
			dirNumber = "00"+dirNumber
		case 5: 
			dirNumber = "0"+dirNumber
		}
	error := os.MkdirAll("filmPac/" + dirNumber + "/"+ dirNumber + folderType,0777)
		if error != nil {
			fmt.Println("Error Creating File: ",dirNumber, folderType)
			fmt.Print(error)
		}
}

func main(){
	for i:=1 ; i < 10001; i++ {
		dirMaker(i, "CF")
		dirMaker(i, "CH")
		dirMaker(i, "PD")
		dirMaker(i, "PF")
		dirMaker(i, "PI")
		dirMaker(i, "PS")
		dirMaker(i, "PT")
		dirMaker(i, "SL")
	}
	fmt.Println("finished")
}

