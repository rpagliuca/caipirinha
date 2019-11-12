package caipirinha

import (
	"fmt"
	"github.com/kr/pretty"
)

func pivot(data []map[string]interface{}, groupBy []string, accumulator string) []map[string]interface{} {
	for _, group := range groupBy {
		data = sort(data, group)
	}
	out := make([]map[string]interface{}, 0)
	totals := make(map[int]float64, 0)
	var previous map[string]interface{}
	data = append(data, nil)
	for k, row := range data {
		pretty.Println("row", k)
		changed := false
		newRow := make(map[string]interface{}, 0)
		for i, group := range groupBy {
			newRow[group] = previous[group]
			if previous != nil && (changed || row[group] != previous[group]) {
				pretty.Println(fmt.Sprintf("totals[%d]", i), totals[i])
				newRow[accumulator] = totals[i]
				out = append(out, newRow)
				pretty.Println("out", out)
				oldRow := newRow
				pretty.Println("oldRow", oldRow)
				newRow = make(map[string]interface{}, 0)
				for k, v := range oldRow {
					newRow[k] = v
				}
				totals[i] = 0
				changed = true
			}
			if row != nil {
				totals[i] += row[accumulator].(float64)
			}
			pretty.Println(totals)
		}
		previous = row
	}
	for _, group := range groupBy {
		out = sort(out, group)
	}
	return out
}

func sort(data []map[string]interface{}, key string) []map[string]interface{} {
	for i := range data {
		for j := i+1; j < len(data); j++ {
			greater := false
			if data[j][key] == nil {
				greater = true
			} else {
				switch val := data[i][key].(type) {
					case string:
						greater = val > data[j][key].(string)
					case int:
						greater = val > data[j][key].(int)
					case float64:
						greater = val > data[j][key].(float64)
					case float32:
						greater = val > data[j][key].(float32)
					default:
						greater = false
				}
			}
			if greater {
				tmp := data[j]
				data[j] = data[i]
				data[i] = tmp
			}
		}
	}
	return data
}
