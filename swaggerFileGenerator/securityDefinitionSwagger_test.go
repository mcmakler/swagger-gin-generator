package swaggerFileGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicSecurityDefinitionSwagger_ToString(t *testing.T) {
	t.Run("Test: BasicSecurityDefinitionSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return "+errorEmptySecurityTitle.Error()+"error", func(t *testing.T) {
			a := NewBasicSecurityDefinition("")
			_, actual := a.ToString()
			expected := errorEmptySecurityTitle
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewBasicSecurityDefinition("title")
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := "\ntitle:" + basicSecurityString
			assert.Equal(t, expected, actual)
		})
	})
}

func TestApiKeySecurityDefinitionSwagger_ToString(t *testing.T) {
	t.Run("Test: ApiKeySecurityDefinitionSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return "+errorEmptySecurityTitle.Error()+"error", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("", "", true)
			_, actual := a.ToString()
			expected := errorEmptySecurityTitle
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return "+errorEmptyApiKeySecurityName.Error()+"error", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("title", "", true)
			_, actual := a.ToString()
			expected := errorEmptyApiKeySecurityName
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("title", "name", true)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := "\ntitle:" + apiKeySecurityString +
				apiKeySecurityNameString + "name" +
				apiKeySecurityInHeaderString
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewApiKeySecurityDefinition("title", "name", false)
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := "\ntitle:" + apiKeySecurityString +
				apiKeySecurityNameString + "name" +
				apiKeySecurityInQueryString
			assert.Equal(t, expected, actual)
		})
	})
}

func TestOauth2SecurityDefinitionSwagger_ToString(t *testing.T) {
	t.Run("Test: Oauth2SecurityDefinitionSwagger.ToString()", func(t *testing.T) {
		t.Run("Should: return "+errorEmptySecurityTitle.Error()+"error", func(t *testing.T) {
			a := NewOauth2PasswordSecurityDefinition("", "")
			_, actual := a.ToString()
			expected := errorEmptySecurityTitle
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return "+errorEmptyOauth2SecurityUrl.Error()+"error", func(t *testing.T) {
			a := NewOauth2ApplicationSecurityDefinition("title", "")
			_, actual := a.ToString()
			expected := errorEmptyOauth2SecurityUrl
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: return "+errorEmptyOauth2SecurityUrl.Error()+"error", func(t *testing.T) {
			a := NewOauth2ImplicitSecurityDefinition("title", "")
			_, actual := a.ToString()
			expected := errorEmptyOauth2SecurityUrl
			assert.Equal(t, expected, actual)
		})

		t.Run("Should: no errors", func(t *testing.T) {
			a := NewOauth2AccessCodeSecurityDefinition("title", "authUrl", "tokenUrl")
			actual, err := a.ToString()
			assert.NoError(t, err)

			expected := "\ntitle:" + oauth2SecurityString +
				oauth2SecurityFlowString + "accessCode" +
				oauth2SecurityAuthUrlString + "authUrl" +
				oauth2SecurityTokenUrlString + "tokenUrl"
			assert.Equal(t, expected, actual)
		})
	})
}
