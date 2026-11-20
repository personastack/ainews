package content

func init() {
	posts = append([]Post{
		{
			Title:   "Anthropic Opens at $111.50, Closes at $97.17 After Record $83 IPO",
			Slug:    "anthropic-record-ipo-83-dollars",
			Date:    "November 19, 2026",
			Tag:     "Industry",
			Summary: "Anthropic (Nasdaq: ANT) opened its first day at $111.50, about 34% above its $83 offer price, then closed at $97.17, up about 17%. The IPO closed on Nov. 20 with a record $93.5 billion in gross proceeds, including the full underwriters' option.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"Anthropic is a public company. Shares of ANT opened at $111.50 on Thursday, Nov. 19, according to Nasdaq's IPO recap. The opening price was $28.50, or roughly 34%, above the $83 offer price set the night before.",
						"The pop did not hold. According to the Associated Press, ANT closed its first day at $97.17, about 17% above the IPO price and about 13% below the opening trade, as the broader market turned lower after an early rally.",
					},
				},
				{
					Heading: "What happened",
					Paragraphs: []string{
						"Anthropic priced above its marketed range of $70 to $78, at $83. At $83 the company's fully diluted valuation works out to about $1.5 trillion. Reports put its fully diluted valuation near $2 trillion at the opening trade; on the prospectus share count alone (undiluted), the open implied roughly $1.47 trillion. Those figures are for the opening price, not the close.",
						"Anthropic announced on November 20 that the offering has closed. The underwriters exercised in full their option to buy an additional 147 million shares, bringing the total to 1.127 billion Class A shares sold at $83 each. That is roughly $93.5 billion in gross proceeds, of which about $91.6 billion went to Anthropic. The SEC final prospectus covered the base deal of 980 million shares, or about $81.34 billion; the option shares are what lifted the total.",
						"Separately, Anthropic sold 313,253,007 Class N shares at $83 to AMD, Broadcom, Google, Micron, Microsoft, NVIDIA and TSMC Partners for about $26.0 billion. That sale is not part of the public offering figure; combined gross is about $119.5 billion before fees and expenses.",
					},
				},
				{
					Heading: "The IPO record",
					Paragraphs: []string{
						"Anthropic's $93.5 billion overtakes SpaceX's previous IPO record of about $86.25 billion. Both figures are gross public offering proceeds and include the full underwriters' options. SpaceX still leads on valuation, at about $1.77 trillion at pricing, versus Anthropic's roughly $1.5 trillion.",
					},
				},
				{
					Heading: "Who can buy, who can sell",
					Paragraphs: []string{
						"Up to 20% of the shares were reserved for U.S. retail investors through brokerages such as Fidelity, Robinhood, Charles Schwab, SoFi and E*TRADE, per Fox Business. The same report says up to 315 million shares held by non-founder insiders became eligible for sale immediately, while founders and major holders face staggered restrictions of up to a year, and that founders keep 50.1% of voting power through a special share class. We have not independently confirmed those details against the prospectus.",
					},
				},
				{
					Heading: "What to watch",
					Paragraphs: []string{
						"1. The weeks after the debut. A 17% close above $83 says demand outran the offer price, but the slide from the $111.50 open shows how quickly first-day enthusiasm can fade. Where the stock settles will show whether the market thinks $1.5 trillion was fair or a floor.",
						"2. Insider supply. Early eligibility for hundreds of millions of shares is a classic source of pressure after a hot debut.",
						"3. The AI financing cycle. A listing this size gives Anthropic cash and a public currency for compute commitments and acquisitions, and sets a benchmark other frontier labs will be measured against.",
						"The takeaway: a first-day pop, and a partial give-back, is a headline, not a verdict. The real test for ANT is whether revenue growth justifies a valuation that still rivals the largest companies on earth.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Associated Press, ANT first-day close ($97.17, about 17% above the $83 IPO price): https://apnews.com/article/b6a74b8949021207dc162fcc876c4649",
						"Nasdaq's Nov. 19 IPO recap ($111.50 opening price) https://www.nasdaq.com/newsroom/anthropic-marks-ipo-at-nasdaq",
						"Anthropic closing announcement and full exercise of the underwriters' option https://www.anthropic.com/news/ipo-close",
						"Anthropic final prospectus for the base offering https://www.sec.gov/Archives/edgar/data/2130357/000162828026074944/anthropicpbc-424b4.htm",
						"SpaceX Form 10-Q, IPO closing and full exercise of the underwriters' option https://www.sec.gov/Archives/edgar/data/1181412/000162828026052535/spcx-20260630.htm",
						"CNN live coverage https://www.cnn.com/2026/11/19/business/live-news/anthropic-ipo",
						"Fox Business https://www.foxbusiness.com/markets/anthropic-begins-trading-nasdaq-record-breaking-ipo",
						"CNBC live updates https://www.cnbc.com/2026/11/19/anthropic-ipo-ant-live-updates.html",
						"TipRanks https://www.tipranks.com/news/anthropic-stock-opens-34-higher-as-early-nasdaq-rally-fades",
					},
				},
			},
			Related: []Link{
				{Title: "OpenAI's GPT-6.2 Cuts Astra's Price 20% and Claims Wins Over Claude. The Independent Scores Tell a Quieter Story.", Slug: "openai-gpt-6-2-astra-sol-price-cut-benchmarks-2026"},
				{Title: "Nvidia Says Open Source Stays After Jensen Huang's Exit. The Hugging Face Deal Is the Test.", Slug: "nvidia-open-source-ai-strategy-after-jensen-huang-exit-2026"},
			},
		},
	}, posts...)
}
