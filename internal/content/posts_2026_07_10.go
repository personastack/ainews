package content

func init() {
	posts = append([]Post{
		{
			Title:   "Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million.",
			Slug:    "openai-gpt-5-6-general-availability-government-gate-precedent-2026",
			Date:    "July 10, 2026",
			Tag:     "Policy",
			Summary: "OpenAI's GPT-5.6 spent its first weeks behind the first U.S. government security review of a frontier model. On Thursday the gate lifted and the model went on sale in three tiers. The precedent it set didn't lift with it.",
			Related: []Link{
				{
					Title: "OpenAI Shipped Its Most Powerful Model. Only 20 Companies — All Government-Approved — Can Use It.",
					Slug:  "openai-gpt-5-6-sol-government-gated-frontier-release-2026",
				},
				{
					Title: "Washington Wrote the Rulebook for Frontier AI — And the First Lab It Touched Is Suing",
					Slug:  "us-ai-national-security-executive-order-anthropic-lawsuit-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"For about twelve days this summer, the most capable model OpenAI had ever built existed in a strange limbo. GPT-5.6 was finished — previewed on June 26, benchmarked, and handed to roughly twenty partner organizations — but you couldn't have it. At the request of the U.S. government, OpenAI shipped its new flagship, Sol, only to a small circle of trusted customers while the Department of Commerce ran a security review. On Thursday, July 9, that review ended. GPT-5.6 went generally available across ChatGPT, Codex, and the OpenAI API, with the global rollout finishing over the following day.",
						"If you only read the release notes, this looks like a routine launch. It is not. It is the first time a leading American AI lab has held a finished model off the open market specifically because Washington asked it to — and then released it on Washington's clearance. The gate is gone. The template it leaves behind is what matters.",
					},
				},
				{
					Heading: "Why this model, and not the last one",
					Paragraphs: []string{
						"The trigger was narrow and specific: cybersecurity. OpenAI's own framing describes Sol as its most capable model yet for long-horizon security work — the kind of patient, multi-step reasoning that vulnerability research and exploitation actually require. A model that can find and chain software flaws is dual-use in the most literal sense: the same capability that helps a defender audit their code helps an attacker breach it. That is the threshold that moved GPT-5.6 from \"ship it\" to \"let's have someone look first.\"",
						"The someone was the Commerce Department's Center for AI Standards and Innovation, or CAISI — the body that, under a previous name, was the U.S. AI Safety Institute. According to reporting around the clearance, CAISI ran additional testing and OpenAI sent technical staff to Washington to work through questions in person before the department signed off on a broad release. Call it what it is: a pre-market safety review, conducted by the government, on a commercial product, triggered by a capability the company itself flagged.",
					},
				},
				{
					Heading: "A dollar a million",
					Paragraphs: []string{
						"The other half of Thursday's news is the price sheet, and it tells a different story about where the frontier is going. GPT-5.6 didn't ship as one model. It shipped as three. Sol, the flagship, costs $5 per million input tokens and $30 per million output. Terra, the mid-tier, is half that — $2.50 and $15 — and is pitched as the everyday workhorse. Luna, the floor, runs $1 input and $6 output, built for high-volume pipelines where latency and cost matter more than the last few points of benchmark.",
						"Read those two facts together and you get the shape of the 2026 frontier. The most powerful tier — the one with the cyber capability that drew the government's attention — is priced as a premium, agentic product. The bottom tier is priced as a commodity, cheap enough to sit inside a background job you never think about. The same family of weights now spans \"national-security-review-worthy\" and \"a dollar a million.\" That is a lot of range to hold under one version number.",
					},
				},
				{
					Heading: "The gate becomes a norm",
					Paragraphs: []string{
						"Here is the part worth sitting with. The review was fast — days, not months — and it ended in a yes. From the outside, it looks less like a regulatory wall and more like a checkpoint: capability crosses a line, the government takes a look, the model ships. No statute compelled it. No public rule defined the threshold. It happened through a voluntary framework and a working relationship, and it worked well enough that nobody had to fight about it in public.",
						"That is exactly what makes it consequential. A one-time courtesy is easy to walk back. A precedent that both sides found tolerable is not. The next frontier model that crosses a capability threshold — cyber, bio, autonomy — now has a path already worn into the ground: preview, gate, review, clear. We spent 2026 arguing about whether AI would be governed by hard law or left to self-govern. GPT-5.6 suggests the real answer, at least in the United States, is a third thing — discretionary, capability-triggered, and negotiated case by case behind closed doors.",
						"The gate lifted on Thursday. Whether it ever fully closes again is the question the price tag can't answer.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Sources: OpenAI, \"Previewing GPT-5.6 Sol\"; OpenAI Help Center GPT-5.6 article; Simon Willison, \"The new GPT-5.6 family: Luna, Terra, Sol\" (July 9, 2026); Dataconomy, \"OpenAI Launches GPT-5.6 With Sol, Terra, And Luna\" (July 10, 2026); FinanceFeeds, Quartz, and TechTimes reporting on Department of Commerce CAISI clearance for broad GPT-5.6 release (July 9, 2026). Pricing per 1M tokens: Sol $5 input/$30 output, Terra $2.50 input/$15 output, Luna $1 input/$6 output.",
						"Author article handoff: https://docs.google.com/document/d/1x8X0RjI4yQiIhjLsRndobNro8upzLFxzIX2zYU3YnNY/edit",
					},
				},
			},
		},
	}, posts...)
}
