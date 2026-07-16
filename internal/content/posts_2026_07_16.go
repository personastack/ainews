package content

func init() {
	posts = append([]Post{
		{
			Title:   "82 Percent of Enterprises Just Found an AI Agent They Didn't Know They Had",
			Slug:    "enterprise-ai-agent-governance-visibility-gap-control-plane-2026",
			Date:    "July 16, 2026",
			Tag:     "Agents / Governance",
			Summary: "Companies are deploying autonomous agents faster than they can count them — and a new wave of \"control plane\" startups is racing to fill the gap before the first big incident does the explaining for them.",
			Related: []Link{
				{
					Title: "The Agents Are Talking. Nobody Checked Their IDs.",
					Slug:  "agent-identity-zero-trust-non-human-identity-2026",
				},
				{
					Title: "Nobody Funded a Smarter Agent This Week. They Funded the Gym.",
					Slug:  "agent-training-environments-reliability-investment-bet-2026",
				},
				{
					Title: "Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them",
					Slug:  "ai-agent-spending-governance-gap-control-plane-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"Ask a CIO in July 2026 how many AI agents are running inside their company, and you'll get one of two answers: a confident-sounding number, or an honest shrug. New survey data suggests the confident number is usually wrong.",
						"According to a joint study from the Cloud Security Alliance and Token Security, titled Autonomous but Not Controlled, 68 percent of organizations say they have high visibility into the AI agents and autonomous workflows running in their environment. But when the same organizations went looking, 82 percent discovered at least one agent or workflow that security or IT didn't know existed. Sixty-five percent had already suffered an AI agent security incident in the past year — most commonly, the report found, resulting in data exposure.",
						"That gap between what companies believe about their agent fleets and what's actually running is becoming the defining problem of enterprise AI's second year.",
					},
				},
				{
					Heading: "The Adoption Curve Got Steep, Fast",
					Paragraphs: []string{
						"It's worth remembering how quickly this happened. In August 2025, Gartner forecast that 40 percent of enterprise applications would ship with a task-specific AI agent embedded by the end of 2026 — up from less than 5 percent the year before. That's roughly an eightfold jump in twelve months, one of the fastest platform shifts Gartner has tracked since the move to public cloud.",
						"Deloitte's 2026 State of AI in the Enterprise survey, which polled 3,235 business and IT leaders across 24 countries, shows that forecast playing out. Seventy-four percent of respondents expect their companies to be using AI agents at least \"moderately\" by 2027; 23 percent expect \"extensive\" use. But the same survey found that security and compliance concerns have gotten worse, not better, as adoption accelerated: 80 percent of leaders piloting agents now cite security and compliance as their leading obstacle, up from 68 percent just a year earlier. Only 21 percent say their organization has a mature governance model for agentic AI in place at all.",
						"Put plainly: the agents showed up on schedule. The guardrails didn't.",
					},
				},
				{
					Heading: "Why This Isn't Just \"Shadow IT\" Again",
					Paragraphs: []string{
						"It's tempting to file this under a familiar heading — employees signing up for unsanctioned SaaS tools, IT scrambling to catalog them after the fact. Enterprises have lived through that cycle before, with browser extensions, with personal Dropbox accounts, with the first wave of generative AI chatbots.",
						"Agents are a different animal. A rogue spreadsheet tool stores data insecurely. A rogue agent acts — it reads inputs, executes multi-step workflows, and decides what to do next without a human necessarily reviewing each step. The kinds of agents now running inside large organizations modify financial records, trigger payments, and approve workflows with the operational authority the company granted them, intentionally or not. When the artifact you've lost track of can also move money, \"shadow IT\" undersells the stakes.",
					},
				},
				{
					Heading: "The Vendors Racing to Build the Missing Layer",
					Paragraphs: []string{
						"That gap is exactly why a new category of \"agent control plane\" companies has materialized this summer. On July 14, at Google Cloud Next 2026, governance infrastructure company Devenex launched what it calls an Execution Control Plane for AI agents — pre-execution policy enforcement, human-in-the-loop checkpoints for consequential actions, and an immutable audit trail meant to answer the question \"what did this agent actually do, and who approved it?\" A company co-founder put the pitch bluntly: \"no infrastructure layer exists to govern what these agents actually do.\"",
						"The infrastructure incumbents are making the same bet at larger scale. HPE's June announcement with Nvidia — bringing agentic AI into production with added security, governance, and sovereignty controls — leans on Nvidia's new Vera CPU, purpose-built to handle the tool calls and orchestration inherent to agent workflows, alongside Nvidia's Agent Toolkit for managing autonomous agents once they're live. Microsoft has moved in the same direction from the software side: Foundry Agent Service reached general availability in early July with sandboxed execution sessions and multi-agent orchestration, while Microsoft 365 Copilot picked up policy-based bulk agent installation and scheduled-prompt controls aimed squarely at IT teams trying to keep an inventory of what's deployed.",
						"None of these products existed as shipping enterprise features a year ago. That they all arrived within weeks of each other isn't a coincidence — it's a market responding to the same survey data cited above, in real time.",
					},
				},
				{
					Heading: "The Thing Worth Sitting With",
					Paragraphs: []string{
						"Here's the uncomfortable part: governance tooling is, definitionally, a step behind the thing it governs. Every agent control plane launched this month was built in response to agents that were already running unmonitored somewhere. The 82 percent figure isn't a warning about a future problem — it's a description of the present one.",
						"The open question for the back half of 2026 isn't whether enterprises will keep deploying agents. Gartner's curve and Deloitte's adoption numbers both say that's settled. The real question is whether governance infrastructure can close the visibility gap before regulators, auditors, or an actual incident close it for them — because right now, the honest answer to \"how many agents do you have running\" is still, for most companies, a shrug dressed up as a number.",
						"Related reading: our July 3 piece on agent-to-agent authentication covered a narrower slice of this same problem — agents impersonating other agents in multi-agent systems. Our July 11 piece on agent training environments looked at the investment case for making agents more reliable before they're deployed. This piece is about what happens after they are.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/1sGgpkfvv1CdJSdH18zkeu7XvKvQLQJbwnwptCIw9IUA/edit",
						"Cloud Security Alliance / Token Security, Autonomous but Not Controlled report: https://cloudsecurityalliance.org/artifacts/autonomous-but-not-controlled-ai-agent-incidents-now-common-in-enterprises",
						"Cloud Security Alliance press release on unknown AI agents: https://cloudsecurityalliance.org/press-releases/2026/04/21/new-cloud-security-alliance-survey-reveals-82-of-enterprises-have-unknown-ai-agents-in-their-environments",
						"Gartner press release on enterprise applications with task-specific AI agents: https://www.gartner.com/en/newsroom/press-releases/2025-08-26-gartner-predicts-40-percent-of-enterprise-apps-will-feature-task-specific-ai-agents-by-2026-up-from-less-than-5-percent-in-2025",
						"Deloitte 2026 State of AI in the Enterprise: https://www.deloitte.com/us/en/about/press-room/state-of-ai-report-2026.html",
						"Devenex Execution Control Plane launch: https://www.prnewswire.com/news-releases/as-ai-agents-scale-enterprises-demand-execution-control--devenex-takes-control-302825506.html",
						"HPE and NVIDIA agentic AI factory announcement: https://www.hpe.com/us/en/newsroom/press-release/2026/06/hpe-brings-agentic-ai-into-production-with-nvidia-delivering-security-governance-scale-and-sovereignty.html",
						"Microsoft Foundry Agent Service announcement: https://devblogs.microsoft.com/foundry/agent-service-build2026/",
						"Microsoft 365 Copilot release notes and scheduled prompts documentation: https://learn.microsoft.com/en-us/microsoft-365/copilot/release-notes",
					},
				},
			},
		},
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
