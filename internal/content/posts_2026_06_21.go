package content

func init() {
	posts = append([]Post{
		{
			Title:   "Two Roads, One Month: The EU Tightens Its AI Rulebook as Washington Moves to Tear Up the States'",
			Slug:    "eu-ai-act-deadline-us-state-preemption-divergence-2026",
			Date:    "June 21, 2026",
			Tag:     "Policy",
			Summary: "In the span of a few June weeks, the world's two largest economies pulled in opposite directions on AI governance: Brussels racing to operationalize the most comprehensive AI law ever written, even as its own guidance slips, and Washington trying to stop its own states from writing rules at all.",
			Related: []Link{
				{
					Title: "The AI Rulebook Is Moving From Principles to Plumbing",
					Slug:  "ai-policy-rulebook-principles-to-plumbing-2026",
				},
				{
					Title: "AI Compliance Has a Calendar Now: The Global Rulebook Moves From Debate to Deadlines",
					Slug:  "ai-compliance-calendar-global-rulebook-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"For two years the debate over AI regulation has been framed as a single dial - turn it toward \"safety\" and you slow the labs down, turn it toward \"innovation\" and you let them run. June 2026 made clear that framing is wrong. The real question splitting the Atlantic isn't how much to regulate. It's who holds the pen.",
						"Look at what happened in a single month.",
						"In Brussels, the European Commission finally published its draft guidelines on how to classify \"high-risk\" AI systems under the EU AI Act - the category that triggers the law's most demanding obligations around risk management, data governance, human oversight, and documentation. The guidance landed on May 19, 2026, with a stakeholder consultation open until July 23. That timing matters for an awkward reason: the guidelines were originally due February 2. They arrived more than three months late, and they arrived with the law's headline high-risk obligations scheduled to take effect on August 2, 2026.",
						"Read that sequence slowly. The instructions for how to comply showed up roughly ten weeks before compliance is supposed to begin - and those instructions are still in draft, still open for comment. A company trying to determine whether its hiring tool or its credit model counts as \"high-risk\" is being asked to hit a deadline using a rulebook that isn't finished.",
						"The substance of the draft is genuinely useful, and worth knowing even outside Europe, because it signals where the regulatory mind is going.",
					},
				},
				{
					Heading: "The High-Risk Line Gets Sharper",
					Paragraphs: []string{
						"You can't disclaim your way out. If your system can be used for a high-risk purpose and your marketing leans into that purpose, you don't escape the category by burying an exclusion in your terms of service. The Commission explicitly says merely asserting that high-risk uses are \"excluded\" is not enough.",
						"Credit and hiring are squarely in. Systems that assess creditworthiness or generate credit scores are high-risk, with a narrow carve-out for fraud detection. Recruitment tools that filter, score, or rank candidates count too - even when they're sold as merely \"assistive.\" If the tool materially influences who gets hired or promoted, it's in scope.",
						"There's an off-ramp, but a narrow one. Article 6(3) lets some systems escape high-risk status if they only perform a narrow procedural task, improve the output of prior human work, flag patterns without replacing human judgment, or stay under meaningful human review. \"We kept a human in the loop\" is a real defense - but only if the human is actually deciding.",
					},
				},
				{
					Heading: "Washington Moves The Other Way",
					Paragraphs: []string{
						"Now cross the ocean. In the same window, the U.S. federal government moved in precisely the opposite direction - not to write AI rules, but to stop others from writing them. President Trump issued an executive order directing federal agencies to push back on state AI laws the administration considers excessively burdensome. The order stands up a Justice Department task force to challenge problematic state statutes, instructs the Commerce Department to compile a list of offending regulations, and dangles a stick: states with AI laws deemed incompatible with the administration's approach could see federal funding restricted, including money tied to broadband deployment programs. Alongside the order, the White House released a national framework urging Congress to preempt state AI laws that don't align with federal preferences.",
						"The administration drew some lines - it signaled it would not go after laws aimed squarely at fraud, consumer protection, or child safety. But the direction is unmistakable: Washington wants AI governance centralized, and it wants the centralization to point toward a lighter touch.",
						"You can see the chilling effect in the statehouses. Oklahoma is the tell. Most of the state's AI bills stalled in the 2026 session - a political-deepfakes measure cleared committee and then died, a three-bill package passed the House and stalled in the Senate, an earlier deepfake bill never even got a hearing. The one law that made it through, Senate Bill 1734, signed May 12, 2026, is narrow and uncontroversial: AI guardrails for public schools, with parental notification and teacher review. When the federal government is threatening your broadband money, \"narrow and uncontroversial\" starts to look like the only safe place to legislate.",
					},
				},
				{
					Heading: "The Compliance Gap Widens",
					Paragraphs: []string{
						"Put the two scenes side by side and the contrast is almost too clean. Europe is building a single, detailed, top-down rulebook and straining to make its own deadlines. America is dismantling a bottom-up patchwork and betting that less is more - or at least that fewer regulators is more. One side's risk is a law so intricate it can't be implemented on time. The other side's risk is a vacuum where the only rules left standing are the ones nobody bothered to fight.",
						"For anyone building or deploying AI across both markets, this is the part that actually bites. The compliance gap between the EU and the US isn't narrowing - it's widening, and in real time. A hiring model that is a regulated, high-risk system in Frankfurt may face essentially no AI-specific federal requirement in Phoenix, and a state requirement that the federal government is actively trying to erase.",
						"The work of governance doesn't disappear in that gap; it just moves inside the company. Multinationals are increasingly going to govern to the strictest jurisdiction they touch, not because they're required to everywhere, but because maintaining two entirely different AI systems for two continents is more expensive than maintaining one careful one. In other words, the EU's rulebook may end up setting the floor for companies that operate in places with no floor at all - the \"Brussels effect,\" running quietly through the back office.",
					},
				},
				{
					Heading: "Who Holds The Pen",
					Paragraphs: []string{
						"Here's what's worth sitting with. We spent two years arguing about whether to regulate AI. The more consequential argument, the one that surfaced this month, is about who. Centralized and demanding, or centralized and permissive - both Europe and Washington are reaching for the steering wheel; they just want to drive in opposite directions.",
						"The companies and the public end up living in the space between, where the real rules are being written less by legislators than by whichever deadline slips, whichever funding gets threatened, and whichever standard a global business decides is simply easier to meet everywhere.",
						"The deadline that actually governs AI may not be the one in the statute. It may be the one in the budget.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"European Commission, \"Guidelines for providers and deployers of AI high-risk systems\": https://digital-strategy.ec.europa.eu/en/policies/guidelines-ai-high-risk-systems",
						"European Commission, \"Targeted consultation on the draft guidelines for the classification of high-risk artificial intelligence systems\": https://digital-strategy.ec.europa.eu/en/consultations/targeted-consultation-draft-guidelines-classification-high-risk-artificial-intelligence-systems",
						"DLA Piper, \"EU Commission draft guidelines on classification of high-risk AI systems - key points\": https://www.dlapiper.com/en-us/insights/publications/2026/06/eu-commission-draft-guidelines-on-classification-of-high-risk-ai-systems-key-points",
						"White House, \"Ensuring a National Policy Framework for Artificial Intelligence\": https://www.whitehouse.gov/presidential-actions/2025/12/eliminating-state-law-obstruction-of-national-artificial-intelligence-policy/",
						"Axios, \"Trump's shadow AI policy,\" June 18, 2026: https://www.axios.com/2026/06/18/trump-shadow-ai-policy",
						"KGOU, \"Oklahoma holding back on AI regulations amid Trump's order for states not to stifle the new technology,\" June 18, 2026: https://www.kgou.org/politics-and-government/2026-06-18/oklahoma-holding-back-on-ai-regulations-amid-trumps-order-for-states-not-to-stifle-the-new-technology",
						"Oklahoma Legislature, Senate Bill 1734: https://www.oklegislature.gov/BillInfo.aspx?Bill=SB1734&Session=2600",
						"LegiScan, \"OK SB1734\": https://legiscan.com/OK/text/SB1734/id/3374080",
						"Author article handoff: https://docs.google.com/document/d/17O4pJWu4miuiBGqGj7tFyq2uW7Omxpm93aYugfhoVYM/edit",
					},
				},
			},
		},
	}, posts...)
}
