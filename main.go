package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const MIN_MERGE = 32 

// UTILITY FUNCTIONS //
func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}

func copySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func measureTime(fn func ()) float64 {
	start := time.Now()
	fn()
	return float64(time.Since(start).Nanoseconds()) / 1_000_000
}

// SORTING FUNCTIONS //
func insertionSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func merge(arr []int, left, mid, right int) {
	leftArr := append([]int{}, arr[left:mid+1]...)
	rightArr := append([]int{}, arr[mid+1:right+1]...)

	i, j, k := 0, 0, left
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}
	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// TIM SORT //
func timSortRecursiveHelper(arr []int, left, right int) {
	if right-left+1 <= MIN_MERGE {
		insertionSort(arr, left, right)
		return
	}

	mid := (left + right) / 2
	timSortRecursiveHelper(arr, left, mid)
	timSortRecursiveHelper(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func timSortRecursive(arr []int) {
	if len(arr) > 1 {
		timSortRecursiveHelper(arr, 0, len(arr)-1)
	}
}

func timSortIterative(arr []int) {
	n := len(arr)
	for i := 0; i < n; i += MIN_MERGE {
		insertionSort(arr, i, min(i+MIN_MERGE-1, n-1))
	}
	for size := MIN_MERGE; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := left + size - 1
			right := min(left+2*size-1, n-1)
			if mid < right {
				merge(arr, left, mid, right)
			}
		}
	}
}

// DATA & BENCHMARK //
func generateData(n int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, n)
	for i := range data {
		data[i] = rand.Intn(n)
	}
	return data
}

type Result struct {
	n    int
	rec  float64
	iter float64
}

// OUTPUT //
func printTable(results []Result) {
	fmt.Println("\n========================================== ")
	fmt.Println("TIM SORT - PERBANDINGAN WAKTU EKSEKUSI (s) ")
	fmt.Println("========================================== ")
	fmt.Printf("%-10s %-15s %-15s\n", "N Data", "Recursive", "Iterative")
	fmt.Println("========================================== ")
	for _, r := range results {
		fmt.Printf("%-10d %-15.5f %-15.5f\n", r.n, r.rec, r.iter)
	}	
}

func printBarChart(results []Result) {
	fmt.Println("\n==============================================")
	fmt.Println(" DIAGRAM BATANG (█)")
	fmt.Println("==============================================")

	const scale = 0.5
	for _, r := range results {
		fmt.Printf("\nn = %d\n", r.n)
		fmt.Printf("Recursive: %s%.2f s\n", strings.Repeat("█", int(r.rec*scale)), r.rec)
		fmt.Printf("Iterative: %s%.2f s\n", strings.Repeat("█", int(r.iter*scale)), r.iter)
	}
}

// INPUT //
func askInput() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan Ukuran data : ")
	text, _ := reader.ReadString('\n')

	fields := strings.Fields(text)
	sizes := []int{}

	for _, f := range fields {
		if n, err := strconv.Atoi(f); err == nil && n > 0 {
			sizes = append(sizes, n)
		}
	}
	return sizes
}

// MAIN //
func main() {
	fmt.Print("==========================\n")
	fmt.Printf("MIN_MERGE = %d\n", MIN_MERGE)
	sizes := askInput()
	if len(sizes) == 0 {
		fmt.Println(" Inputan Tidak Valid ")
		return
	}
	results := []Result{}
	fmt.Println("Benchmark Sedang Berjalan... ")
	fmt.Println("--------------------------")
	for _, n := range sizes {
		data := generateData(n)
		recArr := copySlice(data)
		recTime :=  measureTime(func() {
			timSortRecursive(recArr)
		})
		iterArr := copySlice(data)
		iterTime := measureTime(func() {
			timSortIterative(iterArr)
		})
		results = append(results, Result{n, recTime, iterTime})
		fmt.Printf("✓ n = %d selesai\n", n)
	}

	printTable(results)
	printBarChart(results)
}