package main

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type Conditions struct {
	column string
	value  interface{}
}

type SplitInfo struct {
	splitData []string
	splitDta  [][]string
}

func main() {
	GetRandomDataWithCondition()
}

func GetRandomDataWithCondition() {
	df := dataframe.LoadRecords(
		[][]string{
			{"A", "B", "C", "D"},
			{"a", "4", "5.1", "true"},
			{"k", "5", "7.0", "true"},
			{"k", "4", "6.0", "true"},
			{"a", "2", "7.1", "false"},
		},
	)
	fil := df.Filter(
		dataframe.F{
			Colname:    "A",
			Comparator: series.Eq,
			Comparando: "a",
		},
		dataframe.F{
			Colname:    "B",
			Comparator: series.Greater,
			Comparando: 4,
		},
	)
	fil2 := fil.Filter(
		dataframe.F{
			Colname:    "D",
			Comparator: series.Eq,
			Comparando: true,
		},
	)
	fmt.Println("filter: ", fil)
	fmt.Println("filter2: ", fil2)
}
