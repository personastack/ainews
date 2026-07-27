package content

func init() {
	posts = append([]Post{
		{
			Title:   "OpenAI Wants a $500 Billion Data Center. It Needed Nvidia to Cosign the Lease.",
			Slug:    "nvidia-openai-ohio-datacenter-250b-backstop-circular-financing-2026",
			Date:    "July 27, 2026",
			Tag:     "Business",
			Summary: "Nvidia is reportedly negotiating up to $250 billion in lease and construction guarantees, plus a separate $350 billion chip-financing package, to get OpenAI's next data center campus built on the site of a decommissioned Ohio uranium plant. The arrangement exists because OpenAI can't get an investment-grade credit rating on its own — and it's exactly the kind of circular financing that bond-market regulators have started warning about.",
			Related: []Link{
				{
					Title: "Nvidia's Roadmap Just Hit the Reticle Limit",
					Slug:  "nvidia-rubin-ultra-dual-die-redesign-reticle-limit-2026",
				},
				{
					Title: "Anthropic Is Racing OpenAI to Wall Street. Its Own Revenue Number May Not Survive the Trip.",
					Slug:  "anthropic-ipo-openai-race-revenue-accounting-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`Nvidia is reportedly negotiating up to $250 billion in lease and construction guarantees, plus a separate $350 billion chip-financing package, to get OpenAI's next data center campus built on the site of a decommissioned Ohio uranium plant. The arrangement exists because OpenAI can't get an investment-grade credit rating on its own — and it's exactly the kind of circular financing that bond-market regulators have started warning about.`,
						`According to a Wall Street Journal report picked up by outlets including Yahoo Finance, Tom's Hardware, HotHardware, and ZeroHedge, Nvidia is in talks to guarantee roughly $250 billion covering the lease and construction financing for a 10-gigawatt data center campus that SB Energy — SoftBank's power subsidiary — is developing in Piketon, Ohio, about 50 miles south of Columbus. Separately, Nvidia is discussing financing up to $350 billion more to cover the cost of the chips that would fill the buildings. Combined with construction, the full project is projected to exceed $500 billion, which would make it the largest data center commitment announced to date. Nothing is signed. Terms haven't been finalized, and multiple outlets note the arrangement could still fall apart.`,
					},
				},
				{
					Heading: "A uranium site becomes an AI campus",
					Paragraphs: []string{
						`The site itself carries a strange kind of symbolism: it's the former home of a Cold War-era uranium enrichment facility, now being retrofitted to house the power-hungry GPU clusters of the AI boom. The first phase, targeted for 2028, is expected to bring roughly 800 megawatts online, on the way to the full 10-gigawatt build-out. Power for the site is tangled up in trade policy, too — Japan has committed $33 billion to a natural-gas power project on the same federal land, part of its broader pledge to invest in U.S. infrastructure in exchange for lower tariffs, with Japan recovering its investment first and the U.S. receiving 90 percent of proceeds after that.`,
					},
				},
				{
					Heading: "Why Nvidia's guarantee matters",
					Paragraphs: []string{
						`What makes the Nvidia guarantee necessary is simpler than the geopolitics: OpenAI has never turned a profit, and a private company without profits doesn't get an investment-grade credit rating on its own. Without that rating, SB Energy would have to borrow at much less favorable terms to build a $500 billion campus for a single, unprofitable tenant. Nvidia's guarantee is what turns that debt into something bond markets will actually buy at a reasonable price.`,
						`It's also, unavoidably, Nvidia financing demand for its own product. The $350 billion chip package would fund OpenAI's purchase of what reporting describes as hundreds of thousands of next-generation GPUs — GPUs Nvidia manufactures and sells. That's the arrangement critics have started calling circular financing: the chipmaker isn't just selling hardware anymore, it's underwriting the real estate that houses it and bankrolling the purchase of the hardware itself, all to guarantee a customer keeps buying. ZeroHedge's analysis pegs the broader pattern of these vendor-financed AI infrastructure deals at roughly $2 trillion industry-wide, resting on the assumption that OpenAI and Anthropic will eventually turn a profit on more than $1.5 trillion in compute commitments that aren't yet backed by matching revenue.`,
					},
				},
				{
					Heading: "The bust risk",
					Paragraphs: []string{
						`That assumption is exactly what has the Bank for International Settlements on edge. The BIS has warned that "a disappointment in returns could trigger a sudden pullback in financing and turn the capex boom into a protracted investment bust, with potential knock-on effects on financial conditions." It's a warning worth sitting with: the same week this financing news broke, Moonshot AI's 2.8-trillion-parameter Kimi K3 went fully open-weight and free to download — the kind of competitive pressure from cheaper, openly available models that could squeeze the revenue these trillion-dollar bets are counting on.`,
						`For now, the Ohio campus is still a negotiation, not a building. But the shape of the deal already tells you something about where the AI industry's financial plumbing has ended up: the world's most valuable AI startup needs its own chip supplier to cosign the lease.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Yahoo Finance, Nvidia in talks to back OpenAI Ohio data center: https://finance.yahoo.com/technology/ai/articles/nvidia-talks-back-openai-ohio-114515389.html",
						"Tom's Hardware, Nvidia weighs $250 billion guarantee for OpenAI lease: https://www.tomshardware.com/tech-industry/data-centers/nvidia-weighs-250-billion-guarantee-so-openai-can-lease-softbanks-10-gigawatt-ohio-campus",
						"HotHardware, OpenAI seeks $250 billion Nvidia guarantee for Ohio data center: https://hothardware.com/news/openai-seeks-250-billion-nvidia-guarantee-ohio-data-center",
						"ZeroHedge, Nvidia weighs $250 billion backstop for OpenAI data center campus: https://www.zerohedge.com/ai/nvidia-weighs-250-billion-backstop-openais-gargantuan-ohio-data-center-campus",
					},
				},
			},
		},
	}, posts...)
}
