package main

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("handle get", func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/post/", handleRequest(&FakePost{}))
		writer := httptest.NewRecorder()

		request, _ := http.NewRequest("GET", "/post/5", nil)
		mux.ServeHTTP(writer, request)

		if writer.Code != 200 {
			GinkgoT().Errorf("Response code is %v", writer.Code)
		}

		var post Post
		json.Unmarshal(writer.Body.Bytes(), &post)
		if post.Id != 5 {
			GinkgoT().Error("Cannnot retrieve JSON post")
		}
	})
	It("handle put", func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/post/", handleRequest(&FakePost{}))
		writer := httptest.NewRecorder()

		json := strings.NewReader(`{"content":"Updated post", "author":"Sau SHeong"`)
		request, _ := http.NewRequest("PUT", "/post/1", json)
		mux.ServeHTTP(writer, request)

		if writer.Code != 200 {
			GinkgoT().Errorf("Response code is %#v", writer.Code)
		}
	})
})

func setUp() {

}
