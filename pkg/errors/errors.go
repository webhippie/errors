package errors

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/webhippie/errors/pkg/config"
	"gopkg.in/yaml.v3"
)

var (
	defaultErrors = List{
		400: "The server cannot or will not process the request.",
		401: "You are not authorized to request this resource.",
		403: "The server is refusing to respond to your request.",
		404: "The page you are looking for was not found.",
		405: "The server doesn't accept your request method.",
		406: "The headers can't be accepted by the server.",
		407: "The client must first authenticate itself with the proxy.",
		408: "The server timed out waiting for the request.",
		409: "The request could not be processed because of conflict.",
		410: "The resource is no longer available and will not be available again.",
		411: "The request did not specify the length of its content.",
		412: "The server does not meet one of the preconditions.",
		413: "The request is larger than the server is willing or able to process.",
		414: "The URI provided was too long for the server to process.",
		415: "The request has a media type which the server does not support.",
		416: "The requested range can't be satisfied.",
		417: "The server cannot meet the requirements of the Expect header.",
		422: "The request was well-formed but was unable to be followed.",
		423: "The resource that is being accessed is locked.",
		424: "The request failed due to failure of a previous request.",
		426: "The client should switch to a different protocol.",
		428: "The origin server requires the request to be conditional.",
		429: "The user has sent too many requests in a given amount of time.",
		431: "The server is unwilling to process the request.",
		451: "The request have been blocked because of legal reasons.",
		500: "I messed it up, but this is not your fault.",
		501: "The server does not recognize the request method.",
		502: "The server received an invalid response from upstream.",
		503: "The server is currently unavailable, this is a temporary state.",
		504: "The server did not receive a timely response from upstream.",
		505: "The server does not support the HTTP protocol version.",
		506: "Transparent content negotiation results in a circular reference.",
		507: "The server is unable to store the result of the request.",
		508: "The server detected an infinite loop while processing the request.",
		510: "Extensions to the request are required for the server to fulfil it.",
		511: "The client needs to authenticate to gain network access.",
	}
)

// List defines the list of available errors.
type List map[int]string

// Load initializes the errors list.
func Load(cfg *config.Config) List {
	if cfg.Server.Errors != "" {
		result := List{}

		if _, err := os.Stat(cfg.Server.Errors); os.IsNotExist(err) {
			log.Error().
				Err(err).
				Str("file", cfg.Server.Errors).
				Msg("Failed to find custom errors")

			return defaultErrors
		}

		content, err := os.ReadFile(cfg.Server.Errors)

		if err != nil {
			log.Error().
				Err(err).
				Str("file", cfg.Server.Errors).
				Msg("Failed to read custom errors")

			return defaultErrors
		}

		switch ext := filepath.Ext(cfg.Server.Errors); ext {
		case ".json":
			if err := json.Unmarshal(content, &result); err != nil {
				log.Error().
					Err(err).
					Str("file", cfg.Server.Errors).
					Msg("Failed to parse custom errors")

				return defaultErrors
			}
		case ".yaml":
			if err := yaml.Unmarshal(content, &result); err != nil {
				log.Error().
					Err(err).
					Str("file", cfg.Server.Errors).
					Msg("Failed to parse custom errors")

				return defaultErrors
			}
		default:
			log.Error().
				Str("file", cfg.Server.Errors).
				Msg("Unknown file extension for custom errors")

			return defaultErrors
		}

		return result
	}

	return defaultErrors
}
