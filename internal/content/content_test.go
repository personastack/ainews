package content

import (
	"testing"
	"time"
)

func containsSlug(posts []Post, slug string) bool {
	for _, post := range posts {
		if post.Slug == slug {
			return true
		}
	}

	return false
}

func TestPostsReturnsPublishedPosts(t *testing.T) {
	got := Posts()
	if len(got) == 0 {
		t.Fatal("Posts() returned no posts")
	}

	for _, post := range got {
		if post.Title == "" {
			t.Fatal("Posts() returned a post with an empty title")
		}
		if post.Slug == "" {
			t.Fatal("Posts() returned a post with an empty slug")
		}
		if post.Date == "" {
			t.Fatalf("post %q has an empty date", post.Slug)
		}
		if len(post.Sections) == 0 {
			t.Fatalf("post %q has no sections", post.Slug)
		}
	}
}

func TestPostsDoNotExceedCurrentUTCDate(t *testing.T) {
	now := time.Now().UTC()

	for _, post := range Posts() {
		postDate, err := time.Parse("January 2, 2006", post.Date)
		if err != nil {
			t.Fatalf("parse date for post %q: %v", post.Slug, err)
		}

		if postDate.After(now) {
			t.Fatalf("post %q is dated %s, which is after current UTC time %s", post.Slug, post.Date, now.Format(time.RFC3339))
		}
	}
}

func TestPublishedPostsAppliesFutureDateGate(t *testing.T) {
	onPublicationJune22 := time.Date(2026, time.June, 22, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublicationJune22), "enterprise-ai-roi-gap-pilots-production-ownership-2026") {
		t.Fatal("publishedPosts() did not include enterprise AI ROI gap article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune22), "retrieval-layer-small-embedding-models-rag-accuracy-2026") {
		t.Fatal("publishedPosts() did not include retrieval layer small embedding models article on publication date")
	}

	onPublicationJune21 := time.Date(2026, time.June, 21, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune21), "enterprise-ai-roi-gap-pilots-production-ownership-2026") {
		t.Fatal("publishedPosts() included enterprise AI ROI gap article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune21), "retrieval-layer-small-embedding-models-rag-accuracy-2026") {
		t.Fatal("publishedPosts() included retrieval layer small embedding models article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune21), "agentic-ai-verification-oracle-chip-design-2026") {
		t.Fatal("publishedPosts() did not include agentic AI verification oracle article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune21), "eu-ai-act-deadline-us-state-preemption-divergence-2026") {
		t.Fatal("publishedPosts() did not include EU AI Act and US state preemption article on publication date")
	}

	onPublicationJune20 := time.Date(2026, time.June, 20, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune20), "agentic-ai-verification-oracle-chip-design-2026") {
		t.Fatal("publishedPosts() included agentic AI verification oracle article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune20), "eu-ai-act-deadline-us-state-preemption-divergence-2026") {
		t.Fatal("publishedPosts() included EU AI Act and US state preemption article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune20), "ai-silicon-photonics-interconnect-light-2026") {
		t.Fatal("publishedPosts() did not include AI silicon photonics interconnect article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune20), "ai-drug-discovery-clinic-not-approval-2026") {
		t.Fatal("publishedPosts() did not include AI drug discovery clinic article on publication date")
	}

	onPublicationJune19 := time.Date(2026, time.June, 19, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune19), "ai-silicon-photonics-interconnect-light-2026") {
		t.Fatal("publishedPosts() included AI silicon photonics interconnect article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune19), "ai-drug-discovery-clinic-not-approval-2026") {
		t.Fatal("publishedPosts() included AI drug discovery clinic article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune19), "glm-5-2-open-weights-frontier-coding-cost-2026") {
		t.Fatal("publishedPosts() did not include GLM-5.2 open-weights coding article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune19), "ai-cost-meter-copilot-cowork-deepseek-2026") {
		t.Fatal("publishedPosts() did not include Microsoft AI cost meter article on publication date")
	}

	onPublicationJune18 := time.Date(2026, time.June, 18, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune18), "glm-5-2-open-weights-frontier-coding-cost-2026") {
		t.Fatal("publishedPosts() included GLM-5.2 open-weights coding article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune18), "ai-cost-meter-copilot-cowork-deepseek-2026") {
		t.Fatal("publishedPosts() included Microsoft AI cost meter article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune18), "ai-builds-itself-recursive-self-improvement-coordinated-pause-2026") {
		t.Fatal("publishedPosts() did not include AI builds itself article on publication date")
	}

	onPublicationJune17 := time.Date(2026, time.June, 17, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune17), "ai-builds-itself-recursive-self-improvement-coordinated-pause-2026") {
		t.Fatal("publishedPosts() included AI builds itself article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune17), "reasoning-models-test-time-compute-think-smarter-2026") {
		t.Fatal("publishedPosts() did not include reasoning models test-time compute article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune17), "ai-real-bottleneck-power-memory-not-chips-2026") {
		t.Fatal("publishedPosts() did not include AI power and memory bottleneck article on publication date")
	}

	onPublicationJune16 := time.Date(2026, time.June, 16, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune16), "reasoning-models-test-time-compute-think-smarter-2026") {
		t.Fatal("publishedPosts() included reasoning models test-time compute article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune16), "ai-real-bottleneck-power-memory-not-chips-2026") {
		t.Fatal("publishedPosts() included AI power and memory bottleneck article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune16), "world-models-predict-action-physical-ai-2026") {
		t.Fatal("publishedPosts() did not include world models physical AI article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune16), "ai-agents-demo-to-production-control-plane-2026") {
		t.Fatal("publishedPosts() did not include AI agents demo-to-production control plane article on publication date")
	}

	onPublicationJune15 := time.Date(2026, time.June, 15, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune15), "world-models-predict-action-physical-ai-2026") {
		t.Fatal("publishedPosts() included world models physical AI article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune15), "ai-agents-demo-to-production-control-plane-2026") {
		t.Fatal("publishedPosts() included AI agents demo-to-production control plane article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune15), "ny-synthetic-performer-ai-ad-law-state-patchwork-2026") {
		t.Fatal("publishedPosts() did not include New York synthetic performer ad law article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune15), "uk-ai-hardware-plan-first-customer-chips-2026") {
		t.Fatal("publishedPosts() did not include UK AI hardware first customer article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune15), "self-driving-labs-ai-runs-experiments-2026") {
		t.Fatal("publishedPosts() did not include self-driving labs AI experiments article on publication date")
	}

	onPublicationJune14 := time.Date(2026, time.June, 14, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune14), "ny-synthetic-performer-ai-ad-law-state-patchwork-2026") {
		t.Fatal("publishedPosts() included New York synthetic performer ad law article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune14), "uk-ai-hardware-plan-first-customer-chips-2026") {
		t.Fatal("publishedPosts() included UK AI hardware first customer article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune14), "self-driving-labs-ai-runs-experiments-2026") {
		t.Fatal("publishedPosts() included self-driving labs AI experiments article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune14), "openai-ona-codex-enterprise-runtime-2026") {
		t.Fatal("publishedPosts() did not include OpenAI Ona Codex enterprise runtime article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune14), "medical-ai-specialist-moat-llm-benchmark-2026") {
		t.Fatal("publishedPosts() did not include medical AI specialist moat article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune14), "claude-tcs-systems-integrator-regulated-ai-2026") {
		t.Fatal("publishedPosts() did not include Claude TCS systems integrator article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune14), "chatgpt-visa-agentic-commerce-payments-2026") {
		t.Fatal("publishedPosts() did not include ChatGPT Visa agentic commerce payments article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune14), "healthcare-ai-operating-office-cms-2026") {
		t.Fatal("publishedPosts() did not include healthcare AI operating office article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune14), "ai-safety-left-lab-courts-g7-2026") {
		t.Fatal("publishedPosts() did not include AI safety left the lab article on publication date")
	}

	onPublicationJune13 := time.Date(2026, time.June, 13, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune13), "openai-ona-codex-enterprise-runtime-2026") {
		t.Fatal("publishedPosts() included OpenAI Ona Codex enterprise runtime article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune13), "medical-ai-specialist-moat-llm-benchmark-2026") {
		t.Fatal("publishedPosts() included medical AI specialist moat article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune13), "claude-tcs-systems-integrator-regulated-ai-2026") {
		t.Fatal("publishedPosts() included Claude TCS systems integrator article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune13), "chatgpt-visa-agentic-commerce-payments-2026") {
		t.Fatal("publishedPosts() included ChatGPT Visa agentic commerce payments article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune13), "healthcare-ai-operating-office-cms-2026") {
		t.Fatal("publishedPosts() included healthcare AI operating office article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune13), "ai-safety-left-lab-courts-g7-2026") {
		t.Fatal("publishedPosts() included AI safety left the lab article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune13), "fable-5-mythos-5-export-control-shutdown-2026") {
		t.Fatal("publishedPosts() did not include Fable 5 Mythos 5 export control article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune13), "agent-registry-security-perimeter-agentic-ai-2026") {
		t.Fatal("publishedPosts() did not include agent registry security perimeter article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune13), "botsitting-hidden-cost-ai-productivity-boom-2026") {
		t.Fatal("publishedPosts() did not include botsitting productivity article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune13), "opentext-ireland-sovereign-agentic-ai-2026") {
		t.Fatal("publishedPosts() did not include OpenText Ireland sovereign agentic AI article on publication date")
	}

	onPublicationJune12 := time.Date(2026, time.June, 12, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune12), "fable-5-mythos-5-export-control-shutdown-2026") {
		t.Fatal("publishedPosts() included Fable 5 Mythos 5 export control article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune12), "agent-registry-security-perimeter-agentic-ai-2026") {
		t.Fatal("publishedPosts() included agent registry security perimeter article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune12), "botsitting-hidden-cost-ai-productivity-boom-2026") {
		t.Fatal("publishedPosts() included botsitting productivity article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune12), "opentext-ireland-sovereign-agentic-ai-2026") {
		t.Fatal("publishedPosts() included OpenText Ireland sovereign agentic AI article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune12), "meta-ai-glasses-blind-veterans-accessibility-wearables-2026") {
		t.Fatal("publishedPosts() did not include Meta AI glasses accessibility article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune12), "agents-need-managers-enterprise-ai-infrastructure-2026") {
		t.Fatal("publishedPosts() did not include enterprise agent infrastructure article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune12), "tito-ai-fast-forwards-molecular-simulation-2026") {
		t.Fatal("publishedPosts() did not include TITO molecular simulation article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune12), "ai-agents-data-cyera-trust-layer-market-2026") {
		t.Fatal("publishedPosts() did not include Cyera trust layer market article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune12), "ai-review-machine-anthropic-veto-power-2026") {
		t.Fatal("publishedPosts() did not include Anthropic AI review machine article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune12), "ai-startups-infrastructure-not-chatbots-2026") {
		t.Fatal("publishedPosts() did not include AI startup infrastructure article on publication date")
	}

	onPublicationJune11 := time.Date(2026, time.June, 11, 0, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(onPublicationJune11), "meta-ai-glasses-blind-veterans-accessibility-wearables-2026") {
		t.Fatal("publishedPosts() included Meta AI glasses accessibility article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune11), "agents-need-managers-enterprise-ai-infrastructure-2026") {
		t.Fatal("publishedPosts() included enterprise agent infrastructure article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune11), "tito-ai-fast-forwards-molecular-simulation-2026") {
		t.Fatal("publishedPosts() included TITO molecular simulation article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune11), "ai-agents-data-cyera-trust-layer-market-2026") {
		t.Fatal("publishedPosts() included Cyera trust layer market article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune11), "ai-review-machine-anthropic-veto-power-2026") {
		t.Fatal("publishedPosts() included Anthropic AI review machine article before publication date")
	}
	if containsSlug(publishedPosts(onPublicationJune11), "ai-startups-infrastructure-not-chatbots-2026") {
		t.Fatal("publishedPosts() included AI startup infrastructure article before publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune11), "hong-kong-sfc-ai-cyber-risk-financial-regulation-2026") {
		t.Fatal("publishedPosts() did not include Hong Kong SFC AI cyber risk article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune11), "atos-agent-365-governance-fleet-management-2026") {
		t.Fatal("publishedPosts() did not include Atos agent governance article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune11), "ai-agents-fraud-aml-verafin-2026") {
		t.Fatal("publishedPosts() did not include Verafin fraud and AML agents article on publication date")
	}

	onPublicationJune10 := time.Date(2026, time.June, 10, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublicationJune10), "cohere-north-mini-code-local-coding-agents-2026") {
		t.Fatal("publishedPosts() did not include Cohere North Mini Code article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune10), "chatgpt-reasoning-effort-product-ux-2026") {
		t.Fatal("publishedPosts() did not include ChatGPT reasoning effort picker article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune10), "claude-fable-5-safety-routed-agent-infrastructure-2026") {
		t.Fatal("publishedPosts() did not include Claude Fable 5 article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune10), "ai-agents-devops-infrastructure-rebuild-2026") {
		t.Fatal("publishedPosts() did not include AI agents DevOps infrastructure article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune10), "ai-compliance-calendar-global-rulebook-2026") {
		t.Fatal("publishedPosts() did not include AI compliance calendar article on publication date")
	}

	onPublicationJune9 := time.Date(2026, time.June, 9, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublicationJune9), "ai-benchmarks-enter-agent-era-2026") {
		t.Fatal("publishedPosts() did not include AI benchmark agent era article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune9), "ai-science-workflow-era-2026") {
		t.Fatal("publishedPosts() did not include AI science workflow article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune9), "uk-ai-hardware-plan-compute-industrial-policy-2026") {
		t.Fatal("publishedPosts() did not include UK AI hardware plan article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune9), "ai-hardware-race-moves-from-chips-to-systems-2026") {
		t.Fatal("publishedPosts() did not include AI hardware systems article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune9), "sap-erp-control-plane-autonomous-enterprise-2026") {
		t.Fatal("publishedPosts() did not include SAP ERP control plane article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune9), "quantum-adapters-llm-compression-hardware-path-2026") {
		t.Fatal("publishedPosts() did not include quantum adapters article on publication date")
	}

	onPublicationJune7 := time.Date(2026, time.June, 7, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublicationJune7), "ibm-google-cloud-gemini-enterprise-ai-modernization-2026") {
		t.Fatal("publishedPosts() did not include IBM and Google Cloud enterprise AI article on publication date")
	}

	onPublicationJune8 := time.Date(2026, time.June, 8, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublicationJune8), "ai-policy-rulebook-principles-to-plumbing-2026") {
		t.Fatal("publishedPosts() did not include AI policy rulebook article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune8), "intel-rackscale-agentic-inference-cpu-comeback-2026") {
		t.Fatal("publishedPosts() did not include Intel rackscale article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune8), "microsoft-build-2026-agent-memory-layer") {
		t.Fatal("publishedPosts() did not include Microsoft Build 2026 memory layer article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune8), "google-search-agent-workbench-gemini-2026") {
		t.Fatal("publishedPosts() did not include Google Search agent workbench article on publication date")
	}

	onPublicationJune6 := time.Date(2026, time.June, 6, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublicationJune6), "intel-crescent-island-agentic-inference-2026") {
		t.Fatal("publishedPosts() did not include Intel Crescent Island article on publication date")
	}

	onPublicationJune5 := time.Date(2026, time.June, 5, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublicationJune5), "federal-ai-oversight-frontier-labs-2026") {
		t.Fatal("publishedPosts() did not include federal AI oversight article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune5), "enterprise-ai-roi-gap-95-17-2026") {
		t.Fatal("publishedPosts() did not include enterprise AI ROI article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune5), "copilot-300k-seats-india-it-2026") {
		t.Fatal("publishedPosts() did not include Copilot 300K seats article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune5), "ai-power-grid-578-billion-gap-2026") {
		t.Fatal("publishedPosts() did not include AI power grid article on publication date")
	}
	if !containsSlug(publishedPosts(onPublicationJune5), "agentic-coding-orchestration-battleground-2026") {
		t.Fatal("publishedPosts() did not include agentic coding orchestration article on publication date")
	}

	beforePublication := time.Date(2026, time.May, 25, 12, 0, 0, 0, time.UTC)
	if containsSlug(publishedPosts(beforePublication), "alibaba-cloud-agentic-ai-offensive-qwen3-7-max") {
		t.Fatal("publishedPosts() included Alibaba Cloud offensive article before publication date")
	}
	if containsSlug(publishedPosts(beforePublication), "consumer-ai-vs-hype-reality") {
		t.Fatal("publishedPosts() included consumer AI hype article before publication date")
	}
	if containsSlug(publishedPosts(beforePublication), "ai-early-disease-detection-breakthroughs") {
		t.Fatal("publishedPosts() included early disease detection post before publication date")
	}
	if containsSlug(publishedPosts(beforePublication), "the-agentic-ai-revolution-governance-gap") {
		t.Fatal("publishedPosts() included governance gap post before publication date")
	}

	onPublication := time.Date(2026, time.May, 27, 0, 0, 0, 0, time.UTC)
	if !containsSlug(publishedPosts(onPublication), "alibaba-cloud-agentic-ai-offensive-qwen3-7-max") {
		t.Fatal("publishedPosts() did not include Alibaba Cloud offensive article on publication date")
	}
	if !containsSlug(publishedPosts(onPublication), "consumer-ai-vs-hype-reality") {
		t.Fatal("publishedPosts() did not include consumer AI hype article on publication date")
	}
	if !containsSlug(publishedPosts(onPublication), "ai-early-disease-detection-breakthroughs") {
		t.Fatal("publishedPosts() did not include early disease detection post on publication date")
	}
	if !containsSlug(publishedPosts(onPublication), "the-agentic-ai-revolution-governance-gap") {
		t.Fatal("publishedPosts() did not include governance gap post on publication date")
	}
}

func TestFindBySlug(t *testing.T) {
	enterpriseAgentROIPost, ok := FindBySlug("enterprise-ai-roi-gap-pilots-production-ownership-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find enterprise AI ROI gap article")
	}
	if enterpriseAgentROIPost.Title != "Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?" {
		t.Fatalf("FindBySlug() returned %q for enterprise AI ROI gap article", enterpriseAgentROIPost.Title)
	}
	if len(enterpriseAgentROIPost.Related) != 2 {
		t.Fatalf("enterprise AI ROI gap article related count = %d, want 2", len(enterpriseAgentROIPost.Related))
	}

	retrievalPost, ok := FindBySlug("retrieval-layer-small-embedding-models-rag-accuracy-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find retrieval layer small embedding models article")
	}
	if retrievalPost.Title != "The Smartest Model in Your Stack Might Be the Smallest" {
		t.Fatalf("FindBySlug() returned %q for retrieval layer small embedding models article", retrievalPost.Title)
	}
	if len(retrievalPost.Related) != 2 {
		t.Fatalf("retrieval layer small embedding models article related count = %d, want 2", len(retrievalPost.Related))
	}

	agenticVerificationPost, ok := FindBySlug("agentic-ai-verification-oracle-chip-design-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find agentic AI verification oracle article")
	}
	if agenticVerificationPost.Title != "The Real Test for AI Agents Isn't Autonomy — It's Whether They Can Check Their Own Work" {
		t.Fatalf("FindBySlug() returned %q for agentic AI verification oracle article", agenticVerificationPost.Title)
	}
	if len(agenticVerificationPost.Related) != 2 {
		t.Fatalf("agentic AI verification oracle article related count = %d, want 2", len(agenticVerificationPost.Related))
	}

	euAIActPost, ok := FindBySlug("eu-ai-act-deadline-us-state-preemption-divergence-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find EU AI Act and US state preemption article")
	}
	if euAIActPost.Title != "Two Roads, One Month: The EU Tightens Its AI Rulebook as Washington Moves to Tear Up the States'" {
		t.Fatalf("FindBySlug() returned %q for EU AI Act and US state preemption article", euAIActPost.Title)
	}
	if len(euAIActPost.Related) != 2 {
		t.Fatalf("EU AI Act and US state preemption article related count = %d, want 2", len(euAIActPost.Related))
	}

	siliconPhotonicsPost, ok := FindBySlug("ai-silicon-photonics-interconnect-light-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI silicon photonics interconnect article")
	}
	if siliconPhotonicsPost.Title != "The Wire Became the Bottleneck — So AI Is Rebuilding It Out of Light" {
		t.Fatalf("FindBySlug() returned %q for AI silicon photonics interconnect article", siliconPhotonicsPost.Title)
	}

	glmPost, ok := FindBySlug("glm-5-2-open-weights-frontier-coding-cost-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find GLM-5.2 open-weights coding article")
	}
	if glmPost.Title != "An Open-Weights Model Just Caught the Frontier on Coding — at One-Sixth the Price" {
		t.Fatalf("FindBySlug() returned %q for GLM-5.2 open-weights coding article", glmPost.Title)
	}

	powerMemoryPost, ok := FindBySlug("ai-real-bottleneck-power-memory-not-chips-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI power and memory bottleneck article")
	}
	if powerMemoryPost.Title != "The Chip Stopped Being the Bottleneck — Now It's Power and Memory" {
		t.Fatalf("FindBySlug() returned %q for AI power and memory bottleneck article", powerMemoryPost.Title)
	}

	worldModelsPost, ok := FindBySlug("world-models-predict-action-physical-ai-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find world models physical AI article")
	}
	if worldModelsPost.Title != "World Models Grew Up: AI Stopped Generating Scenes and Started Predicting Actions" {
		t.Fatalf("FindBySlug() returned %q for world models physical AI article", worldModelsPost.Title)
	}

	ukHardwareFirstCustomerPost, ok := FindBySlug("uk-ai-hardware-plan-first-customer-chips-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find UK AI hardware first customer article")
	}
	if ukHardwareFirstCustomerPost.Title != "Britain's £1.1 Billion Bet: Become the AI Chip Industry's First Customer" {
		t.Fatalf("FindBySlug() returned %q for UK AI hardware first customer article", ukHardwareFirstCustomerPost.Title)
	}

	medicalAIPost, ok := FindBySlug("medical-ai-specialist-moat-llm-benchmark-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find medical AI specialist moat article")
	}
	if medicalAIPost.Title != "Medical AI's Specialist Moat Just Cracked" {
		t.Fatalf("FindBySlug() returned %q for medical AI specialist moat article", medicalAIPost.Title)
	}

	openAIOnaPost, ok := FindBySlug("openai-ona-codex-enterprise-runtime-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenAI Ona Codex enterprise runtime post")
	}
	if openAIOnaPost.Title != "The Laptop Becomes a Handoff: OpenAI's Ona Deal Turns Codex Into an Enterprise Runtime" {
		t.Fatalf("FindBySlug() returned %q for OpenAI Ona Codex enterprise runtime post", openAIOnaPost.Title)
	}

	claudeTCSPost, ok := FindBySlug("claude-tcs-systems-integrator-regulated-ai-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Claude TCS systems integrator post")
	}
	if claudeTCSPost.Title != "Claude's Next Market Is the Systems Integrator" {
		t.Fatalf("FindBySlug() returned %q for Claude TCS systems integrator post", claudeTCSPost.Title)
	}

	chatGPTVisaPost, ok := FindBySlug("chatgpt-visa-agentic-commerce-payments-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find ChatGPT Visa agentic commerce payments post")
	}
	if chatGPTVisaPost.Title != "ChatGPT Can Shop. Visa Wants to Decide How AI Agents Pay" {
		t.Fatalf("FindBySlug() returned %q for ChatGPT Visa agentic commerce payments post", chatGPTVisaPost.Title)
	}

	healthcareAIPost, ok := FindBySlug("healthcare-ai-operating-office-cms-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find healthcare AI operating office post")
	}
	if healthcareAIPost.Title != "Healthcare AI Just Got an Operating Office" {
		t.Fatalf("FindBySlug() returned %q for healthcare AI operating office post", healthcareAIPost.Title)
	}

	aiSafetyPost, ok := FindBySlug("ai-safety-left-lab-courts-g7-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI safety left the lab post")
	}
	if aiSafetyPost.Title != "AI Safety Has Left the Lab" {
		t.Fatalf("FindBySlug() returned %q for AI safety left the lab post", aiSafetyPost.Title)
	}

	fableMythosPost, ok := FindBySlug("fable-5-mythos-5-export-control-shutdown-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Fable 5 Mythos 5 export control post")
	}
	if fableMythosPost.Title != "Fable 5 Was Built for Safer Access. Washington Shut It Down Anyway." {
		t.Fatalf("FindBySlug() returned %q for Fable 5 Mythos 5 export control post", fableMythosPost.Title)
	}

	agentRegistryPost, ok := FindBySlug("agent-registry-security-perimeter-agentic-ai-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find agent registry security perimeter post")
	}
	if agentRegistryPost.Title != "The Agent Registry Is Becoming the New Security Perimeter" {
		t.Fatalf("FindBySlug() returned %q for agent registry security perimeter post", agentRegistryPost.Title)
	}

	botsittingPost, ok := FindBySlug("botsitting-hidden-cost-ai-productivity-boom-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find botsitting productivity post")
	}
	if botsittingPost.Title != "Botsitting Is the Hidden Cost of the AI Productivity Boom" {
		t.Fatalf("FindBySlug() returned %q for botsitting productivity post", botsittingPost.Title)
	}

	openTextPost, ok := FindBySlug("opentext-ireland-sovereign-agentic-ai-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenText Ireland sovereign agentic AI post")
	}
	if openTextPost.Title != "OpenText's Ireland Bet Shows Enterprise Agents Need Borders" {
		t.Fatalf("FindBySlug() returned %q for OpenText Ireland sovereign agentic AI post", openTextPost.Title)
	}

	metaAIGlassesPost, ok := FindBySlug("meta-ai-glasses-blind-veterans-accessibility-wearables-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Meta AI glasses accessibility post")
	}
	if metaAIGlassesPost.Title != "Meta's AI Glasses Find the Use Case Wearables Needed" {
		t.Fatalf("FindBySlug() returned %q for Meta AI glasses accessibility post", metaAIGlassesPost.Title)
	}

	agentInfrastructurePost, ok := FindBySlug("agents-need-managers-enterprise-ai-infrastructure-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find enterprise agent infrastructure post")
	}
	if agentInfrastructurePost.Title != "Agents Need Managers Now: Enterprise AI Enters Its IAM and FinOps Era" {
		t.Fatalf("FindBySlug() returned %q for enterprise agent infrastructure post", agentInfrastructurePost.Title)
	}

	titoPost, ok := FindBySlug("tito-ai-fast-forwards-molecular-simulation-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find TITO molecular simulation post")
	}
	if titoPost.Title != "AI Is Learning to Fast-Forward Molecules" {
		t.Fatalf("FindBySlug() returned %q for TITO molecular simulation post", titoPost.Title)
	}

	cyeraTrustLayerPost, ok := FindBySlug("ai-agents-data-cyera-trust-layer-market-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Cyera trust layer market post")
	}
	if cyeraTrustLayerPost.Title != "AI Agents Need Data. Cyera's $600 Million Round Shows the Trust Layer Is Becoming a Market" {
		t.Fatalf("FindBySlug() returned %q for Cyera trust layer market post", cyeraTrustLayerPost.Title)
	}

	anthropicReviewPost, ok := FindBySlug("ai-review-machine-anthropic-veto-power-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Anthropic AI review machine post")
	}
	if anthropicReviewPost.Title != "Washington Is Building an AI Review Machine. Anthropic Wants a Veto Button" {
		t.Fatalf("FindBySlug() returned %q for Anthropic AI review machine post", anthropicReviewPost.Title)
	}

	aiStartupInfrastructurePost, ok := FindBySlug("ai-startups-infrastructure-not-chatbots-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI startup infrastructure post")
	}
	if aiStartupInfrastructurePost.Title != "The Next AI Startup Wave Is Infrastructure, Not Chatbots" {
		t.Fatalf("FindBySlug() returned %q for AI startup infrastructure post", aiStartupInfrastructurePost.Title)
	}

	hongKongSFCPost, ok := FindBySlug("hong-kong-sfc-ai-cyber-risk-financial-regulation-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Hong Kong SFC AI cyber risk post")
	}
	if hongKongSFCPost.Title != "Hong Kong Shows How AI Cyber Risk Becomes Financial Regulation" {
		t.Fatalf("FindBySlug() returned %q for Hong Kong SFC AI cyber risk post", hongKongSFCPost.Title)
	}

	atosPost, ok := FindBySlug("atos-agent-365-governance-fleet-management-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Atos agent governance post")
	}
	if atosPost.Title != "Atos Has 19,000 AI Agents. Now Comes the Hard Part." {
		t.Fatalf("FindBySlug() returned %q for Atos agent governance post", atosPost.Title)
	}

	verafinPost, ok := FindBySlug("ai-agents-fraud-aml-verafin-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Verafin fraud and AML agents post")
	}
	if verafinPost.Title != "The First Serious AI Agents May Work in Fraud, Not Chat" {
		t.Fatalf("FindBySlug() returned %q for Verafin fraud and AML agents post", verafinPost.Title)
	}

	cohereNorthMiniCodePost, ok := FindBySlug("cohere-north-mini-code-local-coding-agents-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Cohere North Mini Code post")
	}
	if cohereNorthMiniCodePost.Title != "Cohere's North Mini Code Shows the Next Coding Agent Race Is About Control" {
		t.Fatalf("FindBySlug() returned %q for Cohere North Mini Code post", cohereNorthMiniCodePost.Title)
	}

	chatGPTPickerPost, ok := FindBySlug("chatgpt-reasoning-effort-product-ux-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find ChatGPT reasoning effort picker post")
	}
	if chatGPTPickerPost.Title != "Reasoning Becomes a Button: ChatGPT's New Picker Turns Compute Into UX" {
		t.Fatalf("FindBySlug() returned %q for ChatGPT reasoning effort picker post", chatGPTPickerPost.Title)
	}

	claudeFablePost, ok := FindBySlug("claude-fable-5-safety-routed-agent-infrastructure-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Claude Fable 5 post")
	}
	if claudeFablePost.Title != "Claude Fable 5 Shows the Next AI Race Is About Autonomy and Control" {
		t.Fatalf("FindBySlug() returned %q for Claude Fable 5 post", claudeFablePost.Title)
	}

	aiAgentsDevOpsPost, ok := FindBySlug("ai-agents-devops-infrastructure-rebuild-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI agents DevOps infrastructure post")
	}
	if aiAgentsDevOpsPost.Title != "When AI Agents Join the Build Pipeline, DevOps Has to Rebuild" {
		t.Fatalf("FindBySlug() returned %q for AI agents DevOps infrastructure post", aiAgentsDevOpsPost.Title)
	}

	policyCalendarPost, ok := FindBySlug("ai-compliance-calendar-global-rulebook-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI compliance calendar post")
	}
	if policyCalendarPost.Title != "AI Compliance Has a Calendar Now: The Global Rulebook Moves From Debate to Deadlines" {
		t.Fatalf("FindBySlug() returned %q for AI compliance calendar post", policyCalendarPost.Title)
	}

	ukAIHardwarePlanPost, ok := FindBySlug("uk-ai-hardware-plan-compute-industrial-policy-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find UK AI hardware plan post")
	}
	if ukAIHardwarePlanPost.Title != "Britain's AI Hardware Bet: The Chip Plan That Turns Compute Into Industrial Policy" {
		t.Fatalf("FindBySlug() returned %q for UK AI hardware plan post", ukAIHardwarePlanPost.Title)
	}

	aiBenchmarksPost, ok := FindBySlug("ai-benchmarks-enter-agent-era-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI benchmark agent era post")
	}
	if aiBenchmarksPost.Title != "The Leaderboard Is No Longer the Product: AI Benchmarks Enter Their Agent Era" {
		t.Fatalf("FindBySlug() returned %q for AI benchmark agent era post", aiBenchmarksPost.Title)
	}

	aiScienceWorkflowPost, ok := FindBySlug("ai-science-workflow-era-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI science workflow post")
	}
	if aiScienceWorkflowPost.Title != "AI Science Enters the Workflow Era" {
		t.Fatalf("FindBySlug() returned %q for AI science workflow post", aiScienceWorkflowPost.Title)
	}

	aiHardwarePost, ok := FindBySlug("ai-hardware-race-moves-from-chips-to-systems-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI hardware systems post")
	}
	if aiHardwarePost.Title != "The AI Hardware Race Has Moved From Chips to Systems" {
		t.Fatalf("FindBySlug() returned %q for AI hardware systems post", aiHardwarePost.Title)
	}

	sapEnterprisePost, ok := FindBySlug("sap-erp-control-plane-autonomous-enterprise-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find SAP ERP control plane post")
	}
	if sapEnterprisePost.Title != "SAP Wants ERP to Become the Control Plane for Enterprise AI" {
		t.Fatalf("FindBySlug() returned %q for SAP ERP control plane post", sapEnterprisePost.Title)
	}

	quantumAdaptersPost, ok := FindBySlug("quantum-adapters-llm-compression-hardware-path-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find quantum adapters post")
	}
	if quantumAdaptersPost.Title != "Quantum Adapters Offer a Small but Real Hardware Path for LLM Efficiency" {
		t.Fatalf("FindBySlug() returned %q for quantum adapters post", quantumAdaptersPost.Title)
	}

	ibmGoogleCloudPost, ok := FindBySlug("ibm-google-cloud-gemini-enterprise-ai-modernization-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find IBM and Google Cloud enterprise AI post")
	}
	if ibmGoogleCloudPost.Title != "The AI Pilot Is Over: IBM and Google Cloud Turn Gemini Into a Modernization Program" {
		t.Fatalf("FindBySlug() returned %q for IBM and Google Cloud enterprise AI post", ibmGoogleCloudPost.Title)
	}

	chainOfCommandPost, ok := FindBySlug("frontier-ai-chain-of-command-nspm-11-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find frontier AI chain of command post")
	}
	if chainOfCommandPost.Title != "Frontier AI Enters the Chain of Command" {
		t.Fatalf("FindBySlug() returned %q for frontier AI chain of command post", chainOfCommandPost.Title)
	}

	memoryPlatformPost, ok := FindBySlug("chatgpt-memory-ai-platform-layer-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find ChatGPT memory platform post")
	}
	if memoryPlatformPost.Title != "ChatGPT Memory Becomes the Next AI Platform Layer" {
		t.Fatalf("FindBySlug() returned %q for ChatGPT memory platform post", memoryPlatformPost.Title)
	}

	intelCrescentIslandPost, ok := FindBySlug("intel-crescent-island-agentic-inference-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Intel Crescent Island post")
	}
	if intelCrescentIslandPost.Title != "Intel Skips the Training War: Why Crescent Island's 480GB Bet Is All About Agentic AI" {
		t.Fatalf("FindBySlug() returned %q for Intel Crescent Island post", intelCrescentIslandPost.Title)
	}

	intelRackscalePost, ok := FindBySlug("intel-rackscale-agentic-inference-cpu-comeback-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Intel rackscale post")
	}
	if intelRackscalePost.Title != "The CPU Returns to the AI Story: Intel's Rackscale Bet on Agentic Inference" {
		t.Fatalf("FindBySlug() returned %q for Intel rackscale post", intelRackscalePost.Title)
	}

	aiPolicyPost, ok := FindBySlug("ai-policy-rulebook-principles-to-plumbing-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI policy rulebook post")
	}
	if aiPolicyPost.Title != "The AI Rulebook Is Moving From Principles to Plumbing" {
		t.Fatalf("FindBySlug() returned %q for AI policy rulebook post", aiPolicyPost.Title)
	}

	federalOversightPost, ok := FindBySlug("federal-ai-oversight-frontier-labs-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find federal AI oversight post")
	}
	if federalOversightPost.Title != "The Labs Want a Federal Referee — Inside the Week AI's Frontier Asked Washington to Step In" {
		t.Fatalf("FindBySlug() returned %q for federal AI oversight post", federalOversightPost.Title)
	}

	enterpriseROIPost, ok := FindBySlug("enterprise-ai-roi-gap-95-17-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find enterprise AI ROI post")
	}
	if enterpriseROIPost.Title != "95% Say It Worked, 17% Say It Wowed: Enterprise AI's ROI Reckoning Arrives" {
		t.Fatalf("FindBySlug() returned %q for enterprise AI ROI post", enterpriseROIPost.Title)
	}

	copilotSeatsPost, ok := FindBySlug("copilot-300k-seats-india-it-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Copilot 300K seats post")
	}
	if copilotSeatsPost.Title != "300,000 Seats in Six Months — How India's IT Giants Dragged Copilot Out of Pilot Purgatory" {
		t.Fatalf("FindBySlug() returned %q for Copilot 300K seats post", copilotSeatsPost.Title)
	}

	powerGridPost, ok := FindBySlug("ai-power-grid-578-billion-gap-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI power grid post")
	}
	if powerGridPost.Title != "The $578 Billion Gap: AI's Real Bottleneck Isn't Chips, It's the Power Grid" {
		t.Fatalf("FindBySlug() returned %q for AI power grid post", powerGridPost.Title)
	}

	agenticCodingPost, ok := FindBySlug("agentic-coding-orchestration-battleground-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find agentic coding orchestration post")
	}
	if agenticCodingPost.Title != "The Models Tied, So the Fight Moved: How Orchestration Became the Real Agentic-Coding Battleground in 2026" {
		t.Fatalf("FindBySlug() returned %q for agentic coding orchestration post", agenticCodingPost.Title)
	}

	alibabaPost, ok := FindBySlug("alibaba-cloud-agentic-ai-offensive-qwen3-7-max")
	if time.Now().UTC().Before(time.Date(2026, time.May, 27, 0, 0, 0, 0, time.UTC)) {
		if ok {
			t.Fatal("FindBySlug() returned Alibaba Cloud offensive article before publication date")
		}
	} else {
		if !ok {
			t.Fatal("FindBySlug() did not find Alibaba Cloud offensive article")
		}
		if alibabaPost.Title != "Alibaba's AI Offensive: How Qwen3.7-Max and a New Skills Portal Challenge Western Cloud Giants" {
			t.Fatalf("FindBySlug() returned %q for Alibaba Cloud offensive article", alibabaPost.Title)
		}
	}

	consumerAIPost, ok := FindBySlug("consumer-ai-vs-hype-reality")
	if time.Now().UTC().Before(time.Date(2026, time.May, 27, 0, 0, 0, 0, time.UTC)) {
		if ok {
			t.Fatal("FindBySlug() returned consumer AI hype article before publication date")
		}
	} else {
		if !ok {
			t.Fatal("FindBySlug() did not find consumer AI hype article")
		}
		if consumerAIPost.Title != "Consumer AI vs. The Hype Machine: What's Real, What's Bullshit, and What Matters" {
			t.Fatalf("FindBySlug() returned %q for consumer AI hype article", consumerAIPost.Title)
		}
	}

	federalPolicyPost, ok := FindBySlug("federal-ai-executive-order-coordination-over-control-june-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find federal AI order post")
	}
	if federalPolicyPost.Title != "Washington Picks a Lane: Inside the Federal AI Order That Bets on Coordination Over Control" {
		t.Fatalf("FindBySlug() returned %q for federal AI order post", federalPolicyPost.Title)
	}

	googleSearchAgentPost, ok := FindBySlug("google-search-agent-workbench-gemini-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Google Search agent workbench post")
	}
	if googleSearchAgentPost.Title != "When Search Becomes a Workplace: Google's Agentic Gemini Strategy Comes Into Focus" {
		t.Fatalf("FindBySlug() returned %q for Google Search agent workbench post", googleSearchAgentPost.Title)
	}

	microsoftBuildMemoryPost, ok := FindBySlug("microsoft-build-2026-agent-memory-layer")
	if !ok {
		t.Fatal("FindBySlug() did not find Microsoft Build 2026 memory layer post")
	}
	if microsoftBuildMemoryPost.Title != "The Memory Layer Arrives: Microsoft's Build 2026 Shows What Enterprise Agents Need Next" {
		t.Fatalf("FindBySlug() returned %q for Microsoft Build 2026 memory layer post", microsoftBuildMemoryPost.Title)
	}

	rtxSparkPost, ok := FindBySlug("nvidia-rtx-spark-local-ai-supercomputer-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find NVIDIA RTX Spark article")
	}
	if rtxSparkPost.Title != "The Supercomputer Moves In: NVIDIA's RTX Spark and the Quiet Return of Local AI" {
		t.Fatalf("FindBySlug() returned %q for NVIDIA RTX Spark article", rtxSparkPost.Title)
	}

	safetyBreakthroughsPost, ok := FindBySlug("ai-safety-breakthroughs-responsible-development-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI safety breakthroughs post")
	}
	if safetyBreakthroughsPost.Title != "AI Safety Breakthroughs Signal a New Era of Responsible AI Development" {
		t.Fatalf("FindBySlug() returned %q for AI safety breakthroughs post", safetyBreakthroughsPost.Title)
	}

	edgeMovesToDevicePost, ok := FindBySlug("edge-ai-intelligence-moves-to-device-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find edge AI moves to device post")
	}
	if edgeMovesToDevicePost.Title != "Edge AI Revolution: Intelligence Moves to the Device" {
		t.Fatalf("FindBySlug() returned %q for edge AI moves to device post", edgeMovesToDevicePost.Title)
	}

	geminiSearchPost, ok := FindBySlug("google-gemini-3-5-flash-search-revolution-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Gemini search revolution post")
	}
	if geminiSearchPost.Title != "Google's Gemini 3.5 Flash Search Revolution: Conversational AI Takes Over" {
		t.Fatalf("FindBySlug() returned %q for Gemini search revolution post", geminiSearchPost.Title)
	}

	scientificRevolutionPost, ok := FindBySlug("ai-scientific-revolution-materials-climate")
	if !ok {
		t.Fatal("FindBySlug() did not find scientific revolution post")
	}
	if scientificRevolutionPost.Title != "AI Ignites a Scientific Revolution: From Materials to Climate Solutions" {
		t.Fatalf("FindBySlug() returned %q for scientific revolution post", scientificRevolutionPost.Title)
	}

	postBlackwellHardwarePost, ok := FindBySlug("post-blackwell-ai-hardware-competition-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find post-Blackwell hardware competition post")
	}
	if postBlackwellHardwarePost.Title != "Post-Blackwell AI Hardware Competition: Quantum, Photonic and Neuromorphic Bets Accelerate" {
		t.Fatalf("FindBySlug() returned %q for post-Blackwell hardware competition post", postBlackwellHardwarePost.Title)
	}

	productivityParadoxPost, ok := FindBySlug("ai-productivity-paradox-enterprise-governance")
	if !ok {
		t.Fatal("FindBySlug() did not find productivity paradox post")
	}
	if productivityParadoxPost.Title != "The AI Productivity Paradox: Governance Gaps in Enterprise Adoption" {
		t.Fatalf("FindBySlug() returned %q for productivity paradox post", productivityParadoxPost.Title)
	}

	agentRevolutionPost, ok := FindBySlug("ai-agents-production-revolution-may-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find agent revolution post")
	}
	if agentRevolutionPost.Title != "The Agent Revolution: From Chatbots to Autonomous Co-workers (May 2026)" {
		t.Fatalf("FindBySlug() returned %q for agent revolution post", agentRevolutionPost.Title)
	}

	regulatoryDividePost, ok := FindBySlug("ai-regulatory-divide-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find regulatory divide post")
	}
	if regulatoryDividePost.Title != "The Great AI Regulatory Divide: EU vs US vs China (May 2026)" {
		t.Fatalf("FindBySlug() returned %q for regulatory divide post", regulatoryDividePost.Title)
	}

	openAIGenuineDiscoveryPost, ok := FindBySlug("2026-05-24-openai-genuine-mathematical-discovery")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenAI genuine mathematical discovery post")
	}
	if openAIGenuineDiscoveryPost.Title != "OpenAI's AI Makes Genuine Mathematical Discovery" {
		t.Fatalf("FindBySlug() returned %q for OpenAI genuine mathematical discovery post", openAIGenuineDiscoveryPost.Title)
	}

	agenticEnterprisePost, ok := FindBySlug("2026-05-24-agentic-ai-enterprise-deployment")
	if !ok {
		t.Fatal("FindBySlug() did not find agentic AI enterprise deployment post")
	}
	if agenticEnterprisePost.Title != "Agentic AI Moves from Demos to Enterprise Deployment" {
		t.Fatalf("FindBySlug() returned %q for agentic AI enterprise deployment post", agenticEnterprisePost.Title)
	}

	openAIGeometryPost, ok := FindBySlug("2026-05-24-openai-discrete-geometry-breakthrough")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenAI geometry breakthrough post")
	}
	if openAIGeometryPost.Title != "AI's New Math Frontier: OpenAI's Geometry Breakthrough Signals Revolution in Discovery" {
		t.Fatalf("FindBySlug() returned %q for OpenAI geometry breakthrough post", openAIGeometryPost.Title)
	}

	computeArmsRacePost, ok := FindBySlug("2026-05-24-ai-compute-arms-race")
	if !ok {
		t.Fatal("FindBySlug() did not find AI compute arms race post")
	}
	if computeArmsRacePost.Title != "The $500 Billion AI Compute Arms Race: Power, Politics, and Planet at Stake" {
		t.Fatalf("FindBySlug() returned %q for AI compute arms race post", computeArmsRacePost.Title)
	}

	powerCrunchPost, ok := FindBySlug("2026-05-24-ai-power-crunch")
	if !ok {
		t.Fatal("FindBySlug() did not find AI power crunch post")
	}
	if powerCrunchPost.Title != "The AI Power Crunch: When Electricity Becomes the New Bottleneck" {
		t.Fatalf("FindBySlug() returned %q for AI power crunch post", powerCrunchPost.Title)
	}

	agentFailurePost, ok := FindBySlug("2026-05-24-agent-failure-rates")
	if !ok {
		t.Fatal("FindBySlug() did not find enterprise AI agent failure post")
	}
	if agentFailurePost.Title != "Why 89% of Enterprise AI Agent Pilots Are Failing" {
		t.Fatalf("FindBySlug() returned %q for enterprise AI agent failure post", agentFailurePost.Title)
	}

	aiAgentsProductionPost, ok := FindBySlug("from-prototypes-to-production-ai-agents-that-actually-work")
	if !ok {
		t.Fatal("FindBySlug() did not find AI agents production post")
	}
	if aiAgentsProductionPost.Title != "From Prototypes to Production: AI Agents That Actually Work" {
		t.Fatalf("FindBySlug() returned %q for AI agents production post", aiAgentsProductionPost.Title)
	}

	gemini3_5FlashPost, ok := FindBySlug("gemini-3-5-flash-agentic-speed")
	if !ok {
		t.Fatal("FindBySlug() did not find Gemini 3.5 Flash post")
	}
	if gemini3_5FlashPost.Title != "Google's Gemini 3.5 Flash: The Speed Demon Powering Tomorrow's AI Agents" {
		t.Fatalf("FindBySlug() returned %q for Gemini 3.5 Flash post", gemini3_5FlashPost.Title)
	}

	prototypeToProductionPost, ok := FindBySlug("from-prototypes-to-production-the-maturing-world-of-ai-agents-in-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find prototypes to production post")
	}
	if prototypeToProductionPost.Title != "From Prototypes to Production: The Maturing World of AI Agents in 2026" {
		t.Fatalf("FindBySlug() returned %q for prototypes to production post", prototypeToProductionPost.Title)
	}

	geminiSparkProactivePost, ok := FindBySlug("the-end-of-to-do-lists-geminis-proactive-ai")
	if !ok {
		t.Fatal("FindBySlug() did not find Gemini Spark proactive AI post")
	}
	if geminiSparkProactivePost.Title != "The End of To-Do Lists: Gemini's Proactive AI" {
		t.Fatalf("FindBySlug() returned %q for Gemini Spark proactive AI post", geminiSparkProactivePost.Title)
	}

	mangroveGSPost, ok := FindBySlug("ai-cancer-metastasis-prediction-mangrovegs-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find MangroveGS cancer metastasis post")
	}
	if mangroveGSPost.Title != "AI Breakthrough: MangroveGS Predicts Cancer Metastasis with 80% Accuracy" {
		t.Fatalf("FindBySlug() returned %q for MangroveGS cancer metastasis post", mangroveGSPost.Title)
	}

	openAIAWSPost, ok := FindBySlug("openai-aws-partnership-shift-cloud-ai-war-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenAI AWS partnership shift post")
	}
	if openAIAWSPost.Title != "Cloud AI War Escalates: OpenAI Ends Microsoft Exclusivity for AWS" {
		t.Fatalf("FindBySlug() returned %q for OpenAI AWS partnership shift post", openAIAWSPost.Title)
	}

	openAIMathGeometryPost, ok := FindBySlug("openai-ai-solves-80-year-math-problem")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenAI math geometry post")
	}
	if openAIMathGeometryPost.Title != "Mathematical Milestone: OpenAI's AI Autonomously Cracks 80-Year Geometry Problem" {
		t.Fatalf("FindBySlug() returned %q for OpenAI math geometry post", openAIMathGeometryPost.Title)
	}

	openAIMathematicalBreakthroughPost, ok := FindBySlug("openai-mathematical-breakthrough")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenAI mathematical breakthrough post")
	}
	if openAIMathematicalBreakthroughPost.Title != "AI Solves 80-Year-Old Math Problem: OpenAI's Autonomous Discovery Milestone" {
		t.Fatalf("FindBySlug() returned %q for OpenAI mathematical breakthrough post", openAIMathematicalBreakthroughPost.Title)
	}

	googleIOGeminiSearchPost, ok := FindBySlug("google-io-2026-gemini-search")
	if !ok {
		t.Fatal("FindBySlug() did not find Google I/O Gemini search post")
	}
	if googleIOGeminiSearchPost.Title != "Google I/O 2026: Gemini Becomes the Heart of Search and Personal AI" {
		t.Fatalf("FindBySlug() returned %q for Google I/O Gemini search post", googleIOGeminiSearchPost.Title)
	}

	researchFrontiersPost, ok := FindBySlug("ai-research-frontiers-may-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI research frontiers post")
	}
	if researchFrontiersPost.Title != "Beyond the Model Race: May 2026's Most Exciting AI Research Frontiers" {
		t.Fatalf("FindBySlug() returned %q for AI research frontiers post", researchFrontiersPost.Title)
	}

	regulationDivergencePost, ok := FindBySlug("eu-us-ai-regulation-divergence-may-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find EU/US regulatory divergence post")
	}
	if regulationDivergencePost.Title != "EU AI Act vs America's Regulatory Patchwork: Divergence That Could Split Global AI (May 2026)" {
		t.Fatalf("FindBySlug() returned %q for EU/US regulatory divergence post", regulationDivergencePost.Title)
	}

	powerWallPost, ok := FindBySlug("data-center-power-constraints-ai-may-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find data center power constraints post")
	}
	if powerWallPost.Title != "Data Centers Hit the Power Wall: Why Energy, Not Chips, Is Now AI's Biggest Constraint (May 2026)" {
		t.Fatalf("FindBySlug() returned %q for data center power constraints post", powerWallPost.Title)
	}

	blackwellReasoningPost, ok := FindBySlug("nvidia-blackwell-ultra-powering-the-next-wave-of-ai-reasoning-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find NVIDIA Blackwell Ultra reasoning post")
	}
	if blackwellReasoningPost.Title != "NVIDIA Blackwell Ultra: Powering the Next Wave of AI Reasoning in 2026" {
		t.Fatalf("FindBySlug() returned %q for NVIDIA Blackwell Ultra reasoning post", blackwellReasoningPost.Title)
	}

	edgeAIPost, ok := FindBySlug("from-pilots-to-production-the-maturing-edge-ai-revolution-in-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find edge AI production post")
	}
	if edgeAIPost.Title != "From Pilots to Production: The Maturing Edge AI Revolution in 2026" {
		t.Fatalf("FindBySlug() returned %q for edge AI production post", edgeAIPost.Title)
	}

	orbitalVisionPost, ok := FindBySlug("the-final-frontier-for-ai-spacexs-orbital-data-center-vision")
	if !ok {
		t.Fatal("FindBySlug() did not find SpaceX orbital data center vision post")
	}
	if orbitalVisionPost.Title != "The Final Frontier for AI: SpaceX's Orbital Data Center Vision" {
		t.Fatalf("FindBySlug() returned %q for SpaceX orbital data center vision post", orbitalVisionPost.Title)
	}

	edgeMainstreamPost, ok := FindBySlug("from-prototype-to-production-how-edge-ai-went-mainstream-in-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find edge AI mainstream post")
	}
	if edgeMainstreamPost.Title != "From Prototype to Production: How Edge AI Went Mainstream in 2026" {
		t.Fatalf("FindBySlug() returned %q for edge AI mainstream post", edgeMainstreamPost.Title)
	}

	spacexEyesTheStarsPost, ok := FindBySlug("spacex-eyes-the-stars-for-ai-could-orbital-data-centers-power-the-next-ai-boom")
	if !ok {
		t.Fatal("FindBySlug() did not find SpaceX Eyes the Stars post")
	}
	if spacexEyesTheStarsPost.Title != "SpaceX Eyes the Stars for AI: Could Orbital Data Centers Power the Next AI Boom?" {
		t.Fatalf("FindBySlug() returned %q for SpaceX Eyes the Stars post", spacexEyesTheStarsPost.Title)
	}

	orbitalInferencePost, ok := FindBySlug("spacexs-orbital-ai-data-centers-unlimited-solar-power-meets-global-ai-inference")
	if !ok {
		t.Fatal("FindBySlug() did not find SpaceX orbital AI data centers post")
	}
	if orbitalInferencePost.Title != "SpaceX's Orbital AI Data Centers: Unlimited Solar Power Meets Global AI Inference" {
		t.Fatalf("FindBySlug() returned %q for SpaceX orbital AI data centers post", orbitalInferencePost.Title)
	}

	orbitalDataCentersPost, ok := FindBySlug("orbital-ai-data-centers-spacexs-2t-vision-for-the-future-of-compute")
	if !ok {
		t.Fatal("FindBySlug() did not find orbital AI data centers post")
	}
	if orbitalDataCentersPost.Title != "Orbital AI Data Centers: SpaceX's $2T Vision for the Future of Compute" {
		t.Fatalf("FindBySlug() returned %q for orbital AI data centers post", orbitalDataCentersPost.Title)
	}

	multimodalBenchmarkPost, ok := FindBySlug("the-multimodal-benchmark-race-is-moving-beyond-recognition")
	if !ok {
		t.Fatal("FindBySlug() did not find multimodal benchmark post")
	}
	if multimodalBenchmarkPost.Title != "The Multimodal Benchmark Race Is Moving Beyond Recognition" {
		t.Fatalf("FindBySlug() returned %q for multimodal benchmark post", multimodalBenchmarkPost.Title)
	}

	aiRulebookPost, ok := FindBySlug("the-new-ai-rulebook-europe-tightens-washington-picks-a-national-standard")
	if !ok {
		t.Fatal("FindBySlug() did not find AI rulebook post")
	}
	if aiRulebookPost.Title != "The New AI Rulebook: Europe Tightens, Washington Picks A National Standard" {
		t.Fatalf("FindBySlug() returned %q for AI rulebook post", aiRulebookPost.Title)
	}

	blackwellUltraPost, ok := FindBySlug("nvidia-blackwell-ultra-ramps-up-powering-the-next-wave-of-ai-factories")
	if !ok {
		t.Fatal("FindBySlug() did not find NVIDIA Blackwell Ultra post")
	}
	if blackwellUltraPost.Title != "NVIDIA Blackwell Ultra Ramps Up: Powering the Next Wave of AI Factories" {
		t.Fatalf("FindBySlug() returned %q for NVIDIA Blackwell Ultra post", blackwellUltraPost.Title)
	}

	googleIOMay21Post, ok := FindBySlug("post-google-io-2026-geminis-new-agentic-capabilities-signal-a-shift-in-enterprise-ai")
	if !ok {
		t.Fatal("FindBySlug() did not find Google I/O May 21 post")
	}
	if googleIOMay21Post.Title != "Post-Google I/O 2026: Gemini's New Agentic Capabilities Signal a Shift in Enterprise AI" {
		t.Fatalf("FindBySlug() returned %q for Google I/O May 21 post", googleIOMay21Post.Title)
	}

	enterpriseAgentsProductionPost, ok := FindBySlug("beyond-the-pilot-graveyard-what-actually-works-for-enterprise-ai-agents-in-production")
	if !ok {
		t.Fatal("FindBySlug() did not find enterprise agents production post")
	}
	if enterpriseAgentsProductionPost.Title != "Beyond the Pilot Graveyard: What Actually Works for Enterprise AI Agents in Production" {
		t.Fatalf("FindBySlug() returned %q for enterprise agents production post", enterpriseAgentsProductionPost.Title)
	}

	efficientPost, ok := FindBySlug("efficient-ai-models-balancing-performance-and-sustainability-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find efficient AI models post")
	}
	if efficientPost.Title != "The Rise of Efficient AI Models: Balancing Performance and Sustainability in 2026" {
		t.Fatalf("FindBySlug() returned %q for efficient AI models post", efficientPost.Title)
	}

	toolSwitchingPost, ok := FindBySlug("seamless-multimodal-agents-tool-switching-tax-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find seamless multimodal agents post")
	}
	if toolSwitchingPost.Title != "Towards Seamless Multimodal Agents: Conquering the Tool-Switching Tax" {
		t.Fatalf("FindBySlug() returned %q for seamless multimodal agents post", toolSwitchingPost.Title)
	}

	fdaPost, ok := FindBySlug("fda-first-ai-warning-letter-ai-compliance-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find FDA AI warning letter post")
	}
	if fdaPost.Title != "FDA's First AI Warning Letter: Why 'The AI Didn't Tell Us' Is No Defense" {
		t.Fatalf("FindBySlug() returned %q for FDA AI warning letter post", fdaPost.Title)
	}

	cloudflarePost, ok := FindBySlug("cloudflares-global-llm-inference-infrastructure-agents-week-2026-deep-dive")
	if !ok {
		t.Fatal("FindBySlug() did not find Cloudflare inference infrastructure post")
	}
	if cloudflarePost.Title != "Cloudflare's Global LLM Inference Infrastructure: Agents Week 2026 Deep Dive" {
		t.Fatalf("FindBySlug() returned %q for Cloudflare inference infrastructure post", cloudflarePost.Title)
	}

	grokSprintPost, ok := FindBySlug("xai-grok-4-sprint-may-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find xAI Grok 4 sprint post")
	}
	if grokSprintPost.Title != "xAI's Grok 4 Sprint: Three Releases in Six Weeks Chasing GPT-5.5" {
		t.Fatalf("FindBySlug() returned %q for xAI Grok 4 sprint post", grokSprintPost.Title)
	}

	enterpriseAgentsPost, ok := FindBySlug("enterprise-ai-agent-failure-86-percent-governance-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find enterprise AI agent failure post")
	}
	if enterpriseAgentsPost.Title != "The 86% Enterprise AI Agent Failure Rate: Governance Crisis Explained" {
		t.Fatalf("FindBySlug() returned %q for enterprise AI agent failure post", enterpriseAgentsPost.Title)
	}

	reasoningPost, ok := FindBySlug("maturing-reasoning-models-adaptive-thinking-takes-center-stage")
	if !ok {
		t.Fatal("FindBySlug() did not find maturing reasoning models post")
	}
	if reasoningPost.Title != "Maturing Reasoning Models: Adaptive Thinking Takes Center Stage" {
		t.Fatalf("FindBySlug() returned %q for maturing reasoning models post", reasoningPost.Title)
	}

	geminiPost, ok := FindBySlug("gemini-3-1-ultra-native-multimodal-may-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Gemini 3.1 Ultra post")
	}
	if geminiPost.Title != "Gemini 3.1 Ultra: Why Google's Native Multimodal Architecture Is The Real Story" {
		t.Fatalf("FindBySlug() returned %q for Gemini 3.1 Ultra post", geminiPost.Title)
	}

	multiTokenPredictionPost, ok := FindBySlug("google-multi-token-prediction-drafters-3x-inference-revolution")
	if !ok {
		t.Fatal("FindBySlug() did not find Google multi-token prediction post")
	}
	if multiTokenPredictionPost.Title != "Google's Multi-Token Prediction Drafters: The 3x Inference Revolution That Changes Everything" {
		t.Fatalf("FindBySlug() returned %q for Google multi-token prediction post", multiTokenPredictionPost.Title)
	}

	costCollapsePost, ok := FindBySlug("ai-inference-cost-collapse-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find AI inference cost collapse post")
	}
	if costCollapsePost.Title != "From $30 to $0.40 Per Million Tokens: The AI Inference Cost Collapse That Redefines Enterprise AI" {
		t.Fatalf("FindBySlug() returned %q for AI inference cost collapse post", costCollapsePost.Title)
	}

	qwenPost, ok := FindBySlug("qwen-overtakes-llama-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find Qwen overtakes Llama post")
	}
	if qwenPost.Title != "Qwen Surpasses Llama: China's Open-Source AI Dominance and What It Means for Global Developers" {
		t.Fatalf("FindBySlug() returned %q for Qwen overtakes Llama post", qwenPost.Title)
	}

	mythosStandoffPost, ok := FindBySlug("mythos-national-security-standoff-2026-05-14")
	if !ok {
		t.Fatal("FindBySlug() did not find Mythos national security standoff post")
	}
	if mythosStandoffPost.Title != "Mythos National Security Standoff: The AI Model America Can't Agree On" {
		t.Fatalf("FindBySlug() returned %q for Mythos national security standoff post", mythosStandoffPost.Title)
	}

	caisiOversightPost, ok := FindBySlug("caisi-pre-release-ai-oversight-2026-05-14")
	if !ok {
		t.Fatal("FindBySlug() did not find CAISI pre-release oversight post")
	}
	if caisiOversightPost.Title != "CAISI Framework: US Government's Quiet Pivot to Pre-Release AI Oversight" {
		t.Fatalf("FindBySlug() returned %q for CAISI pre-release oversight post", caisiOversightPost.Title)
	}

	modelRushLatest, ok := FindBySlug("the-may-2026-model-rush-gpt-5-5-instant-subqs-long-context-and-grok-4-3")
	if !ok {
		t.Fatal("FindBySlug() did not find latest May model rush post")
	}
	if modelRushLatest.Title != "The May 2026 Model Rush: GPT-5.5 Instant, SubQ's Long Context, and Grok 4.3" {
		t.Fatalf("FindBySlug() returned %q for latest May model rush post", modelRushLatest.Title)
	}

	agenticPost, ok := FindBySlug("agentic-ai-trends-marketing-scale-enterprise-adoption-and-the-cost-paradox")
	if !ok {
		t.Fatal("FindBySlug() did not find agentic AI trends post")
	}
	if agenticPost.Title != "Agentic AI Trends: Marketing Scale, Enterprise Adoption, and the Cost Paradox" {
		t.Fatalf("FindBySlug() returned %q for agentic AI trends post", agenticPost.Title)
	}

	cyberPost, ok := FindBySlug("frontier-ai-as-cyber-weapons-gpt-5-5-tops-aisi-benchmarks-raising-urgent-safety-alarms")
	if !ok {
		t.Fatal("FindBySlug() did not find frontier cyber weapons post")
	}
	if cyberPost.Title != "Frontier AI as Cyber Weapons: GPT-5.5 Tops AISI Benchmarks, Raising Urgent Safety Alarms" {
		t.Fatalf("FindBySlug() returned %q for frontier cyber weapons post", cyberPost.Title)
	}

	modelRushPost, ok := FindBySlug("may-2026-ai-model-rush-12m-contexts-flash-speed-and-specialized-agents")
	if !ok {
		t.Fatal("FindBySlug() did not find May 2026 model rush post")
	}
	if modelRushPost.Title != "May 2026 AI Model Rush: 12M Contexts, Flash Speed, and Specialized Agents" {
		t.Fatalf("FindBySlug() returned %q for May 2026 model rush post", modelRushPost.Title)
	}

	ioPost, ok := FindBySlug("google-io-2026-preview-agentic-ai-gemma-4-and-the-cosmo-ghost-layer")
	if !ok {
		t.Fatal("FindBySlug() did not find Google I/O preview post")
	}
	if ioPost.Title != "Google I/O 2026 Preview: Agentic AI, Gemma 4, and COSMO's Ghost Layer" {
		t.Fatalf("FindBySlug() returned %q for Google I/O preview post", ioPost.Title)
	}

	ioKeynotePost, ok := FindBySlug("google-io-2026-ai-innovations-gemini-updates-and-android-xr-on-the-horizon")
	if !ok {
		t.Fatal("FindBySlug() did not find Google I/O 2026 keynote post")
	}
	if ioKeynotePost.Title != "Google I/O 2026: AI Innovations, Gemini Updates, and Android XR on the Horizon" {
		t.Fatalf("FindBySlug() returned %q for Google I/O 2026 keynote post", ioKeynotePost.Title)
	}

	ioTeaserPost, ok := FindBySlug("google-io-2026-ai-takes-center-stage-gemini-agentic-tools-android-xr-teasers")
	if !ok {
		t.Fatal("FindBySlug() did not find Google I/O 2026 teaser post")
	}
	if ioTeaserPost.Title != "Google I/O 2026: AI Innovations Take Center Stage with Gemini and Android XR Teasers" {
		t.Fatalf("FindBySlug() returned %q for Google I/O 2026 teaser post", ioTeaserPost.Title)
	}

	androidPost, ok := FindBySlug("googles-android-show-2026-android-17-gemini-4-and-the-next-wave-of-mobile-ai")
	if !ok {
		t.Fatal("FindBySlug() did not find Android Show 2026 post")
	}
	if androidPost.Title != "Google's Android Show 2026: Android 17, Gemini 4.0, and the Next Wave of Mobile AI" {
		t.Fatalf("FindBySlug() returned %q for Android Show 2026 post", androidPost.Title)
	}

	latestPost, ok := FindBySlug("us-governments-ai-policy-u-turn-caisi-framework-and-mythos-catalyst")
	if !ok {
		t.Fatal("FindBySlug() did not find CAISI policy u-turn post")
	}
	if latestPost.Title != "US Government's AI Policy U-Turn: The CAISI Framework and the 'Mythos' Catalyst" {
		t.Fatalf("FindBySlug() returned %q for CAISI policy u-turn post", latestPost.Title)
	}

	frontierPost, ok := FindBySlug("the-frontier-firm-is-here-microsoft-says-ai-has-moved-from-tool-to-operating-model")
	if !ok {
		t.Fatal("FindBySlug() did not find frontier firm post")
	}
	if frontierPost.Title != "The Frontier Firm Is Here: Microsoft Says AI Has Moved From Tool to Operating Model" {
		t.Fatalf("FindBySlug() returned %q for frontier firm post", frontierPost.Title)
	}

	securityPost, ok := FindBySlug("the-government-that-fears-its-own-weapon-how-mythos-became-americas-most-dangerous-ai-secret")
	if !ok {
		t.Fatal("FindBySlug() did not find Mythos security post")
	}
	if securityPost.Title != "The Government That Fears Its Own Weapon: How Mythos Became America's Most Dangerous AI Secret" {
		t.Fatalf("FindBySlug() returned %q for Mythos security post", securityPost.Title)
	}

	aiHackerPost, ok := FindBySlug("the-ai-hacker-how-claude-mythos-is-changing-cybersecurity-forever")
	if !ok {
		t.Fatal("FindBySlug() did not find AI hacker post")
	}
	if aiHackerPost.Title != "The AI Hacker: How Claude Mythos Is Changing Cybersecurity Forever" {
		t.Fatalf("FindBySlug() returned %q for AI hacker post", aiHackerPost.Title)
	}

	memoryPost, ok := FindBySlug("samsungs-trillion-dollar-moment-and-the-memory-bottleneck-that-will-define-ais-next-year")
	if !ok {
		t.Fatal("FindBySlug() did not find Samsung memory post")
	}
	if memoryPost.Title != "Samsung's Trillion-Dollar Moment and the Memory Bottleneck That Will Define AI's Next Year" {
		t.Fatalf("FindBySlug() returned %q for Samsung memory post", memoryPost.Title)
	}

	sciencePost, ok := FindBySlug("the-lab-without-scientists-when-ai-became-its-own-researcher")
	if !ok {
		t.Fatal("FindBySlug() did not find autonomous science post")
	}
	if sciencePost.Title != "The Lab Without Scientists: When AI Became Its Own Researcher" {
		t.Fatalf("FindBySlug() returned %q for autonomous science post", sciencePost.Title)
	}

	hardwarePost, ok := FindBySlug("the-chip-that-stays-home-inside-chinas-race-to-build-robotics-ai-hardware")
	if !ok {
		t.Fatal("FindBySlug() did not find China hardware post")
	}
	if hardwarePost.Title != "The Chip That Stays Home: Inside China's Race to Build Robotics AI Hardware" {
		t.Fatalf("FindBySlug() returned %q for China hardware post", hardwarePost.Title)
	}

	platformPost, ok := FindBySlug("apples-multi-ai-gambit-what-ios-27-reveals-about-the-platform-wars")
	if !ok {
		t.Fatal("FindBySlug() did not find Apple multi-AI post")
	}
	if platformPost.Title != "Apple's Multi-AI Gambit: What iOS 27 Reveals About the Platform Wars" {
		t.Fatalf("FindBySlug() returned %q for Apple multi-AI post", platformPost.Title)
	}

	specializedAgentsPost, ok := FindBySlug("rise-of-specialized-ai-agents-in-enterprise-workflows-2026-trends")
	if !ok {
		t.Fatal("FindBySlug() did not find specialized AI agents post")
	}
	if specializedAgentsPost.Title != "The Rise of Specialized AI Agents in Enterprise Workflows: 2026 Trends" {
		t.Fatalf("FindBySlug() returned %q for specialized AI agents post", specializedAgentsPost.Title)
	}

	agenticIndustryPost, ok := FindBySlug("rise-of-agentic-ai-autonomous-agents-reshaping-industries-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find agentic AI industries post")
	}
	if agenticIndustryPost.Title != "The Rise of Agentic AI: Autonomous Agents Reshaping Industries in 2026" {
		t.Fatalf("FindBySlug() returned %q for agentic AI industries post", agenticIndustryPost.Title)
	}

	openaiPost, ok := FindBySlug("openai-strategic-expansion-voice-tech-finance-tools-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find OpenAI strategic expansion post")
	}
	if openaiPost.Title != "OpenAI's Strategic Expansion: Acquiring Voice Tech and Launching Finance Tools" {
		t.Fatalf("FindBySlug() returned %q for OpenAI strategic expansion post", openaiPost.Title)
	}

	orbitalPost, ok := FindBySlug("orbital-data-centers-environmental-cost-ai-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find orbital data centers post")
	}
	if orbitalPost.Title != "The Next Frontier: Orbital Data Centers and the Environmental Cost of AI" {
		t.Fatalf("FindBySlug() returned %q for orbital data centers post", orbitalPost.Title)
	}

	post, ok := FindBySlug("sprint-is-real-inside-xai-grok-4-race-to-the-top")
	if !ok {
		t.Fatal("FindBySlug() did not find Grok sprint post")
	}
	if post.Title != "The Sprint Is Real: Inside xAI's Grok 4 Race to the Top" {
		t.Fatalf("FindBySlug() returned %q for Grok sprint post", post.Title)
	}

	policyPost, ok := FindBySlug("caisi-reversal-what-washingtons-sudden-policy-pivot-means-for-ai-development")
	if !ok {
		t.Fatal("FindBySlug() did not find CAISI reversal post")
	}
	if policyPost.Title != "The CAISI Reversal: What Washington's Sudden Policy Pivot Means for AI Development" {
		t.Fatalf("FindBySlug() returned %q for CAISI reversal post", policyPost.Title)
	}

	engineeringPost, ok := FindBySlug("googles-75-percent-stat-wake-up-call-software-engineers-needed")
	if !ok {
		t.Fatal("FindBySlug() did not find expected post")
	}

	if engineeringPost.Title != "Google's 75% Stat is the Wake-Up Call Software Engineers Needed" {
		t.Fatalf("FindBySlug() returned %q", engineeringPost.Title)
	}

	governancePost, ok := FindBySlug("broken-promise-worth-134-billion-openai-trial-ai-governance-under-oath")
	if !ok {
		t.Fatal("FindBySlug() did not find governance post")
	}
	if governancePost.Title != "A Broken Promise Worth $134 Billion: The OpenAI Trial Putting AI Governance Under Oath" {
		t.Fatalf("FindBySlug() returned %q for governance post", governancePost.Title)
	}

	openSourcePost, ok := FindBySlug("the-new-open-source-king-how-qwen-quietly-dethroned-llama")
	if !ok {
		t.Fatal("FindBySlug() did not find open source post")
	}
	if openSourcePost.Title != "The New Open-Source King: How Qwen Quietly Dethroned Llama" {
		t.Fatalf("FindBySlug() returned %q for open source post", openSourcePost.Title)
	}

	modelExplosionPost, ok := FindBySlug("may-2026s-ai-model-explosion-open-weight-models-reshape-the-landscape")
	if !ok {
		t.Fatal("FindBySlug() did not find open-weight model explosion post")
	}
	if modelExplosionPost.Title != "May 2026's AI Model Explosion: Open-Weight Models Reshape the Landscape" {
		t.Fatalf("FindBySlug() returned %q for open-weight model explosion post", modelExplosionPost.Title)
	}

	scienceDiscoveryPost, ok := FindBySlug("ai-accelerates-scientific-discovery-breakthroughs-in-medicine-and-research-may-2026")
	if !ok {
		t.Fatal("FindBySlug() did not find scientific discovery post")
	}
	if scienceDiscoveryPost.Title != "AI Accelerates Scientific Discovery: Breakthroughs in Medicine and Research May 2026" {
		t.Fatalf("FindBySlug() returned %q for scientific discovery post", scienceDiscoveryPost.Title)
	}

	publicationDate := time.Date(2026, time.May, 27, 0, 0, 0, 0, time.UTC)
	earlyDiseasePost, ok := FindBySlug("ai-early-disease-detection-breakthroughs")
	if time.Now().UTC().Before(publicationDate) {
		if ok {
			t.Fatal("FindBySlug() returned unpublished early disease detection post before publication date")
		}
	} else {
		if !ok {
			t.Fatal("FindBySlug() did not find early disease detection post on or after publication date")
		}
		if earlyDiseasePost.Title != "AI's Early Warning System: Detecting Diseases Years Before Symptoms" {
			t.Fatalf("FindBySlug() returned %q for early disease detection post", earlyDiseasePost.Title)
		}
	}

	governanceGapPost, ok := FindBySlug("the-agentic-ai-revolution-governance-gap")
	if time.Now().UTC().Before(publicationDate) {
		if ok {
			t.Fatal("FindBySlug() returned unpublished governance gap post before publication date")
		}
	} else {
		if !ok {
			t.Fatal("FindBySlug() did not find governance gap post on or after publication date")
		}
		if governanceGapPost.Title != "The Agentic AI Revolution: Adoption Surges But Governance Lags Dangerously Behind" {
			t.Fatalf("FindBySlug() returned %q for governance gap post", governanceGapPost.Title)
		}
	}

	if _, ok := FindBySlug("missing-post"); ok {
		t.Fatal("FindBySlug() reported a match for a missing slug")
	}
}

func TestPostsReturnsDeepCopy(t *testing.T) {
	got := Posts()
	got[0].Title = "mutated"
	got[0].Sections[0].Paragraphs[0] = "mutated"

	refetched := Posts()
	if refetched[0].Title == "mutated" {
		t.Fatal("Posts() exposed mutable post data")
	}
	if refetched[0].Sections[0].Paragraphs[0] == "mutated" {
		t.Fatal("Posts() exposed mutable section paragraph data")
	}

	post, ok := FindBySlug(refetched[0].Slug)
	if !ok {
		t.Fatal("FindBySlug() did not return the refetched post")
	}
	post.Sections[0].Paragraphs[0] = "mutated again"

	refetchedPost, ok := FindBySlug(refetched[0].Slug)
	if !ok {
		t.Fatal("FindBySlug() did not return the post on second lookup")
	}
	if refetchedPost.Sections[0].Paragraphs[0] == "mutated again" {
		t.Fatal("FindBySlug() exposed mutable section paragraph data")
	}
}
