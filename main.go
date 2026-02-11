package main

import (
	"fmt"
	"koda-b6-golang3/lib"
)

// pencarian nama dg melakukan passing data slice melalui parameter pertama dan
// keyword pencarian melalui parameter kedua
// kembalikan slice hasil pencarian, jika tidak ada, kembalikan slice kosong
// contoh searchPerson(users, "John")
// buat interaktif

func main() {
	nameList := []string{
		"Arya", "Anandita", 
		"Bobo", "Budi",
		"Caca", "Chichi",
		"Danu", "Doni",
		"Eli", "Edo",
	}
	
	var name string

	fmt.Print("Search: ")
	fmt.Scan(&name)

	lib.SearchPerson(nameList, name)
}