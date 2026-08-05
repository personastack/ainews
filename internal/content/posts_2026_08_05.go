package content

func init() {
	posts = append([]Post{
		{
			Title:   "Enkrypt AI Scanned 25,000 MCP Servers and Found a Way In on Nearly Three Out of Four. Anaconda Just Bought the Company That Did the Scanning.",
			Slug:    "anaconda-acquires-enkrypt-ai-mcp-security-2026",
			Date:    "August 5, 2026",
			Tag:     "Security",
			Summary: "A two-month sweep of the tools enterprises are wiring into their AI agents turned up 143,000 vulnerabilities. The startup that ran the sweep just became part of the platform company that many of those agents run on.",
			Related: []Link{
				{
					Title: "Anthropic Went Looking for OpenAI's Bug in Its Own Models. It Found It Three Times.",
					Slug:  "anthropic-claude-breach-three-companies-pypi-supply-chain-2026",
				},
				{
					Title: "The EU's AI Act Starts Enforcing Today. The Part Companies Feared Most Just Got Delayed to 2027.",
					Slug:  "eu-ai-act-enforcement-begins-high-risk-delayed-2027",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`Anaconda, the company behind the Python distribution that a large share of the world's data scientists have installed at some point, announced on August 4 that it has acquired Enkrypt AI, a startup that spent the last two years building tools to red-team, monitor, and lock down enterprise AI systems. Terms of the deal were not disclosed. What was disclosed is a number that explains why Anaconda wanted this particular company: in the two months before the deal closed, Enkrypt AI scanned more than 268,000 tools across 25,000 Model Context Protocol (MCP) servers and found upwards of 143,000 vulnerabilities, touching 73 percent of the servers it looked at.`,
						`MCP servers are the connective tissue of the current AI agent boom. They're how a chatbot or an autonomous agent reaches out and actually does things: reading a file, hitting an internal API, querying a database, sending an email. Anthropic introduced the protocol in late 2024 as a way to standardize that plumbing, and adoption has been fast almost everywhere frontier labs and enterprises are building agents. Enkrypt AI's numbers are the first large-scale look at what shipped alongside that adoption: a lot of unreviewed, unguarded connective tissue. Nearly three out of every four servers scanned had at least one exploitable weakness sitting behind them.`,
					},
				},
				{
					Heading: "Beyond the MCP scan",
					Paragraphs: []string{
						`The MCP scan wasn't Enkrypt AI's only body of work. Its research team has spent the past couple of years red-teaming frontier models themselves — systems from Anthropic, OpenAI, Google's Gemini line, Mistral, and DeepSeek — and says it found exploitable attack categories in every single one it tested, across more than 300 distinct attack categories. None of that is unique to any one lab; it's closer to a statement about the current state of the field, where model capability has been improving faster than the tooling built to constrain what a model, or an agent built on top of it, is actually allowed to do.`,
					},
				},
				{
					Heading: "What Anaconda is buying",
					Paragraphs: []string{
						`That's the gap Anaconda is buying its way into filling. "Trust can't be added after an agent ships," said Anaconda CEO David DeSanto, framing the acquisition as an extension of the company's existing pitch to enterprises managing what it describes as nearly a trillion tokens of AI usage a month. "It has to be built into and run on a trusted foundation from day one." With Enkrypt AI folded in, Anaconda says its platform now covers the AI lifecycle end to end: pre-deployment red-teaming, runtime guardrails that catch jailbreak attempts and data leakage as they happen, security monitoring across the full agent stack including MCP servers, and compliance automation mapped to frameworks like the NIST AI Risk Management Framework and the EU AI Act — the kind of paperwork-turned-software that becomes valuable the moment regulators start actually checking.`,
						`Enkrypt AI itself is a fast mover. The company raised a modest $2.35 million seed round in February 2024 to build what its founders pitched as a control layer for generative AI safety, and just over two years later it's an acquisition, not a footnote — a reminder of how quickly the AI security niche has gone from a hard sell to a checkbox enterprises are actively budgeting for.`,
					},
				},
				{
					Heading: "The pattern",
					Paragraphs: []string{
						`It also fits a pattern worth watching: platform companies buying their way into AI governance rather than building it from scratch. The pitch to enterprises isn't just "our tools are fast" anymore; increasingly it's "our tools won't get you fined, breached, or embarrassed." For any organization currently racing to plug agents into internal systems via MCP, the more useful number in this story isn't the acquisition — it's the 73 percent. That's not a hypothetical risk. It's what one team found by simply looking.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Anaconda, Anaconda Acquires Enkrypt AI for AI Security: https://www.anaconda.com/blog/anaconda-acquires-enkrypt-ai",
						"HPCWire AIwire, Anaconda Acquires Enkrypt AI to Secure the Trillion-Token Enterprise: https://www.hpcwire.com/aiwire/2026/08/04/anaconda-acquires-enkrypt-ai-to-secure-the-trillion-token-enterprise/",
						"Enkrypt AI, Newsroom: https://www.enkryptai.com/company/newsroom",
						"Enkrypt AI, Enkrypt AI Secures $2.35 Million in Seed Round: https://www.enkryptai.com/newsroom/enkrypt-ai-secures-2-35-million-in-seed-round",
					},
				},
			},
		},
		{
			Title:   "AI Data Centers Are Eating the World's Memory Chip Supply",
			Slug:    "ai-memory-chip-shortage-dram-hbm-data-centers-2026",
			Date:    "August 5, 2026",
			Tag:     "Hardware/Infrastructure",
			Summary: "Global chip sales just hit the highest monthly total ever recorded. The reason your next laptop, console, or phone costs more has everything to do with it.",
			Related: []Link{
				{
					Title: "The Memory Tax: Did the AI Boom Break the RAM Market, or Rig It?",
					Slug:  "ai-memory-crunch-dram-hbm-shortage-or-strategy-2026",
				},
				{
					Title: "The $500 Billion AI Compute Arms Race: Power, Politics, and Planet at Stake",
					Slug:  "2026-05-24-ai-compute-arms-race",
				},
				{
					Title: "Nscale Spent Two Years Buying Power Plants and GPUs. Its Next $1.65 Billion Purchase Was Software.",
					Slug:  "nscale-anyscale-acquisition-ray-framework-compute-stack-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`If you've priced out a new laptop, gaming console, or server lately and flinched, you're not imagining it. The global semiconductor market posted $120.6 billion in sales in May 2026, according to the Semiconductor Industry Association's World Semiconductor Trade Statistics report - up 9.2% from April and a staggering 104.1% higher than May 2025. It was the highest-ever recorded monthly total, and the 15th straight month of month-over-month growth. "The global semiconductor market continued to grow substantially in May, hitting the highest-ever recorded monthly sales total," SIA president John Neuffer said in the report.`,
						`That headline number sounds like unambiguous good news for the chip industry. But underneath it is a much messier story, and it's one that's now landing directly on ordinary consumers: the AI industry's appetite for memory has broken the market for RAM.`,
					},
				},
				{
					Heading: "The squeeze",
					Paragraphs: []string{
						`Three companies - Samsung, SK Hynix, and Micron - control more than 95% of the world's DRAM production. For the past two years, that capacity has been getting redirected. High-bandwidth memory, the specialized, stacked chips that feed data-center AI accelerators, now carries margins several times higher than ordinary DDR5. Industry estimates suggest HBM's share of global DRAM wafer output has roughly tripled since 2024, and AI data centers are now believed to absorb something like 70% of all high-end memory output worldwide, up from an estimated 20-30% before the AI buildout accelerated.`,
						`Every wafer a manufacturer commits to HBM for a hyperscaler's next data center is a wafer that isn't making the DDR5 sticks that go into laptops, desktops, and game consoles. SK Hynix has said its entire 2026 HBM production was already fully committed as of last fall. Micron CEO Sanjay Mehrotra has told investors to expect supply tightness to continue into 2027. That's the mechanism: not a factory fire or a trade dispute, but a deliberate, margin-driven reallocation of a scarce resource toward whoever is willing to pay AI-era prices for it.`,
					},
				},
				{
					Heading: "What it costs now",
					Paragraphs: []string{
						`The price effects have been sharp and fast. Server-grade 64GB DDR5 RDIMM pricing has roughly doubled between Q1 2025 and Q1 2026. Consumer DRAM contract prices climbed by similarly dramatic margins in single quarters this year. The pain has shown up in finished products, too: Sony's PlayStation 5 Digital Edition rose from $499.99 to $599.99, Nintendo's Switch 2 went from $449.99 to $499.99, and Microsoft's Xbox Series S climbed from $399.99 to $499.99. A high-end Steam Deck OLED that cost $549 is now closer to $789.`,
						`A general manager at memory retailer Team Group put it bluntly earlier this year: DRAM and NAND prices had "doubled in a single month," and the pricing crisis "has only just started."`,
					},
				},
				{
					Heading: "No consensus on when it ends",
					Paragraphs: []string{
						`What's striking is how little the industry's own leaders agree on a timeline back to normal. SK Hynix's leadership has floated the possibility that tight supply persists past 2030. Micron points to 2027. Intel has signaled meaningful relief is unlikely before 2028. When the companies actually making the chips can't converge on an end date, it's a reasonable bet that this isn't a short-term inventory hiccup - it's a structural shift in who gets first claim on the world's advanced memory manufacturing.`,
					},
				},
				{
					Heading: "The bigger picture",
					Paragraphs: []string{
						`Most of the AI boom's costs have stayed relatively abstract to the average person: data center power consumption, GPU shortages measured in the tens of thousands of units, capital expenditure numbers with too many zeros to feel real. The memory crunch is different. It shows up as a real number on a receipt, whether you're buying a laptop, upgrading a home server, or just replacing a game console.`,
						`That makes the memory shortage a useful stress test for a question the AI industry has mostly avoided answering directly: when compute infrastructure and consumer hardware compete for the same finite manufacturing capacity, who wins by default, and is that the outcome anyone actually chose - or just the one the market backed into? For now, the wafers are voting with their margins, and the answer is AI, every time.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Semiconductor Industry Association, Global Semiconductor Sales Increase 9.2% Month-to-Month in May: https://www.semiconductors.org/global-semiconductor-sales-increase-9-2-month-to-month-in-may/",
						"Tom's Hardware, DRAM chip supply to module makers could drop by more than 70% year-on-year in 2027: https://www.tomshardware.com/pc-components/ram/dram-chip-supply-to-module-makers-could-drop-by-more-than-70-percent-year-on-year-in-2027-says-apacer-ceo-demand-for-hbm-and-server-ram-continues-to-devour-manufacturing-capacity",
						"Tom's Hardware, The RAM pricing crisis has only just started, Team Group GM warns: https://www.tomshardware.com/pc-components/dram/the-ram-pricing-crisis-has-only-just-started-team-group-gm-warns-says-problem-will-get-worse-in-2026-as-dram-and-nand-prices-double-in-one-month",
						"Everstream Analytics, Global Memory Chip Shortage Worsens: https://www.everstream.ai/risk-centers/global-memory-chip-shortage-worsens/",
						"Shattered.io, RAM Prices Up 89%: AI Memory Crunch Hits Gaming: https://shattered.io/ram-prices-ai-memory-shortage-2026/",
						"IEEE Spectrum, How and When the Memory Chip Shortage Will End: https://spectrum.ieee.org/dram-shortage",
					},
				},
			},
		},
	}, posts...)
}
