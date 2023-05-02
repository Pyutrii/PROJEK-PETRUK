package main

import "fmt"

type Siswa struct {
	nim           int
	nama_siswa    string
	Jenis_kelamin string
	alamat        string
}
type Guru struct {
	nip       int
	nama_guru string
	alamat    string
}
type Sekolah1 struct {
	data Siswa
	next *Sekolah1
}
type Sekolah2 struct {
	data1 Guru
	next  *Sekolah2
}

func View_menu1() {
	fmt.Println(" PILIHAN MENU AWAL")
	fmt.Println("1. Siswa ")
	fmt.Println("2. Guru ")
	fmt.Println("3. EXIT ")
}
func View_menu2() {
	fmt.Println("PILIHAN MENU MANIPULASI DATA")
	fmt.Println("1. Insert ")
	fmt.Println("2. Update ")
	fmt.Println("3. Delete ")
	fmt.Println("4. View All ")
	fmt.Println("5. View By ID ")
	fmt.Println("6. BACK ")
}

func Insert_siswa(sekolah1 *Sekolah1, siswa Siswa) {
	newSekolah1 := Sekolah1{}
	newSekolah1.data = siswa
	temp := sekolah1
	if temp.next == nil {
		temp.next = &newSekolah1
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &newSekolah1
	}
}
func Insert_guru(sekolah2 *Sekolah2, guru Guru) {
	newSekolah2 := Sekolah2{}
	newSekolah2.data1 = guru
	temp := sekolah2
	if temp.next == nil {
		temp.next = &newSekolah2
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &newSekolah2
	}
}

func Update_siswa(sekolah1 *Sekolah1, NIM int) {
	var nama string
	temp := sekolah1
	for temp != nil {
		if temp.data.nim == NIM {
			fmt.Println("MASUKAN DATA BARU")
			fmt.Print("Masukan nama :")
			fmt.Scan(&nama)
			temp.data.nama_siswa = nama
			return
		}
		temp = temp.next
	}
}
func Update_guru(sekolah2 *Sekolah2, Nip int) {
	var nama string
	temp := sekolah2
	for temp != nil {
		if temp.data1.nip == Nip {
			fmt.Println("MASUKAN DATA BARU")
			fmt.Print("Masukan nama :")
			fmt.Scan(&nama)
			temp.data1.nama_guru = nama
			return
		}
		temp = temp.next
	}
}
func Delete_datasiswa(sekolah1 *Sekolah1, Letak int) {
	letak := 1
	temp := sekolah1
	for temp != nil {
		if letak == Letak {
			temp.next = temp.next.next
		}
		letak++
		temp = temp.next
	}
}
func Delete_dataguru(sekolah2 *Sekolah2, Letak int) {
	letak := 1
	temp := sekolah2
	for temp != nil {
		if letak == Letak {
			temp.next = temp.next.next
		}
		letak++
		temp = temp.next
	}
}
func Viewaall_siswa(sekolah1 *Sekolah1) {
	temp := sekolah1.next
	for temp != nil {
		fmt.Println("nim:", temp.data.nim)
		fmt.Println("nama:", temp.data.nama_siswa)
		fmt.Println("Jenis Kelamin:", temp.data.Jenis_kelamin)
		fmt.Println("Alamat:", temp.data.alamat)
		temp = temp.next
	}
}
func Viewaall_guru(sekolah2 *Sekolah2) {
	temp := sekolah2.next
	for temp != nil {
		fmt.Println("Nip :", temp.data1.nip)
		fmt.Println("Nama:", temp.data1.nama_guru)
		fmt.Println("Alamat:", temp.data1.alamat)
		temp = temp.next
	}
}
func Viewbyid_siswa(sekolah1 *Sekolah1, NIM int) {
	temp := sekolah1
	for temp != nil {
		if NIM == temp.data.nim {
			fmt.Println("nim:", temp.data.nim)
			fmt.Println("nama:", temp.data.nama_siswa)
			fmt.Println("Jenis Kelamin:", temp.data.Jenis_kelamin)
			fmt.Println("Alamat:", temp.data.alamat)
		}
		temp = temp.next
	}
}
func Viewbyid_guru(sekolah2 *Sekolah2, NIP int) {
	temp := sekolah2
	for temp != nil {
		if NIP == temp.data1.nip {
			fmt.Println("Nip :", temp.data1.nip)
			fmt.Println("Nama :", temp.data1.nama_guru)
			fmt.Println("Alamat :", temp.data1.alamat)
		}
		temp = temp.next
	}
}

func main() {
	var in int = 0
	for in != 3 {
		View_menu1()
		fmt.Print("masukkan pilihan menu :")
		fmt.Scan(&in)
		switch in {
		case 1:
			fmt.Println("ANDA MASUK PILIHAN 1:SISWA")
			datasiswa := Sekolah1{}
			var input int = 0
			for input != 6 {
				View_menu2()
				fmt.Print("masukkan pilihan menu :")
				fmt.Scan(&input)
				switch input {
				case 1:
					fmt.Println("anda masuk pilihan 1: INSERT")
					var nim int
					var nama string
					var al string
					var jk string
					fmt.Print("masukkan nim siswa: ")
					fmt.Scan(&nim)
					fmt.Print("masukkan nama siswa: ")
					fmt.Scan(&nama)
					fmt.Print("masukkan  Jenis Kelamin siswa: ")
					fmt.Scan(&jk)
					fmt.Print("masukkan  alamat  siswa: ")
					fmt.Scan(&al)

					data := Siswa{
						nim:           nim,
						nama_siswa:    nama,
						Jenis_kelamin: jk,
						alamat:        al,
					}
					Insert_siswa(&datasiswa, data)
					fmt.Println("Data berhasil ditambahkan")
				case 2:
					fmt.Println("anda masuk pilihan 2:UPDATE")
					fmt.Print("Masukan Nim siswa yang akan diupdate:")
					var Nim int
					fmt.Scan(&Nim)
					Update_siswa(&datasiswa, Nim)
				case 3:
					fmt.Println("anda masuk pilihan 3:DELETE")
					var index int
					fmt.Print("Hapus data ke-:")
					fmt.Scan(&index)
					Delete_datasiswa(&datasiswa, index)
				case 4:
					fmt.Println("anda masuk pilihan 4:VIEW ALL")
					Viewaall_siswa(&datasiswa)
				case 5:
					fmt.Println("anda masuk pilihan 5:VIEW BY ID")
					var Nim int
					fmt.Print("Mencari nim siswa ke-:")
					fmt.Scan(&Nim)
					Viewbyid_siswa(&datasiswa, Nim)
				}
			}
		case 2:
			View_menu2()
			fmt.Println("ANDA MASUK PILIHAN 1:GURU")
			dataguru := Sekolah2{}
			var input int = 0
			for input != 6 {
				fmt.Print("masukkan pilihan menu :")
				fmt.Scan(&input)
				switch input {
				case 1:
					fmt.Println("anda masuk pilihan 1: INSERT")
					var nip int
					var nama string
					var al string
					fmt.Print("masukkan nip guru: ")
					fmt.Scan(&nip)
					fmt.Print("masukkan nama guru: ")
					fmt.Scan(&nama)
					fmt.Print("masukkan  alamat guru: ")
					fmt.Scan(&al)

					data1 := Guru{
						nip:       nip,
						nama_guru: nama,
						alamat:    al,
					}
					Insert_guru(&dataguru, data1)
				case 2:
					fmt.Println("anda masuk pilihan 2 : UPDATE ")
					fmt.Print("Masukan Nip guru yang akan diupdate:")
					var Nip int
					fmt.Scan(&Nip)
					Update_guru(&dataguru, Nip)
				case 3:
					fmt.Println("anda masuk pilihan 3 : DELETE ")
					var index int
					fmt.Print("Mencari Nip guru:")
					fmt.Scan(&index)
					Delete_dataguru(&dataguru, index)

				case 4:
					fmt.Println("anda masuk pilihan 4 : VIEW ALL ")
					Viewaall_guru(&dataguru)
				case 5:
					fmt.Println("anda masuk pilihan 5 : VIEW BY ID")
					var Nip int
					fmt.Print("Mencari nip guru ke-:")
					fmt.Scan(&Nip)
					Viewbyid_guru(&dataguru, Nip)
				}
			}
		}
	}
}
