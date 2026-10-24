package site

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/personadock/ainews/internal/content"
)

func TestPostRendersStructuredSectionsSafely(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	server.render(recorder, http.StatusOK, "post", postPageData{Post: content.Post{
		Sections: []content.Section{{
			Paragraphs: []string{"Before the list"},
			Items:      []string{"First <script>point</script>", "Second & final"},
			Sources: []content.Source{
				{Title: "Public <report>", URL: "https://example.com/report?a=1&b=2"},
				{Title: "Unsafe source", URL: "javascript:alert(1)"},
			},
		}},
	}})
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d", recorder.Code)
	}
	body := recorder.Body.String()
	previous := -1
	for _, want := range []string{
		"<p>Before the list</p>",
		"<li>First &lt;script&gt;point&lt;/script&gt;</li>",
		"<li>Second &amp; final</li>",
		`<ul class="sources">`,
		`<a href="https://example.com/report?a=1&amp;b=2">Public &lt;report&gt;</a>`,
		`<a href="#ZgotmplZ">Unsafe source</a>`,
	} {
		position := strings.Index(body, want)
		if position <= previous {
			t.Errorf("rendered article missing or reordered %q", want)
		}
		previous = position
	}
}
