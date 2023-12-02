package project

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func KacDefa(anahtarIfade, dosyaYoli string, s bool) error {
	dosya, err := os.Open(dosyaYoli)
	if err != nil {
		return err
	}
	defer dosya.Close()

	scanner := bufio.NewScanner(dosya)
	butunKontrol := 0
	for scanner.Scan() {
		metin := scanner.Text()
		varmi, err := regexp.MatchString(anahtarIfade, metin)
		if err != nil {
			return err
		}
		if varmi {
			butunKontrol += 1
		}
	}
	if butunKontrol == 0 {
		fmt.Println("aranan ifade bulunumadi")
	}else{
		fmt.Printf("aranan ifade toplam %d defa dosya icinde gecmektedir",butunKontrol)
	}


	return nil
}
