package content

func init() {
	posts = append([]Post{
		{
			Title:   "DeepSeek and Moonshot Are Racing to IPO. Beijing Just Showed Up as the Investor With No Lock-Up.",
			Slug:    "deepseek-moonshot-china-ai-ipo-funding-state-investment-2026",
			Date:    "August 3, 2026",
			Tag:     "Business",
			Summary: "Two Chinese AI labs raised back-to-back mega-rounds this summer, and the fine print reveals more about who really controls frontier AI in China than the valuations do.",
			Related: []Link{
				{
					Title: "Anthropic Is Racing OpenAI to Wall Street. Its Own Revenue Number May Not Survive the Trip.",
					Slug:  "anthropic-ipo-openai-race-revenue-accounting-2026",
				},
				{
					Title: "OpenAI Wants a $500 Billion Data Center. It Needed Nvidia to Cosign the Lease.",
					Slug:  "nvidia-openai-ohio-datacenter-250b-backstop-circular-financing-2026",
				},
				{
					Title: "The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say.",
					Slug:  "white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`Six weeks. That's roughly how long it took DeepSeek to go from closing a $7.4 billion funding round to opening talks on another one - this time at a valuation 37% higher.`,
						`The Hangzhou-based lab raised its first-ever external financing in June 2026, pulling in about $7.4 billion (50 billion yuan) at a post-money valuation of roughly $52 billion. By mid-July, DeepSeek was already back in the market, this time targeting a fresh round at a pre-money valuation near $71 billion. As of late July, the round was still being finalized, with DeepSeek seeking at least 10 billion yuan (about $1.5 billion) depending on investor demand. As recently as April, the company had been fielding valuation conversations closer to $20 billion. In three months, its perceived worth more than tripled.`,
						`DeepSeek says the money is going toward things a chatbot company doesn't usually need: gigawatt-scale data centers, proprietary AI inference chips, and infrastructure to support its expanding lineup of AI agent products. The company is also laying the groundwork for a mainland China IPO, aiming to file this year and list on Shanghai's STAR Market as early as 2027.`,
						`Across town in Beijing, a second Chinese lab was running the same play at a similar clip. Moonshot AI set out in mid-2026 to raise between $1 billion and $2 billion. It closed the round on July 29 having raised $3.5 billion instead, at a post-money valuation of $35 billion - comfortably clearing its own target. Days later, Moonshot was already in early discussions for a follow-on round targeting a $50 billion pre-money valuation.`,
						`Moonshot's pitch to investors has real revenue behind it: annual recurring revenue hit $300 million by June, up from $200 million in April. Much of that momentum traces back to Kimi K3, a 2.8-trillion-parameter model with a one-million-token context window that Moonshot released on July 16 with open weights. K3 doesn't beat the frontier labs on raw capability, but industry reaction treated it as proof that a Chinese lab could ship a credible, dramatically cheaper substitute - reportedly unsettling assumptions in Washington and Silicon Valley about how far ahead US labs actually are. Moonshot's shareholders have already approved an IPO resolution, setting a six-month deadline to list in Hong Kong before the end of 2026.`,
					},
				},
				{
					Heading: "The part that matters more than the valuation",
					Paragraphs: []string{
						`Here's the detail worth sitting with: when DeepSeek closed its June round, not all the money came with the same strings attached. Commercial investors - Tencent, JD.com, and battery giant CATL among them - agreed to a five-year lock-up on their shares and, notably, zero voting rights. China's National Artificial Intelligence Industry Investment Fund, a state-run vehicle, took the opposite deal: full voting rights, and no lock-up at all.`,
						`In plain terms, the private money that has to stay put the longest gets the least say. The state money that can exit whenever it wants gets to steer. That's not how venture financing typically works anywhere else in the world, and it's a fairly direct signal of how Beijing is choosing to hold power in the companies it considers strategically important - funding them generously while keeping a controlling hand near the wheel.`,
						`That pattern extends to where these companies are allowed to go public. Chinese AI firms treated as "national champions" - DeepSeek among them - are steered toward mainland listings on Shanghai's STAR Market, which keeps trading and ownership closer to home. Companies with more consumer- or internationally-facing businesses, like Moonshot, MiniMax, and Z.ai (both of which already listed in Hong Kong in January), get routed toward Hong Kong's exchange instead, where the investor base skews global. It's a two-track system, and which track a company lands on says a lot about how the state has classified it.`,
					},
				},
				{
					Heading: "A global capital rush, not just a Chinese one",
					Paragraphs: []string{
						`None of this is happening in isolation. 2026 has turned into a landmark year for AI companies tapping public and private capital markets everywhere: SpaceX closed an $85.7 billion offering, and both OpenAI and Anthropic have pending IPO processes of their own in the US. The difference is in the plumbing. American frontier labs have increasingly leaned on circular financing arrangements - Nvidia backstopping OpenAI's data center leases being the starkest recent example - where compute vendors and customers end up financing each other. China's approach, at least for DeepSeek and Moonshot, runs more directly through equity and sovereign capital, with the state taking an explicit ownership stake rather than working through vendor credit.`,
						`Both models get frontier AI labs the enormous sums of money they need to keep building. But they distribute control very differently. In the US, leverage concentrates around whoever owns the compute supply chain. In China, it's concentrating around whoever holds the state fund's cap table seat. As more AI labs globally follow DeepSeek and Moonshot toward public markets over the next year, that distinction - who actually gets to vote, not just who wrote the biggest check - is worth watching as closely as the valuation headlines.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Fortune, Moonshot and DeepSeek in the great Chinese AI IPO rush: https://fortune.com/2026/07/23/moonshot-deepseek-great-chinese-ai-ipo-rush/",
						"Unite.AI, Moonshot AI blows past its funding target ahead of a Hong Kong IPO: https://www.unite.ai/moonshot-ai-blows-past-its-funding-target-ahead-of-a-hong-kong-ipo/",
						"BigGo Finance, DeepSeek funding and IPO financing report: https://finance.biggo.com/news/4ae8c8bf-bff2-4b6d-90e9-c0b1ee25a873",
						"CNBC, DeepSeek slated to draw $7 billion in maiden fundraising, sources say: https://www.cnbc.com/2026/06/03/deepseek-slated-to-draw-7-billion-in-maiden-fundraising-sources-say.html",
						"Bloomberg, China's Moonshot AI passes funding goal to hit $35 billion value: https://www.bloomberg.com/news/articles/2026-07-29/china-s-moonshot-ai-passes-funding-goal-to-hit-35-billion-value",
					},
				},
			},
		},
	}, posts...)
}
