package swaggerFileGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicSecurityDefinitionSwagg_ToString(t *testing.T) {
	t.Run("Test: BasicSecurityDefinitionSwagg.ToString()", func(t *testing.T) {
		t.Run("Should: return "+errorEmptySecurityTitle.Error()+"error", func(t *testing.T) {
			a := NewBasicSecurityDefinition("")
			expexted := errorEmptySecurityTitle
			_, actual := a.ToString()
			assert.Equal(t, expexted, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewBasicSecurityDefinition("title")
			expexted := "\ntitle:" + basicSecurityString
			actual, _ := a.ToString()
			assert.Equal(t, expexted, actual)
		})
	})
}

func TestApiKeySecurityDefinitionSwagg_ToString(t *testing.T) {
	t.Run("Test: ApiKeySecurityDefinitionSwagg.ToString()", func(t *testing.T) {
		t.Run("Should: return "+errorEmptySecurityTitle.Error()+"error", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("", "", true)
			expexted := errorEmptySecurityTitle
			_, actual := a.ToString()
			assert.Equal(t, expexted, actual)
		})

		t.Run("Should: return "+errorEmptyApiKeySecurityName.Error()+"error", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("title", "", true)
			expexted := errorEmptyApiKeySecurityName
			_, actual := a.ToString()
			assert.Equal(t, expexted, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("title", "name", true)
			expexted := "\ntitle:" + apiKeySecurityString +
				apiKeySecurityNameString + "name" +
				apiKeySecurityInHeaderString
			actual, _ := a.ToString()
			assert.Equal(t, expexted, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("title", "name", false)
			expexted := "\ntitle:" + apiKeySecurityString +
				apiKeySecurityNameString + "name" +
				apiKeySecurityInQueryString
			actual, _ := a.ToString()
			assert.Equal(t, expexted, actual)
		})
	})
}

func TestOauth2SecurityDefinitionSwagg_ToString(t *testing.T) {
	t.Run("Test: Oauth2SecurityDefinitionSwagg.ToString()", func(t *testing.T) {
		t.Run("Should: return "+errorEmptySecurityTitle.Error()+"error", func(t *testing.T) {
			a := NewOauth2PasswordSecurityDefinition("", "")
			expexted := errorEmptySecurityTitle
			_, actual := a.ToString()
			assert.Equal(t, expexted, actual)
		})

		t.Run("Should: return "+errorEmptyOauth2SecurityUrl.Error()+"error", func(t *testing.T) {
			a := NewOauth2ApplicationSecurityDefinition("title", "")
			expexted := errorEmptyOauth2SecurityUrl
			_, actual := a.ToString()
			assert.Equal(t, expexted, actual)
		})

		t.Run("Should: return "+errorEmptyOauth2SecurityUrl.Error()+"error", func(t *testing.T) {
			a := NewOauth2ImplicitSecurityDefinition("title", "")
			expexted := errorEmptyOauth2SecurityUrl
			_, actual := a.ToString()
			assert.Equal(t, expexted, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewOauth2AccessCodeSecurityDefinition("title", "authUrl", "tokenUrl")
			expexted := "\ntitle:" + oauth2SecurityString +
				oauth2SecurityFlowString + "accessCode" +
				oauth2SecurityAuthUrlString + "authUrl" +
				oauth2SecurityTokenUrlString + "tokenUrl"
			actual, _ := a.ToString()
			assert.Equal(t, expexted, actual)
		})
	})
}
