# Cli-TODO

Basit bir komut satırı TODO yöneticisi. Görev ekleyebilir, listeleyebilir ve tamamlandı olarak işaretleyebilirsin.

## Kurulum

Projeyi klonladıktan sonra:

```bash
# derleme
go build -o tdo

# veya PATH'ine kurulacak şekilde
go install ./...
```

## Kullanım

Yeni iş ekleme:

```bash
tdo add "alışveriş yap"
# öncelik vererek
tdo add -p 1 "önemli iş"
```

Listeleme:

```bash
tdo list          # yalnızca tamamlanmamışlar
tdo list --done   # yalnızca tamamlanmışlar
tdo list --all    # hepsi
```

Bir işi tamamlandı işaretleme:

```bash
tdo done 1
```

## Veri Dosyası

Görevler, varsayılan olarak `.tridos.json` adlı bir JSON dosyasında saklanır. 
Ev dizinindeki dosya kullanılır, ancak çalıştığın dizinde `.tridos.json` varsa o dosya tercih edilir.

## Lisans

Bu proje MIT lisansı ile lisanslanmıştır. Ayrıntılar için `LICENSE` dosyasına bak.
