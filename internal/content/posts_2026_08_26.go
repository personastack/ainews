package content

func init() {
	posts = append([]Post{
		{
			Title:   "Thomson Reuters Built Its Own LLM for $40 Million. On Legal Work, It Beats GPT-5.4.",
			Slug:    "thomson-reuters-thomson-llm-westlaw-cocounsel-domain-2026",
			Date:    "August 26, 2026",
			Tag:     "Industry Trends",
			Summary: "Thomson Reuters' $40 million legal model is a case study in how proprietary data, expert alignment, and an open-weight base can outperform general-purpose systems on a tightly defined professional task.",
			Sections: []Section{
				{Paragraphs: []string{
					`For two decades, the conventional wisdom about large language models held roughly constant: bigger wins. More parameters, more compute, more data. The biggest models, from the biggest labs, with the biggest funding rounds, were supposed to own the market.`,
					`Thomson Reuters just complicated that story.`,
					`On August 24, the legal publishing giant launched Thomson 1.0, its proprietary large language model trained on the company's accumulated knowledge: more than 40,000 databases from Westlaw, decades of Practical Law guidance, Checkpoint tax research, and Reuters news content. The company says the effort cost roughly $40 million over two years, with the final training run costing less than $450,000.`,
					`On the tasks that matter to legal professionals, Thomson Reuters says the model does not just hold its own. Its internal benchmarks show stronger completeness and factuality than GPT-5.4 and Claude Sonnet 5 when the systems are connected to Thomson Reuters content. Independent verification is still developing, and the company has promised a technical report, but its framing is notably careful: broadly competitive with frontier models on web-only access, roughly equal or slightly better when plugged into TR data, and meaningfully ahead on document tasks where citations and authority matter most.`,
					`"Thomson needs to set the frontier of intelligence for legal," said Joel Hron, Thomson Reuters' chief technology officer.`,
				}},
				{Heading: "Built on Qwen, Powered by Decades of Law", Paragraphs: []string{
					`Thomson 1.0's most recent foundation is Qwen 3.5, the open-weight model from Alibaba Cloud. TR did not build from scratch. Instead, the team used a layered specialization approach: pre-training on proprietary content, targeted post-training guided by hundreds of subject-matter expert lawyers and tax professionals, then reinforcement learning that taught the model to navigate Westlaw and Practical Law tool interfaces directly.`,
					`That last step may be the most important. Westlaw's more than 40,000 databases span generations of legal publishing, and knowing how to retrieve from them is a skill that general-purpose models trained on public internet data fundamentally lack. TR's model does not just know legal content; it has been trained to navigate the retrieval system that organizes that content for professional use.`,
					`In academic testing cited by Thomson Reuters, the model's responses included links to treatises, making them more transparent and useful. Less than 10% of TR's available proprietary content has been used for training so far, according to the company.`,
				}},
				{Heading: "CoCounsel Goes Agentic — on Claude", Paragraphs: []string{
					`The architecture of TR's CoCounsel Legal AI platform shows how enterprise AI is actually being built in 2026. CoCounsel — the agentic layer that handles planning, multi-step reasoning, and workflow execution — is built on Anthropic's Claude Agent SDK. Thomson will be the default model for specific tasks within that system, beginning with Tabular Analysis, the high-volume document-review function where structured extraction and factual accuracy matter most.`,
					`The layers work like this: Claude Agent SDK orchestrates the workflow, understands intent, and manages the back-and-forth with users. Thomson handles legal-domain work — lookups, citations, and synthesis of dense primary sources — where proprietary training data gives it an edge.`,
					`This is not a single-model story. It is a layered architecture in which one vendor's agentic framework works alongside another organization's domain intelligence. CoCounsel administrators can configure alternative models for specific tasks, but the default bet is on TR's own model where its content advantage matters most.`,
				}},
				{Heading: "A New Business Model: API Licensing", Paragraphs: []string{
					`TR has not announced standalone pricing or availability for Thomson. The company says it has begun early conversations with large law firms and corporate legal departments about direct access and fine-tuning capabilities.`,
					`A smaller open-weight version is also planned for Hugging Face under a noncommercial academic license, a nod to the research and law-school communities that help define what legal AI tools need to do well. Customer data is not used for model training; Thomson Reuters says the model was built on its own proprietary and licensed content with expert guidance throughout.`,
				}},
				{Heading: "The Broader Pattern", Paragraphs: []string{
					`Thomson Reuters is not the first enterprise to pursue a domain-specific model strategy. Bloomberg built BloombergGPT, and Adobe trained Firefly on licensed image content. What has changed is that open-weight base models have become capable enough for a well-resourced enterprise to start from a strong foundation and achieve frontier-level performance within its domain without building a base model from scratch.`,
					`The reported $450,000 final training run is the number enterprise technology leaders will circle. It is within the budget of a major organization sitting on a large, clean proprietary data lake: legal publishers, financial-data providers, medical-record systems, and scientific publishers. The three ingredients TR assembled — a strong open-weight base, deep proprietary domain data, and expert-guided alignment — are potentially replicable by organizations with comparable assets.`,
					`Whether that matters depends on what you are building for. For tasks where specialized knowledge, citation accuracy, and domain-specific tool navigation matter, the TR playbook is now well documented. For general-purpose tasks or work requiring broad world knowledge, frontier models still hold the advantage. TR's competitive-on-the-web framing acknowledges that distinction.`,
					`The conventional wisdom is not wrong: bigger still wins in many domains. But Thomson is a signal that bigger is increasingly measured not in total parameter count, but in depth of relevant expertise. For a company that has spent generations accumulating exactly that expertise, the reframing is a considerable advantage.`,
				}},
				{Heading: "Sources", Paragraphs: []string{
					`Thomson Reuters — Thomson: a purpose-built foundation model for professionals: https://www.thomsonreuters.com/en/thomson-llm`,
					`Thomson Reuters — Next-generation CoCounsel Legal launch: https://www.thomsonreuters.com/en/press-releases/2026/august/thomson-reuters-launches-next-generation-of-cocounsel-legal-the-ai-ecosystem-built-for-legal-professionals`,
					`LawSites — Thomson Reuters launches Thomson: https://www.lawnext.com/2026/08/thomson-reuters-launches-thomson-its-own-proprietary-llm-trained-on-westlaw-and-practical-law-content.html`,
				}},
			},
			Related: []Link{
				{Title: "Not an Acquisition. Not an Acquihire. Nvidia Just Invented Something New.", Slug: "nvidia-poolside-6-billion-model-factory-license-2026"},
			},
		},
		{
			Title:   "The Inference Chip Nvidia Paid $20 Billion For Is Now Shipping. Meet the Groq 3 LPX.",
			Slug:    "nvidia-groq-3-lpx-full-production-hot-chips-2026-agentic-inference",
			Date:    "August 26, 2026",
			Tag:     "Infrastructure",
			Summary: "Nvidia's Groq 3 LPX is now in full production, turning the token-by-token decode bottleneck of agentic AI into a dedicated rack-scale hardware problem.",
			Sections: []Section{
				{Paragraphs: []string{
					`In December 2025, Nvidia reportedly paid roughly $20 billion to license inference technology from Groq Inc. and hired its founders. It was not an acquisition in the traditional sense — Groq the company still exists — but the technology, team, and architecture behind one of the fastest inference accelerators moved into Nvidia's AI factory platform.`,
					`At Hot Chips 2026 last week, Nvidia revealed what that deal was actually for.`,
					`The Groq 3 LPX — the third-generation version of the chip originally developed by Groq founder Jonathan Ross and his team — is now in full production. Production shipments begin this fall.`,
					`The headline number is 3,400 output tokens per second running Gemma 4 31B with a 100,000-token context window. Artificial Analysis benchmarked the configuration; Nvidia says that is four times faster in responsiveness than the nearest alternative platform for latency-sensitive workloads.`,
				}},
				{Heading: "The Agentic Bottleneck", Paragraphs: []string{
					`To understand why this matters, it helps to understand where AI systems actually slow down. Training a large model is slow by design, and prefill — the step where a model ingests a long document or conversation history — can tolerate some latency. The decode step, where the model generates tokens one at a time to build a response, is where users feel the difference between a system that feels alive and one that feels like it is buffering.`,
					`For agentic AI workflows — systems that reason through problems, plan actions, call external tools, and loop back to revise their work — decode runs constantly. An agent that writes code, checks its output, runs tests, and iterates may execute dozens of decode cycles per task. Slow token generation compounds at every step in that chain.`,
					`That was Groq's original insight: build silicon optimized for the sequential, memory-bandwidth-intensive nature of token-by-token generation. Not for training, and not primarily for prefill — for making tokens come out fast and predictably.`,
					`Nvidia's Vera Rubin architecture reflects the split directly. Rubin GPUs handle context ingestion and the portions of inference best suited to high-throughput compute, while Groq 3 LPX accelerates latency-sensitive decode work. Together, Nvidia says, the pair expands the usual throughput-versus-response-time tradeoff. That remains a vendor claim, but it is a specific architectural bet rather than just a faster-chip slogan.`,
				}},
				{Heading: "Platform Architecture: Five Racks, Seven Chips", Paragraphs: []string{
					`The Vera Rubin platform combines Groq 3 LPX with the Vera CPU, Rubin GPU, BlueField-4 storage, and Spectrum-6 networking in five purpose-built rack types that operate as a unified AI supercomputer. A single LPX rack contains 256 Groq 3 LPU accelerators connected by high-bandwidth chip interconnects.`,
					`The LPX system itself is designed around deterministic, compiler-scheduled execution and SRAM-first memory. Nvidia lists 128 GB of total SRAM, 40 PB/s of on-chip SRAM bandwidth, and 640 TB/s of scale-up bandwidth for the rack. The purpose is to minimize the jitter that becomes visible in small-batch, interactive inference.`,
					`The broader NVL72 platform's capacity, power, cooling, and networking choices are consequential because this is not a component intended to live in isolation. Nvidia is pitching a full-stack operating model for running AI at factory scale, including liquid-cooled, cableless compute trays and a software layer that can route work between GPUs and LPUs.`,
				}},
				{Heading: "First Customer: Nebius Token Factory", Paragraphs: []string{
					`Nebius is the first publicly confirmed customer planning to deploy Groq 3 LPX, integrating it into its Token Factory production inference platform. The name is worth noticing.`,
					`Token Factory reflects a shift in how enterprise AI infrastructure is being conceptualized: not simply as a cloud API routing requests to shared models, but as a production process with defined throughput, quality, and reliability targets. Buying dedicated inference silicon rather than renting capacity from a hyperscaler is a bet on predictable demand at scale.`,
					`The economics of that bet depend on volumes most organizations do not yet have. For cloud inference providers such as Nebius, however, the hardware math can favor ownership once utilization is consistently high.`,
				}},
				{Heading: "What Comes Next", Paragraphs: []string{
					`Production shipments begin in fall 2026, so most buyers will not have hardware before year end. Nebius is the only confirmed public commitment so far, and the 3,400-tokens-per-second figure is specific to one model and context length. Real-world performance across workloads will vary.`,
					`Still, the signal from Hot Chips is clear: Nvidia is treating agentic AI as a primary design target for its next platform generation, not an accommodation for an existing GPU architecture. The explicit split between context processing and token generation — handled by silicon tuned for each task — is Nvidia's architectural bet on how AI workloads will be structured for the next several years.`,
					`That bet is tied directly to the Groq deal. Nvidia did not need to spend $20 billion to build a general-purpose inference chip. It needed to address the bottleneck that makes agentic AI feel slow, and the fastest route was to license technology from the team that had already concentrated on that problem.`,
				}},
				{Heading: "Sources", Paragraphs: []string{
					`NVIDIA Newsroom — NVIDIA Groq 3 LPX now in full production: https://nvidianews.nvidia.com/news/nvidia-groq-3-lpx-now-in-full-production-with-world-class-speed-for-agentic-ai`,
					`NVIDIA Technical Blog — Inside NVIDIA Groq 3 LPX: https://developer.nvidia.com/blog/inside-nvidia-groq-3-lpx-the-low-latency-inference-accelerator-for-the-nvidia-vera-rubin-platform/`,
					`NVIDIA Technical Blog — Groq 3 LPX long-context benchmark details: https://developer.nvidia.com/blog/how-nvidia-groq-3-lpx-unlocks-ultrafast-interactivity-at-long-context-on-nvidia-vera-rubin/`,
				}},
			},
			Related: []Link{
				{Title: "Not an Acquisition. Not an Acquihire. Nvidia Just Invented Something New.", Slug: "nvidia-poolside-6-billion-model-factory-license-2026"},
				{Title: "AI Gets Cheaper to Use. The Hardware That Runs It Is About to Get Much More Expensive.", Slug: "nvidia-ai-server-prices-15-percent-dram-hbm-shortage-2027"},
			},
		},
	}, posts...)
}
