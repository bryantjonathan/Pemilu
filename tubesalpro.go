package main

import "fmt"

const NMAX int = 500
const threshold int = 500

// Struct untuk menyimpan data calon legislatif
type candidate struct {
	nama, nama_partai, jenis_kelamin, umur, id string
	banyak_pemilih                             int
}

// Struct untuk menyimpan data pemilih
type voter struct {
	nama, jenis_kelamin, NIK, umur, pilihan, partai_kandidat string
}

type arrCandidate [NMAX]candidate

type arrVoter [NMAX]voter

// Prosedur untuk menambah data calon legislatif
func addCandidate(C *arrCandidate, n *int) {
	var c candidate

	fmt.Print("Nama           : ")
	fmt.Scanln(&c.nama)
	fmt.Print("Nama Partai    : ")
	fmt.Scanln(&c.nama_partai)
	fmt.Print("Jenis Kelamin  : ")
	fmt.Scanln(&c.jenis_kelamin)
	fmt.Print("Umur           : ")
	fmt.Scanln(&c.umur)
	fmt.Print("ID             : ")
	fmt.Scanln(&c.id)

	if *n < NMAX {
		C[*n] = c
		*n++
		fmt.Println("Data calon telah ditambahkan.")
	} else {
		fmt.Println("Batas maksimum calon telah tercapai.")
	}
}

// Prosedur untuk mengubah (edit) data calon legislatif berdasarkan ID
func editCandidate(C *arrCandidate, n int) {
	var id string
	var i int = 0
	var found bool = false

	fmt.Print("Masukkan ID calon yang ingin diubah: ")
	fmt.Scanln(&id)

	for i < n && !found {
		if C[i].id == id {
			fmt.Print("Nama           : ")
			fmt.Scanln(&C[i].nama)
			fmt.Print("Nama Partai    : ")
			fmt.Scanln(&C[i].nama_partai)
			fmt.Print("Jenis Kelamin  : ")
			fmt.Scanln(&C[i].jenis_kelamin)
			fmt.Print("Umur           : ")
			fmt.Scanln(&C[i].umur)

			fmt.Println("Data calon telah diubah.")
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Calon tidak ditemukan.")
	}
}

// Prosedur untuk menghapus data calon legislatif berdasarkan ID
func deleteCandidate(C *arrCandidate, n *int) {
	var j int
	var id string
	var i int = 0
	var found bool = false

	fmt.Print("Masukkan ID calon yang ingin dihapus: ")
	fmt.Scanln(&id)

	for i < *n && !found {
		if C[i].id == id {
			for j = i; j < *n-1; j++ {
				C[j] = C[j+1]
			}
			*n--
			fmt.Println("Data calon telah dihapus.")
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Calon tidak ditemukan.")
	}
}

// Prosedur untuk menampilkan data calon legislatif
func showCandidates(C arrCandidate, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("-----------------------------")
		fmt.Println("ID              :", C[i].id)
		fmt.Println("Nama            :", C[i].nama)
		fmt.Println("Partai          :", C[i].nama_partai)
		fmt.Println("Jenis Kelamin   :", C[i].jenis_kelamin)
		fmt.Println("Umur            :", C[i].umur)
		fmt.Println("Jumlah pemilih  :", C[i].banyak_pemilih)
	}
}

// Prosedur untuk mengurutkan dengan Insertion Sort dan menampilkan data calon legislatif berdasarkan jumlah votes secara DESCENDING
func showCandidatesByVotes(C arrCandidate, n int) {
	var pass, i int
	var temp candidate
	
	pass = 1
	for pass < n {
		i = pass
		temp = C[pass]
		for i > 0 && temp.banyak_pemilih > C[i-1].banyak_pemilih {
			C[i] = C[i-1]
			i--
		}
		C[i] = temp
		pass++
	}
	fmt.Print("Data calon terurut berdasarkan banyaknya pemilih: ")
	showCandidates(C, n)
}

// Prosedur untuk mengurutkan dengan Selection Sort dan menampilkan data calon legislatif berdasarkan nama partai secara ASCENDING
func showCandidatesByParty(C arrCandidate, n int) {
	var pass, i, idx int
	var temp candidate

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if C[idx].nama_partai == C[i].nama_partai {
				if C[idx].nama > C[i].nama {
					idx = i
				} 
			} else if C[idx].nama_partai > C[i].nama_partai {
				idx = i
			}
			i++
		}
		temp = C[pass-1]
		C[pass-1] = C[idx]
		C[idx] = temp
		pass++
	}
	fmt.Print("Data calon terurut berdasarkan partai: ")
	showCandidates(C, n)
}

// Prosedur untuk mengurutkan dengan Selection Sort dan menampilkan data calon legislatif berdasarkan nama calon legislatif dan nama partainya secara ASCENDING
func showCandidatesByNameAndParty(C arrCandidate, n int) {
	var pass, i, idx int
	var temp candidate

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if C[idx].nama == C[i].nama {
				if C[idx].nama_partai > C[i].nama_partai {
					idx = i
				}
			} else if C[idx].nama > C[i].nama {
				idx = i
			}
			i++
		}
		temp = C[pass-1]
		C[pass-1] = C[idx]
		C[idx] = temp
		pass++
	}
	fmt.Print("Data calon terurut berdasarkan nama kandidat dan partai: ")
	showCandidates(C, n)
}

// Prosedur untuk menambahkan data pemilih calon legislatif
func addVoter(C *arrVoter, n *int, candidate *arrCandidate, n_candidate int) {
	var v voter
	var i, tanggal, bulan, tahun int

	fmt.Println("Tanggal pemilihan antara 16 Mei 2023 - 21 Mei 2023 (16/6/2023 - 21/6/2023)")
	fmt.Print("Tanggal pemilihan              : ")
	fmt.Scanln(&tanggal)
	fmt.Print("Bulan pemilihan (dalam angka)  : ")
	fmt.Scanln(&bulan)
	fmt.Print("Tahun pemilihan                : ")
	fmt.Scanln(&tahun)

	if *n < NMAX {
		if isInVotingDuration(tanggal, bulan, tahun) == true {
			fmt.Print("Nama              : ")
			fmt.Scanln(&v.nama)
			fmt.Print("Jenis Kelamin     : ")
			fmt.Scanln(&v.jenis_kelamin)
			fmt.Print("NIK               : ")
			fmt.Scanln(&v.NIK)
			fmt.Print("Umur              : ")
			fmt.Scanln(&v.umur)
			fmt.Print("Kandidat Pilihan  : ")
			fmt.Scanln(&v.pilihan)
			fmt.Print("Partai Kandidat   : ")
			fmt.Scanln(&v.partai_kandidat)
			C[*n] = v
			*n++
			for i = 0; i < n_candidate; i++ {
				if candidate[i].nama == v.pilihan && candidate[i].nama_partai == v.partai_kandidat {
					candidate[i].banyak_pemilih++
				}
			}
			fmt.Println("Data pemilih telah ditambahkan.")
		}
	} else {
		fmt.Println("Batas maksimum pemilih telah tercapai.")
	}

	if isInVotingDuration(tanggal, bulan, tahun) == false {
		fmt.Println("Tanggal pemilihan berada di luar durasi waktu pemilihan.")
		fmt.Print("Data Calon: ")
		showCandidates(*candidate, n_candidate)
	}
}

// Fungsi untuk mengembalikan nilai true apabila pemilih berada di dalam durasi waktu pemilihan dan mengembalikan nilai false jika sebaliknya
func isInVotingDuration(tanggal, bulan, tahun int) bool {
	var condition bool = false

	if tahun <= 2023 {
		if bulan <= 6 {
			if tanggal >= 16 && tanggal <= 21 {
				condition = true
			}
		}
	}
	return condition
}

// Prosedur untuk menampilkan data pemilih calon legislatif
func showVoters(C arrVoter, n int) {
	if n == 0 {
		fmt.Print("-")
		fmt.Println()
	} else {
		fmt.Println()
		for i := 0; i < n; i++ {
			fmt.Println("-----------------------------")
			fmt.Println("Nama              :", C[i].nama)
			fmt.Println("Jenis Kelamin     :", C[i].jenis_kelamin)
			fmt.Println("NIK               :", C[i].NIK)
			fmt.Println("Umur              :", C[i].umur)
			fmt.Println("Kandidat Pilihan  :", C[i].pilihan)
			fmt.Println("Partai Kandidat   :", C[i].partai_kandidat)
		}
	}
}
// Fungsi untuk mencari kandidat yang terpilih
func searchElected(C arrCandidate, n, threshold int) int {
	var idx int = -1
	var i int
	
	for i < n && idx == -1 {
		if C[i].banyak_pemilih >= threshold {
			idx = i
		}
		i++
	}
	return idx
}

// Prosedur untuk menampilkan kandidat yang terpilih
func printElected(C arrCandidate, n, threshold int) {
	var i int
	
	i = searchElected(C, n, threshold)
	
	if i == -1 {
		fmt.Println()
		fmt.Println("Tidak ada kandidat yang menang")
	} else {
		fmt.Println()
		fmt.Println("============== PEMENANG =============")
		fmt.Println("ID              :", C[i].id)
		fmt.Println("Nama            :", C[i].nama)
		fmt.Println("Partai          :", C[i].nama_partai)
		fmt.Println("Jenis Kelamin   :", C[i].jenis_kelamin)
		fmt.Println("Umur            :", C[i].umur)
		fmt.Println("Jumlah pemilih  :", C[i].banyak_pemilih)
	}
}

// Prosedur untuk mencari data calon legislatif berdasarkan nama, nama partai, atau pemilih dari calon tersebut
func searchCandidate(C arrCandidate, V arrVoter, nc, nv int, searchBy string) {
	var temp [NMAX]int
	var counter int
	var nama, partai, pemilih string
	
	if searchBy == "nama" {
		fmt.Print("Masukkan Nama Calon: ")
		fmt.Scanln(&nama)
		for i := 0; i < nc; i++ {
			if C[i].nama == nama {
				temp[counter] = i
				counter++
			}
		}
	} else if searchBy == "partai" {
		fmt.Print("Masukkan Nama Partai: ")
		fmt.Scanln(&partai)
		for i := 0; i < nc; i++ {
			if C[i].nama_partai == partai {
				temp[counter] = i
				counter++
			}
		}
	} else if searchBy == "pemilih" {
		fmt.Print("Masukkan nama pemilih: ")
		fmt.Scanln(&pemilih)
		for i := 0; i < nv; i++ {
			if V[i].nama == pemilih {
				for j := 0; j < nc; j++ {
					if V[i].pilihan == C[j].nama && V[i].partai_kandidat == C[j].nama_partai {
						temp[counter] = j
						counter++
					}
				}
			}
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
	
	if counter > 0 {
		if searchBy == "partai" {
			fmt.Println("==============", partai, "===============")
			for k := 0; k < counter; k++ {
				fmt.Println("Nama            :", C[temp[k]].nama)
				fmt.Println("Nama Partai     :", C[temp[k]].nama_partai)
				fmt.Println("Jenis Kelamin   :", C[temp[k]].jenis_kelamin)
				fmt.Println("Umur            :", C[temp[k]].umur)
				fmt.Println("Jumlah pemilih  :", C[temp[k]].banyak_pemilih)
				fmt.Println()
			}
			fmt.Println("=======================================")
		} else if searchBy == "nama" {
			fmt.Println("==============", nama, "===============")
			for k := 0; k < counter; k++ {
				fmt.Println("Nama            :", C[temp[k]].nama)
				fmt.Println("Nama Partai     :", C[temp[k]].nama_partai)
				fmt.Println("Jenis Kelamin   :", C[temp[k]].jenis_kelamin)
				fmt.Println("Umur            :", C[temp[k]].umur)
				fmt.Println("Jumlah pemilih  :", C[temp[k]].banyak_pemilih)
				fmt.Println()
			}
			fmt.Println("=======================================")
		} else if searchBy == "pemilih" {
			fmt.Println("===============", pemilih, "===============")
			for k := 0; k < counter; k++ {
				fmt.Println("Nama            :", C[temp[k]].nama)
				fmt.Println("Nama Partai     :", C[temp[k]].nama_partai)
				fmt.Println("Jenis Kelamin   :", C[temp[k]].jenis_kelamin)
				fmt.Println("Umur            :", C[temp[k]].umur)
				fmt.Println("Jumlah pemilih  :", C[temp[k]].banyak_pemilih)
				fmt.Println()
			}
			fmt.Println("=======================================")
		}
	} else {
		fmt.Println("Data tidak ada")
	}
}

func main() {
	var candidate arrCandidate
	var voter arrVoter
	var n_candidate, n_voter, menu int
	var threshold int
	var searchBy string

	candidate = arrCandidate{
		{"agus", "pdi", "L", "60", "1001", 0},
		{"suharyuni", "merdeka", "P", "40", "1002", 0},
		{"agus", "merdeka", "L", "45", "1003", 0},
		{"ganjar", "pdi", "L", "50", "1004", 0},
		{"anis", "perindo", "L", "55", "1005", 0},
	}
	n_candidate = 5

	for menu != 14 {
		fmt.Println("=============== PEMILU ================")
		fmt.Println("1. Add Candidate")
		fmt.Println("2. Edit Candidate")
		fmt.Println("3. Delete Candidate")
		fmt.Println("4. Print Candidate")
		fmt.Println("5. Sort Candidate by Votes")
		fmt.Println("6. Sort Candidate by Party")
		fmt.Println("7. Sort Candidate by Name and Party")
		fmt.Println("8. Add Voter")
		fmt.Println("9. Print Voter")
		fmt.Println("10. Set Threshold")
		fmt.Println("11. Show Threshold")
		fmt.Println("12. Print Elected")
		fmt.Println("13. Search Candidate")
		fmt.Println("14. Exit")

		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&menu)
		fmt.Println()

		if menu == 1 {
			addCandidate(&candidate, &n_candidate)
		} else if menu == 2 {
			editCandidate(&candidate, n_candidate)
		} else if menu == 3 {
			deleteCandidate(&candidate, &n_candidate)
		} else if menu == 4 {
			fmt.Print("Data Calon: ")
			showCandidates(candidate, n_candidate)
		} else if menu == 5 {
			showCandidatesByVotes(candidate, n_candidate)
		} else if menu == 6 {
			showCandidatesByParty(candidate, n_candidate)
		} else if menu == 7 {
			showCandidatesByNameAndParty(candidate, n_candidate)
		} else if menu == 8 {
			addVoter(&voter, &n_voter, &candidate, n_candidate)
		} else if menu == 9 {
			fmt.Print("Data Pemilih: ")
			showVoters(voter, n_voter)
		} else if menu == 10 {
			fmt.Print("Masukkan nilai threshold: ")
			fmt.Scanln(&threshold)
		} else if menu == 11 {
			fmt.Println()
			fmt.Println("Threshold:", threshold)
		} else if menu == 12 {
			printElected(candidate, n_candidate, threshold)
		} else if menu == 13 {
			fmt.Print("Cari berdasarkan [nama] [partai] [pemilih]: ")
			fmt.Scanln(&searchBy)
			searchCandidate(candidate, voter, n_candidate, n_voter, searchBy)
		}
		fmt.Println()
	}
}