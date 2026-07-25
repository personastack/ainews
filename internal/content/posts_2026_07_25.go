package content

func init() {
	posts = append([]Post{
		{
			Title:   "The White House Says China Cloned Claude to Build Kimi K3. There Wasn't Enough Time, Researchers Say.",
			Slug:    "white-house-moonshot-kimi-k3-anthropic-fable-distillation-2026",
			Date:    "July 25, 2026",
			Tag:     "Policy",
			Summary: "OSTP director Michael Kratsios accused Moonshot AI of distilling Anthropic's Fable model and routing banned Nvidia chips through Thailand to build Kimi K3. The chip-access claim is easier to test. The distillation claim has a timing problem.",
			Related: []Link{
				{
					Title: "The Chip Industry Just Had Its Best Quarter Ever. Wall Street Sold It Anyway.",
					Slug:  "chip-earnings-record-profits-stock-selloff-kimi-k3-2026",
				},
				{
					Title: "The AI Industry Graded Its Own Safety Homework. Nobody Passed.",
					Slug:  "ai-safety-index-summer-2026-anthropic-c-plus-pause-pledges-erode",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"Six days after Kimi K3 became the largest open-weight AI model ever released, the White House accused its maker of stealing the model that made it possible.",
						"On July 22, Michael Kratsios, director of the White House Office of Science and Technology Policy, accused Moonshot AI of distilling Anthropic's Fable model into Kimi K3 through a large-scale internal platform that switched between access methods to avoid detection. Kratsios also alleged Moonshot acquired GB300-equipped servers and accessed GB300s in Thailand to train AI models, even though Nvidia's Blackwell-generation GB300 chips are restricted from export to China.",
						"The charge landed at a delicate moment. Moonshot has reportedly been preparing a final pre-IPO funding round that could value the company at $50 billion ahead of a possible Hong Kong listing. Kimi K3 itself is a 2.8 trillion-parameter mixture-of-experts model with a one-million-token context window, and Moonshot has promoted it as the largest open-weight model released so far.",
					},
				},
				{
					Heading: "The timeline problem",
					Paragraphs: []string{
						"Fable 5 became publicly reachable on July 1. Anthropic had launched it on June 9, but U.S. export controls forced the company to suspend access three days later until the restrictions were lifted on June 30. Kimi K3 shipped on July 16. That gives a maximum public distillation window of fifteen days.",
						"Researchers interviewed by TechCrunch said that timeline does not fit the strongest version of Washington's claim. Braden Hancock, co-founder of Snorkel AI, said he did not think a model could get that strong that quickly through distillation alone, even though distillation is a real technique. He also noted that Moonshot employs legitimate researchers and engineers, including people trained at Carnegie Mellon.",
						"Nathan Lambert of the Allen Institute for AI raised the compute problem. Reinforcement-learning distillation at this scale requires enormous numbers of agent rollouts; doing that through a rival's paid API would be expensive and bottlenecked by access limits. Supervised fine-tuning on another model's answers helps, but near the frontier it usually has diminishing returns. If it were enough by itself, he argued, catching up would be easy for everyone.",
						"Kratsios did not publish technical evidence for the distillation allegation. Moonshot had not responded to requests for comment in the public reports the Author used for this article.",
					},
				},
				{
					Heading: "The chip trail is a separate story",
					Paragraphs: []string{
						"The Nvidia allegation is different. Thailand is not covered by the same export restrictions as China, and Southeast Asian data-center capacity has expanded as hyperscalers diversify where they run compute. A Chinese lab could plausibly rent or route training capacity across a border, though the public reporting has not independently verified that Moonshot did so.",
						"That makes the chip-access claim narrower and easier to investigate than the claim that Kimi K3 was cloned from Fable. Regulators do not need to settle the training-method debate to ask where restricted GPUs were physically located, who controlled the servers, and whether any provider knowingly helped a Chinese lab evade U.S. export rules.",
					},
				},
				{
					Heading: "A pattern, not an isolated incident",
					Paragraphs: []string{
						"This is not the first accusation of a Chinese lab harvesting U.S. model outputs. In a June 10 letter to the Senate Banking Committee, Anthropic accused Alibaba's Qwen team of running what it called the largest known distillation campaign against Claude: roughly 25,000 fraudulent accounts and more than 28.8 million exchanges from April 22 to June 5. Alibaba has denied wrongdoing, and the public record still depends heavily on the accusing company's telemetry and analysis.",
						"Congress has been circling the same issue since April, with House committees investigating whether Chinese AI labs have integrated U.S. model outputs into their own systems. The Moonshot claim fits the broader political pattern: Chinese open-weight models are closing capability gaps faster than Washington expected, and one increasingly common explanation is that they copied U.S. frontier systems. The open question is whether the evidence in this case supports that explanation.",
					},
				},
				{
					Heading: "What to watch",
					Paragraphs: []string{
						"Two threads matter now. The first is whether the White House or Anthropic publishes evidence specific enough to convince outside researchers that distillation explains Kimi K3's capability jump. The second is whether the Thailand chip-access allegation leads to enforcement against compute providers or intermediaries. The second case may be easier for Washington: it depends less on reconstructing Moonshot's training recipe and more on following the servers.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Michael Kratsios statement on X: https://x.com/mkratsios47/status/2079933645888880708",
						"CyberScoop, White House accuses Moonshot AI of Anthropic model distillation: https://cyberscoop.com/white-house-accuses-moonshot-ai-anthropic-model-distillation/",
						"TechCrunch, Experts say exploiting Anthropic's Fable isn't how Kimi K3 got so good: https://techcrunch.com/2026/07/23/experts-say-exploiting-anthropics-fable-isnt-how-kimi-k3-got-so-good/",
						"Bloomberg, China's Moonshot in Talks on Pre-IPO Funds at $50 Billion Value: https://www.bloomberg.com/news/articles/2026-07-21/china-s-moonshot-in-talks-on-pre-ipo-funds-at-50-billion-value",
						"TechTimes, Alibaba Ran Largest Known AI Theft Campaign Against Claude, Anthropic Tells Senate: https://www.techtimes.com/articles/319105/20260625/alibaba-ran-largest-known-ai-theft-campaign-against-claude-anthropic-tells-senate.htm",
						"Gizmochina, Kimi K3: Moonshot AI unleashes 2.8 trillion parameter model for free: https://www.gizmochina.com/2026/07/19/kimi-k3-moonshot-ai-unleashes-2-8-trillion-parameter-model-for-free/",
						"Tom's Hardware, China's 2.8-trillion-parameter Kimi K3 beats Claude Fable 5 in Frontend Code Arena benchmark: https://www.tomshardware.com/tech-industry/artificial-intelligence/moonshot-releases-2-8-trillion-parameter-kimi-k3",
						"Anthropic Newsroom, Redeploying Claude Fable 5: https://www.anthropic.com/news/redeploying-fable-5",
					},
				},
			},
		},
	}, posts...)
}
