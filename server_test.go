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

func Test_Get_Cep(t *testing.T) {
	wrongCep := getCep("0000000")
	expect(t, "<h2>Bad Request (400)</h2>", wrongCep)

	cepJson := getCep("60170150")
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
}

func Test_Cache(t *testing.T) {
	id := "89201405"
	getCep(id) // Add to /tmp/cep89201405

	if _, err := os.Stat("/tmp/cep" + id); err != nil {
		t.Errorf("Cache doesn't work - %v", err)
	}
}