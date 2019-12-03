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

type SecurityDefinitionSwagg interface {
	ToString() (string, error)
}

type basicSecurityDefinitionSwagg struct {
	title string
}

func NewBasicSecurityDefinition(title string) SecurityDefinitionSwagg {
	return &basicSecurityDefinitionSwagg{
		title: title,
	}
}
func (b *basicSecurityDefinitionSwagg) ToString() (string, error) {
	if b.title == "" {
		return "", errorEmptySecurityTitle
	}
	res := "\n" + b.title + ":" + basicSecurityString
	return res, nil
}

type apiKeySecurityDefinitionSwagg struct {
	title    string
	name     string
	inHeader bool
}

func NewApiKeySecurityDefinition(title, name string, inHeader bool) SecurityDefinitionSwagg {
	return &apiKeySecurityDefinitionSwagg{
		title:    title,
		name:     name,
		inHeader: inHeader,
	}
}

func (b *apiKeySecurityDefinitionSwagg) ToString() (string, error) {
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

type oauth2SecurityDefinitionSwagg struct {
	title    string
	flow     string
	tokenURL string
	authURL  string
}

func NewOauth2ImplicitSecurityDefinition(title, authorizationUrl string) SecurityDefinitionSwagg {
	return &oauth2SecurityDefinitionSwagg{
		flow:     "implicit",
		title:    title,
		tokenURL: "",
		authURL:  authorizationUrl,
	}
}

func NewOauth2PasswordSecurityDefinition(title, tokenURL string) SecurityDefinitionSwagg {
	return &oauth2SecurityDefinitionSwagg{
		flow:     "password",
		title:    title,
		tokenURL: tokenURL,
		authURL:  "",
	}
}

func NewOauth2ApplicationSecurityDefinition(title, tokenURL string) SecurityDefinitionSwagg {
	return &oauth2SecurityDefinitionSwagg{
		flow:     "application",
		title:    title,
		tokenURL: tokenURL,
		authURL:  "",
	}
}

func NewOauth2AccessCodeSecurityDefinition(title, authorizationUrl, tokenURL string) SecurityDefinitionSwagg {
	return &oauth2SecurityDefinitionSwagg{
		flow:     "accessCode",
		title:    title,
		tokenURL: tokenURL,
		authURL:  authorizationUrl,
	}
}

func (b *oauth2SecurityDefinitionSwagg) ToString() (string, error) {
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
