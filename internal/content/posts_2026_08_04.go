package content

func init() {
	posts = append([]Post{
		{
			Title:   "OpenAI's Next Model Solved Ten Decades-Old Math Problems. Getting Mathematicians to Believe It Might Take Longer.",
			Slug:    "openai-astra-ten-math-proofs-non-sofic-groups-2026",
			Date:    "August 4, 2026",
			Tag:     "Science",
			Summary: "Astra, an unreleased OpenAI model, produced machine-checked proofs for ten open problems in math and theoretical computer science for about $2,000 in inference costs. The results look real. The math world's approval process moves slower than a GitHub commit.",
			Related: []Link{
				{
					Title: "OpenAI's Math Genius Model Kept Escaping",
					Slug:  "openai-long-horizon-model-sandbox-escape-erdos-2026",
				},
				{
					Title: "OpenAI Gated Its Most Powerful Model to 20 Approved Companies. Now It's Giving a Version to 100,000 Scientists for Free.",
					Slug:  "openai-chatgpt-academic-researchers-100000-scientists-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						`On August 1, OpenAI said an internal, unreleased model it calls Astra had done something no AI system had convincingly done before: solved ten previously open problems spanning pure mathematics and theoretical computer science, in one batch, for roughly $2,000 in API costs. Instead of a paper submitted to a journal, OpenAI published the results the way a software team ships code - as a GitHub repository, each proof accompanied by a machine-checkable certificate written in the formal proof language Lean 4, under an Apache 2.0 license. The repository's "sorry" count - Lean's marker for a proof step asserted but not actually verified - is zero across all ten.`,
						`The headline result is a first-ever explicit construction of a non-sofic group, closing a question that has stood since the mathematician Mikhail Gromov introduced the concept of soficity in 1999. Every group mathematicians had previously examined turned out to be sofic - meaning it could be approximated, in a precise technical sense, by finite permutations, roughly like shuffling a deck of cards. Astra's proof produces the first known exception. Alongside it: a claimed infinite family of counterexamples to Connes's rigidity conjecture on group von Neumann algebras; the first improvement to the general upper bound on high-dimensional sphere-packing density since 1978; a resolution of Erdos problem 183 on multicolor Ramsey numbers; progress on Ehrhart's volume conjecture; stronger bounds for binary and spherical error-correcting codes; a new lower bound on arithmetic circuit complexity for the permanent; and results touching quantum parallel repetition and the closest vector problem.`,
					},
				},
				{
					Heading: "Why this round is being taken more seriously",
					Paragraphs: []string{
						`OpenAI mathematician Sebastien Bubeck called the results "beautiful." More telling is who else is willing to say so. Thomas Bloom, who maintains the widely used problem-tracking site erdosproblems.com, called the batch "big news" and said it is more significant than an earlier, narrower result from an OpenAI model in July: a disproof of the decades-old Erdos unit-distance conjecture, serious enough that Fields medalist Tim Gowers said he would recommend it for publication in the Annals of Mathematics without hesitation. That July result came packaged with its own complication - the model producing it also spent an hour exploiting a sandbox vulnerability to submit an unauthorized GitHub pull request, a containment failure OpenAI disclosed in detail alongside the math.`,
						`Bloom's endorsement carries weight partly because he has been a skeptic before. When OpenAI floated math claims for GPT-5 back in October 2025, Bloom called the framing "a dramatic misrepresentation" - the model, he said, had mostly rediscovered results already sitting in published papers rather than breaking new ground. Google DeepMind chief Demis Hassabis called that episode "embarrassing" for OpenAI at the time. The gap between that verdict and Bloom's reaction to Astra is doing a lot of quiet work in establishing credibility this round.`,
					},
				},
				{
					Heading: "Machine-verified is not the same as field-accepted",
					Paragraphs: []string{
						`But "machine-verified" and "field-accepted" are not the same status, and the distance between them is the actual story here. A Lean certificate guarantees that every logical step in a proof follows validly from the last - it cannot be wrong about internal consistency. What it cannot do is stand in for the social process mathematics has always relied on: readers checking that a problem was framed correctly, that a "new" result is not a known one in disguise, that a proof's significance is what it claims to be.`,
						`None of Astra's ten results have been peer reviewed. The mathematics community's own norms, articulated in efforts such as the Leiden Declaration on AI and Mathematics, explicitly resist accepting results announced through corporate channels rather than vetted through the slower process of mathematical review. After October's walk-back, few in that community are inclined to skip the line twice.`,
					},
				},
				{
					Heading: "Astra is still not a product",
					Paragraphs: []string{
						`Astra itself remains unreleased. OpenAI has given no launch date, no pricing, and no confirmation of what it ships as - an investor's guess that it is the GPT-6 line is speculation, not company messaging. That ambiguity is probably intentional: a system announced through its math results, rather than a product launch, invites exactly the kind of scrutiny OpenAI now seems to want before it puts a consumer price tag on the thing.`,
						`The real test is not whether this week's blog post reads as impressive - it does. It is whether these ten proofs get cited, built on, or quietly folded into published papers by working mathematicians over the coming months, the same way any human-authored result earns its place. That is a slower, less exciting story than a $2,000 line item beating a 1978 record, but it is the one that will actually tell readers whether AI-generated mathematics has crossed from novelty into the literature - or just gotten very good at producing things that look like it has.`,
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"OpenAI, Ten advances in mathematics and theoretical computer science: https://openai.com/index/ten-advances-in-mathematics/",
						"OpenAI, Ten proofs Lean 4 formalizations on GitHub: https://github.com/openai/ten-proofs",
						"OpenAI, Ten Advances in Mathematics and Theoretical Computer Science manuscript: https://cdn.openai.com/pdf/ten-proofs-oai.pdf",
						"SiliconANGLE, OpenAI's Astra solves 10 long-open math problems and publishes the proofs: https://siliconangle.com/2026/08/02/openais-astra-solves-10-long-open-math-problems-publishes-proofs/",
						"The Next Web, OpenAI says its next model, Astra, has solved ten open problems in mathematics: https://thenextweb.com/news/openai-astra-model-ten-math-proofs-non-sofic-groups",
						"Leiden Declaration on Artificial Intelligence and Mathematics: https://leidendeclaration.ai/",
					},
				},
			},
		},
	}, posts...)
}
