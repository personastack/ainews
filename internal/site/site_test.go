package site

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
	for _, want := range []string{
		"The Rise of Efficient AI Models: Balancing Performance and Sustainability in 2026",
		"The Rise of Specialized AI Agents in Enterprise Workflows: 2026 Trends",
		"OpenAI&#39;s Strategic Expansion: Acquiring Voice Tech and Launching Finance Tools",
		"The Next Frontier: Orbital Data Centers and the Environmental Cost of AI",
		"Towards Seamless Multimodal Agents: Conquering the Tool-Switching Tax",
		"FDA&#39;s First AI Warning Letter: Why &#39;The AI Didn&#39;t Tell Us&#39; Is No Defense",
		"Cloudflare&#39;s Global LLM Inference Infrastructure: Agents Week 2026 Deep Dive",
		"xAI&#39;s Grok 4 Sprint: Three Releases in Six Weeks Chasing GPT-5.5",
		"The 86% Enterprise AI Agent Failure Rate: Governance Crisis Explained",
		"Maturing Reasoning Models: Adaptive Thinking Takes Center Stage",
		"Gemini 3.1 Ultra: Why Google&#39;s Native Multimodal Architecture Is The Real Story",
		"From $30 to $0.40 Per Million Tokens",
		"Qwen Surpasses Llama",
		"Mythos National Security Standoff",
		"CAISI Framework",
		"The May 2026 Model Rush",
		"Agentic AI Trends",
		"Frontier AI as Cyber Weapons",
		"May 2026 AI Model Rush",
		"Google I/O 2026: AI Innovations, Gemini Updates, and Android XR on the Horizon",
		"Google I/O 2026 Preview",
		"Google&#39;s Android Show 2026",
		"The Frontier Firm Is Here",
		"The Government That Fears Its Own Weapon",
		"Samsung&#39;s Trillion-Dollar Moment",
		"The Lab Without Scientists",
		"The Chip That Stays Home",
		"Apple&#39;s Multi-AI Gambit",
		"The Sprint Is Real",
		"The CAISI Reversal",
		"Three AIs, Three Laws",
		"From Text-to-Video to Intent-to-Video",
		"Wake-Up Call Software Engineers Needed",
		"Broken Promise Worth $134 Billion",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing %q", want)
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

	if len(posts) != 38 {
		t.Fatalf("len(posts) = %d, want 38", len(posts))
	}

	if got := posts[0]["slug"]; got != "google-io-2026-ai-innovations-gemini-updates-and-android-xr-on-the-horizon" {
		t.Fatalf("first post slug = %q, want newest Google I/O post", got)
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
