package site

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/personadock/ainews/internal/content"
)

//go:embed templates/*.html
var templatesFS embed.FS

const postsPerPage = 12

type Server struct {
	templates *template.Template
	mux       *http.ServeMux
}

type listPageData struct {
	Title      string
	Deck       string
	Posts      []content.Post
	Page       int
	TotalPages int
	TotalPosts int
	HasPrev    bool
	HasNext    bool
	PrevURL    string
	NextURL    string
	Start      int
	End        int
}

type postPageData struct {
	Title string
	Post  content.Post
}

type apiPost struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Summary string `json:"summary"`
	Tag     string `json:"tag"`
	Slug    string `json:"slug"`
	URL     string `json:"url"`
}

func New() (*Server, error) {
	tmpl, err := template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	s := &Server{
		templates: tmpl,
		mux:       http.NewServeMux(),
	}

	s.routes()
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) routes() {
	s.mux.HandleFunc("/", s.handleIndex)
	s.mux.HandleFunc("/posts/", s.handlePost)
	s.mux.HandleFunc("/api/posts", s.handlePostsAPI)
	s.mux.HandleFunc("/healthz", s.handleHealthz)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	posts := content.Posts()
	page := pageFromRequest(r)
	pagePosts, totalPages, start, end := paginatePosts(posts, page)

	data := listPageData{
		Title:      "AI News Briefs",
		Deck:       "Fresh technical briefings on the model race, regulation, and product shifts shaping AI right now.",
		Posts:      pagePosts,
		Page:       page,
		TotalPages: totalPages,
		TotalPosts: len(posts),
		HasPrev:    page > 1,
		HasNext:    page < totalPages,
		PrevURL:    pageURL(page - 1),
		NextURL:    pageURL(page + 1),
		Start:      start,
		End:        end,
	}

	s.render(w, http.StatusOK, "index", data)
}

func pageFromRequest(r *http.Request) int {
	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		return 1
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		return 1
	}

	totalPages := totalPages(len(content.Posts()))
	if page > totalPages {
		return totalPages
	}

	return page
}

func paginatePosts(posts []content.Post, page int) ([]content.Post, int, int, int) {
	totalPages := totalPages(len(posts))
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	start := (page - 1) * postsPerPage
	if start > len(posts) {
		start = len(posts)
	}

	end := start + postsPerPage
	if end > len(posts) {
		end = len(posts)
	}

	return posts[start:end], totalPages, start + 1, end
}

func totalPages(totalPosts int) int {
	if totalPosts == 0 {
		return 1
	}

	pages := totalPosts / postsPerPage
	if totalPosts%postsPerPage != 0 {
		pages++
	}

	return pages
}

func pageURL(page int) string {
	if page <= 1 {
		return "/"
	}

	return fmt.Sprintf("/?page=%d", page)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/posts/")
	if slug == "" || strings.Contains(slug, "/") {
		http.NotFound(w, r)
		return
	}

	post, ok := content.FindBySlug(slug)
	if !ok {
		http.NotFound(w, r)
		return
	}

	s.render(w, http.StatusOK, "post", postPageData{
		Title: post.Title + " | AI News Briefs",
		Post:  post,
	})
}

func (s *Server) handlePostsAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	apiPosts := make([]apiPost, 0, len(content.Posts()))
	for _, post := range content.Posts() {
		apiPosts = append(apiPosts, apiPost{
			Title:   post.Title,
			Date:    post.Date,
			Summary: post.Summary,
			Tag:     post.Tag,
			Slug:    post.Slug,
			URL:     "/posts/" + post.Slug,
		})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(apiPosts)
}

func (s *Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok"))
}

func (s *Server) render(w http.ResponseWriter, status int, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := s.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}
