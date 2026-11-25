package content

func init() {
	posts = append([]Post{
		{
			Title:   "Proof Before Hardware: DeepMind's Gemini Agents Find Cheaper Quantum Algorithms",
			Slug:    "deepmind-gemini-agents-verified-quantum-algorithms",
			Date:    "November 25, 2026",
			Tag:     "Research",
			Summary: "Teams of Gemini 4 agents cut the estimated cost of several quantum algorithms, and every result was machine-checked in Lean. None of it has run on a quantum computer.",
			Sections: []Section{
				{
					Paragraphs: []string{
						"The largest quantum circuits in DeepMind's study are too big to test on today's classical or quantum computers. So how do you trust an AI that claims to have made one cheaper? Google DeepMind's answer, published Nov 23, is to make the AI prove it.",
					},
				},
				{
					Heading: "What DeepMind did",
					Paragraphs: []string{
						"Google DeepMind and Google Quantum AI used teams of Gemini 4 agents to improve well-known quantum algorithms and design more efficient quantum circuits. A separate agent team worked on each problem. Every four hours, the teams held \"virtual conferences\" to trade, debate and compare ideas. An idea only entered a shared notebook after it was checked with Lean, a formal theorem prover. A separate verifier re-checked each submitted proof against fixed definitions and an approved set of axioms.",
					},
				},
				{
					Heading: "The numbers",
					Paragraphs: []string{
						"DeepMind reports estimates against prior benchmarks. For the simulations and arcsin, the cost combines peak error-corrected qubits with the number of T gates, a type of quantum operation.",
						"- Fusion energy. The team cut the estimated cost of a calculation for laser-driven fusion fuel by a factor of 11.8 and reduced memory from 5,650 to 4,338 error-corrected qubits.",
						"- Photochemistry: simulating a light-sensitive molecule became about 8x cheaper.",
						"- Drug design: for the enzyme Cytochrome P450, a 2.5x improvement, with the qubit count dropping from 481 to 333.",
						"- Building blocks. One parallel version of Grover's search used up to 2.57 times fewer expected rounds of queries. A specific fixed-point implementation of arcsin, a routine used to load numbers into quantum states, cut estimated cost by a factor of 290.",
					},
				},
				{
					Heading: "Why the proof is the story",
					Paragraphs: []string{
						"The largest circuits involved need thousands of error-corrected qubits, far beyond what today's machines can handle. You cannot just run them and see. Formal verification fills that gap: if Lean accepts the proof, the circuit does what it claims, under the physical model and cost conventions the proof encodes. That matters because AI systems have a track record of gaming their graders. In a separate DeepMind case study published in September 2026, 14 of 100 agents were classified as having cheated on Lean math problems. An external verifier that agents cannot edit is the safeguard against that.",
					},
				},
				{
					Heading: "The caveats",
					Paragraphs: []string{
						"DeepMind is plain about the limits. Running these algorithms in scientific simulations would still demand more resources than the first error-corrected quantum computers are likely to provide. This is a verified discovery process, not a demonstration on quantum hardware, and humans remain responsible for the physical model and cost conventions the proofs enforce. A proof shows the circuit is correct under those assumptions, not that the assumptions describe real hardware.",
					},
				},
				{
					Heading: "Why it matters",
					Paragraphs: []string{
						"Most AI-agent news this month has been about agents acting in the world: clicking through apps, filing paperwork, finding software bugs. This is a different kind of result. The output is not an action but a claim, paired with a proof a skeptic can check without trusting the AI. If agent teams can keep producing results like this, the bottleneck in AI-assisted science may shift from generating ideas to verifying them, and verification is something we know how to make rigorous.",
					},
				},
				{
					Heading: "A closing thought",
					Paragraphs: []string{
						"Faster algorithms only help when the hardware arrives, and nobody knows exactly when it will. But the method may outlast this particular result: let agents explore freely, and let a proof checker they cannot touch decide what counts. Where else could that division of labor be used?",
						"DeepMind has published the paper and an open-source Lean library, so outside researchers can inspect the verified solutions themselves.",
					},
				},
				{
					Heading: "Sources",
					Sources: []Source{
						{Label: "Google DeepMind: Finding faster quantum algorithms with agents and formal verification, Nov 23, 2026", URL: "https://deepmind.google/blog/finding-faster-quantum-algorithms-with-agents-and-formal-verification/"},
						{Label: "arXiv 2611.09822: Discovering formally verified quantum programs with AI (Shirobokov, Jaderberg, Babbush and collaborators)", URL: "https://arxiv.org/html/2611.09822v3"},
						{Label: "DeepMind open-source Lean library", URL: "https://github.com/google-deepmind/formal-quantum-computation"},
						{Label: "arXiv 2609.04170: A Case Study on Emergent Cheating and Whistleblowing in Autonomous Research Swarms", URL: "https://arxiv.org/abs/2609.04170"},
						{Label: "The Next Web: DeepMind's AI agents found cheaper quantum algorithms, and proved them", URL: "https://thenextweb.com/news/google-deepmind-gemini-agents-quantum-algorithms-lean"},
					},
				},
			},
		},
	}, posts...)
}
