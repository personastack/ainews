package content

func init() {
	posts = append([]Post{
		{
			Title:   "The First Big American AI Law Was Supposed to Take Effect Yesterday. It No Longer Exists.",
			Slug:    "colorado-ai-act-repealed-first-us-ai-law-deadline-2026",
			Date:    "July 1, 2026",
			Tag:     "Policy",
			Summary: "On June 30, 2026, Colorado's landmark AI Act was set to become the first comprehensive AI law in the United States to actually bind companies. Instead, the deadline came and went over a statute the state had quietly already repealed and rewritten — and the provision that got cut is the one that mattered most.",
			Related: []Link{
				{
					Title: "Two Roads, One Month: The EU Tightens Its AI Rulebook as Washington Moves to Tear Up the States'",
					Slug:  "eu-ai-act-deadline-us-state-preemption-divergence-2026",
				},
				{
					Title: "Washington Wrote the Rulebook for Frontier AI — And the First Lab It Touched Is Suing",
					Slug:  "us-ai-national-security-executive-order-anthropic-lawsuit-2026",
				},
				{
					Title: "The AI Rulebook Is Moving From Principles to Plumbing",
					Slug:  "ai-policy-rulebook-principles-to-plumbing-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"By the middle of 2026, June 30 was the date every compliance team in America circled. It was the day Colorado's SB 24-205 — signed in May 2024 as the first comprehensive state AI law in the country, and later delayed from its original February 1 effective date — was finally supposed to take effect. Yesterday, that day arrived. And there was nothing there.",
						"The law had already been repealed. Six weeks before its own deadline, Colorado gutted and replaced it. What was supposed to be a historic first is instead a case study in how hard it turns out to be for a single state to regulate frontier AI — and in which protections are the first to be traded away when the pressure comes.",
					},
				},
				{
					Heading: "What SB 24-205 actually did",
					Paragraphs: []string{
						"When Governor Jared Polis signed SB 24-205 in 2024, it was genuinely ambitious. The law reached both the companies that build \"high-risk\" AI systems and the companies that deploy them, across the decisions that shape people's lives: employment, housing, lending, insurance, health care, and education.",
						"Three obligations gave the law its teeth. Developers and deployers had to run a formal risk-management program. They had to conduct impact assessments before putting a high-risk system into use. And — the heart of the statute — they had to exercise \"reasonable care\" to protect consumers from algorithmic discrimination. That duty of care was the part civil-rights advocates cared about and the part industry feared, because it created real, litigable liability when an automated system produced discriminatory outcomes.",
						"That combination made Colorado the American analog to the EU AI Act's risk-based approach: not a ban, but a duty to document, assess, and prevent harm.",
					},
				},
				{
					Heading: "Death by a thousand delays",
					Paragraphs: []string{
						"The law never got the chance to work. Its original effective date was February 1, 2026. Colorado then pushed it to June 30, 2026 via SB 25B-004, signed in August 2025. That extra time mattered: it gave a business coalition more room to argue the compliance burden was unworkable, especially for smaller firms that deploy AI without building it.",
						"Then the courts got involved. In early 2026, a lawsuit filed by xAI led a court to temporarily suspend enforcement — and the U.S. Department of Justice intervened to support the challenge, a striking signal that the federal government intended to actively contest state-level AI rules rather than defer to them.",
						"By the time June 30 rolled around, the deadline was purely symbolic. In May 2026 the Colorado General Assembly passed SB 26-189 through the legislature, and Governor Polis signed it on May 14. The new law repealed and reenacted the AI Act as something substantially smaller.",
					},
				},
				{
					Heading: "What survived — and what didn't",
					Paragraphs: []string{
						"Here is the part worth slowing down on, because it tells you where the fight really was.",
						"The replacement law narrows the target from \"high-risk AI systems\" to \"automated decision-making technology\" (ADMT) that \"materially influences\" consequential decisions in seven domains — education, employment, housing, financial services, insurance, health care, and essential government services. In their place it imposes four operational duties: notify people before an ADMT is used on them; disclose an adverse outcome within 30 days in plain language; correct inaccurate personal data on request; and provide meaningful human review and reconsideration \"to the extent commercially reasonable.\"",
						"Those are real transparency rights. But notice what left the building. The three most-contested obligations — the risk-management program, the impact assessment, and the duty of reasonable care to prevent algorithmic discrimination — are all gone. The single provision that turned Colorado's law from a disclosure regime into an anti-discrimination regime was the first thing cut.",
						"And the new law does not take effect until January 1, 2027. So the practical situation today, July 1, 2026, is that Colorado — the state that was supposed to be first — currently has no operative comprehensive AI law at all.",
					},
				},
				{
					Heading: "Why this is bigger than Colorado",
					Paragraphs: []string{
						"It is tempting to read this as one state's cold feet. It is more than that, because Colorado is not being left alone to figure it out.",
						"The Trump administration's December 2025 executive order on AI singled out the Colorado AI Act by name — the only state law specifically called out — arguing that its anti-discrimination obligations would effectively \"force AI models to produce false results.\" That order sits on top of a broader federalization push, including the National Policy Framework for AI released earlier this year, aimed at replacing the emerging state patchwork with a lighter-touch federal standard.",
						"Stack the pieces up and a pattern appears. A state passes an ambitious AI law. Industry warns it is unworkable and wins repeated delays. A lawsuit, backed by the federal government, freezes enforcement. The state rewrites the law into a narrower, disclosure-focused version — dropping the anti-discrimination duty of care — and pushes the date out another year. Meanwhile the federal government moves to preempt the whole category.",
						"Colorado was the test case for whether a single state could hold the line on AI accountability on its own. The answer, for now, is that it could not.",
					},
				},
				{
					Heading: "The takeaway",
					Paragraphs: []string{
						"The most revealing thing about America's first comprehensive AI law is not that it was delayed. It is which provision died. Disclosure rules — tell me a machine decided, tell me when it went against me, let me ask a human — survived, because almost no one opposes transparency in principle. The duty to actually prevent discriminatory outcomes did not survive, because that is where liability and cost live.",
						"That is the fault line to watch as the fight moves to Washington and to the next state to try. Transparency is comparatively easy to legislate. Accountability for outcomes is the expensive, contested part — and so far, when it comes under pressure, it is the part that goes first. June 30, 2026 was supposed to be the day American AI regulation grew teeth. It turned out to be the day we learned how quickly they come out.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Colorado General Assembly, SB 26-189 bill page: https://leg.colorado.gov/bills/sb26-189",
						"Colorado General Assembly, SB 24-205 bill page: https://leg.colorado.gov/bills/sb24-205",
						"Crowell & Moring, \"Colorado Hits Reset on AI Regulation,\" May 2026.",
						"Norton Rose Fulbright, \"Colorado enacts revised AI law,\" May 2026.",
						"Baker Botts, \"Colorado Repeals and Replaces AI Act,\" May 2026.",
						"White House, \"Ensuring a National Policy Framework for Artificial Intelligence,\" December 11, 2025: https://www.whitehouse.gov/presidential-actions/2025/12/eliminating-state-law-obstruction-of-national-artificial-intelligence-policy/",
						"Author article handoff: https://docs.google.com/document/d/1wvP0L0z89n3i2CrmDSwkt7OzfhDGlCdRa9j7OAjf9pY/edit",
					},
				},
			},
		},
	}, posts...)
}
