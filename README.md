# Caipirinha

This library allows you to calculate subtotals of a dataset

## Usage

    data := []map[string]interface{} {
        { "col1": "v1", "col2": "v2", "col3": "v3", "quantity": 2.5 },
        { "col1": "v2", "col2": "v2", "col3": "v4", "quantity": 7.5 },
        { "col1": "v1", "col2": "v1", "col3": "v3", "quantity": 1.5 },
        { "col1": "v2", "col2": "v3", "col3": "v4", "quantity": 1.0 },
        { "col1": "v2", "col2": "v2", "col3": "v3", "quantity": 2.5 },
    }

    summary := Pivot(data, []string{"col1"}, "quantity")

    /*
    []map[string]interface{} {
        { "col1": "v1", "quantity": 4.0 },
        { "col1": "v2", "quantity": 11.0 },
    }
    /*
