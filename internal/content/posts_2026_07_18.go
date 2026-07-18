package content

func init() {
	posts = append([]Post{
		{
			Title:   "The Chip Industry Just Had Its Best Quarter Ever. Wall Street Sold It Anyway.",
			Slug:    "chip-earnings-record-profits-stock-selloff-kimi-k3-2026",
			Date:    "July 18, 2026",
			Tag:     "Hardware",
			Summary: "Record results from TSMC, Samsung, and ASML collided with a semiconductor bear market after Moonshot's Kimi K3 release forced investors to rethink whether AI compute demand is a training arms race, an inference boom, or both.",
			Related: []Link{
				{
					Title: "Nvidia's Roadmap Just Hit the Reticle Limit",
					Slug:  "nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026",
				},
				{
					Title: "The Rocket Company Ships a Coding Model — And the Benchmark Depends on Who's Grading",
					Slug:  "grok-4-5-spacexai-cursor-coding-benchmark-harness-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"In the same week that TSMC, Samsung, and ASML each reported the strongest results in their history, the Philadelphia Semiconductor Index fell into a bear market. That's not a contradiction that resolves itself with a little more context — it's the story. The companies making AI's physical infrastructure have never been more profitable, and investors have rarely been more nervous about what that profitability is actually worth.",
						"Start with the numbers, because they're genuinely startling. TSMC posted second-quarter revenue of $40.2 billion, up 36% year over year, with net income up 77.4% — a record for the fifth consecutive quarter, driven by full utilization of its advanced-node capacity and an AI-heavy product mix that now makes up 77% of revenue. Gross margin hit 67.7%, an all-time high. CEO C.C. Wei used the earnings call to announce another $100 billion for TSMC's Arizona buildout, bringing the company's total committed U.S. investment to $265 billion. Samsung's preliminary Q2 operating profit came in at 89.4 trillion won — roughly $59 billion, about 19 times what it earned a year earlier, beating analyst estimates by more than 6% on the back of simultaneous price increases across DRAM, NAND, and HBM. And ASML, the sole supplier of the extreme ultraviolet lithography tools that make advanced chips possible, raised its full-year sales guidance for the second time this year, to €43–45 billion, and said its order book is essentially full through 2027 with meaningful demand already visible for 2028.",
						"There was a genuine engineering milestone tucked into that same week, too. On July 15, Intel became the first company to ship high-volume logic chips patterned with ASML's High-NA (0.55 numerical aperture) EUV scanners — years ahead of TSMC, which has ruled out the technology through 2029. The chips are select layers of Intel's Panther Lake laptop processors, built on the 18A node at Intel's Hillsboro, Oregon fab, and ASML says those layers are already yielding at parity with the older NXE EUV platform. It's a real technical win for Intel's foundry ambitions — worth noting precisely because it's narrow. It applies to specific layers of chips already shipping to consumers, not a verdict on 18A's overall profitability for the external customers Intel is trying to win back; separate reporting this month suggested those customers may not see profitable yields out of 18A until late 2026 or 2027. Genuine progress and unresolved risk are sitting right next to each other in the same node.",
						"None of that record-breaking is what moved the stock market. On Friday, July 17, the SOX benchmark shed as much as 5.7% in a single session, extending a slide that has now erased more than 20% of the index's value since its late-June record — the technical threshold for a bear market — and roughly $3.3 trillion in global chip-sector market value since June 22 alone. It was the worst weekly rout for the sector since April 2025. Marvell, Arm, and Intel are each down more than 30% from that June peak. This is the same index that had rallied 105% from its March low just a few weeks earlier — so the reversal isn't a slow deflation, it's a fast one.",
						"The trigger was a piece of software, not a piece of hardware: Kimi K3, an open-weight model released July 16 by Moonshot AI, one of China's \"Six AI Tigers.\" Kimi K3 has 2.8 trillion parameters — the first open-weight model to reach that scale, nearly triple its predecessor — a million-token context window, and a mixture-of-experts design that routes each request through just 16 of 896 available expert subnetworks to keep inference costs down. It trails Claude Fable 5 and GPT-5.6 on most benchmarks, but beats Claude Opus 4.8 and GPT-5.5, at a price of $3 per million input tokens and $15 per million output tokens. Full weights land publicly on July 27.",
						"Why would a chatbot rattle a chip index that just posted its best quarter in years? Because the entire trillion-dollar hardware buildout — TSMC's $265 billion Arizona bet included — is a wager that frontier AI will keep demanding exponentially more compute, indefinitely. A capable open-weight model built more efficiently, out of China, at a fraction of the assumed cost, is the specific scenario that wager is exposed to. If good-enough intelligence gets cheap enough, the argument goes, maybe the industry doesn't need quite so much silicon after all.",
						"It's worth being skeptical of that logic, though, and not just because it triggered a $3.3 trillion round-trip in three weeks. Cheaper, more efficient inference has historically increased total compute consumption rather than shrinking it — more people running more queries more often, a semiconductor-flavored version of the Jevons paradox that showed up when GPT costs fell in 2024 and usage exploded rather than contracted. There's already a market signal pointing that direction: General Compute, an inference-focused cloud startup, just closed a $400 million loan from Upper90 using its inference chips as loan collateral — reportedly the first deal of its kind. That's a bet that inference capacity specifically, not training capacity, is the durable, financeable asset — a more precise read on where AI compute demand is actually heading than \"more of everything, forever.\"",
						"The chip industry's fundamentals didn't get worse this week. What got clearer is that the market pricing those fundamentals is still trying to figure out whether it's underwriting a training arms race, an inference boom, or both — and a single well-timed model release from Beijing was enough to make it flinch. That's the part worth watching over the next earnings cycle, not the headline number on any one company's income statement.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/1T6En1BbwlrR6PCNHAW2Tl6q0xfN62PYWCGoGQa36bgY/edit",
						"TSMC Q2 2026 earnings release: https://investor.tsmc.com/english/encrypt/files/encrypt_file/reports/2026-07/a80d7933be643644081584087731f73b22ea5a2c/2Q26%20EarningsRelease.pdf",
						"Yahoo Finance coverage of TSMC's Q2 results and Arizona investment: https://finance.yahoo.com/technology/articles/tsmc-adds-100bn-us-arizona-081156641.html",
						"The Next Web coverage of Samsung's preliminary Q2 2026 profit: https://thenextweb.com/news/samsung-q2-2026-operating-profit-record-ai-memory",
						"Yahoo Finance coverage of ASML's Q2 2026 guidance raise: https://finance.yahoo.com/technology/articles/asml-raises-2026-guidance-second-114432617.html",
						"Tom's Hardware coverage of Intel and ASML's High-NA EUV milestone: https://www.tomshardware.com/tech-industry/semiconductors/intel-becomes-the-first-company-to-ship-high-volume-logic-chips-made-with-asmls-high-na-euv-select-panther-lake-layers-on-18a-are-now-dual-qualified-for-0-55-na-scanners",
						"MarketWatch coverage of the SOX bear-market move: https://www.marketwatch.com/story/chip-stocks-enter-bear-market-territory-a-bofa-analyst-says-not-to-panic-4d34df17",
						"Tom's Hardware coverage of Moonshot Kimi K3: https://www.tomshardware.com/tech-industry/artificial-intelligence/moonshot-releases-2-8-trillion-parameter-kimi-k3",
						"TechCrunch coverage of General Compute's $400 million inference-chip-backed loan: https://techcrunch.com/2026/07/17/why-the-first-gpu-financiers-are-turning-to-inference-chips-in-a-400-million-deal/",
					},
				},
			},
		},
		{
			Title:   "Apple Says OpenAI Turned Job Interviews Into a Trade Secrets Pipeline",
			Slug:    "apple-openai-trade-secret-lawsuit-io-products-2026",
			Date:    "July 18, 2026",
			Tag:     "Legal",
			Summary: "A 41-page lawsuit accuses OpenAI's hardware chief of asking Apple engineers to bring prototypes to job interviews, a former Apple engineer of quietly reading company files after he'd already left, and Jony Ive's design studio of borrowing a metal-finishing trick it never had permission to use.",
			Related: []Link{
				{
					Title: "Nvidia's Roadmap Just Hit the Reticle Limit",
					Slug:  "nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"On July 10, Apple filed suit against OpenAI in the U.S. District Court for the Northern District of California. The case, Apple Inc. v. Liu (5:26-cv-07078), reads less like a routine corporate dispute and more like a workplace thriller. Apple's lawyers argue the theft wasn't the work of one rogue employee cashing out on the way to a new job — it was happening, in the company's words, \"at every level, from members of its Technical Staff to its Chief Hardware Officer, and in coordination with business partners.\"",
						"That Chief Hardware Officer is Tang Tan, a former Apple vice president who spent years running product design for the iPhone and Apple Watch before joining OpenAI to lead its push into consumer hardware. Apple's complaint alleges Tan didn't just recruit Apple engineers — he allegedly turned the interview process itself into an extraction point. According to the filing, Tan directed job candidates who were still employed at Apple to bring \"actual parts,\" CAD files, and prototypes to their interviews for what the complaint calls \"show and tell\" sessions, and used Apple's own internal project code names while doing the recruiting. One candidate, the filing says, was surprised to learn Apple hardware was even allowed to leave the building.",
						"The second thread involves Chang Liu, a former Apple senior systems electrical engineer who spent eight years at the company before moving to OpenAI. Apple alleges Liu exploited an authentication flaw to reach Apple's internal network storage through a colleague's work computer after he'd left the company, texting \"LOL, I found out I can access the [network storage]\" and later noting \"I still have another computer\" — which Apple's lawyers read as a plan to keep the access going. The complaint also says Liu never returned his Apple-issued laptop and used it to download confidential technical documents, including engineering presentations and specifications tied to unannounced products.",
						"The third and strangest thread points at io Products, the hardware design studio OpenAI acquired for $6.5 billion in 2025, co-founded by former Apple design chief Jony Ive. Apple alleges io used a \"secret, proprietary industrial design technique\" related to metal finishing by misleading one of Apple's own manufacturing partners into believing OpenAI had Apple's permission to use it — and that OpenAI approached suppliers using confidential Apple terminology it shouldn't have had. If true, it suggests the leakage wasn't limited to code or blueprints; it reached into the physical supply chain that actually builds Apple's products.",
						"Apple says none of this came out of nowhere. According to the complaint, the company first raised its concerns directly with OpenAI back in February 2026 and didn't get a response it considered adequate, which is part of why it's now suing rather than negotiating privately. The scale Apple cites is also striking context: more than 400 former Apple employees are now working at OpenAI, a flow of talent that has made Cupertino noticeably nervous as the AI industry's hardware ambitions have grown.",
						"OpenAI has pushed back, saying it takes the lawsuit seriously but believes the allegations lack merit. The company's position leans on a real legal distinction: trade secret law generally protects specific proprietary information, not the general skills and knowledge an employee carries in their head from one job to the next. California in particular has strong protections for worker mobility and a near-total ban on non-compete agreements, which is part of why Silicon Valley job-hopping is normally treated as healthy competition rather than theft.",
						"What makes this case worth watching isn't just the gossip-column details, though there are plenty of those. It's what it signals about where the AI industry's real bottleneck has shifted. For the past few years, the story was about who had the best model. Increasingly, it's about who can actually build a device people want to hold — and that requires exactly the kind of manufacturing know-how, supplier relationships, and industrial design instincts that companies like Apple have spent decades protecting. If Apple wins, it could make the next generation of AI hardware startups think twice about how aggressively they recruit from incumbents. If OpenAI wins, or if the case settles quietly, it will reinforce that in California, talent — and everything it remembers — is fair game.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/13cABBIhzEUCUabPvpJzi9-U2p-A-HzDH_bnwnrcbj34/edit",
						"CourtListener docket for Apple Inc. v. Liu, 5:26-cv-07078: https://www.courtlistener.com/docket/73602437/apple-inc-v-liu/",
						"TechCrunch, Apple sues OpenAI over alleged trade secret theft: https://techcrunch.com/2026/07/10/apple-sues-openai-over-alleged-trade-secret-theft/",
						"TechCrunch, the wildest allegations in Apple's trade secrets lawsuit against OpenAI: https://techcrunch.com/2026/07/13/the-wildest-allegations-in-apples-trade-secrets-lawsuit-against-openai/",
					},
				},
			},
		},
	}, posts...)
}
