package main
import "fmt"

const NMAX int = 1000

type gym struct {
	namaLatihan    string
	durasi, kalori int
}

type tabGym [NMAX]gym

func inputJadwal(A *tabGym, n *int) {
	fmt.Print("Enter exercise name : ")
	fmt.Scan(&A[*n].namaLatihan)

	fmt.Print("Enter excercise duration (minutes) : ")
	fmt.Scan(&A[*n].durasi)

	fmt.Print("Enter calories burned : ")
	fmt.Scan(&A[*n].kalori)

	*n++
}

func tampilkanJadwal(A *tabGym, n int) {
	fmt.Println("|===============================================================|")
	fmt.Println("|                           SCHEDULE                            |")
	fmt.Println("|===============================================================|")

	for i := 0; i < n; i++ {
		fmt.Printf("| %-2d | %-10s | %-10d Minutes | %-17d kkal |\n", i+1, A[i].namaLatihan, A[i].durasi, A[i].kalori)
		fmt.Println("|===============================================================|")
	}
}

func editJadwal(A *tabGym, n int) {
	var nomor int

	fmt.Println("Enter 0 to Back to menu")
	fmt.Print("Choose exercise number to change: ")
	fmt.Scan(&nomor)
	index := nomor - 1

	if index >= 0 && index < n {
		fmt.Print("Enter new exercise name : ")
		fmt.Scan(&A[index].namaLatihan)

		fmt.Print("Enter new exercise duration (minutes) : ")
		fmt.Scan(&A[index].durasi)

		fmt.Print("Enter new calories burned : ")
		fmt.Scan(&A[index].kalori)
	} else {
		if nomor == 0 {
			fmt.Println("Back to menu")
		} else {
			fmt.Println("Invalid number")
		}
	}
}

func hapusJadwal(A *tabGym, n *int) {
	var nomor int

	fmt.Println("Enter 0 to back to menu")
	fmt.Print("Choose exercise number to delete : ")
	fmt.Scan(&nomor)
	index := nomor - 1

	if index >= 0 && index < *n {
		// Geser elemen setelahnya ke kiri
		for i := index; i < *n-1; i++ {
			(*A)[i] = (*A)[i+1]
		}
		*n = *n - 1
		fmt.Println("Exercise successfully deleted.")
	} else {
		if nomor == 0 {
			fmt.Println("Back to menu")
		} else {
			fmt.Println("Invalid number")
		}
	}
}

func rekomendasiWorkout(A *tabGym, n int) {
	if n == 0 {
		fmt.Println("There are no exercises yet, cannot give a recommendation!")
		return
	}

	var sudahDicek [NMAX]bool
	var namaJarang string
	var minCount = NMAX

	for i := 0; i < n; i++ {
		if sudahDicek[i] {
			continue
		}

		count := 1
		sudahDicek[i] = true

		for j := i + 1; j < n; j++ {
			if A[i].namaLatihan == A[j].namaLatihan {
				count++
				sudahDicek[j] = true
			}
		}

		if count < minCount {
			minCount = count
			namaJarang = A[i].namaLatihan
		}
	}

	fmt.Println("================== Workout Recommendation ===================")
	fmt.Println("Recommended Exercises : ", namaJarang)
	fmt.Println("Because this exercise has only been done ", minCount, "times.")
}

func selectionSortDurasiAscending(A *tabGym, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if A[j].durasi < A[minIndex].durasi {
				minIndex = j
			}
		}
		if minIndex != i {
			A[i], A[minIndex] = A[minIndex], A[i]
		}
	}
}

func selectionSortDurasiDescending(A *tabGym, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if A[j].durasi > A[minIndex].durasi {
				minIndex = j
			}
		}
		if minIndex != i {
			A[i], A[minIndex] = A[minIndex], A[i]
		}
	}
}

func insertionSortKaloriAscending(A *tabGym, n int) {
	for i := 1; i < n; i++ {
		temp := A[i]
		j := i - 1
		for j >= 0 && A[j].kalori > temp.kalori {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func insertionSortKaloriDescending(A *tabGym, n int) {
	for i := 1; i < n; i++ {
		temp := A[i]
		j := i - 1
		for j >= 0 && A[j].kalori < temp.kalori {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func laporanWorkoutTerakhir(A *tabGym, n int) {
	fmt.Println("|===============================================================|")
	fmt.Println("|                   Last 10 Activities Report                   |")
	fmt.Println("|===============================================================|")
	start := 0
	if n > 10 {
		start = n - 10
	}

	totalKalori := 0
	totalDurasi := 0

	for i := start; i < n; i++ {
		fmt.Printf("| %-2d | %-10s | %-10d Minutes | %-17d kkal |\n", i+1, A[i].namaLatihan, A[i].durasi, A[i].kalori)
		fmt.Println("|===============================================================|")
		totalKalori += A[i].kalori
		totalDurasi += A[i].durasi
	}

	fmt.Println("\n=============== Summary ===============")
	fmt.Println("Total Calories Burned :", totalKalori)
	fmt.Println("Total Durations :", totalDurasi, "minutes")
}

func main() {
	var data tabGym
	var n int
	var pilihan int

	for {
		fmt.Println("\n=== Daily Workout Management App ===")
		fmt.Println("[1] Add Schedule")
		fmt.Println("[2] Show Schedule")
		fmt.Println("[3] Edit Schedule")
		fmt.Println("[4] Delete Schedule")
		fmt.Println("[5] Find Exercise")
		fmt.Println("[6] Exercise Recommendation")
		fmt.Println("[7] Sort Exercises")
		fmt.Println("[8] Workout Reports")
		fmt.Println("[9] Exit")
		fmt.Print("Input Choice : ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			inputJadwal(&data, &n)
		case 2:
			tampilkanJadwal(&data, n)
		case 3:
			editJadwal(&data, n)
		case 4:
			hapusJadwal(&data, &n)
		case 5:
			var keyword string
			fmt.Print("Enter Exercise name to find : ")
			fmt.Scan(&keyword)
			cariJenisOlahraga(&data, n, keyword)
		case 6:
			rekomendasiWorkout(&data, n)
		case 7:
			var urut int
			fmt.Println("=== Sorting Method ===")
			fmt.Println("[1] Duration (Ascending)")
			fmt.Println("[2] Duration (Descending)")
			fmt.Println("[3] Calories (Ascending)")
			fmt.Println("[4] Calories (Descending)")
			fmt.Print("Choose Method : ")
			fmt.Scan(&urut)

			switch urut {
			case 1:
				selectionSortDurasiAscending(&data, n)
				fmt.Println("Data Sorted based on Duration (ascending).")
			case 2:
				selectionSortDurasiDescending(&data, n)
				fmt.Println("Data Sorted based on Duration (descending).")
			case 3:
				insertionSortKaloriAscending(&data, n)
				fmt.Println("Data Sorted based on Calories (ascending).")
			case 4:
				insertionSortKaloriDescending(&data, n)
				fmt.Println("Data Sorted based on Calories (descending).")
			default:
				fmt.Println("Invalid Choice.")
			}
		case 8:
			laporanWorkoutTerakhir(&data, n)
		case 9:
			fmt.Println("Thankyou for using this App!")
			return
		default:
			fmt.Println("Invalid Choice!")
		}
	}
}
