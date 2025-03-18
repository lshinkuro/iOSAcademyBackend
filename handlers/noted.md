// redis-cli -u redis://default:ouFsknrJmW1QaLTe3uX1387jLCucdy6a@redis-16695.c98.us-east-1-4.ec2.redns.redis-cloud.com:16695

# Redis CLI Tutorial

## **Pendahuluan**
Redis CLI (`redis-cli`) adalah command-line interface untuk berinteraksi dengan Redis. Dengan `redis-cli`, kita dapat menyimpan, mengambil, dan menghapus data dengan mudah.

---

## **1. Instalasi Redis**
Jika Redis belum terinstal, silakan instal terlebih dahulu sesuai dengan sistem operasi Anda:

### **MacOS (Homebrew)**
```sh
brew install redis
```

### **Ubuntu/Debian**
```sh
sudo apt update
sudo apt install redis-server -y
```

### **Windows (Menggunakan WSL)**
```sh
wsl --install -d Ubuntu
sudo apt update
sudo apt install redis-server -y
```

Pastikan Redis berjalan dengan perintah:
```sh
redis-server
```

---

## **2. Mengakses Redis CLI**
Jalankan perintah berikut untuk masuk ke Redis CLI:
```sh
redis-cli
```
Atau jika Redis berjalan di server tertentu:
```sh
redis-cli -h <host> -p <port>
```
Contoh:
```sh
redis-cli -h 127.0.0.1 -p 6379
```

---

## **3. Perintah Dasar Redis CLI**

### **Menyimpan Data**
```sh
SET myKey "Hello, Redis!"
```

### **Mengambil Data**
```sh
GET myKey
```
Output:
```
"Hello, Redis!"
```

### **Melihat Semua Key yang Tersimpan**
```sh
KEYS *
```
Atau jika menggunakan Redis terbaru:
```sh
SCAN 0
```

### **Menghapus Data**
```sh
DEL myKey
```

### **Mengecek Apakah Key Ada**
```sh
EXISTS myKey
```
Output:
- `1` jika key ada
- `0` jika key tidak ada

### **Melihat Waktu Kedaluwarsa dan Tipe Data**
```sh
TTL myKey    # Melihat waktu kedaluwarsa key
TYPE myKey   # Melihat tipe data key
```

### **Menyimpan Data dengan Expiry Time**
```sh
SETEX session 60 "User123"
```
Key `session` akan dihapus otomatis setelah 60 detik.

---

## **4. Keluar dari Redis CLI**
Untuk keluar dari Redis CLI, gunakan perintah:
```sh
EXIT
```

---

## **5. Bantuan dan Dokumentasi**
Jika membutuhkan bantuan tambahan dalam Redis CLI, jalankan perintah berikut:
```sh
HELP
```

---

## **Penutup**
Redis CLI adalah alat yang sangat berguna untuk berinteraksi dengan database Redis. Dengan perintah-perintah di atas, Anda dapat dengan mudah menyimpan, mengambil, dan mengelola data di Redis.

Selamat mencoba! 🚀

