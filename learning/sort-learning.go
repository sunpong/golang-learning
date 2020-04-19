package main

import (
	"fmt"
	"sort"
)

type MyStringList []int

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {

	names := MyStringList{3, 2, 1, 4}
	sort.Sort(names)

	for _, v := range names {
		fmt.Println(v)
	}

}
