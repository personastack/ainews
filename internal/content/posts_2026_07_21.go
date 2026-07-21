package content

func init() {
	posts = append([]Post{
		{
			Title:   `TSMC Just Pushed Its Arizona Bet to $265 Billion. The New Money Finally Targets the Part Critics Called a "Paperweight."`,
			Slug:    "tsmc-arizona-265-billion-packaging-bottleneck-2026",
			Date:    "July 21, 2026",
			Tag:     "Hardware",
			Summary: "The world's biggest chipmaker says it's building a dozen facilities in the desert, and for the first time enough of them are for packaging, not just wafers, to answer years of skepticism. Whether the timeline survives contact with reality is a separate question.",
			Related: []Link{
				{
					Title: "The Chip Industry Just Had Its Best Quarter Ever. Wall Street Sold It Anyway.",
					Slug:  "chip-earnings-record-profits-stock-selloff-kimi-k3-2026",
				},
				{
					Title: "Nvidia's Roadmap Just Hit the Reticle Limit",
					Slug:  "nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`Taiwan Semiconductor Manufacturing Company told investors on its July 16 earnings call that it's adding another $100 billion to its Arizona buildout, pushing total committed investment there to $265 billion — the largest foreign direct investment in US history, according to the Trump administration, which negotiated it as part of a broader US-Taiwan trade and investment deal struck in January. TSMC now says the full site could eventually include 12 facilities: 10 wafer fabs, two advanced-packaging plants, and a research and development center.`,
						`That packaging detail is easy to skim past. It shouldn't be.`,
					},
				},
				{
					Heading: "Why \"just add fabs\" was never the whole plan",
					Paragraphs: []string{
						`For years, the standard critique of TSMC's Arizona push was that it made half a product. Wafers came out of the ground fabs looking like finished chips, but they weren't — advanced AI accelerators need a separate assembly step called CoWoS (chip-on-wafer-on-substrate) that stacks the logic die together with high-bandwidth memory into a single package. Without it, a wafer is expensive sand. Earlier Arizona output reportedly had to be shipped back to Taiwan for that step, which is the origin of critics calling the first-generation plant a "paperweight."`,
						`CoWoS has since become the tightest link in the entire AI supply chain. TSMC's own CEO, C.C. Wei, told shareholders at the company's June annual meeting that CoWoS capacity was "extremely tight and sold out through 2026," with the company's three advanced-packaging backend facilities booked through 2027 and lead times running 52 to 78 weeks. Nvidia alone has locked up more than 70% of TSMC's CoWoS-L capacity, leaving AMD, Broadcom, and Marvell to split what's left. TSMC has been racing to scale CoWoS output from roughly 35,000 wafers a month in late 2024 toward a projected 130,000 a month by the end of this year — nearly 4x growth in under two years — and it's still not enough. The $100 billion increment, and the two packaging plants now on the books alongside the fabs, is TSMC finally building the part of the pipeline that made "Made in Arizona" chips incomplete.`,
					},
				},
				{
					Heading: "The politics wrapped around the announcement",
					Paragraphs: []string{
						`The investment lands as the practical output of a trade deal Washington has been pushing hard. Commerce Secretary Howard Lutnick framed it directly: "TSMC's announcement of an additional $100 billion investment following our historic deal on trade and investment with Taiwan will create tens of thousands of American jobs and bring advanced semiconductor manufacturing back to America." Lutnick's stated goal has been to relocate 40% of Taiwan's entire chip supply chain to the US, using the threat of tariffs as high as 100% on Taiwanese chipmakers who don't build domestically as leverage. Taiwanese officials have pushed back hard on that target, calling it "impossible" given how deeply the island's semiconductor ecosystem — thousands of specialized suppliers, decades of institutional expertise — is rooted at home.`,
						`TSMC's local economic footprint is already real: the company employs roughly 3,500 people in Arizona today, mostly locally hired, and Steve Zylstra, president of the Arizona Technology Council, says the investment is drawing a broader magnet effect — Nvidia and Apple have both relocated operations to the state, and roughly 50 supply-chain companies have followed TSMC there. TSMC estimates the new construction alone will support 40,000 jobs over the next four years, with Zylstra projecting 6,000 permanent positions once just the first three fabs are running, many requiring only a high school diploma plus an 11-week training course now offered free through local universities.`,
					},
				},
				{
					Heading: "Why the skepticism is earned, not just cynicism",
					Paragraphs: []string{
						`None of that means the $265 billion figure should be read as a delivery date. TSMC's own language on the new fabs is conditional — the company said the expansion would "likely" produce four additional plants, with pacing set by "market situation," and no construction timeline has been published. The planned facility count has crept upward before: from an original eight plants to today's claimed 12, without a matching acceleration in what's actually been built. Cost is a real constraint too — fabricating chips in Arizona is estimated to run several times more expensive than in Taiwan, a gap TSMC has absorbed so far but hasn't eliminated. And TSMC still keeps its most advanced process nodes at home: US fabs are currently limited to producing chips roughly three generations behind Taiwan's cutting edge, meaning Apple's newest devices still can't be built domestically. Apple, notably, is simultaneously exploring Intel and Samsung as alternative foundry partners, a hedge that could reshape the "market situation" TSMC says will determine its own Arizona pace.`,
					},
				},
				{
					Heading: "Capital isn't waiting to find out",
					Paragraphs: []string{
						`Zoom out from the mega-project, and the same bet is showing up at a completely different scale. Etched, an AI inference-chip startup founded in 2022 by three Harvard dropouts, is reportedly negotiating two separate funding rounds at once: one led by Sequoia Capital around a $10 billion valuation, another led by Jane Street around $20 billion — roughly quadrupling where the company stood before, on the strength of what it says is $1 billion in customer interest for its inference-specific chips. Neither deal has closed, and terms could still shift. But the willingness of two separate investor groups to negotiate at that range, for a company that hasn't begun large-scale commercial shipping yet, says something on its own.`,
						`It's a useful contrast to the story AINews covered on July 18, when record quarterly earnings from TSMC, Samsung, and ASML collided with a semiconductor stock selloff that erased roughly $3.3 trillion in market value in a matter of weeks. Public markets, spooked by questions about whether AI compute demand is durable, sold the chip complex anyway. Private capital — from TSMC's own board approving record capex, to venture firms fighting over Etched's cap table — doesn't appear to be having the same argument.`,
					},
				},
				{
					Heading: "What to actually watch",
					Paragraphs: []string{
						`The dollar figure attached to TSMC's Arizona plan will keep growing in press releases regardless of what gets poured in concrete. The more useful number to track is CoWoS and advanced-packaging capacity specifically, since that's the stage of production that's currently gating how many finished AI accelerators the entire industry can ship, not wafer starts. If Arizona's two new packaging plants come online on anything close to schedule, it would mark the first time the US builds a complete AI chip, start to finish, on its own soil. If they slip the way earlier Arizona timelines have, the "paperweight" critique won't have been solved — it'll just have gotten $100 billion more expensive to keep making.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/16-f0UOZArS6tmgu-XxRqm13dWM4qizVo0tddxlGCtDY/edit",
						"KTAR, TSMC Arizona $100 billion investment: https://ktar.com/arizona-business/tsmc-arizona-100-billion-invest/5889392/",
						"Slashdot, TSMC to invest additional $100 billion in Arizona: https://news.slashdot.org/story/26/07/16/2158219/tsmc-to-invest-additional-100-billion-in-arizona",
						"KJZZ, TSMC will invest another $100 million in Arizona. How significant will it be for the state?: https://www.kjzz.org/the-show/2026-07-20/tsmc-will-invest-another-100-million-in-arizona-how-significant-will-it-be-for-the-state",
						"9to5Mac, TSMC says it may build 12 Arizona chip plants in total, but be skeptical: https://9to5mac.com/2026/07/16/tsmc-says-it-may-build-12-arizona-chip-plants-in-total-but-be-skeptical/",
						"TechTimes, TSMC lifts Arizona to $265 billion after record quarter: https://www.techtimes.com/articles/320841/20260717/tsmc-lifts-arizona-265-billion-after-record-quarter-four-fabs-target-ai-packaging-bottleneck.htm",
						"BigGo Finance, TSMC Q2 2026 earnings cross-check: https://finance.biggo.com/news/b08d2d2c-156a-4597-9fae-53142441ea1f",
						"MarketScale, Etched targets a $20 billion valuation: https://www.marketscale.com/industries/software-and-technology/etched-targets-a-20-billion-valuation-with-back-to-back-rounds-as-inference-chip-demand-hits-1-billion",
						"PYMNTS, AI chip startup Etched eyes $20 billion valuation: https://www.pymnts.com/news/artificial-intelligence/2026/ai-chip-startup-etched-eyes-20-billion-valuation/",
					},
				},
			},
		},
		{
			Title:   "Brussels Just Ordered Google to Share Android's AI Controls With Rivals. Google Says the Order Is the Privacy Risk.",
			Slug:    "eu-forces-google-share-android-ai-search-data-rivals-2026",
			Date:    "July 21, 2026",
			Tag:     "Policy",
			Summary: "Two weeks after Google exhausted its last appeal on a $4.7 billion antitrust fine, EU regulators hit it with a bigger ask: open Android's AI plumbing to ChatGPT and Claude, and hand competitors the search data Gemini runs on.",
			Related: []Link{
				{
					Title: "Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million.",
					Slug:  "openai-gpt-5-6-general-availability-government-gate-precedent-2026",
				},
				{
					Title: "Two AI Superpowers Just Built Rival Alliances. One Country Already Joined Both.",
					Slug:  "two-ai-superpowers-rival-alliances-one-country-joined-both-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`On July 16, the European Commission adopted two binding decisions under the Digital Markets Act that go further than anything it has ordered Google to do since the law took effect. The first forces Google to open roughly a dozen system-level Android capabilities — voice activation on par with "Hey Google," the ability to delegate tasks like booking a ride or drafting an email, access to phone sensors and background processes, and in-call chat suggestions — to competing AI assistants, including ChatGPT and Claude, on the same terms Google currently reserves for its own Gemini. The second requires Google to share anonymized search query, click, and ranking data with rival search engines and AI chatbots that want to build search-grade answers without indexing the web themselves.`,
						`The timeline is deliberately staged. Google has to publish template licensing agreements and sample test data by September 2026. A finalized pricing structure for the search data and the start of actual data access follow in January 2027. The Android changes land later still, timed to the next major Android release and expected sometime around mid-2027 — giving Google roughly a year to rebuild parts of its stack rather than flip a switch overnight. Non-compliance with either decision carries a penalty of up to 10% of Alphabet's global annual revenue, a ceiling that, on recent revenue, would run well past $30 billion.`,
						`Google's response arrived the same day, in a blog post from Kent Walker, the company's president of global affairs, who called the decisions a threat to "vital privacy and security guardrails for millions of Europeans." It's a notable pivot in framing: Google isn't contesting that it dominates these markets — Search holds north of 90% of query volume in the EU, and Android runs roughly 60% of the bloc's phones — it's arguing that opening the vault is itself the harm. The Commission has heard this argument before and rejected it once already this year: an earlier Google proposal to share search data was turned down for anonymizing so aggressively that it stripped out 90 to 100% of unique queries, rendering the dataset close to useless for a rival trying to build a real competing product. The new decision responds directly to that failure, specifying a "multi-layered" anonymization method built with outside privacy experts and aligned with draft joint guidance the Commission is developing with the EU's data protection board — an attempt to thread the needle between the DMA's competition goals and GDPR's privacy floor.`,
						`Henna Virkkunen, the Commission's executive vice-president for tech sovereignty, framed the goal in market terms rather than privacy terms: the measures are meant to give "emerging alternatives to Google Search and Google's AI services, such as Gemini" a real shot at competing, rather than a theoretical one. That's the crux of the disagreement. Google's position is that data-sharing at this scale is a privacy problem no matter how it's anonymized. The Commission's position is that Google has had every opportunity to propose an anonymization method that satisfies both goals, and the one it offered on its own didn't.`,
						`The timing isn't a coincidence, either. Just two weeks before the new decisions, on July 2, the EU's top court, the Court of Justice, dismissed Google and Alphabet's final appeal of a separate, decade-old Android case — confirming a €4.1 billion fine (about $4.7 billion) for using pre-installation deals to lock phone makers into bundling Google Search and Chrome. That fine was already reduced once, from an original €4.34 billion in 2018 down to €4.1 billion by a lower court in 2022; the July 2 ruling closed off Google's last legal avenue to challenge it at all. It's the kind of case that takes eight years to fully resolve — and its resolution appears to have cleared the runway for the Commission to move faster on the newer, AI-era version of the same argument: that whoever controls the default experience on a phone people didn't choose to leave gets to pick who wins the next platform.`,
						`There's more coming. Separate from either decision, the Commission is reportedly preparing a binding compliance order — expected before the end of July — on a different DMA track entirely, covering Google's alleged self-preferencing of Google Shopping, Flights, and Hotels results in Search, plus ongoing Play Store compliance gaps. Reports put the likely fines in the hundreds of millions of euros, a fraction of the Android case but a sign the Commission isn't done opening fronts against the same company in the same month.`,
						`What makes the AI interoperability order different from the decade of Android antitrust litigation that preceded it is the target. The 2018 case was about who gets to be the default browser or search bar on a phone — a fight over defaults in a market Google already controlled. This one is about whether a phone's operating system is allowed to treat one company's AI assistant as more privileged than another's, at a moment when AI assistants, not search bars, are becoming the thing people actually talk to first. Google has a year and a half to build the interoperability the Commission is asking for. Whether "Hey Google" and "Hey ChatGPT" end up on equal footing on the same phone, or whether Google's privacy objections turn into the next decade-long court fight, is the part nobody — regulators or Google — gets to decide alone.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/1Ztef-hf2Zpt8G-YNAB3DKq5ADUoytsC5KGF64K2equw/edit",
						"European Commission, Digital Markets Act, Commission provides guidance to Google for AI interoperability on Android and sharing of Google Search data under the Digital Markets Act: https://digital-markets-act.ec.europa.eu/",
						"Unite.AI, EU Compels Google to Share Android and Search Data With Rival AI Assistants: https://www.unite.ai/",
						"The Register, EU forces Google to share its toys with the other AI and search kids: https://www.theregister.com/",
						"techxplore.com / AFP, EU top court upholds record €4.1bn Google fine: https://techxplore.com/",
						"CNBC, Google loses fight over record $4.7 billion EU antitrust fine: https://www.cnbc.com/",
					},
				},
			},
		},
	}, posts...)
}
