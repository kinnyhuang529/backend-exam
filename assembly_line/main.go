package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Employee 員工
type Employee struct {
	ID        int // 員工編號
	ItemCount int // 處理物品數量
}

// Item1 物品1 (100ms)
type Item1 struct {
	ID int
}

func (i Item1) Process() {
	time.Sleep(100 * time.Millisecond)
}

// Item2 物品2 (200ms)
type Item2 struct {
	ID int
}

func (i Item2) Process() {
	time.Sleep(200 * time.Millisecond)
}

// Item3 物品3 (300ms)
type Item3 struct {
	ID int
}

func (i Item3) Process() {
	time.Sleep(300 * time.Millisecond)
}

type Item interface {
	// Process 這是一個耗時操作
	Process()
}

func main() {
	// 初始化5個員工
	employees := make([]*Employee, 5)
	for i := 0; i < 5; i++ {
		employees[i] = &Employee{ID: i + 1}
	}

	// 初始化3種物品
	var items []Item
	for i := 1; i <= 10; i++ {
		items = append(items, Item1{ID: i}, Item2{ID: i}, Item3{ID: i})
	}

	// 打亂物品順序
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	// 建立流水線 Channel
	pipeline := make(chan Item, len(items))
	var wg sync.WaitGroup
	startTime := time.Now()

	// --------------------

	// 把員工放到崗位 開始工作囉
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(e *Employee) {
			defer wg.Done()
			for item := range pipeline {
				// 打印開始記錄
				fmt.Printf("[員工 %d] 開始處理: %v\n", e.ID, item)

				item.Process()
				e.ItemCount++

				// 打印結束記錄
				fmt.Printf("[員工 %d] 完成處理: %v\n", e.ID, item)
			}
		}(employees[i])
	}

	// 把物品丟進流水線
	for _, itm := range items {
		pipeline <- itm
	}
	close(pipeline)

	// 等員工們完成
	wg.Wait()
	totalTime := time.Since(startTime)

	// 統計結果
	fmt.Println("------------------------------")
	fmt.Println("總處理時間: ", totalTime)
	for _, e := range employees {
		fmt.Printf("員工 %d 處理了 %d 件物品\n", e.ID, e.ItemCount)
	}
	fmt.Println("------------------------------")
}
