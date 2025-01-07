// TUGAS BESAR ALPRO 2 KELOMPOK 10 S1IF-11-01
// DIVA OCTAVIANI      (2311102006)
// ARJUN WERDHO KUMORO (2311102009)
// RICO ADE PRATAMA    (2311102138)
// PROGRAM JUAL BELI BARANG UNTUK PEGAWAI TOKO
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Struktur untuk menyimpan informasi Data Barang
type Item struct {
	ID         int
	Name       string
	Harga_Beli float64
	Harga_Jual float64
	Stock      int
}

// Struktur untuk menyimpan informasi transaksi penjualan
type Transaction struct {
	ItemID     int
	ItemName   string
	Quantity   int
	TotalPrice float64
	SaleDate   time.Time
}

// Struktur untuk menyimpan Riwayat/History
type History struct {
	Action   string
	Category string
	ItemID   int
	Details  string
	Date     time.Time
}

// Struktur untuk Top 5 Barang Terlaris
type SoldItem struct {
	ItemID    int
	TotalSold int
}

// Daftar Data Barang yang tersedia
var items = [10]Item{
	{101, "Televisi", 1000000, 1500000, 15},
	{102, "AC", 800000, 1000000, 5},
	{103, "Kipas Angin", 30000, 50000, 20},
	{104, "Dispenser", 40000, 75000, 10},
	{105, "Mesin Cuci", 100000, 200000, 25},
}

// Daftar transaksi penjualan dan History Riwayat
var transactions [100]Transaction
var riwayat [100]History
var transactionCount, historyCount int

// Fungsi untuk format harga
func formatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

// Fungsi untuk menambahkan Data Barang
func addItem(ID int, Nama string, Harga_Beli, Harga_Jual float64, Stock int) {
	for _, item := range items {
		if item.ID == ID {
			fmt.Printf("## Gagal menambahkan. ID %d sudah ada.\n", ID)
			return
		}
	}
	for i := range items {
		if items[i].ID == 0 {
			items[i] = Item{ID, Nama, Harga_Beli, Harga_Jual, Stock}
			riwayat[historyCount] = History{
				Action:   "add",
				Category: "Data",
				ItemID:   ID,
				Details:  fmt.Sprintf("Nama: %s, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d", Nama, Harga_Beli, Harga_Jual, Stock),
				Date:     time.Now(),
			}
			historyCount++
			fmt.Printf("- %s (ID: %d, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d) pada %s\n",
				Nama, ID, Harga_Beli, Harga_Jual, Stock, time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("Data Barang berhasil ditambahkan.")
			return
		}
	}
	fmt.Println("## Tidak dapat menambahkan, kapasitas array penuh.")
}

// Fungsi untuk mengubah (edit) Data Barang
func editItem(ID int, Nama string, Harga_Beli, Harga_Jual float64, Stock int) {
	for i := range items {
		if items[i].ID == ID {
			if Harga_Beli <= 0 || Harga_Jual <= 0 || Stock <= 0 {
				fmt.Println("## Input Salah!! Harga beli, harga jual, dan stok harus lebih besar dari 0.")
				return
			}
			oldDetails := fmt.Sprintf("Nama: %s, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d",
				items[i].Name, items[i].Harga_Beli, items[i].Harga_Jual, items[i].Stock)
			items[i].Name = Nama
			items[i].Harga_Beli = Harga_Beli
			items[i].Harga_Jual = Harga_Jual
			items[i].Stock = Stock
			riwayat[historyCount] = History{
				Action:   "edit",
				Category: "Data",
				ItemID:   ID,
				Details: fmt.Sprintf("Old: %s | New: Nama: %s, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d",
					oldDetails, Nama, Harga_Beli, Harga_Jual, Stock),
				Date: time.Now(),
			}
			historyCount++
			fmt.Printf("- %s (ID: %d, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d) pada %s\n",
				Nama, ID, Harga_Beli, Harga_Jual, Stock, time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("Data Barang berhasil diperbarui.")
			return
		}
	}
	fmt.Println("## Data Barang dengan ID tersebut tidak ditemukan.")
}

// Fungsi untuk menghapus Data Barang
func deleteItem(ID int) {
	for i := range items {
		if items[i].ID == ID {
			oldDetails := fmt.Sprintf("Nama: %s, Harga Beli: Rp%.2f, Harga Jual: Rp%.2f, Stok: %d pada %s",
				items[i].Name, items[i].Harga_Beli, items[i].Harga_Jual, items[i].Stock, time.Now().Format("2006-01-02 15:04:05"))
			riwayat[historyCount] = History{
				Action:   "delete",
				Category: "Data",
				ItemID:   ID,
				Details:  fmt.Sprintf("Old: %s", oldDetails),
				Date:     time.Now(),
			}
			historyCount++
			for j := i; j < len(items)-1; j++ {
				items[j] = items[j+1]
			}
			items[len(items)-1] = Item{}
			fmt.Printf("Data Barang dengan ID %d berhasil dihapus pada %s\n", ID, time.Now().Format("2006-01-02 15:04:05"))
			return
		}
	}
	fmt.Println("## Data Barang dengan ID tersebut tidak ditemukan.")
}

// Fungsi untuk menambahkan transaksi penjualan
func addTransaction(itemID, quantity int) {
	if quantity <= 0 {
		fmt.Println("## Input Salah!! Jumlah Data Barang yang dijual harus lebih besar dari 0.")
		return
	}
	for i := range items {
		if items[i].ID == itemID {
			if items[i].Stock < quantity {
				fmt.Println("## Maaf, Stok tidak mencukupi!")
				return
			}
			totalPrice := float64(quantity) * items[i].Harga_Jual
			transactions[transactionCount] = Transaction{
				ItemID:     items[i].ID,
				ItemName:   items[i].Name,
				Quantity:   quantity,
				TotalPrice: totalPrice,
				SaleDate:   time.Now(),
			}
			transactionCount++
			items[i].Stock -= quantity
			riwayat[historyCount] = History{
				Action:   "add",
				Category: "transaksi",
				ItemID:   itemID,
				Details:  fmt.Sprintf("Data Barang: %s, Jumlah: %d, Total Harga: Rp%.2f", items[i].Name, quantity, totalPrice),
				Date:     time.Now(),
			}
			historyCount++
			fmt.Printf("- %s (ID: %d, Jumlah: %d, Total Harga: Rp%.2f, Diedit pada: %s, Status: Berhasil)\n",
				items[i].Name, items[i].ID, quantity, totalPrice, time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("Transaksi berhasil ditambahkan.")
			return
		}
	}
	fmt.Println("## Data Barang dengan ID tersebut tidak ditemukan.")
}

// Fungsi untuk mengedit transaksi
func editTransaction(transactionIndex, newQuantity int) {
	if transactionCount == 0 {
		fmt.Println("## Tidak ada transaksi yang tersedia untuk diedit.")
		return
	}
	if newQuantity <= 0 {
		fmt.Println("## Input Salah!! Jumlah Data Barang yang dijual harus lebih besar dari 0.")
		return
	}
	if transactionIndex < 0 || transactionIndex >= transactionCount {
		fmt.Println("## Data transaksi tidak bisa diedit, indeks transaksi tidak ditemukan.")
		return
	}
	transaction := &transactions[transactionIndex]
	itemIndex := -1
	for i, item := range items {
		if item.ID == transaction.ItemID {
			itemIndex = i
			break
		}
	}
	if itemIndex == -1 {
		fmt.Println("## Data transaksi tidak bisa diedit, item tidak ditemukan.")
		return
	}
	deltaQuantity := newQuantity - transaction.Quantity
	if deltaQuantity > 0 && items[itemIndex].Stock < deltaQuantity {
		fmt.Println("## Data transaksi tidak bisa diedit, stok tidak mencukupi.")
		return
	}
	oldDetails := fmt.Sprintf("Jumlah: %d, Total Harga: Rp%.2f", transaction.Quantity, transaction.TotalPrice)
	items[itemIndex].Stock = items[itemIndex].Stock + transaction.Quantity - newQuantity
	transaction.Quantity = newQuantity
	transaction.TotalPrice = float64(newQuantity) * items[itemIndex].Harga_Jual
	riwayat[historyCount] = History{
		Action:   "edit",
		Category: "transaksi",
		ItemID:   transaction.ItemID,
		Details:  fmt.Sprintf("Old: %s | New: Jumlah: %d, Total Harga: Rp%.2f", oldDetails, newQuantity, transaction.TotalPrice),
		Date:     time.Now(),
	}
	historyCount++
	fmt.Printf("\n- %s (ID: %d, Jumlah: %d, Total Harga: Rp%.2f, Diedit pada: %s, Status: Berhasil)\n",
		items[itemIndex].Name, items[itemIndex].ID, newQuantity, transaction.TotalPrice, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Transaksi berhasil diperbarui.")
}

// Fungsi untuk menghapus transaksi penjualan
func deleteTransaction(transactionIndex int) {
	if transactionIndex >= 0 && transactionIndex < transactionCount {
		transaction := transactions[transactionIndex]
		found := false
		for i := range items {
			if items[i].ID == transaction.ItemID && !found {
				items[i].Stock += transaction.Quantity
				riwayat[historyCount] = History{
					Action:   "delete",
					Category: "transaksi",
					ItemID:   transaction.ItemID,
					Date:     time.Now(),
					Details: fmt.Sprintf("Data Barang: %s, Jumlah: %d, Total Harga: Rp%.2f pada %s",
						items[i].Name, transaction.Quantity, transaction.TotalPrice, time.Now().Format("2006-01-02 15:04:05")),
				}
				historyCount++
				found = true
			}
		}
		for j := transactionIndex; j < transactionCount-1; j++ {
			transactions[j] = transactions[j+1]
		}
		transactions[transactionCount-1] = Transaction{}
		transactionCount--
		fmt.Printf("Transaksi dengan indeks %d berhasil dihapus pada %s\n", transactionIndex, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println("## Data ini tidak termasuk data Transaksi, harus terjadi transaksi dulu. Pastikan indeks yang dimasukkan benar!")
	}
}

// Fungsi untuk menampilkan Data Barang yang tersedia
func displayItems() {
	fmt.Println("Daftar Barang:")
	fmt.Printf("+-----+--------------------+------------------+------------------+-----------------+\n")
	fmt.Printf("| ID  | Nama Barang        | Harga Beli (Rp)  | Harga Jual (Rp)  | Stok Tersedia   |\n")
	fmt.Printf("+-----+--------------------+------------------+------------------+-----------------+\n")
	for _, item := range items {
		if item.ID != 0 && item.Name != "" && item.Harga_Beli > 0 && item.Harga_Jual > 0 && item.Stock > 0 {
			fmt.Printf("| %-3d | %-18s | %-16s | %-16s | %-15d |\n",
				item.ID, item.Name, formatPrice(item.Harga_Beli), formatPrice(item.Harga_Jual), item.Stock)
		}
	}
	fmt.Printf("+-----+--------------------+------------------+------------------+-----------------+\n")
}

// Fungsi untuk mengurutkan Data Barang berdasarkan kategori
func sortItems(by, order string) {
	switch by {
	case "ID":
		if order == "asc" {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].ID < items[:][j].ID })
		} else {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].ID > items[:][j].ID })
		}
	case "Nama":
		if order == "asc" {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Name < items[:][j].Name })
		} else {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Name > items[:][j].Name })
		}
	case "Harga_Beli":
		if order == "asc" {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Harga_Beli < items[:][j].Harga_Beli })
		} else {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Harga_Beli > items[:][j].Harga_Beli })
		}
	case "Harga_Jual":
		if order == "asc" {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Harga_Jual < items[:][j].Harga_Jual })
		} else {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Harga_Jual > items[:][j].Harga_Jual })
		}
	case "Stock":
		if order == "asc" {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Stock < items[:][j].Stock })
		} else {
			selectionSort(items[:], func(i, j int) bool { return items[:][i].Stock > items[:][j].Stock })
		}
	default:
		fmt.Println("Kriteria urutkan tidak valid.")
		return
	}
	displayItems()
}

// Fungsi Selection Sort
func selectionSort(items []Item, compare func(i, j int) bool) {
	for i := 0; i < len(items)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(items); j++ {
			if compare(j, minIndex) {
				minIndex = j
			}
		}
		if minIndex != i {
			items[i], items[minIndex] = items[minIndex], items[i]
		}
	}
}

// Fungsi untuk mencari Data Barang berdasarkan kata kunci
func searchItems(keyword string) {
	fmt.Println("Hasil Pencarian:")
	fmt.Printf("+-----+--------------------+------------------+------------------+-----------------+\n")
	fmt.Printf("| ID  | Nama Barang        | Harga Beli (Rp)  | Harga Jual (Rp)  | Stok Tersedia   |\n")
	fmt.Printf("+-----+--------------------+------------------+------------------+-----------------+\n")
	found := false
	for _, item := range items {
		if strings.Contains(strings.ToLower(strconv.Itoa(item.ID)), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(item.Name), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(fmt.Sprintf("%.2f", item.Harga_Beli)), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(fmt.Sprintf("%.2f", item.Harga_Jual)), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(strconv.Itoa(item.Stock)), strings.ToLower(keyword)) {
			fmt.Printf("| %-3d | %-18s | %-16s | %-16s | %-15d |\n",
				item.ID,
				item.Name,
				formatPrice(item.Harga_Beli),
				formatPrice(item.Harga_Jual),
				item.Stock)
			found = true
		}
	}
	if !found {
		fmt.Println("|                     Tidak ada Data yang sesuai.                                  |")
	}
	fmt.Printf("+-----+--------------------+------------------+------------------+-----------------+\n")
}

// Fungsi untuk menampilkan laporan modal dan pendapatan
func displayReport() {
	var totalHarga_Beli, totalHarga_Jual, totalProfit float64
	for _, transaction := range transactions {
		for _, item := range items {
			if item.ID == transaction.ItemID {
				totalHarga_Beli += float64(transaction.Quantity) * item.Harga_Beli
				totalHarga_Jual += float64(transaction.Quantity) * item.Harga_Jual
				totalProfit += (item.Harga_Jual - item.Harga_Beli) * float64(transaction.Quantity)
			}
		}
	}
	fmt.Printf("Total Modal      : Rp %.2f\n", totalHarga_Beli)
	fmt.Printf("Total Pendapatan : Rp %.2f\n", totalHarga_Jual)
	fmt.Printf("Total Keuntungan : Rp %.2f\n", totalProfit)
}

// Fungsi untuk mengurutkan SoldItem berdasarkan TotalSold
func sortSoldItemsDescending(items []SoldItem) {
	for i := 0; i < len(items)-1; i++ {
		maxIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[j].TotalSold > items[maxIndex].TotalSold {
				maxIndex = j
			}
		}
		if maxIndex != i {
			items[i], items[maxIndex] = items[maxIndex], items[i]
		}
	}
}

// Fungsi Top 5 Barang Terlaris
func displayTopSoldItems() {
	soldItemsMap := make(map[int]int)
	for _, transaction := range transactions {
		soldItemsMap[transaction.ItemID] += transaction.Quantity
	}
	var soldItems [10]SoldItem
	count := 0
	for itemID, totalSold := range soldItemsMap {
		if count < len(soldItems) {
			soldItems[count] = SoldItem{ItemID: itemID, TotalSold: totalSold}
			count++
		}
	}
	sortSoldItemsDescending(soldItems[:count])
	fmt.Println("TOP 5 Barang yang Paling Banyak Terjual:")
	for i := 0; i < 5; i++ {
		if i < count {
			itemFound := false
			for _, item := range items {
				if soldItems[i].ItemID == item.ID {
					fmt.Printf("%d. Nama Barang: %s, Terjual: %d\n", i+1, item.Name, soldItems[i].TotalSold)
					itemFound = true
					break
				}
			}
			if !itemFound {
				fmt.Printf("%d. Nama Barang: Tidak ditemukan, Terjual: %d\n", i+1, soldItems[i].TotalSold)
			}
		} else {
			fmt.Printf("%d. Belum ada data\n", i+1)
		}
	}
}

// Fungsi untuk menampilkan Riwayat/History
func showHistory() {
	if len(riwayat) == 0 {
		fmt.Println("## Tidak ada riwayat tindakan.")
		return
	}
	fmt.Println("### Riwayat Aksi:")
	for _, history := range riwayat {
		if history.Action != "" && !history.Date.IsZero() {
			fmt.Printf("%s - %s: %s (ID: %d, %s) pada %s\n",
				history.Date.Format("2006-01-02 15:04:05"), history.Action, history.Category,
				history.ItemID, history.Details, history.Date.Format("2006-01-02 15:04:05"))
		}
	}
}

// Fungsi utama untuk menjalankan aplikasi
func main() {
	var choice string
	for {
		fmt.Println("\n===================================================================")
		fmt.Println(">>>>>>>>>>> Menu Aplikasi Jual Beli Untuk Pegawai Toko <<<<<<<<<<<<")
		fmt.Println("===================================================================")
		fmt.Println("1.  Tampilkan Data                              (•⊙ω⊙•)")
		fmt.Println("2.  Tambah Data                                 (＾∀＾)")
		fmt.Println("3.  Edit Data                                   (｡◕‿◕｡)")
		fmt.Println("4.  Hapus Data                                  ᕙ(⇀‸↼‶)ᕗ")
		fmt.Println("5.  Tambah Transaksi Penjualan                  (๑>ᴗ<๑)")
		fmt.Println("6.  Edit Transaksi Penjualan                    ( ͡° ͜ʖ ͡°)")
		fmt.Println("7.  Hapus Transaksi Penjualan                   (╥﹏╥)")
		fmt.Println("8.  Urutkan Data                                (•̀‿⊹ )")
		fmt.Println("9.  Pencarian Data                              q(◉ᴥ◉)p")
		fmt.Println("10. Lihat Laporan Modal dan Pendapatan          [̲̅$̲̅(̲̅5̲̅)̲̅$̲̅]")
		fmt.Println("11. TOP 5 Data Terlaris                         (⌐▨_▨)")
		fmt.Println("12. History Data dan Transaksi                  (☞ﾟヮﾟ)☞")
		fmt.Println("13. Keluar                                      (✖╭╮✖)")
		fmt.Println("===================================================================")
		fmt.Print("Pilih menu (1-13): ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			displayItems()
		case "2":
			var ID int
			var Nama string
			var Harga_Beli, Harga_Jual float64
			var Stock int
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Masukkan ID Barang   : ")
			fmt.Scanln(&ID)
			fmt.Print("Masukkan nama Barang : ")
			Nama, _ = reader.ReadString('\n')
			Nama = strings.TrimSpace(Nama)
			fmt.Print("Masukkan harga beli  : ")
			fmt.Scanln(&Harga_Beli)
			fmt.Print("Masukkan harga jual  : ")
			fmt.Scanln(&Harga_Jual)
			fmt.Print("Masukkan jumlah stok : ")
			fmt.Scanln(&Stock)
			addItem(ID, Nama, Harga_Beli, Harga_Jual, Stock)
		case "3":
			var ID int
			var Nama string
			var Harga_Beli, Harga_Jual float64
			var Stock int
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Masukkan ID Barang yang ingin diedit : ")
			fmt.Scanln(&ID)
			fmt.Print("Masukkan nama barang baru : ")
			Nama, _ = reader.ReadString('\n')
			Nama = strings.TrimSpace(Nama)
			fmt.Print("Masukkan harga beli baru  : ")
			fmt.Scanln(&Harga_Beli)
			fmt.Print("Masukkan harga jual baru  : ")
			fmt.Scanln(&Harga_Jual)
			fmt.Print("Masukkan jumlah stok baru : ")
			fmt.Scanln(&Stock)
			editItem(ID, Nama, Harga_Beli, Harga_Jual, Stock)
		case "4":
			var ID int
			fmt.Print("Masukkan ID Barang yang ingin dihapus : ")
			fmt.Scanln(&ID)
			deleteItem(ID)
		case "5":
			var ID, qty int
			fmt.Print("Masukkan ID Barang yang terjual : ")
			fmt.Scanln(&ID)
			fmt.Print("Masukkan jumlah Data Barang yang terjual : ")
			fmt.Scanln(&qty)
			addTransaction(ID, qty)
		case "6":
			var index, qty int
			fmt.Print("Masukkan indeks transaksi yang ingin diedit   : ")
			fmt.Scanln(&index)
			fmt.Print("Masukkan jumlah baru Data Barang yang terjual : ")
			fmt.Scanln(&qty)
			editTransaction(index, qty)
		case "7":
			var index int
			fmt.Print("Masukkan indeks transaksi yang ingin dihapus : ")
			fmt.Scanln(&index)
			deleteTransaction(index)
		case "8":
			var sortBy, order string
			fmt.Print("Urutkan berdasarkan (pilih : ID/Nama/Harga_Beli/Harga_Jual/Stock): ")
			fmt.Scanln(&sortBy)
			fmt.Print("Urutkan (asc/desc): ")
			fmt.Scanln(&order)
			sortItems(sortBy, order)
		case "9":
			var keyword string
			fmt.Print("Kata Kunci pencarian bisa berupa apa saja (misal dari ID/Nama/Harga/Huruf/Angka/dll)")
			fmt.Print("\nMasukkan Kata Kunci pencarian : ")
			fmt.Scanln(&keyword)
			searchItems(keyword)
		case "10":
			displayReport()
		case "11":
			displayTopSoldItems()
		case "12":
			showHistory()
		case "13":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini! semoga berkah dan sukses selalu (^‿^)")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
