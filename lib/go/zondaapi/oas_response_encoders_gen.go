// Code generated by ogen, DO NOT EDIT.

package zondaapi

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

func encodeCheckHealthResponse(response *CheckHealthOK, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}
