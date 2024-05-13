package tools

import (
	"net/http"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rs/zerolog/log"
)

func GraphError(err error, w http.ResponseWriter) {
	if err != nil {
		switch err := err.(type) {
		case *odataerrors.ODataError:
			log.Error().Any("Graph Error", err.GetErrorEscaped()).Msg("")
			if terr := err.GetErrorEscaped(); terr != nil {
				log.Error().Str("code", *terr.GetCode()).Msg("")
				log.Error().Str("msg", *terr.GetMessage()).Msg("")
			}
		default:
			log.Error().Err(err).Msg("")
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
