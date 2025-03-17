package mercadopago

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	configMP "github.com/mercadopago/sdk-go/pkg/config"

	//"github.com/mercadopago/sdk-go/pkg/merchantorder"

	//"github.com/mercadopago/sdk-go/pkg/merchantorder"
	"github.com/mercadopago/sdk-go/pkg/preference"

	"lanchonete/bootstrap"
	"github.com/tidwall/gjson"
)

func WebhookHandler(c *gin.Context) {
	access_token := bootstrap.NewEnv().AccessToken
	paymentID := c.Query("id")

	if paymentID == "" {
		println("Payment ID is empty")
	}

	url := fmt.Sprintf("https://api.mercadopago.com/v1/payments/%s", paymentID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check for non-200 responses
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Request failed with status: %d - %s", resp.StatusCode, string(body))

	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return
	}
	//body := string(io.ReadAll(resp.Body))
	status := gjson.Get(string(body), "status").String()
	fmt.Println("Status:", status)

	// Step 2: Extract the "id" field inside the "items" array
	itemID := gjson.Get(string(body), "additional_info.items.0.id").String()
	fmt.Println("Item ID:", itemID)

	//fmt.Println("Raw response body:", string(body)) // Debugging step

	// var responseMap map[string]interface{}
	// if err := json.Unmarshal(body, &responseMap); err != nil {
	// 	log.Printf("Error unmarshaling response body: %v. Raw response: %s", err, string(body))
	// 	return
	// }

	// if responseMap["items"] != nil {
	// 	fmt.Println("id")
	// 	fmt.Println(responseMap["items"])
	// }

	// if responseMap["status"] != nil {
	// 	fmt.Println("status")
	// 	fmt.Println(responseMap["status"])
	// }

}

func GetPagamento(c *gin.Context, access_token string) {

	//access_token := bootstrap.NewEnv().AccessToken
	paymentID := c.Param("id")

	if paymentID == "" {
		println("Payment ID is empty")
	}

	url := fmt.Sprintf("https://api.mercadopago.com/v1/payments/%s", paymentID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check for non-200 responses
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Request failed with status: %d - %s", resp.StatusCode, string(body))

	}

	// body := resp.Body.Read([]byte)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return
	}
	//fmt.Println("Raw response body:", body) // Debugging step
	// fmt.Println("Teste") 
	//fmt.Println(body)
	//fmt.Println("Raw response body:", string(body)) // Debugging step
	fmt.Println("Raw response body:", string(body)) // Debugging step

	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		log.Printf("Error unmarshaling response body: %v. Raw response: %s", err, string(body))
		return
	}
	//fmt.Println(responseMap)
	if responseMap["items"] != nil {
		fmt.Println("id")
		fmt.Println(responseMap["items"])
	}

	if responseMap["status"] != nil {
		fmt.Println("status")
		fmt.Println(responseMap["status"])
	}

	// cfg, err := configMP.New(access_token)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// // client := preference.NewClient(cfg)
	// // ID := "104272673767"
	// // resource, err := client.Get(context.Background(), ID)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // 	return
	// // }

	// // fmt.Println(resource)

	// client := payment.NewClient(cfg)
	// ID := 104272673767
	// resource, err := client.Get(context.Background(), ID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(resource)

}

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Error generating random string: %v", err)
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

func MercadoPagoHandler(c *gin.Context, access_token string) {
	cfg, err := configMP.New(access_token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				ID:          generateRandomString(10),
				Title:       "Pedido",
				UnitPrice:   1,
				Quantity:    1,
				Description: "Combo",
				CurrencyID:  "BRL",
				CategoryID:  "food",
			},
		},
		PaymentMethods: &preference.PaymentMethodsRequest{
			ExcludedPaymentMethods: []preference.ExcludedPaymentMethodRequest{
				{
					ID: "bolbradesco",
				},
			},
		},
		NotificationURL: "https://webhook.site/90c7fca7-f67b-443b-8e75-2a0fa167c9e9",
	}

	client := preference.NewClient(cfg)
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"init_point": resource})

	// request := payment.Request{
	// 	TransactionAmount: 11,
	// 	PaymentMethodID : "pix",
	// 	Description: "Teste",
	// 	Payer: &payment.PayerRequest{
	// 		Email: "teste@teste.com",
	// 	},
	// 	//Token:        "{{CARD_TOKEN}}",
	// 	Installments: 1,
	// }

	// //payment_response = sdk.payment().create(payment_data, request_options)
	// //payment = payment_response["response"]

	// client := payment.NewClient(cfg)
	// result, err := client.Create(context.Background(), request)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(200, gin.H{"res": result})

	// AutoReturn: "all",
	// BackURLs: &preference.BackURLsRequest{
	// 	Success: "http://localhost:8080/pagamento/sucesso",
	// 	Pending: "http://localhost:8080/pagamento/pendente",
	// 	Failure: "http://localhost:8080/pagamento/falha",
	// },

	//c.JSON(200, gin.H{"init_point": resource.InitPoint})
}

// func BuscarPagamento(c *gin.Context, access_token string, id string) {
// 	cfg, err := config.New(access_token)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	client := preference.NewClient(cfg)
// 	resource, err := client.Get(context.Background(), id)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"status": resource.InitPoint})
// }

func PagamentoAtualizado(c *gin.Context) {
	c.JSON(200, gin.H{"response": "success"})
}
