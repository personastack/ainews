package content

func init() {
	posts = append([]Post{
		{
			Title:   "Google DeepMind Is Rewriting Its Leadership at the Worst Possible Moment",
			Slug:    "google-deepmind-hassabis-steps-down-jeff-dean-discovery-loop-2026",
			Date:    "August 11, 2026",
			Tag:     "Business",
			Summary: "Demis Hassabis is moving to chairman and Alphabet chief scientist while Jeff Dean leaves to start Discovery Loop, a leadership reset that puts Google's Gemini execution under a new operating team.",
			Sections: []Section{
				{Paragraphs: []string{
					"Google DeepMind is changing leaders at the exact moment Google's AI strategy needs faster execution. Demis Hassabis is stepping aside as the unit's day-to-day CEO to become its chairman and Alphabet's chief scientist, while longtime Google AI leader Jeff Dean is leaving to build a new company called Discovery Loop.",
					"Koray Kavukcuoglu, Google DeepMind's chief technology officer, is taking over operational leadership as senior vice president and will report to Sundar Pichai.",
				}},
				{Heading: "One lab, two departures", Paragraphs: []string{
					"Dean's departure is not a routine executive move. The 27-year Google veteran is taking Sanjay Ghemawat, Oriol Vinyals, and Quoc Le into Discovery Loop, a public-benefit company focused on using machine learning to automate scientific discovery. Google will invest in the startup and provide cloud support, preserving a relationship while losing a core internal research center of gravity.",
					"Google says the Hassabis and Dean moves are not connected. The timing nevertheless makes them read as one organizational event: the scientist who led DeepMind is moving upward into long-horizon strategy while the engineer who shaped Google's AI infrastructure is moving outward.",
				}},
				{Heading: "The Gemini clock is still running", Paragraphs: []string{
					"The changes arrive alongside reports that Gemini 3.5 Pro is months behind schedule, departures of senior researchers, and employee concerns about morale and direction. Google still has unusual advantages — custom chips, cloud distribution, and profitable products that can fund enormous AI spending — but those assets do not remove the need to ship competitive models on time.",
					"Kavukcuoglu now has to turn organizational consolidation into a clearer release path. Hassabis can focus on longer-term scientific work, but the market will judge the reshuffle first by whether Gemini execution accelerates.",
				}},
				{Heading: "Sources", Paragraphs: []string{
					"Axios, Google DeepMind CEO Demis Hassabis is stepping aside: https://www.axios.com/2026/08/05/google-deepmind-demis-hassabis-ai",
					"Axios, Google's AI leadership shuffle: https://www.axios.com/2026/08/06/googles-ai-leadership-shuffle",
					"ITPro, DeepMind CEO Demis Hassabis steps aside: https://www.itpro.com/business/leadership/deepmind-ceo-demis-hassabis-steps-aside-amid-google-leadership-shake-up",
				}},
			},
		},
		{
			Title:   "OpenAI Just Gave Its Most Dangerous Model to Defenders First",
			Slug:    "openai-daybreak-gpt-5-6-cyber-launch-2026",
			Date:    "August 11, 2026",
			Tag:     "Security",
			Summary: "OpenAI's Daybreak program separates defensive access to GPT-5.6 Sol from a tightly vetted GPT-5.6-Cyber tier built for exploit validation, making procedural controls the price of offensive capability.",
			Sections: []Section{
				{Paragraphs: []string{
					"OpenAI's Daybreak program is an unusually direct answer to a difficult security question: what should a lab do when its model is capable of finding and validating real exploits? The company is giving defenders access first, but separating ordinary defensive work from a much more capable offensive-security tier.",
					"Daybreak Blue runs on GPT-5.6 Sol with guardrails tuned for vulnerability detection, malware analysis, incident response, and secure code review. Daybreak Red unlocks GPT-5.6-Cyber for vetted exploit validation, penetration testing, and vulnerability discovery.",
				}},
				{Heading: "A capability jump with a gate around it", Paragraphs: []string{
					"OpenAI reports that GPT-5.6 Sol completed 1.5% of sensitive exploit-chain and privilege-escalation requests on its Advanced Cybersecurity Completion benchmark, compared with 95% for GPT-5.6-Cyber. The company also says the model found previously unknown vulnerabilities in Chrome's V8 engine and a mobile operating system during testing.",
					"Those are company-reported results, not an independent certification. They still explain why OpenAI created separate access tiers: a model useful enough to discover serious bugs is also useful to attackers.",
				}},
				{Heading: "The guardrail is procedural", Paragraphs: []string{
					"Daybreak Red requires identity verification, legal declarations, activity monitoring, and isolated environments. Hardware security keys become mandatory for accounts on September 1, according to OpenAI's announcement. These controls make misuse harder to hide and raise the cost of casual abuse, but they do not make a powerful model technically incapable of misuse.",
					"That is the central trade: OpenAI is trusting a vetted, traceable defender population to handle capabilities that the base model is still trained to refuse. The strategy may give defenders a head start, but it also makes account security and oversight part of the model's safety case.",
				}},
				{Heading: "Sources", Paragraphs: []string{
					"OpenAI, Expanding Daybreak as the Cyber Defense Window Narrows: https://openai.com/index/expanding-daybreak-as-the-cyber-defense-window-narrows/",
					"Cybersecurity News, OpenAI expands Daybreak Cyber: https://cybersecuritynews.com/openai-expands-daybreak-cyber/",
					"The Decoder, OpenAI launches GPT-5.6-Cyber: https://the-decoder.com/openai-launches-gpt-5-6-cyber-to-help-defenders-find-vulnerabilities-before-attackers-do/",
				}},
			},
		},
	}, posts...)
}
