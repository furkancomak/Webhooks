## İçindekiler
  * [Bilgisayarıma Neler Kurmalıyım?](#bilgisayarıma-neler-kurmalıyım?)
    + [VPN Kurulumu](#vpn-kurulumu)
    + [AWS Kurulumu](#aws-kurulumu)
  * [Projeler](#projeler)
    + [Kütüphaneler](#kütüphaneler)
    + [Mikroservisler](#mikroservisler)
  * [Takımlar](#takımlar)
      - [F1](#f1)
      - [APOLLO](#apollo)
      - [ATLAS](#atlas)
      - [MİGPACK](#MİGPACK)
      - [DATA](#data)
  * [İş Geliştirme Süreci](#İş-geliştirme-süreci)
      - [Scrum Süreçleri](#scrum-süreçleri)
      - [Jira Süreci](#jira-süreci)
      - [Environmentlar ve Test Ortamları](#environmentlar-ve-test-ortamları)
      - [Monitoring ve Logging Süreçleri](#monitoring-ve-logging-süreçleri)
    + [Slack](#slack)
  * [İK İşlemleri](#İK-İşlemleri)

# Aramıza Hoşgeldin

  Bu yeni yolculuğunda sana yardımcı olabilecek bazı bilgileri seninle paylaşmak isteriz. Aşağıda kullandığımız programları, deployment sürecinde nasıl ilerlediğimizi ve projelerimiz hakkında kısa açıklamalar bulabilirsin. Takıldığın konularda herkes sana yardımcı olacaktır. Aramıza hoşgeldin :)

## Bilgisayarıma Neler Kurmalıyım?

  Kurulması gereken çok şey var. Tabi ki hepsini tek tek kurmaya vaktini harcamamalısın.

- [ ] Bunun için [buradaki](https://github.com/startupheroes/dotfiles) dosyaları kullanabilirsin. Gerekli talimatlar içeride yazıyor.

- [ ] Eğer (-sh veya migros uzantılı) mail adresin açılmadıysa **Mert Sarıboğa** sana yardımcı olacaktır.

- [ ] Mail adresin var mı? Jira, github ve slack'e katılman için **Mert Sarıboğa**'e gidebilirsin.

- [ ] Homebrew, programların konsol aracılığıyla yönetimini sağlayan bir paket yönetim sistemidir. [Sitesindeki](https://brew.sh/index_tr) talimatlara bakarak kurulumunu yapabilirsin.

- [ ] IDE olarak jetbrainsin araçlarını kullanıyoruz. [Buraya tıklayarak](https://www.jetbrains.com) Jetbrains'in sitesine ulaşabilirsin. Jetbrains hesabı için  **Ömer Başar** ile görüşebilirsin.

  - Ide : Java için IntelliJ Idea, Javascript için Webstorm'u kullanıbilirsin.  Kurulumu dotfiles ile kurulan intellij-toolbox üzerinden veya *brew cask install intellij-idea* , *brew cask install webstorm* komutlarını terminale yazarak yapabilirsin.

    - Checkstyle : Intellij ayarlari icin (not: her yeni intellij projesinde yapilmasi gerekiyor) kullanılıyor. [Buradan](https://github.com/startupheroes/startupheroes-checkstyle/wiki/3.-How-to-use-StartupHeroes-Checkstyle-in-Intellij-IDEA) siteye ulaşarak gerekli bilgileri bulabilirsin. Checkstyle statik kod analizi yapılıp class içerisinde kendi iç kurallarımızı uygulayabilmemizi sağlıyor.

    - Codestyle : Java ve android projelerinde Intellij IDEA kod (boşluk ayarlaması, satır sınırı, satırbaşı ayarlaması) düzenlemesini sağlıyor. [Buradan](https://github.com/startupheroes/startupheroes-codestyle) dosyalara ve kurulum bilgilerine ulaşabilirsin.

  - Database Client :Database erişimi için DataGrip'i indirebilirsin. *brew cask install datagrip* komutunu terminale yazarak da kurulumunu sağlayabilirsin.

  - Terminal : Iterm veya Mac'in kendi terminalini kullanıyoruz.

    > jenv:  Java versiyonunu kontrol edebilmek ve bazı uygulamalarda java versiyonumuzu set etmek için kullanıyoruz. [Sitesinden](https://www.jenv.be) indirip kullanım ve kurulum detaylarını bulabilirsin.

  ### VPN Kurulumu

  - VPN'e bağlanmak için tunnelblick kullanıyoruz. Yukarıda kurduğun dosyaların içerisinde mevcut.

  - Slack üzerinde _ops-vpn kanalına girerek !generate migros-firstnamelastname-device komutunu girerek bir *QR kod* ve *config* dosyası elde ediyoruz.

    Örnek kullanım: *!generate migros-furkancomak-macbook*

  - QR kodu telefonumuzdaki Authenticator'a okutuyoruz ve OTP şifremizi elde ediyoruz.

  - Tunnelblick'e indirdiğimiz config dosyasını atıyoruz ve ayarlardan VPN log level'ı **2.4.8-OpenSSL v1.0.2u** olarak ayarlayarak vpn e bağlanabiliriz.

  ### AWS Kurulumu

  - Öncelikle Keybase hesabı açıp bir tane key tanımlamalısın.

  - Oluşturduğun key'i **Mehmet Başal'a** göndermelisin.
  - İhtiyacına göre console giriş bilgilerini ya da access key ve secret key bilgilerini Mehmet Başal sana gönderecektir.
  - Console'a giriş yapmak için [aws.amazon](signin.aws.amazon.com) 'a girip
    - *Account or Alias=emigros*
    - *Username=Belirlediğin kullanıcı adı*
    - *Password=Belirlediğin şifre*
  - Command line üzerinde **aws configure** komutunu kullanıp
    - Access Key
    - Secret Key
    - Region=eu-west1 bilgilerini girerek giriş yapabilirsin.
  - AWS configure için **brew install awscli** komutunu kullanıyoruz.

- [ ] Son olarak AWS, VPN ve Keybase hesaplarında kafana takılan bir konu olursa **Mehmet Başal** ile görüşebilirsin.


> VPN, AWS ve github'a girişte OTP kullanılıyoruz. Bunun için telefonuna Google Authenticator indirebilirsin.



## Projeler

**Sanalmarket:** Sanalmarket mikroservisi [website](www.sanalmarket.com.tr)

**Macrocenter:** Macrocenter mikroservisi [website](www.macrocenter.com.tr)

**Tazedirekt:** Tazedirekt mikroservisi[website](https://www.tazedirekt.com)

**Migros Hemen:** Migros Hemen mikroservisi [website](https://www.migroshemen.com)

**Cockpit:** Ana yönetici uygulaması [website](https://cockpit.migros.com.tr/login)

- Cockpit hesabı açman gerekiyor. Bunun için **Mert Sarıboğa** Sana yardımcı olacaktır.

**Collectify:** Toplama ve dağıtım servisi [website](http://collectify.tazedirekt.com/)

**Migpack:** Ürün Paketleme servisi [website](https://cockpit.migpack.app/)

**Cronify:** Zamanlı işlerin yönetim servisi

### Kütüphaneler

  **Clovify:** Uygulamaların kullanım altyapısı

  **Trink:** Ödeme altyapısı

  **Dotfiles:** Bilgisayara kurulumuna yardımcı dosyaların projesi

###       Mikroservisler

​     **Actor:** Kullanıcı ve cihaz servisleri

​     **Lotus:** Ödeme ve sipariş servisleri

​     **Notify:** Mail, sms, app ve web bildirim servisleri

​     **Cerebro:** Arama servisi

​     **Proseidon:** Product yönetim servisi

​     **Hermes:** Dış kaynaklarla entegrasyon servisi

​     **Shipify:** Adres ve dağıtım servisi

​     **Catalog:** Marka ve kategori servisi

​     **Persona:** Kullanıcı kişisel özellikleri yönetim servisi

​     **Markomation:** Kampanya altyapısı

​     **Dali:** Resim kesme servisi

​     **Pathfinder:** Url kısaltma, yeniden yönlendirmeler gibi path servisleri

​     **Stockroom:** Stok servisi

​     **Harmony:** Tamim, Ays, ETL gibi migros ile konustugumuz servisler

​       **- Tamim:** Migros ile iletişim kurduğumuz dış servislerden biridir. Fiziksel mağaza ile sanal mağaza arasında ürün, kategori, kampanya için data kaynağını sağlar. Yani ürün akışı, kategoriler ve kampanya kurguları Tamim'den gelmektedir. Böylece fiziksel mağaza ortamı sanal ortama taşınabilmektedir.


## Takımlar

####      F1

Team id= F1

####      APOLLO

Team id= AP

####      ATLAS

Team id= CL

####      MİGPACK

Team id= MIP

####      DATA

Team id= AI


## İş Geliştirme Süreci

##### Scrum Süreçleri

- İş geliştirme programımız **ikişer haftalık sprintler** temelinde ilerliyor.

- Sprint toplantılarında hangi işlerin yapılacağına karar veriliyor ve iş puanlaması yapılıyor.

    - İşin değerine, bir backlog işinin, diğer backlog isleri ile karşılaştırıldığında tahmini olarak ne kadar çaba gerektireceği baz alınarak fibonacci serisine göre puanlama yapılır. Puanlama için telefonuna app (Scrum Time, Scrum Poker vb) kurman gerekiyor.

- Her sabah yaklaşık 15 dakikalık **daily** yapılıyor.Bu toplantılarda bir önceki gün neler yaptığımızı ve o gün neler yapacağımızı kısa bir şekilde takım arkadaşlarımıza anlatıyoruz.

- Sprint ortasında, sprintte ne kadar ilerlediğimizi konuşmak ve bir sonraki sprint hedeflerini ve geliştirilecek ürün listesine nelerin eklenmesini tartışmak için **Grooming** toplantısı yapıyoruz.

- Ayda bir defa olarak da **Retrospective** toplantılarımız oluyor. Bu toplantının içeriği ise önceki sprint'in takım için işleyişin, alınan kararların nasıl gittiğini, takım içi veya ortam konusunda nasıl geliştirmelerin yapılabileceğini konuştuğumuz ve bunların planlamasını yaptığımız toplantıdır.

  ##### Jira Süreci

- İşler jira üzerinden açılıyor. Yapılacak iş için gerekli bilgiler iş detayında anlatılıyor.

- İş detayından üzerinde çalışacağın mikroservisi görüp orada çalışmaya başlayabilirsin.

- Master branch üzerinde çalışmamaya dikkat ediyoruz.

- Branch isimlerini **"f-Team_id-Jira Kodu"** şeklinde koyuyoruz. Yeni bir feature için *f-* bug için ise *b-* ile başlayabilirsin.

  - Örnek kullanım: b-cl-193

- Commit için ise **"Team_id-Jira kodu: Açıklama"** taslağını kullanıyoruz. Açıklama kısmının fiil (emir kipi) ile ve büyük harf ile başlaması gerekiyor yapılan işle alakalı kısa bir açıklama yeterlidir. Noktalama işaretlerini kullanmıyoruz.

  - Örnek kullanım: CL-1246: Add log to manual advance

- Terminal veya IDE üzerinden branch açma ve commitleme işinden sonra Github'ta olduğun repositorye giderek Pull Request açman gerekiyor.

- Pull Request açarken doldurman gereken 3 kısım var.

  - Assignee olarak kendini atamalısın.
  - Reviewer olarak takımından veya o işle ilgilenen minimum 2 kişi atamalısın.
  - Type kısmında da hazır typeları seçerek (bug, enhancement vb) işin ne ile alakalı olduğunu belirlemelisin.

  ![Checklist](https://github.com/startupheroes/wiki/blob/master/onboarding/Resources/checklist.png)

- Pull Requesti açtık ve Github üzerindeki bazı checklerin tamamlanmasını bekledik.

##### Environmentlar ve Test Ortamları

- "test, stage ve prod" kullanmakta olduğumuz environmentlardır.
  - Test environment'ı testleri yaptığımız ortamımızdır. Testi kendimiz veya test ekibindekilerle beraber yapabiliriz. İlk olarak kodumuzu bu environmenta atmamız gerekiyor. Eğer bir sorun çıkmadıysa stage ortamına taşıyabiliriz.

    - **Sanal Market Test Ortamları :**
      - test.migros.com.tr
      - test1.migros.com.tr
      - test2.migros.com.tr

    - **Tazedirekt Test Ortamı**
      - test.tazedirekt.com
      - test1.tazedirekt.com

    - **Collectify Test Ortamı**
      - test.collectify.tazedirekt.com

    - **Cockpit Test Ortamı**
      - cockpit.[website].com / com.tr

  - Stage environment'ı prod ile aynı databesi kullanmakta fakat dışarıdaki kullanıcılara etkisi olmayan bir ortamdır. Burada da gerçek verilerle tekrar test edilir, eğer başka kısımları etkilemiyor ve testlerden de başarıyla geçtiyse kodumuzu prod ortamına taşıyabiliriz.

  - Prod environment'ı ise dışarıdaki kullanıcıların da ürününüzü kullanabildiği ortamdır. Kodumuzu bu ortama taşıdığımızda artık kodumuz canlı ortamdadır.

- Check listesi onaylandıysa, Slack üzerindeki **_ops-deploy** kanalına geliyoruz. **!deploy commit hash kodu project_name/environment** formatını kullanarak kodu istenen environmenta atabilirsin.
- Örnek kullanım: !deploy 3c776cc notify/test

##### Monitoring ve Logging Süreçleri

- **Scalyr:**

  - Mikroservislerin loglarını takip edebilmek için [Scalyr](https://www.scalyr.com) hesabı oluşturman gerekiyor.
  - Hesabı oluşturduktan sonra organizasyona davetini **Mehmet Başal**'dan isteyebilirsin.
  - Scalyr sayesinde log levellerini mikroservis durumunu vb görüntüleyebilirsin.

- Deploy komutu sonrasında kodumuzun istediğimiz ortama geçip geçmediğini kontrol etmek için test.migros.com.tr/info (örnek kullanım) sitesine girebiliriz. Kullanılan api'ları görebilmek için de rest.migros.com.tr/info (örnek kullanım) sitesine girebiliriz. Giriş bilgileri aşağıdaki gibidir.

  - Kullanıcı adı: **swagger** veya **tester**           şifre: **qaz890***

    > Maven'a girebilmek için Settings.xml dosyasına ihtiyacın var. Takımındaki herhangi bir kişiden isteyebilirsin.

  ### Slack

  - Deploy, Restart, Rollback, VPN açma gibi işlemleri slack botumuz üzerinden rahatlıkla yapabilirsin.
  - *!list projects* komutuyla da kullandığımız projelerin tamamını görebilirsin.
  - Yukarıdaki komutların kullanımıyla alakalı gerekli bilgileri _ops-deploy kanalına !help yazarak ulaşabilirsin.

## İK İşlemleri

- Kalan izin günleri, izin talep belgesi ve diğer ik işlemlerini:
  - Startupheroes bordroluysan: [Kolay İK](https://kolayik.com/app/#/login/home)
  - Migros bordroluysan [Otonet](http://10.34.1.80/otonet/) siteleri üzerinden yapabilirsin.

>  Fotoğrafınızı, Jira, Github, KolayIK ve Slack'e eklemeniz insanlara sizi bulma konusunda yardımcı olacaktır. Lütfen fotoğrafınızı yüklemeyi unutmayın :).
