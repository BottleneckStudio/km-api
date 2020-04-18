package auth

import (
	"io/ioutil"
	"testing"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("empty key set", func(t *testing.T) {
		p := New(nil)

		assert.Nil(t, p)
	})

	t.Run("Not empty keySet", func(t *testing.T) {
		p := New(&jwk.Set{})

		assert.NotNil(t, p)
	})
}

func TestParse(t *testing.T) {
	t.Run("parse error", func(t *testing.T) {
		tokenString := ""
		p := New(initializeKeySets())

		_, err := p.Parse(tokenString)

		assert.NotNil(t, err)
	})

	t.Run("parse error", func(t *testing.T) {
		tokenString := ""
		p := New(initializeKeySets())

		_, err := p.Parse(tokenString)

		assert.NotNil(t, err)
	})

	t.Run("valid token", func(t *testing.T) {
		tokenString := ""

		p := New(initializeKeySets())
		_, err := p.Parse(tokenString)

		assert.NotNil(t, err)

	})

}

func initializeKeySets() *jwk.Set {
	content, err := ioutil.ReadFile("jwks.json")
	if err != nil {
		return nil
	}
	set, err := jwk.ParseBytes(content)
	if err != nil {
		return nil
	}
	return set
}
