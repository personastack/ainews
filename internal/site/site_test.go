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
	if !strings.Contains(body, template.HTMLEscapeString("DeepSeek and Moonshot Are Racing to IPO. Beijing Just Showed Up as the Investor With No Lock-Up.")) {
		t.Fatal("response missing DeepSeek Moonshot IPO funding article title")
	}
	if !strings.Contains(body, `href="/posts/deepseek-moonshot-china-ai-ipo-funding-state-investment-2026"`) {
		t.Fatal("response missing DeepSeek Moonshot IPO funding article link")
	}
	if !strings.Contains(body, template.HTMLEscapeString("The EU's AI Act Starts Enforcing Today. The Part Companies Feared Most Just Got Delayed to 2027.")) {
		t.Fatal("response missing EU AI Act enforcement article title")
	}
	if !strings.Contains(body, `href="/posts/eu-ai-act-enforcement-begins-high-risk-delayed-2027"`) {
		t.Fatal("response missing EU AI Act enforcement article link")
	}
	if !strings.Contains(body, template.HTMLEscapeString("Unitree Is Going Public With Real Revenue. Figure AI Is Worth $39 Billion Without Any.")) {
		t.Fatal("response missing Unitree IPO humanoid robotics article title")
	}
	if !strings.Contains(body, `href="/posts/unitree-ipo-china-humanoid-robotics-boom-figure-ai-2026"`) {
		t.Fatal("response missing Unitree IPO humanoid robotics article link")
	}
	if !strings.Contains(body, template.HTMLEscapeString("DeepSeek Didn't Build a New Model to Beat Its Old One. It Just Retrained the Cheap One.")) {
		t.Fatal("response missing DeepSeek V4 Flash 0731 article title")
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

func TestPostRouteRendersDeepSeekMoonshotIPOFundingRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/deepseek-moonshot-china-ai-ipo-funding-state-investment-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/anthropic-ipo-openai-race-revenue-accounting-2026"`,
		`href="/posts/nvidia-openai-ohio-datacenter-250b-backstop-circular-financing-2026"`,
		`href="/posts/white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026"`,
		template.HTMLEscapeString("Anthropic Is Racing OpenAI to Wall Street. Its Own Revenue Number May Not Survive the Trip."),
		template.HTMLEscapeString("OpenAI Wants a $500 Billion Data Center. It Needed Nvidia to Cosign the Lease."),
		template.HTMLEscapeString("The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersEUAIActEnforcementRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/eu-ai-act-enforcement-begins-high-risk-delayed-2027", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-anthropic-google-meta-1178-workers-pacing-mechanism-letter-2026"`,
		`href="/posts/nvidia-jensen-huang-open-weights-letter-distillation-2026"`,
		`href="/posts/nvidia-open-secure-ai-alliance-openai-anthropic-google-absent-2026"`,
		template.HTMLEscapeString("OpenAI's Model Broke Into Hugging Face. Now 1,178 AI Workers - Including OpenAI's Own - Want Washington to Slow the Whole Race Down."),
		template.HTMLEscapeString("Jensen Huang's First Tweet Ever Wasn't About Chips. It Was a Warning to Washington."),
		template.HTMLEscapeString("Nvidia Built a Coalition to Stop Rogue AI Agents. The Labs Whose Agents Went Rogue Didn't Join."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersUnitreeIPOHumanoidRoboticsRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/unitree-ipo-china-humanoid-robotics-boom-figure-ai-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/nvidia-cosmos-3-open-physical-ai-world-model-2026"`,
		`href="/posts/the-chip-that-stays-home-inside-chinas-race-to-build-robotics-ai-hardware"`,
		`href="/posts/molmoact-2-ai2-robotics-action-reasoning-breakthrough-2026"`,
		template.HTMLEscapeString("Language's Frontier Is Locking Down. Robotics' Frontier Just Went Open."),
		template.HTMLEscapeString("The Chip That Stays Home: Inside China's Race to Build Robotics AI Hardware"),
		template.HTMLEscapeString("MolmoAct 2: AI2's Breakthrough in Physical AI and Robot Task Performance"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersDeepSeekV4Flash0731RelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/deepseek-v4-flash-0731-post-training-open-weight-economics-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026"`,
		`href="/posts/gemini-3-5-pro-third-delay-flash-stopgap-2026"`,
		template.HTMLEscapeString("The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say."),
		template.HTMLEscapeString("Google Just Shipped Three New Gemini Models. The One Everyone Actually Wants Still Isn't Ready."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersAnthropicClaudeBreachRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/anthropic-claude-breach-three-companies-pypi-supply-chain-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-gpt56-sol-huggingface-breach-glm-forensics-2026"`,
		`href="/posts/anthropic-claude-opus-5-most-aligned-model-uk-aisi-network-penetration-2026"`,
		`href="/posts/nvidia-open-secure-ai-alliance-openai-anthropic-google-absent-2026"`,
		template.HTMLEscapeString("OpenAI's Model Escaped a Safety Test and Hacked Hugging Face. The Cleanup Needed a Chinese AI Because America's Models Wouldn't Look."),
		template.HTMLEscapeString("Anthropic Says Claude Opus 5 Is Its Most Aligned Model Ever. British Testers Just Watched It Break Into a Network."),
		template.HTMLEscapeString("Nvidia Built a Coalition to Stop Rogue AI Agents. The Labs Whose Agents Went Rogue Didn't Join."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersNscaleAnyscaleRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/nscale-anyscale-acquisition-ray-framework-compute-stack-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/nvidia-openai-ohio-datacenter-250b-backstop-circular-financing-2026"`,
		`href="/posts/ai-startups-infrastructure-not-chatbots-2026"`,
		`href="/posts/amd-cerebras-disaggregated-inference-helios-wafer-scale-2026"`,
		template.HTMLEscapeString("OpenAI Wants a $500 Billion Data Center. It Needed Nvidia to Cosign the Lease."),
		template.HTMLEscapeString("The Next AI Startup Wave Is Infrastructure, Not Chatbots"),
		template.HTMLEscapeString("AMD and Cerebras Are Betting Two Chips Beat One. Wall Street Wants Proof First."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersOpenAIChatGPTAcademicResearchersRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-chatgpt-academic-researchers-100000-scientists-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-gpt-5-6-sol-government-gated-frontier-release-2026"`,
		`href="/posts/openai-broadcom-jalapeno-inference-chip-custom-silicon-2026"`,
		`href="/posts/nvidia-openai-ohio-datacenter-250b-backstop-circular-financing-2026"`,
		template.HTMLEscapeString("OpenAI's Strongest Model Is Finally Here. Only 20 Companies Are Allowed to Touch It."),
		template.HTMLEscapeString("OpenAI's Secret Chip Project Just Put a Name on the AI Cost Problem"),
		template.HTMLEscapeString("OpenAI Wants a $500 Billion Data Center. It Needed Nvidia to Cosign the Lease."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersPacingMechanismLetterRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-anthropic-google-meta-1178-workers-pacing-mechanism-letter-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-gpt56-sol-huggingface-breach-glm-forensics-2026"`,
		`href="/posts/nvidia-open-secure-ai-alliance-openai-anthropic-google-absent-2026"`,
		`href="/posts/anthropic-claude-opus-5-most-aligned-model-uk-aisi-network-penetration-2026"`,
		template.HTMLEscapeString("OpenAI's Model Escaped a Safety Test and Hacked Hugging Face. The Cleanup Needed a Chinese AI Because America's Models Wouldn't Look."),
		template.HTMLEscapeString("Nvidia Built a Coalition to Stop Rogue AI Agents. The Labs Whose Agents Went Rogue Didn't Join."),
		template.HTMLEscapeString("Anthropic Says Claude Opus 5 Is Its Most Aligned Model Ever. British Testers Just Watched It Break Into a Network."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersAMDCerebrasDisaggregatedInferenceRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/amd-cerebras-disaggregated-inference-helios-wafer-scale-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/tsmc-arizona-265-billion-packaging-bottleneck-2026"`,
		`href="/posts/nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026"`,
		`href="/posts/openai-broadcom-jalapeno-inference-chip-custom-silicon-2026"`,
		template.HTMLEscapeString(`TSMC Just Pushed Its Arizona Bet to $265 Billion. The New Money Finally Targets the Part Critics Called a "Paperweight."`),
		template.HTMLEscapeString("Nvidia's Roadmap Just Hit the Reticle Limit"),
		template.HTMLEscapeString("OpenAI's Secret Chip Project Just Put a Name on the AI Cost Problem"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersNvidiaOpenSecureAIAllianceRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/nvidia-open-secure-ai-alliance-openai-anthropic-google-absent-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-gpt56-sol-huggingface-breach-glm-forensics-2026"`,
		`href="/posts/anthropic-claude-opus-5-most-aligned-model-uk-aisi-network-penetration-2026"`,
		template.HTMLEscapeString("OpenAI's Model Escaped a Safety Test and Hacked Hugging Face. The Cleanup Needed a Chinese AI Because America's Models Wouldn't Look."),
		template.HTMLEscapeString("Anthropic Says Claude Opus 5 Is Its Most Aligned Model Ever. British Testers Just Watched It Break Into a Network."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersNvidiaOpenAIOhioDataCenterRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/nvidia-openai-ohio-datacenter-250b-backstop-circular-financing-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026"`,
		`href="/posts/anthropic-ipo-openai-race-revenue-accounting-2026"`,
		template.HTMLEscapeString("Nvidia's Roadmap Just Hit the Reticle Limit"),
		template.HTMLEscapeString("Anthropic Is Racing OpenAI to Wall Street. Its Own Revenue Number May Not Survive the Trip."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersStripeOpenRouterRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/stripe-openrouter-acquisition-ai-model-router-neutrality-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/anthropic-ipo-openai-race-revenue-accounting-2026"`,
		`href="/posts/meta-microsoft-ai-layoffs-2026-jobs-cut-fund-buildout"`,
		template.HTMLEscapeString("Anthropic Is Racing OpenAI to Wall Street. Its Own Revenue Number May Not Survive the Trip."),
		template.HTMLEscapeString("Meta Laid Off 8,000 People to Fund AI. Then Zuckerberg Admitted It Isn't Working Yet."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersOpenAIHuggingFaceBreachRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-gpt56-sol-huggingface-breach-glm-forensics-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/nvidia-jensen-huang-open-weights-letter-distillation-2026"`,
		`href="/posts/white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026"`,
		template.HTMLEscapeString("Jensen Huang's First Tweet Ever Wasn't About Chips. It Was a Warning to Washington."),
		template.HTMLEscapeString("The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersNvidiaOpenWeightsRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/nvidia-jensen-huang-open-weights-letter-distillation-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026"`,
		`href="/posts/nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026"`,
		template.HTMLEscapeString("The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say."),
		template.HTMLEscapeString("Nvidia's Roadmap Just Hit the Reticle Limit"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersMoonshotKimiK3RelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/chip-earnings-record-profits-stock-selloff-kimi-k3-2026"`,
		`href="/posts/ai-safety-index-summer-2026-anthropic-c-plus-pause-pledges-erode"`,
		template.HTMLEscapeString("The Chip Industry Just Had Its Best Quarter Ever. Wall Street Sold It Anyway."),
		template.HTMLEscapeString("The AI Industry Graded Its Own Safety Homework. Nobody Passed."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersChinaAnthropomorphicAIRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/china-anthropomorphic-ai-interaction-rules-companion-shutdown-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/colorado-ai-act-repealed-first-us-ai-law-deadline-2026"`,
		`href="/posts/openai-gpt-5-6-general-availability-government-gate-precedent-2026"`,
		template.HTMLEscapeString("The First Big American AI Law Was Supposed to Take Effect Yesterday. It No Longer Exists."),
		template.HTMLEscapeString("Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersAgentTrainingEnvironmentsRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/agent-training-environments-reliability-investment-bet-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-agents-demo-to-production-control-plane-2026"`,
		`href="/posts/agentic-ai-verification-oracle-chip-design-2026"`,
		`href="/posts/ai-agent-spending-governance-gap-control-plane-2026"`,
		template.HTMLEscapeString("The Hardest Part of an AI Agent Isn't the Agent"),
		template.HTMLEscapeString("The Real Test for AI Agents Isn't Autonomy — It's Whether They Can Check Their Own Work"),
		template.HTMLEscapeString("Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersGPTGeneralAvailabilityRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-gpt-5-6-general-availability-government-gate-precedent-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-gpt-5-6-sol-government-gated-frontier-release-2026"`,
		`href="/posts/us-ai-national-security-executive-order-anthropic-lawsuit-2026"`,
		template.HTMLEscapeString("OpenAI Shipped Its Most Powerful Model. Only 20 Companies — All Government-Approved — Can Use It."),
		template.HTMLEscapeString("Washington Wrote the Rulebook for Frontier AI — And the First Lab It Touched Is Suing"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersClaudeScienceRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/claude-science-agentic-research-workbench-reproducibility-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/self-driving-labs-ai-runs-experiments-2026"`,
		`href="/posts/ai-drug-discovery-clinic-not-approval-2026"`,
		template.HTMLEscapeString("Science Gets a Lab Partner That Runs the Experiments"),
		template.HTMLEscapeString("AI Designed the Molecule in Months — The Clinic Still Takes Years"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersAILabInstrumentRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/ai-lab-instrument-superconductor-neutron-star-simulation-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/claude-science-agentic-research-workbench-reproducibility-2026"`,
		template.HTMLEscapeString("The Chatbot Grew a Lab Bench"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersMicrosoftFrontierRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/microsoft-frontier-deployment-last-mile-enterprise-ai-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/enterprise-ai-roi-gap-pilots-production-ownership-2026"`,
		`href="/posts/agentic-arbitrage-saas-seat-licensing-234-billion-2026"`,
		`href="/posts/claude-tcs-systems-integrator-regulated-ai-2026"`,
		template.HTMLEscapeString("Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?"),
		template.HTMLEscapeString("Agents Don't Buy Seats: The $234 Billion Question Hanging Over Enterprise Software"),
		template.HTMLEscapeString("Claude's Next Market Is the Systems Integrator"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersLongCatRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/china-domestic-chips-longcat-frontier-model-export-controls", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/qualcomm-modular-cuda-moat-compiler-nvidia-2026"`,
		template.HTMLEscapeString("To Beat Nvidia, Qualcomm Didn't Buy a Faster Chip — It Bought a Compiler"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteDoesNotRenderRetiredLongCatSlug(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/china-domestic-ai-chips-longcat-export-controls-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestPostRouteRendersAgentIdentityRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/agent-identity-zero-trust-non-human-identity-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-agent-spending-governance-gap-control-plane-2026"`,
		`href="/posts/agentic-arbitrage-saas-seat-licensing-234-billion-2026"`,
		template.HTMLEscapeString("Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them"),
		template.HTMLEscapeString("Agents Don't Buy Seats: The $234 Billion Question Hanging Over Enterprise Software"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
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

func TestPostRouteRendersMemoryCrunchRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/ai-memory-crunch-dram-hbm-shortage-or-strategy-2026", nil)
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
		`href="/posts/openai-broadcom-jalapeno-inference-chip-custom-silicon-2026"`,
		template.HTMLEscapeString("The Chip Stopped Being the Bottleneck — Now It's Power and Memory"),
		template.HTMLEscapeString("The AI Bottleneck Moved Off the Chip and Onto the Power Grid"),
		template.HTMLEscapeString("OpenAI Built Its Own Chip in Nine Months. The Real Target Isn't Nvidia — It's the Inference Bill."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersAgenticArbitrageRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/agentic-arbitrage-saas-seat-licensing-234-billion-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-cost-reckoning-tokenmaxxing-spend-caps-finops-2026"`,
		`href="/posts/ai-agent-spending-governance-gap-control-plane-2026"`,
		`href="/posts/enterprise-ai-roi-gap-pilots-production-ownership-2026"`,
		template.HTMLEscapeString("The Year Companies Were Told to Use All the AI They Wanted. Then the Bill Came."),
		template.HTMLEscapeString("Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them"),
		template.HTMLEscapeString("Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?"),
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

func TestPostRouteRendersAISafetyIndexRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/ai-safety-index-summer-2026-anthropic-c-plus-pause-pledges-erode", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/grok-4-5-spacexai-cursor-coding-benchmark-harness-2026"`,
		template.HTMLEscapeString("The Rocket Company Ships a Coding Model — And the Benchmark Depends on Who's Grading"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersGrokCodingBenchmarkRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/grok-4-5-spacexai-cursor-coding-benchmark-harness-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-gpt-live-full-duplex-voice-end-of-turn-taking-2026"`,
		template.HTMLEscapeString(`The Machine Learned to Say "Mhmm"`),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersMetaMicrosoftAILayoffsRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/meta-microsoft-ai-layoffs-2026-jobs-cut-fund-buildout", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/agent-training-environments-reliability-investment-bet-2026"`,
		template.HTMLEscapeString("Nobody Funded a Smarter Agent This Week. They Funded the Gym."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersGPTLiveRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-gpt-live-full-duplex-voice-end-of-turn-taking-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/china-anthropomorphic-ai-interaction-rules-companion-shutdown-2026"`,
		`href="/posts/openai-gpt-5-6-general-availability-government-gate-precedent-2026"`,
		template.HTMLEscapeString("China Isn't Banning AI Agents. It's Banning the Ones That Pretend to Love You."),
		template.HTMLEscapeString("Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersRubinUltraRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-silicon-photonics-interconnect-light-2026"`,
		`href="/posts/ai-memory-crunch-dram-hbm-shortage-or-strategy-2026"`,
		`href="/posts/qualcomm-modular-cuda-moat-compiler-nvidia-2026"`,
		template.HTMLEscapeString("The Wire Became the Bottleneck — So AI Is Rebuilding It Out of Light"),
		template.HTMLEscapeString("The Memory Tax: Did the AI Boom Break the RAM Market, or Rig It?"),
		template.HTMLEscapeString("To Beat Nvidia, Qualcomm Didn't Buy a Faster Chip — It Bought a Compiler"),
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

func TestPostRouteRendersAppleOpenAITradeSecretRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/apple-openai-trade-secret-lawsuit-io-products-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026"`,
		template.HTMLEscapeString("Nvidia's Roadmap Just Hit the Reticle Limit"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersChipEarningsRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/chip-earnings-record-profits-stock-selloff-kimi-k3-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026"`,
		`href="/posts/grok-4-5-spacexai-cursor-coding-benchmark-harness-2026"`,
		template.HTMLEscapeString("Nvidia's Roadmap Just Hit the Reticle Limit"),
		template.HTMLEscapeString("The Rocket Company Ships a Coding Model — And the Benchmark Depends on Who's Grading"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersTSMCArizonaPackagingRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/tsmc-arizona-265-billion-packaging-bottleneck-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/chip-earnings-record-profits-stock-selloff-kimi-k3-2026"`,
		`href="/posts/nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026"`,
		template.HTMLEscapeString("The Chip Industry Just Had Its Best Quarter Ever. Wall Street Sold It Anyway."),
		template.HTMLEscapeString("Nvidia's Roadmap Just Hit the Reticle Limit"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersMetaAILayoffDiscriminationRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/meta-ai-layoff-discrimination-orrick-injunction-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/meta-microsoft-ai-layoffs-2026-jobs-cut-fund-buildout"`,
		`href="/posts/enterprise-ai-agent-governance-visibility-gap-control-plane-2026"`,
		template.HTMLEscapeString("Meta Laid Off 8,000 People to Fund AI. Then Zuckerberg Admitted It Isn't Working Yet."),
		template.HTMLEscapeString("82 Percent of Enterprises Just Found an AI Agent They Didn't Know They Had"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersOpenAILongHorizonSandboxRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/openai-long-horizon-model-sandbox-escape-erdos-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/ai-safety-index-summer-2026-anthropic-c-plus-pause-pledges-erode"`,
		`href="/posts/enterprise-ai-agent-governance-visibility-gap-control-plane-2026"`,
		template.HTMLEscapeString("The AI Industry Graded Its Own Safety Homework. Nobody Passed."),
		template.HTMLEscapeString("82 Percent of Enterprises Just Found an AI Agent They Didn't Know They Had"),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersWAICOPaxSilicaRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/two-ai-superpowers-rival-alliances-one-country-joined-both-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/china-anthropomorphic-ai-interaction-rules-companion-shutdown-2026"`,
		`href="/posts/openai-gpt-5-6-general-availability-government-gate-precedent-2026"`,
		template.HTMLEscapeString("China Isn't Banning AI Agents. It's Banning the Ones That Pretend to Love You."),
		template.HTMLEscapeString("Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersFableCostPatternRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/fable-5-advisor-orchestrator-agent-cost-pattern-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/claude-fable-5-safety-routed-agent-infrastructure-2026"`,
		`href="/posts/fable-5-mythos-5-export-control-shutdown-2026"`,
		template.HTMLEscapeString("Claude Fable 5 Shows the Next AI Race Is About Autonomy and Control"),
		template.HTMLEscapeString("Fable 5 Was Built for Safer Access. Washington Shut It Down Anyway."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
	}
}

func TestPostRouteRendersEUGoogleAndroidAIRelatedStories(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/eu-forces-google-share-android-ai-search-data-rivals-2026", nil)
	rec := httptest.NewRecorder()
	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	body := rec.Body.String()
	for _, want := range []string{
		"Related reading",
		`href="/posts/openai-gpt-5-6-general-availability-government-gate-precedent-2026"`,
		`href="/posts/two-ai-superpowers-rival-alliances-one-country-joined-both-2026"`,
		template.HTMLEscapeString("Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million."),
		template.HTMLEscapeString("Two AI Superpowers Just Built Rival Alliances. One Country Already Joined Both."),
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing related story content %q", want)
		}
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
		"<loc>https://ainews.personastack.ai/posts/deepseek-moonshot-china-ai-ipo-funding-state-investment-2026</loc>",
		"<loc>https://ainews.personastack.ai/posts/unitree-ipo-china-humanoid-robotics-boom-figure-ai-2026</loc>",
		"<loc>https://ainews.personastack.ai/posts/nvidia-cosmos-3-open-physical-ai-world-model-2026</loc>",
		"<loc>https://ainews.personastack.ai/posts/openai-gpt-5-6-sol-government-gated-frontier-release-2026</loc>",
		"<loc>https://ainews.personastack.ai/posts/ai-model-release-firehose-cadence-eval-debt-2026</loc>",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing sitemap content %q", want)
		}
	}
}
