package cproxy

import "net/http"

type defaultFilter struct{}

func newFilter() *defaultFilter { return &defaultFilter{} }

func (this *defaultFilter) IsAuthorized(*http.Request) bool { return true }
