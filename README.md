# Tugas Besar Analisis Kompleksitas Algoritma

Tugas Besar Mata Kuliah **Analisis Kompleksitas Algoritma** ini membahas
analisis performa algoritma pengurutan **Tim Sort**, dengan fokus pada
**perbandingan pendekatan rekursif dan iteratif** dalam proses
pengurutan dan penggabungan data.

## Deskripsi Umum

**Tim Sort** adalah algoritma pengurutan berbasis *comparison* yang
merupakan gabungan dari **Insertion Sort** dan **Merge Sort**. Algoritma
ini dirancang untuk bekerja sangat efisien pada data dunia nyata
(*real-world data*), khususnya data yang sudah memiliki sebagian urutan
(*partially sorted*).

Tim Sort digunakan secara luas pada bahasa pemrograman modern seperti
**Python** dan **Java**, karena mampu memberikan performa optimal pada
berbagai kondisi input.

## Prinsip Kerja Tim Sort

Secara umum, Tim Sort bekerja melalui beberapa tahapan utama:

1.  **Pembentukan Run**
    -   Data dibagi menjadi bagian-bagian kecil yang disebut *run*
    -   Setiap run memiliki panjang minimum tertentu (*minRun*)
    -   Run diurutkan menggunakan **Insertion Sort (iteratif)**
2.  **Penyimpanan Run dalam Stack**
    -   Run yang telah terurut disimpan ke dalam sebuah *stack*
    -   Stack ini digunakan untuk mengatur urutan proses penggabungan
        (*merge*)
3.  **Proses Merge**
    -   Run-run digabungkan menggunakan teknik **Merge Sort**
    -   Proses merge dapat dilakukan secara:
        -   **Iteratif** (menggunakan loop dan stack)
        -   **Rekursif** (pemanggilan fungsi merge secara berulang)
4.  **Galloping Mode**
    -   Optimisasi ketika salah satu run lebih dominan
    -   Menggunakan binary search untuk mempercepat merge

## Perbandingan Rekursif dan Iteratif

### Pendekatan Iteratif
-   Menggunakan loop dan stack
-   Lebih hemat memori
-   Tidak menggunakan call stack

### Pendekatan Rekursif
-   Menggunakan pemanggilan fungsi berulang
-   Lebih mudah dipahami
-   Membutuhkan memori tambahan

## Analisis Kompleksitas

  Kasus          Kompleksitas
  -------------- --------------
  Best Case      O(n)
  Average Case   O(n log n)
  Worst Case     O(n log n)

Kompleksitas ruang: - O(n) untuk merge - Tambahan call stack pada
pendekatan rekursif

## Tujuan Tugas Besar
1.  Menganalisis kompleksitas Tim Sort
2.  Membandingkan rekursif dan iteratif
3.  Mengkaji efisiensi algoritma
4.  Memahami optimisasi Tim Sort

## Kesimpulan
Tim Sort adalah algoritma adaptif yang sangat efisien untuk data nyata
dan memberikan performa optimal dibandingkan algoritma pengurutan
klasik.

### Anggota Kelompok
- Muhammad Rasyid Ridho - 103122400018
- Rahmadanis Danang Kumala - 103122400066
- Fatikhah Sukma Arti - 103122400019
