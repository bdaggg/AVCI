# AVCI
![Screenshot_2023-12-02_10_49_19](https://github.com/bdaggg/AVCI/assets/110742864/0beba60f-fcc2-4f53-b46d-87c5d985a799)

yeni aracım olan AVCI'yı sizlere tanıtmak istiyorum. Tamamen Türkçe olarak yazdığım bu aracımın ilk sürümünde bir dosyada tarama yaparak 
istediğiniz herhangi bir ifadenin var olup olmadığını veya kaç defa bu ifadenin dosya içerisinde geçtiğini bulabilirsiniz. 

![Screenshot_2023-12-02_10_49_52](https://github.com/bdaggg/AVCI/assets/110742864/10e2daf2-590d-420d-8467-47e34c915f36)

Şekilde görüldüğü gibi go run ./main.go ile AVCI'yı çalıştırıp devamında --yardim veya -y ile kullanım hakkında yardım alabilirsiniz. 

![Screenshot_2023-12-02_10_50_40](https://github.com/bdaggg/AVCI/assets/110742864/f3a2ac3d-fc0f-45a3-b5bd-e46974b98da8)

Eğer bilgisayarınızda herhangi bir dosyada herhangi bir ifadenin geçtiği satrları görmek isterseniz yukarıda görülüdüğü gibi 
go run ./main.go --aranan "aranan_ifade" --dosya "dosyanın_konumu_buraya_gelecek"
bu şekilde çalıştırabilirsiniz

![Screenshot_2023-12-02_10_50_16](https://github.com/bdaggg/AVCI/assets/110742864/44dbcd71-af80-4ff5-b0f0-c74cadc282b1)

satırları değilde sadece ifadenin kaç defa dosya içerisinde geçtiğini görmek için ise  aynı şekilde aranan kelime ve dosya konumu verilir ek olarak
--sayac eklenebilir. go run ./main.go --aranan "aranan_ifade" --dosya "dosyanın_konumu_buraya_gelecek" --sayac bu şekilde. 


AVCI'nın sonraki sürümlerinde eklenecekler:
-- ismi verilen bir dosyanın bilgisayar içerisindeki konumunu bulma işlemini yapabilecek (linux sistemlerindeki find komutu gibi düşünülebilir)
-- dosya okuma işlemini yapabilecek (linux sistemlerdeki cat komutu gibi düşünülebilir)
-- dosya şifreleme işlemini yapabilecek
