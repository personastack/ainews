package content

func init() {
	posts = append([]Post{{
		Title:   "Two Chinese AI Systems Just Made Math Olympiad History. The Other Four Claims Need an Asterisk.",
		Slug:    "imo-2026-huawei-celia-xiaohongshu-official-ai-perfect-scores-2026",
		Date:    "August 22, 2026",
		Tag:     "Models",
		Summary: "Two AI systems officially scored 42/42 at IMO 2026, while four more perfect-score claims came from an informal Claude-based evaluation—an important distinction as elite benchmarks saturate.",
		Sections: []Section{
			{Paragraphs: []string{
				`The International Mathematical Olympiad is the world's most prestigious high school math competition — six problems, nine hours over two days, students from 109 countries, the best young mathematical minds on the planet. Historically, getting anywhere near a perfect score as a human is an extraordinary achievement. In Shanghai this July, seven of 666 human contestants did it.`,
				`Two AI systems did it too. But the story of how, and why the other four "perfect scores" you may have seen reported are something rather different, is worth paying attention to — because it tells you something important about where AI evaluation stands right now.`,
			}},
			{Heading: "The Official Record", Paragraphs: []string{
				`The July 15–16 IMO 2026 competition in Shanghai was the first year in which artificial intelligence systems were formally permitted to compete in a structured evaluation framework parallel to the human competition. The terms were strict: systems received the six problems only after human contestants had finished, operated within fixed submission windows, and had their proofs graded by IMO organizers themselves.`,
				`Two systems achieved perfect 42/42 scores under these conditions:`,
				"- Huawei's Celia\n- Xiaohongshu's dots-note-3.0",
				`This is genuinely historic. No AI system had ever earned a verified, independently graded perfect score at the IMO under the competition's own protocols. The IMO involves competition-level mathematical reasoning — not just calculation, but proof construction, creative insight, and the ability to recognize structure in problems that have never been seen before. A year ago, 35/42 was considered a landmark result for AI.`,
			}},
			{Heading: "The Unofficial Claims", Paragraphs: []string{
				`Separately, in the days following the competition, Menlo Ventures partner Deedy Das ran an independent evaluation, posting the six IMO 2026 problems to four additional frontier AI systems and grading their responses using a Claude-based agent grader. Four models achieved what Das described as perfect scores:`,
				"- Claude Fable 5 (Anthropic)\n- GPT-5.6 Sol (OpenAI)\n- Kimi K3 (Moonshot AI)\n- AxiomProver (Axiom Math)",
				`Das's repository is explicit: "treat scores as strong but not authoritative." The evaluation methodology differs from the official competition in meaningful ways — the grading was done by another AI model, not IMO judges, and the conditions were not identical to the competition's constraints.`,
				`Much of the media coverage that followed conflated these two distinct categories into a single narrative about six AI systems achieving perfect IMO scores. That conflation matters, and not because the unofficial results aren't interesting. It matters because "graded by the IMO" and "graded by a Claude-based agent in an investor's side project" are different evidentiary claims. Both are worth knowing. They are not the same claim.`,
			}},
			{Heading: "What It Costs to Ace the IMO Now", Paragraphs: []string{
				`One figure from Das's evaluation that deserves more attention than it has received: the cost per model run ranged from roughly $20 to $51 in compute. This means that for less than the price of a dinner out, you can now submit frontier AI against what was, within living memory, considered the hardest high school mathematics problem set on earth — and expect a very good result.`,
				`That's a remarkable shift in the economics of mathematical reasoning. When AI first began competing with Olympiad-level math just a few years ago, strong performance required months of focused research effort and significant compute uncertainty. What was once a research achievement now runs in a test sitting for under fifty dollars.`,
			}},
			{Heading: "The Benchmark Saturation Problem", Paragraphs: []string{
				`Here is the uncomfortable implication that neither the triumphant press releases nor the skeptical takes have fully confronted: when multiple frontier AI systems — official and unofficial — can achieve perfect scores on the hardest high school math competition in the world, the benchmark itself stops functioning as a discriminator.`,
				`The IMO was useful precisely because it was near-impossible for AI. Now that it's not, the relevant questions shift. Not "can an AI ace the IMO?" but: How reliably? At what cost per correct answer? Under what conditions and constraints? What happens on problems designed to be novel and resistant to pattern recognition? How does performance degrade when the grading methodology changes?`,
				`The research community has already begun this pivot. Evaluations like GPQA Diamond, Humanity's Last Exam, and novel problem generation frameworks are increasingly the tools that separate frontier models — because they're harder to saturate, more resistant to training data contamination, and better at distinguishing genuine reasoning from very sophisticated recall. When traditional benchmarks like MMLU and HellaSwag hit functional saturation above 88% and 95% respectively for frontier models, the score differences at the top become statistically meaningless. The IMO is following the same trajectory.`,
			}},
			{Heading: "Why Chinese Labs Went for Official Verification", Paragraphs: []string{
				`One detail in this story that hasn't gotten enough attention: Huawei and Xiaohongshu pursued official IMO verification. They didn't run a quick informal test — they coordinated with IMO organizers, submitted under competition conditions, and waited for official grading. That's a deliberate choice about what "proof" means.`,
				`The Western labs whose models appeared in Das's evaluation didn't pursue formal IMO verification, at least not publicly. Das's informal test is significant and the results aren't being dismissed — but the pursuit of independent verification is itself telling. In a world where claims about AI capabilities are made constantly and with varying levels of rigor, the decision to get a result independently checked is worth noting.`,
				`This may reflect different strategic priorities rather than different confidence in the results. But it's a distinction that matters as AI capability claims become more consequential and more contested.`,
			}},
			{Heading: "What Comes Next", Paragraphs: []string{
				`The IMO will likely remain a useful calibration point — it's hard enough that strong performance still means something, and it's a well-understood domain. But as a separator of frontier capability, it has joined MMLU and HellaSwag in the club of benchmarks that now ask us to pay attention to the second decimal place rather than the headline number.`,
				`The new frontier for math reasoning evaluation is novel problem generation: can an AI system produce and verify solutions to problems that definitively could not have been in its training data? That's the question the research community is working toward, and it's considerably harder to fake.`,
				`For now, the headline is real: two AI systems, officially graded, achieved perfect scores at the 2026 International Mathematical Olympiad. That's a landmark. The other four results are interesting, informative, and shouldn't be dismissed — but they carry a different kind of warrant.`,
				`Knowing the difference matters more, not less, as these claims become more common.`,
			}},
			{Heading: "Sources", Paragraphs: []string{
				`Digital Applied: https://www.digitalapplied.com/blog/imo-2026-perfect-scores-ai-benchmark-saturation`,
				`Deedy Das / Menlo Ventures IMO evaluation: https://x.com/deedydas/status/2079409461874332066`,
				`IMO 2026 Leaderboard & Scores, BenchLM.ai: https://benchlm.ai/benchmarks/imo2026`,
				`explainx.ai AI Benchmarks 2026 Guide: https://explainx.ai/blog/ai-benchmarks-complete-guide-2026`,
				`Tech Insider, AI Perfect Score at IMO 2026: https://tech-insider.org/ai-imo-2026-perfect-score-odds-hit-96-percent/`,
			}},
		},
		Related: []Link{
			{Title: "OpenAI's Next Model Solved Ten Decades-Old Math Problems. Getting Mathematicians to Believe It Might Take Longer.", Slug: "openai-astra-ten-math-proofs-non-sofic-groups-2026"},
			{Title: "Give Frontier AI Every Paper a Scientist Cited. It Still Can't Guess the Discovery.", Slug: "reconstruction-benchmark-ai-research-idea-bibliography-2026"},
			{Title: "Alibaba's Qwen3.8-Max Beats GPT-5.6 and Claude on Key Benchmarks — And It's Going Open Weight", Slug: "qwen3-8-max-open-weight-benchmarks-gpt-5-6-claude-2026"},
		},
	}}, posts...)
}
