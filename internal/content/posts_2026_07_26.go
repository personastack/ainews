package content

func init() {
	posts = append([]Post{
		{
			Title:   "Jensen Huang's First Tweet Ever Wasn't About Chips. It Was a Warning to Washington.",
			Slug:    "nvidia-jensen-huang-open-weights-letter-distillation-2026",
			Date:    "July 26, 2026",
			Tag:     "Policy",
			Summary: "Nvidia and more than 20 other companies signed a letter defending open-weight AI models — including the exact technique the White House accused China's Kimi K3 of stealing just two days earlier.",
			Related: []Link{
				{
					Title: "The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say.",
					Slug:  "white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026",
				},
				{
					Title: "Nvidia's Roadmap Just Hit the Reticle Limit",
					Slug:  "nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`Jensen Huang joined X in June and said nothing for weeks. When Nvidia's CEO finally posted, on July 24, he didn't share a chip roadmap or an earnings tease. He shared a letter.`,
						`"For my first post, I'm sharing a letter NVIDIA signed on why open models matter," Huang wrote. "AI will transform every industry, power every company, and be built by every country." Open models, he added, "strengthen safety and cybersecurity, accelerate innovation and diffusion, and enable sovereignty."`,
						`The letter itself, titled "Open Weights and American AI Leadership," is a policy document, not a product announcement. It's signed by more than 20 companies and organizations spanning the AI industry's usual rivalries: Nvidia, Meta, Microsoft, Palantir, IBM, Hugging Face, Andreessen Horowitz, Mozilla, Perplexity, Dell, CrowdStrike, ServiceNow, Y Combinator, and the Linux Foundation among them. Its core argument is that Washington should not restrict open-weight AI models — the kind anyone can download, inspect, and modify — as it weighs new AI trade and export policy.`,
					},
				},
				{
					Heading: "The case the letter makes",
					Paragraphs: []string{
						`The signatories lay out three main arguments. Open models widen economic access, letting startups and universities build on frontier-adjacent technology without paying a closed lab's API bill. They improve security, the letter argues, because broad researcher access means vulnerabilities get found and patched rather than hidden inside "single points of failure" at a handful of closed labs. And they support what the letter calls sovereignty — the idea that no country, including the United States, should have its AI industry dependent on a small number of proprietary systems it doesn't control.`,
						`The letter is blunt about what a closed-only future would look like: "A world dominated by two or three closed frontier labs...is a world with fewer buyers." And on the country's overall competitive position: "Our AI leadership will be judged not by one frontier AI model, but by whether the United States builds a strong, open ecosystem that diffuses into every sector."`,
						`Huang's framing leans on history. The letter compares today's open-weight debate to open-source software in the 1980s and 90s — code that skeptics dismissed as amateur or unsafe, and that now quietly runs most of the internet. The implied warning: don't repeat that mistake with AI models.`,
					},
				},
				{
					Heading: "The distillation problem",
					Paragraphs: []string{
						`Buried in the letter's more technical section is the line most likely to raise eyebrows in Washington. The signatories describe distillation — training a smaller model to mimic a larger one's outputs — as "a legitimate and widely used technique for model improvement," and argue it shouldn't be conflated with unlawful misappropriation of a closed model's weights or training data.`,
						`That's not an abstract stance. Two days before Huang's post, White House Office of Science and Technology Policy director Michael Kratsios had publicly accused Chinese startup Moonshot AI of doing exactly that: distilling Anthropic's Fable 5 model to help build its Kimi K3 system, launched July 16. Independent researchers pushed back on the timeline — Fable 5 had only been publicly accessible again since July 1, after a two-week export-control suspension, leaving what they called an implausibly short window for the kind of distillation campaign Kratsios described. Nvidia's letter doesn't name Kimi K3 or Moonshot. It doesn't have to. It arrives right as the administration is reportedly weighing trade restrictions on freely downloadable AI weights from international competitors, and it stakes out the industry's position before those rules get written: distillation is normal, not theft, and restricting it broadly would hurt American open-weight developers more than it hurts anyone in Beijing.`,
					},
				},
				{
					Heading: "Huang's China problem",
					Paragraphs: []string{
						`The letter also lands amid a rockier week for Huang personally. In comments to Axios shortly before his X debut, he argued that American companies should be free to use Chinese AI models — remarks that put him at odds with Treasury Secretary Scott Bessent, who has warned of possible sanctions tied to Chinese AI adoption. Huang has separately described Chinese models as strong performers whose popularity, if anything, expands demand for the chips that run them, regardless of which country's lab trained the weights.`,
						`That's worth sitting with, because it points to something the letter doesn't say outright: Nvidia's business model doesn't actually care whether the winning models are open or closed, American or Chinese. It sells the compute either way. What Nvidia and its co-signatories are really lobbying against is a world with fewer AI companies buying fewer chips — and an open-weight ecosystem, almost by definition, has more of both. The safety and sovereignty arguments may be genuine. They're also, conveniently, the arguments that keep the customer base as wide as possible.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Nvidia CEO Jensen Huang makes X debut backing open-weight AI — Yahoo Tech: https://tech.yahoo.com/ai/articles/nvidia-ceo-jensen-huang-makes-155300228.html",
						"Jensen Huang just used his first ever X post to warn the AI industry — Fortune: https://fortune.com/2026/07/24/jensen-huang-open-source-letter-nvidia-kimi/",
						"NVIDIA CEO Jensen Huang Backs Open Models In First Post On X — OfficeChai: https://officechai.com/ai/nvidia-ceo-jensen-huang-backs-open-models-in-first-post-on-x/",
						"NVIDIA CEO Jensen Huang Uses X Debut To Push Open-Weight AI Models — HotHardware: https://hothardware.com/news/nvidia-ceo-jensen-huang-x-debut-push-open-weight-ai-models",
					},
				},
			},
		},
	}, posts...)
}
