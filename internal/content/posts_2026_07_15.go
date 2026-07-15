package content

func init() {
	posts = append([]Post{
		{
			Title:   "The Rocket Company Ships a Coding Model — And the Benchmark Depends on Who's Grading",
			Slug:    "grok-4-5-spacexai-cursor-coding-benchmark-harness-2026",
			Date:    "July 15, 2026",
			Tag:     "Models",
			Summary: "Grok 4.5 is the first model to come out of SpaceXAI, the merged SpaceX-xAI-Cursor colossus, and it undercuts Claude on price while splitting benchmarks down the middle. The more interesting story than \"who wins\" is why the same test gives opposite answers depending on who's holding the stopwatch.",
			Related: []Link{
				{
					Title: `The Machine Learned to Say "Mhmm"`,
					Slug:  "openai-gpt-live-full-duplex-voice-end-of-turn-taking-2026",
				},
			},
			Sections: []Section{
				{
					Paragraphs: []string{
						"Six months ago, SpaceX built rockets and xAI built chatbots. Today they're the same company, that company just spent $60 billion in stock buying the AI coding editor Cursor, and its first product under the new arrangement is a language model tuned specifically to write software. If you'd predicted that sequence of events in January, you'd have been laughed out of the room. It happened anyway, and on July 8th the model — Grok 4.5 — shipped.",
						"The corporate mechanics are worth a paragraph before the technology, because they explain what Grok 4.5 is actually for. xAI merged with X in early 2025, then that combined entity merged into SpaceX at a board-set valuation of roughly $1 trillion for SpaceX alone. SpaceX went public on the Nasdaq in June, and within days used its newly-public stock as acquisition currency to buy Cursor. By July 6th the whole thing had a new name: SpaceXAI. Two days later, Grok 4.5 arrived as its first release — trained \"alongside Cursor,\" in xAI's own words, and built into every tier of Cursor's product from day one. A rocket company's flagship AI move, it turns out, was to become a developer tools company.",
						"On paper, Grok 4.5 is a capable, aggressively priced model. It carries a 500,000-token context window, an adjustable reasoning-effort dial (low, medium, or high — though unlike most competitors, it can't be switched off entirely), and both text and image input. Pricing lands at $2 per million input tokens and $6 per million output, with cached input at 50 cents — roughly a third of what Anthropic charges for Opus 4.8 on input, and less than a quarter on output. It's available through Cursor, the xAI API, Grok's own apps, Microsoft Office add-ins, and most of the major inference gateways, though not yet in the EU, where a rollout is expected later this month.",
						"The benchmark story is where it gets genuinely interesting — not because Grok 4.5 wins or loses, but because it does both, on the same test, depending on who's grading. On the independent Artificial Analysis Intelligence Index, Grok 4.5 ranks fourth overall, behind Anthropic's Fable 5 and Opus 4.8 and OpenAI's GPT-5.5. It posts a strong 93.1% on GPQA Diamond and takes the top spot on agentic tool use. On Terminal-Bench 2.1, a coding-agent benchmark, it beats Opus 4.8 handily: 83.3% to 78.9%.",
						"But look at DeepSWE, a benchmark xAI itself chose to publish, and the story flips depending on the harness. Run under xAI's own provider harness — the test environment the vendor controls — Grok 4.5 beats Opus 4.8, 62.0% to 55.75%. Run under a neutral, third-party harness on the same benchmark family, Opus 4.8 beats Grok 4.5, 59% to 53%. Same two models, same underlying task, opposite winner, and the only variable that changed is who built the scaffolding the models were tested in.",
						"That's not a scandal — harness effects are a known, documented issue in AI benchmarking, and xAI published both numbers rather than hiding the less flattering one. But it's a useful reminder for anyone reading a vendor's launch-day benchmark chart: \"beats the competition\" is doing a lot of quiet work in that sentence, and the honest version of the claim usually includes the word \"sometimes,\" followed by a footnote about who ran the test. Grok 4.5 also produces roughly 60% fewer output tokens than Opus 4.8 to solve the same Intelligence Index tasks — a real efficiency edge that matters more at production scale than any single benchmark percentage point, and one worth watching independently of who's ahead this week.",
						"Zoom out and the more durable story isn't the scorecard, it's the shape of the company that produced it. A rocket-and-satellite business, a chatbot lab, and a code editor are now one balance sheet, and the first thing that balance sheet chose to build was a model that writes software faster and cheaper than the incumbent. That's a bet on where the money is, not on what's flashiest. As AI labs increasingly own the whole stack — the model, the coding tool it plugs into, and the benchmark harness used to judge it — independent, third-party evaluation matters more than ever, not less. The vendor with the most control over the test is rarely the one who should be trusted to grade it alone.",
					},
				},
				{
					Heading: "Sources",
					Paragraphs: []string{
						"Author article handoff and archive doc: https://docs.google.com/document/d/1NFXhAZ-FnnwCPiZD_kfi69Lfi-qLAwxUsn_MgbZdZiM/edit",
						"Author reported independent WebSearch and WebFetch sourcing against kingy.ai, benchlm.ai, Yahoo Finance, and general search as of July 15, 2026.",
					},
				},
			},
		},
	}, posts...)
}
