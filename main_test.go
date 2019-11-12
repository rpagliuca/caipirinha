package caipirinha

import (
	"fmt"
	"testing"
	"errors"
	"github.com/kr/pretty"
)

func TestPivot(t *testing.T) {
	type testCase struct {
		data []map[string]interface{}
		groupBy []string
		accumulator string
		expected []map[string]interface{}
	}

	data := []map[string]interface{} {
		{ "col1": "v1", "col2": "v2", "col3": "v3", "quantity": 2.5 },
		{ "col1": "v2", "col2": "v2", "col3": "v4", "quantity": 7.5 },
		{ "col1": "v1", "col2": "v1", "col3": "v3", "quantity": 1.5 },
		{ "col1": "v2", "col2": "v3", "col3": "v4", "quantity": 1.0 },
		{ "col1": "v2", "col2": "v2", "col3": "v3", "quantity": 2.5 },
	}

	data2 := []map[string]interface{} {
		{ "col1": "v1", "col2": "v1", "quantity": 2.5 },
		{ "col1": "v1", "col2": "v2", "quantity": 7.5 },

	}

	testCases := []testCase {
		{
			data,
			[]string {
				"col1",
			},
			"quantity",
			[]map[string]interface{} {
				{ "col1": "v1", "quantity": 4.0 },
				{ "col1": "v2", "quantity": 11.0 },
			},
		},

		{
			data2,
			[]string {
				"col1", "col2",
			},
			"quantity",
			[]map[string]interface{} {
				{ "col1": "v1", "quantity": 10.0 },
				{ "col1": "v1", "col2": "v1", "quantity": 2.5 },
				{ "col1": "v1", "col2": "v2", "quantity": 7.5 },
			},

			/*
			c1v1 2.5
			c2v1 2.5

			c1 10.0
			c2 7.5 -> "c1v1 2.5 c2v1 2.5"

			c1 nil -> "c1v1 10.0"
			c2 nil -> "c1v1 10.0 c2v1 7.5"
			*/
		},

		//{
		//	data,
		//	[]string {
		//		"col1",
		//		"col2",
		//	},
		//	"quantity",
		//	[]map[string]interface{} {
		//		{ "col1": "v1", "col2": "v1", "quantity": 1.5},
		//		{ "col1": "v1", "col2": "v2", "quantity": 2.5},
		//		{ "col1": "v1", "quantity": 4.0},
		//		{ "col1": "v2", "col2": "v2", "quantity": 10.0},
		//		{ "col1": "v2", "col2": "v3", "quantity": 1.0},
		//		{ "col1": "v2", "quantity": 11.0},
		//	},
		//},
	}
	for i := range testCases {
		c := testCases[i]
		pretty.Println("Test case" , i)
		got := pivot(c.data, c.groupBy, c.accumulator)
		err := assertSlicesEqual(got, c.expected)
		if err != nil {
			pretty.Println("got", got)
			pretty.Println("expected", c.expected)
			t.Error(fmt.Sprintf("Test case %d: ", i), err.Error())
			return
		}
	}
}

func assertSlicesEqual(slice1 []map[string]interface{}, slice2 []map[string]interface{}) error {
	if len(slice1) != len(slice2) {
		return errors.New(fmt.Sprintf("Length of slice1 (%v) != length of slice 2 (%v)\n", len(slice1), len(slice2)))
	}
	for i := range slice1 {
		if len(slice1[i]) != len(slice2[i]) {
			return errors.New(fmt.Sprintf("Length of row %v differs\n", i))
		}
		for key, _ := range slice1[i] {
			if slice1[i][key] != slice2[i][key] {
				return errors.New(fmt.Sprintf("Value %v differs from %v\n", slice1[i][key], slice2[i][key]))
			}
		}
		for key, _ := range slice2[i] {
			if slice1[i][key] != slice2[i][key] {
				return errors.New(fmt.Sprintf("Value %v differs from %v\n", slice1[i][key], slice2[i][key]))
			}
		}
	}
	return nil
}
