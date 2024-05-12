package tools

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rs/zerolog/log"
)

func GraphError(err error) {
	if err != nil {
		switch err := err.(type) {
		case *odataerrors.ODataError:
			log.Error().Any("Graph Error", err.GetErrorEscaped())
			if terr := err.GetErrorEscaped(); terr != nil {
				log.Error().Str("code", *terr.GetCode())
				log.Error().Str("msg", *terr.GetMessage())
			}
		default:
			log.Error().Err(err).Msg("")
		}
	}
}