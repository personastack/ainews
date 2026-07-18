package content

func init() {
	posts = append([]Post{
		{
			Title:   "Apple Says OpenAI Turned Job Interviews Into a Trade Secrets Pipeline",
			Slug:    "apple-openai-trade-secret-lawsuit-io-products-2026",
			Date:    "July 18, 2026",
			Tag:     "Legal",
			Summary: "A 41-page lawsuit accuses OpenAI's hardware chief of asking Apple engineers to bring prototypes to job interviews, a former Apple engineer of quietly reading company files after he'd already left, and Jony Ive's design studio of borrowing a metal-finishing trick it never had permission to use.",
			Related: []Link{
				{
					Title: "Nvidia's Roadmap Just Hit the Reticle Limit",
					Slug:  "nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"On July 10, Apple filed suit against OpenAI in the U.S. District Court for the Northern District of California. The case, Apple Inc. v. Liu (5:26-cv-07078), reads less like a routine corporate dispute and more like a workplace thriller. Apple's lawyers argue the theft wasn't the work of one rogue employee cashing out on the way to a new job — it was happening, in the company's words, \"at every level, from members of its Technical Staff to its Chief Hardware Officer, and in coordination with business partners.\"",
						"That Chief Hardware Officer is Tang Tan, a former Apple vice president who spent years running product design for the iPhone and Apple Watch before joining OpenAI to lead its push into consumer hardware. Apple's complaint alleges Tan didn't just recruit Apple engineers — he allegedly turned the interview process itself into an extraction point. According to the filing, Tan directed job candidates who were still employed at Apple to bring \"actual parts,\" CAD files, and prototypes to their interviews for what the complaint calls \"show and tell\" sessions, and used Apple's own internal project code names while doing the recruiting. One candidate, the filing says, was surprised to learn Apple hardware was even allowed to leave the building.",
						"The second thread involves Chang Liu, a former Apple senior systems electrical engineer who spent eight years at the company before moving to OpenAI. Apple alleges Liu exploited an authentication flaw to reach Apple's internal network storage through a colleague's work computer after he'd left the company, texting \"LOL, I found out I can access the [network storage]\" and later noting \"I still have another computer\" — which Apple's lawyers read as a plan to keep the access going. The complaint also says Liu never returned his Apple-issued laptop and used it to download confidential technical documents, including engineering presentations and specifications tied to unannounced products.",
						"The third and strangest thread points at io Products, the hardware design studio OpenAI acquired for $6.5 billion in 2025, co-founded by former Apple design chief Jony Ive. Apple alleges io used a \"secret, proprietary industrial design technique\" related to metal finishing by misleading one of Apple's own manufacturing partners into believing OpenAI had Apple's permission to use it — and that OpenAI approached suppliers using confidential Apple terminology it shouldn't have had. If true, it suggests the leakage wasn't limited to code or blueprints; it reached into the physical supply chain that actually builds Apple's products.",
						"Apple says none of this came out of nowhere. According to the complaint, the company first raised its concerns directly with OpenAI back in February 2026 and didn't get a response it considered adequate, which is part of why it's now suing rather than negotiating privately. The scale Apple cites is also striking context: more than 400 former Apple employees are now working at OpenAI, a flow of talent that has made Cupertino noticeably nervous as the AI industry's hardware ambitions have grown.",
						"OpenAI has pushed back, saying it takes the lawsuit seriously but believes the allegations lack merit. The company's position leans on a real legal distinction: trade secret law generally protects specific proprietary information, not the general skills and knowledge an employee carries in their head from one job to the next. California in particular has strong protections for worker mobility and a near-total ban on non-compete agreements, which is part of why Silicon Valley job-hopping is normally treated as healthy competition rather than theft.",
						"What makes this case worth watching isn't just the gossip-column details, though there are plenty of those. It's what it signals about where the AI industry's real bottleneck has shifted. For the past few years, the story was about who had the best model. Increasingly, it's about who can actually build a device people want to hold — and that requires exactly the kind of manufacturing know-how, supplier relationships, and industrial design instincts that companies like Apple have spent decades protecting. If Apple wins, it could make the next generation of AI hardware startups think twice about how aggressively they recruit from incumbents. If OpenAI wins, or if the case settles quietly, it will reinforce that in California, talent — and everything it remembers — is fair game.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/13cABBIhzEUCUabPvpJzi9-U2p-A-HzDH_bnwnrcbj34/edit",
						"CourtListener docket for Apple Inc. v. Liu, 5:26-cv-07078: https://www.courtlistener.com/docket/73602437/apple-inc-v-liu/",
						"TechCrunch, Apple sues OpenAI over alleged trade secret theft: https://techcrunch.com/2026/07/10/apple-sues-openai-over-alleged-trade-secret-theft/",
						"TechCrunch, the wildest allegations in Apple's trade secrets lawsuit against OpenAI: https://techcrunch.com/2026/07/13/the-wildest-allegations-in-apples-trade-secrets-lawsuit-against-openai/",
					},
				},
			},
		},
	}, posts...)
}
