package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements заполняет слайс заданного размера случайными положительными числами.
// Если размер равен нулю или отрицательный, результатом будет пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for index := range data {
		// случайное положительное целое число в рамках границ int
		data[index] = rand.Intn(int(^uint(0)>>1)) + 1
	}
	return data
}

// maximum находит наибольшее число в слайсе. Для пустого слайса возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks делит слайс на CHUNKS частей и ищет максимум в каждой отдельной горутиной.
// Для пустого слайса возвращает 0.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxima := make([]int, CHUNKS)
	chunkSize, remainder := len(data)/CHUNKS, len(data)%CHUNKS

	// примитив синхронизации
	var workers sync.WaitGroup
	// делим работу на 8 потоков
	workers.Add(CHUNKS)

	for chunkIndex := range CHUNKS {
		start := chunkIndex*chunkSize + min(chunkIndex, remainder)
		end := start + chunkSize
		if chunkIndex < remainder {
			end++
		}
		// запуск отдельной горутины
		go func(index int, chunk []int) {
			// уменьшает счетчик для группы
			defer workers.Done()
			maxima[index] = data[0]
			if len(chunk) > 0 {
				// максимум в чанке
				maxima[index] = maximum(chunk)
			}
		}(chunkIndex, data[start:end])
	}
	// ожидаем окончания всех потоков
	workers.Wait()

	// максимум из всех чанков
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxVal := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", maxVal, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	maxVal = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", maxVal, elapsed)
}
