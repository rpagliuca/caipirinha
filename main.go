package caipirinha

func Pivot(data []map[string]interface{}, groupBy []string, accumulators []string) []map[string]interface{} {
	data = Sort(data, groupBy)
	out := make([]map[string]interface{}, 0)
	totals := make(map[string]map[int]float64, 0)
	for _, accumulator := range accumulators {
		totals[accumulator] = make(map[int]float64, 0)
	}
	var previous map[string]interface{}
	data = append(data, nil)
	for _, row := range data {
		changed := false
		newRow := make(map[string]interface{}, 0)
		for i, group := range groupBy {
			newRow[group] = previous[group]
			if previous != nil && (changed || row[group] != previous[group]) {
				for _, accumulator := range accumulators {
					newRow[accumulator] = totals[accumulator][i]
				}
				out = append(out, newRow)
				oldRow := newRow
				newRow = make(map[string]interface{}, 0)
				for k, v := range oldRow {
					newRow[k] = v
				}
				for _, accumulator := range accumulators {
					totals[accumulator][i] = 0
				}
				changed = true
			}
			if row != nil {
				for _, accumulator := range accumulators {
					totals[accumulator][i] += row[accumulator].(float64)
				}
			}
		}
		previous = row
	}
	out = Sort(out, groupBy)
	return out
}

func Sort(data []map[string]interface{}, keys []string) []map[string]interface{} {
	for i := range data {
		for j := i+1; j < len(data); j++ {
			var greater bool
			out:
			for _, key := range keys {
				_, oki := data[i][key]
				_, okj := data[j][key]
				if !oki && okj {
					// We do not have column in record i, but we have in record j,
					// so record i cannot be positioned before record j
					greater = false
					break out
				} else if !oki && !okj {
					// We do not have column in record i, neither record j,
					// so we cannot infer anything from this column
					continue
				} else if oki && !okj {
					// We have column in record i, but not in record j,
					// so we are sure that record i should be positioned before record j
					greater = true
					break out
				} else {
					if data[i][key] == nil && data[j][key] != nil {
						// Value for column in record i is nil, and in record j is not nil,
						// so we are sure that record i cannot be positioned before record j
						greater = false
						break out
					} else if data[i][key] == nil && data[j][key] == nil {
						// Both values are nil,
						// so we cannot infer anything from this column
						continue
					} else if data[i][key] != nil && data[j][key] == nil {
						// Value for column in record i is not nil, and in record j is nil,
						// so we are sure that record i should be positioned before record j
						greater = true
						break out
					}
					switch val := data[i][key].(type) {
						case string:
							if val > data[j][key].(string) {
								greater = true
								break out
							} else if val < data[j][key].(string) {
								greater = false
								break out
							}
						case int:
							if (val > data[j][key].(int)) {
								greater = true
								break out
							} else if (val < data[j][key].(int)) {
								greater = false
								break out
							}
						case float64:
							if (val > data[j][key].(float64)) {
								greater = true
								break out
							} else if (val < data[j][key].(float64)) {
								greater = false
								break out
							}
						case float32:
							if (val > data[j][key].(float32)) {
								greater = true
								break out
							} else if (val < data[j][key].(float32)) {
								greater = false
								break out
							}
					}
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
