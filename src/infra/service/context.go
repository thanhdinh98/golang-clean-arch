package service

import (
	"errors"
	"togo/src"

	"github.com/mitchellh/mapstructure"
)

type Context struct {
	token      string
	tokenData  *src.TokenData
	jwtService src.IJWTService
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (this *Context) GetTokenData() *src.TokenData {
	return this.tokenData
}

func (this *Context) HasPermission(scopes []string) (bool, error) {
	data, err := this.jwtService.VerifyToken(this.token)
	if err != nil {
		return false, err
	}

	for _, scope := range scopes {
		if !contains(data.Permissions, scope) {
			return false, nil
		}
	}

	this.tokenData = data

	return true, nil
}

func (this *Context) LoadContext(data interface{}) error {

	var header struct {
		AccessToken []string `mapstructure:"access_token"`
	}

	if err := mapstructure.Decode(data, &header); err != nil {
		return err
	}

	if len(header.AccessToken) == 0 {
		return errors.New("TOKEN_NOT_PROVIED")
	}

	this.token = header.AccessToken[0]
	return nil
}

func NewServiceContext() src.IContextService {
	return &Context{
		jwtService: NewJWTService(),
	}
}
