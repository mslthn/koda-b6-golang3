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
//buat readme, jelaskan package apa aja, kegunaannya apa aja

func main() {
	nameList := []string{
		"Arya", "Anandita",
		"Bobo", "Emil",
		"Andi", "Doni",
		"Edo", "Bandi", "Cici",
		"Caca", "Budi", 
		"Danu", "Cucuk", "Edi",
		"Eli", "Chichi",
	}

	var name string

	fmt.Print("Search: ")
	fmt.Scan(&name)

	lib.SearchPerson(nameList, name)
}
