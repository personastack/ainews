package content

func init() {
	posts = append([]Post{
		{
			Title:   "The EU's AI Act Starts Enforcing Today. The Part Companies Feared Most Just Got Delayed to 2027.",
			Slug:    "eu-ai-act-enforcement-begins-high-risk-delayed-2027",
			Date:    "August 2, 2026",
			Tag:     "Policy",
			Summary: "Brussels begins policing chatbot disclosures, deepfake labels, and foundation-model transparency this week - but a last-minute simplification deal quietly bought the toughest high-risk rules two more years.",
			Related: []Link{
				{
					Title: "OpenAI's Model Broke Into Hugging Face. Now 1,178 AI Workers - Including OpenAI's Own - Want Washington to Slow the Whole Race Down.",
					Slug:  "openai-anthropic-google-meta-1178-workers-pacing-mechanism-letter-2026",
				},
				{
					Title: "Jensen Huang's First Tweet Ever Wasn't About Chips. It Was a Warning to Washington.",
					Slug:  "nvidia-jensen-huang-open-weights-letter-distillation-2026",
				},
				{
					Title: "Nvidia Built a Coalition to Stop Rogue AI Agents. The Labs Whose Agents Went Rogue Didn't Join.",
					Slug:  "nvidia-open-secure-ai-alliance-openai-anthropic-google-absent-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`Today is the date European regulators have been circling on their calendars since the AI Act became law in 2024. As of August 2, 2026, the European Commission's AI Office and national regulators across the bloc can formally enforce it, and a new wave of transparency obligations lands alongside them. Chatbots and other interactive AI systems now have to tell users they're talking to a machine. Deepfakes have to be labeled. AI-generated or altered images, video, and audio need machine-readable marks baked in so platforms and detection tools can identify them. The Commission has also switched on a public complaints portal and a whistleblower tool, giving anyone - a competitor, a journalist, an employee - a formal channel to flag a violation.`,
						`The obligations landing today aren't limited to consumer-facing labeling. Providers of general-purpose AI models - the foundation models behind everything from chatbots to coding assistants - are now subject to Article 53 of the Act, which sets baseline documentation and copyright-compliance duties for every GPAI model and layers on extra scrutiny for models the Commission judges to carry "systemic risk," a category aimed squarely at the handful of frontier labs training the largest systems. More than 180 organizations have already signed onto the Commission-backed transparency code of practice, a voluntary framework designed as a soft landing into binding rules on AI-generated content. Enforcement authority is split: the AI Office in Brussels handles GPAI providers and oversight of certain AI systems directly, while national regulators keep jurisdiction over sensitive domains like law enforcement, border management, courts, and financial services. National AI regulatory sandboxes, meanwhile, no longer land today: the Digital Omnibus deal moved that deadline to August 2, 2027.`,
					},
				},
				{
					Heading: "The fines are real money",
					Paragraphs: []string{
						`Under Article 99, deploying a prohibited AI practice - the Act's short list of banned uses, like social scoring or manipulative subliminal techniques - can cost a company up to EUR35 million or 7% of global annual turnover, whichever is larger. Breaching most operator obligations, including transparency duties under Article 50, tops out at EUR15 million or 3% of turnover. Supplying incorrect, incomplete, or misleading information to notified bodies or national authorities can draw penalties up to EUR7.5 million or 1% of turnover. For a large AI company, the percentage-of-turnover figures dwarf the flat euro amounts.`,
					},
				},
				{
					Heading: "The deadline that moved",
					Paragraphs: []string{
						`But the headline date obscures a quieter story that unfolded over the past three months. The original 2024 text of the AI Act set today as the deadline not just for transparency and GPAI rules, but for the much heavier compliance regime covering "high-risk" AI systems - the category that includes hiring algorithms, credit-scoring tools, biometric identification, and AI embedded in medical devices or critical infrastructure. That's the part of the law that actually forces companies to run conformity assessments, produce technical documentation, and register systems in an EU database before deploying them. It's also the part industry groups lobbied hardest against, arguing the compliance burden was unworkable on the original timeline.`,
						`They got their extension. On May 7, 2026, the Council presidency and European Parliament negotiators reached a provisional deal under the EU's Digital Omnibus package to simplify parts of the Act. The Council gave final approval on June 29. The agreement pushes the compliance deadline for standalone high-risk systems under Annex III of the Act to December 2, 2027, and gives high-risk AI embedded in already regulated products - think medical devices or industrial machinery - until August 2, 2028. It also softens the AI-content watermarking requirement for systems already on the market before today, giving providers until December 2, 2026, to add transparency solutions retroactively, though anything launched from today onward has to comply immediately. In exchange for the breathing room, the omnibus deal added new prohibitions on AI-generated non-consensual intimate imagery and child sexual abuse material, and carved out lighter-touch accommodations for small and mid-cap companies.`,
					},
				},
				{
					Heading: "A narrower launch than advertised",
					Paragraphs: []string{
						`The net effect is a regulatory regime that goes live today in name, but arrives narrower than the law's authors originally intended. The rules that touch the biggest AI labs directly - GPAI documentation, systemic-risk scrutiny, chatbot and deepfake disclosure, and a functioning enforcement and penalty apparatus - are active as of this week. The rules that would have forced a much broader swath of European businesses to formally certify AI systems used in hiring, lending, and healthcare now have until the end of 2027 to comply, a concession won only weeks before it would have otherwise bitten. Whether that trade counts as pragmatic sequencing or a retreat under lobbying pressure will likely become one of the more contested storylines in AI policy over the next eighteen months, especially as the AI Office starts publishing its first enforcement actions under the parts of the law that did not get a reprieve.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"European Commission, Commission starts enforcing AI Act rules and new transparency requirements on 2 August: https://digital-strategy.ec.europa.eu/en/news/commission-starts-enforcing-ai-act-rules-and-new-transparency-requirements-2-august",
						"Council of the EU, Artificial Intelligence: Council and Parliament agree to simplify and streamline rules: https://www.consilium.europa.eu/en/press/press-releases/2026/05/07/artificial-intelligence-council-and-parliament-agree-to-simplify-and-streamline-rules/",
						"Council of the EU, Artificial Intelligence: Council gives final green light to simplify and streamline rules: https://www.consilium.europa.eu/en/press/press-releases/2026/06/29/artificial-intelligence-council-gives-final-green-light-to-simplify-and-streamline-rules/",
						"Stibbe, AI Act Reloaded? What the Latest AI Act Changes Mean in Practice: https://www.stibbe.com/publications-and-insights/ai-act-reloaded-what-the-latest-ai-act-changes-mean-in-practice",
						"EU Artificial Intelligence Act reference site, Article 99: Penalties: https://artificialintelligenceact.eu/article/99/",
					},
				},
			},
		},
		{
			Title:   "Unitree Is Going Public With Real Revenue. Figure AI Is Worth $39 Billion Without Any.",
			Slug:    "unitree-ipo-china-humanoid-robotics-boom-figure-ai-2026",
			Date:    "August 2, 2026",
			Tag:     "Robotics",
			Summary: "This week, China's best-selling humanoid robot maker opens IPO subscriptions backed by audited books and real revenue growth. Its Silicon Valley rival is worth four times as much on numbers nobody outside the company has ever seen.",
			Related: []Link{
				{
					Title: "Language's Frontier Is Locking Down. Robotics' Frontier Just Went Open.",
					Slug:  "nvidia-cosmos-3-open-physical-ai-world-model-2026",
				},
				{
					Title: "The Chip That Stays Home: Inside China's Race to Build Robotics AI Hardware",
					Slug:  "the-chip-that-stays-home-inside-chinas-race-to-build-robotics-ai-hardware",
				},
				{
					Title: "MolmoAct 2: AI2's Breakthrough in Physical AI and Robot Task Performance",
					Slug:  "molmoact-2-ai2-robotics-action-reasoning-breakthrough-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`On August 5, Unitree Robotics opens its preliminary price inquiry on the Shanghai Stock Exchange's STAR Market. Subscriptions open five days later, on August 10. When the shares start trading, Unitree becomes the first humanoid robot maker anywhere to list on a public exchange - and the first time investors anywhere get an audited look at what it actually costs, and earns, to build and sell humanoid robots at scale.`,
						`The numbers behind that listing are the real story. Unitree plans to raise about 4.2 billion yuan (roughly $618 million) by issuing just over 40 million new shares, 10 percent of its enlarged share capital. First-half 2026 revenue is projected between 1.05 and 1.13 billion yuan (about $155 to $166 million), up 35.6 to 45.4 percent year over year. In 2025, the company shipped more than 5,500 humanoid units, more than any other manufacturer in the world, and China-built quadruped robots - Unitree's other product line - captured close to 70 percent of global sales in the first half of this year. Founder Wang Xingxing will keep 65.31 percent of voting control after the float. Regulators approved the listing in 104 days, the fastest full review in STAR Market history, which one Xinhua report framed as recognition of embodied intelligence's strategic importance to Beijing.`,
					},
				},
				{
					Heading: "A very crowded moment to go public",
					Paragraphs: []string{
						`Unitree isn't stepping into a quiet market. Five weeks before its subscription window opened, two other Chinese humanoid makers hit the same valuation on the same day. On June 29, Shenzhen-based AI² Robotics (unrelated to Seattle's Allen Institute for AI) disclosed a $735 million round valuing it at $2.8 billion, while X Square Robots closed the fourth round in a Series C run that landed it at the identical $2.8 billion mark. AI²'s wheeled AlphaBot line - 34-plus degrees of freedom, a wheeled base instead of legs to simplify both engineering and the path through public-space safety approvals - is backed by a genuinely unusual investor list: state funds like the National Small and Medium Enterprises Development Fund alongside Moutai Group, the liquor conglomerate, and Sino Biopharmaceutical.`,
						`Unitree isn't even the only one racing for a stock ticker. LimX Dynamics raised $200 million in a pre-IPO round in July at a $2.21 billion valuation, explicitly positioning for its own public listing. Ant Group, meanwhile, has made a dozen robotics investments since the start of 2025 - Unitree and Galaxea among them - and added humanoid startup Zeroth to that list in early July with a $73.6 million round. "Listing is a must," one Chinese humanoid executive told CNBC of the sudden IPO rush, a sentiment that reads less like ambition and more like a description of how competitive the fundraising environment has become.`,
						`The capital math around all of this is large enough to be its own story. Global robotics startups had raised $18.8 billion by late June of this year, according to Crunchbase News - already ahead of the $15 billion the entire sector raised in all of 2025, and past the prior high-water mark of $14.1 billion set in 2021's venture boom, with five months of the year still left to run. China's humanoid robot output alone is projected to top 100,000 units in 2026. None of that includes whatever Unitree raises next week.`,
					},
				},
				{
					Heading: "The valuation nobody has to prove",
					Paragraphs: []string{
						`Now set that against Figure AI, the Bay Area humanoid maker still widely treated as the sector's bellwether. Figure's valuation was set at $39 billion in a September 2025 Series C that pulled in investors including Nvidia, Microsoft, Brookfield Asset Management, the OpenAI Startup Fund, and Jeff Bezos personally. Figure has never disclosed revenue. Third-party estimates put it in the low tens of millions at best, which would make $39 billion a valuation built almost entirely on a bet about what Figure's robots might eventually do, rather than a record of what they currently earn. That is a perfectly normal way for a private, venture-backed company to be valued. It is also a bet nobody outside Figure's boardroom can actually check.`,
						`Unitree is about to lose that shield. Once its shares trade on the STAR Market, its quarterly filings become public record - real shipment volumes, real margins, real R&D spend, all subject to the same disclosure rules that apply to any other listed Chinese company. That is a meaningfully higher bar than a press release announcing a new valuation, and it is a bar Figure, Apptronik, Agility Robotics, and the rest of the venture-funded humanoid field in the US have not had to clear.`,
					},
				},
				{
					Heading: "What to watch",
					Paragraphs: []string{
						`The interesting question isn't which humanoid maker builds the better robot. It's what happens to the sector's private valuations once one of its biggest players has to publish real numbers every quarter. If Unitree's disclosed margins turn out to be thin - and hardware margins usually are, at this stage of any category - that's a data point every investor pricing Figure, Apptronik, or the next Chinese contender at billions of dollars per funding round will have to reckon with, whether they want to or not. Beijing sped a 104-day regulatory review specifically to get a humanoid maker onto public markets first. Whether that turns into a genuine transparency advantage for Chinese robotics, or whether Silicon Valley's humanoid bets simply keep raising private rounds and skip the public scoreboard altogether, is the thing worth watching once the trading bell actually rings on August 10.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Xinhua, Unitree Robotics wins approval for STAR Market IPO: https://english.news.cn/20260731/20953e97618041788ceb70d7ee451a02/c.html",
						"Caixin Global, Unitree Robotics Wins Approval for $618 Million STAR Market IPO: https://www.caixinglobal.com/2026-07-03/unitree-robotics-wins-approval-for-618-million-star-market-ipo-102460136.html",
						"SiliconANGLE, Chinese robotics outfits AI2 Robotics and X Square Robots secure funding at $2.8B valuation: https://siliconangle.com/2026/06/29/chinese-robotics-outfits-ai2-robotics-x-square-robots-secure-funding-2-8b-valuation/",
						"Crunchbase News, Robotics startup venture funding surges in 2026: https://news.crunchbase.com/robotics/startup-venture-funding-surges-2026-data/",
						"CNBC, Chinese humanoid startups race toward IPOs: https://www.cnbc.com/2026/07/13/chinese-humanoid-startups-ipo-limx-unitree.html",
						"Sacra, Figure AI company revenue estimates: https://sacra.com/c/figure-ai/",
						"AI Business, Chinese tech vendors converge on humanoid robotics and embodied AI: https://aibusiness.com/robotics/chinese-tech-vendors-converge-humanoid-robotics-embodied-ai",
					},
				},
			},
		},
	}, posts...)
}
