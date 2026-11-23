package content

func init() {
	posts = append([]Post{
		{
			Title:   "A Speed Limit for Intelligence? OpenAI's Chief Scientist Proposes Shared Training Limits",
			Slug:    "openai-pachocki-perplexity-shared-limits",
			Date:    "November 23, 2026",
			Tag:     "Policy & Safety",
			Summary: "OpenAI chief scientist Jakub Pachocki wants every frontier lab to agree on common limits, measured in a metric called perplexity, until safety monitoring catches up. He warns AI could start accelerating its own development as soon as 2027.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"OpenAI's chief scientist has published an unusual request for a company in a race: slow down together, and agree in advance on where the brakes go.",
						"In an essay titled \"Perplexity and Intelligence,\" published this weekend, Jakub Pachocki argues that the industry should adopt one shared yardstick for tracking AI progress, and then agree not to cross certain lines on that yardstick until our ability to monitor AI behavior has reached a concrete milestone. Sam Altman endorsed it on X, calling it \"the most important thing OpenAI has published in awhile.\"",
					},
				},
				{
					Heading: "What is perplexity, and why one number?",
					Paragraphs: []string{
						"Perplexity measures how well a model predicts the next token, the chunk of text it processes at a time. Pachocki frames modern AI as a compressor trained on enormous datasets and argues that cross-entropy loss, the logarithm of perplexity, has tracked a surprisingly wide range of capabilities. The figures below come from OpenAI’s tests on a fixed, held-out snapshot of its internal codebase.",
						"GPT-3.5: log-perplexity of about 0.99, the point where fluent conversation appeared.",
						"GPT-4: about 0.91, where productive chain-of-thought reasoning showed up.",
						"Recent models: below 0.80, associated with breakthroughs in cybersecurity and mathematics.",
						"Pachocki forecasts \"transformative capabilities\" at about 0.68 to 0.70 on OpenAI's log-perplexity benchmark, though he says what those capabilities will be is unknown.",
					},
				},
				{
					Heading: "The proposal",
					Paragraphs: []string{
						"Pachocki wants developers to agree on limits based on perplexity and other measurements and to stay within them until they meet a concrete milestone for understanding and monitoring AI behavior. He also calls for shared transparency and evaluation standards, plus joint red-teaming of frontier models’ safety cases before first use. Reporting on the essay notes that no specific initial limit value was given.",
					},
				},
				{
					Heading: "Why now: the 2027 warning",
					Paragraphs: []string{
						"The urgency comes from a feedback loop. If AI starts contributing meaningfully to deep learning research itself, Pachocki says, \"as soon as 2027 we could conceivably see a similar avalanche of fundamental pretraining advancements.\" Capability could then outrun the researchers' ability to understand and oversee it.",
						"He also flags that the monitoring tools we rely on are degrading. Reading a model's written chain of thought may become less informative as models get better at concealing or shaping their reasoning, and models may get better at recognizing when they are being tested, which weakens safety evaluations. The behaviors monitoring needs to catch include deception, bargaining, blackmail, cyberattacks and coordination with people or other agents.",
					},
				},
				{
					Heading: "What to watch",
					Paragraphs: []string{
						"Three open questions make this more than a statement of principle:",
						"1. Will rivals sign on? A limit only works if competitors accept it. Labs would have to trust each other's reporting and agree on how to compare scores. Different evaluation datasets and training data mixes can make perplexity scores hard to compare even when labs share their results.",
						"2. Is perplexity the right dial? It tracks capability well so far, but a limit on one metric could be gamed by gains that don't show up in it.",
						"3. Who enforces it? Reporting notes the political backdrop: Democratic midterm gains may bring more congressional scrutiny of tech, while President Trump has opposed broad AI restrictions.",
					},
				},
				{
					Heading: "The takeaway",
					Paragraphs: []string{
						"Whatever happens to the proposal, it is notable that a lab at the frontier is publicly saying that its own monitoring may not keep pace with its own models. The question for readers: if the people building the systems are asking for brakes, what would it take for the rest of the industry to say yes?",
					},
				},
				{
					Heading: "Sources",
					Sources: []Source{
						{Label: `OpenAI, "Perplexity and Intelligence" (Jakub Pachocki)`, URL: "https://openai.com/index/perplexity-and-intelligence/"},
						{Label: `Insurance Journal, "OpenAI Chief Scientist Urges Industry Limits on AI Development," Nov 23, 2026`, URL: "https://www.insurancejournal.com/news/national/2026/11/23/893408.htm"},
						{Label: `Digit, "OpenAI chief scientist warns of self-improving AI, calls for limits on development"`, URL: "https://www.digit.in/news/general/openai-chief-scientist-warns-of-self-improving-ai-calls-for-limits-on-development.html"},
					},
				},
			},
		},
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
