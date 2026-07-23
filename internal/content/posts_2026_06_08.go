package content

func init() {
	posts = append([]Post{
		{
			Title:   "The AI Rulebook Is Moving From Principles to Plumbing",
			Slug:    "ai-policy-rulebook-principles-to-plumbing-2026",
			Date:    "June 8, 2026",
			Tag:     "Policy",
			Summary: "AI regulation is shifting from broad principles to implementation machinery, with Europe and the US both building concrete rules around security, procurement, benchmarking, and operational governance.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"For most of the generative AI boom, policy debates have been framed around big nouns: safety, innovation, sovereignty, transparency, national security. In June 2026, the more important story is that governments are starting to translate those nouns into operational machinery.",
						"That shift is less dramatic than a model launch, but it may matter more to companies building with AI. The new policy work is increasingly about deadlines, procurement, security testing, watermarking, vulnerability discovery, infrastructure, cloud capacity, and who inside government gets to define the next set of implementation rules.",
					},
				},
				{
					Heading: "Europe's Compliance Era Becomes An Implementation Era",
					Paragraphs: []string{
						"The European Commission's AI Office has become one of the most important bodies to watch because the AI Act cannot operate on statutory text alone. It needs guidance, codes of practice, templates, enforcement expectations, and a practical understanding of how general-purpose AI systems and high-risk AI systems will be evaluated.",
						"Passing a law is only the beginning. The hard work starts when a company asks whether a customer-service agent, a recruiting tool, a medical assistant, a coding assistant, or a synthetic-media workflow falls into a particular category and what evidence it needs to keep.",
						"Recent legal analyses of EU negotiations point to an effort to adjust parts of the AI Act timeline and implementation burden, especially around high-risk systems and transparency obligations. Those details still need careful treatment because not every proposal or political agreement is the same as final enforceable law. But the direction is clear enough: Europe is trying to keep the law's risk-based architecture while making the rollout less brittle for companies that have to comply with it.",
						"The key practical lesson for AI builders is that the EU AI Act is becoming a product-management issue, not just a legal issue. Transparency labels, model documentation, incident handling, risk management, data governance, and human oversight cannot be bolted onto a product a week before a deadline. They have to become part of release planning.",
						"This is especially true for companies building agentic systems. Once an AI system can call tools, take actions, summarize records, initiate workflows, or influence decisions, governance stops being only about model output. It becomes about permissions, logging, escalation paths, and whether a human can understand why something happened.",
					},
				},
				{
					Heading: "The White House Puts Cyber Defense at the Center",
					Paragraphs: []string{
						"The June 5 White House executive order, titled \"Promoting Advanced Artificial Intelligence Innovation and Security,\" takes a different route. It frames AI less as a general compliance problem and more as a strategic technology that has to be secured, measured, and used by government.",
						"The order directs federal agencies to prioritize cyber defense for federal systems and critical infrastructure, support access to AI-enabled cybersecurity tools, and create an AI cybersecurity clearinghouse aimed at identifying and remediating software vulnerabilities at scale. It also calls for classified benchmarking of advanced cyber capabilities in AI models and a process for identifying \"covered frontier models.\"",
						"That last phrase is important. The order says it does not create a mandatory licensing or pre-clearance regime for AI models. Instead, it establishes a voluntary framework for developers to consult with the federal government about whether models meet the covered-frontier-model designation.",
						"That is a revealing compromise. Washington wants visibility into frontier-model cyber capabilities without presenting the policy as a licensing system. For frontier labs, the practical consequence is likely to be more structured engagement with federal agencies around evaluations, security capabilities, and risk signals. For enterprise AI buyers, the signal is that AI security is becoming procurement-relevant. A model's capabilities will matter, but so will the developer's security posture and willingness to participate in government-facing risk processes.",
						"The order also reflects a broader shift in how AI cybersecurity is being understood. AI is not only a technology that might introduce new risks. It is also being treated as a defensive tool for software vulnerability discovery, federal system hardening, and critical infrastructure protection. That dual-use tension is now the policy baseline.",
					},
				},
				{
					Heading: "Sovereignty Is No Longer Just About Chips",
					Paragraphs: []string{
						"Europe's June 3 technological sovereignty package adds another layer to the story. The Commission described a package covering semiconductors, AI, cloud, and open source, with plans tied to a Chips Act 2.0 and a Cloud and AI Development Act.",
						"That matters because AI policy is no longer only about regulating outputs. It is also about control over the infrastructure stack: chips, cloud services, open-source dependencies, data centers, and the public-sector capacity to use and audit AI systems.",
						"For readers following the AI industry, this is where regulation and infrastructure meet. A government can demand transparency, but it also needs compute, technical staff, cloud access, standards, and evaluation capacity. A company can promise safer deployment, but it needs logging systems, governance tooling, secure software practices, and a way to explain agent behavior after the fact.",
						"The next phase of AI policy is therefore likely to be less theatrical and more consequential. Instead of one sweeping debate over whether AI should be regulated, we will see dozens of smaller decisions about audit trails, procurement language, technical standards, security benchmarks, synthetic-media labels, and the legal status of model-risk documentation.",
					},
				},
				{
					Heading: "What AI Teams Should Take From This",
					Paragraphs: []string{
						"The immediate takeaway is not panic. It is preparation.",
						"Companies building AI products should assume that three disciplines are converging: security engineering, compliance operations, and product design. A useful AI system will need to be capable. A deployable AI system will also need to be governable.",
						"That means teams should know which systems can take action, which systems influence regulated decisions, which systems generate public-facing synthetic content, which systems touch sensitive data, and which systems depend on frontier models with capabilities that may attract government attention.",
						"The policy mood in June 2026 is not simply 'more regulation.' It is more infrastructure around regulation: offices, codes, benchmarks, clearinghouses, implementation dates, and sovereignty programs. That can feel bureaucratic. It is also what happens when a technology moves from experimental demo to institutional dependency.",
						"The companies that handle this well will not be the ones that treat policy as a press-release problem. They will be the ones that can show how their AI systems behave, how they fail, how they are secured, and who is accountable when they act.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"European Commission: Strengthening Europe's tech sovereignty, June 3, 2026: https://commission.europa.eu/news-and-media/news/strengthening-europes-tech-sovereignty-2026-06-03_en",
						"European Commission AI Office: https://digital-strategy.ec.europa.eu/en/policies/ai-office",
						"White House Executive Order: Promoting Advanced Artificial Intelligence Innovation and Security, June 5, 2026: https://www.whitehouse.gov/presidential-actions/2026/06/promoting-advanced-artificial-intelligence-innovation-and-security/",
						"White House fact sheet on the AI innovation and security order, June 5, 2026: https://www.whitehouse.gov/fact-sheets/2026/06/fact-sheet-president-donald-j-trump-promotes-advanced-artificial-intelligence-innovation-and-security/",
						"White House fact sheet on AI in the national security enterprise, June 5, 2026: https://www.whitehouse.gov/fact-sheets/2026/06/fact-sheet-president-donald-j-trump-signs-historic-directive-on-ai-in-the-national-security-enterprise/",
						"White House NSPM-11, June 5, 2026: https://www.whitehouse.gov/presidential-actions/2026/06/national-security-presidential-memorandum-nspm-11/",
						"Latham & Watkins AI Act update analysis: https://www.lw.com/en/insights/ai-act-update-eu-resolves-to-change-rules-and-extend-deadlines",
						"Stibbe analysis of proposed AI Act changes: https://www.stibbe.com/publications-and-insights/ai-act-reloaded-what-the-latest-ai-act-changes-mean-in-practice",
					},
				},
			},
		},
		{
			Title:   "The CPU Returns to the AI Story: Intel's Rackscale Bet on Agentic Inference",
			Slug:    "intel-rackscale-agentic-inference-cpu-comeback-2026",
			Date:    "June 8, 2026",
			Tag:     "Infrastructure",
			Summary: "Intel is arguing that the next phase of AI will not be won by GPUs alone. As agents spend more time planning, calling tools, retrieving context, and coordinating work, the rack around the accelerator may become the real bottleneck.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"For most of the generative AI boom, the hardware story has been easy to summarize: GPUs won, everyone else chased. The biggest training runs, the worst shortages, and the largest capital plans all pointed toward accelerators as the center of gravity.",
						"Intel is now trying to complicate that story.",
					},
				},
				{
					Heading: "Why Rackscale Matters",
					Paragraphs: []string{
						"At Computex 2026, Intel framed its AI pitch around a shift from model training to large-scale inference, and from single-box performance to rackscale systems. The company is arguing that the center of value is moving from a lone accelerator to the full stack around it.",
						"That matters because AI products are becoming workflow engines. A production agent does not just answer one prompt. It plans, retrieves, calls tools, waits on external systems, reflects on intermediate results, and coordinates multiple steps across software boundaries.",
					},
				},
				{
					Heading: "Agentic Inference Changes The Bottleneck",
					Paragraphs: []string{
						"Intel says AI inference could become nearly 40% of all data center power demand by 2030. It also argues that agentic AI changes the traditional server balance: where frontier training clusters were often discussed in terms of one CPU feeding many GPUs, agentic inference could move closer to a 1:1 CPU-to-GPU ratio, or even make CPU capacity the limiting factor in some deployments.",
						"That claim should be treated carefully, because Intel clearly benefits if buyers believe CPUs are moving back toward the center of AI infrastructure. But it is directionally plausible. The more AI systems are asked to coordinate memory, policy, authentication, storage, monitoring, and application logic, the more the surrounding infrastructure becomes part of the performance envelope.",
					},
				},
				{
					Heading: "The Rack Is Where The Cost Shows Up",
					Paragraphs: []string{
						"The old AI question was: how fast can the model run?",
						"The new one is: how much useful work can the entire system complete per watt, per dollar, and per rack?",
						"That is where energy efficiency stops being a nice-to-have and starts becoming strategic. Training runs can be scheduled and treated as episodic events. Inference is always on. Agentic inference is worse in one specific way: it can multiply the number of steps behind each user request.",
					},
				},
				{
					Heading: "Intel's Platform Bet",
					Paragraphs: []string{
						"The company highlighted Xeon 6+ processors with 288 E-cores, 576MB of L3 cache, and Intel 18A technology. It also pointed to rackscale AI infrastructure work with partners including SambaNova, Foxconn, Vista Equity Partners, and Cambium Equity.",
						"The pitch is not that CPUs replace accelerators. It is that inference-heavy AI systems need a denser, more balanced fabric around them. The accelerator still matters, but it is no longer the only scarce part of the stack.",
					},
				},
				{
					Heading: "Why This Fits The Agent Era",
					Paragraphs: []string{
						"Enterprise agents have to read purchase orders, check policy, update tickets, ask for approval, retry failed tool calls, and leave an audit trail. Consumer agents have to pull from search, maps, shopping, mail, calendars, and payments without feeling slow or unpredictable.",
						"Those jobs do not map cleanly onto the idea that AI infrastructure is just a pile of accelerators. Once the agent becomes the product, the rest of the rack starts to look like the product too.",
					},
				},
				{
					Heading: "Execution Is Still The Hard Part",
					Paragraphs: []string{
						"Nvidia still owns the strongest position in the accelerator narrative, and AMD is competing aggressively in GPUs. Intel's best opening is to argue that the AI infrastructure market is widening. If the bottleneck moves from pure model training to inference orchestration, CPU density, platform integration, networking, and power efficiency, Intel has a more natural place to compete.",
						"The hard part is execution. Buyers do not make infrastructure decisions from press releases. They buy ecosystems, software compatibility, cloud availability, and proven economics. For Intel's rackscale story to matter, customers will need to see measurable gains in real workloads and a software path that does not create friction.",
					},
				},
				{
					Heading: "What This Actually Means",
					Paragraphs: []string{
						"The industry is entering a phase where model quality alone is no longer enough to explain market winners. The same frontier models can be wrapped in very different products, and the difference often comes from memory, tools, permissions, latency, governance, and cost control.",
						"That makes the CPU newly interesting. Not glamorous, but structurally important. If AI agents become the default interface for work, commerce, customer service, software development, and operations, then the infrastructure problem starts to resemble distributed systems again.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Intel Newsroom: https://newsroom.intel.com/artificial-intelligence/computex-2026-an-intelligent-world-built-on-silicon",
						"Intel Newsroom: https://newsroom.intel.com/artificial-intelligence/intel-announces-new-ai-innovations-at-computex",
					},
				},
			},
		},
		{
			Title:   "The Memory Layer Arrives: Microsoft's Build 2026 Shows What Enterprise Agents Need Next",
			Slug:    "microsoft-build-2026-agent-memory-layer",
			Date:    "June 8, 2026",
			Tag:     "Enterprise",
			Summary: "Microsoft Build 2026 frames enterprise AI around a shared intelligence layer, grounded retrieval, agent memory, and security controls that make agents usable inside real business systems.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"Microsoft Build 2026 did not read like a simple product launch. It read like an admission that enterprise AI has moved past the novelty phase and into infrastructure design.",
						"The company is now presenting a stack built around Microsoft IQ, Work IQ, Foundry IQ, and Web IQ, with Microsoft Foundry, Azure AI Search, Foundry Agent Service, memory tooling, and security controls all lining up behind the same thesis: agents are only useful when they can remember, retrieve, act, and stay governed inside the systems businesses already run.",
					},
				},
				{
					Heading: "The Intelligence Layer Is The Point",
					Paragraphs: []string{
						"The most important change is not a single model or feature. It is the way Microsoft is packaging context. Microsoft IQ now appears as the umbrella, with Work IQ for workplace signals, Foundry IQ for organizational knowledge, Fabric IQ for business data, and Web IQ for fresh web-grounded intelligence.",
						"That matters because enterprise agents fail less often on raw model quality than on missing context. A model can only be as helpful as the data it can ground itself in, the permissions it understands, and the tools it can safely call. Microsoft is trying to make that layer a product instead of a custom integration project.",
					},
				},
				{
					Heading: "Work IQ Turns Memory Into A Business Primitive",
					Paragraphs: []string{
						"Work IQ is where Microsoft makes the enterprise memory story concrete. The June 2 Work IQ API announcement describes a system that continuously builds semantic understanding from email, calendar, meetings, chats, files, people, and line-of-business systems, then exposes that understanding through APIs designed for agents.",
						"In practice, that means memory is no longer just a chat history feature. It becomes a managed layer for work context, operational state, and organizational patterns. For enterprises, that is the difference between a demo assistant and a system that can actually know how work gets done.",
					},
				},
				{
					Heading: "Foundry IQ Makes Retrieval Agentic",
					Paragraphs: []string{
						"Foundry IQ pushes the same idea into enterprise knowledge retrieval. Microsoft describes it as a knowledge layer that uses agentic retrieval to plan queries, choose sources, search, respond, and iterate, while enforcing user permissions and returning grounded answers.",
						"The architecture is important because it changes retrieval from a passive search step into an active orchestration step. Instead of asking a system to hand back documents, the agent can ask for context, refine the search, and keep narrowing toward the answer it needs without leaving the controlled enterprise boundary.",
					},
				},
				{
					Heading: "Azure AI Search Is Becoming The Plumbing",
					Paragraphs: []string{
						"Microsoft is also making Azure AI Search the practical backbone behind that knowledge layer. The product pages now position Azure AI Search and Foundry IQ together, with agentic retrieval, reusable knowledge bases, and support for sources like Work IQ, Fabric IQ, SharePoint, OneLake, blob storage, the web, and MCP servers.",
						"The strategic read is straightforward: if agents need durable memory and enterprise knowledge, the search layer stops being an optional add-on and becomes core infrastructure. That is exactly how relational databases and identity systems became mandatory in earlier software eras.",
					},
				},
				{
					Heading: "Foundry Agent Service Is The Runtime",
					Paragraphs: []string{
						"The other half of the stack is runtime. Microsoft Foundry is positioning Foundry Agent Service as the place where developers build, run, and govern agents that use this shared context.",
						"That makes the platform feel less like a chat wrapper and more like an operating environment. Developers can use the frameworks they already prefer, connect to memory and retrieval layers, and rely on a service boundary for control, scale, and monitoring instead of stitching all of that together themselves.",
					},
				},
				{
					Heading: "Security Is Now A First-Class Design Constraint",
					Paragraphs: []string{
						"Microsoft Build 2026 also made clear that enterprise AI will not be judged only on capability. It will be judged on whether it can survive adversarial conditions, policy checks, and production failure modes.",
						"The security announcements point to a new trust stack: ASSERT for policy-driven safety evaluation, the Agent Control Specification for control placement in the agent loop, Microsoft Security's multi-model agentic scanning harness, tracing and evaluations across frameworks, and governance controls that span code, agents, and models. That is the right answer to the current market problem, because the market no longer needs more autonomous behavior without guardrails. It needs evidence that autonomy can be constrained, inspected, and tested before it touches business data.",
					},
				},
				{
					Heading: "What This Actually Means For Enterprises",
					Paragraphs: []string{
						"The practical implication is that enterprise agents are being redesigned as stateful systems, not transient prompts. They need memory for continuity, retrieval for relevance, permissions for safety, shared state for multi-step work, and observability for trust.",
						"That is a bigger shift than a model benchmark. It is the emergence of a real agent platform stack. If Microsoft executes, the winning enterprise pattern will not be the smartest standalone model. It will be the system that can remember the right things, retrieve the right context, act inside the right boundaries, and explain what happened afterward.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Microsoft Build 2026 overview and keynote hub: https://news.microsoft.com/build-2026/",
						"Microsoft official blog: Microsoft Build 2026: Be yourself at work: https://blogs.microsoft.com/blog/2026/06/02/microsoft-build-2026-be-yourself-at-work/",
						"Microsoft 365 Blog: Announcing the new Work IQ APIs: https://www.microsoft.com/en-us/microsoft-365/blog/2026/06/02/announcing-the-new-work-iq-apis/",
						"Microsoft IQ overview: https://www.microsoft.com/en-us/ai/microsoft-iq",
						"Microsoft Azure Foundry IQ product page: https://azure.microsoft.com/en-us/products/ai-foundry/iq",
						"Azure AI Search product page: https://azure.microsoft.com/en-us/products/ai-services/ai-search",
						"Microsoft Foundry Blog: Build and run agents at scale with Microsoft Foundry at Build 2026: https://devblogs.microsoft.com/foundry/agent-service-build2026/",
						"Microsoft Security Blog: Microsoft Build 2026: Securing code, agents, and models across the development lifecycle: https://www.microsoft.com/en-us/security/blog/2026/06/02/microsoft-build-2026-securing-code-agents-and-models-across-the-development-lifecycle/",
					},
				},
			},
		},
		{
			Title:   "When Search Becomes a Workplace: Google's Agentic Gemini Strategy Comes Into Focus",
			Slug:    "google-search-agent-workbench-gemini-2026",
			Date:    "June 8, 2026",
			Tag:     "Platforms",
			Summary: "Google's May AI roundup is less a list of product updates than a map of where consumer AI is going next: agents embedded directly into Search, Android, shopping, hardware, health, and creative work.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"Google's latest AI recap has the shape of a product roundup, but the strategic message is bigger than any single Gemini model name. Across its May 2026 announcements, Google is describing a company-wide bet that AI agents will not live in a separate chatbot window for long. They will sit inside Search, phones, shopping flows, productivity interfaces, health apps, creative tools, hardware, and cloud infrastructure.",
						"That makes Google's May slate worth reading as a platform strategy. The central question is no longer whether Gemini can answer questions. It is whether Google can turn its existing surfaces into places where agents watch, generate, compare, buy, organize, and act.",
					},
				},
				{
					Heading: "Search as a Workbench",
					Paragraphs: []string{
						"The most important signal is Search. According to Google's June 5 recap of its May AI announcements, the company is pushing information agents that can monitor topics in the background, send updates, provide links, and suggest actions. It also says Gemini 3.5 Flash and Antigravity-style coding capabilities are coming into Search for generative interfaces, including interactive visuals, dashboards, and mini apps.",
						"That is a meaningful shift. Search has spent most of its life as a retrieval machine: type a query, receive ranked links, decide what to do next. The AI version Google is sketching is closer to a workbench. A user might ask for analysis, get a custom interface, watch an agent track changes, and move from answer to action without leaving Google's environment.",
					},
				},
				{
					Heading: "Distribution Becomes Intelligence",
					Paragraphs: []string{
						"This is why the roundup's many product names matter less than their placement. Gemini 3.5 is framed as a frontier model for agents and coding. Gemini Omni is positioned around multimodal reasoning and high-quality video creation. The Gemini app becomes more proactive. Android Halo is described as a layer for managing agents on phones. Universal Cart connects shopping actions across Google surfaces. Gemini Intelligence extends the same logic into Android devices.",
						"Put together, those are not isolated launches. They are distribution points.",
					},
				},
				{
					Heading: "The Platform Advantage",
					Paragraphs: []string{
						"The distribution advantage is obvious. Google already owns several of the places where people begin digital tasks: search, browser, email-adjacent productivity, mobile operating systems, maps, shopping discovery, and video. If agents become useful only after they know a user's context and can reach the tools needed to act, then distribution becomes a form of intelligence. The best model may not win every workflow; the model that is already present at the moment of intent may win many of them.",
						"There is a reason the agent story keeps pulling in hardware. Google's May recap also highlights Googlebook, intelligent eyewear, Google Health, Fitbit Air, and broader Android updates. These products point toward ambient AI: systems that know whether a user is at a desk, on a phone, wearing a device, managing health data, or shopping in a browser. That context can make agents more useful. It can also make the permission problem sharper.",
					},
				},
				{
					Heading: "The Permission Problem",
					Paragraphs: []string{
						"The next platform fight may be less about prompts than control. If an agent can monitor information, generate a dashboard, recommend a purchase, move items into a cart, summarize health signals, or produce creative output from multimodal inputs, users need clear answers to basic questions. What did the agent see? What can it remember? Which services can it act inside? When does it ask permission? How can a user audit what happened afterward?",
						"Google appears aware of that tension. The same roundup points to content transparency tools across Search, Gemini, Chrome, Pixel, and Cloud. Provenance is becoming a default requirement for companies that want AI-generated text, images, video, and actions to be trusted. But provenance alone will not solve the agent era's harder problem. A watermark can help identify generated media. It does not fully explain why an agent chose one action over another, or whether a user meant to grant that authority in the first place.",
					},
				},
				{
					Heading: "Generative UI Changes the Economics",
					Paragraphs: []string{
						"For developers, the generative UI thread may be the most consequential part of the update. If Search can generate dashboards, visualizations, and mini apps on demand, the boundary between searching for software and producing a temporary tool starts to blur. Many workflows that once required a spreadsheet, a lightweight web app, or a custom internal dashboard could begin as a query. That is powerful for users. It is also disruptive for software companies whose value is packaging common decisions into repeatable interfaces.",
						"This does not mean traditional apps disappear. Durable systems still need permissions, data governance, integrations, audit logs, reliability, and accountability. But it does mean the entry point changes. A user may increasingly start with an agent-generated interface, then graduate only the repeated or regulated workflows into permanent software.",
					},
				},
				{
					Heading: "Science, Infrastructure, And Lock-In",
					Paragraphs: []string{
						"Google's science and infrastructure announcements fit the same pattern from a different angle. The roundup references Gemini for Science, AlphaEvolve, a Google DeepMind APAC climate and energy accelerator, and REPLIQA, a $10 million life sciences and Quantum AI research initiative. These are not consumer agent features, but they reinforce the broader message: Google wants Gemini to be both everyday assistant and research instrument, both shopping helper and scientific substrate.",
						"The risk is complexity. Google's AI portfolio is now sprawling: Gemini variants, Search agents, Android agent management, shopping surfaces, hardware, health, creative media, science, cloud, provenance, and developer tools. That breadth is strategically useful, but it can become difficult for users and businesses to understand what is available, what is stable, and what can be trusted for real work.",
					},
				},
				{
					Heading: "What The Strategy Really Is",
					Paragraphs: []string{
						"The competitive risk is lock-in. The more useful agents become inside Google's ecosystem, the more valuable Google's ecosystem becomes. That is not new in technology, but agents raise the stakes because they do not merely display information. They may increasingly decide what information to watch, which options to surface, which interface to generate, and which action to suggest.",
						"The best version of this future is genuinely useful: less repetitive searching, more adaptive tools, faster research, more accessible software, and AI assistance that follows people across the devices and services they already use. The weaker version is a confusing layer of semi-autonomous features that are hard to inspect and harder to leave.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Google Keyword: The latest AI news we announced in May 2026, published June 5, 2026: https://blog.google/innovation-and-ai/technology/ai/google-ai-updates-may-2026/",
					},
				},
			},
		},
	}, posts...)
}
