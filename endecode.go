package HttpDo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
)

func (h *HttpDo) MarshalJSON(v any) *HttpDo {
	if h.httpErr != nil {
		return h
	}
	bodybytes, err := json.Marshal(v)
	if err != nil {
		h.httpErr = err
		return h
	}
	h.requestBody = bytes.NewReader(bodybytes)
	return h
}

func (h *HttpDo) MarshalXML(v any) *HttpDo {
	if h.httpErr != nil {
		return h
	}
	bodybytes, err := xml.Marshal(v)
	if err != nil {
		h.httpErr = err
		return h
	}
	h.requestBody = bytes.NewReader(bodybytes)
	return h
}

func (h *HttpDo) UnmarshalJSONResponse(v any) error {
	bodybytes, err := io.ReadAll(h.ResponseBody)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodybytes, v)
	if err != nil {
		return err
	}
	return nil
}

func (h *HttpDo) UnmarshalXMLResponse(v any) (any, error) {
	bodybytes, err := io.ReadAll(h.ResponseBody)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(bodybytes, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
