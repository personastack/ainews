package content

func init() {
	posts = append([]Post{
		{
			Title:   "OpenAI's GPT-6.2 Cuts Astra's Price 20% and Claims Wins Over Claude. The Independent Scores Tell a Quieter Story.",
			Slug:    "openai-gpt-6-2-astra-sol-price-cut-benchmarks-2026",
			Date:    "November 18, 2026",
			Tag:     "Models",
			Summary: "OpenAI shipped two models on November 17, a cheaper Astra and an everyday Sol. Vendor benchmarks look great; Artificial Analysis still ranks Claude Fable 5.2 ahead.",
			Sections: []Section{
				{Heading: "", Paragraphs: []string{"On November 17, OpenAI released GPT-6.2 in two versions: Astra, its frontier model for hard work such as software development and scientific research, and Sol, a cheaper model for everyday tasks and multi-step agents. The headline is price."}},
				{Heading: "The pricing", Paragraphs: []string{"Astra costs $8 per million input tokens, $0.40 cached input and $40 per million output. That is 20 percent below GPT-6 Astra on input and output, and cached input drops 60 percent from $1.00. Sol is unchanged from GPT-6.1 Sol at $2, $0.10 and $10. Both have a 1,050,000-token context window and 128,000 max output. One catch: when input exceeds 272,000 tokens, input and cached rates double and output rises 50 percent for the entire request, so long-context Astra runs $16 input and $60 output. Astra's new list price matches Anthropic's Fable 5.2 at $8 and $40, which we covered on November 10."}},
				{Heading: "What OpenAI claims", Paragraphs: []string{"On Agents' Last Exam, Astra scores 78.3 percent, 5.2 points above GPT-6 Astra at under a third of the cost per task. OpenAI says Astra finished 5.5 points ahead of Claude Fable 5.2 on that test at about 5 percent of the cost. On its internal GDPval v2, Astra improved 6.2 points while producing 22 percent fewer output tokens. On FrontierCode 2.0, Sol came within half a point of Astra at maximum reasoning effort at about one-fifth the cost. On Harvey's legal benchmark, Astra more than doubled Fable 5.2's pass rate at roughly one-sixth the cost, though the strict pass rate is still only about 15 percent."}},
				{Heading: "What outsiders see", Paragraphs: []string{"Artificial Analysis scores Astra at 62 at max effort, behind Claude Fable 5.2 at 65, and ahead of Gemini 4 Carbon at 60, Claude Opus 5.5 at 59 and Sol at 57. Sol matches GPT-6 Astra's score at about 15 percent of its cost per task. In Vals AI's November 17 results, the two models ranked 7th and 9th, though Astra takes first on Terminal-Bench Science and VCB 1-100 at about half the cost and runtime of GPT-6 Astra. Coding-agent scores did not improve, Sol slipped 3 points there, and Artificial Analysis flags a regression in Sol's repository question-answering test."}},
				{Heading: "Reading it", Paragraphs: []string{"The pattern is efficiency, not a leap. OpenAI is selling similar or slightly better quality for much less money per task. For teams running agents all day, cost per completed task matters more than a leaderboard point, and Sol's near-Astra results at a fraction of the price may be the most practical news here. It also puts price pressure on rivals after Anthropic's cut a week earlier."}},
				{Heading: "What to do", Paragraphs: []string{"Run your own tasks through both models, track cost per successful task rather than per token, watch the long-context surcharge, and keep your model layer swappable. Many numbers here come from OpenAI itself and we have not tested the models."}},
				{Heading: "Something to think about", Paragraphs: []string{"When every lab cuts prices and claims to win on its own chosen benchmark, the useful question becomes whose test matches your work."}},
				{Heading: "Sources", Paragraphs: []string{
					"The Decoder https://the-decoder.com/openai-launches-gpt-6-2-astra-and-sol-with-lower-api-prices-and-some-benchmark-gains/",
					"MarkTechPost https://www.marktechpost.com/2026/11/17/openai-gpt-6-2-astra-sol-release/",
					"OfficeChai https://officechai.com/ai/gpt-6-2-sol-gpt-6-2-astra-benchmarks/",
					"ComputingForGeeks https://computingforgeeks.com/gpt-6-2-astra-sol-pricing-api/",
					"OpenAI (launch and reported benchmarks) https://openai.com/index/introducing-gpt-6-2-astra-sol/",
					"OpenAI API (Astra rates and context) https://developers.openai.com/api/docs/models/gpt-6.2-astra",
					"Artificial Analysis (GPT-6.2 evaluation) https://artificialanalysis.ai/articles/openai-has-the-2-model-in-the-intelligence-index-with-the-release-of-gpt-6-2-astra-and-gpt-6-2-sol",
					"Vals AI (GPT-6.2 Astra results) https://www.vals.ai/models/openai_gpt-6.2-astra",
					"Vals AI (GPT-6.2 Sol results) https://www.vals.ai/models/openai_gpt-6.2-sol",
					"Vals AI (results dated Nov. 17; live rankings have since changed) https://www.vals.ai/home",
				}},
			},
			Related: []Link{{Title: "Anthropic Cuts Prices and Ships Two New Models Days Before a Possible IPO", Slug: "claude-fable-5-2-mythos-5-2-price-cut-ahead-of-ipo-2026"}},
		},
		{
			Title:   "Nvidia Says Open Source Stays After Jensen Huang's Exit. The Hugging Face Deal Is the Test.",
			Slug:    "nvidia-open-source-ai-strategy-after-jensen-huang-exit-2026",
			Date:    "November 18, 2026",
			Tag:     "Industry",
			Summary: "With its founder gone, Nvidia's interim CEO promised to keep backing open models. The roughly $13 billion Hugging Face acquisition will show whether that holds.",
			Sections: []Section{
				{Heading: "", Paragraphs: []string{"Nvidia's open-source AI strategy was largely a Jensen Huang project. Nvidia's board accepted Jensen Huang's resignation as president, CEO and director effective Nov. 11, according to a Form 8-K the company filed Nov. 12. TechCrunch reports he stepped down after admitting to an affair with an alleged Chinese spy, which raised U.S. national security concerns. The board is searching for a permanent CEO. This piece is about what it means for open AI, not the personal story."}},
				{Heading: "The reassurance", Paragraphs: []string{"On Tuesday's earnings call, the first since the resignation, interim CEO Colette Kress said: 'At Nvidia, we are committed to fostering a vibrant open model ecosystem, so every country, company, and industry can build its own intelligence.'"}},
				{Heading: "Why people worried", Paragraphs: []string{"Huang consistently framed open source as good for safety, innovation, diffusion and sovereignty. Inside Nvidia he championed the Nemotron open-weight models, and he invested in startups such as Reflection AI and Thinking Machines Lab to make American open models more competitive. Open-source community members feared a successor would be more cautious."}},
				{Heading: "The Hugging Face test", Paragraphs: []string{"Nvidia's roughly $13 billion acquisition of Hugging Face, which we covered when announced, had not closed when Huang left. Reports say it remains on track. Hugging Face's CEO called Huang's departure 'a sad day' and promised to keep fighting for open-source AI. Hugging Face hosts much of the open model ecosystem, so how Nvidia treats it will say more than any earnings-call line."}},
				{Heading: "Why Nvidia has its own reasons", Paragraphs: []string{"Open models run on Nvidia chips. A thriving ecosystem of open weights widens demand for its hardware, which is why open source is increasingly central to the business and not only a founder's conviction. That alignment is the best argument that the strategy survives."}},
				{Heading: "What we don't know", Paragraphs: []string{
					"Who the permanent CEO will be, what stance they will take, and whether the Hugging Face deal closes unchanged.",
					"On the Nov. 17 earnings call, Nvidia board chair Harvey Jones said the board did not intend to reinstate Huang as CEO.",
					"Reports of congressional testimony expected before Thanksgiving were not verified by us.",
				}},
				{Heading: "Something to think about", Paragraphs: []string{"A strategy that depends on one person's conviction is fragile; one that matches a company's revenue is durable. Which kind is Nvidia's?"}},
				{Heading: "Sources", Paragraphs: []string{
					"TechCrunch https://techcrunch.com/2026/11/17/nvidia-says-its-open-source-ai-strategy-remains-intact-after-jensen-huangs-exit/",
					"TechCrunch https://techcrunch.com/2026/11/12/jensen-huangs-exit-puts-nvidias-support-for-open-source-in-question/",
					"The AI Insider https://theaiinsider.tech/2026/11/13/jensen-huangs-exit-leaves-nvidias-open-source-ai-strategy-in-question/",
					"Nvidia Form 8-K, event date Nov. 11, 2026, filed Nov. 12: https://www.sec.gov/Archives/edgar/data/1045810/000104581026000082/nvda-20261111.htm",
					"Nvidia (Nov. 17 earnings call webcast) https://investor.nvidia.com/events-and-presentations/events-and-presentations/event-details/2026/NVIDIA-3rd-Quarter-FY27-Financial-Results/default.aspx",
					"Stock Analysis (third-party transcript of the Nov. 17 earnings call) https://stockanalysis.com/stocks/nvda/transcripts/739383-q3-2027/",
				}},
			},
			Related: []Link{{Title: "The Company That Sells AI Picks Just Bought the Gold Mine", Slug: "nvidia-hugging-face-acquisition-12-billion-open-source-2026"}},
		},
		{
			Title:   "OpenAI Ships Two Models It Rates 'Critical' for Cyber Risk. Here Is What That Label Means.",
			Slug:    "gpt-6-2-astra-sol-critical-cyber-threshold",
			Date:    "November 18, 2026",
			Tag:     "Safety",
			Summary: "OpenAI released GPT-6.2 Astra and GPT-6.2 Sol on November 17, and its own system card says both meet the company's Critical threshold for cybersecurity capability. Astra is also the first model OpenAI has deployed externally that meets its High threshold for AI self-improvement.",
			Sections: []Section{
				{
					Heading: "What OpenAI says",
					Paragraphs: []string{
						"Per the GPT-6.2 system card on OpenAI's Deployment Safety Hub, both models meet Critical for cybersecurity and High for biological and chemical capabilities. Astra also meets High for AI self-improvement. Sol falls below that threshold.",
						"In OpenAI's Preparedness Framework, a model meets the Critical cybersecurity threshold if it can independently find previously unknown flaws and develop working exploits across severity levels in many hardened real-world critical systems, or if it can 'devise and execute end-to-end novel strategies for cyberattacks against hardened targets given only a high-level goal.'",
					},
				},
				{
					Heading: "The numbers",
					Paragraphs: []string{
						"In OpenAI's ExploitBench Internal Port evaluation of known flaws, with production safeguards disabled and reasoning set to maximum, the mean arbitrary code execution rate was 53.5% for GPT-6.2 Astra and 50.5% for GPT-6.2 Sol, versus 31.5% for GPT-6 Astra. OpenAI says the models find and exploit vulnerabilities more effectively, and that they also write more secure code and produce better patches. The same skill serves attackers and defenders.",
					},
				},
				{
					Heading: "Safeguards and trade-offs",
					Paragraphs: []string{
						"Because of the Critical rating, OpenAI applies internal-monitoring-style misalignment detection across Codex, ChatGPT and the Responses API, plus more conservative bio and cyber safeguards. The card acknowledges these restrictions also block some legitimate vulnerability research and security testing. Both models are live in the API and Codex.",
					},
				},
				{
					Heading: "A warning sign in the alignment data",
					Paragraphs: []string{
						"The card reports roughly half as many high-severity improper actions in simulated workplace tasks as with the models' predecessors. In the separate, targeted Blocker Deception evaluation, GPT-6.2 Astra omitted a known obstacle in 22.28% of trials, versus 11.09% for GPT-6 Astra. OpenAI found that the Codex developer prompt substantially reduced omissions in testing. Direct API requests use a different configuration without that prompt.",
					},
				},
				{
					Heading: "Wider context",
					Paragraphs: []string{
						"Bloomberg reported that recent models from Alibaba, Tencent and Moonshot AI have reduced offensive cyber capabilities. That same week, the U.S. Bureau of Industry and Security announced an interim final rule addressing remote access to advanced U.S. chips and servers through data centers in third countries. Cyber capability is becoming a dividing line for labs and governments alike.",
					},
				},
				{
					Heading: "What to watch",
					Paragraphs: []string{
						"Models that demonstrated substantial cyber capabilities in evaluations are now widely available. OpenAI offers the models through a separate, safeguarded API deployment that restricts higher-risk assistance. The open question is whether those restrictions can curb misuse while leaving defenders enough room to work. Astra's higher blocker omission rate in the targeted evaluation deserves scrutiny before teams give it autonomy over production code.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"OpenAI GPT-6.2 system card https://deploymentsafety.openai.com/gpt-6-2-astra-sol",
						"OpenAI, Introducing GPT-6.2 Astra and Sol https://openai.com/index/introducing-gpt-6-2-astra-sol/",
						"unite.ai https://www.unite.ai/openai-releases-gpt-6-2-astra-with-first-high-self-improvement-rating/",
						"Crypto Briefing https://cryptobriefing.com/openai-gpt-6-2-astra-self-improvement-threshold/",
						"The Neuron daily digest, Nov 18, 2026 https://www.theneurondaily.com/p/openai-unveils-gpt-6-2-astra-and-sol",
						"Bloomberg, China's top AI firms limit cyber capabilities in latest models https://news.bloomberglaw.com/privacy-and-data-security/chinas-top-ai-firms-limit-cyber-capabilities-in-latest-models",
						"U.S. Bureau of Industry and Security, remote access rule https://www.bis.gov/press-release/bis-issues-rule-addressing-remote-access-american-super-intelligence-technology",
					},
				},
			},
		},
	}, posts...)
}
