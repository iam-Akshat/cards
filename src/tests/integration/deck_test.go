package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iam-Akshat/cards/models"
	routes "github.com/iam-Akshat/cards/routes"
	testUtils "github.com/iam-Akshat/cards/tests/test_utils"
	dba "github.com/iam-Akshat/cards/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestDeck(t *testing.T) {
	dbConfig := dba.DatabaseConfig{
		Host:     "localhost",
		Port:     "5431",
		User:     "postgres",
		Password: "postgres",
		DBName:   "test_db",
	}
	err := testUtils.ExecutePostgresDockerContainer(&dbConfig)
	if err != nil {
		t.Error(err)
	}
	gorm, err := dba.NewDatabaseConnection(dbConfig)
	if err != nil {
		t.Error(err)
	}
	db = gorm

	t.Run("TestCreateDeck", TestCreateDeck)
	t.Run("TestGetDeck", TestGetDeck)

	t.Run("TestDrawCards", TestDrawCards)
}

func TestCreateDeck(t *testing.T) {
	t.Run("TestCreateValidDeck", TestCreateValidDeck)

	t.Run("TestCreateDeckShuffled", TestCheckShuffling)
}

func TestCreateValidDeck(t *testing.T) {
	testUtils.DropAllTables(db)
	app := fiber.New()
	routes.SetupRoutes(app, db)
	reqBody := `{
		"shuffled": true,
		"cards": []
	}`
	reader := strings.NewReader(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/deck", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}
	var res models.CreateDeckResponse

	json.NewDecoder(resp.Body).Decode(&res)

	if res.Remaining != 52 {
		t.Errorf("Expected remaining cards to be 52, got %d", res.Remaining)
	}
	if res.DeckId == uuid.Nil {
		t.Error("Expected deck id to be a valid uuid")
	}
	if res.Shuffled != true {
		t.Errorf("Expected shuffled to be true, got %t", res.Shuffled)
	}

}

func TestCheckShuffling(t *testing.T) {
	testUtils.DropAllTables(db)

	app := fiber.New()
	routes.SetupRoutes(app, db)
	reqBody := `{
		"shuffled": false,
		"cards": []
	}`
	reader := strings.NewReader(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/deck", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}
	var res models.CreateDeckResponse
	json.NewDecoder(resp.Body).Decode(&res)
	if res.Shuffled != false {
		t.Errorf("Expected shuffled to be false, got %t", res.Shuffled)
	}

	req, _ = http.NewRequest(http.MethodGet, "/api/v1/deck/"+res.DeckId.String(), nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	var deck models.Deck
	json.NewDecoder(resp.Body).Decode(&deck)
	// hacky test to check if the deck is shuffled
	if deck.Cards.Data[0].Code == "KS" && deck.Cards.Data[51].Value == "AD" && deck.Cards.Data[5].Value == "8S" {
		t.Errorf("Expected deck to be shuffled, got %v", deck.Cards.Data)
	}
}

func TestGetDeck(t *testing.T) {
	t.Run("TestGetValidDeck", TestGetValidDeck)
	t.Run("TestGetInvalidDeck", TestGetInvalidDeck)
}

func TestGetValidDeck(t *testing.T) {
	testUtils.DropAllTables(db)

	app := fiber.New()
	routes.SetupRoutes(app, db)
	// create a deck
	reqBody := `{
		"shuffled": true,
		"cards": []
	}`
	reader := strings.NewReader(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/deck", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}

	var res models.CreateDeckResponse
	json.NewDecoder(resp.Body).Decode(&res)

	req, _ = http.NewRequest(http.MethodGet, "/api/v1/deck/"+res.DeckId.String(), nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	var deck models.GetDeckResponse
	json.NewDecoder(resp.Body).Decode(&deck)
	if deck.DeckId != res.DeckId {
		t.Errorf("Expected deck id to be %s, got %s", res.DeckId, deck.DeckId)
	}
	if deck.Remaining != 52 {
		t.Errorf("Expected remaining cards to be 52, got %d", deck.Remaining)
	}
	if deck.Shuffled != true {
		t.Errorf("Expected shuffled to be true, got %t", deck.Shuffled)
	}
	if len(deck.Cards.Data) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(deck.Cards.Data))
	}
}

func TestGetInvalidDeck(t *testing.T) {
	testUtils.DropAllTables(db)

	app := fiber.New()
	routes.SetupRoutes(app, db)
	// create a deck
	reqBody := `{
		"shuffled": true,
		"cards": []
	}`
	reader := strings.NewReader(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/deck", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}

	var res models.CreateDeckResponse
	json.NewDecoder(resp.Body).Decode(&res)

	req, _ = http.NewRequest(http.MethodGet, "/api/v1/deck/"+uuid.New().String(), nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", resp.StatusCode)
	}
	var response map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&response)
	if response["message"] != "Record not found" {
		t.Errorf("%v", response)
		t.Errorf("Expected message to be 'Record not found', got %s", response["message"])
	}
}

func TestDrawCards(t *testing.T) {
	t.Run("TestDrawValidCards", TestDrawValidCards)
}

func TestDrawValidCards(t *testing.T) {
	testUtils.DropAllTables(db)

	app := fiber.New()
	routes.SetupRoutes(app, db)
	// create a deck
	reqBody := `{
		"shuffled": false,
		"cards": []
	}`
	reader := strings.NewReader(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/deck", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}

	var res models.CreateDeckResponse
	json.NewDecoder(resp.Body).Decode(&res)

	var reqBody2 = `{
		"count": 5
	}`
	req, _ = http.NewRequest(http.MethodPost, "/api/v1/deck/"+res.DeckId.String()+"/draw", strings.NewReader(reqBody2))
	req.Header.Set("Content-Type", "application/json")
	resp, err = app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	var deck []models.Card
	json.NewDecoder(resp.Body).Decode(&deck)

	if len(deck) != 5 {
		t.Errorf("Expected 5 cards, got %d", len(deck))
	}
	var first5CardsCode = []string{"KS", "QS", "JS", "10S", "9S"}
	for i, card := range deck {
		if string(card.Code) != first5CardsCode[i] {
			t.Errorf("Expected card code to be %s, got %s", first5CardsCode[i], card.Code)
		}
	}

	req, _ = http.NewRequest(http.MethodGet, "/api/v1/deck/"+res.DeckId.String(), nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	var deck2 models.GetDeckResponse
	json.NewDecoder(resp.Body).Decode(&deck2)
	if deck2.Remaining != 47 {
		t.Errorf("Expected remaining cards to be 47, got %d", deck2.Remaining)
	}

}

func TestDrawInvalidCards(t *testing.T) {
	// count validations mostly
	testUtils.DropAllTables(db)

	testUtils.DropAllTables(db)

	app := fiber.New()
	routes.SetupRoutes(app, db)
	// create a deck
	reqBody := `{
		"shuffled": false,
		"cards": []
	}`
	reader := strings.NewReader(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/deck", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}
	var countValidationTests = []struct {
		count int
	}{
		{count: 0},
		{count: 53},
		{count: -1},
	}
	for _, test := range countValidationTests {
		var reqBody2 = fmt.Sprintf(`{
			"count": %d
		}`, test.count)
		req, _ = http.NewRequest(http.MethodPost, "/api/v1/deck/"+uuid.New().String()+"/draw", strings.NewReader(reqBody2))
		req.Header.Set("Content-Type", "application/json")
		resp, err = app.Test(req)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code 400, got %d", resp.StatusCode)
		}
	}
}
