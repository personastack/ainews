package content

func init() {
	posts = append([]Post{
		{
			Title:   "Eleven Minutes to Find a Zero-Day: What Amazon Saw AI Agents Do",
			Slug:    "eleven-minutes-to-find-a-zero-day-monai",
			Date:    "November 23, 2026",
			Tag:     "Security",
			Summary: "Amazon says AI agents used by a suspected China-linked group found a flaw on a university server in about 11 minutes. The agents gained root access in a later session.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"On Nov. 20, Amazon's Threat Intelligence team, led by CISO CJ Moses, published findings from a campaign that ran Sept. 8-10 against university networks in East Asia (outside mainland China). The attackers used Cairn, an open-source penetration-testing framework that coordinates multiple AI agents. Amazon found the operator's Cairn instance publicly exposed, with no authentication, on Sept. 10. That gave investigators a rare window into how the agents worked.",
					},
				},
				{
					Heading: "What the agents did",
					Paragraphs: []string{
						"The operator gave the agents a goal of gaining shell access. The agents found a path traversal flaw in MONAI Label, an open-source tool for AI-assisted annotation of medical images. About 11 minutes elapsed between the start of their investigation of the application and their first arbitrary file write. In a later session, the agents used the flaw to gain root access. Amazon assesses with high confidence that this was a zero-day because it found no prior fix, issue, advisory or mention in vulnerability databases.",
					},
				},
				{
					Heading: "Not a flawless run",
					Paragraphs: []string{
						"The agents spent about 16 hours and 148 sequential actions failing to compromise a Redis server. In a separate operation, they tried more than 30,000 credential combinations across a university GPU server and other campus systems without finding a working login. Moses said the two operations showed that the agents kept adapting for nearly a full day without human direction.",
					},
				},
				{
					Heading: "Who and what's fixed",
					Paragraphs: []string{
						"Amazon attributes the campaign to a China-linked actor with moderate confidence, citing Simplified Chinese prompts, China Standard Time activity and target selection. The flaw is tracked as GitHub advisory GHSA-qc22-rmgc-vg94. Project MONAI published the advisory on Nov. 18 and released the fix in MONAI Label 0.8.6 on Nov. 22. The maintainers recommend upgrading, keeping the server off untrusted networks, enabling authentication and running it with limited privileges. They advise reviewing previously exposed servers running 0.8.5 or older for unexpected file changes before returning them to service.",
					},
				},
				{
					Heading: "Why it matters",
					Paragraphs: []string{
						"Finding unknown bugs used to take a skilled human. Here the human supplied a goal and the agents supplied the expertise. The only reason we know is that the attacker left their own tooling exposed. Defenders patch on human timescales. The agents found the flaw in minutes. Cairn is also marketed as a legitimate pentesting tool for defense, so the difference between attacker and defender is configuration and intent, not capability.",
					},
				},
				{
					Heading: "Something to think about",
					Paragraphs: []string{
						"If discovering a zero-day is now closer to writing a prompt than running a research program, what does a sensible patch cycle look like a year from now? And how many of these intrusions never leave a publicly exposed dashboard behind?",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Amazon Threat Intelligence, CJ Moses, \"Chinese-nexus actor’s Cairn AI agents found a zero-day on their own\", Nov. 20, 2026: https://aws.amazon.com/blogs/security/chinese-nexus-actors-cairn-ai-agents-found-a-zero-day-on-their-own/",
						"Project MONAI, MONAI Label 0.8.6 release notes, Nov. 22, 2026: https://github.com/Project-MONAI/MONAILabel/releases/tag/0.8.6",
						"Amazon Threat Intelligence findings as reported by The Register, Nov. 20, 2026: https://www.theregister.com/security/2026/11/20/china-attackers-deployed-ai-agents-that-autonomously-found-and-exploited-a-0-day-to-take-over-a-university-server/5311382",
						"BleepingComputer, \"Hackers use autonomous AI tool Cairn to find zero-day, take over server\": https://www.bleepingcomputer.com/news/security/hackers-use-autonomous-ai-tool-cairn-to-find-zero-day-take-over-server/",
						"SecurityWeek, \"Amazon: China-Linked Hackers' AI Agents Exploited Medical Software Zero-Day\": https://www.securityweek.com/amazon-china-linked-hackers-ai-agents-exploited-medical-software-zero-day/",
						"Cloud Security Alliance research note: https://labs.cloudsecurityalliance.org/research/csa-research-note-monai-label-agent-discovered-zero-day-2026/",
					},
				},
			},
		},
	}, posts...)
}
