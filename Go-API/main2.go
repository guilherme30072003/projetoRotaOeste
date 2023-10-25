package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"database/sql"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SeuStruct struct {
	Id string `json:"id"`
	Occurrences string `json:"occurrences"`
	MachineLinearTime string `json:"machineLinearTime"`
	Bus string `json:"bus"`
	Time string `json:"time"`
	Color string `json:"color"`
	Severity string `json:"severity"`
	AcknowledgementStatus string `json:"acknowledgementStatus"`
	Ignored string `json:"ignored"`
	Invisible string `json:"invisible"`
	Rel string `json:"rel"`
	Uri string `json:"uri"`
	DurationSeconds string `json:"durationSeconds"`
	EngineHoursReading string `json:"engineHoursReading"`
	SuspectParameterName string `json:"suspectParameterName"`
	FailureModeIndicator string `json:"failureModeIndicator"`
	DefinitionBus string `json:"definition"`
	SourceAddress string `json:"sourceAddress"`
	ThreeLetterAcronym string `json:"threeLetterAcronym"`
	DefinitionId string `json:"definitionId"`
	Description string `json:"description"`
	DefinitionRel string `json:"definitionRel"`
	DefinitionUri string `json:"definitionUri"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func main() {
	// Ler o arquivo JSON
	data, err := ioutil.ReadFile("received.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo JSON:", err)
		return
	}

	// Deserializar o JSON em uma slice de sua struct
	var seuStructs SeuStruct
	err = json.Unmarshal(data, &seuStructs)
	if err != nil {
		fmt.Println("Erro ao deserializar o JSON:", err)
		return
	}

	fmt.Println("Dados deserializados com sucesso:", seuStructs.Id)

	// Defina a string de conexão ao banco de dados
	connString := "server=localhost;user id=sa;password=Mamute1912;database=AlertsMachine"

	// Conecte ao banco de dados
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Verifique a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o banco de dados bem-sucedida!")

	// Supondo que você já tenha conectado ao banco de dados (passos anteriores)

	// Loop sobre seus dados e insira-os no banco de dados
	query := "INSERT INTO receivedAlerts (alertId, occurrences, machineLinearTime, bus, time, color, severity, acknowledgementStatus, ignored, invisible,rel, uri, durationSeconds, engineHoursReading,suspectParameterName, failureModeIndicator, definitionBus, sourceAddress, threeLetterAcronym, definitionId, description, definitionRel, definitionUri, lat, lon) VALUES (@param1, @param2, @param3, @param4, @param5, @param6, @param7, @param8, @param9, @param10, @param11, @param12, @param13, @param14, @param15, @param16, @param25, @param17, @param18, @param19, @param20, @param21, @param22, @param23, @param24)"
	consultaResult, err := db.Exec(query, sql.Named("param1",seuStructs.Id), sql.Named("param2",seuStructs.Occurrences), sql.Named("param3",seuStructs.MachineLinearTime), sql.Named("param4",seuStructs.Bus), sql.Named("param5",seuStructs.Time), sql.Named("param6",seuStructs.Color), sql.Named("param7",seuStructs.Severity), sql.Named("param8",seuStructs.AcknowledgementStatus), sql.Named("param9",seuStructs.Ignored), sql.Named("param10",seuStructs.Invisible), sql.Named("param11",seuStructs.Rel), sql.Named("param12",seuStructs.Uri), sql.Named("param13",seuStructs.DurationSeconds), sql.Named("param14",seuStructs.EngineHoursReading), sql.Named("param15",seuStructs.SuspectParameterName), sql.Named("param16",seuStructs.FailureModeIndicator), sql.Named("param17",seuStructs.SourceAddress), sql.Named("param18",seuStructs.ThreeLetterAcronym), sql.Named("param19",seuStructs.DefinitionId), sql.Named("param20",seuStructs.Description), sql.Named("param21",seuStructs.DefinitionRel), sql.Named("param22",seuStructs.DefinitionUri), sql.Named("param23",seuStructs.Lat), sql.Named("param24", seuStructs.Lon), sql.Named("param25",seuStructs.DefinitionBus))
	fmt.Println("Inserção Concluída!", consultaResult)
	if err != nil {
	    log.Println("Erro ao inserir dados:", err)
	}
	router := gin.Default()
	// Rota para obter todos os dados
	router.GET("/dados", func(c *gin.Context) {
		// Execute a consulta SQL para obter todos os dados
		rows, err := db.Query("SELECT * FROM receivedAlerts")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var results []SeuStruct
		for rows.Next() {
			var result SeuStruct
			var teste string
			err := rows.Scan(&result.Id, &result.Occurrences, &result.MachineLinearTime , &result.Bus , &result.Time, &result.Color, &result.Severity, &result.AcknowledgementStatus , &result.Ignored , &result.Invisible, &result.Rel, &result.Uri, &result.DurationSeconds , &result.EngineHoursReading , &result.SuspectParameterName, &result.FailureModeIndicator, &result.DefinitionBus, &result.SourceAddress, &result.ThreeLetterAcronym , &result.DefinitionId , &result.Description, &result.DefinitionRel,  &result.DefinitionUri,  &result.Lat,  &result.Lon, &teste)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			results = append(results, result)
		}

		c.IndentedJSON(http.StatusOK, results)
	})

	// Rota para filtragem de dados (você precisará implementar isso)

	router.Run("localhost:8080")
}
