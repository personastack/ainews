package content

func init() {
	posts = append([]Post{
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
