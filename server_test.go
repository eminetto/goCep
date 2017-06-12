package main

import (
	"reflect"
	"testing"
	"encoding/json"
	"os"
)

type Cep struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Unidade string `json:"unidade"`
	Ibge string `json:"ibge"`
	Gia string `json:"gia"`
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Remove_Cache_File(id string) {
	os.Remove(getCacheFilename(id))
}

func Test_Get_Cache_Filename(t *testing.T) {
  id := "89201405"
  id_with_dash := "89201-405"

  expect(t, getCacheFilename(id), os.TempDir()+"/cep"+id)
  expect(t, getCacheFilename(id_with_dash), os.TempDir()+"/cep"+id)
}

func Test_Get_Cep(t *testing.T) {
	id := "0000000"
	wrongCep := getCep(id)
	expect(t, "<h2>Bad Request (400)</h2>", wrongCep)

	Remove_Cache_File(id)

	id = "60170150"
	cepJson := getCep(id)
	res := Cep{}
	json.Unmarshal([]byte(cepJson), &res)

	expect(t, "60170-150", res.Cep)
	expect(t, "Rua Vicente Leite", res.Logradouro)
	expect(t, "até 879/880", res.Complemento)
	expect(t, "Meireles", res.Bairro)
	expect(t, "Fortaleza", res.Localidade)
	expect(t, "CE", res.Uf)
	expect(t, "", res.Unidade)
	expect(t, "2304400", res.Ibge)
	expect(t, "", res.Gia)

	Remove_Cache_File(id)
}

func Test_Cache(t *testing.T) {
	id := "89201405"
	getCep(id) // Add to temporary_directory_path/cep89201405

	if _, err := os.Stat(getCacheFilename(id)); err != nil {
		t.Errorf("Cache doesn't work - %v", err)
	}

	Remove_Cache_File(id)
}
