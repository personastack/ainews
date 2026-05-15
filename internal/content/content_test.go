package content

import "testing"

func TestPostsReturnsPublishedPosts(t *testing.T) {
	got := Posts()
	if len(got) != 25 {
		t.Fatalf("Posts() returned %d posts, want 25", len(got))
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
