package caipirinha

import (
	"fmt"
	"testing"
	"errors"
)

func TestPivot(t *testing.T) {
	type testCase struct {
		data []map[string]interface{}
		groupBy []string
		accumulators []string
		expected []map[string]interface{}
	}

	data := []map[string]interface{} {
		{ "col1": "v1", "col2": "v2", "col3": "v3", "quantity": 2.5, "other_quantity": 1.0 },
		{ "col1": "v2", "col2": "v2", "col3": "v4", "quantity": 7.5, "other_quantity": 2.0 },
		{ "col1": "v1", "col2": "v1", "col3": "v3", "quantity": 1.5, "other_quantity": 3.0 },
		{ "col1": "v2", "col2": "v3", "col3": "v4", "quantity": 1.0, "other_quantity": 4.0 },
		{ "col1": "v2", "col2": "v2", "col3": "v3", "quantity": 2.5, "other_quantity": 5.0 },
	}

	data2 := []map[string]interface{} {
		{ "col1": "v1", "col2": "v1", "quantity": 2.5, "other_quantity": 1.0 },
		{ "col1": "v1", "col2": "v2", "quantity": 7.5, "other_quantity": 2.0 },

	}

	testCases := []testCase {
		{
			data,
			[]string {
				"col1", "col2", "col3",
			},
			[]string {
				"quantity", "other_quantity",
			},
			[]map[string]interface{} {
				{ "col1": "v1", "quantity": 4.0, "other_quantity": 4.0 },
				{ "col1": "v1", "col2": "v1", "quantity": 1.5, "other_quantity": 3.0 },
				{ "col1": "v1", "col2": "v1", "col3": "v3", "quantity": 1.5, "other_quantity": 3.0 },
				{ "col1": "v1", "col2": "v2", "quantity": 2.5, "other_quantity": 1.0 },
				{ "col1": "v1", "col2": "v2", "col3": "v3", "quantity": 2.5, "other_quantity": 1.0 },
				{ "col1": "v2", "quantity": 11.0, "other_quantity": 11.0},
				{ "col1": "v2", "col2": "v2", "quantity": 10.0, "other_quantity": 7.0 },
				{ "col1": "v2", "col2": "v2", "col3": "v3", "quantity": 2.5, "other_quantity": 5.0 },
				{ "col1": "v2", "col2": "v2", "col3": "v4", "quantity": 7.5, "other_quantity": 2.0 },
				{ "col1": "v2", "col2": "v3", "quantity": 1.0, "other_quantity": 4.0 },
				{ "col1": "v2", "col2": "v3", "col3": "v4", "quantity": 1.0, "other_quantity": 4.0 },
			},
		},
		{
			data,
			[]string {
				"col1",
			},
			[]string {
				"quantity",
			},
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
			[]string {
				"quantity",
			},
			[]map[string]interface{} {
				{ "col1": "v1", "quantity": 10.0 },
				{ "col1": "v1", "col2": "v1", "quantity": 2.5 },
				{ "col1": "v1", "col2": "v2", "quantity": 7.5 },
			},
		},

		{
			data,
			[]string {
				"col1",
				"col2",
			},
			[]string {
				"quantity",
			},
			[]map[string]interface{} {
				{ "col1": "v1", "quantity": 4.0},
				{ "col1": "v1", "col2": "v1", "quantity": 1.5},
				{ "col1": "v1", "col2": "v2", "quantity": 2.5},
				{ "col1": "v2", "quantity": 11.0},
				{ "col1": "v2", "col2": "v2", "quantity": 10.0},
				{ "col1": "v2", "col2": "v3", "quantity": 1.0},
			},
		},

		{
			data,
			[]string {
				"col2",
				"col1",
			},
			[]string {
				"quantity",
			},
			[]map[string]interface{} {
				{ "col2": "v1", "quantity": 1.5},
				{ "col2": "v1", "col1": "v1", "quantity": 1.5},
				{ "col2": "v2", "quantity": 12.5},
				{ "col2": "v2", "col1": "v1", "quantity": 2.5},
				{ "col2": "v2", "col1": "v2", "quantity": 10.0},
				{ "col2": "v3", "quantity": 1.0},
				{ "col2": "v3", "col1": "v2", "quantity": 1.0},
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		got := Pivot(tc.data, tc.groupBy, tc.accumulators)
		err := assertSlicesEqual(got, tc.expected)
		if err != nil {
			t.Error(fmt.Sprintf("Test case %d: ", i), err.Error())
			return
		}
	}
}

func TestPivot2(t *testing.T) {
	type testCase struct {
		data []map[string]interface{}
		groupBy []string
		accumulators []string
		expected []map[string]interface{}
	}

	data := []map[string]interface{} {
		{ "col1": "v1", "col2": nil, "col3": nil, "quantity": 17.0},
		{ "col1": "v1", "col2": "v2", "col3": "v3", "quantity": 2.5},
		{ "col1": "v2", "col2": "v2", "col3": "v4", "quantity": 7.5},
		{ "col1": "v1", "col2": "v1", "col3": "v3", "quantity": 1.5},
		{ "col1": "v2", "col2": "v3", "col3": "v4", "quantity": 1.0},
		{ "col1": "v2", "col2": "v2", "col3": "v3", "quantity": 2.5},
	}

	testCases := []testCase {
		{
			data,
			[]string {
				"col1", "col2", "col3",
			},
			[]string {
				"quantity",
			},
			[]map[string]interface{} {
				{ "col1": "v1", "quantity": 21.0},
				{ "col1": "v1", "col2": nil, "quantity": 17.0},
				{ "col1": "v1", "col2": nil, "col3": nil, "quantity": 17.0},
				{ "col1": "v1", "col2": "v1", "quantity": 1.5},
				{ "col1": "v1", "col2": "v1", "col3": "v3", "quantity": 1.5},
				{ "col1": "v1", "col2": "v2", "quantity": 2.5},
				{ "col1": "v1", "col2": "v2", "col3": "v3", "quantity": 2.5},
				{ "col1": "v2", "quantity": 11.0},
				{ "col1": "v2", "col2": "v2", "quantity": 10.0},
				{ "col1": "v2", "col2": "v2", "col3": "v3", "quantity": 2.5},
				{ "col1": "v2", "col2": "v2", "col3": "v4", "quantity": 7.5},
				{ "col1": "v2", "col2": "v3", "quantity": 1.0},
				{ "col1": "v2", "col2": "v3", "col3": "v4", "quantity": 1.0},
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		got := Pivot(tc.data, tc.groupBy, tc.accumulators)
		err := assertSlicesEqual(got, tc.expected)
		if err != nil {
			fmt.Printf("got: %+v\n", got)
			fmt.Printf("expected: %+v\n", tc.expected)
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

func TestSort(t *testing.T) {
	type TestCase struct {
		data []map[string]interface{}
		keys []string
		expected []map[string]interface{}
	}

	data1 := []map[string]interface{}{
		{"col1": 1},
		{"col1": 2},
		{},
		{"col1": 0},
	}

	expected1 := []map[string]interface{}{
		{},
		{"col1": 0},
		{"col1": 1},
		{"col1": 2},
	}

	data2 := []map[string]interface{}{
		{"col1": 1, "col2": 1},
		{"col1": 2, "col2": 1},
		{},
		{"col1": 1},
		{"col1": 2},
	}

	expected2a := []map[string]interface{}{
		{},
		{"col1": 1},
		{"col1": 2},
		{"col1": 2, "col2": 1},
		{"col1": 1, "col2": 1},
	}

	expected2b := []map[string]interface{}{
		{},
		{"col1": 1},
		{"col1": 1, "col2": 1},
		{"col1": 2},
		{"col1": 2, "col2": 1},
	}

	testCases := []TestCase {
		TestCase{data1, []string{"col1"}, expected1},
		TestCase{data2, []string{"col2"}, expected2a},
		TestCase{data2, []string{"col1", "col2"}, expected2b},
	}

	for _, tc := range testCases {
		got := Sort(tc.data, tc.keys)
		err := assertSlicesEqual(got, tc.expected)
		if err != nil {
			t.Error("Error sorting")
		}
	}
}
