package content

func init() {
	posts = append([]Post{
		{
			Title:   "AI Agents Learn to Click: Gemini Takes Over the Mac as OpenAI Moves Into Hospital Paperwork",
			Slug:    "ai-agents-learn-to-click-gemini-mac-openai-healthcare",
			Date:    "November 21, 2026",
			Tag:     "Agents",
			Summary: "In one week, Google gave Gemini control of Mac desktop apps and OpenAI put computer-use agents to work on healthcare referrals and insurance authorizations. The chatbot is becoming an operator.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"For years, the assistant in your browser tab could only talk. This week it learned to click.",
						"On November 16, OpenAI announced computer-use agents for healthcare. Three days later, on November 19, Google added \"computer use\" to the Gemini app for Mac. Two very different audiences, one shared idea: if software has no API, let the AI use the screen like a person would.",
					},
				},
				{
					Heading: "Gemini gets a Mac of its own",
					Paragraphs: []string{
						"Google's Gemini desktop app for macOS can now navigate native apps, browse websites in the background, and finish multi-step tasks such as working through menus and filling out forms (TechCrunch, 9to5Google, Nov 19). You summon it by adding @Computer use to a prompt.",
						"The design choices are the interesting part:",
						"- It works in an isolated background window with a live preview in the chat, so your own screen and cursor stay free.",
						"- You can pause, stop, or approve actions, and you choose which apps it may touch.",
						"- New \"Desktop\" connectors, including Blender and Autodesk Fusion, join the existing Git and Obsidian integrations.",
						"- It is rolling out in US English to Google AI Plus, Pro and Ultra subscribers, requires an Apple Silicon Mac, and is not available in the EU or UK.",
						"TechCrunch notes Google is a little late: OpenAI, Anthropic and Meta already offer agentic features in their desktop apps. The underlying capability reached developers in June with Gemini 3.5 Flash.",
					},
				},
				{
					Heading: "OpenAI heads for the clinic back office",
					Paragraphs: []string{
						"OpenAI's announcement is less flashy and arguably more consequential. Its computer-use agents are now available for healthcare through the API, ChatGPT for Healthcare and ChatGPT for Clinicians. Rather than integrating with each hospital system, the agents operate existing clinic software and insurer portals by clicking and typing, the way a staff member would (The Decoder, Interesting Engineering, Nov 16-17).",
						"In OpenAI's demo, an agent drafted an MRI order, a medical record entry and a reply to the patient for a doctor to review, then opened an insurer's portal and prepared the prior-authorization application. Clinicians and authorization specialists review the output before anything is sent.",
						"Reported early users include:",
						"- CVS Health, using agents for referrals and document handling; one report says it cleared a months-long document backlog in a single weekend, with work reviewed through screenshots and action logs.",
						"- Cigna Healthcare, helping case managers pull coverage, provider and appointment details from more than 10 separate systems.",
						"- Ambience demonstrated an agent preparing an authorization request with clinician approval, and Abridge is developing agents for inpatient care coordination.",
						"On the controls side, ChatGPT for Healthcare offers role-based access and audit logs, and OpenAI says organizational data is not used to train its base models. OpenAI reported a cost of $0.085 per task for GPT-6 Luna on Stanford researchers’ HealthAdminBench at maximum reasoning effort. In a separate OpenAI evaluation of GPT-6 Astra on electronic health record inbox work, reviewers rated 99% of 325 reviewed tasks safe based on whether an error, omission or action could plausibly harm a patient. The result applies to that evaluation.",
					},
				},
				{
					Heading: "Why the screen is the new API",
					Paragraphs: []string{
						"Healthcare is the clearest example of why this matters. Its workflows are stitched together from portals, faxes and legacy systems that were never designed to talk to each other. An agent that can read a screen sidesteps years of integration work.",
						"But the same property is the risk. An agent can act within the permissions it is given, and a wrong click in an insurance portal can affect a patient’s care. Google lets users choose apps, watch the agent work and pause it. OpenAI emphasizes restricted access, review and audit logs. Those guardrails are part of what teams need to evaluate.",
					},
				},
				{
					Heading: "What to watch",
					Paragraphs: []string{
						"- Whether review stays meaningful at volume. A specialist approving hundreds of agent-drafted authorizations a day may start rubber-stamping.",
						"- Whether Windows and enterprise desktop versions follow Gemini's Mac debut.",
						"- Whether regulators treat UI-driven agents in medical workflows differently from ordinary software.",
						"The question for readers: when the software you use every day can be driven by an AI, who is actually responsible for what it clicks?",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"TechCrunch, \"Google gives Gemini the ability to use your Mac for you\", Nov 19, 2026: https://techcrunch.com/2026/11/19/google-gives-gemini-the-ability-to-use-your-mac-for-you/",
						"9to5Google, \"Gemini app for Mac adds computer use to control desktop apps\", Nov 19, 2026: https://9to5google.com/2026/11/19/gemini-mac-computer-use/",
						"The Decoder, \"AI agents are starting to fill out medical forms and file insurance requests in place of human staff\": https://the-decoder.com/ai-agents-are-starting-to-fill-out-medical-forms-and-file-insurance-requests-in-place-of-human-staff/",
						"Interesting Engineering, \"OpenAI agents enter US healthcare after 99% clinical safety rating\": https://interestingengineering.com/ai-robotics/openai-healthcare-computer-use-agents",
						"Financial News, \"OpenAI Unveils Hospital Administration Agent\", Nov 17, 2026: https://en.fnnews.com/news/202611170756368840",
						"Google: Let Gemini tackle repetitive chores on your Mac, Nov 19, 2026: https://blog.google/innovation-and-ai/products/gemini-app/computer-use-mac/",
						"OpenAI: Advancing computer-use agents for healthcare teams, Nov 16, 2026: https://openai.com/index/advancing-computer-use-for-healthcare-teams/",
					},
				},
			},
		},
	}, posts...)
}
