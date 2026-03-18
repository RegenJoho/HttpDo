package HttpDo

// Sets the HTTP Request Method.
func (h *HttpDo) Method(method string) *HttpDo {
	if h.httpErr != nil {
		return h
	}
	h.method = method
	return h
}

// Sets the HTTP Request Method to GET.
func (h *HttpDo) GET() *HttpDo {
	h.Method("GET")
	return h
}

// Sets the HTTP Request Method to HEAD.
func (h *HttpDo) HEAD() *HttpDo {
	h.Method("HEAD")
	return h
}

// Sets the HTTP Request Method to OPTIONS.
func (h *HttpDo) OPTIONS() *HttpDo {
	h.Method("OPTIONS")
	return h
}

// Sets the HTTP Request Method to TRACE.
func (h *HttpDo) TRACE() *HttpDo {
	h.Method("TRACE")
	return h
}

// Sets the HTTP Request Method to PUT.
func (h *HttpDo) PUT() *HttpDo {
	h.Method("PUT")
	return h
}

// Sets the HTTP Request Method to DELETE
func (h *HttpDo) DELETE() *HttpDo {
	h.Method("DELETE")
	return h
}

// Sets the HTTP Request Method to POST.
func (h *HttpDo) POST() *HttpDo {
	h.Method("POST")
	return h
}

// Sets the HTTP Request Method to PATCH.
func (h *HttpDo) PATCH() *HttpDo {
	h.Method("PATCH")
	return h
}

// Sets the HTTP Request Method to CONNECT.
func (h *HttpDo) CONNECT() *HttpDo {
	h.Method("CONNECT")
	return h
}
