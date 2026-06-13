package content

func init() {
	posts = append([]Post{
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
