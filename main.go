package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB //variabel untuk menggunakan syntax yang ada didalam modul Sequel "db"

type account struct {
	email    string
	password string
	BB       float64
	TB       float64
}

var akun account

func login() {
	allAccount() // memanggil fungsi allAccount untuk menampilkan list akun yang sudah terdaftar
	fmt.Println()

	var id, name, email, password string
	var err error // Deklarasi variabel dengan tipe data error, untuk menampung error yang mungkin terjadi
	fmt.Println("Please enter your Email and Password")
	fmt.Print("Email : ")
	fmt.Scan(&email)
	fmt.Print("Password : ")
	fmt.Scan(&password)
	fmt.Println()

	result := db.QueryRow(`SELECT * FROM account WHERE email = ? AND password = ?`, email, password)
	/*
		- "db.QueryRow" untuk memeriksa ada atau tidaknya data di dalam database, memeriksa per baris
		- "?" untuk memberi value dari Golang ke SQL
	*/
	err = result.Scan(&id, &name, &email, &password)
	/*
		- "result.Scan" untuk memberikan value yang ada di Golang ke SQL
		- "err" untuk mengecek apakah Querynya ada atau tidak (jika tidak maka akan error)
	*/

	if err != nil {
		log.Fatal("Account not registered")
	}
	/*
		- "nil" adalah kosong, artinya tidak ada isi apapun di dalamnya (zero value)
		- Jika ada error pada Database, program akan dihentikan
		- "log" sama seperti "fmt", tetapi lebih lengkap
		- "Fatal()" untuk memberitahu apa isi errornya
	*/

	fmt.Println("You successfully logged in!")
	akun.email = email // untuk menyimpan email yang sudah diinput agar bisa lanjut ke fungsi selanjutnya
}

func register() {
	var name, email, password string
	var err error
	fmt.Println("Please fill this form to Create Account")
	fmt.Print("Name : ")
	fmt.Scan(&name)
	fmt.Print("Email : ")
	fmt.Scan(&email)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	_, err = db.Exec(`INSERT INTO account (name, email, password) VALUES (?, ?, ?)`, name, email, password)
	/*
		- "?" untuk memberi value dari Golang ke SQL
		- Jika ada hacker yang hendak meretas database, fungsi "?" untuk melindungi dari peretasan
	*/
	if err != nil {
		log.Fatal(err)
	}
}

func allAccount() {
	checkTableAccount := `SELECT * FROM account`
	result, err := db.Query(checkTableAccount)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close() // "defer", syntax yang tugasnya untuk menjalankan fungsi setelah fungsi di sekitarnya selesai

	fmt.Println("|====|=========|===========================|======================|")
	fmt.Printf("| %-1s | %-7s | %-25s | %-20s |\n", "Id", "Name", "Email", "Password")
	fmt.Println("|====|=========|===========================|======================|")

	for result.Next() {
		// fungsi dari result.Next adalah untuk mendapatkan data selanjutnya yang ada di database
		var id int
		var name, email, password string

		result.Scan(&id, &name, &email, &password)
		fmt.Printf("| %-2d | %-7s | %-25s | %-20s |\n", id, name, email, password)
		fmt.Println("|----|---------|---------------------------|----------------------|")
	}
	// untuk mendapatkan semua fungsi di tabel database
}

func BMI(BB, TB float64) float64 {
	return BB / ((TB / 100) * (TB / 100))
}

var currentWorkout int
var sudahPilihWorkout bool
var programBerjalan = true

func workOutMenu() {
	var option int
	fmt.Println("Please choose your path")
	fmt.Println("[1] Bulking 		(Body Maximum Index is below 19)")
	fmt.Println("[2] Cutting 		(Body Maximum Index is above 25)")
	fmt.Println("[3] Muscle Building	(Body Maximum Index is normal, between 19 to 25)")
	fmt.Println("[4] Exit")
	fmt.Print("Choose : ")
	fmt.Scan(&option)
	fmt.Println()

	switch {
	case option == 1:
		currentWorkout = 1
		bulk()
	case option == 2:
		currentWorkout = 2
		cut()
	case option == 3:
		currentWorkout = 3
		muscle()
	case option == 4:
		fmt.Println("Goodbye!")
		programBerjalan = false
	}
}

var riwayatLatihanBulk []latihan
var riwayatLatihanCut []latihan
var riwayatLatihanMuscle []latihan

func bulk() {
	fmt.Println("Wellcome to the Bulking Session!")
	fmt.Println("In this section, you must do as follows : ")
	fmt.Println("1. This WorkOut will take 4 days of exercises and 3 days rest")
	fmt.Println("2. You must do 3 sets of each exercise with 8-12 reps")
	fmt.Println("3. There is no cardio sessions, but Leg day is mandatory!")
	fmt.Println("4. You must eat more calories than your daily needs!")
	fmt.Println("5. Please sleep at least 7 hours a day for a better recovery")
	fmt.Println()

	// Menampilkan list latihan
	fmt.Println("Here is the schedule : ")
	fmt.Println("1. Chest WO (Incline Press, Bench Press, Fly Machine)")
	fmt.Println("2. Back WO (Lat Pulldown, Seated Row, Barbel Row)")
	fmt.Println("3. Leg WO (Leg Extension, Leg Curl, Leg Press, Squat)")
	fmt.Println("4. Full Arm WO (Lateral Raise, Rear Delt, Shoulder Press, Biceps Curl, Triceps Pushdown, Dumbel Curl)")
	fmt.Println("5. Rest Day")

	operation(&riwayatLatihanBulk)
}

func cut() {
	fmt.Println("Wellcome to the Cutting Session!")
	fmt.Println("In this section, you must do as follows : ")
	fmt.Println("1. This WorkOut will take 6 days with 1 rest day in a week")
	fmt.Println("2. A lot of Cardio exercises and less muscle building exercises")
	fmt.Println("3. You have to do 3-5 minutse for each variation of cardio")
	fmt.Println("4. For each muslce building exercises, you need 2-4 sets for each variation")
	fmt.Println("5. Maintain your food, you must eat less than your daily calorie intake")
	fmt.Println()

	// Menampilkan list latihan
	fmt.Println("Here is the schedule : ")
	fmt.Println("1. Chest WO (Incline Press, Bench Press, Fly Machine)")
	fmt.Println("2. Back WO (Lat Pulldown, Seated Row, Barbel Row)")
	fmt.Println("3. Leg WO (Leg Extension, Leg Curl, Leg Press, Squat)")
	fmt.Println("4. Full Arm WO (Lateral Raise, Rear Delt, Shoulder Press, Biceps Curl, Triceps Pushdown, Dumbel Curl")
	fmt.Println("5. Rest Day")
	fmt.Println("6. Long Run (30 minutes - 1 hours)")
	fmt.Println("7. Jogging (15 minutes - 45 minutes)")

	operation(&riwayatLatihanCut)
}

func muscle() {
	fmt.Println("Wellcome to the Muscle Building Session!")
	fmt.Println("In this section, you must do as follows : ")
	fmt.Println("1. This WorkOut will take 4-5 days in a week with 2-3 rest day")
	fmt.Println("2. You can combine 2 exercises in 1 day, but Chest cannot be combined with Back!")
	fmt.Println("3. You must do 3 sets of 8-12 reps for each exercise")
	fmt.Println("4. Don't take to much weight, just focus building your muscle!")
	fmt.Println("5. Sleep for 8 hours is necessary for the better recover and please eat healty to maintain your muscle")
	fmt.Println()

	// Menampilkan list latihan
	fmt.Println("Here is the schedule : ")
	fmt.Println("1. Chest WO (Incline Press, Bench Press, Fly Machine)")
	fmt.Println("2. Back WO (Lat Pulldown, Seated Row, Barbel Row)")
	fmt.Println("3. Leg WO (Leg Extension, Leg Curl, Leg Press, Squat)")
	fmt.Println("4. Shoulder WO (Lateral Raise, Rear Delt, Shoulder Press)")
	fmt.Println("5. Arm WO (Biceps Curl, Triceps Pushdown, Dumbel Curl)")
	fmt.Println("6. Rest Day")

	operation(&riwayatLatihanMuscle)
}

type latihan struct {
	NamaLatihan string
	DurasiMenit int
	Kalori      int
}

/*
kami akan menggunakan Slice untuk membuat 3 fungsi berikut ini, penjelasan Slice terkait :
- Slice adalah tipe data yang dapat menampung nilai-nilai dari tipe data lainnya
- Slice mempunyai "leght" -> jumlah elemen yang digunakan
- Slice mempunyai "capacity" -> jumlah elemen maksimum yang bisa ditampung
- Slice sebenarnya hanya "Jendela" ke array yang sesungguhnya
Singkatnya, slice adalah array fleksibel dan bisa bertambah ukuran
*/
func addSchedule(jadwal []latihan, nama string, durasi int, kalori int) []latihan {
	// fungsi menerima slice jadwal, nama, durasi, dan kalori
	data := latihan{
		NamaLatihan: nama,
		DurasiMenit: durasi,
		Kalori:      kalori,
	}
	return append(jadwal, data) // menambahkan data ke dalam slice
	/*
		- fungsi ini berguna untuk menambahkan jadwal latihan yang diinginkan user ke dalam slice
		- append() digunakan untuk menambahkan elemen ke slice, sederhananya append() membantu untuk memperbesar isi dari sebuah slice.
	*/
}

func editSchedule(jadwal []latihan, index int, nama string, durasi int, kalori int) {
	// fungsi menerima slice jadwal dan index elemen yang ingin diubah, dan nilai-nilai baru
	jadwal[index].NamaLatihan = nama
	jadwal[index].DurasiMenit = durasi
	jadwal[index].Kalori = kalori
	/*
		- fungsi ini berguna untuk mengubah jadwal latihan yang sebelumnya sudah di-input oleh user
		- langsung mengubah elemen ke - index dalam slice
	*/
}

func deleteSchedule(jadwal []latihan, index int) []latihan {
	// fungsi menerima slice jadwal dan index elemen yang akan dihapus, lalu :
	return append(jadwal[:index], jadwal[index+1:]...)
	/*
		- jadwal[:index] -> mengambil bagian sebelum elemen yang ingin dihapus
		- jadwal[index+1:] -> ambil bagian setelah elemen itu
		- '...' -> gabungkan keduanya, hasilnya adalah slice baru tanpa elemen pada posisi index
		- "..." disebut variadic operator , fungsinya adalah "membuka" slice agar elemen-elemennya bisa ditambahkan satu-persatu
	*/
}

func rekomendasi(jadwal []latihan) {
	if len(jadwal) == 0 {
		fmt.Println("There is no exercises to analyze yet.")
		return
	}

	count := make(map[string]int)
	for _, j := range jadwal {
		count[j.NamaLatihan]++ // menghitung jumlah latihan
	}

	var palingSering string
	max := 0
	for nama, frekuensi := range count {
		if frekuensi > max {
			max = frekuensi
			palingSering = nama
		}
	}
	fmt.Println("The next recommended exercise based on your history is:", palingSering)
}

func sequentialSearch(jadwal []latihan, target string) int {
	for i, j := range jadwal {
		if j.NamaLatihan == target {
			return i
		}
	}
	return -1
}

func selectionSortByDuration(jadwal []latihan) {
	n := len(jadwal)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if jadwal[j].DurasiMenit < jadwal[min].DurasiMenit {
				min = j
			}
		}
		jadwal[i], jadwal[min] = jadwal[min], jadwal[i]
	}
}

func insertionSortByCalories(jadwal []latihan) {
	for i := 1; i < len(jadwal); i++ {
		key := jadwal[i]
		j := i - 1
		for j >= 0 && jadwal[j].Kalori > key.Kalori {
			jadwal[j+1] = jadwal[j]
			j--
		}
		jadwal[j+1] = key
	}
}

func exercisesReport(jadwal []latihan) {
	if len(jadwal) == 0 {
		fmt.Println("There is no workout history!")
		return
	}

	// 10 aktivitas terakhir
	fmt.Println("======================================================================")
	fmt.Println("		            Last 10 Workouts		           ")
	fmt.Println("======================================================================")
	start := 0
	if len(jadwal) > 10 {
		start = len(jadwal) - 10
	}
	for i := start; i < len(jadwal); i++ {
		item := jadwal[i]
		fmt.Printf("| %-2d | %-10s | %-6s %-6d %-6s | %-7s %-6d %-4s |\n", i+1, item.NamaLatihan, "Duration:", item.DurasiMenit, "Minutes", "Calories:", item.Kalori, "kkal")
		fmt.Println("----------------------------------------------------------------------")
	}

	// Total kalori dalam periode tertentu
	var awal, akhir int
	fmt.Printf("\nEnter range of workout (1 to %d)\n", len(jadwal))
	fmt.Print("Start from (index): ")
	fmt.Scan(&awal)
	fmt.Print("Until (index): ")
	fmt.Scan(&akhir)

	if awal < 1 || akhir > len(jadwal) || awal > akhir {
		fmt.Println("Invalid range.")
		return
	}

	totalKalori := 0
	for i := awal - 1; i < akhir; i++ {
		totalKalori += jadwal[i].Kalori
	}
	fmt.Printf("Total calories burned from workout %d to %d: %d kkal\n", awal, akhir, totalKalori)

	// Total seluruh waktu latihan
	totalDurasi := 0
	for _, item := range jadwal {
		totalDurasi += item.DurasiMenit
	}
	fmt.Printf("Total workout time: %d minutes\n", totalDurasi)
	fmt.Println()
}

func operation(riwayatLatihan *[]latihan) {
	var pilihan int

	fmt.Println("\nMenu:")
	fmt.Println("[1] Add your workout schedule")
	fmt.Println("[2] See all schedule")
	fmt.Println("[3] Edit schedule")
	fmt.Println("[4] Delete schedule")
	fmt.Println("[5] Exercises recommendation")
	fmt.Println("[6] Search exercise")
	fmt.Println("[7] Short Schedule")
	fmt.Println("[8] View Report")
	fmt.Println("[9] Go back to Main Menu")
	fmt.Print("choose : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var nama string
		var durasi, kalori int

		fmt.Print("Input your schedule (without space) : ")
		fmt.Scan(&nama)
		fmt.Print("Input Duration (minutes) : ")
		fmt.Scan(&durasi)
		fmt.Print("Input calories burned : ")
		fmt.Scan(&kalori)

		*riwayatLatihan = addSchedule(*riwayatLatihan, nama, durasi, kalori)
		fmt.Println("Schedule added!")
		fmt.Println()

	case 2:
		if len(*riwayatLatihan) == 0 {
			fmt.Println("There is no schedule yet!")
			break
		}
		fmt.Println("Schedule : ")
		for i, item := range *riwayatLatihan {
			fmt.Println("===============================================================")
			fmt.Printf("| %-2d | %-10s | %-15d minutes | %-10d kkal |\n", i+1, item.NamaLatihan, item.DurasiMenit, item.Kalori)
		}
		fmt.Println("===============================================================")

	case 3:
		if len(*riwayatLatihan) == 0 {
			fmt.Println("There is no schedule yet!")
			fmt.Println()
			break
		}
		var index, durasiBaru, kaloriBaru int
		var namaBaru string

		fmt.Println("Choose schedule to edit : ")
		for i, item := range *riwayatLatihan {
			fmt.Printf("%d. %s | Duration : %d minutes | Calories : %d kkal\n", i+1, item.NamaLatihan, item.DurasiMenit, item.Kalori)
		}
		fmt.Print("Choose schedule number to edit : ")
		fmt.Scan(&index)

		if index < 1 || index > len(*riwayatLatihan) {
			fmt.Println("Invalid schedule number!")
			fmt.Println()
			break
		}

		fmt.Print("Enter new schedule (without space) : ")
		fmt.Scan(&namaBaru)
		fmt.Print("Enter new duration (minutes) : ")
		fmt.Scan(&durasiBaru)
		fmt.Print("Enter new calories burned : ")
		fmt.Scan(&kaloriBaru)

		editSchedule(*riwayatLatihan, index-1, namaBaru, durasiBaru, kaloriBaru)
		fmt.Println("Schedule edited!")
		fmt.Println()

	case 4:
		if len(*riwayatLatihan) == 0 {
			fmt.Println("There is no schedule yet!")
			fmt.Println()
			break
		}
		var index int
		fmt.Println("Choose shcedule to delete : ")
		for i, item := range *riwayatLatihan {
			fmt.Printf("%d. %s | Duration : %d minutes | Calories : %d kkal\n", i+1, item.NamaLatihan, item.DurasiMenit, item.Kalori)
		}
		fmt.Print("Choose schedule number to delete : ")
		fmt.Scan(&index)

		if index < 1 || index > len(*riwayatLatihan) {
			fmt.Println("Invalid schedule number!")
			fmt.Println()
			break
		}

		*riwayatLatihan = deleteSchedule(*riwayatLatihan, index-1)
		fmt.Println("Schedule deleted!")
		fmt.Println()

	case 5:
		rekomendasi(*riwayatLatihan)

	case 6:
		var target string
		fmt.Println("Enter workout name to search (without space): ")
		fmt.Scan(&target)

		// Pencarian Sequential
		idx := sequentialSearch(*riwayatLatihan, target)
		if idx != -1 {
			fmt.Printf("Sequential Search: found at index %d: %s\n", idx+1, (*riwayatLatihan)[idx].NamaLatihan)
		} else {
			fmt.Println("Sequential Search: not found.")
		}

	case 7:
		var metode int
		fmt.Println("Sort by:")
		fmt.Println("[1] Duration (Selection Sort)")
		fmt.Println("[2] Calories (Insertion Sort)")
		fmt.Print("Choose: ")
		fmt.Scan(&metode)

		if metode == 1 {
			selectionSortByDuration(*riwayatLatihan)
			fmt.Println("Sorted by duration.")
		} else if metode == 2 {
			insertionSortByCalories(*riwayatLatihan)
			fmt.Println("Sorted by calories.")
		} else {
			fmt.Println("Invalid choice.")
		}

	case 8:
		exercisesReport(*riwayatLatihan)

	case 9:
		fmt.Println()
		workOutMenu()

	default:
		fmt.Println("Invalid choice!")
		fmt.Println()
	}
}

func main() {
	var err error                                 // untuk mendeklarasi variabel Error
	db, err = sql.Open("sqlite", "./Database.db") //untuk mengakses file Database
	if err != nil {
		// "nil" adalah kosong, artinya tidak ada isi apapun di dalam Database
		log.Fatal(err)
		/*
			- Jika ada error pada Database, program akan dihentikan
			- "log" sama seperti "fmt", tetapi lebih lengkap
			- "Fatal(err)" untuk memberitahu bahwa errornya sangat fatal
		*/
	}
	defer db.Close() // "defer", syntax yang tugasnya untuk menjalankan fungsi setelah fungsi di sekitarnya selesai

	fmt.Println("\nWellcome to 2R Gym App! please login to access")

	for programBerjalan {
		createTableAccount := `CREATE TABLE IF NOT EXISTS account (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		);`
		_, err = db.Exec(createTableAccount)
		if err != nil {
			log.Fatal(err)
		}
		/*
			- "_," adalah penanda bahwa kita tidak mengambil value di Database
			- "db.Exec" adalah syntax untuk mengeksekusi Tabel SQL
			- createTableAccount diletakkan didalam loop agar setiap kali user membuat akun baru,
		*/

		if akun.email == "" && akun.password == "" {
			fmt.Println("[1] Sign In")
			fmt.Println("[2] Create Account")

			var choose int
			fmt.Print("choose : ")
			fmt.Scan(&choose)
			fmt.Println()

			switch {
			case choose == 1:
				login()
			case choose == 2:
				register()
			}

		} else {
			if akun.BB == 0 && akun.TB == 0 {
				fmt.Println("Please calculate your Body Max Index before choosing your path!")
				fmt.Print("Insert your Weight : ")
				fmt.Scan(&akun.BB)
				fmt.Print("Insert your Height : ")
				fmt.Scan(&akun.TB)
				fmt.Printf("Your BMI is : %.2f\n ", BMI(akun.BB, akun.TB))
				fmt.Println()

			} else {
				if !sudahPilihWorkout {
					workOutMenu()
					sudahPilihWorkout = true
				} else {
					switch currentWorkout {
					case 1:
						operation(&riwayatLatihanBulk)
					case 2:
						operation(&riwayatLatihanCut)
					case 3:
						operation(&riwayatLatihanMuscle)
					}
				}
			}
		}
	}
}
