package content

func init() {
	posts = append([]Post{
		{
			Title:   "The AI Industry Graded Its Own Safety Homework. Nobody Passed.",
			Slug:    "ai-safety-index-summer-2026-anthropic-c-plus-pause-pledges-erode",
			Date:    "July 16, 2026",
			Tag:     "Policy",
			Summary: "The Future of Life Institute's third AI Safety Index is out, and the best score any company earned was a C+. The more interesting finding isn't the grade — it's what the report says labs have quietly stopped promising.",
			Related: []Link{
				{
					Title: "The Rocket Company Ships a Coding Model — And the Benchmark Depends on Who's Grading",
					Slug:  "grok-4-5-spacexai-cursor-coding-benchmark-harness-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"Every six months or so, seven independent AI researchers and governance experts sit down and do something the AI industry has mostly stopped doing to itself: grade its safety work on a curve with real consequences. The Future of Life Institute published its Summer 2026 AI Safety Index this month, running nine of the world's largest AI developers through 37 indicators across six domains — risk assessment, current harms, safety frameworks, existential safety, governance, and information sharing. The top score, earned by Anthropic, was a C+.",
						"Nobody got a B. Nobody got an A. That's not new — no lab has ever cracked a B on this index since it launched in 2024 — but the gap between \"still trying\" and \"not really trying\" has become the more useful story to tell.",
					},
				},
				{
					Heading: "The Scorecard",
					Paragraphs: []string{
						"Anthropic came out on top again, leading five of the six domains on the strength of comparatively strong transparency, a more developed safety framework, and consistent technical research output. OpenAI and Google DeepMind followed with a C each, with OpenAI notably taking the lead in the Risk Assessment domain thanks to a broader external evaluation suite. Meta improved to a D+. Three companies failed outright: xAI, DeepSeek, and Mistral each received an F — one flagship failing grade apiece from the US, China, and Europe.",
						"Worth flagging: the report's evaluation window closed June 3, 2026, so the \"xAI\" graded here is the pre-merger company — this is before the July 6 rebrand that folded xAI into SpaceXAI alongside SpaceX and Cursor, and more than a month before that combined entity shipped Grok 4.5. The F grade belongs to the standalone company that existed at evaluation time, not a verdict on whatever SpaceXAI does next. It's a reminder that safety report cards, like credit reports, are always a snapshot of the past — useful, but not predictive of a company that's actively restructuring underneath the grade.",
					},
				},
				{
					Heading: "The Part That Should Worry You More Than the Letters",
					Paragraphs: []string{
						"The grades are the headline. The trend line underneath is the actual news. The report documents that Anthropic, OpenAI, Google DeepMind, and Meta — the four highest-scoring companies on the index — have each weakened or voided prior pledges to unilaterally pause development if specific risk \"redlines\" were crossed. Some of the revised commitments are now conditional on what competitors do, which the panel's reviewers singled out with a specific, unflattering label: \"moving goalpost.\" Their assessment was blunt — this pattern has \"undermined safety frameworks across the board,\" not just at the companies doing the walking-back.",
						"A second, related retreat: several labs that had explicitly banned military applications of their technology in earlier years have reversed that position, pursuing defense-sector partnerships and national security agreements instead. The report frames this as a genuine strategic shift, not a one-off exception — companies that once drew a bright line around weapons and warfare work are now treating that line as negotiable.",
						"Neither of these is a scandal in the tabloid sense — no leaked memo, no whistleblower. It's slower and more structural than that: a group of companies that made public commitments during a calmer, more cautious period are, one by one, finding reasons those commitments no longer quite apply. Grades can recover next cycle. A norm, once abandoned by the market leaders, is harder to rebuild.",
					},
				},
				{
					Heading: "Where Everyone Actually Fails",
					Paragraphs: []string{
						"If you want the single most damning number in the report, skip the overall letter grades and look at the Existential Safety domain specifically — the category that asks whether labs have credible plans for controlling systems that could eventually exceed human-level capability across the board. No company scored above a C- here. Most scored a D or worse.",
						"It's not for lack of trying to show work. Anthropic points to its constitutional classifiers. OpenAI points to its public calls for new governance institutions. Google DeepMind points to its monitoring commitments. Meta points to loss-of-control provisions in its safety documentation. The panel looked at all of it and rendered a verdict that applies industry-wide, not to any single company: \"entirely inadequate.\"",
						"That phrase is worth sitting with. These are the four companies actually building the most capable systems on Earth, employing the researchers whose entire job is to think about exactly this problem, and an outside panel of specialists in that same field concluded that none of their current answers are good enough. It's not that the industry has no theory of how it might keep more capable systems under control — it's that the theories that exist don't survive contact with expert scrutiny.",
					},
				},
				{
					Heading: "The Gap Is Widening, Not Closing",
					Paragraphs: []string{
						"The report's broader conclusion is the one that should stick with you longer than any individual letter grade: the distance between the three market leaders and everyone else is growing, even as commercial safety progress across the leaders themselves has \"largely plateaued or regressed.\" Translate that out of report-speak: the companies with the most resources to invest in safety are spending less of their improving lead on it than they were, while the companies without the resources are falling further behind on both capability and safety simultaneously.",
						"None of this means the industry is reckless by design, and it doesn't mean every pledge revision is made in bad faith — competitive pressure is real, and a unilateral pause that a rival ignores can look, from inside a boardroom, like unilateral disarmament. But that's precisely the dynamic the Future of Life Institute is trying to put a number on: safety commitments that hold only as long as they're not competitively inconvenient aren't really commitments. They're forecasts. And on the evidence of this report, the forecast keeps getting revised downward.",
						"The next Index is due in roughly six months. The question worth tracking between now and then isn't whether Anthropic can turn a C+ into a B — it's whether \"moving goalpost\" becomes the industry's permanent operating mode, or a phase it grows out of once the current capability race cools down even slightly.",
						"Related reading: for the model these same labs are racing to ship next, see our coverage of Grok 4.5 and the coding-benchmark harness — released by the entity xAI became one month after this report's evaluation window closed.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/1e6IIJEcYMVv9_skCRHT1PvC51-EtHAEsipF8i0Wa4co/edit",
						"Future of Life Institute, AI Safety Index Summer 2026: https://futureoflife.org/ai-safety-index-summer-2026/",
						"InsideAIPolicy, Future of Life Institute hands out middling grades for major AI developers' safety efforts: https://insideaipolicy.com/ai-daily-news/future-life-institute-hands-out-middling-grades-major-ai-developers-safety-efforts",
						"Axios, AI companies retreat from safety pledges: https://www.axios.com/2026/07/07/report-ai-safety-pledges",
					},
				},
			},
		},
	}, posts...)
}
