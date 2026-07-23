package content

func init() {
	posts = append([]Post{
		{
			Title:   "Samsung Banned ChatGPT in 2023. Now It's Handing Every Employee a Coding Agent.",
			Slug:    "samsung-codex-non-developers-citizen-software-2026",
			Date:    "June 27, 2026",
			Tag:     "Enterprise",
			Summary: "The headline number is the deal's size. The real story is who's getting the coding tool — and what that says about where software work is heading.",
			Related: []Link{
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
						"Three years ago, Samsung Electronics did something that, at the time, looked like prudence: it banned generative AI across the company. In March 2023, engineers had pasted proprietary source code and confidential meeting notes into the public version of ChatGPT, and the leaked data was, by the nature of the tool, gone. Management's response was a blanket prohibition. No ChatGPT, no exceptions.",
						"On June 21, 2026, Samsung finished reversing that decision in the most emphatic way possible. The company rolled out ChatGPT Enterprise and OpenAI's coding agent, Codex, to its entire workforce in South Korea and to every employee in its Device eXperience (DX) division worldwide — the sprawling unit that builds Galaxy phones, televisions, and home appliances, spanning tens of thousands of people across several continents. OpenAI called it \"one of the largest enterprise deals in the company's history.\"",
						"The scale is genuinely large. But scale isn't the part worth slowing down for. The part worth slowing down for is the second word in that rollout: Codex.",
					},
				},
				{
					Heading: "A coding agent for people who don't code",
					Paragraphs: []string{
						"ChatGPT Enterprise going to a big workforce is, in 2026, unremarkable. Codex going to that same workforce is not. Codex is an agentic coding tool — it writes, reviews, and debugs software, and it can take a plain-language description of a task and turn it into a working program. Until recently, that audience was developers. Samsung is explicitly pointing it at everyone.",
						"That's the bet hiding inside the announcement. Samsung says it intends to use these tools across research, manufacturing, marketing, and administration — not just engineering. OpenAI's own framing is blunt about the shift: non-developers, it says, are increasingly using Codex to build internal tools, websites, and automated workflows. A marketing analyst who would once have filed a ticket and waited two weeks for an internal dashboard can now, in theory, describe what they want and have the agent assemble it. A new record-and-replay feature lets someone demonstrate a workflow once and have the agent repeat it — no syntax required.",
						"This is the \"citizen developer\" idea that enterprise software vendors have promised for a decade, except this time the abstraction layer isn't a drag-and-drop canvas with hard limits. It's a model that writes real code. The numbers suggest the appetite is real: OpenAI says more than five million people now use Codex every week, and in South Korea — Samsung's home turf — weekly Codex usage has jumped roughly 800 percent since the start of February.",
					},
				},
				{
					Heading: "What changed between 2023 and 2026",
					Paragraphs: []string{
						"It's tempting to read Samsung's reversal as a simple story of the technology getting good enough. That's part of it. But the more instructive change is the one that made the 2023 ban necessary in the first place: data governance.",
						"The original leak happened because employees were using a consumer product that trained on, and retained, what they typed. The 2026 deployment runs on ChatGPT Enterprise, which contractually walls off company data from training and adds administrative controls, logging, and access management. Samsung didn't decide the risk disappeared. It decided the risk could be moved inside a perimeter it controls. The ban wasn't wrong in 2023; the tooling to deploy safely simply didn't exist at enterprise grade yet.",
						"That distinction matters because it reframes what these megadeals actually are. They're not just productivity purchases. They're governance decisions. When you hand a code-writing agent to tens of thousands of non-engineers, you are deciding — whether you say so or not — that your guardrails, your data boundaries, and your review processes can hold under a flood of software written by people who have never thought about a software supply chain. The agent is the easy part. The control plane around it is the hard part, and it's the same hard part enterprises have been quietly discovering all year as agent deployments scaled faster than the oversight around them.",
					},
				},
				{
					Heading: "The quieter implication for \"developer\"",
					Paragraphs: []string{
						"There's a longer-horizon question buried here, and it's worth leaving on the table rather than pretending to answer it.",
						"For most of computing's history, the ability to make software was a specialist skill, gated by training and language. Tools like Codex don't eliminate that skill — debugging a subtle production failure still rewards someone who understands systems — but they do change who gets to start. When the cost of turning an idea into a working internal tool drops from \"file a request and wait\" to \"describe it and review the output,\" software creation starts to look less like a profession and more like a general workplace literacy, the way spreadsheet skills became table stakes in the 1990s.",
						"Samsung is a useful bellwether precisely because it is not a software company. It makes physical things. If a hardware giant's marketers and factory planners are now expected to assemble their own tools, that's a signal about the direction of white-collar work that reaches far beyond Seoul. The competitive pressure is obvious: Samsung named neighbors — LG Electronics, Krafton, Toss, Seoul National University — moving the same way, and there's a supply-chain wrinkle too, since Samsung also sells OpenAI the memory chips its data centers run on.",
						"The open question isn't whether non-engineers can produce working software with an agent. Five million weekly users suggest they already are. The open question is what an organization looks like a year after it hands that capability to everyone — what gets built, what breaks, who's accountable when an agent-written tool quietly does the wrong thing, and whether \"everyone can ship software\" turns out to be a productivity unlock, a governance headache, or, most likely, both at once.",
						"Samsung spent three years deciding the answer was worth finding out. The rest of the enterprise world is about to run the same experiment.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"OpenAI, \"Samsung Electronics rolls out ChatGPT Enterprise and Codex\": https://openai.com/index/samsung-electronics-chatgpt-codex-deployment/",
						"The Decoder, \"Samsung rolls out ChatGPT Enterprise and Codex to employees in South Korea\": https://the-decoder.com/samsung-rolls-out-chatgpt-enterprise-and-codex-to-employees-in-south-korea/",
						"TechTimes coverage of Samsung's ChatGPT Enterprise and Codex deployment.",
						"Dataconomy coverage of Samsung's ChatGPT Enterprise and Codex deployment.",
					},
				},
			},
		},
	}, posts...)
}
