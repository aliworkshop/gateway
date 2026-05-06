package gateway_test

import (
	"net/http"
	"testing"

	"github.com/aliworkshop/errors"
	"github.com/aliworkshop/gateway/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestPaginatorDefaults(t *testing.T) {
	p := gateway.NewPaginator()
	if got := p.GetPage(); got != 1 {
		t.Errorf("default page = %d, want 1", got)
	}
	if got := p.GetPageSize(); got != 10 {
		t.Errorf("default page size = %d, want 10", got)
	}
	if got := p.Total(); got != 0 {
		t.Errorf("default total = %d, want 0", got)
	}
}

func TestPaginatorSetGet(t *testing.T) {
	p := gateway.NewPaginator()
	p.SetPage(5)
	p.SetPageSize(50)
	p.SetTotal(123)
	if got := p.GetPage(); got != 5 {
		t.Errorf("page = %d, want 5", got)
	}
	if got := p.GetPageSize(); got != 50 {
		t.Errorf("page size = %d, want 50", got)
	}
	if got := p.Total(); got != 123 {
		t.Errorf("total = %d, want 123", got)
	}
}

func TestPaginatorClampsNonPositivePage(t *testing.T) {
	p := gateway.NewPaginator()
	p.SetPage(0)
	if got := p.GetPage(); got != 1 {
		t.Errorf("page(0) = %d, want clamped to 1", got)
	}
	p.SetPage(-3)
	if got := p.GetPage(); got != 1 {
		t.Errorf("page(-3) = %d, want clamped to 1", got)
	}
}

func TestSorterDefault(t *testing.T) {
	s := gateway.NewSorter()
	if got := s.GetSortBy(); got != "-id" {
		t.Errorf("default sort = %q, want -id", got)
	}
}

func TestSorterSetGet(t *testing.T) {
	s := gateway.NewSorter()
	s.SetSort("created_at")
	if got := s.GetSortBy(); got != "created_at" {
		t.Errorf("sort = %q, want created_at", got)
	}
}

func TestLanguageLocalizeFallback(t *testing.T) {
	bundle := i18n.NewBundle(language.English)
	lang := gateway.NewLanguage(bundle, "en")
	got, err := lang.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: "hello", Other: "Hello, world"},
	})
	if err != nil {
		t.Fatalf("Localize: unexpected error: %v", err)
	}
	if got != "Hello, world" {
		t.Errorf("Localize = %q, want Hello, world", got)
	}
}

type stubRouter struct {
	called   string
	path     string
	handlers int
}

func (s *stubRouter) Group(string) gateway.RouterGroupModel { return s }
func (s *stubRouter) READ(p string, h ...gateway.Handler) {
	s.called, s.path, s.handlers = "READ", p, len(h)
}
func (s *stubRouter) CREATE(p string, h ...gateway.Handler) {
	s.called, s.path, s.handlers = "CREATE", p, len(h)
}
func (s *stubRouter) UPDATE(p string, h ...gateway.Handler) {
	s.called, s.path, s.handlers = "UPDATE", p, len(h)
}
func (s *stubRouter) DELETE(p string, h ...gateway.Handler) {
	s.called, s.path, s.handlers = "DELETE", p, len(h)
}
func (s *stubRouter) STATIC(string)                                {}
func (s *stubRouter) ServeHttp(http.ResponseWriter, *http.Request) {}
func (s *stubRouter) Middleware(...gateway.Handler)                {}

type noopHandler struct{}

func (noopHandler) Handle(gateway.HttpRequester) (any, errors.ErrorModel) {
	return nil, nil
}

func TestRegisterRoutersDispatches(t *testing.T) {
	cases := []struct {
		action gateway.Action
		want   string
	}{
		{gateway.Read, "READ"},
		{gateway.Create, "CREATE"},
		{gateway.Update, "UPDATE"},
		{gateway.Delete, "DELETE"},
	}
	for _, tc := range cases {
		t.Run(string(tc.action), func(t *testing.T) {
			r := &stubRouter{}
			gateway.RegisterRouters(r, "/things", tc.action, noopHandler{}, noopHandler{})
			if r.called != tc.want {
				t.Errorf("dispatched to %q, want %q", r.called, tc.want)
			}
			if r.path != "/things" {
				t.Errorf("path = %q, want /things", r.path)
			}
			if r.handlers != 2 {
				t.Errorf("handlers = %d, want 2", r.handlers)
			}
		})
	}
}

func TestActionConstants(t *testing.T) {
	if gateway.Create != "CREATE" || gateway.Read != "READ" ||
		gateway.Update != "UPDATE" || gateway.Delete != "DELETE" {
		t.Error("action constant values drifted")
	}
}

func TestStatusConstants(t *testing.T) {
	want := map[gateway.Status]string{
		gateway.StatusOK:        "OK",
		gateway.StatusCreated:   "CREATED",
		gateway.StatusNoContent: "NO_CONTENT",
		gateway.StatusBadInput:  "BAD REQUEST",
		gateway.StatusConflict:  "CONFLICT",
	}
	for got, str := range want {
		if string(got) != str {
			t.Errorf("status %q = %q", str, string(got))
		}
	}
}
