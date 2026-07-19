package content

func init() {
	posts = append([]Post{
		{
			Title:   "Two AI Superpowers Just Built Rival Alliances. One Country Already Joined Both.",
			Slug:    "two-ai-superpowers-rival-alliances-one-country-joined-both-2026",
			Date:    "July 19, 2026",
			Tag:     "Policy",
			Summary: "China's 29-nation World AI Cooperation Organization is Beijing's answer to the US-led Pax Silica supply-chain bloc, but Kazakhstan has already joined both, showing that many governments may treat AI governance rivalry as a menu rather than a fork in the road.",
			Related: []Link{
				{
					Title: "China Isn't Banning AI Agents. It's Banning the Ones That Pretend to Love You.",
					Slug:  "china-anthropomorphic-ai-interaction-rules-companion-shutdown-2026",
				},
				{
					Title: "Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million.",
					Slug:  "openai-gpt-5-6-general-availability-government-gate-precedent-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"For most of 2026, the fight over who gets to set the rules for artificial intelligence has played out in speeches, white papers, and export-control filings. This month it got an org chart.",
						"On July 16, China formally stood up the World AI Cooperation Organization (WAICO), an intergovernmental body headquartered in Shanghai with 29 founding member states, including Indonesia, Brazil, Malaysia, South Africa, Senegal, Pakistan, and Russia. President Xi Jinping unveiled it the next day at the opening of the World AI Conference, telling the audience that \"AI development should not be a solo performance by a single country, but a symphony of international cooperation\" -- a line aimed, without much subtlety, at Washington.",
						"WAICO isn't the first mover here. It's a response. In December 2025, the US State Department launched Pax Silica, a coalition built around securing the AI supply chain -- chips, critical minerals, energy, and advanced manufacturing -- among trusted partners. After a second summit in Washington on June 25 and 26, the State Department confirmed Pax Silica had grown to 24 signatory countries, adding Argentina, Chile, Costa Rica, El Salvador, Germany, Greece, the Netherlands, Panama, Kazakhstan, and the European Union as a bloc. Under Secretary Jacob Helberg also announced a Stanford-designed advanced-manufacturing curriculum for partner economies and a pilot program with Panama to digitize customs tracking for AI hardware shipments.",
						"So now there are two clubs, and they're built on different premises. Pax Silica gates membership around trust and alignment -- its logic is that if you want access to US-designed chips, cloud platforms, and secure supply chains, you sign on to a shared security posture first. WAICO does the opposite: no values test, no regime-type screen, open to any sovereign state that wants in. Xi has pitched it explicitly at the Global South, promising 5,000 AI training slots for developing nations over five years, access to a Chinese meteorological AI system for 30 countries, and joint \"AI application centers\" built with ASEAN, the Arab League, the African Union, the Latin American and Caribbean bloc CELAC, the Shanghai Cooperation Organisation, and BRICS.",
						"The pitch to a country without a domestic chip industry or a seat at the table in Washington is straightforward: cheaper, open-weight Chinese models, homegrown processors, and infrastructure that doesn't depend on the most advanced US export-controlled hardware. For nations priced out of the frontier-compute race entirely, that's not an abstract governance debate -- it's the difference between deploying AI now or waiting.",
						"Washington, for its part, hasn't issued a direct response to WAICO's launch, though US officials and companies including Anthropic have spent recent months accusing Chinese AI labs of extracting proprietary techniques from Western models, allegations Beijing calls groundless. That friction sits underneath both alliances: each is as much about which supply chain and standards regime a country ends up dependent on as it is about safety frameworks or research cooperation.",
						"But the tidy story -- the world splitting into an American bloc and a Chinese one -- is already breaking down in the fine print. Kazakhstan signed the Pax Silica Declaration in June. It's also one of WAICO's 29 founding members. Nothing in either organization's charter forces an exclusive choice, and for a mid-sized economy trying to build out AI infrastructure, there's no real incentive to make one. Take the training slots, the chip access, and the investment from whichever side offers them, and worry about the alignment questions later.",
						"That's probably the more useful way to read this moment than \"two blocs.\" What China and the US have actually built are two competing service offerings for the same customers -- most of the world's governments, which have no frontier AI industry of their own and every reason to extract maximum benefit from a rivalry they didn't start. Whether that hedging strategy survives the next round of export controls, or the next accusation of IP theft, is the real question hanging over both organizations. For now, the fastest-growing membership category in global AI governance isn't Team Washington or Team Beijing. It's both.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/18gLxfTBoUixPEIhBkpInFqYKv9zhrJTt29uJhJybAb4/edit",
						"Al Jazeera coverage of China's WAICO launch and AI governance pitch: https://www.aljazeera.com/news/2026/7/17/chinas-xi-jinping-launches-new-ai-alliance-what-is-it",
						"Xinhua coverage of Xi's World AI Conference remarks and WAICO unveiling: https://english.news.cn/20260717/ce32e833ab5d47f883ad44e1f73cb634/c.html",
						"Chinese government/Xinhua release on the 29-country WAICO founding agreement: https://english.www.gov.cn/news/202607/17/content_WS6a59a226c6d00ca5f9a0c432.html",
						"U.S. Department of State outcomes from the second Pax Silica Summit: https://www.state.gov/releases/office-of-the-spokesperson/2026/06/outcomes-of-the-second-pax-silica-summit",
						"U.S. Department of State Pax Silica declaration page: https://www.state.gov/pax-silica",
						"Astana Times coverage of Kazakhstan joining Pax Silica: https://astanatimes.com/2026/06/kazakhstan-joins-us-led-pax-silica-initiative-to-advance-ai-economy/",
						"Brownstein analysis of the 2026 Pax Silica Summit expansion: https://www.bhfs.com/insight/state-department-expands-pax-silica-initiative-at-2026-summit/",
					},
				},
			},
		},
	}, posts...)
}
