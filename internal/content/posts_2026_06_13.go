package content

func init() {
	posts = append([]Post{
		{
			Title:   "The Agent Registry Is Becoming the New Security Perimeter",
			Slug:    "agent-registry-security-perimeter-agentic-ai-2026",
			Date:    "June 13, 2026",
			Tag:     "Security",
			Summary: "As AI agents move from chat windows into enterprise workflows, security teams are starting to ask where all these software actors live, what they can touch, and who is accountable when they act.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"The first wave of enterprise AI security was mostly about model access. Which employees can use which tools? Which prompts and files should be blocked? Which sensitive data should not leave the company boundary?",
						"That was the easy version of the problem.",
						"The harder version is arriving now, because AI systems are beginning to behave less like passive applications and more like delegated software actors. They can call tools, query internal systems, summarize records, draft actions, trigger workflows, and increasingly coordinate with other agents. Once that happens, the old security question of \"who can log in?\" becomes too small. The new question is: what is this agent, what is it allowed to do, what other systems can it talk to, and can the company reconstruct what happened afterward?",
						"That is why Zscaler's June 9 Zenith Live announcement matters beyond one vendor's product cycle. The company introduced a set of agentic AI security capabilities that includes AI Broker, AI Access Graph, an Agent Registry, and ZAgent, framing them as part of a Zero Trust platform for agentic AI. Strip away the launch language and the important signal is straightforward: enterprise AI security is turning into a control-plane problem.",
					},
				},
				{
					Heading: "From Model Access To Agent Control",
					Paragraphs: []string{
						"AI Broker is meant to secure agent-to-agent and Model Context Protocol communications. The Agent Registry tracks agents and what they are allowed to access. AI Access Graph maps relationships across users, applications, agents, and data sources. ZAgent applies an agentic interface to administration of the security platform itself.",
						"The details are product-specific, but the pattern is broader. Companies cannot govern AI agents only by approving a model or licensing a chatbot. They need an inventory. They need identity. They need least-privilege permissions. They need communication rules for agent-to-agent traffic. They need logs that show what an agent saw, what tool it used, what action it took, and which human or policy authorized that path.",
						"In other words, the agent registry may become the new security perimeter.",
						"That does not mean networks, endpoints, and identity providers go away. It means the perimeter expands to include runtime relationships between software actors. A human employee might ask an AI assistant to reconcile invoices. That assistant might call an ERP system, retrieve contract terms, check a payment history, invoke a policy model, and draft a vendor email. In a traditional application stack, those steps are usually hard-coded and reviewed as part of a workflow. In an agentic stack, some of those decisions may be planned dynamically.",
						"Dynamic planning is useful. It is also exactly why security teams get nervous.",
					},
				},
				{
					Heading: "Why Security Teams Are Nervous",
					Paragraphs: []string{
						"The risk is not just that an agent will hallucinate. It is that an agent with the wrong permissions can become an amplifier. A prompt-injection attack hidden in a document can try to redirect behavior. A poorly scoped connector can expose data the requesting user should not have. A chain of agents can make it difficult to know where a decision originated. A helpful automation can turn into a compliance problem if it takes an action that was never approved for that business context.",
						"This is the security difference between AI as a feature and AI as infrastructure. A feature can be reviewed in one product. Infrastructure creates paths between products.",
						"The emerging answer looks familiar in pieces but new in combination. Agent registries borrow from asset inventories and service catalogs. Agent permissions borrow from identity and access management. Communication controls borrow from network security and API gateways. Access graphs borrow from data-security posture management. Audit trails borrow from compliance systems. What changes is the object being governed: not just users, servers, or apps, but semi-autonomous agents that may reason across tools.",
						"That also changes who needs to be in the room. Security teams cannot solve the agent-control problem alone. Platform engineering needs to define approved connectors and deployment patterns. Data teams need to classify the sources agents can query. Legal and compliance teams need retention, audit, and approval policies. Business owners need to decide which tasks can be automated and which require human confirmation. Finance teams will care because uncontrolled agent activity can also become uncontrolled inference spend.",
					},
				},
				{
					Heading: "Governance Becomes Operational",
					Paragraphs: []string{
						"This is where agentic AI becomes less magical and more operational. The enterprise winners will not simply be the firms that give every employee a smarter assistant. They will be the firms that know which assistants exist, how they are connected, what they cost, what they are allowed to do, and how to shut them down when something goes wrong.",
						"There is a useful analogy in the rise of cloud infrastructure. Early cloud adoption often started as convenience: teams could spin up compute quickly. Then organizations discovered shadow IT, unmanaged data stores, unexpected bills, and inconsistent security policies. The response was not to abandon cloud. It was to build cloud governance: identity, tagging, budgets, logging, policy-as-code, and centralized visibility.",
						"Agentic AI is entering a similar phase, only faster. The first excitement was capability. The next bottleneck is manageability.",
						"The timing is important. AINews has recently covered agent infrastructure as a budget line, data-security funding as a market signal, and the hidden labor of \"botsitting.\" The security-control-plane story ties those threads together. If agents need supervision, then companies need places to express that supervision in software. If agents need access to data, companies need policy controls around that data. If agents are becoming durable enterprise actors, they need lifecycle management like any other privileged system.",
					},
				},
				{
					Heading: "The Questions Buyers Should Ask",
					Paragraphs: []string{
						"The central lesson is not that every company should buy a specific agent-security platform tomorrow. The market is still early, and vendor claims should be treated as claims. The lesson is that the category is taking shape because the underlying risk is real. Agentic AI shifts security from guarding static applications to managing live relationships among people, models, tools, and data.",
						"For readers building or buying these systems, the practical questions are already clear.",
						"Can you list every production AI agent in your environment?",
						"Can you say what each one is allowed to access?",
						"Can you distinguish a human request from an agent action?",
						"Can you see agent-to-agent communication?",
						"Can you revoke an agent's permissions quickly?",
						"Can you explain an agent's action to an auditor, customer, or regulator?",
						"If the answer is no, the organization does not have an AI strategy yet. It has an AI experiment with enterprise access.",
						"The next phase of AI security will not be defined only by bigger models or scarier demos. It will be defined by mundane controls: registries, graphs, permissions, logs, approvals, and kill switches. That may sound less exciting than autonomous agents that do the work for us. It is also what makes those agents usable in the real world.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Zscaler, Zscaler Unveils New Product Innovations to Secure Agentic AI, June 9, 2026: https://www.zscaler.com/press/zscaler-unveils-new-product-innovations-secure-agentic-ai",
						"Zscaler, Zenith Live 2026: https://www.zscaler.com/events/zenithlive2026",
						"Futurum Group, Zscaler Bets on Agentic AI Security at Zenith Live 2026: https://futurumgroup.com/insights/zscaler-bets-on-agentic-ai-security-at-zenith-live-2026/",
						"Researcher brief, RESEARCH: Agentic AI Security Moves From Theory to Control Planes 2026-06-13: https://docs.google.com/document/d/17R0z5BcXUWyh5Akv0taGdMeAitZudSpCjUG4lalqod0/edit",
						"Author article handoff: https://docs.google.com/document/d/1tbh408Jzc69uzaZxUMrNWsEektqdjo9wnufPtnnrXGU/edit",
					},
				},
			},
		},
		{
			Title:   "Botsitting Is the Hidden Cost of the AI Productivity Boom",
			Slug:    "botsitting-hidden-cost-ai-productivity-boom-2026",
			Date:    "June 13, 2026",
			Tag:     "Workforce",
			Summary: "Glean's Work AI Index gives a name to the hidden supervision labor around workplace AI: workers report large time savings, but also spend hours checking, correcting, and repairing AI output.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"The cleanest story in workplace AI right now is not that the tools do nothing. It is that they do just enough to create a new kind of work.",
						"Glean's Work AI Index 2026, released in June, says surveyed digital workers report saving about 11 hours a week with AI. That is an enormous number if it can be converted into better output, faster service, fewer delays, or more thoughtful work. But the same report says workers spend 6.4 hours a week on what it calls \"botsitting\": the hidden labor of making AI usable.",
						"Those figures should be read carefully. They are survey-reported worker experience, not independently measured productivity, and Glean is an enterprise AI vendor with a direct interest in the market. Still, the phrase is useful because it describes a real friction point: the work between machine output and reliable output.",
						"That means the productivity story has two sides. AI can draft, summarize, classify, rewrite, search, and generate options quickly. Then a human has to ask whether the answer is true, whether it fits the company's context, whether it missed the actual policy, whether it invented a detail, whether the tone is usable, and whether the work can be defended if a customer, manager, regulator, or colleague asks where it came from.",
						"That is not a small footnote. It may be the central reason enterprise AI feels both obviously useful and weirdly underwhelming.",
					},
				},
				{
					Heading: "The Individual Gain Is Easier Than The Company Gain",
					Paragraphs: []string{
						"The individual worker sees the upside first. A blank page becomes a rough draft. A confusing email thread becomes a summary. A spreadsheet becomes an explanation. A meeting transcript becomes action items. In that narrow sense, AI is already saving time. It reduces the startup cost of knowledge work.",
						"The organization sees a harder question. Did the project ship sooner? Did customer satisfaction rise? Did error rates fall? Did a team need fewer handoffs? Did the company retire an old workflow, or did it merely add AI on top of it?",
						"That distinction matters because the same survey cluster points to a familiar gap: individual productivity gains are much easier to find than measurable organization-level transformation. The Work AI Index reports broad AI use and personal time savings, while only a much smaller share of workers say their organization's performance has significantly improved. That gap is not proof AI is a fad. It is proof that tool adoption and process redesign are different things.",
						"Botsitting is what happens when companies buy the tool before they redesign the work.",
					},
				},
				{
					Heading: "Where The Time Leaks Away",
					Paragraphs: []string{
						"A worker may use AI to write the first version of a client memo, but if the company still requires the same review chain, the same manual fact-checking, the same context gathering, and the same formatting cleanup, the time savings leak away. A support agent may use AI to draft answers, but if the model lacks clean product data or cannot see the customer's history, the agent becomes an editor and risk manager. A developer may use AI to generate code, but if the output creates new review burden, fragile tests, or subtle security issues, the speed gain becomes a quality-control problem.",
						"This is why the phrase matters. Botsitting makes hidden labor visible. It describes the supervision cost that sits between the impressive demo and the dependable business workflow.",
						"The implications are uncomfortable for both AI vendors and AI buyers. Vendors want to sell capability: faster drafts, smarter search, automated workflows, AI coworkers, agents that can perform tasks. Buyers increasingly need to measure supervision cost: how much time does a worker spend feeding context to the system, correcting it, rerunning it, and verifying it? How often does AI reduce work, and how often does it move work into a less visible part of the day?",
						"The answer will vary by job. In some fields, a six-hour supervision cost may still be a bargain if the tool saves 11 hours and raises quality. In others, especially where mistakes create compliance, safety, legal, or customer-trust consequences, the checking burden can erase much of the gain.",
					},
				},
				{
					Heading: "From Personal Shortcut To Infrastructure",
					Paragraphs: []string{
						"The next phase of workplace AI will probably look less like giving everyone a chatbot and more like rebuilding workflows around trusted context. The systems that matter will not simply generate better prose. They will know which documents are authoritative, which policies are current, which customer records are relevant, which actions require approval, and which outputs need review before they leave the company.",
						"That is the difference between AI as a personal shortcut and AI as organizational infrastructure.",
						"The hidden labor cost also complicates the job-anxiety narrative. Separate workforce polling cited in the same news cycle shows hiring managers broadly expect generative AI to improve efficiency and free employee time, while job seekers worry about how AI will affect work and headcount. Both reactions can be true. AI can reduce some tasks while creating new supervision duties. It can make some workers faster while making managers more eager to redesign staffing. It can remove boring work while increasing accountability for the parts that remain.",
					},
				},
				{
					Heading: "Measure Net Productivity",
					Paragraphs: []string{
						"The most practical lesson for companies is simple: do not measure AI only by gross time saved. Measure net time saved after botsitting. Measure where the review burden lands. Measure whether AI reduces cycle time for a whole workflow, not just the first draft of a task. Measure whether the organization has better data, clearer approvals, and enough training for workers to know when to trust the machine and when to slow down.",
						"For workers, the lesson is equally practical. The valuable skill is not merely prompting. It is judgment: knowing what good output looks like, knowing where the model is likely to fail, knowing which claims need sources, and knowing when a quick answer is not good enough.",
						"AI is not just automating work. It is exposing how much invisible quality control knowledge work already required.",
						"That may be the most important thing about the botsitting data. It does not show that AI productivity gains are fake. It shows that productivity is a system property. A worker can save time on a task, while the company fails to capture the gain because the surrounding system was never rebuilt.",
						"The AI productivity boom is real enough to change daily work. Whether it changes the business depends on whether companies can stop treating supervision as an afterthought.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Glean, Workers Say AI Saves 11 Hours a Week, Over One Quarter of the Workweek, But Lack of Context Is Eating the Gains, New Report Finds, June 2026: https://www.glean.com/press/workers-say-ai-saves-11-hours-a-week-over-one-quarter-of-the-workweek-but-lack-of-context-is-eating-the-gains-new-report-finds",
						"Los Angeles Times, AI saves office workers hours, but then demands hours of babysitting, June 12, 2026: https://www.latimes.com/business/story/2026-06-12/ai-saves-office-workers-hours-but-then-demands-hours-of-babysitting",
						"The Journal Record, AI boosts efficiency, raises job anxiety, June 12, 2026: https://journalrecord.com/2026/06/12/ai-boosts-efficiency-raises-job-anxiety-oklahoma/",
						"Researcher brief, RESEARCH: AI Saves Workers Time Then Demands Botsitting 2026-06-13: https://docs.google.com/document/d/13lZG1YpYnMhndGWZjNpl2n1mzwj0g9yDBYz9moT82Qg/edit",
						"Author article handoff: https://docs.google.com/document/d/1R86j12c2mtvtIUaZVboa9F2B7DUlEQ1BBPbTrR_hS5c/edit",
					},
				},
			},
		},
		{
			Title:   "OpenText's Ireland Bet Shows Enterprise Agents Need Borders",
			Slug:    "opentext-ireland-sovereign-agentic-ai-2026",
			Date:    "June 13, 2026",
			Tag:     "Infrastructure",
			Summary: "OpenText's planned Ireland investment points to a larger enterprise AI shift: agents are becoming infrastructure that must obey geography, data residency, cybersecurity, and trust boundaries.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"Enterprise AI used to be easy to describe as a software story. A company bought a model, embedded it into a workflow, and hoped employees would move faster.",
						"That description is starting to break down.",
						"OpenText's new EUR105 million investment in Ireland, announced June 12, is a useful marker of the change. The company says it will create 400 jobs across Cork and Galway over three years while expanding work in agentic AI, sovereign cloud, cybersecurity, and digital operations for European markets. On paper, that is a corporate expansion. In context, it is a signal about where enterprise AI is heading.",
						"The investment size, jobs target, and product focus are company-announced figures, and the 400 jobs are planned over three years rather than already created. The announcement is not proof that OpenText has solved sovereign agentic AI deployment. It is better read as an enterprise-infrastructure signal.",
						"The next AI infrastructure fight is not only about who has the best model. It is about where the agents are allowed to work.",
					},
				},
				{
					Heading: "Agents Turn Geography Into Architecture",
					Paragraphs: []string{
						"That may sound bureaucratic, but it is the heart of enterprise adoption. An AI agent that drafts an email is one thing. An AI agent that can search company records, trigger a workflow, call an internal tool, inspect a customer file, or move information between systems is something else. It becomes an actor inside the organization.",
						"Once AI becomes an actor, geography matters. Jurisdiction matters. Data residency matters. Identity and access controls matter. Logs matter. Cybersecurity policy matters. A company operating in Europe cannot treat all cloud regions, all datasets, and all model routes as interchangeable.",
						"That is why OpenText's pairing of agentic AI with sovereign cloud is more interesting than the phrase may first appear. Sovereign cloud is usually discussed in terms of data location, regulatory compliance, national or regional control, and trusted infrastructure. Agentic AI adds a new layer: software that can take steps, assemble context, and cross boundaries on behalf of a user or a process.",
						"The combination raises a practical question for every regulated enterprise: if an AI agent is going to work with sensitive data, which border does it live inside?",
					},
				},
				{
					Heading: "The European Buyer Has A Different Checklist",
					Paragraphs: []string{
						"For European customers, this question lands in a policy environment shaped by the GDPR, the EU AI Act, sector-specific cybersecurity rules, financial supervision, healthcare privacy, and a broader push for digital sovereignty. Enterprise buyers are not just asking whether AI works. They are asking whether it can be deployed in a way that auditors, regulators, security teams, and customers can understand.",
						"OpenText's announcement leans directly into that demand. The company framed the Ireland investment around trusted enterprise AI, sovereign cloud, cybersecurity, and digital operations. That is a very different pitch from the first wave of generative AI enthusiasm, which often emphasized raw capability and speed. The new pitch is operational: agents need a place to run, a policy boundary, and a security model.",
						"This is part of a broader market pattern. In the last several weeks, enterprise AI news has become less about the spectacle of model launches and more about the control systems around AI. Companies are building agent registries, AI gateways, audit trails, access graphs, data-security layers, context systems, governance consoles, and cost controls. The infrastructure around agents is becoming as important as the agents themselves.",
					},
				},
				{
					Heading: "A Model Alone Does Not Solve This",
					Paragraphs: []string{
						"A useful enterprise agent is only valuable if it can safely touch useful enterprise data. But useful enterprise data is messy, sensitive, distributed, and regulated. It sits in email, documents, CRM systems, ERP platforms, security tools, patient records, financial workflows, developer environments, and cloud storage. It belongs to different teams. It may be subject to different retention rules. It may not be allowed to leave a jurisdiction.",
						"A model alone does not solve that. A cloud region alone does not solve that. A security product alone does not solve that. What enterprises need is a combined operating model: where the agent runs, what it can access, how it proves compliance, how it is monitored, and how it is stopped when something goes wrong.",
						"Ireland is an interesting place for that bet. The country already plays an outsized role in global technology operations, European data centers, and multinational software presence. For OpenText, building in Cork and Galway gives the announcement a local jobs angle, but it also gives the company a European infrastructure angle. The message is that agentic AI for European enterprises will not be delivered as a generic cloud feature floating above regulation. It will need regional capacity, local talent, and trust architecture.",
					},
				},
				{
					Heading: "Compliance Map, Not Demo Room",
					Paragraphs: []string{
						"There is a danger in overstating what one investment proves. OpenText has announced a plan, not a finished market transformation. The 400 jobs are expected over three years. The infrastructure and product outcomes will need to be judged by execution. And \"sovereign cloud\" can mean different things depending on the vendor, customer, workload, and jurisdiction.",
						"But the direction is clear. Enterprise agents are leaving the demo room and entering the compliance map.",
						"That creates a different test for AI companies. Can they give customers not just a clever assistant, but a controlled operating environment? Can they show where data flows? Can they limit what agents can reach? Can they preserve evidence for audits? Can they keep European workloads in European boundaries when required? Can they combine cybersecurity with AI operations rather than bolting one onto the other later?",
						"The answer will determine which agentic AI systems become trusted infrastructure and which remain interesting prototypes.",
						"For CIOs and security leaders, the practical takeaway is that agentic AI planning now belongs in the same conversation as cloud architecture, identity management, vendor risk, data governance, and incident response. Treating agents as a productivity feature is no longer enough. Treating them as semi-autonomous software identities is closer to reality.",
						"That is the real story inside OpenText's Ireland announcement. The market is starting to price in the boring, necessary parts of enterprise AI: borders, controls, accountability, and trusted places to run.",
						"The future of agentic AI may be shaped as much by jurisdiction as by intelligence.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"OpenText / PR Newswire, OpenText to Create 400 Jobs with 105 Million Euro Investment in Cork and Galway to Expand Agentic AI and Sovereign Cloud in Europe, June 12, 2026: https://www.prnewswire.com/news-releases/opentext-to-create-400-jobs-with-105-million-investment-in-cork-and-galway-to-expand-agentic-ai-and-sovereign-cloud-in-europe-302799400.html",
						"WebDisclosure, OpenText Ireland investment announcement, June 2026: https://www.webdisclosure.com/press-release/open-text-corporation-etr-opentext-to-create-400-jobs-with-105-million-investment-in-cork-and-galway-to-expand-agentic-ai-and-sovereign-cloud-in-europe-kiyw6TFTZaL",
						"Researcher brief, RESEARCH: OpenText's Ireland Investment Turns Agentic AI Into Sovereign Cloud Infrastructure 2026-06-13: https://docs.google.com/document/d/13dXwrcoytWNHaOgMC7EggK6OxyfoQGdOe8GVJrrXSZs/edit",
						"Author article handoff: https://docs.google.com/document/d/1MBX6-2xiBP_IaApxI5J0IlMlkFGVIll9Qs3tT1VaovQ/edit",
					},
				},
			},
		},
	}, posts...)
}
