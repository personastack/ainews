package content

func init() {
	posts = append([]Post{
		{
			Title:   "Alibaba's Qwen3.8-Max Beats GPT-5.6 and Claude on Key Benchmarks — And It's Going Open Weight",
			Slug:    "qwen3-8-max-open-weight-benchmarks-gpt-5-6-claude-2026",
			Date:    "August 6, 2026",
			Tag:     "Models",
			Summary: "Alibaba's 2.4-trillion-parameter Qwen3.8-Max beats GPT-5.6 Sol on PaperBench and Claude Fable 5 on IFBench, with open weights promised within days.",
			Related: []Link{
				{
					Title: "Alibaba's AI Offensive: How Qwen3.7-Max and a New Skills Portal Challenge Western Cloud Giants",
					Slug:  "alibaba-cloud-agentic-ai-offensive-qwen3-7-max",
				},
				{
					Title: "DeepSeek and Moonshot Are Racing to IPO. Beijing Just Showed Up as the Investor With No Lock-Up.",
					Slug:  "deepseek-moonshot-china-ai-ipo-funding-state-investment-2026",
				},
				{
					Title: "Fable 5 Was Built for Safer Access. Washington Shut It Down Anyway.",
					Slug:  "fable-5-mythos-5-export-control-shutdown-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`A 2.4-trillion-parameter model just out-scored OpenAI's flagship on reproducing scientific papers and trounced Anthropic's high-end model on following instructions. In a week or so, anyone will be able to download its weights and run it themselves.`,
						`Alibaba released Qwen3.8-Max on August 3, 2026, and it's the largest model the company has ever shipped: 2.4 trillion total parameters built as a mixture-of-experts system, with roughly 95 billion active on any given request. It handles text, images, and video in a single context window stretching to one million tokens. Access starts through Alibaba's QwenCloud API at $2.00 per million input tokens and $6.00 per million output tokens, with cached input priced at just $0.25 per million — a fraction of what OpenAI and Anthropic charge for their top-tier models.`,
					},
				},
				{
					Heading: "The benchmark split",
					Paragraphs: []string{
						`The pricing is competitive. The benchmark story is more interesting, because it isn't a clean sweep in either direction.`,
						`On Terminal-Bench 2.1, a test of a model's ability to complete real command-line tasks, Qwen3.8-Max scored 86.6, trailing GPT-5.6 Sol's 88.8. On MRCR v2, which measures long-context recall, it posted 92.9 against Sol's 93.8 — close, but still second. Those are the categories where OpenAI's flagship holds its ground.`,
						`Then the picture flips. On PaperBench, which tests whether a model can actually reproduce the results of a published research paper from scratch, Qwen3.8-Max scored 93.0 — ahead of GPT-5.6 Sol's 90.5. And on IFBench, a benchmark for precisely following complicated instructions, Qwen3.8-Max hit 82.8, well clear of Sol's 72.7 and miles ahead of Claude Fable 5's 63.5. Against its own predecessor, Qwen3.7-Max, the generational jump is stark on agentic coding work — FrontierSWE climbed from 40.7 to 73.5 — while reasoning benchmarks like GPQA barely moved (92.4 to 92.6), suggesting Alibaba spent this training cycle on usability and tool use rather than raw reasoning horsepower.`,
						`On Arena.AI's public leaderboard, Qwen3.8-Max now ranks as the highest-scoring Chinese text model available and sits second in the world on visual-analysis tasks, trailing only Claude Fable 5.`,
					},
				},
				{
					Heading: "The export-control backdrop",
					Paragraphs: []string{
						`That comparison carries some irony. Fable 5 spent nearly three weeks locked out of the rest of the world this summer: the U.S. government issued an export control order on June 12 suspending all foreign access to Fable 5 and its cybersecurity-focused sibling Mythos 5, citing a discovered jailbreak method as a national security concern. Anthropic restored global access on July 1 after the order was lifted. Qwen3.8-Max wasn't built to answer that episode directly, but the timing is a reminder of how tightly the current AI race is bound up with export policy, not just model quality — a frontier-class model briefly became something Washington treated like a controlled export, and a few weeks later a Chinese lab is benchmarking directly against it in public.`,
					},
				},
				{
					Heading: "The open-weights move",
					Paragraphs: []string{
						`The bigger structural story here might be the open-weights decision. Alibaba says full open weights for Qwen3.8-Max are coming "next week," alongside a smaller Qwen3.8-27B variant meant for teams that can't run a 2.4-trillion-parameter model on their own infrastructure — and to be clear, almost nobody can; the full checkpoint requires multi-node datacenter deployment, while the 27B model is the realistic on-premise option for most organizations. What makes this notable is that it's the first time Alibaba has open-sourced a model at its Max tier. Previous Qwen-Max releases stayed behind an API, similar to how OpenAI and Anthropic keep their top models closed. Pairing genuine frontier-class benchmark results with an open license is a different competitive move than undercutting on price, which is the move Alibaba and other Chinese labs have leaned on for most of the year.`,
					},
				},
				{
					Heading: "Alibaba's broader pattern",
					Paragraphs: []string{
						`It also continues a pattern for Alibaba specifically. Back in May, Qwen3.7-Max launched alongside an MCP-compatible skills portal aimed squarely at enterprise agent development — a direct challenge to AWS, Azure, and Google Cloud on their own turf. Qwen3.8-Max extends that ambition from tooling into raw model capability, arriving at a moment when Chinese-built models have already been pulling a growing share of enterprise token usage on routing platforms like OpenRouter throughout 2026.`,
					},
				},
				{
					Heading: "What to watch",
					Paragraphs: []string{
						`None of this means Qwen3.8-Max is simply "better" than GPT-5.6 Sol or Claude Fable 5 — the benchmark split is genuinely mixed, and headline leaderboard rankings don't capture reliability, safety tooling, or enterprise support, which is where the closed frontier labs still make their case. But a model that beats OpenAI's flagship at reproducing real research and beats Anthropic's top model at following instructions, priced well below both and headed toward open weights within days, is not a story to file under "China catches up eventually." That gap, in at least these two dimensions, is already closed. The question worth watching now is whether Alibaba can sustain the pace of point releases it's kept up all year, and whether OpenAI and Anthropic respond by opening up their own weights or by doubling down on the safety and reliability guarantees that closed models can still promise more credibly.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Qwen, Qwen3.8-Max: A New Bar for Coding and Cowork: https://qwen.ai/blog?id=qwen3.8",
						"MarkTechPost, Alibaba Qwen Releases Qwen3.8-Max: https://www.marktechpost.com/2026/08/03/alibaba-qwen-releases-qwen3-8-max/",
						"The Decoder, Alibaba's open-weight Qwen3.8-Max takes on long-horizon AI tasks with 2.4 trillion parameters: https://the-decoder.com/alibabas-open-weight-qwen3-8-max-takes-on-long-horizon-ai-tasks-with-2-4-trillion-parameters/",
						"Anthropic, Statement on the US government directive to suspend Fable 5 and Mythos 5 access: https://www.anthropic.com/news/fable-mythos-access",
						"Anthropic, Redeploying Claude Fable 5: https://www.anthropic.com/news/redeploying-fable-5",
					},
				},
			},
		},
	}, posts...)
}
