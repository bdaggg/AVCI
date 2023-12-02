package project

import "fmt"

func Logo() {
	logo := `
	____     __    __     ____    _____  
   (    )    ) )  ( (    / ___)  (_   _) 
   / /\ \   ( (    ) )  / /        | |   
  ( (__) )   \ \  / /  ( (         | |   
   )    (     \ \/ /   ( (         | |   
  /  /\  \     \  /     \ \___    _| |__ 
 /__(  )__\     \/       \____)  /_____( 
										 
 `
	fmt.Println(logo)

}

func Yardim() {
	yardim := `
   --aranan    'aramak istediginiz ifadeyi bu degiskenle verebilirsiniz'
   -a          'aramak istediginiz ifadeyi bu degiskenle verebilirsiniz'
   --dosya     'icersinde arama yapmak istediginiz dosyanin yolunu bu degiskenle verebilirsiniz'
   -d          'icersinde arama yapmak istediginiz dosyanin yolunu bu degiskenle verebilirsiniz'
   --yardim    'yardim isteme komutu'
   -y          'yardim isteme komutu'
   --sayac     'sadece aranan ifdenin kaç defa dosya icersinde gectigini gormek icin kullanilir'
   -s          'sadece aranan ifdenin kaç defa dosya icersinde gectigini gormek icin kullanilir'

   ornek kullanim : 
   
   go run .\main.go --aranan "siber" --dosya "C:\Users\bulen\Desktop\cyberSecurity\notes\cyber security edu-1.txt"

   go run .\main.go --aranan "siber" --dosya "C:\Users\bulen\Desktop\cyberSecurity\notes\cyber security edu-1.txt" --sayac

   go run .\main.go --yardim


   `
	fmt.Println(yardim)
}
