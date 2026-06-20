package content

func init() {
	posts = append([]Post{
		{
			Title:   "AI Designed the Molecule in Months — The Clinic Still Takes Years",
			Slug:    "ai-drug-discovery-clinic-not-approval-2026",
			Date:    "June 20, 2026",
			Tag:     "Science",
			Summary: "AI can now design drug candidates in months, but as of mid-2026 zero AI-discovered drugs have FDA approval; the real test is still the clinic.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"This week, the biotech world gathered in Boston for the BIO 2026 International Convention, and one theme ran through nearly every AI-focused booth and panel: the molecules are no longer the hard part. Companies like Insilico Medicine arrived showcasing platforms that can take a disease, find a plausible biological target, and design a drug-like molecule against it in a fraction of the time the industry once needed. The pitch decks are dazzling. The timelines are real. And yet, as of mid-2026, the number of AI-discovered drugs approved by the U.S. Food and Drug Administration sits at exactly zero.",
						"That gap - between a field that has gotten breathtakingly fast at the beginning of drug discovery and one that still has nothing across the finish line - is the most honest way to understand where AI in medicine actually stands right now.",
						"Start with the speed, because it is genuinely impressive. The traditional path from \"we think this protein matters\" to an identified drug target used to take roughly 18 months of painstaking lab work; AI-driven platforms now routinely compress that to around 9 months. Designing a candidate molecule once meant five years or more of iterative chemistry. Generative models - the same broad family of technology behind the chatbots you use, retrained on the language of proteins and small molecules - have pulled that down to under a year in the best cases. Protein-structure predictors like AlphaFold and RoseTTAFold turned a problem that won a Nobel Prize into a routine first step. Target identification, molecular design, and biologics structure prediction have all been quietly transformed from artisanal crafts into computational pipelines.",
					},
				},
				{
					Heading: "The Proof Point",
					Paragraphs: []string{
						"The proof that this can actually produce a real drug already exists. Insilico Medicine's ISM001-055, also known as rentosertib, is the clearest example. The company used AI to both nominate the biological target and design the molecule, then ran it through a Phase 2a trial in idiopathic pulmonary fibrosis - a brutal, progressive lung-scarring disease with few good options. In a study of 71 patients across 21 sites in China, the highest-dose group didn't just slow the decline in lung function; on the forced vital capacity measure, patients on placebo lost an average of 62 mL of lung capacity while those on the 60 mg dose gained about 98 mL.",
						"The whole molecule went from target to Phase 2 in roughly five years, against an industry norm of well over a decade. It remains the first time an AI-designed drug aimed at an AI-discovered target has shown a clinical efficacy signal in humans.",
					},
				},
				{
					Heading: "Why the Approval Count Is Still Zero",
					Paragraphs: []string{
						"So why the zero on the approvals board?",
						"Because the part of drug development that AI has revolutionized - discovery and design - was never the part that kills most drugs. The graveyard of pharmaceutical research is the clinic. Roughly nine out of ten drugs that enter human trials fail, most of them in Phase 2 and Phase 3, where a molecule that looked perfect in a model and promising in a dish turns out to be toxic, ineffective, or simply no better than what already exists. AI can help you arrive at that test faster and cheaper. It cannot make a human body respond.",
						"A 12-week signal in 71 patients, however encouraging, is not the same as a multi-year, multi-thousand-patient Phase 3 readout - and the history of medicine is littered with early signals that evaporated at scale.",
					},
				},
				{
					Heading: "The Recalibration",
					Paragraphs: []string{
						"This is the quiet recalibration happening across the most serious players in the space. Isomorphic Labs, the DeepMind spinout built around AlphaFold, spent 2025 telling the world its first human trials were imminent. By January 2026, CEO Demis Hassabis had walked that back to \"by the end of 2026\" - not a failure, but a useful dose of realism from a company that has raised $600 million and signed partnerships worth roughly $3 billion with Novartis, Eli Lilly, and others.",
						"Even the best-funded, most technically credible AI-drug company on earth is still working to get its first molecule into a single human. Meanwhile, the merger of Recursion and Exscientia into a single end-to-end platform - pairing automated lab screening with AI chemistry - reflects a dawning consensus that no single clever model wins this game. Owning the whole pipeline, from biology to bedside, might.",
					},
				},
				{
					Heading: "The Useful Frame",
					Paragraphs: []string{
						"For readers trying to calibrate the hype, here is the useful frame. 2026 is not the year of the AI-drug approval. It is the year of the AI-drug readout - the season when the first wave of candidates designed or discovered with machine learning starts reporting real clinical data. Some of those readouts will be good. Many will disappoint, exactly as they would for conventionally discovered drugs, because biology is harder than chemistry and humans are harder than models.",
						"The companies that endure will be the ones that treated AI as a way to take more, better-designed shots on goal - not as a magic wand that repeals the base rates of clinical failure.",
						"That is the thing worth sitting with. The most important number in AI drug discovery over the next two years won't be how fast a molecule was designed. It will be how many of these AI-born candidates survive contact with an actual patient. The algorithms got us to the starting line in record time. The race itself hasn't changed.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Insilico Medicine, \"A Phase 2 Readout Generates Excitement for the Potential of AI-Driven Drug Discovery\": https://insilico.com/blog/1112",
						"Insilico Medicine, \"Insilico Announces Nature Medicine Publication of Phase IIa Results of Rentosertib\": https://insilico.com/news/tnrecuxsc1-insilico-announces-nature-medicine-publi",
						"Nature Medicine via PubMed, \"A generative AI-discovered TNIK inhibitor for idiopathic pulmonary fibrosis\": https://pubmed.ncbi.nlm.nih.gov/40461817/",
						"IntuitionLabs, \"Isomorphic Labs & AlphaFold: AI Drug Discovery in Trials\": https://intuitionlabs.ai/articles/isomorphic-labs-alphafold-ai-drug-discovery-trials",
						"News-Medical, \"Insilico Medicine to showcase AI-driven drug breakthroughs at BIO 2026 International Convention,\" June 18, 2026: https://www.news-medical.net/news/20260618/Insilico-Medicine-to-showcase-AI-driven-drug-breakthroughs-at-BIO-2026-International-Convention.aspx",
						"FDA, \"Artificial Intelligence for Drug Development\": https://www.fda.gov/about-fda/center-drug-evaluation-and-research-cder/artificial-intelligence-drug-development",
						"Author article handoff: https://docs.google.com/document/d/1408TJIM5PqcbrksyWilkOd4X9H6k6Tn_26UIiPS6uDk/edit",
					},
				},
			},
		},
	}, posts...)
}
