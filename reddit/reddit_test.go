package reddit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mock reddit news
type RDServer struct{}

func (srv *RDServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check method and path
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/r/golang/.rss" {
		w.Header().Add("Content-Type", "application/xml")
		dat, _ := ioutil.ReadFile("testdata/topstories.xml")
		w.Write([]byte(dat))
		return
	}

	if r.URL.Path == "/r/golang/comments/bnu47l/.rss" {
		w.Header().Add("Content-Type", "application/xml")
		dat, _ := ioutil.ReadFile("testdata/bnu47l.xml")
		w.Write([]byte(dat))
		return
	}

	// nothing match
	w.WriteHeader(http.StatusNotFound)
	return
}

func TestTopArticles(t *testing.T) {
	srv := RDServer{}
	rdsrv := httptest.NewServer(&srv)
	defer rdsrv.Close()

	apiBase := fmt.Sprintf("%v/r/golang", rdsrv.URL)
	rdr := NewReader(apiBase)

	ids, _ := rdr.TopArticles(5)
	target := []string{"bnu47l", "bnvik4", "bnyrzp", "bnmcwk", "bnzcxy"}
	if len(ids) != len(target) {
		t.Errorf("got: %#v; expect: %#v", ids, target)
	}

	for i, v := range target {
		if v != ids[i] {
			t.Errorf("got: %#v; expect: %#v", v, ids[i])
		}
	}

}

func TestGetArticle(t *testing.T) {
	srv := RDServer{}
	rdsrv := httptest.NewServer(&srv)
	defer rdsrv.Close()

	apiBase := fmt.Sprintf("%v/r/golang", rdsrv.URL)
	rdr := NewReader(apiBase)

	article, _ := rdr.GetArticle("bnu47l")
	if id := article.GetID(); id != "bnu47l" {
		t.Errorf("got: %#v; expect: %#v", id, "bnu47l")
	}
	if title := article.GetTitle(); title != "High Performance DICOM Medical Image Parser in Golang" {
		t.Errorf("got: %#v; expect: %#v", title, "High Performance DICOM Medical Image Parser in Golang")
	}
	if auth := article.GetAuthor(); auth != "/u/suyashkumar" {
		t.Errorf("got: %#v; expect: %#v", auth, "/u/suyashkumar")
	}
	if ts := article.GetTimestamp(); ts != 1557694857 {
		t.Errorf("got: %#v; expect: %#v", ts, 1557694857)
	}
}
