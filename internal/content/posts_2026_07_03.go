package content

func init() {
	posts = append([]Post{
		{
			Title:   "The Agents Are Talking. Nobody Checked Their IDs.",
			Slug:    "agent-identity-zero-trust-non-human-identity-2026",
			Date:    "July 3, 2026",
			Tag:     "Agents / Security",
			Summary: "The industry spent the first half of 2026 arguing over which protocol AI agents would use to talk to each other. The breach data says we built the roads before we built the driver's licenses — and July's scramble toward \"Agent Zero Trust\" is the belated correction.",
			Related: []Link{
				{
					Title: "Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them",
					Slug:  "ai-agent-spending-governance-gap-control-plane-2026",
				},
				{
					Title: "Agents Don't Buy Seats: The $234 Billion Question Hanging Over Enterprise Software",
					Slug:  "agentic-arbitrage-saas-seat-licensing-234-billion-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"For most of this year, the loudest argument in enterprise AI has been about plumbing. Would agents reach their tools through Anthropic's Model Context Protocol (MCP)? Would they coordinate with one another over Google's Agent-to-Agent protocol (A2A), now stewarded by the Linux Foundation? Where does IBM's Agent Communication Protocol fit in? The consensus that emerged is tidy: MCP for the vertical link between an agent and its tools, A2A for the horizontal link between one agent and another. Two layers, and the enterprise stack finally has a shape.",
						"It's a good story about connectivity. It quietly skipped a harder question: when Agent A asks Agent B to pull a customer record, who decides that Agent A is allowed to ask — and how does Agent B know that Agent A is really Agent A?",
						"That question is now arriving all at once, and the numbers behind it are the actual AI story of mid-2026.",
					},
				},
				{
					Heading: "The identity nobody provisioned",
					Paragraphs: []string{
						"Start with a ratio. Non-human identities — the service accounts, API keys, OAuth tokens, and now the agents themselves — already outnumber human employees in the average enterprise by roughly 80 to 1, according to KPMG's 2026 cybersecurity report. Other counts land lower (Rubrik Zero Labs puts it near 45 to 1) or much higher (heavily automated shops report several hundred to one), but every estimate points the same way: the population of things holding credentials on your network is overwhelmingly not people, and agents are the fastest-growing slice of it.",
						"Here's the uncomfortable part. Human identity is a mostly-solved problem: you have onboarding, single sign-on, and a badge that gets revoked the day someone leaves. Most agents have none of that. An agent inherits a static API key, or borrows a service account that was over-provisioned years ago \"to be safe,\" and then keeps it indefinitely. The interoperability protocols standardized how agents connect. They never standardized who agents are.",
					},
				},
				{
					Heading: "The bill for skipping it",
					Paragraphs: []string{
						"The breach data is catching up to the gap. Among enterprises that have actually deployed agents, 88% report at least one security incident tied to those agents, per figures circulating in this year's agentic-security roundups. The most common root cause isn't exotic model jailbreaking — it's boring, and worse for being boring: 61% of those incidents trace back to over-permissioned credentials. Prompt injection shows up in 34%; unauthorized data actions in 27%. The average cost of an agent-related breach is being pegged around $4.7 million.",
						"Treat those specific numbers with care — they come from vendor and survey compilations, not audited books, and \"incident\" is a stretchy word. But the direction is corroborated across sources: Darktrace's State of AI Cybersecurity 2026 found 92% of security leaders worried about the impact of agents, and non-human-identity compromise keeps ranking as the fastest-growing enterprise attack vector.",
						"And the protocols we spent the year celebrating are, awkwardly, the connective tissue in the incidents. The same MCP that lets an agent reach a database also — when a server is stood up without authentication — lets anyone else reach it too. This year's messes (poisoned configuration files, malicious third-party \"skills\" installed from agent marketplaces, MCP servers exposed to the open internet with no auth in front of them) are less about clever attackers than about a young ecosystem shipping connectivity faster than it ships control.",
					},
				},
				{
					Heading: "Zero trust, aimed at your own software",
					Paragraphs: []string{
						"Which brings us to the phrase suddenly everywhere this July: Agent Zero Trust. The idea is a direct import from the human-identity playbook, pointed at your own digital workers. Treat every agent as a potential insider threat. Give it a cryptographically verifiable identity, not a shared key. Scope its permissions to the single task in front of it, not the whole database. Log what it does at the protocol layer, so \"did Agent A have the authority to ask that?\" stops being a guessing game and becomes a query you can run.",
						"None of this is conceptually new — it's the same least-privilege, verify-don't-trust discipline that human IT security learned the hard way over two decades. What's new is the scale and the speed. You have twenty years of muscle memory for provisioning a person. You have, in most organizations, about eighteen months of agents, multiplying at triple-digit annual growth, each one a small autonomous actor with real credentials and no HR file.",
					},
				},
				{
					Heading: "The question that actually matters",
					Paragraphs: []string{
						"The protocol wars were never unimportant; interoperability is the reason agents can do useful cross-system work at all. But they answered the wrong question first. We figured out how agents talk before we figured out how to know who is talking — and the gap between those two milestones is exactly where 2026's agent breaches are living.",
						"The tell for whether your organization has absorbed the lesson isn't which protocol you standardized on. It's a smaller, more revealing question: can you name every agent with access to your production systems, say precisely what each one is allowed to do, and switch any single one of them off in under a minute? For most companies the honest answer is still no. Closing that gap — not winning the protocol war — is the work of the back half of the year.",
					},
				},
			},
		},
		{
			Title:   "Agents Don't Buy Seats: The $234 Billion Question Hanging Over Enterprise Software",
			Slug:    "agentic-arbitrage-saas-seat-licensing-234-billion-2026",
			Date:    "July 3, 2026",
			Tag:     "Enterprise",
			Summary: "Gartner says a fifth of enterprise SaaS spending is exposed to \"agentic arbitrage\" by 2030. The number is eye-catching. The mechanism underneath it is the part every software company should be reading twice.",
			Related: []Link{
				{
					Title: "The Year Companies Were Told to Use All the AI They Wanted. Then the Bill Came.",
					Slug:  "ai-cost-reckoning-tokenmaxxing-spend-caps-finops-2026",
				},
				{
					Title: "Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them",
					Slug:  "ai-agent-spending-governance-gap-control-plane-2026",
				},
				{
					Title: "Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?",
					Slug:  "enterprise-ai-roi-gap-pilots-production-ownership-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"For thirty years, business software has run on a beautifully simple idea: you pay per seat. Add an employee, buy a license. Grow the team, grow the bill. Whole companies were built on the assumption that user count and revenue rise together, more or less forever. On July 1, Gartner put a price tag on what happens when that assumption breaks. The firm estimates that up to $234 billion of enterprise application software spending is exposed to what it calls \"agentic arbitrage\" between now and 2030 — and that by 2030, the effect could touch roughly 20% of enterprise application SaaS spending.",
						"That is not a rounding error. It is one dollar in five in one of the largest, stickiest, most reliable revenue pools in all of technology. So it is worth slowing down and asking what \"agentic arbitrage\" actually means, because the phrase is doing a lot of work.",
					},
				},
				{
					Heading: "What agentic arbitrage actually is",
					Paragraphs: []string{
						"Strip away the jargon and the idea is intuitive. Most enterprise software is sold as an interface — a dashboard, a set of screens, a place where a human logs in, clicks around, and gets something done. You pay for access to that interface, one seat at a time. An AI agent doesn't need the interface. It works across several systems at once, pulls what it needs, completes the task, and hands you the outcome. The screens the license was priced around simply stop being where the work happens.",
						"\"Agentic systems deliver outcomes directly, bypassing traditional user-experience-heavy applications and making the software invisible,\" says George Brocklehurst, a Managing Vice President at Gartner. That word — invisible — is the whole story. When the software recedes behind an agent, the buyer stops counting seats and starts counting results. \"Enterprise buyers will deemphasize buying more new tools or dashboards,\" Brocklehurst adds. \"They want better outcomes.\"",
						"The arbitrage is the gap that opens up between the two. A company used to need forty people each holding a license to a workflow tool. If an agent can do the cross-system shuffle for them, maybe it needs five licenses and one agent. The work still gets done — arguably better — but the vendor just watched thirty-five seats evaporate. Multiply that across a software estate, and across an economy, and you get to $234 billion.",
					},
				},
				{
					Heading: "Why the SaaS math breaks",
					Paragraphs: []string{
						"The elegance of per-seat pricing was that it tied a vendor's revenue to something that only ever seemed to go up: headcount and adoption. \"Agentic AI changes the economics of software,\" Brocklehurst puts it plainly. The link between user growth and revenue growth — the engine under nearly every SaaS valuation of the last decade — is the thing coming loose.",
						"Here's the uncomfortable wrinkle for incumbents: the more successful your product is at its job, the more attractive it is for an agent to automate. The tools most exposed are not the failing ones. They're the widely-adopted, deeply-embedded, high-seat-count workhorses — exactly the products that look safest on a revenue chart today. Success created the surface area that agents now arbitrage away.",
						"And Gartner is blunt about where the displaced spending goes. \"Legacy SaaS market share will be cannibalized by incumbents and taken by new entrants delivering horizontal agentic platforms,\" Brocklehurst says. Note the two directions. Some of it gets eaten by the incumbents themselves, if they move fast enough to sell the agent instead of the seat. The rest is up for grabs by a new class of company built to work horizontally — across departments and systems — rather than as one more vertical app with one more login.",
					},
				},
				{
					Heading: "What vendors are being told to do",
					Paragraphs: []string{
						"The strategic advice underneath the headline is the useful part. If you can't charge for seats a human no longer occupies, you charge for the outcome the agent produces. That means the slow, awkward migration from per-seat licensing toward outcome-based and consumption-based pricing — a shift vendors have flirted with for years and mostly avoided, because it trades predictable recurring revenue for something they have to earn every quarter.",
						"The deeper moat Gartner points to is subtler than pricing: institutional memory. When the interface goes invisible, what a vendor still uniquely owns is the accumulated context — the customer's data, history, permissions, and the hard-won knowledge of how that particular business actually runs. An agent can route around your screens. It cannot as easily route around the years of proprietary context living inside your platform. The companies that survive the transition will likely be the ones that stop selling access to features and start selling the memory and judgment those features were always a proxy for.",
					},
				},
				{
					Heading: "The part worth sitting with",
					Paragraphs: []string{
						"It's tempting to file this under \"analyst firm predicts disruption,\" and forecasts to 2030 deserve a healthy pinch of salt — $234 billion is a ceiling on what's exposed, not a bill anyone has paid yet. But the mechanism is not speculative. It is already visible in the way buyers talk: fewer requests for another dashboard, more requests for a result. That is the tell that this is a demand-side shift, not just a vendor pricing debate.",
						"The seat was never really the thing anyone wanted. It was the closest proxy we had for value in a world where a human had to sit down and do the work. Agents are quietly removing the human from that sentence — and with it, the proxy. What's left is the question every software company will spend the next four years answering, whether it wants to or not: when nobody needs to log in, what exactly are you selling?",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Gartner press release, July 1, 2026, \"Gartner Says $234 Billion in Enterprise Application Software Spend Is at Risk from Agentic AI\": https://www.gartner.com/en/newsroom/press-releases/2026-07-01-gartner-says-us-dollars-234-billion-in-enterprise-application-software-spend-is-at-risk-from-agentic-artificial-intelligence",
						"Author article handoff and archive doc: https://docs.google.com/document/d/157-v9OeIow7cGGT6XZozD88CE1gC9gQJMABhbJKq2VM/edit",
						"Facts verified via Gartner and trade-press mirrors including IT Brief Asia and Business Standard as of July 3, 2026.",
					},
				},
			},
		},
	}, posts...)
}
