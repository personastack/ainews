package content

func init() {
	posts = append([]Post{
		{
			Title:   "China Just Trained a Frontier Model on 50,000 of Its Own Chips. The Export Controls Were Supposed to Make That Impossible.",
			Slug:    "china-domestic-ai-chips-longcat-export-controls-2026",
			Date:    "July 5, 2026",
			Tag:     "Hardware",
			Summary: "Meituan's LongCat-2.0 is a 1.6-trillion-parameter stress test of the U.S. chip-denial strategy: frontier-scale AI trained and deployed on domestic Chinese accelerators.",
			Related: []Link{
				{
					Title: "To Beat Nvidia, Qualcomm Didn't Buy a Faster Chip — It Bought a Compiler",
					Slug:  "qualcomm-modular-cuda-moat-compiler-nvidia-2026",
				},
				{
					Title: "OpenAI Shipped Its Most Powerful Model. Only 20 Companies — All Government-Approved — Can Use It.",
					Slug:  "openai-gpt-5-6-sol-government-gated-frontier-release-2026",
				},
				{
					Title: "The Chip Stopped Being the Bottleneck — Now It's Power and Memory",
					Slug:  "ai-real-bottleneck-power-memory-not-chips-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"For four years, the load-bearing assumption behind U.S. AI policy has been simple: keep the best chips out of China, and you keep China a step behind the frontier. Deny the H100s and Blackwells, and the largest models stay a Western privilege by default. It was a bet on hardware as a chokepoint.",
						"Over the past week, that bet got its most serious challenge yet — and it did not come from a chip company. It came from a food-delivery giant.",
						"Meituan, the Chinese super-app better known for getting dumplings to your door, open-sourced LongCat-2.0: a 1.6-trillion-parameter Mixture-of-Experts model that the company says was pre-trained and deployed end-to-end on a cluster of more than 50,000 domestic Chinese accelerators — Ascend-class silicon, no restricted U.S. hardware anywhere in the stack. The weights are already on Hugging Face under the meituan-longcat organization, free for anyone to download. On SWE-bench Pro, the company's model card reports a 59.5, squarely in frontier company.",
						"If that claim holds up, it is the first time a model of this scale has completed the full round trip — training from scratch through large-scale inference — without touching an Nvidia or Google chip. And that is the whole story.",
					},
				},
				{
					Heading: "What's actually under the hood",
					Paragraphs: []string{
						"Strip away the geopolitics for a second, because the engineering is the interesting part. LongCat-2.0 is big on paper — 1.6 trillion total parameters — but aggressively sparse in practice: only about 48 billion parameters activate for any given token, thanks to the Mixture-of-Experts design that routes each input to a small slice of the network. That is the same efficiency trick that has quietly become the industry default, and it is what makes a trillion-parameter model economical to actually run.",
						"Meituan says the training run chewed through more than 35 trillion tokens, including long stretches at roughly one-million-token context lengths, on AI ASIC superpods stitched together with communication and parallelism work that mattered as much as the chips themselves. Training a model this size is not bottlenecked by any single chip's peak math; it is bottlenecked by how well tens of thousands of chips can talk to each other without stalling. The hard part China had to solve was not just the transistor. It was the plumbing — the interconnect, the collective-communication libraries, the scheduling — the exact software layer that has always been Nvidia's deepest moat.",
						"Which should sound familiar. Just yesterday we wrote about Qualcomm buying its way toward Nvidia's throne not with a faster chip but with a compiler. Same lesson, opposite side of the Pacific: in 2026, the thing that separates usable silicon from expensive sand is the stack that sits on top of it. LongCat-2.0 is a claim that China has now built enough of that stack to reach the frontier on its own hardware.",
					},
				},
				{
					Heading: "Not a one-off",
					Paragraphs: []string{
						"The temptation is to treat this as a stunt — a single splashy release engineered for headlines. The evidence points the other way. LongCat lands on top of a visible trend line.",
						"DeepSeek's V4-Pro, another 1.6-trillion-parameter MoE model, was reported to have trained or run substantial parts of its stack on Huawei Ascend hardware. Zhipu AI has also claimed major-model training on Ascend processors. And the demand side has caught up fast: ByteDance, Tencent, and Alibaba have all been scrambling for Huawei's Ascend accelerators, while Cambricon plans to ship half a million domestic AI chips this year.",
						"Beijing, for its part, has started turning domestic AI hardware into a procurement category. Recent approved-supplier lists and Anke certifications put Chinese accelerators into government and state-linked purchasing channels while Nvidia and AMD remain outside the domestic-preference frame. Domestic silicon already made up roughly 41% of China's AI-chip market in 2025, about half of that Huawei. LongCat-2.0 is not the beginning of chip substitution. It is the moment substitution reached the top of the model stack.",
					},
				},
				{
					Heading: "The honest caveats",
					Paragraphs: []string{
						"Good news deserves skepticism, and so does this. Nearly every headline number here — the 50,000-card cluster, the 35 trillion tokens, the end-to-end no-U.S.-hardware claim — is self-reported by Meituan and has not been independently audited. Company-stated training details are marketing until someone reproduces them. The public model card says the full training run and large-scale deployment were built entirely on AI ASIC superpods; outside reporting says Meituan described a 50,000-card domestic cluster. Those are important claims, but they are still claims.",
						"It is also true that matching a benchmark score is not the same as matching efficiency. Training on domestic chips can still cost far more power, time, and money per unit of capability than the Nvidia path, and \"we can do it at all\" is a different claim from \"we can do it economically at scale.\" Export controls were arguably always more about raising China's costs than about hard denial, and on that narrower goal they may still be working.",
						"But costs come down. Chokepoints, once routed around, tend not to re-close. The strategic question Washington now has to answer is not whether China can train frontier models on its own chips — LongCat-2.0 is a fairly loud yes. It is whether a policy built to buy time has any time left to buy.",
					},
				},
				{
					Heading: "The part worth sitting with",
					Paragraphs: []string{
						"There is a deeper irony here for the open-source debate. The most powerful Western models of 2026 are drifting behind government gates and approved-buyer lists. Meanwhile, the model that just punched a hole in America's containment strategy was posted for the world to download, for free, by a company that sells lunch.",
						"The wall was built out of chips. It turns out the ladder was made of software — and someone left it open-source.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Meituan LongCat-2.0 model card on Hugging Face: https://huggingface.co/meituan-longcat/LongCat-2.0",
						"South China Morning Post, June 30, 2026, \"China claims biggest AI model trained on local chips, as Meituan releases LongCat-2.0\": https://www.scmp.com/tech/tech-trends/article/3358854/china-debuts-biggest-ai-model-trained-local-chips-meituan-releases-longcat-20",
						"Tom's Hardware, May 27, 2026, \"China adds homegrown AI chips to 'secure and reliable' procurement list for the first time\": https://www.tomshardware.com/tech-industry/semiconductors/china-certifies-nine-domestic-ai-chips-for-government-procurement",
						"Author article handoff and archive doc: https://docs.google.com/document/d/1iDy9NcO3M72ith5Rmsf7waEojy5TxWrZEwwgCU6UIVQ/edit",
					},
				},
			},
		},
	}, posts...)
}
