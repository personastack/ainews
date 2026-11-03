package content

func init() {
	posts = append([]Post{{
		Title:   "Texas Wants a Judge to Approve OpenAI's Next Model. That Is the Real Story.",
		Slug:    "texas-attorney-general-sues-openai-chatgpt-safety-minors-model-approval-2026",
		Date:    "October 25, 2026",
		Tag:     "Policy",
		Summary: "Attorney General Ken Paxton's lawsuit makes familiar safety allegations, but the remedy he is asking for, court-appointed pre-approval of new models, would be a first.",
		Sections: []Section{
			{Heading: "", Paragraphs: []string{"On Friday, October 23, Texas Attorney General Ken Paxton sued OpenAI and CEO Sam Altman in Collin County District Court. The complaint alleges the company misled consumers about ChatGPT's safety, accuracy and reliability, especially for teenagers. Paxton's office says ChatGPT has aided mass shooters, encouraged vulnerable people toward suicide, and hooked minors on a product that collected their data without parental oversight. Press coverage cites eight instances in which users later committed violent acts."}},
			{Heading: "The claims:", Paragraphs: []string{"The state invokes the Texas Deceptive Trade Practices Act and the Securing Children Online through Parental Empowerment (SCOPE) Act, a 2023 law also known as House Bill 18. The AG's office says OpenAI 'pressed ahead with creating these risks while unlawfully deceiving Texans about the product's safety, accuracy and reliability.' The suit alleges products shipped without sufficient parental controls, age verification or safety testing. Per Hoodline's summary, it seeks civil penalties of $10,000 per violation and forfeiture of profits."}},
			{Heading: "The remedy is the headline:", Paragraphs: []string{"Most AI lawsuits ask for money. This one also asks for an injunction with two parts: block minors from using ChatGPT, and bar OpenAI from releasing new ChatGPT models in Texas without approval from a court-appointed independent evaluator. Pre-release licensing of software is something regulators have debated for years but largely not done. Even a narrower version would turn a model launch into something closer to a drug approval, with a third party holding a veto."}},
			{Heading: "Why it may be hard to win:", Paragraphs: []string{"A state court would have to decide that product-safety and consumer-protection law reaches generative model releases, and that an injunction limited to one state is workable for a product served nationwide. A company representative told the Texas Tribune OpenAI has 'worked with experts to mitigate risks and build additional safeguards into its products'; other outlets reported no immediate comment. Reports also note that Elon Musk, whose xAI competes with OpenAI, funds Paxton's Senate campaign, context OpenAI will likely raise."}},
			{Heading: "What we don't know:", Paragraphs: []string{"The allegations and proposed remedies are the state's, not court findings or an existing release-approval regime. The complaint requests the restrictions described above; outlets differ slightly on the filing date (we use Oct 23). The eight violent-act instances are as summarized by press and not independently verified."}},
			{Heading: "Why it matters:", Paragraphs: []string{"October has been heavy for AI litigation, with new suits against Grok and Meta in recent weeks. State attorneys general are filling the gap left by absent federal rules, and each tests a different lever. Texas is testing the biggest: whether a government can sit between a model and its release date. If any court entertains that, every lab's launch calendar changes. Question to watch: whether OpenAI seeks removal to federal court, and whether other states copy the pre-approval request."}},
			{Heading: "Sources", Sources: []Source{
				{Title: "Texas Attorney General — Original petition", URL: "https://www.texasattorneygeneral.gov/sites/default/files/images/press/OpenAI%20Final.pdf"},
				{Title: "Courthouse News", URL: "https://www.courthousenews.com/texas-attorney-general-sues-openai-for-lying-about-chatgpt-safety-and-safeguards/"},
				{Title: "KBTX", URL: "https://www.kbtx.com/2026/10/23/texas-attorney-general-ken-paxton-sues-chatgpt-maker-openai/"},
				{Title: "Hoodline", URL: "https://hoodline.com/2026/10/texas-ag-paxton-sues-openai-says-chatgpt-fueled-suicides-and-mass-shooters/"},
				{Title: "Houston Chronicle", URL: "https://www.houstonchronicle.com/news/houston-texas/texas/article/texas-openai-chatgpt-lawsuit-22491691.php"},
			}},
		},
	}}, posts...)
}
