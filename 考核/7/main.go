package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

func main() {
	data := make(map[int]mapset.Set)
	var nRows int
	fmt.Scanf("%d", &nRows)

	for i := 0; i < nRows; i++ {
		var time int
		fmt.Scanf("%d", &time)

		data[time] = mapset.NewSet()

		var nQuality int
		fmt.Scanf("%d", &nQuality)

		for i := 0; i < nQuality; i++ {
			var quality int
			fmt.Scanf("%d", &quality)
			data[time].Add(quality)
		}

		set := mapset.NewSet()
		for k, v := range data {
			if k <= time && k > time-86400 {
				iter := v.Iter()
				for i := 0; i < v.Cardinality(); i++ {
					set.Add(<-iter)
				}
			}
		}

		fmt.Println(set.Cardinality())
	}

}
