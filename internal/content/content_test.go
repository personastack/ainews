package content

import "testing"

func TestPostsReturnsPublishedPosts(t *testing.T) {
	got := Posts()
	if len(got) != 69 {
		t.Fatalf("Posts() returned %d posts, want 69", len(got))
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

func TestFindBySlug(t *testing.T) {
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
