package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

type TestSuite struct {
	suite.Suite
}

func TestSuiteTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) parse(s string) *html.Node {
	a, err := html.Parse(strings.NewReader(s))
	if err != nil {
		ts.T().Fatalf("failed to parse: %v", err)
	}
	a = a.FirstChild.FirstChild.NextSibling.FirstChild
	return a
}

func (ts *TestSuite) TestExtractHref() {
	cases := []struct {
		name string
		a    string
		href string
	}{
		{
			name: "valid a element",
			a:    `<a href="/login">Login</a> `,
			href: "/login",
		},
		{
			name: "missing a element",
			a:    `<a>Login</a> `,
			href: "",
		},
	}

	for _, c := range cases {
		ts.Run(c.name, func() {
			a := ts.parse(c.a)
			href := extractHref(a)
			assert.Equal(ts.T(), href, c.href)
		})

	}
}

func (ts *TestSuite) TestExtractText() {
	cases := []struct {
		name string
		a    string
		text string
	}{
		{
			name: "valid a element",
			a:    `<a href="/login">Login</a> `,
			text: "Login",
		},
		{
			name: "valid nested element",
			a:    `<a href="/login">Login as <strong>Admin</strong></a> `,
			text: "Login as Admin",
		},
		{
			name: "missing a element",
			a:    `<a href="/login"></a>`,
			text: "",
		},
	}

	for _, c := range cases {
		ts.Run(c.name, func() {
			a := ts.parse(c.a)
			text := extractText(a)
			assert.Equal(ts.T(), text, c.text)
		})

	}
}
