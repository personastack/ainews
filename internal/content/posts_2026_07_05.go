package content

func init() {
	posts = append([]Post{
		{
			Title:   "China Just Trained a Frontier Model on 50,000 of Its Own Chips. The Export Controls Were Supposed to Make That Impossible.",
			Slug:    "china-domestic-chips-longcat-frontier-model-export-controls",
			Date:    "July 5, 2026",
			Tag:     "Hardware/Policy",
			Summary: "Meituan's LongCat-2.0 — 1.6 trillion parameters, open weights, and reportedly not a single Nvidia GPU in the training run — is less a model launch than a stress test of America's entire chip-denial strategy.",
			Related: []Link{
				{
					Title: "To Beat Nvidia, Qualcomm Didn't Buy a Faster Chip — It Bought a Compiler",
					Slug:  "qualcomm-modular-cuda-moat-compiler-nvidia-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"For four years, the load-bearing assumption behind U.S. AI policy has been simple: keep the best chips out of China, and you keep China a step behind the frontier. Deny the H100s and Blackwells, and the largest models stay a Western privilege by default. It was a bet on hardware as a chokepoint.",
						"Over the past week, that bet got its most serious challenge yet — and it didn't come from a chip company. It came from a food-delivery giant.",
						"Meituan, the Chinese super-app better known for getting dumplings to your door, open-sourced LongCat-2.0: a 1.6-trillion-parameter Mixture-of-Experts model that the company says was pre-trained and deployed end-to-end on a cluster of more than 50,000 domestic Chinese accelerators — Ascend-class silicon, no restricted U.S. hardware anywhere in the stack. The weights are already on Hugging Face under the meituan-longcat organization, free to download. On SWE-bench Pro it posts a 59.5 — squarely in frontier company.",
						"If that claim holds up, it is the first time a model of this scale has completed the full round trip — training from scratch through large-scale inference — without touching an Nvidia or Google chip. And that is the whole story.",
					},
				},
				{
					Heading: "What's actually under the hood",
					Paragraphs: []string{
						"LongCat-2.0 is big on paper — 1.6T total params — but aggressively sparse: only ~48B parameters activate per token via the MoE design that routes each input to a small slice of the network. Same efficiency trick that became the industry default; it's what makes a trillion-param model economical to run.",
						"Meituan says the run chewed through 35T+ tokens, including long stretches at ~1M-token context, on a 50,000-card cluster stitched together with Huawei's HCCL communication library and in-house parallelism tweaks. That detail matters: a model this size isn't bottlenecked by any single chip's peak math but by how well tens of thousands of chips talk to each other without stalling. The hard part wasn't the transistor. It was the plumbing — interconnect, collective-communication libraries, scheduling — the exact software layer that's always been Nvidia's deepest moat.",
						"Which should sound familiar. Yesterday we wrote about Qualcomm buying toward Nvidia's throne not with a faster chip but with a compiler. Same lesson, opposite side of the Pacific: in 2026 the thing that separates usable silicon from expensive sand is the stack on top of it. LongCat-2.0 is a claim that China has built enough of that stack to reach the frontier on its own hardware.",
					},
				},
				{
					Heading: "Not a one-off",
					Paragraphs: []string{
						"Treating this as a stunt gets it wrong; it lands on a visible trend line. DeepSeek's V4-Pro (another 1.6T MoE) reportedly trained on Huawei Ascend 950 with no Nvidia parts. Zhipu AI earlier claimed to be first to train a major model entirely on Ascend. ByteDance, Tencent, Alibaba all scrambling for Ascend accelerators; Cambricon plans to ship 500,000 domestic AI chips this year. Beijing publishes an approved AI-hardware supplier list — Huawei and Cambricon on it; Nvidia and AMD not. Domestic silicon was ~41% of China's AI-chip market in 2025, about half Huawei. LongCat isn't the beginning of chip substitution; it's the moment substitution reached the top of the model stack.",
					},
				},
				{
					Heading: "The honest caveats",
					Paragraphs: []string{
						"Nearly every headline number — 50,000-card cluster, 35T tokens, end-to-end no-U.S.-hardware — is self-reported by Meituan and not independently audited; company-stated training details are marketing until reproduced. Matching a benchmark is not matching efficiency: domestic-chip training can cost far more power/time/money per unit of capability, and \"we can do it at all\" differs from \"economically at scale.\" Export controls were arguably always about raising costs more than hard denial, and on that narrower goal may still be working. But costs come down; chokepoints once routed around tend not to re-close. The question isn't whether China can train frontier models on its own chips — LongCat is a loud yes — it's whether a policy built to buy time has time left to buy.",
					},
				},
				{
					Heading: "The part worth sitting with",
					Paragraphs: []string{
						"Deeper irony for the open-source debate: the most powerful Western models of 2026 are drifting behind government gates and approved-buyer lists (newest OpenAI frontier models in limited preview for ~two dozen vetted orgs), while the model that just punched a hole in America's containment strategy was posted for the whole world to download, free, by a company that sells lunch. The wall was built out of chips. It turns out the ladder was made of software — and someone left it open-source.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"SiliconANGLE (June 30, 2026); South China Morning Post; CryptoBriefing; Quasa; AI Weekly; CSIS and CFR export-control analysis; Tom's Hardware (China approved-supplier list); company materials on Hugging Face (meituan-longcat). Meituan's training-run figures are company-reported and not independently verified.",
					},
				},
			},
		},
	}, posts...)
}
