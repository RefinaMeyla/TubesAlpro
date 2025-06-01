package main

import "fmt"

// Kegiatan merepresentasikan sebuah kegiatan UKM
type Kegiatan struct {
	ID            int
	NamaKegiatan  string
	Hari          string
	JumlahPeserta int
}

// StatistikHarian digunakan untuk menyimpan jumlah kegiatan per hari
type StatistikHarian struct {
	Hari   string
	Jumlah int
}

// Variabel Global
const NMAX = 100         
const MAX_HARI_UNIK = 10 

var daftarKegiatan [NMAX]Kegiatan 
var jumlahKegiatanSekarang = 0    

// Fungsi Utama
func main() {
	state := true

	for state {
		fmt.Println("\n+=============================================+")
		fmt.Println("|     Manajemen Jadwal Kegiatan UKM Tel-U     |")
		fmt.Println("+=============================================+")
		fmt.Println("| 1. Tambah Kegiatan                          |")
		fmt.Println("| 2. Lihat Semua Kegiatan                     |")
		fmt.Println("| 3. Hitung Total Peserta                     |")
		fmt.Println("| 4. Update Kegiatan                          |")
		fmt.Println("| 5. Hapus Kegiatan                           |")
		fmt.Println("| 6. Cari Kegiatan                            |")
		fmt.Println("| 7. Statistik Kegiatan per Hari              |")
		fmt.Println("| 8. Keluar                                   |")
		fmt.Println("+=============================================+")

		var pilihan int
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan) 

		switch pilihan {
		case 1:
			tambahKegiatan()
		case 2:
			tampilkanSemuaKegiatan()
		case 3:
			hitungTotalPeserta()
		case 4:
			updateKegiatan()
		case 5:
			hapusKegiatan()
		case 6:
			cariKegiatan()
		case 7:
			statistikHari()
		case 8:
			state = false
			fmt.Println("Terima kasih telah menggunakan program.")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Menyiapkan tempat untuk menyimpan hingga 100 data kegiatan.
func tambahKegiatan() {
	if jumlahKegiatanSekarang >= NMAX {
		fmt.Println("Kapasitas kegiatan sudah penuh!")
		return
	}

	var kegiatan Kegiatan
	kegiatan.ID = jumlahKegiatanSekarang + 1

	fmt.Print("Masukkan Nama Kegiatan (satu kata): ")
	fmt.Scan(&kegiatan.NamaKegiatan)
	if kegiatan.NamaKegiatan == "" {
		fmt.Println("Nama Kegiatan tidak boleh kosong.")
		return
	}

	fmt.Print("Masukkan Hari (Contoh: Senin): ")
	fmt.Scan(&kegiatan.Hari)
	if kegiatan.Hari == "" {
		fmt.Println("Hari tidak boleh kosong.")
		return
	}

	fmt.Print("Masukkan Jumlah Peserta: ")
	fmt.Scan(&kegiatan.JumlahPeserta)
	if kegiatan.JumlahPeserta <= 0 {
		fmt.Println("Jumlah Peserta harus lebih dari 0.")
		return
	}

	daftarKegiatan[jumlahKegiatanSekarang] = kegiatan
	jumlahKegiatanSekarang++
	fmt.Println("Kegiatan berhasil ditambahkan!")

	var urutan string
	var i, j int
	fmt.Print("Urutkan berdasarkan ID (asc/desc)? (ketik 'asc' atau 'desc', lainnya tidak diurutkan): ")
	fmt.Scan(&urutan)

	n := jumlahKegiatanSekarang
	if urutan == "asc" {
		for i = 0; i < n-1; i++ {
			for j = 0; j < n-i-1; j++ {
				if daftarKegiatan[j].ID > daftarKegiatan[j+1].ID {
					temp := daftarKegiatan[j]
					daftarKegiatan[j] = daftarKegiatan[j+1]
					daftarKegiatan[j+1] = temp
				}
			}
		}
		fmt.Println("Daftar kegiatan diurutkan berdasarkan ID (ascending).")
	} else if urutan == "desc" { // Case-sensitive
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if daftarKegiatan[j].ID < daftarKegiatan[j+1].ID {
					temp := daftarKegiatan[j]
					daftarKegiatan[j] = daftarKegiatan[j+1]
					daftarKegiatan[j+1] = temp
				}
			}
		}
		fmt.Println("Daftar kegiatan diurutkan berdasarkan ID (descending).")
	}
}

// untuk mencatat dan menunjukkan jumlah kegiatan yang telah dimasukkan 
// ke dalam sistem, serta sebagai penanda posisi indeks berikutnya dalam arra
func tampilkanSemuaKegiatan() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan yang terdaftar.")
		return
	}

	fmt.Println("\n+------------------------------------------------------+")
	fmt.Println("|                 Daftar Kegiatan UKM                |")
	fmt.Println("+----+---------------------+------------+-------------+")
	fmt.Println("| ID |    Nama Kegiatan    |    Hari    | Jml Peserta |")
	fmt.Println("+----+---------------------+------------+-------------+")
	for i := 0; i < jumlahKegiatanSekarang; i++ {
		kegiatan := daftarKegiatan[i]
		fmt.Printf("| %-2d | %-19s | %-10s | %-11d |\n", kegiatan.ID, kegiatan.NamaKegiatan, kegiatan.Hari, kegiatan.JumlahPeserta)
	}
	fmt.Println("+----+---------------------+------------+-------------+")
}

// untuk menjumlahkan semua nilai peserta dari setiap kegiatan
func hitungTotalPeserta() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan untuk dihitung total pesertanya.")
		return
	}
	total := 0
	for i := 0; i < jumlahKegiatanSekarang; i++ {
		total += daftarKegiatan[i].JumlahPeserta
	}
	fmt.Printf("Total jumlah peserta semua kegiatan: %d\n", total)
}

//mencari posisi (indeks) kegiatan berdasarkan ID yang dimasukkan
func cariIndeksKegiatanByID(id int) int {
	for i := 0; i < jumlahKegiatanSekarang; i++ {
		if daftarKegiatan[i].ID == id {
			return i
		}
	}
	return -1
}

// memperbarui data kegiatan tertentu berdasarkan ID yang dimasukkan, termasuk nama, 
// hari, dan jumlah pesertanya, jika ID tersebut ditemukan dalam daftar.
func updateKegiatan() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan untuk diupdate.")
		return
	}
	tampilkanSemuaKegiatan()

	var id int
	fmt.Print("Masukkan ID kegiatan yang akan diupdate: ")
	fmt.Scan(&id)

	indeks := cariIndeksKegiatanByID(id)

	if indeks != -1 {
		fmt.Printf("Mengupdate kegiatan dengan ID %d (%s):\n", daftarKegiatan[indeks].ID, daftarKegiatan[indeks].NamaKegiatan)

		var namaBaru string
		fmt.Print("Masukkan Nama Kegiatan baru (satu kata): ")
		fmt.Scan(&namaBaru)
		if namaBaru != "" { 
			daftarKegiatan[indeks].NamaKegiatan = namaBaru
		} else {
			fmt.Println("Nama kegiatan baru tidak boleh kosong jika diinput.")
		}

		var hariBaru string
		fmt.Print("Masukkan Hari baru: ")
		fmt.Scan(&hariBaru)
		if hariBaru != "" {
			daftarKegiatan[indeks].Hari = hariBaru
		} else {
			fmt.Println("Hari baru tidak boleh kosong jika diinput.")
		}

		var jumlahPesertaBaru int
		fmt.Print("Masukkan Jumlah Peserta baru (masukkan 0 untuk mengosongkan): ")
		fmt.Scan(&jumlahPesertaBaru)
		if jumlahPesertaBaru < 0 {
			fmt.Println("Jumlah Peserta baru tidak boleh negatif.")
			return 
		}
		daftarKegiatan[indeks].JumlahPeserta = jumlahPesertaBaru

		fmt.Println("Kegiatan berhasil diupdate!")
	} else {
		fmt.Println("Kegiatan dengan ID tersebut tidak ditemukan.")
	}
}

//menghapus kegiatan dari daftar berdasarkan ID 
func hapusKegiatan() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan untuk dihapus.")
		return
	}
	tampilkanSemuaKegiatan()

	var id int
	fmt.Print("Masukkan ID kegiatan yang akan dihapus: ")
	fmt.Scan(&id)

	indeksDitemukan := -1
	for i := 0; i < jumlahKegiatanSekarang; i++ {
		if daftarKegiatan[i].ID == id {
			indeksDitemukan = i
			break
		}
	}

	if indeksDitemukan != -1 {
		for i := indeksDitemukan; i < jumlahKegiatanSekarang-1; i++ {
			daftarKegiatan[i] = daftarKegiatan[i+1]
		}
		daftarKegiatan[jumlahKegiatanSekarang-1] = Kegiatan{}
		jumlahKegiatanSekarang--
		fmt.Println("Kegiatan berhasil dihapus!")
	} else {
		fmt.Println("Kegiatan dengan ID tersebut tidak ditemukan.")
	}
}

// mencari dan menampilkan kegiatan berdasarkan input pengguna apakah dengan nama/hari atau 
//ID secara tepat sesuai data yang tersimpan
func cariKegiatan() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan untuk dicari.")
		return
	}

	var pilihanCari string
	fmt.Print("Cari berdasarkan (nama/id): ")
	fmt.Scan(&pilihanCari)

	found := false
	fmt.Println("\n+------------------------------------------------------+")
	fmt.Println("|                    Hasil Pencarian                   |")
	fmt.Println("+----+---------------------+------------+-------------+")
	fmt.Println("| ID |    Nama Kegiatan    |    Hari    | Jml Peserta |")
	fmt.Println("+----+---------------------+------------+-------------+")

	if pilihanCari == "nama" { 
		var keyword string
		fmt.Print("Masukkan nama kegiatan atau hari yang ingin dicari (EXACT MATCH, case-sensitive, satu kata): ")
		fmt.Scan(&keyword)

		for i := 0; i < jumlahKegiatanSekarang; i++ {
			kegiatan := daftarKegiatan[i]
			if kegiatan.NamaKegiatan == keyword || kegiatan.Hari == keyword { // Exact match
				fmt.Printf("| %-2d | %-19s | %-10s | %-11d |\n", kegiatan.ID, kegiatan.NamaKegiatan, kegiatan.Hari, kegiatan.JumlahPeserta)
				found = true
			}
		}
	} else if pilihanCari == "id" {
		var cariID int
		fmt.Print("Masukkan ID yang ingin dicari: ")
		fmt.Scan(&cariID)

		indeks := cariIndeksKegiatanByID(cariID)
		if indeks != -1 {
			kegiatan := daftarKegiatan[indeks]
			fmt.Printf("| %-2d | %-19s | %-10s | %-11d |\n", kegiatan.ID, kegiatan.NamaKegiatan, kegiatan.Hari, kegiatan.JumlahPeserta)
			found = true
		}
	} else {
		fmt.Println("Pilihan pencarian tidak valid (hanya 'nama' atau 'id').")
	}

	if !found && (pilihanCari == "nama" || pilihanCari == "id") {
		fmt.Println("|            Tidak ada kegiatan yang cocok.            |")
	}
	fmt.Println("+----+---------------------+------------+-------------+")
}

// menghitung dan menampilkan jumlah kegiatan yang terjadi pada 
// setiap hari unik dari daftar kegiatan yang terdaftar.
func statistikHari() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan yang terdaftar untuk statistik.")
		return
	}

	var statistik [MAX_HARI_UNIK]StatistikHarian
	var jumlahHariUnik = 0

	for i := 0; i < jumlahKegiatanSekarang; i++ {
		kegiatan := daftarKegiatan[i]
		hariKegiatan := kegiatan.Hari
		ditemukan := false

		for j := 0; j < jumlahHariUnik; j++ {
			if statistik[j].Hari == hariKegiatan {
				statistik[j].Jumlah++
				ditemukan = true
				break
			}
		}

		if !ditemukan {
			if jumlahHariUnik < MAX_HARI_UNIK {
				statistik[jumlahHariUnik].Hari = hariKegiatan
				statistik[jumlahHariUnik].Jumlah = 1
				jumlahHariUnik++
			} else {
			}
		}
	}

	fmt.Println("\n+--------------------------------------+")
	fmt.Println("|     Statistik Kegiatan per Hari      |")
	fmt.Println("+-------------------+------------------+")
	fmt.Println("|        Hari       | Jumlah Kegiatan  |")
	fmt.Println("+-------------------+------------------+")
	for i := 0; i < jumlahHariUnik; i++ {
		stat := statistik[i]
		fmt.Printf("| %-17s | %-16d |\n", stat.Hari, stat.Jumlah)
	}
	fmt.Println("+-------------------+------------------+")
	fmt.Printf("| Total Semua Kegiatan: %-13d |\n", jumlahKegiatanSekarang)
	fmt.Println("+--------------------------------------+")
}

// Insertion Sort: Mengurutkan jumlah peserta dari kecil ke besar
func insertionSortJumlahPesertaAsc() {
	for i := 1; i < jumlahKegiatanSekarang; i++ {
		temp := daftarKegiatan[i]
		j := i - 1
		for j >= 0 && daftarKegiatan[j].JumlahPeserta > temp.JumlahPeserta {
			daftarKegiatan[j+1] = daftarKegiatan[j]
			j--
		}
		daftarKegiatan[j+1] = temp
	}
	fmt.Println("Data diurutkan berdasarkan Jumlah Peserta (dari terkecil).")
}

// Selection Sort: Mengurutkan jumlah peserta dari besar ke kecil
func selectionSortJumlahPesertaDesc() {
	for i := 0; i < jumlahKegiatanSekarang-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahKegiatanSekarang; j++ {
			if daftarKegiatan[j].JumlahPeserta > daftarKegiatan[maxIdx].JumlahPeserta {
				maxIdx = j
			}
		}
		if i != maxIdx {
			daftarKegiatan[i], daftarKegiatan[maxIdx] = daftarKegiatan[maxIdx], daftarKegiatan[i]
		}
	}
	fmt.Println("Data diurutkan berdasarkan Jumlah Peserta (dari terbesar).")
}

// Binary Search: Mendeteksi tabrakan jadwal berdasarkan hari
func binarySearchHari(target string) bool {
	left := 0
	right := jumlahKegiatanSekarang - 1

	for left <= right {
		mid := (left + right) / 2
		if daftarKegiatan[mid].Hari == target {
			return true
		} else if daftarKegiatan[mid].Hari < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// mengecek apakah sudah ada kegiatan yang berlangsung pada hari tertentu dengan mengurutkan daftar kegiatan 
// berdasarkan hari lalu melakukan pencarian hari tersebut.
func cekTabrakanHari() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan.")
		return
	}

	var hari string
	fmt.Print("Masukkan hari yang ingin dicek tabrakan: ")
	fmt.Scan(&hari)

	// Urutkan daftar berdasarkan hari (bubble sort)
	for i := 0; i < jumlahKegiatanSekarang-1; i++ {
		for j := 0; j < jumlahKegiatanSekarang-i-1; j++ {
			if daftarKegiatan[j].Hari > daftarKegiatan[j+1].Hari {
				daftarKegiatan[j], daftarKegiatan[j+1] = daftarKegiatan[j+1], daftarKegiatan[j]
			}
		}
	}

	if binarySearchHari(hari) {
		fmt.Println(" Tabrakan ditemukan! Sudah ada kegiatan di hari tersebut.")
	} else {
		fmt.Println(" Tidak ada tabrakan di hari tersebut.")
	}
}

// Rekursif: Menampilkan seluruh data kegiatan
func tampilkanRekursif(index int) {
	if index >= jumlahKegiatanSekarang {
		return
	}
	k := daftarKegiatan[index]
	fmt.Printf("ID: %d | Nama: %s | Hari: %s | Jumlah Peserta: %d\n",
		k.ID, k.NamaKegiatan, k.Hari, k.JumlahPeserta)
	tampilkanRekursif(index + 1)
}

// menampilkan semua kegiatan secara rekursif mulai 
// dari indeks pertama hingga akhir daftar kegiatan
func tampilkanSemuaRekursif() {
	if jumlahKegiatanSekarang == 0 {
		fmt.Println("Belum ada kegiatan yang terdaftar.")
		return
	}
	fmt.Println("\n Daftar Kegiatan (rekursif):")
	tampilkanRekursif(0)
}