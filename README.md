# Go Cards

## Dependencies

- `go@1.18`
- `docker` and `docker compose`

## Setup

- `cp .env.example .env` and update the values if needed
- `cd src`
- `go mod download` to download dependencies
- `docker compose -f docker-compose.yml up -d --build` to start app and database

- ```bash
   curl -X POST http://localhost:3000/api/v1/deck  -H 'Content-Type: application/json' \
   -d '{"shuffled":true,"cards":[]}'
   ``` 
   to create a deck
- ```bash
    curl -X GET http://localhost:3000/api/v1/deck/:deck_id  -H 'Content-Type: application/json'
    ``` 
    to get a deck
- ```bash
    curl -X POST http://localhost:3000/api/v1/deck/:deck_id/draw  -H 'Content-Type: application/json' \
    -d '{"count":1}'
    ``` 
    to draw cards from a deck

## Major Design Decisions

- `fiber` as the web framework and `gorm` as the ORM
- Decided to use relational database since for a game, it would be of more use
- Not using `Card` model in db since made no sense to have a table for cards
- `Card` is used as `jsonb` in `Deck` model which is as performant as a nosql db if needed
- There were libraries for handling `docker` containers for testing but decided to write a quick utility instead which is very flaky
- A lot of rough edges were made due to time constraints which can be easily be refactored
- ```golang
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
    ```
    Code like this should be scoped to util function like  `createDeckRequest` and `getDeckRequest` to avoid repetition

# Testing
Decided to focus on integration testing and basic unit testing of models
- Need `docker` since it uses `postgres` for integration testing
- `docker pull postgres:13` the image used for testing
- `cd src` 
- `go test -v ./...` to run all tests