package project

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Grep adinda bir fonksiyon yarattim, aranacak kelimeyi ve dosyanin yoluni verdim
func Kontrol(keyWords, filePath string) error {
	//dosya diye bir degisken yarattim ve os.Open ile dosyami actim
	dosya, err := os.Open(filePath)
	//gerekli hata kontrollerini yaptim
	if err != nil {
		return err
	}
	//isimiz bittikten sonra dosyamizin kapanmasini garantiye alalim defer ile
	defer dosya.Close()
	//dosya satirlarinda islemler yapmamiz lazim bunun icin bir scanner degiskeni olusturmamiz lazim
	scanner := bufio.NewScanner(dosya)
	//satir kontrolu yapilacak bunla
	satirNumarasi := 1
	//aranan ifade bazen dosya icinde olmayabilir bunun kontrolude yapildi
	butunKontrol := 0
	//scanner.Scan() fonksiyonu verilen dosya icerisinde satirlari adim adim kontrol etmemizi saglar
	for scanner.Scan() {
		//bulundugumuz satirin icindeki yazilar metin diye bir degiskene atandi
		metin := scanner.Text()
		// metinimizi satir satir almayi basardik, aranan kelime de belli, artik varmi yokmu kontroli zamani
		varmi, err := regexp.MatchString(keyWords, metin) // --> keyWords, metin icersinde var mi kontrolu yapilir ve boolen deger doner
		// hata kontrolu yapilir
		if err != nil {
			return err
		}
		// varmi true ise satir numarasi ve satirdaki metin cikti olarak verilir
		if varmi {
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Printf("%d. satir => %s \n", satirNumarasi, metin)
			butunKontrol += 1
		}
		// sacnner.Scan() butun satirlar bitene kadar kontrol yapacagi icin hangi satirda oldugumuzun kontrolunu yapmamiz lazim
		satirNumarasi += 1
	}
	//aranan ifade bazen dosya icinde olmayabilir bunun kontrolude yapildi
	if butunKontrol == 0 {
		fmt.Println("aranan ifade bulunumadi")
	}

	return nil
}


