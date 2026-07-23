package content

func init() {
	posts = append([]Post{
		{
			Title:   "The Wire Became the Bottleneck — So AI Is Rebuilding It Out of Light",
			Slug:    "ai-silicon-photonics-interconnect-light-2026",
			Date:    "June 20, 2026",
			Tag:     "Hardware",
			Summary: "Copper is becoming the thing AI clusters trip over, so the next infrastructure race is moving light into the package: co-packaged optics and silicon photonics are turning into shipping interconnects, not just lab demos.",
			Related: []Link{
				{
					Title: "The Chip Stopped Being the Bottleneck — Now It's Power and Memory",
					Slug:  "ai-real-bottleneck-power-memory-not-chips-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"The fastest way to misunderstand the next AI hardware race is to stare only at the accelerator. GPUs still matter. Memory still matters. Power still matters. But inside the largest AI factories, another bottleneck is getting harder to ignore: the wire.",
						"When a training run spans thousands, then tens of thousands, then eventually millions of accelerators, the system is no longer just a pile of chips. It is a communications machine. Every model shard, gradient update, retrieval call, and inference handoff has to move. Copper traces, retimers, pluggable optics, and hot switch cages can only absorb so much bandwidth before they become the cost, power, heat, reliability, and deployment problem.",
						"That is why 2026 is starting to look like the year light moves from photonics conference slides into the AI interconnect stack. Not as science fiction. Not as \"computing with light\" replacing GPUs. As the less glamorous but more immediate shift: replacing electrical distance with optical distance wherever copper has become too lossy, hot, and fragile.",
					},
				},
				{
					Heading: "The Switch Moves Toward The Optics",
					Paragraphs: []string{
						"NVIDIA's Quantum-X Photonics is the clearest sign that co-packaged optics has crossed from research direction into product roadmap. The InfiniBand switch family is built around 115 Tb/s of switching capacity, with 144 ports running at 800 Gb/s each, and is expected in early 2026. Instead of keeping optical conversion out in pluggable modules at the edge of the box, the design moves silicon photonics close to the switch ASIC, shrinking the electrical path before the signal becomes light.",
						"NVIDIA's published pitch is not just more bandwidth. The company says the co-packaged approach uses four times fewer lasers, improves power efficiency by 3.5x, and increases resiliency by 10x compared with traditional optical interconnect approaches. Those are vendor claims, but the direction is easy to believe: fewer hot pluggables, shorter electrical runs, less conversion loss, and fewer active components that can fail in a cluster where a single fabric issue can idle extremely expensive compute.",
						"This is the practical meaning of silicon photonics for AI. The network is becoming part of the compute package. Once the training cluster is large enough, the distinction between accelerator design and fabric design starts to blur.",
					},
				},
				{
					Heading: "Lightmatter Pushes The Interposer",
					Paragraphs: []string{
						"Lightmatter is attacking the same pain point from another angle. Its Passage M1000 photonic interposer is advertised at 114 Tbps of total optical bandwidth across a 4,000 mm2 photonic interposer, with 256 optical fibers and a design meant to connect large die complexes that would be choked by conventional electrical I/O at the package edge.",
						"The important word is interposer. Passage is not trying to win by making a prettier cable. It is trying to make optical movement a native property of the package itself, so chips, memory, and switch fabrics can be arranged around bandwidth rather than around the physical shoreline where electrical pins fit.",
						"That matters because AI systems increasingly fail at the boundaries. A single accelerator can compute faster than the system can feed it. A rack can draw more power than the data center can deliver. A cluster can schedule more work than its fabric can move without wasting cycles. Photonic interconnects do not solve all of those problems, but they directly target the one that gets worse every time the model, context window, or parallelism strategy scales outward.",
					},
				},
				{
					Heading: "The Materials Are Catching Up",
					Paragraphs: []string{
						"The product story is also being pulled forward by materials work. Imec and Ghent University have demonstrated thin-film lithium niobate integration on a silicon photonics platform using micro-transfer printing, including a reported 320 Gb/s unamplified optical link over 2 km of standard single-mode fiber. That is not the same thing as a deployed AI switch, but it is the kind of device-level progress that makes higher-speed, lower-power optical lanes feel less exotic.",
						"Thin-film lithium niobate is attractive because it can produce very fast, efficient modulators. Silicon photonics is attractive because it can borrow manufacturing discipline from the semiconductor world. The hard part is putting those advantages together without inventing an impossible manufacturing flow. Imec's work points toward a path where new photonic materials can be added to silicon platforms in ways that remain compatible with industrial process thinking.",
						"That is the quiet theme underneath the hype: photonics is useful only if it can be manufactured, packaged, tested, and repaired at data-center scale. AI does not need beautiful one-off optical demos. It needs boring, serviceable, repeatable bandwidth.",
					},
				},
				{
					Heading: "The Toolchain Is Now Part Of The Race",
					Paragraphs: []string{
						"Hardware shifts do not become industries until design tools catch up. That is why Arizona State University's June 18 announcement about Jiaqi Gu's NSF CAREER award belongs in the same story. Gu's project is aimed at open-source electronic-photonic design automation, or EPDA, for large-scale photonics-empowered AI systems.",
						"The framing is important. If photonic links are going to move from boutique chips into mainstream AI infrastructure, engineers need tools that understand both electronic and photonic behavior early in design. They need to reason about layout, manufacturing constraints, system performance, and architecture together. Otherwise, photonics remains a specialist craft that scales slower than the AI systems it is supposed to connect.",
						"Open-source EPDA will not ship a switch by itself. But it is part of the same transition from lab demo to ecosystem: components, packages, switches, interposers, and design automation all have to mature at once.",
					},
				},
				{
					Heading: "Light Won The Wires, Not The Computer",
					Paragraphs: []string{
						"There is a temptation to turn every photonics story into a claim that optical computing is about to replace electronic computing. That is not what this moment shows. Computing with light remains an active and important research area, but the near-term commercial win is narrower and more concrete: light is winning the wires.",
						"That is still a big deal. The AI industry spent years treating networking as the thing you bought after the chips. In the largest clusters, networking is becoming one of the things that defines what the chips can do. The next generation of AI infrastructure will be judged not only by teraFLOPS, memory bandwidth, or watts per token, but by how efficiently it can move information across the machine.",
						"The practical future is not a glowing optical brain replacing the GPU. It is a data center where copper retreats to the shortest possible distances, optics moves closer to the silicon, and the AI cluster becomes a light-connected computer. The chip stopped being the only bottleneck. Then power and memory joined the story. Now the wire has, too.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"NVIDIA developer blog, \"Scaling AI Factories with Co-Packaged Optics for Better Power Efficiency\": https://developer.nvidia.com/blog/scaling-ai-factories-with-co-packaged-optics-for-better-power-efficiency/",
						"NVIDIA, \"Silicon Photonics Networking for Agentic AI\": https://www.nvidia.com/en-us/networking/products/silicon-photonics/",
						"Lightmatter, \"Passage M1000 Photonic Superchip\": https://lightmatter.co/products/m1000/",
						"Lightmatter, \"Lightmatter Unveils Passage M1000 Photonic Superchip, World's Fastest AI Interconnect\": https://lightmatter.co/press-release/lightmatter-unveils-passage-m1000-photonic-superchip-worlds-fastest-ai-interconnect/",
						"Imec, \"World first: integrating thin-film LiNbO3 modulators on a silicon photonics platform using micro-transfer printing\": https://www.imec-int.com/en/articles/double-world-first-integrating-lithium-niobate-and-lithium-tantalate-modulators-silicon",
						"ASU News, \"Illuminating the future of photonic AI chip design,\" June 18, 2026: https://news.asu.edu/b/20260618-illuminating-future-photonic-ai-chip-design",
					},
				},
			},
		},
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
					},
				},
			},
		},
	}, posts...)
}
