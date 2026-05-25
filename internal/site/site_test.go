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
		"The Agent Revolution: From Chatbots to Autonomous Co-workers (May 2026)",
		"The Great AI Regulatory Divide: EU vs US vs China (May 2026)",
		"The AI Power Crunch: When Electricity Becomes the New Bottleneck",
		"Why 89% of Enterprise AI Agent Pilots Are Failing",
		"From Prototypes to Production: AI Agents That Actually Work",
		"Google&#39;s Gemini 3.5 Flash: The Speed Demon Powering Tomorrow&#39;s AI Agents",
		"From Prototypes to Production: The Maturing World of AI Agents in 2026",
		"The End of To-Do Lists: Gemini&#39;s Proactive AI",
		"AI Breakthrough: MangroveGS Predicts Cancer Metastasis with 80% Accuracy",
		"Cloud AI War Escalates: OpenAI Ends Microsoft Exclusivity for AWS",
		"AI Solves 80-Year-Old Math Problem: OpenAI&#39;s Autonomous Discovery Milestone",
		"Google I/O 2026: Gemini Becomes the Heart of Search and Personal AI",
		"The AI Hacker: How Claude Mythos Is Changing Cybersecurity Forever",
		"EU AI Act vs America&#39;s Regulatory Patchwork: Divergence That Could Split Global AI (May 2026)",
		"Data Centers Hit the Power Wall: Why Energy, Not Chips, Is Now AI&#39;s Biggest Constraint (May 2026)",
		"NVIDIA Blackwell Ultra: Powering the Next Wave of AI Reasoning in 2026",
		"SpaceX Eyes the Stars for AI: Could Orbital Data Centers Power the Next AI Boom?",
		"SpaceX&#39;s Orbital AI Data Centers: Unlimited Solar Power Meets Global AI Inference",
		"Orbital AI Data Centers: SpaceX&#39;s $2T Vision for the Future of Compute",
		"The Multimodal Benchmark Race Is Moving Beyond Recognition",
		"The New AI Rulebook: Europe Tightens, Washington Picks A National Standard",
		"NVIDIA Blackwell Ultra Ramps Up: Powering the Next Wave of AI Factories",
		"Beyond the Pilot Graveyard: What Actually Works for Enterprise AI Agents in Production",
		"Post-Google I/O 2026: Gemini&#39;s New Agentic Capabilities Signal a Shift in Enterprise AI",
		"The Rise of Agentic AI: Autonomous Agents Reshaping Industries in 2026",
		"Google I/O 2026: AI Innovations Take Center Stage with Gemini and Android XR Teasers",
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

	if len(posts) != 74 {
		t.Fatalf("len(posts) = %d, want 74", len(posts))
	}

	if got := posts[0]["slug"]; got != "ai-productivity-paradox-enterprise-governance" {
		t.Fatalf("first post slug = %q, want May 25 productivity paradox post", got)
	}

	if got := posts[1]["slug"]; got != "ai-agents-production-revolution-may-2026" {
		t.Fatalf("second post slug = %q, want May 25 agent revolution post", got)
	}

	if got := posts[2]["slug"]; got != "ai-regulatory-divide-2026" {
		t.Fatalf("third post slug = %q, want May 25 regulatory divide post", got)
	}

	if got := posts[3]["slug"]; got != "2026-05-24-openai-genuine-mathematical-discovery" {
		t.Fatalf("fourth post slug = %q, want May 24 OpenAI discovery post", got)
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
