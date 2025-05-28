package data

import "testing"


func TestReadJSON(t *testing.T) {
	test_cases := []struct {
		fruit     string
		size      string
		color     string
		path     string


	}{
		{"apple", "large", "red", "http://example.com/data.json"},
		{"banana", "medium", "yellow", "https://example.com/data.json"},
		{"grape", "small", "purple", "../json_file.json"},
	}

	for _, tc := range test_cases {
		_, err := readJSON(tc.path)
		if err != nil {
			t.Errorf("Failed to read JSON from %s: %v", tc.path, err)
		} else {
			t.Logf("Successfully read JSON from %s", tc.path)
		}
	}
}

