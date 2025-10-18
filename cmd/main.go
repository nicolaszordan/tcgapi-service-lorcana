package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Card struct {
	Name            string `json:"name"`
	Set             string `json:"set"`
	CollectorNumber string `json:"collector_number"`
	ImageURL        string `json:"image_url"`
}

func main() {
	slog.Info("tcgapi-service-lorcana starting")

	http.HandleFunc("/health", health)
	http.HandleFunc("/cards", cards)

	slog.Info("starting server on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		slog.Error("server error", "error", err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	slog.Info("health check")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

type CardsQuerryParameters struct {
	Name            string
	Set             string
	CollectorNumber string
}

func cards(w http.ResponseWriter, r *http.Request) {
	slog.Info("cards")

	// Parse querry parameters
	params := get_cards_querry_parameters(r)
	slog.Info("querry parameters", "params", params)

	var cards []Card
	if params.Name != "" {
		cards = get_cards_from_name(params.Name)
	} else if params.Set != "" && params.CollectorNumber != "" {
		cards = get_cards_from_set_and_collector(params.Set, params.CollectorNumber)
	} else {
		slog.Info("no valid querry parameters provided")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: no valid querry parameters provided"))
		return
	}

	// if len(cards) == 0 {
	// 	slog.Info("no cards found")
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("Not Found: no cards found"))
	// 	return
	// }

	slog.Info("cards found", "count", len(cards))

	cardJson, err := json.Marshal(cards)
	if err != nil {
		slog.Error("error marshalling cards to json", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error: error marshalling cards to json"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(cardJson)
}

func get_cards_querry_parameters(r *http.Request) CardsQuerryParameters {
	return CardsQuerryParameters{
		Name:            r.URL.Query().Get("name"),
		Set:             r.URL.Query().Get("set"),
		CollectorNumber: r.URL.Query().Get("collector_number"),
	}
}

func get_cards_from_name(name string) []Card {
	slog.Info("getting cards from name", "name", name)

	return []Card{
		{
			Name:            name,
			Set:             "Sample Set",
			CollectorNumber: "001",
			ImageURL:        "https://example.com/sample_card.jpg",
		},
	}
}

func get_cards_from_set_and_collector(set string, collectorNumber string) []Card {
	slog.Info("getting cards from set and collector number", "set", set, "collector_number", collectorNumber)

	return []Card{
		{
			Name:            "Sample Card",
			Set:             set,
			CollectorNumber: collectorNumber,
			ImageURL:        "https://example.com/sample_card.jpg",
		},
	}
}
