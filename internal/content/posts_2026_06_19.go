package content

func init() {
	posts = append([]Post{
		{
			Title:   "Microsoft Put a Meter on Its AI. Then It Went Shopping for a Cheaper Engine.",
			Slug:    "ai-cost-meter-copilot-cowork-deepseek-2026",
			Date:    "June 19, 2026",
			Tag:     "Enterprise",
			Summary: "Microsoft made Copilot Cowork usage-based on June 16, putting a per-task meter on agentic AI, then reports emerged it is weighing a cheaper DeepSeek V4 engine.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"For three years, enterprise AI was sold like a gym membership: one flat fee per user, all the AI you can use. On June 16, 2026, Microsoft quietly retired that model for its most ambitious product yet - and in doing so, it made the real cost of agentic AI visible for the first time.",
						"Then, almost in the same breath, reports surfaced that Microsoft is shopping for a cheaper engine to put under the hood. Read together, those two moves tell you more about where AI is heading in 2026 than any model launch.",
					},
				},
				{
					Heading: "What Actually Changed",
					Paragraphs: []string{
						"Microsoft made Copilot Cowork generally available on June 16. Cowork is the agentic tier of Copilot - the version meant to take on multi-step, multi-tool jobs on your behalf rather than just autocomplete a sentence. The headline is not the feature set. It is the bill.",
						"Cowork is billed on usage, not headcount. You still need the underlying Microsoft 365 Copilot license as a prerequisite, but Cowork itself runs on a metered system Microsoft calls Copilot Credits. Each credit costs one cent, billed in arrears at the end of the month and pooled across the whole tenant.",
						"The price of any given task is computed from four ingredients: which model it used, how much context it had to retrieve, how many tool calls it made, and how long it ran.",
						"That last detail is the interesting one. Microsoft is no longer charging for access to AI. It is charging for work performed - and it is telling you exactly which levers move the number.",
					},
				},
				{
					Heading: "The Meter Has Real Numbers on It",
					Paragraphs: []string{
						"Microsoft's own June 2026 planning guidance gives a sense of scale. A light task - a quick summary, a simple lookup - runs around 125 credits. A medium task that pulls from several sources and produces real output lands in the 400 to 600 range. A heavy task, the kind of long, agentic, many-step job Cowork was built to show off, can run past 1,500 credits, and Microsoft notes some go well beyond that.",
						"Do the arithmetic at a penny a credit: a heavy task is fifteen dollars or more, every time it runs. A team firing off dozens of those a day is suddenly looking at a number that scales with ambition. Microsoft even built in a cushion - some early-access tenants from its Frontier program get a grace period before billing starts in July - which is itself a tell about how the costs land once the meter is live.",
						"This is the part the flat-rate era hid. When AI was a fixed monthly fee, nobody had to think about the cost of a single agent run. Usage-based pricing drags that cost into the daylight. The more capable the agent, the more legible - and the more uncomfortable - the invoice.",
					},
				},
				{
					Heading: "Enter the Cheaper Engine",
					Paragraphs: []string{
						"Here is where it gets strategically interesting. According to reporting from Axios on the same day, Microsoft is evaluating a fine-tuned version of DeepSeek's V4 model - or another open-source model - hosted on its own Azure infrastructure, as a lower-cost alternative to the OpenAI and Anthropic models that currently power Cowork.",
						"Microsoft has not confirmed a final choice, and it has said a lower-cost version is coming in the weeks ahead with the model to be named later. Treat the specifics as reported intent, not a done deal.",
						"But the direction is unmistakable, and it follows directly from the pricing change. Once you have told customers their bill is driven partly by which model it used, you have a powerful incentive to make the default model cheaper to run. Swapping a frontier model that costs a premium per token for a fine-tuned open-weight model you host yourself is one of the largest cost levers available - and it is invisible to the user, who still just sees Copilot.",
					},
				},
				{
					Heading: "The Quiet Story: the Model Is Becoming a Component",
					Paragraphs: []string{
						"Step back and the two announcements rhyme. Microsoft metered the work, then went looking to cheapen the most expensive ingredient in it. That is not the behavior of a company that thinks the frontier model is the product. It is the behavior of a company that thinks the model is a swappable part inside the product.",
						"This is the commoditization thesis arriving in the most concrete form yet. For two years the assumption was that whoever had the smartest model would win the enterprise. What Cowork's pricing reveals is a different contest: at production scale, with agents doing real multi-step work, the binding question is not which model is smartest, but which model is good enough at a price the meter can bear.",
						"The frontier model becomes the expensive option you reach for when the task demands it - not the default you run on everything.",
						"It also reframes the open-weight conversation. DeepSeek's V4 is not interesting to Microsoft here because it tops a leaderboard. It is interesting because a model you can fine-tune and host yourself collapses the per-task cost and gives you control of the most volatile line item in an agentic product. Open weights stop being a hobbyist story and become a margin story.",
					},
				},
				{
					Heading: "What to Watch",
					Paragraphs: []string{
						"Three things will tell you whether this is a blip or the new shape of enterprise AI. First: does Microsoft actually ship a non-frontier default for Cowork, and does it disclose which tasks route to which model?",
						"Second: do competitors follow Microsoft into transparent, task-level metering - or retreat to flat rates to hide the cost?",
						"Third: do enterprises start optimizing their own workflows the way they once optimized cloud spend, trimming heavy tasks and routing work to the cheapest model that clears the bar?",
						"The last one is the real shift. For the first time, the people deploying AI agents have a number to optimize against. When the bill becomes legible, behavior changes - and the AI you actually run starts being chosen by the accountant as much as the engineer.",
						"That is the lesson buried in a pricing page and an unconfirmed model swap: in 2026, the smartest model and the model you will actually use are quietly becoming two different things.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Microsoft 365 Blog, \"Copilot Cowork is now generally available,\" June 16, 2026: https://www.microsoft.com/en-us/microsoft-365/blog/2026/06/16/copilot-cowork-is-now-generally-available/",
						"Microsoft Copilot Credits Guide, June 16, 2026: https://cdn-dynmedia-1.microsoft.com/is/content/microsoftcorp/microsoft/bade/documents/products-and-services/en-us/ai/Microsoft-Copilot-Credits-Guide-June-16-2026-PUB.pdf",
						"Axios, \"Microsoft weighs DeepSeek for Copilot Cowork,\" June 16, 2026: https://www.axios.com/2026/06/16/microsoft-copilot-cowork-tokenmaxxing-cowork",
						"Author article handoff: https://docs.google.com/document/d/1GfHZCSb1N-3X99xApHVzURp5-96zHbYfiptXKUJ6BdM/edit",
					},
				},
			},
		},
	}, posts...)
}
