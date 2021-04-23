# Pokédex

Aplikasi untuk mengetahui status pokemon, move dan damage yang diberikan oleh pokemon kepada musuh

## Building Pokédex

Prerequisites:

1. Golang 1.14
2. Docker
3. Make utility

**Step 1 Create Database and Import Data**

1. Buat database di MySQL
2. Import data yang ada di file `db/dump.sql` [Credit]([https://link](https://github.com/brianr852/Pokemon-Database))

**Step 2 Checkout Pokédex**

```.bash
$ git clone https://github.com/RioRizkyRainey/pokedex.git
$ cd pokedex
```

**Step 3 Prepare Environment**

Silahkan isi file .env sesuai dengan kebutuhan

**Step 4 Build and Run**

```bash
$ make run
```
Menjalankan app akan otomatis build dengan docker.

Jika kamu menggunakan `Windows`, kamu bisa menjalankan aplikasi dengan cara
```bash
$ docker-compose up -d gateway
```

atau

```bash
$ docker-compose up -d
```

Note: pastikan docker sudah terinstall

## Testing Pokédex

```bash
$ make test
``` 

## Postman Pokédex

1. Buka postman
2. Import file `Pokédex.postman_collection.json` ke postmanmu
3. Have fun! :)

## How it's work

Terdapat 4 Microservice dalam repo ini
1. Gateway Service
   Bertujuan sebagai jembatan/penghubung antara client dengan service-service dibelakangnya
2. Pokémon Service
   Memberikan informasi dari Pokémon, seperti attack point dan defense point
3. Move Service
   Memberikan data jurus(move) dari Pokémon
4. Attack Service
   Memberikan data damage yang akan diterima lawan jika. Attack Service akan berhubungan dengan Pokémon Service dan Move Service untuk mengkalkulasi damage

Semua service berkomunikasi dengan gRPC. Namun, data yang akan ditampikan ke Client (Postman, Android apapun itu :D) akan berupa json.

![Diagram](https://github.com/RioRizkyRainey/pokedex/blob/master/assets/diagram.png)

Setiap service akan memiliki 3 layer
1. Delivery
   Menghandle komunikasi service dengan service lainnya, bisa berupa gRPC/Rest
2. Use Case
   Menghandle bussines logic
3. Model
   Mengambil data yang diperlukan, bisa dari database maupun ambil data dari service lain

![Diagram2](https://github.com/RioRizkyRainey/pokedex/blob/master/assets/diagram2.png)

