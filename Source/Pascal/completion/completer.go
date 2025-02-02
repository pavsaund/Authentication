package completion

import (
	"dolittle.io/pascal/openid"
	"dolittle.io/pascal/sessions"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Completer interface {
	Complete(response *Response, session *sessions.Session) (*oauth2.Token, error)
}

func NewCompleter(validator Validator, exchanger openid.TokenExchanger, logger *zap.Logger) Completer {
	return &completer{
		validator: validator,
		exchanger: exchanger,
		logger:    logger,
	}
}

type completer struct {
	validator Validator
	exchanger openid.TokenExchanger
	logger    *zap.Logger
}

func (c *completer) Complete(response *Response, session *sessions.Session) (*oauth2.Token, error) {
	if sesionIsValid, err := c.validator.Validate(response, session); !sesionIsValid {
		c.logger.Error("session was not valid", zap.Error(err))
		return nil, ErrSessionDoesNotMatchProviderCallback
	}

	token, err := c.exchanger.Exchange(response.Code)
	if err != nil {
		c.logger.Error("could not exhchange code for token", zap.Error(err))
		return nil, err
	}

	return token, nil
}
