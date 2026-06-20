package main

import "fmt"

const NMAX int = 9999

type peserta struct {
	id                               int
	nama, bidangMinat, tanggalDaftar string
	aktif                            bool
}

type tabPeserta [NMAX]peserta

func tambahPeserta(A *tabPeserta, n *int) {

	fmt.Print("ID: ")
	fmt.Scan(&(*A)[*n].id)

	if cekID(*A, *n, (*A)[*n].id) {
		fmt.Println("ID sudah digunakan")
	} else {
		fmt.Print("Nama: ")
		fmt.Scan(&(*A)[*n].nama)

		fmt.Print("Bidang minat: ")
		fmt.Scan(&(*A)[*n].bidangMinat)

		fmt.Print("Tanggal pendaftaran: ")
		fmt.Scan(&(*A)[*n].tanggalDaftar)

		fmt.Print("Status keaktifan: ")
		fmt.Scan(&(*A)[*n].aktif)
		*n = *n + 1
	}
}

func cekID(A tabPeserta, n int, id int) bool {
	var i int
	for i = 0; i < n; i++ {
		if A[i].id == id {
			return true
		}
	}
	return false
}

func editPeserta(A *tabPeserta, n int) {
	var idCari, idx int

	fmt.Print("Masukkan ID peserta yang ingin dicari: ")
	fmt.Scan(&idCari)
	idx = -1
	for i := 0; i < n; i++ {
		if (*A)[i].id == idCari {
			idx = i
		}
	}
	fmt.Print()

	if idx != -1 {
		fmt.Println("ID ditemukan: ")
		fmt.Print()

		fmt.Print("Nama baru: ")
		fmt.Scan(&(*A)[idx].nama)

		fmt.Print("Bidang minat baru: ")
		fmt.Scan(&(*A)[idx].bidangMinat)

		fmt.Print("Tanggal daftar baru: ")
		fmt.Scan(&(*A)[idx].tanggalDaftar)

		fmt.Print("Status keaktifan: ")
		fmt.Scan(&(*A)[idx].aktif)

		fmt.Print()
		fmt.Print("Data berhasil diubah")
	} else {
		fmt.Println("ID tidak ditemukan")
	}
}

func hapusPeserta(A *tabPeserta, n *int) {
	var idCari, idx, i int

	fmt.Print("Masukkan ID peserta yang ingin dihapus: ")
	fmt.Scan(&idCari)

	idx = -1
	for i = 0; i < *n; i++ {
		if (*A)[i].id == idCari {
			idx = i
		}
	}
	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			(*A)[i] = (*A)[i+1]
		}
		*n = *n - 1
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("ID tidak ditemukan")
	}
}

func seqSearchNama(A tabPeserta, n int, cari string) {
	var ketemu bool
	var i int

	for i = 0; i < n; i++ {
		if A[i].nama == cari {
			fmt.Println(A[i])
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func seqSearchMinat(A tabPeserta, n int, cari string) {
	var ketemu bool

	for i := 0; i < n; i++ {
		if A[i].bidangMinat == cari {
			fmt.Println(A[i])
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func binSearchNama(A tabPeserta, n int, cari string) int {
	var left, right, mid int
	var found int

	left = 0
	right = n - 1
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if cari > A[mid].nama {
			left = mid + 1
		} else if cari < A[mid].nama {
			right = mid - 1
		} else {
			found = mid
		}
	}
	return found
}
func binSearchMinat(A tabPeserta, n int, cari string) int {
	var left, right, mid int
	var found int

	left = 0
	right = n - 1
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if cari > A[mid].bidangMinat {
			left = mid + 1
		} else if cari < A[mid].bidangMinat {
			right = mid - 1
		} else {
			found = mid
		}
	}
	return found
}

func insertionSortAscendingID(A *tabPeserta, n int) {
	var pass, i int
	var temp peserta

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*A)[pass]

		for i > 0 && temp.id < (*A)[i-1].id {
			(*A)[i] = (*A)[i-1]
			i--
		}

		(*A)[i] = temp
		pass++
	}
}

func insertionSortDesendingID(A *tabPeserta, n int) {
	var pass, i int
	var temp peserta

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*A)[pass]

		for i > 0 && temp.id > (*A)[i-1].id {
			(*A)[i] = (*A)[i-1]
			i--
		}

		(*A)[i] = temp
		pass++
	}
}

func selectionSortAscendingID(A *tabPeserta, n int) {
	var pass, i, idx int
	var temp peserta

	pass = 0
	for pass < n-1 {
		idx = pass

		i = pass + 1
		for i < n {
			if (*A)[i].id < (*A)[idx].id {
				idx = i
			}
			i++
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp

		pass++
	}
}

func selectionSortDesendingID(A *tabPeserta, n int) {
	var pass, i, idx int
	var temp peserta

	pass = 0
	for pass < n-1 {
		idx = pass

		i = pass + 1
		for i < n {
			if (*A)[i].id > (*A)[idx].id {
				idx = i
			}
			i++
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp

		pass++
	}
}

func insertionSortAscendingNama(A *tabPeserta, n int) {
	var pass, i int
	var temp peserta

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*A)[pass]

		for i > 0 && temp.nama < (*A)[i-1].nama {
			(*A)[i] = (*A)[i-1]
			i--
		}

		(*A)[i] = temp
		pass++
	}
}

func insertionSortDesendingNama(A *tabPeserta, n int) {
	var pass, i int
	var temp peserta

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*A)[pass]

		for i > 0 && temp.nama > (*A)[i-1].nama {
			(*A)[i] = (*A)[i-1]
			i--
		}

		(*A)[i] = temp
		pass++
	}
}

func selectionSortAscendingNama(A *tabPeserta, n int) {
	var pass, i, idx int
	var temp peserta

	pass = 0
	for pass < n-1 {
		idx = pass

		i = pass + 1
		for i < n {
			if (*A)[i].nama < (*A)[idx].nama {
				idx = i
			}
			i++
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp

		pass++
	}
}

func selectionSortDesendingNama(A *tabPeserta, n int) {
	var pass, i, idx int
	var temp peserta

	pass = 0
	for pass < n-1 {
		idx = pass

		i = pass + 1
		for i < n {
			if (*A)[i].nama > (*A)[idx].nama {
				idx = i
			}
			i++
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp

		pass++
	}
}

func selectionSortAscendingBidangMinat(A *tabPeserta, n int) {
	var pass, i, idx int
	var temp peserta

	pass = 0
	for pass < n-1 {
		idx = pass

		i = pass + 1
		for i < n {
			if (*A)[i].bidangMinat < (*A)[idx].bidangMinat {
				idx = i
			}
			i++
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp

		pass++
	}
}

func statistik(A tabPeserta, n int) {
	var aktif, i, j int
	var pertama bool

	fmt.Println("Statistik")

	for i = 0; i < n; i++ {

		//menghitung peserta aktif
		if A[i].aktif {
			aktif++
		}

		pertama = true

		// Kalo bidang minat udah pernah ketemu
		// gak perlu dihitung dan ditampilkan lagi
		for j = 0; j < i; j++ {
			if A[j].bidangMinat == A[i].bidangMinat {
				pertama = false
			}
		}
		// cuma bidang minat yang muncul pertama kaliyang akan ditampilin
		if pertama {
			jumlah := 0

			for j = 0; j < n; j++ {
				if A[j].bidangMinat == A[i].bidangMinat {
					jumlah++
				}
			}

			fmt.Println(A[i].bidangMinat, ":", jumlah)
		}
	}
	fmt.Println("Total peserta aktif:", aktif)
}

func main() {
	var data tabPeserta
	var n int
	var temp tabPeserta
	var jalan bool

	jalan = true

	for jalan {
		var pilihan int

		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Edit Peserta")
		fmt.Println("3. Hapus Peserta")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("5. Mencari Data")
		fmt.Println("6. Mengurutkan Data")
		fmt.Println("7. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahPeserta(&data, &n)

		case 2:
			editPeserta(&data, n)

		case 3:
			hapusPeserta(&data, &n)

		case 4:
			fmt.Println("Tampilkan data: ")
			fmt.Println("1. Awal")
			fmt.Println("2. Setelah sort")
			fmt.Print("Pilihan: ")

			var tampilData int
			fmt.Scan(&tampilData)

			switch tampilData {
			case 1:
				for i := 0; i < n; i++ {
					fmt.Println(data[i].id, data[i].nama, data[i].bidangMinat, data[i].tanggalDaftar, data[i].aktif)
				}
			case 2:
				for i := 0; i < n; i++ {
					fmt.Println(temp[i].id, temp[i].nama, temp[i].bidangMinat, temp[i].tanggalDaftar, temp[i].aktif)
				}
			}

		case 5:
			fmt.Println("Mencari data berdasarkan: ")
			fmt.Println("1. Bidang minat")
			fmt.Println("2. Nama lengkap")
			fmt.Print("Pilihan: ")
			temp = data

			var caridata int
			var dicari string
			var jenisSearch int
			var bin int

			fmt.Scan(&caridata)

			switch caridata {

			case 1:
				fmt.Println("Mencari Menggunakan: ")
				fmt.Println("1. Sequential search")
				fmt.Println("2. Binary search")
				fmt.Print("Metode: ")
				fmt.Scan(&jenisSearch)

				if jenisSearch == 1 {
					fmt.Print("Bidang minat yang dicari: ")
					fmt.Scan(&dicari)
					seqSearchMinat(data, n, dicari)

				} else if jenisSearch == 2 {
					fmt.Print("Bidang minat yang dicari: ")
					fmt.Scan(&dicari)

					selectionSortAscendingBidangMinat(&temp, n)
					bin = binSearchMinat(temp, n, dicari)
					if bin != -1 {
						for i := 0; i < n; i++ {
							if temp[i].bidangMinat == dicari {
								fmt.Println(temp[i])
							}
						}
					} else {
						fmt.Println("Data tidak ditemukan")
					}
				}
			case 2:
				fmt.Println("Mencari Menggunakan: ")
				fmt.Println("1. Sequential search")
				fmt.Println("2. Binary search")
				fmt.Println("Metode: ")
				fmt.Scan(&jenisSearch)

				if jenisSearch == 1 {
					fmt.Print("Nama yang dicari: ")
					fmt.Scan(&dicari)
					seqSearchNama(data, n, dicari)

				} else if jenisSearch == 2 {
					fmt.Print("Nama yang dicari: ")
					fmt.Scan(&dicari)

					selectionSortAscendingNama(&temp, n)
					bin = binSearchNama(temp, n, dicari)
					if bin != -1 {
						for i := 0; i < n; i++ {
							if temp[i].nama == dicari {
								fmt.Println(temp[i])
							}
						}
					} else {
						fmt.Println("Data tidak ditemukan")
					}
				}
			}

		case 6:
			temp = data
			fmt.Println("Mengurutkan data berdasarkan: ")
			fmt.Println("1. ID")
			fmt.Println("2. Nama")
			fmt.Print("Pilihan: ")

			var caraUrut int
			fmt.Scan(&caraUrut)

			switch caraUrut {
			case 1:
				fmt.Println("Mengurutkan menggunakan: ")
				fmt.Println("1. Selection")
				fmt.Println("2. Insertion")
				fmt.Print("Metode: ")

				var jenisSort int

				fmt.Scan(&jenisSort)
				if jenisSort == 1 {
					var urutan string

					fmt.Print("Urutan menaik/menurun: ")
					fmt.Scan(&urutan)

					if urutan == "Ascending" || urutan == "ascending" || urutan == "menaik" {
						selectionSortAscendingID(&temp, n)
					} else if urutan == "Descending" || urutan == "descending" || urutan == "menurun" {
						selectionSortDesendingID(&temp, n)
					} else {
						fmt.Println("Metode tidak valid")
					}
				} else if jenisSort == 2 {
					var urutan string

					fmt.Print("Urutan menaik/menurun: ")
					fmt.Scan(&urutan)

					if urutan == "Ascending" || urutan == "ascending" || urutan == "menaik" {
						insertionSortAscendingID(&temp, n)
					} else if urutan == "Descending" || urutan == "descending" || urutan == "menurun" {
						insertionSortDesendingID(&temp, n)
					} else {
						fmt.Println("Metode tidak valid")
					}
				}

			case 2:
				fmt.Println("Mengurutkan menggunakan: ")
				fmt.Println("1. Selection")
				fmt.Println("2. Insertion")
				fmt.Print("Metode: ")

				var jenisSort int

				fmt.Scan(&jenisSort)
				if jenisSort == 1 {
					var urutan string

					fmt.Print("Urutan menaik/menurun: ")
					fmt.Scan(&urutan)

					if urutan == "Ascending" || urutan == "ascending" || urutan == "menaik" {
						selectionSortAscendingNama(&temp, n)
					} else if urutan == "Descending" || urutan == "descending" || urutan == "menurun" {
						selectionSortDesendingNama(&temp, n)
					} else {
						fmt.Println("Metode tidak valid")
					}
				} else if jenisSort == 2 {
					var urutan string

					fmt.Print("Urutan menaik/menurun: ")
					fmt.Scan(&urutan)

					if urutan == "Ascending" || urutan == "ascending" || urutan == "menaik" {
						insertionSortAscendingNama(&temp, n)
					} else if urutan == "Descending" || urutan == "descending" || urutan == "menurun" {
						insertionSortDesendingNama(&temp, n)
					} else {
						fmt.Println("Metode tidak valid")
					}
				}
			}

		case 7:
			statistik(data, n)

		case 0:
			fmt.Println("Program selesai")
			fmt.Println("===END===")
			jalan = false

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

//satu
//1. Jika input ID selain int itu error
//2. Jika ID sama juga error

//2
//binary belum urutan
//
