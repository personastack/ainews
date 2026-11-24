package content

func init() {
	posts = append([]Post{
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
