package main

import "fmt"

const NMAX int = 999

type Tagihan struct {
	NamaTagihan           string
	Nominal               int
	Tanggal, Bulan, Tahun string
	Status                string
	Kategori              string
}
type tabTagihan [NMAX]Tagihan

func main() {
	var pilihan, n int
	var simtab tabTagihan
	n = 0
	fmt.Println("==========WELCOME TO===========")
	fmt.Println("============SIMTAB=============")

	for pilihan != 9 {
		fmt.Println("\n=== MENU SIMTAB ===")
		fmt.Println("1. Tambah Tagihan")
		fmt.Println("2. Lihat Semua Tagihan")
		fmt.Println("3. Bayar Tagihan")
		fmt.Println("4. Statistik")
		fmt.Println("5. Ubah Tagihan")
		fmt.Println("6. Hapus Tagihan")
		fmt.Println("7. Searching")
		fmt.Println("8. Sorting")
		fmt.Println("9. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahTagihan(&n, &simtab)
			for pilihan != 2 {
				fmt.Println("")
				fmt.Println("=====MENU=====")
				fmt.Println("1. TambahTagihan")
				fmt.Println("2. Kembali")
				fmt.Print("PILIH MENU :")
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					tambahTagihan(&n, &simtab)
				} else if pilihan == 2 {
					fmt.Println("Kembali Ke Menu")
				} else {
					fmt.Println("")
					fmt.Println("Pilihan Tidak Valid")
				}
			}
		} else if pilihan == 2 {
			tampilkanSemua(n, simtab)
			for pilihan != 1 {
				fmt.Println("")
				fmt.Println("==== MENU ====")
				fmt.Println("1. Kembali")
				fmt.Print("Pilih Menu :")
				fmt.Scan(&pilihan)
				if pilihan == 1 {

				} else {
					fmt.Println("")
					fmt.Println("Pilihan Tidak Valid")
				}
			}
		} else if pilihan == 3 {
			bayarTagihan(n, &simtab)
		} else if pilihan == 4 {
			Statistik(n, simtab)
		} else if pilihan == 5 {
			ubahTagihan(n, &simtab)
		} else if pilihan == 6 {
			hapusTagihan(&n, &simtab)
		} else if pilihan == 7 {
			searchingTagihan(n, &simtab)
		} else if pilihan == 8 {
			sortingTagihan(n, &simtab)
		} else if pilihan == 9 {
			fmt.Println("Program selesai.")
		} else {
			fmt.Println("")
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func Statistik(n int, simtab tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan statistik tagihan yang berisi total tagihan, jumlah tagihan yang sudah lunas, jumlah tagihan yang belum lunas, persentase tagihan yang sudah lunas, dan total nominal tagihan yang belum lunas
	var i, jumlahLunas, p int
	var totalBelumTerbayar int
	var persenTerbayar float64

	if n == 0 {
		fmt.Println("")
		fmt.Println("Data Tidak Ada")
	} else {
		jumlahLunas = 0
		totalBelumTerbayar = 0

		for i = 0; i < n; i++ {
			if simtab[i].Status == "LUNAS" {
				jumlahLunas++
			} else {
				totalBelumTerbayar = totalBelumTerbayar + simtab[i].Nominal
			}
		}

		persenTerbayar = float64(jumlahLunas) / float64(n) * 100

		fmt.Println("")
		fmt.Println("============================================")
		fmt.Println("           STATISTIK TAGIHAN               ")
		fmt.Println("============================================")
		fmt.Printf("  Total Tagihan          : %d tagihan\n", n)
		fmt.Printf("  Tagihan Lunas          : %d tagihan\n", jumlahLunas)
		fmt.Printf("  Tagihan Belum Lunas    : %d tagihan\n", n-jumlahLunas)
		fmt.Println("--------------------------------------------")
		fmt.Printf("  Persentase Terbayar    : %.2f%%\n", persenTerbayar)
		fmt.Printf("  Total Nominal Belum    : Rp %d\n", totalBelumTerbayar)
		fmt.Println("============================================")
		fmt.Println("")
		fmt.Println("===MENU===")
		fmt.Println("1. Kembali ke Menu Utama")
		fmt.Print("Pilih Menu :")
		fmt.Scan(&p)
		for p != 1 {
			fmt.Println("PILIHAN TIDAK VALID")
			fmt.Print("Pilih Menu :")
			fmt.Scan(&p)
		}
	}
}
func bayarTagihan(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan semua data tagihan dan melakukan pembayaran tagihan sesuai dengan nomor tagihan yang dipilih dan nominal yang dibayarkan, jika nominal yang dibayarkan kurang dari nominal tagihan maka status tagihan tetap belum lunas, jika nominal yang dibayarkan sama dengan nominal tagihan maka status tagihan berubah menjadi lunas, jika nominal yang dibayarkan lebih dari nominal tagihan maka menampilkan pesan bahwa nominal yang dibayarkan tidak valid
	var idx, p, b, i, jumlahLunas int
	if n == 0 {
		fmt.Println("")
		fmt.Println("Data Tidak Ada")
	} else {
		jumlahLunas = 0
		for i = 0; i < n; i++ {
			if simtab[i].Status == "LUNAS" {
				jumlahLunas++
			}
		}
		if jumlahLunas == n {
			fmt.Println("SEMUA TAGIHAN TELAH LUNAS")
			fmt.Println("")
			fmt.Println("===MENU===")
			fmt.Println("1. Kembali ke Menu Utama")
			fmt.Print("Pilih Menu :")
			fmt.Scan(&p)
			for p != 1 {
				fmt.Println("PILIHAN TIDAK VALID")
				fmt.Print("Pilih Menu :")
				fmt.Scan(&p)
			}
			if p == 1 {

			}
		} else {
			tampilkanSemua(n, *simtab)
			fmt.Println("")
			fmt.Print("Nomor yang dibayar :")
			fmt.Scan(&p)
			for p < 1 || p > n {
				fmt.Println("PILIHAN TIDAK VALID")
				fmt.Print("Nomor yang dibayar :")
				fmt.Scan(&p)
			}
			idx = p - 1
			for simtab[idx].Status == "LUNAS" {
				fmt.Println("TAGIHAN SUDAH LUNAS")
				fmt.Print("Nomor yang dibayar :")
				fmt.Scan(&p)
				for p < 1 || p > n {
					fmt.Println("PILIHAN TIDAK VALID")
					fmt.Print("Nomor yang dibayar :")
					fmt.Scan(&p)
				}
				idx = p - 1
			}
			tampilkanSemua(n, *simtab)
			fmt.Print("Masukkan Nominal yang dibayar :")
			fmt.Scan(&b)
			for b > simtab[idx].Nominal || b <= 0 {
				fmt.Println("Nominal yang dibayar tidak valid. Masukkan nominal yang sesuai.")
				fmt.Print("Masukkan nominal yang dibayar : ")
				fmt.Scan(&b)
			}
			simtab[idx].Nominal = simtab[idx].Nominal - b
			if simtab[idx].Nominal == 0 {
				simtab[idx].Status = "LUNAS"
				simtab[idx].Tanggal = "-"
				simtab[idx].Bulan = "-"
				simtab[idx].Tahun = "-"
			}
			tampilkanSemua(n, *simtab)
			fmt.Println("")
			fmt.Println("===MENU===")
			fmt.Println("1. Kembali ke Menu Utama")
			fmt.Println("2. Kembali ke Menu Bayar Tagihan")
			fmt.Print("Pilih Menu :")
			fmt.Scan(&p)
			for p != 1 && p != 2 {
				fmt.Println("PILIHAN TIDAK VALID")
				fmt.Print("Pilih Menu :")
				fmt.Scan(&p)
			}
			if p == 1 {

			} else if p == 2 {
				bayarTagihan(n, simtab)
			}
		}
	}
}
func searchingTagihan(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sesuai dengan pilihan searching yang dipilih dan menampilkan menu setelah searching
	var p int
	if n == 0 {
		fmt.Println("Data Tidak Ada")
	} else {
		fmt.Println("")
		fmt.Println("Searching Berdasarkan : ")
		fmt.Println("1.Nama Tagihan")
		fmt.Println("2.Nominal")
		fmt.Println("3.Tanggal Jatuh Tempo")
		fmt.Println("4.Kategori")
		fmt.Println("5.Status")
		fmt.Println("6.Kembali")
		fmt.Print("Masukkan Pilihan:")
		fmt.Scan(&p)
		if p == 1 {
			Seq_nama(n, simtab)
		} else if p == 2 {
			Search_nominal(n, simtab)
		} else if p == 3 {
			Seq_JatuhTempo(n, simtab)
		} else if p == 4 {
			Seq_kategori(n, simtab)
		} else if p == 5 {
			Seq_status(n, simtab)
		} else if p == 6 {

		} else {
			fmt.Println("Pilihan Tidak Valid")
			searchingTagihan(n, simtab)
		}
	}
}

func Seq_nama(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sesuai dengan nama tagihan yang ingin dicari, jika ada lebih dari satu data tagihan yang memiliki nama tagihan yang sama dengan nama tagihan yang dicari maka menampilkan semua data tagihan yang memiliki nama tagihan yang sama dengan nama tagihan yang dicari
	var j, pilihan int
	var target string
	var ketemu bool
	j = 0
	ketemu = false
	tampilkanSemua(n, *simtab)
	fmt.Println("")
	fmt.Print("Masukkan Target Nama :")
	fmt.Scan(&target)
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("|%-3s| %-13s | %-15s | %-10s | %-12s | %-6s |\n", "NO", "Nama Tagihan", "Nominal", "Jatuh Tempo", "Kategori", "Status")
	fmt.Println("----------------------------------------------------------------------------------")
	for j < n {
		if (*simtab)[j].NamaTagihan == target {
			fmt.Printf("| %-1d | %-13s | Rp %-12d | %-2s/%-2s/%-5s | %-12s | %-11s |\n", j+1, simtab[j].NamaTagihan, simtab[j].Nominal, simtab[j].Tanggal, simtab[j].Bulan, simtab[j].Tahun, simtab[j].Kategori, simtab[j].Status)
			ketemu = true
		}
		j = j + 1
	}
	if !ketemu {
		fmt.Println("Data Tidak Ditemukan")
	}
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("===MENU===")
	fmt.Println("1. Kembali ke Menu Utama")
	fmt.Println("2. Kembali ke Menu Searching Nama")
	fmt.Println("3. Kembali ke Menu Searching Utama")
	fmt.Print("Pilih Menu :")
	fmt.Scan(&pilihan)
	if pilihan == 1 {

	} else if pilihan == 2 {
		Seq_nama(n, simtab)
	} else if pilihan == 3 {
		searchingTagihan(n, simtab)
	} else {
		fmt.Println("Pilihan Tidak Valid")
	}
}

func Seq_JatuhTempo(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sesuai dengan tanggal jatuh tempo yang ingin dicari, jika ada lebih dari satu data tagihan yang memiliki tanggal jatuh tempo yang sama dengan tanggal jatuh tempo yang dicari maka menampilkan semua data tagihan yang memiliki tanggal jatuh tempo yang sama dengan tanggal jatuh tempo yang dicari
	var j, pilihan int
	var tgl, bln, thn string
	var ketemu bool
	j = 0
	ketemu = false
	tampilkanSemua(n, *simtab)
	fmt.Println("")
	fmt.Print("Masukkan Target JatuhTempo (Tgl Bln Thn) :")
	fmt.Scan(&tgl, &bln, &thn)
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("|%-3s| %-13s | %-15s | %-10s | %-12s | %-11s |\n", "NO", "Nama Tagihan", "Nominal", "Jatuh Tempo", "Kategori", "Status")
	fmt.Println("----------------------------------------------------------------------------------")
	for j < n {
		if ((*simtab)[j].Tanggal == tgl) && ((*simtab)[j].Bulan == bln) && ((*simtab)[j].Tahun == thn) {
			fmt.Printf("| %-1d | %-13s | Rp %-12d | %-2s/%-2s/%-5s | %-12s | %-11s |\n", j+1, simtab[j].NamaTagihan, simtab[j].Nominal, simtab[j].Tanggal, simtab[j].Bulan, simtab[j].Tahun, simtab[j].Kategori, simtab[j].Status)
			ketemu = true
		}
		j = j + 1
	}
	if !ketemu {
		fmt.Println("Data Tidak Ditemukan")
	}
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("===MENU===")
	fmt.Println("1. Kembali ke Menu Utama")
	fmt.Println("2. Kembali ke Menu Searching Utama")
	fmt.Print("Pilih Menu :")
	fmt.Scan(&pilihan)
	if pilihan == 1 {

	} else if pilihan == 2 {
		searchingTagihan(n, simtab)
	} else {
		fmt.Println("Pilihan Tidak Valid")
	}
}

func Search_nominal(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sesuai dengan nominal yang ingin dicari dengan data yang sudah diurutkan, jika ada lebih dari satu data tagihan yang memiliki nominal yang sama dengan nominal yang dicari maka menampilkan semua data tagihan yang memiliki nominal yang sama dengan nominal yang dicari, jika tidak ada data tagihan yang memiliki nominal yang sama dengan nominal yang dicari maka menampilkan pesan bahwa data tidak ditemukan
	var p int
	var target int
	var kr, kn, med int
	var ketemu bool
	ketemu = false
	fmt.Println("")
	fmt.Println("Tipe Searching :")
	fmt.Println("1.Data Tidak Terurut")
	fmt.Println("2.Data Terurut")
	fmt.Println("")
	fmt.Print("Masukan Pilihan :")
	fmt.Scan(&p)
	if p == 1 {
		Seq_Nominal(n, simtab)
	} else if p == 2 {
		Insert_nominal(n, simtab)
		tampilkanSemua(n, *simtab)
		fmt.Println("Masukkan Nominal Target :")
		fmt.Scan(&target)
		fmt.Println("----------------------------------------------------------------------------------")
		fmt.Printf("|%-3s| %-13s | %-15s | %-10s | %-12s | %-11s |\n", "NO", "Nama Tagihan", "Nominal", "Jatuh Tempo", "Kategori", "Status")
		fmt.Println("----------------------------------------------------------------------------------")
		kr = 0
		kn = n - 1

		for kr <= kn && !ketemu {
			med = (kr + kn) / 2

			if (*simtab)[med].Nominal == target {
				ketemu = true
			} else if (*simtab)[med].Nominal > target {
				kn = med - 1
			} else {
				kr = med + 1
			}
		}

		if ketemu {
			kiri := med
			for kiri >= 0 && (*simtab)[kiri].Nominal == target {
				kiri = kiri - 1
			}
			kiri = kiri + 1

			for kiri < n && (*simtab)[kiri].Nominal == target {
				fmt.Printf("| %-1d | %-13s | Rp %-12d | %-2s/%-2s/%-5s | %-12s | %-11s |\n", kiri+1, simtab[kiri].NamaTagihan, simtab[kiri].Nominal, simtab[kiri].Tanggal, simtab[kiri].Bulan, simtab[kiri].Tahun, simtab[kiri].Kategori, simtab[kiri].Status)
				kiri = kiri + 1
			}
		} else {
			fmt.Println("Data Tidak Ditemukan")
		}
		fmt.Println("----------------------------------------------------------------------------------")
		fmt.Println("")
		fmt.Println("===MENU===")
		fmt.Println("1. Kembali ke Menu Utama")
		fmt.Println("2. Kembali ke Menu Searching Utama")
		fmt.Print("Pilih Menu :")
		fmt.Scan(&p)
		if p == 1 {

		} else if p == 2 {
			searchingTagihan(n, simtab)
		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}
}

func Seq_Nominal(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sesuai dengan nominal yang ingin dicari, jika ada lebih dari satu data tagihan yang memiliki nominal yang sama dengan nominal yang dicari maka menampilkan semua data tagihan yang memiliki nominal yang sama dengan nominal yang dicari, jika tidak ada data tagihan yang memiliki nominal yang sama dengan nominal yang dicari maka menampilkan pesan bahwa data tidak ditemukan
	var j, pilihan int
	var target int
	var ketemu bool
	j = 0
	ketemu = false
	tampilkanSemua(n, *simtab)
	fmt.Println("")
	fmt.Print("Masukkan Target Nominal :")
	fmt.Scan(&target)
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("|%-3s| %-13s | %-15s | %-10s | %-12s | %-11s |\n", "NO", "Nama Tagihan", "Nominal", "Jatuh Tempo", "Kategori", "Status")
	fmt.Println("----------------------------------------------------------------------------------")
	for j < n {
		if (*simtab)[j].Nominal == target {
			fmt.Printf("| %-1d | %-13s | Rp %-12d | %-2s/%-2s/%-5s | %-12s | %-11s |\n", j+1, simtab[j].NamaTagihan, simtab[j].Nominal, simtab[j].Tanggal, simtab[j].Bulan, simtab[j].Tahun, simtab[j].Kategori, simtab[j].Status)
			ketemu = true
		}
		j = j + 1
	}
	if !ketemu {
		fmt.Println("Data Tidak Ditemukan")
	}
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("===MENU===")
	fmt.Println("1. Kembali ke Menu Utama")
	fmt.Println("2. Kembali ke Menu Searching Nominal")
	fmt.Println("3. Kembali ke Menu Searching Utama")
	fmt.Print("Pilih Menu :")
	fmt.Scan(&pilihan)
	if pilihan == 1 {

	} else if pilihan == 2 {
		Search_nominal(n, simtab)
	} else if pilihan == 3 {
		searchingTagihan(n, simtab)
	} else {
		fmt.Println("Pilihan Tidak Valid")
	}
}

func Seq_kategori(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sesuai dengan kategori yang ingin dicari, jika ada lebih dari satu data tagihan yang memiliki kategori yang sama dengan kategori yang dicari maka menampilkan semua data tagihan yang memiliki kategori yang sama dengan kategori yang dicari, jika tidak ada data tagihan yang memiliki kategori yang sama dengan kategori yang dicari maka menampilkan pesan bahwa data tidak ditemukan
	var j, pilihan int
	var target string
	var ketemu bool
	j = 0
	ketemu = false
	tampilkanSemua(n, *simtab)
	fmt.Println("")
	fmt.Print("Masukkan Target Kategori (Pribadi/RumahTangga) :")
	fmt.Scan(&target)
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("|%-3s| %-13s | %-15s | %-10s | %-12s | %-11s |\n", "NO", "Nama Tagihan", "Nominal", "Jatuh Tempo", "Kategori", "Status")
	fmt.Println("----------------------------------------------------------------------------------")
	for j < n {
		if (*simtab)[j].Kategori == target {
			fmt.Printf("| %-1d | %-13s | Rp %-12d | %-2s/%-2s/%-5s | %-12s | %-11s |\n", j+1, simtab[j].NamaTagihan, simtab[j].Nominal, simtab[j].Tanggal, simtab[j].Bulan, simtab[j].Tahun, simtab[j].Kategori, simtab[j].Status)
			ketemu = true
		}
		j = j + 1
	}
	if !ketemu {
		fmt.Println("Data Tidak Ditemukan")
	}
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("===MENU===")
	fmt.Println("1. Kembali ke Menu Utama")
	fmt.Println("2. Kembali ke Menu Searching Kategori")
	fmt.Println("3. Kembali ke Menu Searching Utama")
	fmt.Print("Pilih Menu :")
	fmt.Scan(&pilihan)
	if pilihan == 1 {

	} else if pilihan == 2 {
		Seq_kategori(n, simtab)
	} else if pilihan == 3 {
		searchingTagihan(n, simtab)
	} else {
		fmt.Println("Pilihan Tidak Valid")
	}
}

func Seq_status(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sesuai dengan status yang ingin dicari, jika ada lebih dari satu data tagihan yang memiliki status yang sama dengan status yang dicari maka menampilkan semua data tagihan yang memiliki status yang sama dengan status yang dicari, jika tidak ada data tagihan yang memiliki status yang sama dengan status yang dicari maka menampilkan pesan bahwa data tidak ditemukan
	var j, pilihan int
	var target string
	var ketemu bool
	j = 0
	ketemu = false
	tampilkanSemua(n, *simtab)
	fmt.Println("")
	fmt.Print("Masukkan Target Status (y/n) :")
	fmt.Scan(&target)
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("|%-3s| %-13s | %-15s | %-10s | %-12s | %-11s |\n", "NO", "Nama Tagihan", "Nominal", "Jatuh Tempo", "Kategori", "Status")
	fmt.Println("----------------------------------------------------------------------------------")
	for j < n {
		if (*simtab)[j].Status == target {
			fmt.Printf("| %-1d | %-13s | Rp %-12d | %-2s/%-2s/%-5s | %-12s | %-11s |\n", j+1, simtab[j].NamaTagihan, simtab[j].Nominal, simtab[j].Tanggal, simtab[j].Bulan, simtab[j].Tahun, simtab[j].Kategori, simtab[j].Status)
			ketemu = true
		}
		j = j + 1
	}
	if !ketemu {
		fmt.Println("Data Tidak Ditemukan")
	}
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("===MENU===")
	fmt.Println("1. Kembali ke Menu Utama")
	fmt.Println("2. Kembali ke Menu Searching Status")
	fmt.Println("3. Kembali ke Menu Searching Utama")
	fmt.Print("Pilih Menu :")
	fmt.Scan(&pilihan)
	if pilihan == 1 {

	} else if pilihan == 2 {
		Seq_status(n, simtab)
	} else if pilihan == 3 {
		searchingTagihan(n, simtab)
	} else {
		fmt.Println("Pilihan Tidak Valid")
	}
}

func sortingTagihan(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sudah diurutkan sesuai dengan pilihan sorting yang dipilih, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var p int
	if n == 0 {
		fmt.Println("Data Tidak Ada")
	} else {
		fmt.Println("")
		fmt.Println("Sorting Berdasarkan : ")
		fmt.Println("1.Nama Tagihan")
		fmt.Println("2.Nominal")
		fmt.Println("3.Tanggal Jatuh Tempo")
		fmt.Println("4.Kategori")
		fmt.Println("5.Status")
		fmt.Println("6.Kembali")
		fmt.Print("Masukkan Pilihan:")
		fmt.Scan(&p)
		if p == 1 {
			Insert_nama(n, simtab)
			after_sort(n, simtab)
		} else if p == 2 {
			Insert_nominal(n, simtab)
			after_sort(n, simtab)
		} else if p == 3 {
			Insert_jatuhTempo(n, simtab)
			after_sort(n, simtab)
		} else if p == 4 {
			selec_kategori(n, simtab)
			after_sort(n, simtab)
		} else if p == 5 {
			selec_status(n, simtab)
			after_sort(n, simtab)
		} else if p == 6 {

		} else {
			fmt.Println("Pilihan Tidak Valid")
			sortingTagihan(n, simtab)
		}
	}
}

func after_sort(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan data tagihan yang sudah diurutkan sesuai dengan pilihan sorting yang dipilih dan menampilkan menu setelah sorting, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var pilihan int
	tampilkanSemua(n, *simtab)
	fmt.Println("")
	fmt.Println("==== MENU ====")
	fmt.Println("1. Kembali ke Menu Utama")
	fmt.Println("2. Kembali ke Menu Sorting Utama")
	fmt.Print("Pilih Menu :")
	fmt.Scan(&pilihan)
	if pilihan == 1 {

	} else if pilihan == 2 {
		sortingTagihan(n, simtab)
	} else {
		fmt.Println("Pilihan Tidak Valid")
	}
}

func Insert_nama(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Mengurutkan nama tagihan secara ascending atau descending sesuai dengan pilihan yang dipilih, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var pass, k, p int
	var temp Tagihan
	fmt.Println("")
	fmt.Println("=====MENU SORT=====")
	fmt.Println("1.A-Z")
	fmt.Println("2.Z-A")
	fmt.Print("Masukkan Pilihan:")
	fmt.Scan(&p)
	if p == 1 {
		for pass = 1; pass < n; pass++ {
			k = pass
			temp = simtab[k]
			for k > 0 && temp.NamaTagihan < simtab[k-1].NamaTagihan {
				simtab[k] = simtab[k-1]
				k = k - 1
			}
			simtab[k] = temp
		}
	} else if p == 2 {
		for pass = 1; pass < n; pass++ {
			k = pass
			temp = simtab[k]
			for k > 0 && temp.NamaTagihan > simtab[k-1].NamaTagihan {
				simtab[k] = simtab[k-1]
				k = k - 1
			}
			simtab[k] = temp
		}
	} else {
		fmt.Println("")
		fmt.Println("Pilihan Tidak Valid")
	}
}

func Insert_nominal(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Mengurutkan nominal tagihan secara ascending atau descending sesuai dengan pilihan yang dipilih, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var pass, k, p int
	var temp Tagihan
	fmt.Println("")
	fmt.Println("=====MENU SORT=====")
	fmt.Println("1.Nominal Terkecil - Terbesar")
	fmt.Println("2.Nominal Terbesar - Terkecil")
	fmt.Print("Masukkan Pilihan:")
	fmt.Scan(&p)
	if p == 1 {
		for pass = 1; pass < n; pass++ {
			k = pass
			temp = simtab[k]
			for k > 0 && temp.Nominal < simtab[k-1].Nominal {
				simtab[k] = simtab[k-1]
				k = k - 1
			}
			simtab[k] = temp
		}
	} else if p == 2 {
		for pass = 1; pass < n; pass++ {
			k = pass
			temp = simtab[k]
			for k > 0 && temp.Nominal > simtab[k-1].Nominal {
				simtab[k] = simtab[k-1]
				k = k - 1
			}
			simtab[k] = temp
		}
	} else {
		fmt.Println("")
		fmt.Println("Pilihan Tidak Valid")
	}
}

func Insert_jatuhTempo(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Mengurutkan tanggal jatuh tempo tagihan secara ascending atau descending sesuai dengan pilihan yang dipilih, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var pass, k, p int
	var temp Tagihan
	fmt.Println("")
	fmt.Println("=====MENU SORT=====")
	fmt.Println("1.Tercepat - Terlama")
	fmt.Println("2.Terlama - Tercepat")
	fmt.Print("Masukkan Pilihan:")
	fmt.Scan(&p)
	if p == 1 {
		for pass = 1; pass < n; pass++ {
			k = pass
			temp = simtab[k]
			for k > 0 && (temp.Tahun < simtab[k-1].Tahun || (temp.Tahun == simtab[k-1].Tahun && temp.Bulan < simtab[k-1].Bulan) || (temp.Tahun == simtab[k-1].Tahun && temp.Bulan == simtab[k-1].Bulan && temp.Tanggal < simtab[k-1].Tanggal)) {
				simtab[k] = simtab[k-1]
				k = k - 1
			}
			simtab[k] = temp
		}
	} else if p == 2 {
		for pass = 1; pass < n; pass++ {
			k = pass
			temp = simtab[k]
			for k > 0 && (temp.Tahun > simtab[k-1].Tahun || (temp.Tahun == simtab[k-1].Tahun && temp.Bulan > simtab[k-1].Bulan) || (temp.Tahun == simtab[k-1].Tahun && temp.Bulan == simtab[k-1].Bulan && temp.Tanggal > simtab[k-1].Tanggal)) {
				simtab[k] = simtab[k-1]
				k = k - 1
			}
			simtab[k] = temp
		}
	} else {
		fmt.Println("")
		fmt.Println("Pilihan Tidak Valid")
	}
}

func selec_kategori(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Mengurutkan kategori tagihan secara ascending atau descending sesuai dengan pilihan yang dipilih, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var pass, k, p, acuan int
	var temp Tagihan
	fmt.Println("")
	fmt.Println("=====MENU SORT=====")
	fmt.Println("1.Pribadi-RumahTangga")
	fmt.Println("2.RumahTangga-Pribadi")
	fmt.Print("Masukkan Pilihan:")
	fmt.Scan(&p)
	if p == 1 {
		for pass = 1; pass < n; pass++ {
			acuan = pass - 1
			for k = pass; k < n; k++ {
				if simtab[acuan].Kategori > simtab[k].Kategori {
					acuan = k
				}
			}
			temp = simtab[acuan]
			simtab[acuan] = simtab[pass-1]
			simtab[pass-1] = temp
		}
	} else if p == 2 {
		for pass = 1; pass < n; pass++ {
			acuan = pass - 1
			for k = pass; k < n; k++ {
				if simtab[acuan].Kategori < simtab[k].Kategori {
					acuan = k
				}
			}
			temp = simtab[acuan]
			simtab[acuan] = simtab[pass-1]
			simtab[pass-1] = temp
		}
	} else {
		fmt.Println("")
		fmt.Println("Pilihan Tidak Valid")
	}
}

func selec_status(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Mengurutkan status tagihan secara ascending atau descending sesuai dengan pilihan yang dipilih, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var pass, k, p, acuan int
	var temp Tagihan
	fmt.Println("")
	fmt.Println("=====MENU SORT=====")
	fmt.Println("1. Belum Lunas - Lunas")
	fmt.Println("2. Lunas - Belum Lunas")
	fmt.Print("Masukkan Pilihan:")
	fmt.Scan(&p)
	if p == 1 {
		for pass = 1; pass < n; pass++ {
			acuan = pass - 1
			for k = pass; k < n; k++ {
				if simtab[acuan].Status > simtab[k].Status {
					acuan = k
				}
			}
			temp = simtab[acuan]
			simtab[acuan] = simtab[pass-1]
			simtab[pass-1] = temp
		}
	} else if p == 2 {
		for pass = 1; pass < n; pass++ {
			acuan = pass - 1
			for k = pass; k < n; k++ {
				if simtab[acuan].Status < simtab[k].Status {
					acuan = k
				}
			}
			temp = simtab[acuan]
			simtab[acuan] = simtab[pass-1]
			simtab[pass-1] = temp
		}
	} else {
		fmt.Println("")
		fmt.Println("Pilihan Tidak Valid")
	}
}

func ubahTagihan(n int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Mengubah data tagihan sesuai dengan nomor tagihan yang dipilih dan data tagihan yang diubah sesuai dengan inputan terbaru, jika nomor tagihan yang dipilih tidak valid maka menampilkan pesan bahwa pilihan tidak valid, jika data tagihan yang dipilih memiliki status LUNAS maka menampilkan pesan bahwa data tagihan tidak bisa diubah, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var idx, p int
	if n == 0 {
		fmt.Println("Data Tidak Ada")
	} else {
		tampilkanSemua(n, *simtab)
		fmt.Print("Nomor yang diubah :")
		fmt.Scan(&p)
		if p > n || p <= 0 {
			fmt.Println("")
			fmt.Println("Pilihan Tidak Valid")
		} else {
			idx = p - 1
			if simtab[idx].Status == "LUNAS" {
				fmt.Println("TAGIHAN LUNAS, DATA TIDAK BISA DIUBAH")
			} else {
				fmt.Println("Nama Tagihan Lama :", simtab[idx].NamaTagihan)
				fmt.Print("Nama Tagihan Baru :")
				fmt.Scan(&simtab[idx].NamaTagihan)
				fmt.Println("Nominal Lama :", simtab[idx].Nominal)
				fmt.Print("Nominal Baru :")
				fmt.Scan(&simtab[idx].Nominal)
				fmt.Println("JatuhTempo Lama :", simtab[idx].Tanggal, simtab[idx].Bulan, simtab[idx].Tahun)
				fmt.Print("JatuhTempo Baru (Tgl Bln Thn) :")
				fmt.Scan(&simtab[idx].Tanggal)
				fmt.Scan(&simtab[idx].Bulan)
				fmt.Scan(&simtab[idx].Tahun)
				fmt.Println("Kategori Lama :", simtab[idx].Kategori)
				fmt.Print("Kategori Baru (Pribadi/RumahTangga) :")
				fmt.Scan(&simtab[idx].Kategori)
				fmt.Println("Status Lama :", simtab[idx].Status)
				simtab[idx].Status = "BELUM LUNAS"
			}
		}
	}
}
func tambahTagihan(n *int, simtab *tabTagihan) {
	//I.S n dan simtab kosong
	//F.S Menambahkan data tagihan sesuai dengan inputan data tagihan, berupa nama tagihan, nominal, tanggal jatuh tempo, kategori, dan status yang sudah ditentukan
	fmt.Print("Nama Tagihan :")
	fmt.Scan(&simtab[*n].NamaTagihan)
	fmt.Print("Nominal :")
	fmt.Scan(&simtab[*n].Nominal)
	fmt.Print("Tgl JatuhTempo (Tgl Bln Thn) :")
	fmt.Scan(&simtab[*n].Tanggal)
	fmt.Scan(&simtab[*n].Bulan)
	fmt.Scan(&simtab[*n].Tahun)
	fmt.Print("Kategori (Pribadi/RumahTangga) :")
	fmt.Scan(&simtab[*n].Kategori)
	simtab[*n].Status = "BELUM LUNAS"
	*n = *n + 1
}

func tampilkanSemua(n int, simtab tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menampilkan semua data tagihan yang ada, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var i int
	fmt.Println("")
	if n == 0 {
		fmt.Println("Data Tidak Ada")
	} else {
		fmt.Println("----------------------------------------------------------------------------------")
		fmt.Printf("|%-3s| %-13s | %-15s | %-10s | %-12s | %-11s |\n", "NO", "Nama Tagihan", "Nominal", "Jatuh Tempo", "Kategori", "Status")
		fmt.Println("----------------------------------------------------------------------------------")
		for i = 0; i < n; i++ {
			fmt.Printf("| %-1d | %-13s | Rp %-12d | %-2s/%-2s/%-5s | %-12s | %-11s |\n", i+1, simtab[i].NamaTagihan, simtab[i].Nominal, simtab[i].Tanggal, simtab[i].Bulan, simtab[i].Tahun, simtab[i].Kategori, simtab[i].Status)
		}
		fmt.Println("----------------------------------------------------------------------------------")
	}
}

func hapusTagihan(n *int, simtab *tabTagihan) {
	//I.S terdefinisi n dan simtab berisi data tagihan
	//F.S Menghapus data tagihan sesuai dengan nomor tagihan yang dipilih, jika nomor tagihan yang dipilih tidak valid maka menampilkan pesan bahwa pilihan tidak valid, jika data tagihan yang dipilih memiliki status LUNAS maka menampilkan pesan bahwa data tagihan tidak bisa dihapus, jika tidak ada data tagihan yang tersedia maka menampilkan pesan bahwa data tidak ada
	var p, idx, i, pilihan int
	if *n == 0 {
		fmt.Println("")
		fmt.Println("DATA TIDAK ADA")
		for pilihan != 1 {
			fmt.Println("")
			fmt.Println("===MENU===")
			fmt.Println("1. Kembali")
			fmt.Print("Pilih Menu :")
			fmt.Scan(&pilihan)
			if pilihan == 1 {

			} else {
				fmt.Println("Pilihan Tidak Valid")
			}
		}
	} else {
		tampilkanSemua(*n, *simtab)
		fmt.Print("Nomor yang dihapus :")
		fmt.Scan(&p)
		idx = p - 1
		if idx >= 0 && idx < *n {
			if simtab[idx].Status == "LUNAS" {
				fmt.Println("TAGIHAN LUNAS, DATA TIDAK BISA DIHAPUS")
			} else {
				for i = idx; i < *n-1; i++ {
					simtab[i] = simtab[i+1]
				}
				*n = *n - 1
				fmt.Println("DATA BERHASIL DIHAPUS")
			}
		} else {
			fmt.Println("NOMOR TIDAK VALID")
		}
	}
}
