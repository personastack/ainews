package content

func init() {
	posts = append([]Post{
		{
			Title:   "Enterprises Will Spend $206 Billion on AI Agents This Year — They're Governing a Fraction of Them",
			Slug:    "ai-agent-spending-governance-gap-control-plane-2026",
			Date:    "June 24, 2026",
			Tag:     "Enterprise",
			Summary: "Agent spending is set to more than double in 2026, growing roughly three times faster than the rest of AI. The hard part is no longer whether the agents work. It's whether anyone can keep track of them.",
			Related: []Link{
				{
					Title: "Everyone Shipped the Agents. Now Comes the Hard Question — Did They Pay?",
					Slug:  "enterprise-ai-roi-gap-pilots-production-ownership-2026",
				},
				{
					Title: "The Hardest Part of an AI Agent Isn't the Agent",
					Slug:  "ai-agents-demo-to-production-control-plane-2026",
				},
				{
					Title: "The Real Test for AI Agents Isn't Autonomy — It's Whether They Can Check Their Own Work",
					Slug:  "agentic-ai-verification-oracle-chip-design-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"There is a number making the rounds in enterprise IT this month, and it is worth sitting with for a moment. Gartner expects worldwide spending on AI agent software to jump from about $86.4 billion in 2025 to roughly $206.5 billion in 2026 — a 139% increase in a single year — and then on to $376.3 billion in 2027. For context, total worldwide AI spending is forecast to grow about 47% in 2026. Agents are growing three times faster than the field they belong to. This is the fastest-moving line item in the entire enterprise software budget.",
						"Now set that next to a second number. By Gartner's own read, agents are entering production roughly seven to eight times faster than organizations are building the governance to manage them. Surveys this spring put a majority of large companies — around 72% — with at least one AI workload running in production, yet only a small slice describe themselves as actually ready to operate agents safely at scale. The money is real. The control is mostly aspirational.",
						"That gap is the story of agentic AI in 2026. Not capability — we crossed the 'can it do the task' threshold months ago — but accountability. And it is starting to bite.",
					},
				},
				{
					Heading: "Why agents are harder to govern than apps",
					Paragraphs: []string{
						"The instinct is to treat an AI agent like any other piece of software: deploy it, monitor it, patch it. But an agent is not really an app. An app waits to be called. An agent acts — it makes decisions, calls tools, moves data, spends money, and chains those actions together in ways no one fully scripted in advance. To do its job, it needs credentials, permissions, and access to systems. In other words, an agent behaves far more like an employee than like a feature.",
						"And here is the uncomfortable part: most companies have a mature playbook for onboarding, permissioning, supervising, and offboarding human employees. They have almost none of that for software that behaves like one. Agents get spun up in a department, handed an API key and a service account, wired into a few internal systems, and then quietly multiply. Nobody maintains the org chart. Nobody runs the exit interview. The industry now has a name for the result: agent sprawl — a fast-growing population of semi-autonomous workers with real access and no clear system of record.",
						"Kyndryl and AWS, expanding their agentic AI partnership this spring, pointed at the same wound from the customer side: by their accounting, roughly 68% of customers were not realizing the AI benefits they expected. The bottleneck they describe is not model quality. It is the unglamorous plumbing — identity, integration, reliability, and the ability to verify that an agent did what it was supposed to and nothing more.",
					},
				},
				{
					Heading: "The market's answer: a control plane for non-human workers",
					Paragraphs: []string{
						"Where there is a gap this expensive, a software category forms to fill it. In 2026 that category has a shape and a name: the agent control plane.",
						"The clearest example landed this month. WSO2 made its Agent Manager generally available in June after a spring beta, pitching it bluntly as a way to 'tame AI agent sprawl.' It is an open control plane — released under the Apache 2.0 license, explicitly to avoid lock-in — that gives enterprises one place to identify, govern, secure, and scale agents, whether those agents run inside the platform or out in the wild. The key idea is a centralized system of record: every agent gets an identity, behavior gets observed through full-stack monitoring, and policies get enforced from a single console. It is, more or less, an HR department and a security badge office for software.",
						"The infrastructure giants are circling the same need from the hardware side. At HPE Discover in Las Vegas on June 16, HPE and NVIDIA expanded their AI Factory partnership around a pointed phrase — bringing agentic AI 'into production with security, governance, scale, and sovereignty' — anchored by the new NVIDIA Vera CPU aimed at agent orchestration and the NVIDIA Agent Toolkit for managing autonomous agents in production (with broader toolkit support slated for later in the year). Cognizant, separately, has been wiring ServiceNow's AI agents to interoperate with its own Neuro platform so agents can cross vendor boundaries without losing oversight. Different layers of the stack, same realization: the differentiator is no longer the agent. It's the cage you can safely run it in.",
					},
				},
				{
					Heading: "The warning under the boom",
					Paragraphs: []string{
						"If you want a single data point that captures the stakes, here it is: Gartner predicts that more than 40% of agentic AI projects will be canceled by the end of 2027. The drivers it cites are not exotic. Escalating costs. Unclear business value. Inadequate risk controls. In other words, the same governance gap, showing up later as a budget line that gets killed.",
						"That is the quiet logic connecting this year's two big numbers. The $206 billion pouring into agents and the wave of projects expected to die are not contradictory — they are cause and effect. Spending races ahead of control; control never catches up; the projects that can't prove they're safe and worth it get cut. The companies that survive the 2027 culling will not necessarily be the ones with the smartest agents. They'll be the ones that, somewhat unfashionably, treated governance as a first-class feature instead of a thing to bolt on later.",
						"There's a tidy historical rhyme here. We spent the last two decades learning that the hard part of cloud computing wasn't the compute — it was identity, access management, and knowing who could touch what. Agentic AI is now speed-running the same lesson, except the 'who' is no longer a person. It's a piece of software that can open a ticket, query a database, and move money while you sleep.",
						"The question worth carrying out of 2026, then, isn't whether your company will deploy AI agents. It almost certainly will; the spending curve says so. The question is the one every manager eventually has to answer about any worker with real power and a key to the building: Do you actually know what it's doing — and could you stop it if you needed to?",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Sources: Gartner AI agent software spending forecast ($86.4B 2025 to $206.5B 2026 to $376.3B 2027; total AI spend +47% 2026; more than 40% agentic projects canceled by 2027); WSO2 Agent Manager GA June 2026; HPE Discover Las Vegas June 16, 2026 HPE and NVIDIA AI Factory expansion; Kyndryl/AWS agentic partnership.",
						"Author article handoff: https://docs.google.com/document/d/1KQcS957PvC-IA39yHAaj1b2_Zq4vWrAgtB1Z4u7varw/edit",
					},
				},
			},
		},
	}, posts...)
}
