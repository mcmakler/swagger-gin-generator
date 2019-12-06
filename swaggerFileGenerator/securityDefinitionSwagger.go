package swaggerFileGenerator

import "errors"

const (
	basicSecurityString = "\n  type: basic"

	apiKeySecurityString         = "\n  type: apiKey"
	apiKeySecurityNameString     = "\n  name: "
	apiKeySecurityInHeaderString = "\n  in: header"
	apiKeySecurityInQueryString  = "\n  in: query"

	oauth2SecurityString         = "\n  type: oauth2"
	oauth2SecurityFlowString     = "\n  flow: "
	oauth2SecurityAuthUrlString  = "\n  authorizationUrl: "
	oauth2SecurityTokenUrlString = "\n  tokenUrl: "
)

var (
	errorEmptySecurityTitle      = errors.New("EMPTY_SECURITY_NAME")
	errorEmptyApiKeySecurityName = errors.New("EMPTY_API_KEY_NAME")
	errorEmptyOauth2SecurityUrl  = errors.New("EMPTY_OAUTH2_URL")
)

type SecurityDefinitionSwagger interface {
	ToString() (string, error)
}

type basicSecurityDefinitionSwagger struct {
	title string
}

func NewBasicSecurityDefinition(title string) SecurityDefinitionSwagger {
	return &basicSecurityDefinitionSwagger{
		title: title,
	}
}

func (b *basicSecurityDefinitionSwagger) ToString() (string, error) {
	if b.title == "" {
		return "", errorEmptySecurityTitle
	}
	res := "\n" + b.title + ":" + basicSecurityString
	return res, nil
}

type apiKeySecurityDefinitionSwagger struct {
	title    string
	name     string
	inHeader bool
}

func NewApiKeySecurityDefinition(title, name string, inHeader bool) SecurityDefinitionSwagger {
	return &apiKeySecurityDefinitionSwagger{
		title:    title,
		name:     name,
		inHeader: inHeader,
	}
}

func (b *apiKeySecurityDefinitionSwagger) ToString() (string, error) {
	if b.title == "" {
		return "", errorEmptySecurityTitle
	}
	if b.name == "" {
		return "", errorEmptyApiKeySecurityName
	}
	res := "\n" + b.title + ":" + apiKeySecurityString +
		apiKeySecurityNameString + b.name
	if b.inHeader {
		res += apiKeySecurityInHeaderString
		return res, nil
	}
	res += apiKeySecurityInQueryString
	return res, nil
}

type oauth2SecurityDefinitionSwagger struct {
	title    string
	flow     string
	tokenURL string
	authURL  string
}

func NewOauth2ImplicitSecurityDefinition(title, authorizationUrl string) SecurityDefinitionSwagger {
	return &oauth2SecurityDefinitionSwagger{
		flow:     "implicit",
		title:    title,
		tokenURL: "",
		authURL:  authorizationUrl,
	}
}

func NewOauth2PasswordSecurityDefinition(title, tokenURL string) SecurityDefinitionSwagger {
	return &oauth2SecurityDefinitionSwagger{
		flow:     "password",
		title:    title,
		tokenURL: tokenURL,
		authURL:  "",
	}
}

func NewOauth2ApplicationSecurityDefinition(title, tokenURL string) SecurityDefinitionSwagger {
	return &oauth2SecurityDefinitionSwagger{
		flow:     "application",
		title:    title,
		tokenURL: tokenURL,
		authURL:  "",
	}
}

func NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL string) SecurityDefinitionSwagger {
	return &oauth2SecurityDefinitionSwagger{
		flow:     "accessCode",
		title:    title,
		tokenURL: tokenURL,
		authURL:  authorizationUrl,
	}
}

func (b *oauth2SecurityDefinitionSwagger) ToString() (string, error) {
	if b.title == "" {
		return "", errorEmptySecurityTitle
	}
	if b.tokenURL == "" && b.authURL == "" {
		return "", errorEmptyOauth2SecurityUrl
	}
	res := "\n" + b.title + ":" + oauth2SecurityString +
		oauth2SecurityFlowString + b.flow
	if b.authURL != "" {
		res += oauth2SecurityAuthUrlString + b.authURL
	}
	if b.tokenURL != "" {
		res += oauth2SecurityTokenUrlString + b.tokenURL
	}
	return res, nil
}
