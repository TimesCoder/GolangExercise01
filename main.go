package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selamat datang di program CLI Go! 🎉")

	for {
		showMenu()

		fmt.Printf("Masukkan pilihan Anda: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		

		switch input {
		case "1":
			printHelloWorld()
		case "2":
			operasiMatematika()
		case "3":
			inputdanSimpanData()
		case "4":
			fungsiFibonacci()
		case "5":
			fmt.Println("Terima kasih telah menggunakan program CLI Go!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
		}

		fmt.Print("\nTekan 'Enter' untuk melanjutkan...")
		_, _ = reader.ReadString('\n')

		
		}
}

// Menampilkan menu utama
func showMenu() {
	fmt.Println("\nPilih menu:")
	fmt.Println("1. Menampilkan 'Hello, World!'")
	fmt.Println("2. Operasi Matematika Sederhana")
	fmt.Println("3. Menyimpan dan menampilkan data pengguna")
	fmt.Println("4. Fibonacci")
	fmt.Printf("5. Keluar \n")
}

// Fungsi Hello, World!
func printHelloWorld() {
	fmt.Println("Hello, World!")
}

func operasiMatematika() {
	reader := bufio.NewReader(os.Stdin)

	// Meminta input angka pertama
	fmt.Print("Masukkan angka pertama: ")
	angka1Str, _ := reader.ReadString('\n')
	angka1Str = strings.TrimSpace(angka1Str)

	// Konversi angka pertama ke float64
	angka1, err1 := strconv.ParseFloat(angka1Str, 64)
	if err1 != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka yang benar.")
		return
	}

	// Meminta input angka kedua
	fmt.Print("Masukkan angka kedua: ")
	angka2Str, _ := reader.ReadString('\n')
	angka2Str = strings.TrimSpace(angka2Str)

	// Konversi angka kedua ke float64
	angka2, err2 := strconv.ParseFloat(angka2Str, 64)
	if err2 != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka yang benar.")
		return
	}

	// Melakukan operasi matematika penjumlahan, pengurangan, perkalian, dan pembagian
	// %.2f untuk menampilkan 2 angka di belakang koma
	fmt.Printf("Hasil Penjumlahan: %.2f\n", angka1+angka2) 
	fmt.Printf("Hasil Pengurangan: %.2f\n", angka1-angka2)
	fmt.Printf("Hasil Perkalian: %.2f\n", angka1*angka2)

	// Pembagian dengan pengecekan apakah angka kedua adalah 0
	if angka2 != 0 {
		fmt.Printf("Hasil Pembagian: %.2f\n", angka1/angka2)
	} else {
		fmt.Println("Pembagian dengan nol tidak diperbolehkan.")
	}
}

func inputdanSimpanData() {
	// Map untuk menyimpan data pengguna
	dataPengguna := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	for {
		// Meminta input nama pengguna
		fmt.Print("Masukkan nama Anda: ")
		nama, _ := reader.ReadString('\n')
		nama = strings.TrimSpace(nama)

		// Meminta input umur pengguna
		fmt.Print("Masukkan umur Anda: ")
		umur, _ := reader.ReadString('\n')
		umur = strings.TrimSpace(umur)
		_, err := strconv.Atoi(umur)
		if err != nil {
			fmt.Println("Umur harus berupa angka. Silakan coba lagi.")
			continue
		}
		

		// Validasi input, nama dan umur tidak boleh kosong
		if nama == "" || umur == "" {
			fmt.Println("Nama dan umur tidak boleh kosong. Silakan coba lagi.")
			continue
		}

		// Menyimpan data ke dalam map
		dataPengguna[nama] = umur

		// Menanyakan apakah ingin menambah data lagi
		fmt.Print("Apakah ingin menambah data lagi? (ya/tidak): ")
		tentukan, _ := reader.ReadString('\n')
		tentukan = strings.TrimSpace(tentukan)

		// Jika tidak ingin menambah data
		if strings.ToLower(tentukan) != "ya" {
			break
		}
	}

	// Menampilkan data pengguna yang telah disimpan
	fmt.Println("\nData pengguna yang telah disimpan:")
	for key, value := range dataPengguna {
		fmt.Printf("Nama: %s, Umur: %s\n", key, value)
	}
}

func fungsiFibonacci() {
	reader := bufio.NewReader(os.Stdin)

	// Meminta input angka ke-n
	fmt.Print("Masukkan angka ke-n: ")
	nilaiN, _ := reader.ReadString('\n')
	nilaiN = strings.TrimSpace(nilaiN)

	// Konversi angka ke-n ke int
	nilai, err := strconv.Atoi(nilaiN)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka yang benar.")
		return
	}


	// Memanggil fungsi fibonacci
	result := func (nilai int) int {
			if nilai == 0 {
				return 0
			}
			if nilai == 1 {
				return 1
			}
            f := make([]int, nilai+2)
            f[0], f[1] = 0, 1
            for i := 2; i <= nilai; i++ {
                f[i] = f[i-1] + f[i-2]
            }
            return f[nilai]
        }
	fmt.Printf("Hasil fibonacci dari angka ke-%d adalah %d\n", nilai, result(nilai))
}

