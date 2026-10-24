package content

import "testing"

func TestClonePostPreservesAndIsolatesStructuredSections(t *testing.T) {
	original := Post{Sections: []Section{{
		Items:   []string{"First point"},
		Sources: []Source{{Title: "Public report", URL: "https://example.com/report"}},
	}}}
	cloned := clonePost(original)
	if len(cloned.Sections[0].Items) != 1 || len(cloned.Sections[0].Sources) != 1 {
		t.Fatal("clone dropped structured article content")
	}
	if cloned.Sections[0].Items[0] != "First point" || cloned.Sections[0].Sources[0] != original.Sections[0].Sources[0] {
		t.Fatal("clone changed structured article content")
	}
	cloned.Sections[0].Items[0] = "Changed point"
	cloned.Sections[0].Sources[0].Title = "Changed title"
	cloned.Sections[0].Sources[0].URL = "https://example.com/changed"
	if original.Sections[0].Items[0] != "First point" || original.Sections[0].Sources[0].Title != "Public report" || original.Sections[0].Sources[0].URL != "https://example.com/report" {
		t.Fatal("clone exposes mutable structured article content")
	}
}
