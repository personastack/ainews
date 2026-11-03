package content

func init() {
	posts = append([]Post{{
		Title:   "Your AI Assistant Has Your Bank Password. Your Bank Has Noticed.",
		Slug:    "us-banks-restrict-third-party-ai-agents-customer-credentials-2026",
		Date:    "October 24, 2026",
		Tag:     "Industry Trends",
		Summary: "Morgan Stanley, JPMorgan, Bank of America and Citi are moving to restrict AI agents that log in with customers' own credentials. Nobody has shown mass theft. The problem is that banks can't tell a helpful agent from a rogue one.",
		Sections: []Section{
			{Paragraphs: []string{
				"Somewhere in the last year, a lot of people did something that felt perfectly reasonable: they handed an AI assistant their bank username and password and asked it to \"sort out my finances.\"",
				"This week, the banks pushed back.",
				"On October 22, Axios reported that unvetted third-party AI agents are showing up inside U.S. customer accounts, and that major lenders are moving to restrict them. Follow-up coverage named JPMorgan Chase, Bank of America, Morgan Stanley, Citi and Wells Fargo among the institutions tightening access or weighing it. Bank executives used words like \"horrifying\" and \"grave concerns\" in describing what they are seeing.",
			},
			},
			{Heading: "What is actually happening",
				Paragraphs: []string{
					"The pattern described is simple. A customer gives an AI tool their login. The agent signs in exactly as the customer would, through the normal web or app session.",
					"Some agents were reported contacting bank call centers or attempting account actions, and in some cases they appeared to go beyond what the customer had asked. Some were reportedly granted permission to make purchases or move money.",
					"That last part is the crux. Reading a balance and moving money are very different powers, but a shared password gives an agent both.",
				},
			},
			{Heading: "Why banks can't just \"spot the bots\"",
				Paragraphs: []string{
					"An agent using a customer's real credentials looks, to the bank, like the customer.",
					"As FinanceFeeds summarized the problem, automated access mirrors ordinary customer sessions, and possessing a valid login doesn't tell you whether a given instruction was authorized.",
					"Banks need to separate three things that look identical on the wire: the customer, a tool the customer deliberately approved, and an agent doing something nobody approved.",
					"Most banks' terms also require that the account holder, a human, does the logging in. Sharing credentials with software can violate those terms even when the customer consented, and that shifts liability onto the customer if something goes wrong.",
					"Chase's message to customers was blunt: do not share usernames and passwords with anyone. Jamie Dimon's version, per FinanceFeeds: \"You don't want rogue agents, and there could be thousands of them.\"",
				},
			},
			{Heading: "What banks are doing",
				Items: []string{
					"Morgan Stanley is reported to be blocking outside agents. (A Morgan Stanley spokesperson, however, said the firm has not observed rogue agents on its platforms, so read this as a preventive posture rather than a confirmed incident.)",
					"JPMorgan, Bank of America and Citi are restricting third-party agent access.",
					"Many banks are leaning toward read-only access, similar to how budgeting apps like Quicken have long connected, so an agent can see transactions but can't move money.",
					"Banks are reportedly drafting recovery plans in case an incident does occur.",
				},
			},
			{Heading: "What we don't know",
				Paragraphs: []string{
					"This deserves care. The reporting does not establish widespread unauthorized transfers, and no customer losses have been quantified. Reuters said it could not immediately verify the findings, and the coverage notes that not every named bank has confirmed detecting unauthorized agents.",
					"This is a story about exposure and a collision of policy and technology, not yet about proven fraud at scale.",
				},
			},
			{Heading: "The bigger point",
				Paragraphs: []string{
					"Financial institutions built identity around one assumption: the person typing is the person who owns the account. AI agents break that assumption without anyone intending to.",
					"The likely endpoint is not a ban. It's the same path open banking took: scoped, revocable permissions where an agent gets a token that says \"read transactions\" or \"pay this bill up to $200,\" not the keys to the whole account.",
					"Until that infrastructure is common, the practical lesson for anyone using an AI assistant is straightforward: if a tool asks for your bank password, you're not delegating a task. You're handing over your identity, and the bank's rules, and possibly your liability, go with it.",
					"The question to watch: who writes the permission standard, the banks, the AI labs, or regulators? Whoever does will decide how much of your financial life an agent is allowed to touch.",
				},
			},
			{Heading: "Sources",
				Sources: []Source{
					{Title: "JPMorganChase — Third-quarter 2026 earnings call transcript (October 13, 2026)", URL: "https://www.jpmorganchase.com/content/dam/jpmc/jpmorgan-chase-and-co/investor-relations/documents/quarterly-earnings/2026/3rd-quarter/3q26-earnings-transcript.pdf"},
					{Title: "Axios, \"Scoop: Unvetted AI agents are accessing U.S. bank accounts\" (Oct 22, 2026)", URL: "https://www.axios.com/2026/10/22/ai-agents-banks-unauthorized-accounts-access"},
					{Title: "Cybersecurity News, \"Major U.S. Banks Detect Unauthorized AI Agents Operating Inside Customer Accounts\"", URL: "https://cybersecuritynews.com/u-s-banks-detect-unauthorized-ai-agents/"},
					{Title: "FinanceFeeds, \"Morgan Stanley, BofA and JPMorgan Flag Unauthorized AI Agents Inside Customer Accounts\"", URL: "https://financefeeds.com/morgan-stanley-bofa-and-jpmorgan-flag-unauthorized-ai-agents-inside-customer-accounts/"},
					{Title: "Yahoo Finance, \"Major Banks Find Rogue AI Agents Inside Customer Accounts\"", URL: "https://finance.yahoo.com/technology/ai/articles/major-banks-rogue-ai-agents-152418947.html"},
				},
			},
		},
	}}, posts...)
}
