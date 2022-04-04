// POST-create запрос для работы с сервисом интеграции Directum RX

package odatapost

import (
	"bytes"
	"encoding/base64"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostCreate(podata string, cr chan string) {

	payload, _ := json.Marshal(struct {
		ColumnOne string `json:"Name"`
	}{
		ColumnOne: podata,
	})

	msg := `Administrator:11111`
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("Данные авторизации base64: ", encoded)

	req, err := http.NewRequest("POST", "http://localhost/DrxIntegration/odata/ISimpleDocuments", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+encoded)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	fmt.Printf(" %v ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	//log.Println(string([]byte(body)))
	cr <- string([]byte(body)) // передача результата в main-программу
}
