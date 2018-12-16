package main

import "testing"

func TestGetCategories(t *testing.T) {
	actual := parseCategories([]Entry{
		Entry{Category: "A"},
		Entry{Category: "B"},
		Entry{Category: "B"},
		Entry{Category: "C"},
		Entry{Category: "D"},
	})
	expected := []string{"A", "B", "C", "D"}
	if len(actual) != len(expected) {
		t.Fatalf("bad parsing: expected %v, got %v", expected, actual)
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("bad element: expected %q, got %q", actual[i], expected[i])
		}
	}
}

func TestCheckEntryMatches(t *testing.T) {
	entry := Entry{
		API:         "examplesAsAService",
		Description: "provide classic examples",
		Auth:        "apiKey",
		HTTPS:       true,
		Cors:        "Unknown",
		Link:        "http://www.example.com",
		Category:    "Development",
	}
	entryEmptyAuth := Entry{
		API:         "examplesAsAService",
		Description: "provide classic examples",
		Auth:        "",
		HTTPS:       true,
		Cors:        "Unknown",
		Link:        "http://www.example.com",
		Category:    "Development",
	}

	testCases := []struct {
		name       string
		entry      Entry
		search     *SearchRequest
		shouldPass bool
	}{
		{"Full search", entry, &SearchRequest{}, true},
		{"Desc valid full", entry, &SearchRequest{Description: "provide classic examples"}, true},
		{"Desc valid match", entry, &SearchRequest{Description: "provide class"}, true},
		{"Desc invalid", entry, &SearchRequest{Description: "this will not match"}, false},
		{"Auth valid full", entry, &SearchRequest{Auth: "apiKey"}, true},
		{"Auth valid match", entry, &SearchRequest{Auth: "apiK"}, true},
		{"Auth empty", entry, &SearchRequest{Auth: ""}, true},
		{"Auth empty entry", entryEmptyAuth, &SearchRequest{Auth: ""}, true},
		{"Auth null", entry, &SearchRequest{Auth: "null"}, false},
		{"Auth null empty entry", entryEmptyAuth, &SearchRequest{Auth: "null"}, true},
		{"Auth invalid", entry, &SearchRequest{Auth: "foo"}, false},
		{"HTTPS true", entry, &SearchRequest{HTTPS: "1"}, true},
		{"HTTPS false", entry, &SearchRequest{HTTPS: "false"}, false},
		{"CORS valid full", entry, &SearchRequest{Cors: "unknown"}, true},
		{"CORS valid match", entry, &SearchRequest{Cors: "unk"}, true},
		{"CORS invalid", entry, &SearchRequest{Cors: "bar"}, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if checkEntryMatches(tc.entry, tc.search) != tc.shouldPass {
				if tc.shouldPass {
					t.Errorf("was expecting to pass, but failed")
				} else {
					t.Errorf("was expecting to fail, but passed")
				}
			}
		})
	}
}
