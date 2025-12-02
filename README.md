# Cli-TODO

Basit bir komut satırı TODO yöneticisi. Görev ekleyebilir, listeleyebilir, silebilir, düzenleyebilir ve tamamlandı olarak işaretleyebilirsin.

## Kurulum

Projeyi klonladıktan sonra:

```bash
# derleme
go build -o tdo

# veya PATH'ine kurulacak şekilde
go install ./...
```

## Kullanım

Yeni Not Ekleme:

```bash
./tdo add "alışveriş yap"
# öncelik vererek
./tdo add -p 1 "önemli iş"
```

Notları Listeleme:

```bash
./tdo list          # yalnızca tamamlanmamışlar
./tdo list --done   # yalnızca tamamlanmışlar
./tdo list --all    # hepsi
```
Notları Düzenleme:
```bash
./tdo edit 1 "new text"
```

Notları Tamamlandı işaretleme:

```bash
./tdo done 1
```
Notları Silme:
```bash
./tdo delete 1
```

## Veri Dosyası

Görevler, varsayılan olarak `.tridos.json` adlı bir JSON dosyasında saklanır. 
Ev dizinindeki dosya kullanılır, ancak çalıştığın dizinde `.tridos.json` varsa o dosya tercih edilir.

## Lisans

Bu proje MIT lisansı ile lisanslanmıştır. Ayrıntılar için `LICENSE` dosyasına bak.

## Info
Bu proje Bilal Yazıcıoğlu tarafından yapılmış ve github'a atılmıştır.