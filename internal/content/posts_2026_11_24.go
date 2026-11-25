package content

func init() {
	posts = append([]Post{
		{
			Title:   "EDITH: The NHS Puts AI in the Second Reader's Chair",
			Slug:    "nhs-edith-ai-breast-screening-trial",
			Date:    "November 24, 2026",
			Tag:     "AI in Healthcare",
			Summary: "The NHS has begun recruiting for the world's largest trial of AI in breast cancer screening. One study arm will test whether AI can take the place of the second human reader.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"Some AI daily digests this week described the UK's EDITH trial as reporting results. The November announcement was about recruitment in Cambridgeshire. The study will follow participants over several years, and its results could change how a national screening programme works.",
					},
				},
				{
					Heading: "What launched",
					Paragraphs: []string{
						"On Nov 23, 2026, NHS England announced that EDITH had started recruiting at Cambridge University Hospitals. The Health Research Authority says nearly 700,000 women will be invited across 30 UK screening centres over the next three years. The trial registry sets a recruitment target of 660,000. The National Institute for Health and Care Research (NIHR) is funding the £10 million trial. The University of Warwick, the University of Cambridge and Cambridge University Hospitals NHS Foundation Trust lead it, working with NHS England.",
					},
				},
				{
					Heading: "The design: one human, one AI",
					Paragraphs: []string{
						"Today, two specialists independently review each mammogram in NHS breast screening, a practice known as double reading. In England, routine invitations go out every three years from age 50 until a woman's 71st birthday. EDITH compares standard double reading with two AI-assisted approaches. In one, AI acts as the second reader. In the other, AI helps decide whether a mammogram needs one specialist or two. At least one human specialist reviews every mammogram. Co-chief investigator Prof Sian Taylor-Phillips said four AI systems had already been tested for accuracy. The trial's main outcomes are cancer detection rate and recall rate (how many women are called back for further tests). Researchers will also use screening and cancer registry records to track cancers diagnosed between screening rounds.",
					},
				},
				{
					Heading: "Why it matters",
					Paragraphs: []string{
						"About 55,000 women are diagnosed with breast cancer in the UK each year, and double reading is labour-intensive in a field short of radiologists. NHS England says a successful trial could free up hundreds of radiologists and other specialists to see patients who need further tests or treatment, potentially shortening waits. Prof Fiona Gilbert, EDITH co-chief investigator, said the technology can let specialist breast radiologists spend more time caring for patients.",
					},
				},
				{
					Heading: "The caution",
					Paragraphs: []string{
						"The risk is not only missed cancers; it is also over-detection. Prof Sian Taylor-Phillips of the University of Warwick said AI could improve early detection and warned against overdiagnosing milder disease that would never have caused symptoms, as reported by BBC News. A tool that finds more cancers is not automatically better if some of them are ones patients would never have needed treated. That is why a large randomised trial with long follow-up matters more than a vendor accuracy claim.",
					},
				},
				{
					Heading: "A closing thought",
					Paragraphs: []string{
						"Most AI headlines this week are about launches and valuations. EDITH is a slow, careful experiment, with about three years of invitations and more time needed for follow-up. Healthcare AI will likely earn trust this way, through trials that report outcomes for hundreds of thousands of real patients. Worth asking of every AI claim: where is the trial?",
					},
				},
				{
					Heading: "Sources",
					Sources: []Source{
						{Label: "NHS England, first women join the NHS breast screening trial", URL: "https://www.england.nhs.uk/2026/11/first-women-join-world-leading-ai-assisted-nhs-breast-screening-trial/"},
						{Label: "NIHR Cambridge BRC, Cambridgeshire recruitment announcement", URL: "https://cambridgebrc.nihr.ac.uk/2026/11/23/cambridgeshire-women-are-first-to-join-world-leading-ai-breast-screening-trial/"},
						{Label: "Health Research Authority, EDITH trial recruitment and design", URL: "https://www.hra.nhs.uk/about-us/news-updates/nhs-launches-the-worlds-largest-ai-breast-cancer-screening-trial/"},
						{Label: "ISRCTN EDITH trial register", URL: "https://www.isrctn.com/ISRCTN81384017"},
						{Label: "BBC News, Can AI spot breast cancer earlier?", URL: "https://www.bbc.co.uk/news/articles/cmpdwnw3nvzjo"},
						{Label: "Clinical Trials Arena: UK's NIHR backs largest global AI breast cancer screening trial, Nov 2026", URL: "https://www.clinicaltrialsarena.com/news/nihr-backs-ai-breast-cancer-screening-trial-edith-2/"},
						{Label: "AI Weekly: NHS EDITH trial puts AI into 700,000 breast screenings", URL: "https://aiweekly.co/alerts/nhs-edith-trial-puts-ai-into-700000-breast-screenings"},
						{Label: "The Neuron: NHS AI Breast Screening Trial", URL: "https://www.theneuron.ai/news/nhs-ai-breast-cancer-screening-trial/"},
					},
				},
			},
		},
		{
			Title:   "The $63 Billion Chip Lease: How AI's Biggest Compute Bets Are Being Financed Like Aircraft",
			Slug:    "anthropic-63-billion-chip-lease",
			Date:    "November 24, 2026",
			Tag:     "Industry",
			Summary: "A $63 billion financing package lets Anthropic lease AI chips while Broadcom helps back the debt. The deal shows how the supplier and lenders share the risk of falling hardware values.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"The most expensive part of the AI race is no longer just the chips. It is the paperwork that pays for them. On October 26, Broadcom and four financial partners announced a $63 billion package to buy AI accelerators and networking equipment for Anthropic. It looks less like a tech funding round and more like an aircraft lease.",
					},
				},
				{
					Heading: "How the deal works",
					Paragraphs: []string{
						"Broadcom arranged the package with Bank of America, Blackstone, Citigroup and Morgan Stanley. A new special purpose vehicle (SPV) buys the Google TPUs that Broadcom helps develop and leases them to Anthropic. Lenders are repaid from lease payments, and the hardware is the collateral.",
						"Bank of America, Citi and Morgan Stanley arrange the first-lien portion. Blackstone is the sole provider of the committed portion of the second-lien financing.",
						"Broadcom backstops part of the financing with residual-value support tied to what the chips will be worth later. The Wall Street Journal reported that the banks plan to sell part of their share to bond investors once the chips are delivered.",
						"The chips tie to Anthropic's April 2026 agreement with Broadcom and Google. Broadcom said Anthropic would access about 3.5 GW of next-generation TPU capacity through Broadcom as part of the wider agreement, starting in 2027.",
					},
				},
				{
					Heading: "It was not smooth",
					Paragraphs: []string{
						"The Wall Street Journal reported that Broadcom CEO Hock Tan spent nearly two months seeking more than $40 billion in bank commitments. Rising interest rates and turmoil around neocloud companies stalled the deal until banks got higher rates and Broadcom offered protection against losses if Anthropic defaults.",
					},
				},
				{
					Heading: "A separate number to keep straight",
					Paragraphs: []string{
						"Anthropic's Form S-1 describes a separate note facility. The company may issue up to $42.0 billion of convertible promissory notes to Broadcom or a designated financial partner solely to finance quarterly rent under its 2026 equipment lease arrangement with Broadcom. The $63 billion package finances the SPV's purchase of equipment. The note facility would fund Anthropic's rent payments if needed.",
					},
				},
				{
					Heading: "Why it matters",
					Paragraphs: []string{
						"1. Chips as collateral. Lenders are betting AI hardware keeps its value, which is why Broadcom's residual-value guarantee was needed: a GPU or TPU ages faster than a jet or building.",
						"2. Suppliers financing customers. Broadcom, which sells the chips, also shoulders risk to make the sale happen, echoing circular-financing worries in the industry.",
						"3. Compute is going global and local at once. On August 31, AMD, Cisco and Saudi Arabia's HUMAIN said AMD Instinct MI355X systems were live and outlined a next phase of up to 250 MW starting in 2027. On November 19, Deutsche Telekom said its fund invested €100 million in Mistral's €3 billion Series D, ten days after announcing it would expand its Munich AI data center.",
					},
				},
				{
					Heading: "Closing thought",
					Paragraphs: []string{
						"If the model race is a sprint, the financing race is a mortgage. When chips depreciate faster than the loans, who really holds the risk: the lab, the supplier, the banks, or eventually the bond buyers?",
					},
				},
				{
					Heading: "Sources",
					Sources: []Source{
						{Label: "Broadcom and financial partners, October 26 financing announcement", URL: "https://investors.broadcom.com/node/64786/pdf"},
						{Label: "Broadcom, transaction structure and residual-value support", URL: "https://investors.broadcom.com/static-files/f4bda79d-cc66-4748-919b-40d08c658f2d"},
						{Label: "Data Center Dynamics, coverage of the $63 billion chip financing", URL: "https://www.datacenterdynamics.com/en/news/broadcom-organizes-63bn-deal-with-bank-of-america-blackstone-citigroup-and-morgan-stanley-to-fund-anthropic-chip-purchases/"},
						{Label: "The Wall Street Journal, how Broadcom got the deal done", URL: "https://www.wsj.com/finance/broadcom-ceo-intervened-to-seal-63-billion-anthropic-debt-deal-c807754f"},
						{Label: "Broadcom, April 6 Form 8-K on its compute agreements with Google and Anthropic", URL: "https://investors.broadcom.com/static-files/c906d370-921b-4bc2-bb7b-57877dfcf1ae"},
						{Label: "Anthropic, Form S-1 filed with the SEC", URL: "https://www.sec.gov/Archives/edgar/data/2130357/000162828026067168/anthropicpbc-sx1.htm"},
						{Label: "AMD, Cisco and HUMAIN, August 31 announcement on Saudi AI infrastructure", URL: "https://ir.amd.com/news-events/press-releases/detail/1298/amd-cisco-and-humain-expand-saudi-arabias-ai-infrastructure-as-amd-instinct-systems-go-live"},
						{Label: "Deutsche Telekom, November 19 investment announcement", URL: "https://www.telekom.com/en/newsroom/latest-updates/media-information/2026/11/deutsche-telekom-advances-sovereign-ai-in-germany"},
						{Label: "Deutsche Telekom, November 9 Munich expansion announcement", URL: "https://www.telekom.com/en/newsroom/latest-updates/media-information/2026/11/ai-demand-rises-deutsche-telekom-expands-ai-factory-in-munich"},
						{Label: "Mistral AI, Series D announcement", URL: "https://mistral.ai/news/mistral-makes-sovereign-open-weight-ai-to-frontier/"},
					},
				},
			},
			Related: []Link{
				{Slug: "anthropic-record-ipo-83-dollars", Title: "Anthropic Opens at $111.50, Closes at $97.17 After Record $83 IPO"},
				{Slug: "anthropic-volta-bitdeer-10-billion-compute-deal-norway-2026", Title: "Anthropic Just Bet $10 Billion on a Compute Company That Didn't Exist Six Months Ago"},
			},
		},
	}, posts...)
}
