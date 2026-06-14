# AINews Technical Spec

AINews is a small Go web service that publishes static AI news articles at `ainews.personastack.ai`. The service keeps runtime behavior simple: article content is compiled into the binary, templates render the homepage and article pages, and a JSON endpoint exposes the published feed for promotion and monitoring.

## Architecture

- `cmd/ainews/main.go` starts the HTTP server and reads `PORT`, defaulting to `8080`.
- `internal/site` owns routing, template parsing, HTML rendering, the `/api/posts` JSON feed, and `/healthz`.
- `internal/content` owns the in-memory article catalog and publication filtering.
- `Dockerfile` is a multi-stage container build that runs `CGO_ENABLED=0 go test ./...` before compiling a static Linux binary into a distroless runtime image.

## Content Model

Articles are `content.Post` values with:

- `Title`, `Slug`, `Date`, `Tag`, and `Summary` string metadata.
- `Sections`, a list of `content.Section` values with an optional `Heading` and paragraph strings.

The base catalog lives in `internal/content/content.go`. New dated batches live in files named `internal/content/posts_YYYY_MM_DD.go`. Each dated file prepends its posts during `init()` with:

```go
posts = append([]Post{
    // new posts
}, posts...)
```

Prepending keeps newer articles at the front of the homepage and `/api/posts` feed.

## Publishing Rules

Published posts are gated by the `Date` field using Go's `January 2, 2006` layout. `Posts()` and `FindBySlug()` hide posts dated after the current UTC date. Article additions should use the intended publication date and include a final `Sources` section with the backing links from Author or Researcher.

AINews currently stores one primary category in `Post.Tag`. Secondary tags from editorial handoff can be reflected in article copy, but the rendered site only exposes the single primary tag.

The June 10, 2026 content batch includes a DevTools article about Cohere North Mini Code as a local, open, controllable coding-agent layer for enterprises. It should keep the caveat that Cohere's throughput and inter-token latency comparisons are company-reported and should frame the model as routable local or sovereign infrastructure rather than a universal replacement for frontier cloud systems.

The same batch also includes a Platforms article about OpenAI turning ChatGPT reasoning effort into a visible picker control. The story is intentionally framed as product UX and compute allocation, not as a new model launch.

The June 11, 2026 content batch includes an Enterprise article about Nasdaq Verafin's planned agentic AML and fraud analysts. Keep availability timing precise: general availability is expected in Q3 2026, the third-party deployment model is expected to enter beta in the second half of 2026, and alert auto-dispositioning should be framed as a planned roadmap capability rather than broad production autonomy.

The same batch also includes an Enterprise article about Atos managing a large AI-agent rollout with Microsoft. Keep the rollout scale framed as company-reported: 56,000 employees, 54 countries, and 19,000 AI agents. Preserve the product-naming caveat that Microsoft's source uses Microsoft 365 Copilot E7 while Atos also uses Microsoft 365 E7 wording.

The same batch also includes a Security article about Hong Kong's SFC treating AI cyber risk as financial regulation. Preserve the June 2, 2026 circular date, use "AI-driven" when referring directly to the SFC circular's title and framing, and keep prompt injection framed as technical context rather than a direct SFC quote.

The June 12, 2026 content batch includes an AI Infrastructure article about the World Economic Forum's 2026 Technology Pioneers cohort as a signal that AI startup activity is moving toward agent payments, identity, GPU orchestration, energy, grid systems, and vertical deployment infrastructure. Preserve the June 10, 2026 WEF announcement date, the 100-company and 23-country cohort figures, and the framing that Technology Pioneers lists are market-map signals rather than predictions of winners.

The same batch also includes a Policy article about Anthropic asking for legally bounded deployment-blocking authority while the June 2, 2026 White House order builds classified benchmarking, covered-frontier-model designation, and voluntary 30-day pre-release access. Preserve the distinction that the U.S. order does not adopt mandatory licensing, preclearance, permitting, or a deployment-veto regime. Keep the Hong Kong SFC comparison framed as operational-risk supervision for AI-enabled cyber threats, not frontier-model approval or model veto authority.

The same batch also includes a Security article about Cyera's $600 million funding round at a $12 billion valuation as a signal that enterprise AI data governance and trust infrastructure are becoming strategic markets. Preserve the caveat that the round size, valuation, total funding over $2.3 billion, investor list, annual growth, employee/acquisition figures, 95%+ precision claim, and 100+ product capabilities are company-reported or reported funding facts. Do not frame the article as proof that Cyera has solved agentic security; keep the claims to data visibility, access governance, DLP/privacy/identity convergence, and runtime controls for AI use.

The same batch also includes a Science article about TITO, or Transferable Implicit Transfer Operators, as a way to compress molecular dynamics simulation between candidate proposal and lab validation. Preserve the reported scope: more than 10,000x faster than conventional simulations in the studied settings, tested on more than 12,500 organic molecules and more than 1,000 short peptides, demonstrated on small molecular systems and simplified solvent models. Do not frame it as turnkey drug discovery or immediate clinical acceleration.

The same batch also includes an Enterprise article about agent infrastructure entering an IAM, observability, FinOps, and workflow-governance era. Preserve the framing that major platforms are preparing production controls and selected workflows are moving into operational use; do not claim broad autonomous deployment across all enterprises. Treat vendor maturity figures and maturity-index claims as vendor research if used.

The same batch also includes an Accessibility article about Meta's Ray-Ban Meta AI glasses donation program for legally blind U.S. veterans. Preserve the framing as a Meta/BVA/nonprofit partner accessibility distribution and training program, not a Department of Veterans Affairs program based on current sources. Do not describe the glasses as restoring sight, FDA-cleared medical technology, or clinically validated safety-critical vision systems. Keep distribution language careful: eligible veterans can request glasses through the program, but the article should not imply every device has already been delivered.

The June 13, 2026 content batch includes a Workforce article about workplace AI "botsitting" as hidden supervision labor. Preserve the caveat that Glean's 11-hour weekly AI time-saved figure and 6.4-hour weekly botsitting figure are survey-reported worker experience, not independently measured productivity, and that Glean is an enterprise AI vendor. Frame botsitting as the checking, correcting, prompting, and repair work needed to make AI output usable, not as proof that AI is useless.

The same batch also includes an Infrastructure article about OpenText's planned EUR105 million Ireland investment across Cork and Galway as a signal that enterprise agents need geography, data-residency, cybersecurity, and trust boundaries. Preserve that the investment size, 400 planned jobs over three years, and product focus are company-announced figures. Do not imply the 400 jobs already exist or that OpenText has solved sovereign agentic AI deployment.

The same batch also includes a Security article about agent registries and control planes becoming a practical enterprise AI security boundary. Preserve the framing as a security-infrastructure signal rather than proof that Zscaler or any other vendor has solved agentic AI security. Keep Zscaler's AI Broker, AI Access Graph, Agent Registry, and ZAgent details product-specific; do not include unverified CISA/BOD claims, and distinguish agent registry/control-plane needs from broader general agent infrastructure.

The same batch also includes a Policy article about the U.S. export-control directive that led Anthropic to suspend Claude Fable 5 and Claude Mythos 5 access. Preserve the caveats that other Claude models are not reported down, the cutoff should not be described as permanent, the directive text/legal pathway/detailed jailbreak evidence are not public in full, and Anthropic's technical rebuttal is Anthropic's account. Frame the June 2 White House executive-order tension carefully: the EO rejects mandatory licensing/preclearance while a separate national-security/export-control action appears to have caused an emergency model-access cutoff.

The June 14, 2026 content batch includes a Policy article about AI safety moving from voluntary lab governance into courts, consumer-protection enforcement, cybercrime disruption, telecom/FBI coordination, and G7 diplomacy. Preserve the caveats that Google alleges an AI-powered smishing/phishing network and that complaint coverage reports Gemini/AI tools were used to help create phishing infrastructure; OpenAI's state-attorney-general scrutiny should be framed as a reported multistate probe without asserting exact subpoena details as independently confirmed; and G7 AI outputs should be described as political coordination and diplomacy rather than binding domestic law.

The same batch also includes a Policy article about CMS creating an Office of Health Technology and Products as reported by Healthcare Dive. Preserve the caveat that Healthcare Dive is the current public source for the specific OHTP details unless an official CMS announcement is added. Do not frame the office as a new AI regulator, clinical-safety regulator, AI approval authority, or sole regulator of healthcare AI. Keep CMS framed as a central payment and policy actor, and preserve the distinction between AI implementation, digital health tools, healthcare data exchange, and clinical AI validation.

The same batch also includes a Platforms article about Visa's OpenAI collaboration for ChatGPT agentic commerce. Preserve the distinction between Visa's payment-network, tokenization, authorization, agent-identification, and fraud-control role and OpenAI's broader commerce and checkout work. Do not frame this as fully autonomous unrestricted ChatGPT spending. Keep the article focused on bounded agentic commerce infrastructure with user limits, approved merchants, approval steps, authorization, agent identification, fraud monitoring, and tokenized credentials. Financial terms, merchant fee structures, and full rollout mechanics are undisclosed.

The same batch also includes an Enterprise article about TCS joining Anthropic's Claude Partner Network as a Global Premier Partner. Keep the 50,000-employee and 56-country figures framed as official company-reported scale claims. Do not imply Claude is automatically approved for regulated industries or already broadly deployed across every named sector. Preserve the framing as systems integration, deployment, governance, training, certification, reusable workflow packaging, and planned industry offerings such as claims processing and lending advisory.

## Testing

Run:

```sh
TMPDIR=/workspace/ainews/.tmp CGO_ENABLED=0 go test ./...
```

Use `CGO_ENABLED=0` in this environment because the default cgo toolchain may not have a working assembler. Content additions should update tests in `internal/content/content_test.go` for publication gating and slug lookup. If the new article should be visible on the homepage, include its escaped title expectation in `internal/site/site_test.go`.
