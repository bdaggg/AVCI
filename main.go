package main

import (
	"flag"
	"fmt"
	"os"
	project "project/avci"
)

func main() {
	project.Logo()

	var aranan string
	flag.StringVar(&aranan, "aranan", "null", "aranan ifadenin tutuldugu degiskendir")
	flag.StringVar(&aranan, "a", "null", "aranan ifadenin tutuldugu degiskendir")

	var dosya string
	flag.StringVar(&dosya, "dosya", "null", "içinde arama yapilacak dosya yolunu tutan degiskendir")
	flag.StringVar(&dosya, "d", "null", "içinde arama yapilacak dosya yolunu tutan degiskendir")

	var ifadeSayac bool
	flag.BoolVar(&ifadeSayac, "sayac", false, "yardimi tutan degiskendir")
	flag.BoolVar(&ifadeSayac, "s", false, "yardimi tutan degiskendir")

	var yardim bool
	flag.BoolVar(&yardim, "yardim", false, "yardimi tutan degiskendir")
	flag.BoolVar(&yardim, "y", false, "yardimi tutan degiskendir")

	flag.Parse()

	if yardim {
		project.Yardim()
	} else if ifadeSayac {
		err := project.KacDefa(aranan, dosya, ifadeSayac)
		if err != nil {
			fmt.Println("hata olustu!!")
			os.Exit(1)
		}
	} else {
		project.Kontrol(aranan, dosya)
	}

}
