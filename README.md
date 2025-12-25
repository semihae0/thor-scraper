Bu proje, Go (Golang) dili kullanılarak geliştirilmiş, Tor ağı üzerindeki `.onion` uzantılı web sitelerini anonim olarak tarayan, durumlarını raporlayan ve içeriklerini kaydeden bir otomasyon aracıdır.

Proje, siber tehdit istihbaratı (CTI) süreçlerinde veri toplama (Collection) yetkinliğini geliştirmek amacıyla eğitim odaklı hazırlanmıştır.

* **Anonim Bağlantı:** Tüm trafik yerel SOCKS5 proxy üzerinden Tor ağına yönlendirilir.
* **IP Sızıntı Koruması:** Gerçek IP adresini gizler ve bağlantı öncesi Tor IP kontrolü yapar.
* **User-Agent Spoofing:** Modern tarayıcı (Chrome/Windows) taklidi yaparak engellemeleri aşar.
* **Toplu Tarama:** `targets.yaml` dosyasından okunan çok sayıda hedefi sırayla tarar.
* **Raporlama:** Tarama sonuçlarını `scan_report.log` dosyasına kaydeder.
* **Veri Kaydı:** Erişilebilen sitelerin HTML kaynak kodlarını yerel diske indirir.
* **Go (Golang):**
* **Tor Service:** Arka planda çalışan bir Tor servisi.
  *Tor Browser* açıkken (Genellikle Port: 9150)
  *Tor Expert Bundle* servisi (Genellikle Port: 9050)
 
  -- Projeyi klonlayın: --
      git clone https://github.com/semihae0/thor-scraper.git

  -- Gerekli Go modüllerini indirin: --
      go mod tidy

   Yapılandırma **ÖNEMLİ**
  Tarama yapılacak adresleri belirlemek için ana dizinde `targets.yaml` adındaki dosyaya aşşağıdaki formatta linkleri giriniz:
  
adresler:
  - http://örneksite1.onion/
  - http://örneksite2.onion/
  - http://örneksite3.onion/
  - http://örneksite4.onion/
  - http://örneksite5.onion/
 


Sizin için `targets.yaml` adındaki dosyaya yapıştırabileceğiniz örnekler koydum:

adresler:
  - http://zwf5i7hiwmffq2bl7euedg6y5ydzze3ljiyrjmm7o42vhe7ni56fm7qd.onion/
  - http://vu3miq3vhxljfclehmvy7ezclvsb3vksmug5vuivbpw4zovyszbemvqd.onion/
  - http://pz5uprzhnzeotviraa2fogkua5nlnmu75pbnnqu4fnwgfffldwxog7ad.onion/
  - http://2bcbla34hrkp6shb4myzb2wntl2fxdbrroc2t4t7c3shckvhvk4fw6qd.onion/
  - http://pliy7tiq6jf77gkg2sezlx7ljynkysxq6ptmfbfcdyrvihp7i6imyyqd.onion/
  - http://check.torproject.org/
