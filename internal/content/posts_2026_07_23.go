package content

func init() {
	posts = append([]Post{
		{
			Title:   "Google Just Shipped Three New Gemini Models. The One Everyone Actually Wants Still Isn't Ready.",
			Slug:    "gemini-3-5-pro-third-delay-flash-stopgap-2026",
			Date:    "July 23, 2026",
			Tag:     "Models",
			Summary: "Gemini 3.5 Pro has missed three deadlines since Google I/O in May, and this week Google filled the gap with two smaller Flash models and a third built exclusively for cyber defense. The stopgaps are polished. What they reveal about the delay is not.",
			Related: []Link{
				{
					Title: "Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million.",
					Slug:  "openai-gpt-5-6-general-availability-government-gate-precedent-2026",
				},
				{
					Title: "The Rocket Company Ships a Coding Model - And the Benchmark Depends on Who's Grading",
					Slug:  "grok-4-5-spacexai-cursor-coding-benchmark-harness-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`For the third time since it was first promised at Google I/O, Gemini 3.5 Pro has missed its launch date. Google confirmed on July 21 that its next flagship reasoning model remains in partner testing, with no new release date attached - only a line in the company's own announcement: "Gemini 3.5 Pro is currently testing with partners and we plan to make it broadly available as soon as it's ready."`,
						"What Google shipped instead, that same day, was a trio of smaller Gemini models built to hold the line while Pro keeps slipping.",
					},
				},
				{
					Heading: "The workhorse stopgap",
					Paragraphs: []string{
						"Gemini 3.6 Flash is the headline release: a workhorse update priced at $1.50 per million input tokens and $7.50 per million output tokens, positioned as the model most developers will actually touch day to day. Google says it cuts output token usage by 17% compared to 3.5 Flash on the Artificial Analysis Index - and by as much as 65% on the DeepSWE coding benchmark specifically. On DeepSWE itself, 3.6 Flash scores 49% versus 3.5 Flash's 37%; on MLE-Bench, 63.9% versus 49.7%; on OSWorld-Verified, 83.0% versus 78.4%.",
						"Alongside it: Gemini 3.5 Flash-Lite, a cut-rate variant at $0.30/$2.50 per million tokens that Google says outputs 350 tokens per second - fast and cheap enough to out-benchmark Google's own mid-tier Gemini 3 Flash on SWE-Bench Pro (54.2% versus 49.6%) and post a 54% score on Terminal-Bench 2.1, against 3.5 Flash's 31%.",
					},
				},
				{
					Heading: "The cyber-only model",
					Paragraphs: []string{
						`And then there's Gemini 3.5 Flash Cyber, the odd release out: a security-tuned model fine-tuned specifically for vulnerability detection and patching, deployed exclusively through Google's CodeMender security agent and offered, for now, only to governments and "trusted partners" in a limited pilot.`,
						`Three ships in one day - and the fourth, the one carrying the "Pro" badge that actually matters for competitive positioning, nowhere in sight.`,
					},
				},
				{
					Heading: "What slipped",
					Paragraphs: []string{
						"Google's own account of the delay traces back to May 19, when the company told developers at I/O that Gemini 3.5 Pro would arrive the following month. June came and went. By mid-July, Bloomberg-sourced reporting said the model was running months behind its internal schedule, and outlets covering the slip converged on a consistent set of problems: coding performance that fell short of Google's internal targets, plus hallucinations and reliability gaps that kept surfacing in real-world testing. Google has not attached its own name to a specific technical explanation beyond confirming the model remains in partner testing - notable, since Google has historically been willing to name the concrete benchmarks it is chasing when a launch is imminent.",
						"The timing matters as much as the technical detail. Gemini 3.5 Pro was supposed to be Google's answer to a summer in which its two biggest rivals both landed flagship releases: OpenAI took GPT-5.6 to general availability in mid-July, and Anthropic has had Fable 5 and Sonnet 5 running in production enterprise workloads since June. Every month Pro stays in testing is a month enterprise buyers spend building workflows on someone else's model instead - and switching costs, once integrations are live, do not reverse easily.",
					},
				},
				{
					Heading: "Good enough is the strategy",
					Paragraphs: []string{
						`Google's bet, for now, looks less like "wait for the best model" and more like "ship enough good-enough models that nobody notices Pro is missing." The DeepSWE and MLE-Bench numbers suggest the Flash line is getting genuinely more capable, not just cheaper - a jump from 37% to 49% on DeepSWE and 49.7% to 63.9% on MLE-Bench would have counted as flagship news eighteen months ago.`,
						`If Flash-tier models keep closing the gap to what a "Pro" model used to promise, the more interesting question stops being when Gemini 3.5 Pro ships, and starts being whether, by the time it does, anyone still needs it as badly as they would have in June.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Google, Introducing Gemini 3.6 Flash, 3.5 Flash-Lite, and 3.5 Flash Cyber: https://blog.google/innovation-and-ai/models-and-research/gemini-models/gemini-3-6-flash-3-5-flash-lite-3-5-flash-cyber/",
						"9to5Google, Gemini 3.5 Pro delays due to coding performance: https://9to5google.com/2026/07/16/gemini-3-5-pro-delays/",
						"Geeky Gadgets, Stopgap Gemini 3.6 Flash May Launch During Gemini 3.5 Pro Delay: https://www.geeky-gadgets.com/gemini-3-5-pro-delayed-again/",
						"Geeky Gadgets, Google Cuts Gemini 3.6 Flash Tokens for Coding Tasks: https://www.geeky-gadgets.com/gemini-3-6-flash-release/",
					},
				},
			},
		},
	}, posts...)
}
