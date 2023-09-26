package main

import "fmt"

func main20() {
	res := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(res)
}

func findDisappearedNumbers(nums []int) []int {
	key := make([]int, len(nums)+1)

	// Проходим по входному массиву и увеличиваем соответствующий ключ в key[]
	// для каждого числа, которое встречаем в nums
	for _, n := range nums {
			key[n]++
	}

	res := []int{}
	// Проходим по key[] и добавляем индексы с нулевым значением в res[],
	// таким образом, находим числа, которые отсутствуют в исходном массиве
	for i := 1; i < len(key); i++ {
			if key[i] == 0 {
					res = append(res, i)
			}
	}

	return res
}

func FindMissingValues(nums []int) []int {
	dict := make(map[int]int)
	
	//list := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dict[i+1] = i+1
		//list[i] = i+1
	}
	for _, num := range nums {
		delete(dict, num)
		//list = append(list[:num], list[num+1:]...)
	}

	keys := []int{}
  for key, _ := range dict {
      keys = append(keys, key)
  }
	return keys
}
