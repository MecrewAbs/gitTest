package main

import (
	"fmt"
	"gitTest/bins"
	"gitTest/storage"
)

//ggg
func main() {
	store := storage.NewStorage()

	bin := bins.NewBin("mybin", false)
	fmt.Println("Bin was created", bin)

	bin.SetPrivate()
	fmt.Println("Bin was set private", bin)

	err := store.CreateBin(bin)
	if err != nil {
		fmt.Println("Saving error", err)
		return
	}
	retrivedBin, err := store.LookUpBun("mybin")
	if err != nil {
		fmt.Println("LookUp error", err)
		return
	}
	fmt.Println("LookUp Bun was returned", retrivedBin)

	jsonStr, err := retrivedBin.ToJSON()
	if err != nil {
		fmt.Println("JSON error", err)
		return
	}
	fmt.Println("JSON was returned", jsonStr)

	newBin := &bins.Bin{}
	err = newBin.FromJSON(jsonStr)
	if err != nil {
		fmt.Println("JSON error", err)
		return
	}
	fmt.Println("JSON was returned", newBin)
	fmt.Println("Bin list in the storage", store.ListBinNames())
}
