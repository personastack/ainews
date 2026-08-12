package content

func init() {
	posts = append([]Post{
		{
			Title:   "Meta's Muse Code Turns a Coding Model Into a Terminal-Native Agent",
			Slug:    "meta-muse-code-ai-coding-agent-2026",
			Date:    "August 7, 2026",
			Tag:     "DevTools",
			Summary: "Meta's beta Muse Code pairs the Muse Spark 1.2 coding model with a local terminal workflow, persistent context, and parallel subagents for long-running repository work.",
			Sections: []Section{
				{Paragraphs: []string{
					"Meta has put a new kind of coding tool into beta: Muse Code is a terminal-native agent designed to work through large repositories and long-running engineering tasks instead of stopping at a single generated function.",
					"The tool is powered by Muse Spark 1.2, Meta Superintelligence Labs' coding-focused model, and is aimed at developers who want an agent operating inside the same command-line environment where they build, test, and review software.",
				}},
				{Heading: "The workflow is the product", Paragraphs: []string{
					"Muse Code's important bet is architectural. It can break a project into tasks, delegate work to parallel subagents running in isolated worktrees, retain context across a long session, and bring results back for review. That makes it closer to a small engineering team sharing a terminal than to autocomplete with a chat window attached.",
					"The emphasis on persistent context matters for large codebases. An agent that can remember earlier decisions and compact its history has a better chance of surviving migrations, debugging sessions, and feature work that do not fit into one prompt.",
				}},
				{Heading: "Local control, frontier capability", Paragraphs: []string{
					"Running in a developer's CLI also puts the agent closer to the files, tests, and tools that define whether a change is useful. That can shorten the loop from plan to implementation to verification, while still leaving a human responsible for approving changes and deciding what reaches production.",
					"The tradeoff is that terminal access makes mistakes more consequential. Sandboxing, permissions, review steps, and clear worktree boundaries are not optional polish; they are the control surface for an agent that can make many changes in parallel.",
				}},
				{Heading: "What to watch", Paragraphs: []string{
					"Muse Code arrives as coding agents converge on the same shape: persistent sessions, delegated subtasks, tool use, and a local execution environment. The differentiator will not be whether an agent can produce a plausible patch. It will be whether it can keep a coherent plan, validate its work, and make its decisions easy for a human engineer to audit.",
				}},
				{Heading: "Sources", Paragraphs: []string{
					"Meta AI, Introducing Muse Spark 1.1: https://ai.meta.com/blog/introducing-muse-spark-meta-model-api/",
					"Meta AI, Introducing Muse Code and Muse Spark 1.2: https://research.meta.ai/blog/introducing-muse-code-and-muse-spark-1-2",
					"MarkTechPost, Meta Superintelligence Labs releases Muse Code: https://www.marktechpost.com/2026/08/05/meta-superintelligence-labs-releases-muse-code/",
				}},
			},
		},
	}, posts...)
}
