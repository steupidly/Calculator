package main

import "fmt"

func main() {
	var menu int
	var nilai1 int
	var nilai2 int
	history := []string{}
	fmt.Println(("\t Selamat datang di Program Kalkulator Lanjutan"))
	fmt.Println("\t ----------------------------------------------")

	for menu < 9 {
		fmt.Println("Silahkan pilih nomor di menu :")
		fmt.Println("1. Penambahan (+)")
		fmt.Println("2. Pengurangan (-)")
		fmt.Println("3. Perkalian (*)")
		fmt.Println("4. Pembagian (/)")
		fmt.Println("5. History")
		fmt.Println("9. Exit")
		fmt.Print("Masukan pilihan Menu = ")
		fmt.Scanln(&menu)

		if menu == 9 {
			fmt.Println("Selamat Tinggal, See ya !!~")
			break
		}

		switch menu {
		case 1:
			fmt.Print("Silahkan masukan nilai pertama = ")
			fmt.Scanln(&nilai1)

			fmt.Print("Silahkan masukan nilai kedua = ")
			fmt.Scan(&nilai2)
			fmt.Println(nilai1, "+", nilai2, "=", nilai1+nilai2)
			history = append(history, fmt.Sprintf("%d + %d = %d\n", nilai1, nilai2, nilai1+nilai2))
		case 2:
			fmt.Print("Silahkan masukan nilai pertama = ")
			fmt.Scanln(&nilai1)

			fmt.Print("Silahkan masukan nilai kedua = ")
			fmt.Scan(&nilai2)
			fmt.Println(nilai1, "-", nilai2, "=", nilai1-nilai2)
			history = append(history, fmt.Sprintf("%d + %d = %d\n", nilai1, nilai2, nilai1-nilai2))
		case 3:
			fmt.Print("Silahkan masukan nilai pertama = ")
			fmt.Scanln(&nilai1)

			fmt.Print("Silahkan masukan nilai kedua = ")
			fmt.Scan(&nilai2)
			fmt.Println(nilai1, "*", nilai2, "=", nilai1*nilai2)
			history = append(history, fmt.Sprintf("%d + %d = %d\n", nilai1, nilai2, nilai1*nilai2))
		case 4:
			fmt.Print("Silahkan masukan nilai pertama = ")
			fmt.Scanln(&nilai1)

			fmt.Print("Silahkan masukan nilai kedua = ")
			fmt.Scan(&nilai2)
			if nilai2 != 0 {
				fmt.Println("Hasil ", nilai1, "/", nilai2, "=", nilai1/nilai2)
			} else {
				fmt.Println("Hasil ", nilai1, "/", nilai2, "=", 0)
			}
			history = append(history, fmt.Sprintf("%d + %d = %d\n", nilai1, nilai2, nilai1/nilai2))

		case 5:

			for i := 0; i < len(history); i++ {
				fmt.Print(history[i])
			}

		default:
			fmt.Println("Menu pilihant idak ada, silahkan pilih nomor menu")

		}
	}
}
