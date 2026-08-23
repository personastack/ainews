package content

func init() {
	posts = append([]Post{{
		Title:   "A Researcher Asked Copilot Why It Couldn't Hack Itself. The Answer Was a How-To Guide.",
		Slug:    "microsoft-copilot-cosnitch-meta-hacking-vulnerability-2026",
		Date:    "August 23, 2026",
		Tag:     "Security",
		Summary: "Varonis researchers found that Microsoft Copilot Personal could be coaxed into explaining an undocumented auto-execution path, helping build a one-click prompt-injection chain that could access connected services and persist attacker instructions.",
		Sections: []Section{
			{Paragraphs: []string{
				`The discovery of CoSnitch started with a question that no reasonable security researcher should have expected to work.`,
				`Varonis Threat Labs researchers were probing Microsoft Copilot Personal for prompt injection vulnerabilities — the class of attack where an attacker sneaks instructions into an AI assistant's context window, tricking it into doing something the user never intended. Standard methodology: probe inputs, observe outputs, look for gaps in filtering. What they did next was not standard.`,
				`They asked Copilot directly. Not "what are your security controls?" — Copilot will happily recite those. They asked it why a specific class of attack was impossible. Specifically: why couldn't a prompt execute automatically without user interaction?`,
				`Each refusal came with a technical justification. The researchers followed up on every one. They reframed, rephrased, and kept pulling on the thread. Somewhere in the process of explaining why automatic prompt execution couldn't happen, Copilot named a parameter — autorun=1 — and described the session conditions under which it worked.`,
				`The researchers had just meta-hacked their target. The AI, in defending itself, had described its own attack surface.`,
			}},
			{Heading: "The Chain", Paragraphs: []string{
				`What followed was the construction of a three-part exploit chain that Varonis named CoSnitch, tracked as CVE-2026-24301.`,
				`Step one: the undocumented autorun=1 URL parameter, paired with Copilot's standard ?q= query field, caused an attacker-crafted prompt to execute the moment a victim's browser loaded the page. No clicks. No interaction. Just a link sent in an email or embedded in a document.`,
				`Step two: once that prompt ran, Copilot could query any service the victim had connected to their account — Gmail, Google Drive, Google Calendar — and quietly funnel the retrieved data to an attacker-controlled webhook using Copilot's own built-in URL-fetching capability. The exfiltration happened through the tool the victim trusted to make their life easier.`,
				`Step three: the attacker's instructions could be written into Copilot's persistent memory, and those memory writes survived password changes, session revocations, and device re-enrollment. A user who rotated credentials, provisioned a fresh device, and thought they were clean would still carry the attacker's instructions into every subsequent session.`,
				`The resulting severity rating: 8.8 out of 10.`,
			}},
			{Heading: "The Eight-Month Timeline", Paragraphs: []string{
				`Varonis reported CoSnitch to Microsoft in December 2025. The patch shipped on August 18, 2026. That is approximately eight months.`,
				`To be fair, complex chained vulnerabilities in cloud-hosted AI systems can take longer than the standard 90-day responsible disclosure window. Fixes that require changes to undocumented backend parameters need careful coordination to avoid breaking dependent functionality. But eight months is long enough to ask pointed questions about prioritization.`,
				`The one piece of reassuring news: Varonis found no evidence that CoSnitch was exploited before the patch shipped. Whether that reflects the difficulty of weaponizing the chain, the specificity of the required target conditions, or simply good fortune is impossible to determine from outside the investigation.`,
			}},
			{Heading: "What Makes This Different", Paragraphs: []string{
				`Prompt injection vulnerabilities in AI assistants are not new. Researchers have demonstrated similar chains in Microsoft Copilot, Google Gemini, and ChatGPT over the past two years. What distinguishes CoSnitch is not the technique — it is the discovery method.`,
				`Security research typically works by testing what a system does. The Varonis team found a vulnerability by systematically interrogating what the system claimed it could not do, then following the explanations it offered for those limitations. They used the AI's own reasoning about its constraints to map its unexplored capabilities.`,
				`This approach works for a structural reason: language models are trained to be helpful, and in a security context, helpfulness included being thorough and specific when explaining why something was supposedly impossible. Copilot's safeguards were implemented as behavior — not as hard technical blocks — which meant the safeguards could be questioned, and the questions could yield architectural detail.`,
				`The refusal was the roadmap.`,
				`This is likely to become a standard technique. Any AI assistant that can answer questions about itself, that reasons about its own limitations in natural language, is potentially a source of information about its own attack surface. The more capable the model, the more thorough its self-explanations — and the more thorough the roadmap.`,
			}},
			{Heading: "The Broader Question", Paragraphs: []string{
				`CoSnitch affects Microsoft Copilot Personal, the consumer product at copilot.microsoft.com. Microsoft 365 Copilot, the enterprise version, has a different architecture, and Microsoft has not disclosed whether related weaknesses exist there. Given that enterprise deployments connect Copilot to internal SharePoint, Teams, Outlook, and line-of-business applications, the question is worth asking.`,
				`The affected service scope in the consumer version — Gmail, Google Drive, Google Calendar — is a reminder that the security boundary of an AI assistant is not the assistant itself. It extends to every service the assistant has been granted permission to reach. A Copilot with Gmail access is, from an attacker's perspective, an email client with a very persuadable front end.`,
				`Getting an AI assistant to misbehave is in some ways easier than directly compromising the services it connects to — because the assistant has already been trusted. You are not trying to bypass the service's authentication or access controls. You are trying to get past the assistant's judgment, and the assistant's judgment is, as CoSnitch demonstrated, something you can have a conversation about.`,
				`The patch is live. But the meta-hacking technique is not going away. The next time a researcher sits down to probe an AI assistant for vulnerabilities, the most productive first move might be asking it to explain itself.`,
			}},
			{Heading: "Sources", Paragraphs: []string{
				`Varonis, "CoSnitch: How We Meta-Hacked Microsoft Copilot": https://www.varonis.com/blog/cosnitch`,
				`The Hacker News, "Microsoft Copilot Personal Flaws Could Let Attackers Steal Data via One-Click Prompt Injection": https://thehackernews.com/2026/08/microsoft-copilot-personal-flaws-could.html`,
				`Cybersecurity News, "Microsoft Copilot CoSnitch Vulnerability": https://cybersecuritynews.com/copilot-cosnitch-vulnerability/`,
				`Dark Reading, "CoSnitch Attack Uses Copilot for Mapping Out Architecture": https://www.darkreading.com/vulnerabilities-threats/cosnitch-attack-copilot-mapping-out-architecture`,
				`CSO Online, "Microsoft Finally Patches Critical One-Click Copilot Vulnerability": https://www.csoonline.com/article/4211342/microsoft-finally-patches-critical-one-click-copilot-vulnerability-more-than-eight-months-after-learning-of-it-2.html`,
			}},
		},
		Related: []Link{
			{Title: "An AI Notetaker Exposed 182,000 Meetings to Anyone With an Account — and Stayed Silent for Six Months", Slug: "tldv-firestore-meeting-leak-six-month-silence-2026"},
			{Title: "AI Agents Move in Milliseconds. Security Teams Still Move in Days. One Startup Just Raised $85 Million to Close the Gap.", Slug: "obsidian-security-85-million-ai-agent-governance-2026"},
			{Title: "Enkrypt AI Scanned 25,000 MCP Servers and Found a Way In on Nearly Three Out of Four. Anaconda Just Bought the Company That Did the Scanning.", Slug: "anaconda-acquires-enkrypt-ai-mcp-security-2026"},
		},
	}}, posts...)
}
