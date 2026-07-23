package content

func init() {
	posts = append([]Post{
		{
			Title:   `A Judge Wouldn't Stop Meta's Layoffs. He Also Said the AI Discrimination Claims Raise "Serious Questions."`,
			Slug:    "meta-ai-layoff-discrimination-orrick-injunction-2026",
			Date:    "July 23, 2026",
			Tag:     "Legal",
			Summary: "Twenty-six Meta employees on medical, parental, or disability leave say a scoring algorithm graded them out of a job during May's 8,000-person cut. A federal judge let their terminations proceed anyway on Friday - while leaving the door open to reversing course.",
			Related: []Link{
				{
					Title: "Meta Laid Off 8,000 People to Fund AI. Then Zuckerberg Admitted It Isn't Working Yet.",
					Slug:  "meta-microsoft-ai-layoffs-2026-jobs-cut-fund-buildout",
				},
				{
					Title: "82 Percent of Enterprises Just Found an AI Agent They Didn't Know They Had",
					Slug:  "enterprise-ai-agent-governance-visibility-gap-control-plane-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"Twenty-six Meta employees went to federal court in Oakland last week arguing something relatively new for a layoff case: that an algorithm, not a person, decided they should lose their jobs - and that the algorithm could not tell the difference between someone who was not performing and someone who was on approved medical leave. On July 17, a judge agreed the question was serious enough to take seriously. He just would not stop the layoffs to answer it.",
						"The case grows out of the roughly 8,000 job cuts - about 10% of Meta's global workforce - that the company announced in May as part of its push to redirect spending toward AI infrastructure, the same round AINews covered on July 15 when Mark Zuckerberg conceded the AI investment was not yet paying off. What was not clear at the time was exactly how Meta decided which employees, among those thousands, would go.",
					},
				},
				{
					Heading: "What the lawsuit says",
					Paragraphs: []string{
						"According to the lawsuit, filed July 13 in the U.S. District Court for the Northern District of California, Meta ranked employees using a mix of internal AI tools - including its Metamate assistant - alongside keystroke and click-tracking software, AI token-usage dashboards, and algorithmically generated performance scores. That tracking traces back to an internal program the suit calls the Model Capability Initiative, which it says Meta installed on U.S. employees' work laptops starting in April with no opt-out, logging keystrokes, mouse movement, and periodic screenshots.",
						"The plaintiffs - among them employees on medical, parental, pregnancy, disability, and other protected leave or accommodations, according to filings and wire coverage - say the scoring system had no mechanism to account for approved time away from a keyboard. Weeks or months of legally protected leave simply registered as missing data, the complaint alleges, dragging down aggregate scores with no leave-neutral adjustment applied before rankings were finalized.",
						"They are suing under the Family and Medical Leave Act, the Americans with Disabilities Act, the Pregnancy Discrimination Act, the Pregnant Workers Fairness Act, and Title VII's disparate-impact provision - a combination that, if it holds up, would make this one of the first real legal tests of whether AI-assisted layoff selection can itself amount to discrimination.",
					},
				},
				{
					Heading: "Why Orrick let the layoffs proceed",
					Paragraphs: []string{
						`U.S. District Judge William Orrick declined to grant the emergency relief the workers wanted: a temporary restraining order preserving their jobs while the underlying claims are argued in arbitration. Orrick found the plaintiffs had not shown the "irreparable harm" that emergency injunctions require - a familiar, high bar in employment law, where lost wages are usually treated as recoverable rather than irreversible. The layoffs went forward as scheduled on July 22.`,
						`But Orrick did not wave the underlying claims away. According to plaintiffs' attorneys, his written order stated the case raises serious questions going to the merits, and that he could revisit his decision if the parties produce more evidence on whether and how AI actually drove the termination decisions. A separate motion for a longer-term preliminary injunction remains pending. Meta, for its part, maintains there is nothing to revisit: "Workforce management and organizational decisions were and are made by people, not AI," a company spokesperson said, calling the claims meritless.`,
					},
				},
				{
					Heading: "The question arbitration has to answer",
					Paragraphs: []string{
						`That is the crux the case will have to settle in arbitration - not whether Meta used AI tools somewhere in its process, because plenty of large employers now do in one form or another, but whether the specific outputs of those tools drove specific termination decisions, and whether anyone checked those outputs against the legal protections some affected employees were entitled to before finalizing them.`,
						`That is a narrower and more provable question than "did an algorithm do this," and employment lawyers watching the case say it is the kind of question that will define the next wave of AI-in-HR litigation regardless of how this particular dispute resolves. For a company that spent its May earnings call explaining why AI spending would eventually pay off, its lawyers are now spending July explaining, in a different courtroom, exactly how little of that AI was supposedly involved in one of its highest-stakes people decisions of the year.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Reuters via Investing.com, US judge won't block Meta from laying off workers who filed AI discrimination lawsuit: https://www.investing.com/news/stock-market-news/us-judge-wont-block-meta-from-laying-off-workers-who-filed-ai-discrimination-lawsuit-4798876",
						"Associated Press via ABC News, 26 Meta employees sue, alleging layoff picks hit workers on medical, parental leave: https://abcnews.com/Technology/wireStory/26-meta-employees-sue-alleging-ai-driven-layoff-134764788",
						"SFGATE, Meta laid off thousands to prioritize AI. Former employees say AI was used to lay them off: https://www.sfgate.com/tech/article/meta-disability-lawsuit-22347135.php",
						"Complaint, Doe et al. v. Meta Platforms, Inc., N.D. Cal., filed July 13, 2026: https://www.courthousenews.com/wp-content/uploads/2026/07/meta-employee-complaint.pdf",
					},
				},
			},
		},
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
