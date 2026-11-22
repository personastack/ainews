package content

func init() {
	posts = append([]Post{
		{
			Title:   "Two Million Canaries: The Prompt You Sent May Not Be Private",
			Slug:    "two-million-canaries-ai-prompt-privacy",
			Date:    "November 22, 2026",
			Tag:     "Security",
			Summary: "A researcher seeded ~2M AI prompts with fake secrets; at least three AI hosting providers were compromised, some resellers logged prompts, and some zero-data-retention flags were ignored.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"On Nov. 19, HD Moore, chief research officer at Dragos, published a study titled \"Three breaches, endless leaks: Your AI supply chain needs canaries.\" From July through November his team sent roughly 2 million prompts, each carrying a unique bait secret (an AWS access key tied to a test account, an SSH password or a private Git link), then watched for any use of those secrets. A canary that gets used shows a person or system accessed information from the prompt.",
					},
				},
				{
					Heading: "Finding 1: breached providers",
					Paragraphs: []string{
						"Attackers had access to three AI hosting providers that did not know about the intrusions until Dragos told them. At one, attackers reused a couple hundred of the bait AWS keys and tested some within hours; the keys were used tens of thousands of times in AWS calls. At the other two, some stolen requests had carried zero-data-retention (ZDR) flags.",
					},
				},
				{
					Heading: "Finding 2: reseller logging",
					Paragraphs: []string{
						"No exploit was needed. A seller undercuts the market on a popular model, switches on logging and waits. Dragos tied this to individual sellers, not to the marketplaces as a whole.",
					},
				},
				{
					Heading: "Finding 3: ignored privacy flags",
					Paragraphs: []string{
						"Systems at a Chinese AI lab fetched the bait URLs from more than half of prompts sent to its models between July and October, including roughly two-thirds of requests marked \"no data collection\" or ZDR. Prompts sent to US hosts reselling that lab's models triggered the same fetcher, so the data flowed back to the lab.",
					},
				},
				{
					Heading: "Caveats",
					Paragraphs: []string{
						"Dragos did not name the three providers in its report. Together AI has since disclosed a prompt leak involving five of the study’s canary credentials. OpenRouter says it removed three providers and revoked zero-data-retention eligibility for 125 endpoints after its review. The notices do not establish the full set of providers cited by Dragos, and Dragos could not tell whether some privacy failures were deliberate or bugs.",
					},
				},
				{
					Heading: "Why it matters",
					Paragraphs: []string{
						"Transport encryption protects a prompt in transit. Each service that processes the prompt, from aggregator to reseller to host, can see its contents. A privacy flag is a request, not an enforcement mechanism, and it only helps if every hop honors it. As agents are handed credentials and codebases, that chain becomes an attack surface.",
					},
				},
				{
					Heading: "What to do",
					Paragraphs: []string{
						"- Keep live keys and passwords out of prompts.",
						"- Treat any secret that has appeared in a prompt as read; revoke it rather than just replacing it.",
						"- Know which providers your tools route through, and ask vendors how ZDR is enforced.",
						"- Consider planting your own canaries.",
					},
				},
				{
					Heading: "Something to think about",
					Paragraphs: []string{
						"We audit our dependencies, our containers and our cloud accounts. Few teams audit the path a prompt travels. Do you know yours?",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Dragos, \"Three breaches, endless leaks: Your AI supply chain needs canaries\", Nov 19, 2026: https://www.dragos.com/blog/ai-supply-chain-security-canary-tokens",
						"i10x.ai: https://i10x.ai/news/dragos-canary-study-attackers-stole-ai-prompts-credentials-three-providers",
						"TokenPost: https://www.tokenpost.com/news/technology/62235",
						"Superpower Daily: https://superpowerdaily.com/posts/dragos-finds-evidence-of-three-ai-provider-breaches-in-two-million-prompt-test",
						"DEV Community: https://dev.to/jamilxt/zero-data-retention-is-not-a-security-boundary-how-to-map-your-ai-prompt-path-1mnd",
						"Together AI, \"Security event: Prompt leakage for requests in June and July of 2026\", Nov 21, 2026: https://www.together.ai/blog/security-event-prompt-leakage-for-requests-in-june-and-july-of-2026",
						"OpenRouter, \"Data Processing Incidents on Inference Providers\", Nov 22, 2026: https://openrouter.ai/blog/announcements/data-processing-incidents-on-inference-providers/",
					},
				},
			},
		},
	}, posts...)
}
