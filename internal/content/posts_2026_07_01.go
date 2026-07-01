package content

func init() {
	posts = append([]Post{
		{
			Title:   "AI's Quiet Second Frontier: Foundation Models Built for the Data That Actually Runs Your Business",
			Slug:    "tabular-foundation-models-ai-second-frontier-structured-data-2026",
			Date:    "July 1, 2026",
			Tag:     "Models & Research",
			Summary: "While the world argued about the next chatbot, a different kind of foundation model quietly grew up around spreadsheets, databases, and data warehouses — and in 2026 it became a real category, with a billion-euro price tag to match.",
			Related: []Link{
				{
					Title: "The Year Companies Were Told to Use All the AI They Wanted. Then the Bill Came.",
					Slug:  "ai-cost-reckoning-tokenmaxxing-spend-caps-finops-2026",
				},
				{
					Title: "The Smartest Model in Your Stack Might Be the Smallest",
					Slug:  "retrieval-layer-small-embedding-models-rag-accuracy-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"Spend any time following AI in 2026 and you could be forgiven for thinking the whole field is one long argument about chatbots. Which lab shipped the smartest model this fortnight. Whose reasoning is deeper. Whose tokens are cheaper. It is a loud, fast, and genuinely important race.",
						"It is also not the only one.",
						"Off to the side, largely out of the headlines, a second frontier has been forming — one aimed not at language but at the least glamorous and most valuable data most companies own: the rows and columns sitting in their databases, spreadsheets, and warehouses. The models built for it have a clunky name — tabular foundation models, or more broadly, structured-data models — and in the first half of 2026 they stopped being an academic curiosity and became a competitive market. The clearest signal came in May, when enterprise-software giant SAP agreed to buy a small German lab called Prior Labs and committed to invest more than one billion euros over four years to scale it. The deal is expected to close this quarter.",
						"If you have never heard of Prior Labs, that is rather the point.",
					},
				},
				{
					Heading: "The model that skips training",
					Paragraphs: []string{
						"Prior Labs makes a model family called TabPFN, and to understand why SAP paid frontier-lab money for it, you have to understand how strange it is.",
						"Most machine learning on tabular data works like this: you take your dataset — say, ten thousand loan applications with twenty columns each — and you train a fresh model on it, tuning knobs for hours to squeeze out accuracy. Do it again for a new dataset, and you start over. This is the workhorse behind a huge share of real-world prediction: fraud scoring, churn, demand forecasting, medical risk. It is also slow, fiddly, and bespoke every single time.",
						"TabPFN throws that loop out. It is a transformer — the same broad architecture behind large language models — but instead of being trained on your data, it was pretrained once on a vast collection of synthetic tabular datasets, generated from structural causal models. In the process it learned something more general than any single task: how tabular problems tend to look. Show it a new dataset and its labels, and it makes predictions in a single forward pass, no training run required. The technical trick is a two-way attention mechanism where each cell attends both to other features in its row and to other examples in its column, which makes the model indifferent to the order of your rows and columns — exactly the symmetry real tables have.",
						"The original TabPFN v2 was striking enough to land in Nature in early 2025. Its headline result read like a typo: in 2.8 seconds, it matched or beat an ensemble of the strongest classical methods that had been tuned for four hours. But it had a ceiling — it worked best on small datasets, up to roughly ten thousand rows, which is a poor fit for the enterprise, where a million rows is a Tuesday.",
						"That ceiling is what fell away in 2026. TabPFN-3, released May 12, scales the same approach to datasets with up to a million training rows while staying practical on a single H100 GPU, using a reduced memory cache and row-chunked processing to keep the math tractable. On TabArena, a standard benchmark for this world, a single forward pass of TabPFN-3 outperforms every other entry and sits on the best point of the speed-versus-accuracy curve. Founders Frank Hutter, Noah Hollmann, and Sauraj Gambhir had turned a clever paper into something an enterprise could actually deploy — and SAP noticed.",
					},
				},
				{
					Heading: "Not one model, a category",
					Paragraphs: []string{
						"Here is the part that makes this a trend and not just an acquisition: Prior Labs is not alone.",
						"By 2026 the structured-data frontier has fractured into a small, competitive field. There are tabular models that read a single flat table — TabPFN, but also NEXUS from Fundamental, NICL from France's Neuralk AI, TabICL, and CARTE. There is a separate breed of relational foundation models, led by Kumo.ai's KumoRFM, that read many connected tables at once — think five to fifty tables joined by foreign keys — and discover patterns that hop across those boundaries, the way an actual business database is shaped. There are time-series foundation models like Chronos and TimesFM for forecasting. A Chinese effort, LimiX, reframes the whole problem as modeling the joint distribution over a table's variables and its missing values, and even proposes a scaling law for the category. Amazon's Mitra takes yet another angle with curated synthetic priors. Google has entered with a zero-shot tabular model of its own.",
						"The geography is telling too: Germany, France, Canada, China, and the United States all have serious entrants. This is not one company's bet. It is a field deciding, more or less at once, that structured data deserves its own foundation models rather than being awkwardly force-fed to a chatbot.",
					},
				},
				{
					Heading: "Why the boring frontier might matter more",
					Paragraphs: []string{
						"It is worth asking why a company like SAP — whose entire business is the structured data that runs supply chains, payrolls, and ledgers — would spend a billion euros here rather than on a language model.",
						"The answer is a quiet critique of the LLM era. Large language models are extraordinary at text, but ask one to reason over a table of numbers and it is working against its own grain, predicting plausible-looking tokens rather than computing an answer. It can hallucinate a figure with total confidence. Structured-data models are built for the opposite virtues: deterministic, repeatable predictions with clear feature attribution — outputs you can audit, defend to a regulator, and reproduce next quarter. For a bank deciding on a loan or a hospital triaging risk, that auditability is not a nice-to-have. It is the whole job.",
						"There is a thread here that connects to a story we have been tracking all year. Enterprises spent 2025 and early 2026 pouring money into LLM pilots and are now, in the great cost reckoning, asking the uncomfortable question of what actually paid off. A model that skips the training loop, runs on one GPU, and produces auditable numbers is a very different value proposition from a chatbot metered by the token. It may turn out that the most durable enterprise AI of this era was never the one that talked back.",
					},
				},
				{
					Heading: "Something to think about",
					Paragraphs: []string{
						"The chatbot race will keep making the headlines, and it should — it is genuinely reshaping how we write, code, and search. But the history of computing is full of moments where the loud revolution and the important one were not the same thing.",
						"So here is the question worth sitting with: if the last three years taught machines to be fluent in our language, the next few may teach them to be fluent in our data — the ledgers, the sensor logs, the customer tables that quietly run the world. That frontier does not demo well. It will never trend. But it is aimed squarely at the numbers your business actually lives and dies by. Keep one eye on the chatbots. Keep the other on the spreadsheets.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"SAP Newsroom, \"SAP to Acquire Prior Labs,\" May 4, 2026: https://news.sap.com/2026/05/sap-to-acquire-prior-labs-establish-frontier-ai-lab-europe/",
						"Prior Labs, TabPFN-3 Technical Report, arXiv:2605.13986, May 12, 2026: https://arxiv.org/abs/2605.13986",
						"Hollmann et al., \"Accurate predictions on small data with a tabular foundation model,\" Nature, 2025: https://www.nature.com/articles/s41586-024-08328-6",
						"TabArena benchmark: https://tabarena.ai/",
						"LimiX, arXiv:2509.03505: https://arxiv.org/abs/2509.03505",
						"Kumo.ai, KumoRFM: https://kumo.ai/",
						"Google Research, \"Introducing TabFM: A zero-shot foundation model for tabular data\": https://research.google/blog/introducing-tabfm-a-zero-shot-foundation-model-for-tabular-data/",
						"Author article handoff: https://docs.google.com/document/d/12rAAjra_q0DaHZYUGB8KqbSYGRdGbMUctkJb4uyw7tQ/edit",
						"All figures verified against primary sources as of July 1, 2026.",
					},
				},
			},
		},
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
