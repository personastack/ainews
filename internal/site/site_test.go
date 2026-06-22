package site

import (
	"encoding/json"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/personadock/ainews/internal/content"
)

func TestIndexIncludesPublishedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	if !strings.Contains(body, template.HTMLEscapeString("Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?")) {
		t.Fatal("response missing enterprise AI ROI gap article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The Smartest Model in Your Stack Might Be the Smallest")) {
		t.Fatal("response missing retrieval layer small embedding models article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The Real Test for AI Agents Isn't Autonomy — It's Whether They Can Check Their Own Work")) {
		t.Fatal("response missing agentic AI verification oracle article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Two Roads, One Month: The EU Tightens Its AI Rulebook as Washington Moves to Tear Up the States'")) {
		t.Fatal("response missing EU AI Act and US state preemption article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The Wire Became the Bottleneck — So AI Is Rebuilding It Out of Light")) {
		t.Fatal("response missing AI silicon photonics interconnect article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("AI Designed the Molecule in Months — The Clinic Still Takes Years")) {
		t.Fatal("response missing AI drug discovery clinic article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("An Open-Weights Model Just Caught the Frontier on Coding — at One-Sixth the Price")) {
		t.Fatal("response missing GLM-5.2 open-weights coding article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Microsoft Put a Meter on Its AI. Then It Went Shopping for a Cheaper Engine.")) {
		t.Fatal("response missing Microsoft AI cost meter article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("When the AI Starts Building the AI")) {
		t.Fatal("response missing AI builds itself article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Reasoning Models Were Built to Think Longer. 2026 Is Teaching Them When to Stop.")) {
		t.Fatal("response missing reasoning models test-time compute article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The Chip Stopped Being the Bottleneck — Now It's Power and Memory")) {
		t.Fatal("response missing AI power and memory bottleneck article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("World Models Grew Up: AI Stopped Generating Scenes and Started Predicting Actions")) {
		t.Fatal("response missing world models physical AI article title")
	}
	posts := content.Posts()
	for i := 0; i < postsPerPage; i++ {
		want := template.HTMLEscapeString(posts[i].Title)
		if !strings.Contains(body, want) {
			t.Fatalf("response missing page-one title %q", want)
		}
	}

	if got := strings.Count(body, "<article>"); got != postsPerPage {
		t.Fatalf("article count = %d, want %d", got, postsPerPage)
	}

	if strings.Contains(body, template.HTMLEscapeString(posts[postsPerPage].Title)) {
		t.Fatalf("response included first page-two title %q", posts[postsPerPage].Title)
	}

	for _, want := range []string{
		"Showing 1-12 of ",
		"Page 1 of ",
		`href="/?page=2"`,
		"Older posts",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing pagination text %q", want)
		}
	}
}

func TestIndexSecondPageIncludesNextSlice(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/?page=2", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	posts := content.Posts()
	if strings.Contains(body, template.HTMLEscapeString(posts[0].Title)) {
		t.Fatalf("response included newest page-one title %q", posts[0].Title)
	}

	wantTitle := template.HTMLEscapeString(posts[postsPerPage].Title)
	if !strings.Contains(body, wantTitle) {
		t.Fatalf("response missing first page-two title %q", wantTitle)
	}

	for _, want := range []string{
		"Showing 13-24 of ",
		"Page 2 of ",
		`href="/"`,
		`href="/?page=3"`,
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing pagination text %q", want)
		}
	}
}

func TestPostRoute(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/the-frontier-firm-is-here-microsoft-says-ai-has-moved-from-tool-to-operating-model", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if !strings.Contains(rec.Body.String(), "Copilot Cowork") {
		t.Fatalf("response did not render article body")
	}
}

func TestPostRouteRendersRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/retrieval-layer-small-embedding-models-rag-accuracy-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-cost-meter-copilot-cowork-deepseek-2026"`,
		`href="/posts/ai-startups-infrastructure-not-chatbots-2026"`,
		template.HTMLEscapeString("Microsoft Put a Meter on Its AI. Then It Went Shopping for a Cheaper Engine."),
		template.HTMLEscapeString("The Next AI Startup Wave Is Infrastructure, Not Chatbots"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostsAPI(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/posts", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var posts []map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &posts); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	wantPosts := content.Posts()
	if len(posts) != len(wantPosts) {
		t.Fatalf("len(posts) = %d, want %d", len(posts), len(wantPosts))
	}

	if got := posts[0]["slug"]; got != wantPosts[0].Slug {
		t.Fatalf("first post slug = %q, want %q", got, wantPosts[0].Slug)
	}

	if got := posts[1]["slug"]; got != wantPosts[1].Slug {
		t.Fatalf("second post slug = %q, want %q", got, wantPosts[1].Slug)
	}

	if got := posts[2]["slug"]; got != wantPosts[2].Slug {
		t.Fatalf("third post slug = %q, want %q", got, wantPosts[2].Slug)
	}

	if got := posts[3]["slug"]; got != wantPosts[3].Slug {
		t.Fatalf("fourth post slug = %q, want %q", got, wantPosts[3].Slug)
	}

	if got := posts[4]["slug"]; got != wantPosts[4].Slug {
		t.Fatalf("fifth post slug = %q, want %q", got, wantPosts[4].Slug)
	}
}

func TestHealthz(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if got := strings.TrimSpace(rec.Body.String()); got != "ok" {
		t.Fatalf("body = %q, want ok", got)
	}
}
