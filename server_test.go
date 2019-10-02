package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)



func removeCacheFile(t *testing.T, id string) {
	assert.Nil(t, os.Remove(getCacheFilename(id)))
}

func Test_getCacheFilename(t *testing.T) {
  id := "89201405"
  idWithDash := "89201-405"

  assert.Equal(t, getCacheFilename(id), os.TempDir()+"/cep"+id)
  assert.Equal(t, getCacheFilename(idWithDash), os.TempDir()+"/cep"+id)
}

func Test_getCep(t *testing.T) {
	t.Run("invalid cep", func(t *testing.T) {
		id := "0000000"
		wrongCep, err := getCep(id)
		assert.Equal(t, "", wrongCep)
		assert.Error(t, err)
	})
	t.Run("valid cep", func(t *testing.T) {

		id := "60170150"
		cepJson, err := getCep(id)
		assert.Nil(t, err)
		res := Cep{}
		assert.Nil(t, json.Unmarshal([]byte(cepJson), &res))

		assert.Equal(t, "60170-150", res.Cep)
		assert.Equal(t, "Rua Vicente Leite", res.Logradouro)
		assert.Equal(t, "até 879/880", res.Complemento)
		assert.Equal(t, "Meireles", res.Bairro)
		assert.Equal(t, "Fortaleza", res.Localidade)
		assert.Equal(t, "CE", res.Uf)
		assert.Equal(t, "", res.Unidade)
		assert.Equal(t, "2304400", res.Ibge)
		assert.Equal(t, "", res.Gia)

		removeCacheFile(t, id)
	})
}

func Test_Cache(t *testing.T) {
	id := "89201405"
	_, err := getCep(id) // Add to temporary_directory_path/cep89201405
	assert.Nil(t, err)
	_, err = os.Stat(getCacheFilename(id))
	assert.Nil(t, err)
	removeCacheFile(t,id)
}
