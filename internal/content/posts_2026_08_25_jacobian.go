package content

func init() {
	posts = append([]Post{{
		Title:   "A Counterexample Broke the Jacobian Conjecture. The AI Origin Story Needs More Evidence.",
		Slug:    "jacobian-conjecture-claude-fable-5-levent-alpoge-counterexample-2026",
		Date:    "August 25, 2026",
		Tag:     "Research",
		Summary: "A new arXiv account records an explicit three-dimensional counterexample to the 1939 Jacobian conjecture and extends the construction to every higher dimension. Its disclosure also draws an important line between a mathematical result and a model's documented role in producing it.",
		Sections: []Section{
			{Paragraphs: []string{
				"For 87 years, the Jacobian conjecture asked a deceptively compact question about polynomial maps: if a map has a nonzero constant Jacobian determinant, must it have a polynomial inverse? An explicit counterexample in three complex dimensions now says no. That single result refutes the conjecture in its general form.",
				"The development is a genuine mathematical landmark. It is also a useful test of how AI-assisted discoveries should be reported. The public record supports a sharp result and a role for Claude Fable 5 in a subsequent paper. It does not, on the strongest primary source available, establish the more sweeping claim that a model independently found the first counterexample.",
			}},
			{Heading: "What the Counterexample Shows", Paragraphs: []string{
				"The preprint identifies the first counterexample as a polynomial map from C³ to C³ with Jacobian determinant identically -2. It maps three distinct input points to the same output, so it cannot be one-to-one and therefore cannot have the inverse the conjecture would require.",
				"This is not a numerical near miss. The paper gives an explicit map whose three components have degrees 7, 6, and 4, and says its polynomial identities and fiber counts were checked with exact rational arithmetic. Those are the details that turn a striking formula into a result other mathematicians can inspect and reproduce.",
			}},
			{Heading: "Three Dimensions Fall; Two Remain", Paragraphs: []string{
				"The first result is credited in the preprint to L. Alpöge, announced on July 19. The same account describes a one-variable family from Gallagher and a geometric explanation from Speyer, then develops a construction that produces counterexamples in every dimension greater than two.",
				"That scope matters. The general conjecture is refuted once a three-dimensional example exists, but the two-dimensional case remains open. The new work does not close that smaller problem; it instead clarifies exactly where the original universal statement fails.",
			}},
			{Heading: "Where the AI Claim Is Documented", Paragraphs: []string{
				"The available primary paper is authored by Shuhong Gao of Clemson University. Its disclosure says that the paper's main idea and framework are Gao's, and that Claude Fable 5 assisted with proofs and the write-up. The paper separately credits Alpöge with the first three-dimensional counterexample.",
				"In a July 21 exposition, Terence Tao writes that Alpöge credited Claude Fable 5 for assistance. That is meaningful evidence of AI-assisted mathematical work. But it is different from saying that the model independently discovered Alpöge's result, or that the first result should be attributed to an AI system. The public sources reviewed here do not document that stronger origin story.",
			}},
			{Heading: "Why the Distinction Matters", Paragraphs: []string{
				"AI can help produce ideas, explore cases, check algebra, draft proofs, and make a difficult result easier to communicate. Each role changes the research process, but each requires a different kind of evidence. A model-assisted proof in a named author's paper is not the same thing as an independently documented model discovery.",
				"The Jacobian result is exciting without collapsing those categories. Its lasting significance will depend on the usual mathematical process: experts checking the construction, reproducing the exact calculations, and building on a method that now reaches every dimension above two. Clear attribution is part of that verification culture, not a distraction from it.",
			}},
			{Heading: "Sources", Paragraphs: []string{
				"arXiv — Shuhong Gao, Counterexamples to the Jacobian conjecture in dimensions greater than two: https://arxiv.org/abs/2608.00222",
				"arXiv HTML version — author disclosure, introduction, explicit map, and exact-arithmetic verification: https://arxiv.org/html/2608.00222v1",
				"Terence Tao — A digestion of the Jacobian conjecture counterexample: https://terrytao.wordpress.com/2026/07/21/a-digestion-of-the-jacobian-conjecture-counterexample/",
			}},
		},
		Related: []Link{
			{Title: "Nvidia's AVO Put Claude Opus 5 at 100% on ARC-AGI-3's Public Set", Slug: "nvidia-avo-arc-agi-3-perfect-score-scaffolding-2026"},
			{Title: "Two Chinese AI Systems Just Made Math Olympiad History. The Other Four Claims Need an Asterisk.", Slug: "imo-2026-huawei-celia-xiaohongshu-official-ai-perfect-scores-2026"},
		},
	}}, posts...)
}
