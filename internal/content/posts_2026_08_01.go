package content

func init() {
	posts = append([]Post{
		{
			Title:   "DeepSeek Didn't Build a New Model to Beat Its Old One. It Just Retrained the Cheap One.",
			Slug:    "deepseek-v4-flash-0731-post-training-open-weight-economics-2026",
			Date:    "August 1, 2026",
			Tag:     "Models",
			Summary: "V4-Flash-0731 now outscores DeepSeek's own flagship Pro model on every published agent benchmark, at roughly a third of the price - without a single new parameter.",
			Related: []Link{
				{
					Title: "The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say.",
					Slug:  "white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026",
				},
				{
					Title: "Google Just Shipped Three New Gemini Models. The One Everyone Actually Wants Still Isn't Ready.",
					Slug:  "gemini-3-5-pro-third-delay-flash-stopgap-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`On July 31, DeepSeek quietly pushed its V4-Flash model out of preview and into general release under the name DeepSeek-V4-Flash-0731. That sentence alone wouldn't normally be news. Point releases happen every week in this industry. What makes this one worth stopping for is what DeepSeek didn't change: the architecture is identical to the version previewed back in the spring. Same 284-billion-parameter mixture-of-experts design, same 13 billion parameters actually activated per token, same 1-million-token context window. Nothing about the model got bigger. DeepSeek just retrained it, and that alone was enough to leapfrog its own flagship.`,
					},
				},
				{
					Heading: "The benchmark jump was not subtle",
					Paragraphs: []string{
						`According to benchmark data DeepSeek published on Hugging Face and independently verified by Artificial Analysis, V4-Flash-0731 now beats DeepSeek-V4-Pro - the company's larger, more expensive model - on all nine agentic and coding benchmarks the company has released numbers for. The gains aren't marginal. On Terminal-Bench 2.1, a test of an AI system's ability to actually operate a command line and complete real tasks, the new Flash release scores 82.7, up from 61.8 in the spring preview and ahead of Pro's 72.1. On DeepSWE, a software-engineering benchmark, the jump is even sharper: from 7.3 in preview to 54.4 now. NL2Repo, which tests turning natural-language requests into working code changes across a repository, climbed from 39.4 to 54.2. Cybergym, a security-flavored benchmark, went from 38.7 to 76.7.`,
						`DeepSeek's own release notes attribute all of this to re-post-training - adjusting how the model is fine-tuned and aligned after pretraining - rather than a new base model. That distinction matters more than it might sound. Pretraining a frontier-scale model from scratch costs tens of millions of dollars and months of compute. Post-training is dramatically cheaper and faster, which means DeepSeek effectively bought a generational capability jump without the generational price tag. Artificial Analysis's own composite scoring backs up the scale of the jump: V4-Flash-0731 now posts a 50 on the firm's Intelligence Index, good for third among the 101 models it tracks in that class, against a median of just 25 for comparable open-weight systems.`,
					},
				},
				{
					Heading: "Cheap enough to change the argument",
					Paragraphs: []string{
						`The pricing is where the story turns from impressive to disruptive. DeepSeek is charging $0.14 per million input tokens and $0.28 per million output tokens for the new Flash release - about a third of what it charges for Pro's $0.87 output rate, according to Artificial Analysis's pricing data. Cached input tokens run just $0.0028 per million, a discount Artificial Analysis flags as roughly 98% off the fresh-token rate and the cheapest cache pricing it tracks across the board. For comparison, Google's newly launched Gemini 3.6 Flash - itself pitched as a cost-cutting move - prices output tokens at $7.50 per million, more than 26 times DeepSeek's rate, though it's worth noting the two models weren't built for identical workloads and Google's pricing includes first-party enterprise support and tooling DeepSeek doesn't offer directly.`,
						`The weights are MIT-licensed and sitting ungated on Hugging Face, meaning any company willing to run its own infrastructure can self-host the model instead of paying API fees at all. DeepSeek's documentation puts full-precision self-hosting at roughly a four-GPU GB300 node, or as little as 110GB of memory if you're willing to run it at 3-bit quantization - a meaningfully lower bar than the multi-node clusters flagship-scale models typically demand.`,
					},
				},
				{
					Heading: "The moat gets shallower",
					Paragraphs: []string{
						`Step back from the benchmark table and the real headline is about incentives. Every major lab is currently racing to convince enterprises that bigger, pricier, more heavily guarded models are worth paying a premium for. DeepSeek just demonstrated, in public, that a smaller open-weight model can out-benchmark its own more expensive sibling purely through better post-training - the part of the pipeline that's cheapest to iterate on and hardest to keep proprietary. If that trick generalizes across labs, the moat around frontier pricing gets a lot shallower, and fast.`,
						`It's also the latest data point in a pattern regular readers have seen us cover before: Chinese open-weight labs iterating in public at a pace and price point that's forcing everyone else to justify their margins. Washington has already accused one Chinese lab, Moonshot AI, of leaning on Claude's outputs to train its Kimi K3 model - a claim researchers pushed back on for lacking a plausible timeline, as we reported last week. Whatever the truth of that specific dispute, the broader trend it sits inside is real: the cheapest agentic AI money can buy right now isn't coming from Silicon Valley, and it's getting harder to explain to a CFO why that shouldn't matter.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"DeepSeek, DeepSeek-V4-Flash-0731 model card on Hugging Face: https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-0731",
						"Artificial Analysis, DeepSeek V4 Flash 0731 scores 50 on the Artificial Analysis Intelligence Index: https://artificialanalysis.ai/articles/deepseek-v4-flash-0731-scores-50-on-the-artificial-analysis-intelligence-index-10-points-above-previous-deepseek-v4-flash",
						"Artificial Analysis, DeepSeek model evaluation and pricing page: https://artificialanalysis.ai/models/deepseek-v4-flash",
						"MarkTechPost, DeepSeek Upgrades DeepSeek-V4-Flash-0731 with Major Agentic and Coding Gains: https://www.marktechpost.com/2026/07/31/deepseek-upgrades-deepseek-v4-flash-0731-with-major-agentic-and-coding-gains/",
						"Google, Introducing Gemini 3.6 Flash, 3.5 Flash-Lite, and 3.5 Flash Cyber: https://blog.google/innovation-and-ai/models-and-research/gemini-models/gemini-3-6-flash-3-5-flash-lite-3-5-flash-cyber/",
					},
				},
			},
		},
	}, posts...)
}
