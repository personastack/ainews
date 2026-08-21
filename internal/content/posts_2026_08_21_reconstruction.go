package content

func init() {
	posts = append([]Post{{
		Title:   "Give Frontier AI Every Paper a Scientist Cited. It Still Can't Guess the Discovery.",
		Slug:    "reconstruction-benchmark-ai-research-idea-bibliography-2026",
		Date:    "August 21, 2026",
		Tag:     "Research",
		Summary: "Stanford's Reconstruction benchmark finds frontier models rarely recover a paper's central research idea from its bibliography alone, even after multi-agent orchestration raises accuracy to 36 percent.",
		Sections: []Section{
			{Paragraphs: []string{
				`The same models solving decades-old mathematics theorems and writing production software are stumbling on a task any curious PhD student could attempt: look at the papers a researcher referenced, and figure out what they were trying to prove.`,
				`A new benchmark published this week by researchers at Stanford University, Titan Holdings, and Prentis AI reveals a striking gap in what frontier language models can actually do when the rubber meets the road in scientific research. The study, titled *Reconstruction: A Blind Benchmark for Recovering Research Ideas from Pre-Publication Bibliographies*, asked seven of the world's most capable AI systems to deduce a paper's core research idea using only that paper's bibliography — none of the text, none of the conclusions, just the list of what the authors read.`,
				`The result: single models guessed correctly 3 to 15 percent of the time.`,
			}},
			{Heading: "Designing an Honest Test", Paragraphs: []string{
				`The researchers — including Shaolong Chen, Yanlin Fei, Nazhou Liu, Xinmiao Yu, Lei Li, and Rahul Thapa from Stanford, and Madalina Ciobanu, Qingqing Mao, and Ritankar Das from Titan Holdings and Prentis AI — went to considerable lengths to close the obvious escape hatches. They implemented a temporal citation cutoff so no model could reason about papers published after the seed paper; anonymous reference IDs to strip identifying information; and frozen bibliographies to prevent any leakage of the idea itself. An independent LLM judge then assessed whether a model's hypothesis described "the same core research idea" — not just topical overlap, but the identical research question and central claim.`,
				`Across 643 papers spanning six scientific domains — machine learning (120 papers), astronomy (85), chemistry (105), materials science (117), medicine (78), and physics (138) — the best-performing single model was Claude Opus 4.8, which scored 13.3 percent on average, with individual domain scores spanning the full 3–15 percent range. GPT-5.6 Sol Pro, Kimi K3, GLM 5.2, Gemini 3.1 Pro Preview, DeepSeek-V4-Pro, and Qwen3.7-Max all clustered in the same modest range. This isn't a question of one model pulling ahead — the frontier appears to share the same ceiling.`,
			}},
			{Heading: "Throwing More Agents At It", Paragraphs: []string{
				`The researchers didn't stop with single-model evaluation. They built a multi-agent pipeline that combined cross-model review with a Swiss tournament system: models generate parallel hypotheses, align them into slots, peer-review each other's work without access to external search, and then compete in a three-round tournament with strict conflict-of-interest rules — if you proposed a hypothesis, you can't judge it against another.`,
				`The result improved dramatically, but still tells a cautionary story. The multi-agent pipeline reached an average of 36 percent across domains, with medicine performing best at 41.6 percent, physics at 36.4 percent, materials science at 40.1 percent, chemistry at 38.4 percent, astronomy at 36.5 percent, and machine learning lagging at just 22.9 percent. That 2.4x lift over the best single-model baseline is meaningful engineering progress — but it also means that even the most sophisticated multi-agent orchestration available today gets the research idea wrong nearly two-thirds of the time.`,
			}},
			{Heading: "Why This Matters More Than a Benchmark Score", Paragraphs: []string{
				`The timing of this paper is deliberate. In August alone, the AI community has celebrated OpenAI's Astra model working through ten previously unsolved mathematics problems, DARPA successfully flying an AI-controlled F-16, and multiple frontier labs publishing results suggesting AI is becoming a genuine partner in scientific discovery. The Reconstruction benchmark offers a sober counterweight to that narrative.`,
				`Recovering a research idea from a bibliography is, in some sense, what human scientific intuition is for. A skilled researcher reading a pre-submission draft's references — even without the paper itself — builds a mental model of the knowledge gap the author is trying to fill, the tradition they're working within, the move they're about to make. That synthesis is a form of inductive reasoning about research as a creative act, not just pattern-matching over existing text.`,
				`What the Reconstruction data suggests is that current language models, despite their extraordinary breadth of knowledge, have a much weaker grasp of the structure of scientific originality. They know what has been published. They are considerably less reliable at inferring what is about to be discovered — and why.`,
				`The implications ripple outward. An AI system that can't reliably infer a research direction from its own context will struggle as an autonomous scientific co-investigator, regardless of how well it performs at execution once a direction is set. The benchmark illuminates a specific bottleneck: the hypothesis-generation step, the moment when a researcher looks at what they know and asks what question hasn't been asked yet.`,
				`Machine learning performed the worst of the six domains at just 22.9 percent even with multi-agent orchestration — perhaps unsurprising given how saturated ML literature is, and how hard it is to identify a genuinely novel contribution from a bibliography of highly interconnected papers. Medicine and materials science did best, possibly because the research-idea space in those domains has more clearly defined boundaries around known biological mechanisms or material properties.`,
			}},
			{Heading: "A Useful Calibration", Paragraphs: []string{
				`None of this diminishes what frontier AI can already do in research contexts — literature summarization, experimental design assistance, data analysis, and formal proof verification are all real and valuable contributions. But the Reconstruction results suggest the field may be conflating tool-use fluency with something deeper: the capacity for genuine scientific intuition.`,
				`The hardest part of science has never been reading the papers. It has always been knowing which question to ask next.`,
				`At 36 percent with the best available multi-agent orchestration, we're not there yet. The Reconstruction benchmark doesn't close the door on AI as a scientific partner — but it does clarify which room we still need to open.`,
			}},
			{Heading: "Sources", Paragraphs: []string{
				`arXiv 2608.16645: https://arxiv.org/abs/2608.16645`,
				`Reconstruction: A Blind Benchmark for Recovering Research Ideas from Pre-Publication Bibliographies (HTML full paper): https://arxiv.org/html/2608.16645`,
			}},
		},
		Related: []Link{
			{Title: "OpenAI's Next Model Solved Ten Decades-Old Math Problems. Getting Mathematicians to Believe It Might Take Longer.", Slug: "openai-astra-ten-math-proofs-non-sofic-groups-2026"},
			{Title: "AI Isn't Just Answering Physics Questions Anymore — It's Running the Experiments", Slug: "ai-lab-instrument-superconductor-neutron-star-simulation-2026"},
			{Title: "Science Gets a Lab Partner That Runs the Experiments", Slug: "self-driving-labs-ai-runs-experiments-2026"},
		},
	}}, posts...)
}
