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

## Testing

Run:

```sh
TMPDIR=/workspace/ainews/.tmp CGO_ENABLED=0 go test ./...
```

Use `CGO_ENABLED=0` in this environment because the default cgo toolchain may not have a working assembler. Content additions should update tests in `internal/content/content_test.go` for publication gating and slug lookup. If the new article should be visible on the homepage, include its escaped title expectation in `internal/site/site_test.go`.
