package main

import "fmt"

func main() {

	type Point struct {
		X int
		Y int
	}

	pointsMap := map[string]Point{} // объявили и проинициализировали

	pointsMap["a"] = Point{
		// добавление элементов в map
		X: 1,
		Y: 2,
	}
	pointsMap["b"] = Point{
		X: 3,
		Y: 4,
	}

	fmt.Println(pointsMap)        // map[a:{1 2} b:{3 4}]
	fmt.Println(pointsMap["a"])   // {1 2}
	fmt.Println(pointsMap["b"].X) // 3

	pointsMap2 := make(map[int]Point)
	pointsMap2[1] = Point{1, 2} // ри инициализации структуры могу просто написать значения, без ключей
	fmt.Println(pointsMap2)     // map[1:{1 2}]
	fmt.Println(pointsMap2[1])  // {1 2}
	// добавление элементов в map

	//проверяем на пустоту, если да то добавляем

	var pointsMap3 map[string]Point // объявлен но не инициализирован
	if pointsMap3 == nil {
		pointsMap3 = map[string]Point{
			"a": {1, 2},
		}
	}

	// есть ли в мапе значению по ключу
	key := "a"
	value, ok := pointsMap3[key]
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s does in map\n", key)
		fmt.Println(value)

	}

	//	 декодирование мапа в структуру. Из мапа pointsMap4, получить структуру Pointt

	type Pointt struct {
		X int
		Y int
	}

	pointsMap4 := map[string]int{
		"x": 1,
		"y": 2,
	}

}
