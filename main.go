package main

import (
	"fmt"
	"math"
)

// Struktur graf berbobot (adjacency list)
type Graph map[string]map[string]int

// Algoritma Dijkstra
func dijkstra(graph Graph, start string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	prev := make(map[string]string)
	visited := make(map[string]bool)

	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	for {
		minDist := math.MaxInt32
		current := ""

		for node := range graph {
			if !visited[node] && dist[node] < minDist {
				minDist = dist[node]
				current = node
			}
		}

		if current == "" {
			break
		}

		visited[current] = true

		for neighbor, weight := range graph[current] {
			newDist := dist[current] + weight
			if newDist < dist[neighbor] {
				dist[neighbor] = newDist
				prev[neighbor] = current
			}
		}
	}

	return dist, prev
}

// Rekonstruksi rute
func getPath(prev map[string]string, start, end string) []string {
	path := []string{end}
	for end != start {
		end = prev[end]
		path = append([]string{end}, path...)
	}
	return path
}

func main() {
	// ==========================
	// Dataset Graf
	// ========================
	graph := map[string]map[string]int{
		"Kantor": {"A": 4, "B": 2},
		"A":      {"B": 4, "C": 2},
		"B":      {"A": 4, "C": 1, "D": 5},
		"C":      {"A": 2, "B": 1, "D": 8, "E": 10},
		"D":      {"B": 5, "C": 8, "E": 2, "F": 6},
		"E":      {"C": 10, "D": 2, "F": 3},
		"F":      {"D": 6, "E": 3},
	}

	// ==========================
	// Input User
	// ==========================
	var start, target string

	fmt.Println("=== Optimasi Rute Petugas Lapangan ===")
	fmt.Println("Daftar Lokasi:")
	for node := range graph {
		fmt.Println("-", node)
	}
	fmt.Println("masukan dengan huruf kapital")
	fmt.Print("\nMasukkan titik awal   : ")
	fmt.Scanln(&start)

	fmt.Print("Masukkan titik tujuan : ")
	fmt.Scanln(&target)
	// Jalankan Dijkstra
	dist, prev := dijkstra(graph, start)

	// Validasi tujuan
	if dist[target] == math.MaxInt32 {
		fmt.Println("\n❌ Rute tidak ditemukan.")
		return
	}

	// Ambil rute
	path := getPath(prev, start, target)

	// Output Hasil
	fmt.Println("\n=== HASIL ===")
	fmt.Println("Titik Awal   :", start)
	fmt.Println("Titik Tujuan :", target)
	fmt.Println("Rute Optimal :", path)
	fmt.Println("Total Jarak  :", dist[target], "km")
}
