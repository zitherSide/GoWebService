package main

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finishied test -", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finishing Post Test Suite")
}

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TestGetPost(c *C) {
	request, _ := http.NewRequest("GET", "/post/5", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)

	json.Unmarshal(s.writer.Body.Bytes(), &s.post)
	c.Check(s.post.Id, Equals, 5)
}

func (s *PostTestSuite) TestPutPost(c *C) {
	json := strings.NewReader(`{"content":"Updated post", "author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/5", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 5)
	c.Check(s.post.Content, Equals, "Updated post")
}

func TestHandleGet(t *testing.T) {
	// request, _ := http.NewRequest("GET", "/post/5", nil)
	// mux.ServeHTTP(writer, request)

	// if writer.Code != 200 {
	// 	t.Errorf("Response code is %v", writer.Code)
	// }

	// var post Post
	// json.Unmarshal(writer.Body.Bytes(), &post)
	// if post.Id != 5 {
	// 	t.Error("Cannnot retrieve JSON post")
	// }
}

func TestHandlePut(t *testing.T) {
	// json := strings.NewReader(`{"content":"Updated post", "author":"Sau SHeong"`)
	// request, _ := http.NewRequest("PUT", "/post/1", json)
	// mux.ServeHTTP(writer, request)

	// if writer.Code != 200 {
	// 	t.Errorf("Response code is %#v", writer.Code)
	// }
}
