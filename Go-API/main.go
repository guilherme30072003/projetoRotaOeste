package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type Alerta struct{
	IdInterfaceNegociacao	int `json:"idInterfaceNegociacao"`
	Notificacoes []string `json:"notificacoes"`
	EncaminhamentoAtivo []string `json:"encaminhamentoAtivo"`
	Veiculo []string `json:"veiculo"`
	Evento []string `json:"evento"`
}

var alertas = []Alerta{
	{IdInterfaceNegociacao: 0, 
		Notificacoes: []string{"key:", "ENCAMINHAMENTO","value:", "1"}, 
		EncaminhamentoAtivo: []string{"idEncaminhamento:", "1",
		"idEvento:", "1",
		"idUsuarioEncaminhador:", "1",
		"idUsuario:", "1",
		"motivo:", "blablabla",
		"idEmpresa:", "1",
		"encaminhamentoAtivo:", "true",
		"dataInclusao:", "1",
		"dataAlteracao:", "2",
		"usuarioInc:", "blablabla2",
		"usuarioAlt:", "blablabla3",
		"origemRetorno:", "1"},
		Veiculo: []string{"idVeiculo:","789456123",
			"idCliente:", "1",
			"idEmpresa:", "1",
			"placa:", "BLA-1234",
			"cor:", "laranja",
			"chassi:", "123456789",
			"modelo:", "Teste"},
		Evento: []string{"idCliente:", "1",
			"idEmpresa:", "1",
			"idGrupoEvento:", "teste1",
			"idTipoEvento:", "teste2",
			"formaContato:", "teste3",
			"midia:", "teste4",
			"temperatura:", "BOLSAO",
			"marca:", "teste5",
			"modelo:", "teste6",
			"idAgente:", "1",
			"observacao:", "teste7",
			"dataProximaAcao:", "3",},
	},
	{IdInterfaceNegociacao: 1, 
		Notificacoes: []string{"key:", "ENCAMINHAMENTO","value:", "2"}, 
		EncaminhamentoAtivo: []string{"idEncaminhamento:", "2",
		"idEvento:", "2",
		"idUsuarioEncaminhador:", "2",
		"idUsuario:", "2",
		"motivo:", "lorem ipsum",
		"idEmpresa:", "2",
		"encaminhamentoAtivo:", "true",
		"dataInclusao:", "2",
		"dataAlteracao:", "3",
		"usuarioInc:", "lorem ipsum",
		"usuarioAlt:", "lorem ipsum",
		"origemRetorno:", "2"},
		Veiculo: []string{"idVeiculo:","456987",
			"idCliente:", "2",
			"idEmpresa:", "2",
			"placa:", "LOREM-1234",
			"cor:", "amarelo",
			"chassi:", "654123789",
			"modelo:", "Teste2"},
		Evento: []string{"idCliente:", "2",
			"idEmpresa:", "2",
			"idGrupoEvento:", "lorem1",
			"idTipoEvento:", "lorem2",
			"formaContato:", "lorem3",
			"midia:", "lorem4",
			"temperatura:", "BOLSAO",
			"marca:", "lorem5",
			"modelo:", "lorem6",
			"idAgente:", "2",
			"observacao:", "lorem7",
			"dataProximaAcao:", "4",},
	},
	{IdInterfaceNegociacao: 2, 
		Notificacoes: []string{"key:", "ENCAMINHAMENTO","value:", "3"}, 
		EncaminhamentoAtivo: []string{"idEncaminhamento:", "3",
		"idEvento:", "3",
		"idUsuarioEncaminhador:", "3",
		"idUsuario:", "3",
		"motivo:", "mimimi",
		"idEmpresa:", "3",
		"encaminhamentoAtivo:", "true",
		"dataInclusao:", "3",
		"dataAlteracao:", "4",
		"usuarioInc:", "mimimi1",
		"usuarioAlt:", "mimimi2",
		"origemRetorno:", "13"},
		Veiculo: []string{"idVeiculo:","963852",
			"idCliente:", "3",
			"idEmpresa:", "3",
			"placa:", "MIMIMI-1234",
			"cor:", "verde",
			"chassi:", "741258",
			"modelo:", "Teste3"},
		Evento: []string{"idCliente:", "3",
			"idEmpresa:", "3",
			"idGrupoEvento:", "mimimi1",
			"idTipoEvento:", "mimimi2",
			"formaContato:", "mimimi3",
			"midia:", "mimimi4",
			"temperatura:", "BOLSAO",
			"marca:", "mimimi5",
			"modelo:", "mimimi6",
			"idAgente:", "3",
			"observacao:", "mimimi7",
			"dataProximaAcao:", "4",},
	},
}

func getAlertas(context *gin.Context){
	context.IndentedJSON(http.StatusOK, alertas)
}

func main(){
	router := gin.Default()
	router.GET("/alertas", getAlertas)
	router.Run("localhost:8080")
}