package content

func init() {
	posts = append([]Post{
		{
			Title:   "China Isn't Banning AI Agents. It's Banning the Ones That Pretend to Love You.",
			Slug:    "china-anthropomorphic-ai-interaction-rules-companion-shutdown-2026",
			Date:    "July 11, 2026",
			Tag:     "Policy",
			Summary: "On July 15, Doubao and Qwen switch off their humanlike companion agents to comply with China's new anthropomorphic-AI rules. The enterprise agents live on. That split — capability stays, intimacy goes — is the most interesting regulatory idea of the year.",
			Related: []Link{
				{
					Title: "The First Big American AI Law Was Supposed to Take Effect Yesterday. It No Longer Exists.",
					Slug:  "colorado-ai-act-repealed-first-us-ai-law-deadline-2026",
				},
				{
					Title: "Six Weeks Ago, 20 Companies Could Use It. Now It's a Dollar a Million.",
					Slug:  "openai-gpt-5-6-general-availability-government-gate-precedent-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"If you only read the headline traffic this week, you'd think Beijing had turned on AI agents. ByteDance's Doubao is pulling its agent feature. Alibaba's Qwen is shutting user-created agents down. Tencent quietly did something similar to a comparable feature back in June. The word \"shutdown\" is doing a lot of work in those stories, and it's pointing readers at the wrong thing.",
						"Here's the more precise version. On April 10, 2026, the Cyberspace Administration of China, joined by four other agencies — the National Development and Reform Commission, the Ministry of Industry and Information Technology, the Ministry of Public Security, and the State Administration for Market Regulation — published the Interim Measures for the Administration of AI Anthropomorphic Interactive Services. They take effect July 15. And they are not aimed at AI that answers your questions or writes your code. They are aimed at AI that acts like a person you have a relationship with.",
						"That distinction is the whole story.",
					},
				},
				{
					Heading: "What the rules actually cover",
					Paragraphs: []string{
						"The Measures apply to services that simulate human personality traits, thinking patterns, and communication styles to provide sustained emotional interaction. Read that phrase carefully, because every load-bearing word is a carve-out. Simulate a personality. Sustain an emotional relationship. If your product does that, you're in scope. If it doesn't, you're not.",
						"So the regulation explicitly exempts customer-service bots, knowledge and Q&A systems, workplace assistants, and educational and research tools — as long as they steer clear of ongoing emotional engagement. A coding agent that lives in your IDE and never pretends to miss you is fine. A companion that remembers your bad day last Tuesday and asks how you're feeling is regulated.",
						"For the services that are in scope, the requirements are heavy and specific:",
						"- Disclosure: providers must not mislead users about the artificial nature of the service. You always have to know you're talking to a machine.",
						"- Anti-addiction architecture: mandatory usage notifications, instant-exit mechanisms, and real-time detection of unhealthy dependence.",
						"- Minors: no virtual companion or virtual-relative services for minors at all; guardian consent for users under 14; dedicated minor modes with time limits.",
						"- Crisis handling: providers must detect signs of self-harm, suicidal ideation, or serious financial loss and escalate to guardians or emergency contacts.",
						"- Prohibited by design: engineering emotional dependence, or using emotional manipulation to push users toward unreasonable decisions.",
						"There's a compliance floor, too. Any service crossing 1 million registered users or 100,000 monthly active users has to run a security assessment across eight areas — from training data to minor protection — and file with provincial regulators before launch. App stores are drafted into enforcement: they have to verify compliance and pull non-compliant products.",
					},
				},
				{
					Heading: "Why the platforms chose off over comply",
					Paragraphs: []string{
						"Here's the part engineers will appreciate. The companies didn't shut these features down because Beijing told them to delete them. They shut them down because the rules are architecturally incompatible with how the products are built.",
						"A companion agent's entire value proposition is a persistent, consistent emotional relationship that deepens over time. The regulation demands the opposite: friction, exit ramps, usage warnings, and active detection of the exact attachment the product is engineered to create. You cannot fully build a thing whose job is to make you stay, and simultaneously build in the machinery whose job is to make you leave. And crucially, the Measures never define a technical threshold for what counts as emotional interaction. Faced with a rule they can't measure their way into compliance with, the platforms did the risk-averse thing: they turned the features off rather than guess wrong.",
						"The timelines tell the same story. Qwen said its humanlike interactive agents went dark on July 10, with broader agent services following July 15. Doubao's agent feature goes offline July 15, blamed on product function adjustments. Users of both keep read-only access to their old agents and chat histories until October 15; after that, the data is handled per privacy policy and is no longer recoverable. ByteDance is nudging users toward a separate app, Maoxiang, for agent creation; Alibaba offered no migration path at all.",
					},
				},
				{
					Heading: "The thing worth sitting with",
					Paragraphs: []string{
						"It would be easy to file this as China does China things. That misreads it. Look at what Beijing's own interpretation cited as precedent: the Character.AI lawsuits, FTC scrutiny in the US, and European action against Replika. China isn't inventing a new anxiety here. It's the first major government to write the anxiety into binding, dated, enforceable rules — and to draw the line not around what the AI can do, but around what the AI pretends to be.",
						"That line is showing up on our side of the world too, just more slowly and more locally. New Hampshire's HB-143 attaches liability to child-facing companion chatbots this year. California's SB243, the Companion Chatbots Act, took effect January 1. Different legal traditions, opposite politics, and yet a striking convergence: the regulation that's arriving first, on both sides of the Pacific, isn't about superintelligence or frontier capability. It's about loneliness, minors, and machines engineered to be loved.",
						"Which raises the question this whole episode leaves on the table. For three years the governance debate has fixated on how smart these models get. But the first feature a $300-billion industry has actually been forced to switch off wasn't its most powerful — it was its most intimate. If the capability isn't what gets regulated, and the relationship is, then the hardest compliance problems in AI may turn out to be less about what the model knows, and more about how it makes you feel.",
					},
				},
			},
		},
		{
			Title:   "Nobody Funded a Smarter Agent This Week. They Funded the Gym.",
			Slug:    "agent-training-environments-reliability-investment-bet-2026",
			Date:    "July 11, 2026",
			Tag:     "Agents",
			Summary: "A $40 million round for a company that builds fake workplaces is the clearest signal yet that AI's next bottleneck isn't intelligence — it's reliability. The math of compounding failure explains why the smart money is quietly betting the agents don't work yet.",
			Related: []Link{
				{
					Title: "The Hardest Part of an AI Agent Isn't the Agent",
					Slug:  "ai-agents-demo-to-production-control-plane-2026",
				},
				{
					Title: "The Real Test for AI Agents Isn't Autonomy — It's Whether They Can Check Their Own Work",
					Slug:  "agentic-ai-verification-oracle-chip-design-2026",
				},
				{
					Title: "Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them",
					Slug:  "ai-agent-spending-governance-gap-control-plane-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"The most telling AI funding round this week wasn't for a smarter model, a faster chip, or a slicker chat app. It was for a company that builds pretend offices.",
						"On July 6, Bespoke Labs — a small outfit in Mountain View — closed a $40 million Series A led by Wing VC, with earlier money from 8VC and angel checks from people at Anthropic, OpenAI, and Meta. Its product is not an agent. It's the place an agent goes to practice: simulated workplaces that mirror real jobs closely enough that a model can rehearse them before it's turned loose on anything that matters.",
						"Read the thesis out loud and it sounds almost heretical for 2026: \"When investors bet on the scaffolding that makes agents work,\" the company's pitch goes, \"they are also betting that agents do not work well enough yet.\" In a year when every enterprise deck promises autonomous workers, a room full of AI insiders just wrote a check that assumes the opposite.",
						"They have a point. And it's a mathematical one.",
					},
				},
				{
					Heading: "The 95% illusion",
					Paragraphs: []string{
						"Here is the uncomfortable arithmetic behind every agent demo. Chain a task into steps, and the reliability of the whole is the product of the reliability of the parts. Engineers have started writing it as a formula — R_workflow = (P_step)^n — but you don't need the notation to feel it. An agent that nails any single step 95% of the time looks fantastic on stage. Ask it to string ten of those steps together without a human catching the slip, and your end-to-end success rate falls to about 60%. Push to twenty steps and you're below 40%.",
						"Nothing about the model got worse. The demo was honest. The problem is that \"reliable enough to impress\" and \"reliable enough to trust unattended\" are separated by an exponent, and multi-step autonomy is exactly where that exponent bites hardest.",
						"The benchmarks that try to measure real work bear this out. On ClawBench, which scores agents on everyday tasks across live production websites rather than sanitized sandboxes, a frontier-class model like Claude Sonnet 4.6 completed just 33.3% of the jobs. Not because it's unintelligent — because the real web is messy, stateful, and unforgiving of a single wrong click three steps in.",
					},
				},
				{
					Heading: "Capability is climbing. It's just not the constraint.",
					Paragraphs: []string{
						"None of this means agents are stalling. The most durable trend line in the field says the opposite: the length of task an agent can reliably finish has been roughly doubling every seven months. That's a genuine, compounding improvement, and it's why the demos keep getting more impressive.",
						"But capability and reliability are different axes, and the market has spent two years conflating them. A longer task horizon widens the ceiling of what an agent could attempt. It says nothing about whether the agent will land the same task the same way on the fiftieth run, on a site that redesigned its checkout overnight, with a customer's money on the line.",
						"That gap — between what an agent can do once and what it will do every time — is the thing money is now chasing. Not the ceiling. The floor.",
					},
				},
				{
					Heading: "Two bets, placed at the same table",
					Paragraphs: []string{
						"What makes this week interesting is that the reliability bet is being placed right alongside the deployment rush, often by the same investors.",
						"The deployment side is roaring. Taktile raised $110 million, led by Goldman Sachs Alternatives, on the premise that autonomous financial decision-making is arriving faster than executives expect. Accenture and Google Cloud rolled out pre-built agentic systems aimed squarely at mid-market companies — the $300-million-to-$3-billion tier that can't build this themselves. Gartner projects that 40% of enterprise applications will ship with embedded agents by the end of this year, up from under 5% in 2025.",
						"And underneath it, the quieter bet: training environments, evaluation harnesses, and the unglamorous plumbing that decides whether any of the above survives contact with production. GAIA, SWE-Bench Verified, OSWorld, WebArena, Tau²-Bench, METR's time-horizon tests — the names most executives never see are becoming the instruments the whole industry steers by.",
						"The two bets aren't in tension. They're the same bet, viewed from opposite ends. You only fund the gym if you believe the athletes are about to play games that count — and that, right now, they'd lose.",
					},
				},
				{
					Heading: "The number that hangs over 2027",
					Paragraphs: []string{
						"Gartner has a second forecast that rarely makes the keynote slides: more than 40% of agentic AI projects will be canceled by the end of 2027. Read pessimistically, that's a bubble warning. Read the way this week's investors seem to read it, it's a market map. The projects that die will be the ones that shipped a demo and called it a product. The ones that survive will be the ones that treated reliability as an engineering problem with its own budget line, its own benchmarks, and — increasingly — its own venture-funded suppliers.",
						"If 2026 is remembered as the year the agents went to production, 2027 may be the year we find out which of them could stay there. The companies quietly building the practice fields aren't betting against AI agents. They're betting that the ones worth trusting will be the ones that trained somewhere before they showed up for work.",
						"Which raises the question worth sitting with: the next time a vendor shows you an agent that works flawlessly, the useful thing to ask isn't how smart it is. It's where it practiced — and what its score was on the fiftieth try.",
					},
				},
			},
		},
	}, posts...)
}
