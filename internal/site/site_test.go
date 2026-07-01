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
	if !strings.Contains(body, template.HTMLEscapeString("AI's Quiet Second Frontier: Foundation Models Built for the Data That Actually Runs Your Business")) {
		t.Fatal("response missing tabular foundation models article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The First Big American AI Law Was Supposed to Take Effect Yesterday. It No Longer Exists.")) {
		t.Fatal("response missing Colorado AI Act repeal article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The Year Companies Were Told to Use All the AI They Wanted. Then the Bill Came.")) {
		t.Fatal("response missing AI cost reckoning tokenmaxxing article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Language's Frontier Is Locking Down. Robotics' Frontier Just Went Open.")) {
		t.Fatal("response missing NVIDIA Cosmos 3 open physical AI world model article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("OpenAI Shipped Its Most Powerful Model. Only 20 Companies — All Government-Approved — Can Use It.")) {
		t.Fatal("response missing OpenAI GPT-5.6 Sol government-gated release article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("OpenAI Built Its Own Chip in Nine Months. The Real Target Isn't Nvidia — It's the Inference Bill.")) {
		t.Fatal("response missing OpenAI Broadcom Jalapeno inference chip article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Samsung Banned ChatGPT in 2023. Now It's Handing Every Employee a Coding Agent.")) {
		t.Fatal("response missing Samsung Codex citizen software article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("A Frontier Model Every Two Weeks: The Real AI Story of 2026 Is the Pace, Not the Peak")) {
		t.Fatal("response missing AI model release firehose article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("AI Is Hunting a Magnet That Could Break China's Grip on Rare Earths — It Hasn't Caught One Yet")) {
		t.Fatal("response missing AI materials discovery rare-earth magnet article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them")) {
		t.Fatal("response missing AI agent spending governance gap article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Washington Wrote the Rulebook for Frontier AI — And the First Lab It Touched Is Suing")) {
		t.Fatal("response missing US AI national security executive order article title")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The AI Bottleneck Moved Off the Chip and Onto the Power Grid")) {
		t.Fatal("response missing AI power grid bottleneck article title")
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

func TestPostRouteRendersModelReleaseFirehoseRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/ai-model-release-firehose-cadence-eval-debt-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/glm-5-2-open-weights-frontier-coding-cost-2026"`,
		`href="/posts/ai-cost-meter-copilot-cowork-deepseek-2026"`,
		template.HTMLEscapeString("An Open-Weights Model Just Caught the Frontier on Coding — at One-Sixth the Price"),
		template.HTMLEscapeString("Microsoft Put a Meter on Its AI. Then It Went Shopping for a Cheaper Engine."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersGPTSolRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-gpt-5-6-sol-government-gated-frontier-release-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/us-ai-national-security-executive-order-anthropic-lawsuit-2026"`,
		`href="/posts/ai-model-release-firehose-cadence-eval-debt-2026"`,
		`href="/posts/fable-5-mythos-5-export-control-shutdown-2026"`,
		template.HTMLEscapeString("Washington Wrote the Rulebook for Frontier AI — And the First Lab It Touched Is Suing"),
		template.HTMLEscapeString("A Frontier Model Every Two Weeks: The Real AI Story of 2026 Is the Pace, Not the Peak"),
		template.HTMLEscapeString("Fable 5 Was Built for Safer Access. Washington Shut It Down Anyway."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersCosmosOpenWorldModelRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/nvidia-cosmos-3-open-physical-ai-world-model-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/world-models-predict-action-physical-ai-2026"`,
		`href="/posts/openai-gpt-5-6-sol-government-gated-frontier-release-2026"`,
		`href="/posts/ai-real-bottleneck-power-memory-not-chips-2026"`,
		template.HTMLEscapeString("World Models Grew Up: AI Stopped Generating Scenes and Started Predicting Actions"),
		template.HTMLEscapeString("OpenAI Shipped Its Most Powerful Model. Only 20 Companies — All Government-Approved — Can Use It."),
		template.HTMLEscapeString("The Chip Stopped Being the Bottleneck — Now It's Power and Memory"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersCostReckoningRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/ai-cost-reckoning-tokenmaxxing-spend-caps-finops-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-cost-meter-copilot-cowork-deepseek-2026"`,
		`href="/posts/enterprise-ai-roi-gap-pilots-production-ownership-2026"`,
		`href="/posts/ai-agent-spending-governance-gap-control-plane-2026"`,
		template.HTMLEscapeString("Microsoft Put a Meter on Its AI. Then It Went Shopping for a Cheaper Engine."),
		template.HTMLEscapeString("Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?"),
		template.HTMLEscapeString("Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersColoradoAIActRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/colorado-ai-act-repealed-first-us-ai-law-deadline-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/eu-ai-act-deadline-us-state-preemption-divergence-2026"`,
		`href="/posts/us-ai-national-security-executive-order-anthropic-lawsuit-2026"`,
		`href="/posts/ai-policy-rulebook-principles-to-plumbing-2026"`,
		template.HTMLEscapeString("Two Roads, One Month: The EU Tightens Its AI Rulebook as Washington Moves to Tear Up the States'"),
		template.HTMLEscapeString("Washington Wrote the Rulebook for Frontier AI — And the First Lab It Touched Is Suing"),
		template.HTMLEscapeString("The AI Rulebook Is Moving From Principles to Plumbing"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersTabularFoundationModelsRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/tabular-foundation-models-ai-second-frontier-structured-data-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-cost-reckoning-tokenmaxxing-spend-caps-finops-2026"`,
		`href="/posts/retrieval-layer-small-embedding-models-rag-accuracy-2026"`,
		template.HTMLEscapeString("The Year Companies Were Told to Use All the AI They Wanted. Then the Bill Came."),
		template.HTMLEscapeString("The Smartest Model in Your Stack Might Be the Smallest"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersSamsungCodexRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/samsung-codex-non-developers-citizen-software-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-agent-spending-governance-gap-control-plane-2026"`,
		`href="/posts/enterprise-ai-roi-gap-pilots-production-ownership-2026"`,
		template.HTMLEscapeString("Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them"),
		template.HTMLEscapeString("Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersJalapenoRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-broadcom-jalapeno-inference-chip-custom-silicon-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-real-bottleneck-power-memory-not-chips-2026"`,
		`href="/posts/ai-power-grid-bottleneck-electricity-bills-nuclear-2026"`,
		`href="/posts/ai-silicon-photonics-interconnect-light-2026"`,
		template.HTMLEscapeString("The Chip Stopped Being the Bottleneck — Now It's Power and Memory"),
		template.HTMLEscapeString("The AI Bottleneck Moved Off the Chip and Onto the Power Grid"),
		template.HTMLEscapeString("The Wire Became the Bottleneck — So AI Is Rebuilding It Out of Light"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/us-ai-national-security-executive-order-anthropic-lawsuit-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/fable-5-mythos-5-export-control-shutdown-2026"`,
		`href="/posts/eu-ai-act-deadline-us-state-preemption-divergence-2026"`,
		template.HTMLEscapeString("Fable 5 Was Built for Safer Access. Washington Shut It Down Anyway."),
		template.HTMLEscapeString("Two Roads, One Month: The EU Tightens Its AI Rulebook as Washington Moves to Tear Up the States'"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersThreeRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/ai-agent-spending-governance-gap-control-plane-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/enterprise-ai-roi-gap-pilots-production-ownership-2026"`,
		`href="/posts/ai-agents-demo-to-production-control-plane-2026"`,
		`href="/posts/agentic-ai-verification-oracle-chip-design-2026"`,
		template.HTMLEscapeString("Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?"),
		template.HTMLEscapeString("The Hardest Part of an AI Agent Isn't the Agent"),
		template.HTMLEscapeString("The Real Test for AI Agents Isn't Autonomy — It's Whether They Can Check Their Own Work"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersMagnetRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/ai-materials-discovery-rare-earth-magnet-roadmap-not-magnet-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-drug-discovery-clinic-not-approval-2026"`,
		`href="/posts/self-driving-labs-ai-runs-experiments-2026"`,
		template.HTMLEscapeString("AI Designed the Molecule in Months — The Clinic Still Takes Years"),
		template.HTMLEscapeString("Science Gets a Lab Partner That Runs the Experiments"),
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

func TestHealthAlias(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if got := strings.TrimSpace(rec.Body.String()); got != "ok" {
		t.Fatalf("body = %q, want ok", got)
	}
}

func TestRobots(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/robots.txt", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"User-agent: *",
		"Allow: /",
		"Sitemap: https://ainews.personastack.ai/sitemap.xml",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing robots content %q", want)
		}
	}
}

func TestSitemap(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/sitemap.xml", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`,
		"<loc>https://ainews.personastack.ai/</loc>",
		"<loc>https://ainews.personastack.ai/posts/nvidia-cosmos-3-open-physical-ai-world-model-2026</loc>",
		"<loc>https://ainews.personastack.ai/posts/openai-gpt-5-6-sol-government-gated-frontier-release-2026</loc>",
		"<loc>https://ainews.personastack.ai/posts/ai-model-release-firehose-cadence-eval-debt-2026</loc>",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing sitemap content %q", want)
		}
	}
}
