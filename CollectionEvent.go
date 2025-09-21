package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type CollectionEvent struct {
	ID           string    `json:"id"`
	Timestamp    string    `json:"timestamp"`
	CollectorID  string    `json:"collectorId"`
	Species      string    `json:"species"`
	Lat          float64   `json:"lat"`
	Lon          float64   `json:"lon"`
	WeightKg     float64   `json:"weightKg"`
	MoisturePct  float64   `json:"moisturePct"`
	PhotoHash    string    `json:"photoHash"`
	Status       string    `json:"status"` // "RECEIVED", "REJECTED", "PENDING"
	Notes        string    `json:"notes"`
}

func (s *SmartContract) SubmitCollectionEvent(ctx contractapi.TransactionContextInterface, payload string) error {
	// identity / MSP check (example: ensure clientMSP == "FarmersMSP")
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client MSPID: %v", err)
	}
	if clientMSPID != "FarmersMSP" && clientMSPID != "CollectorsMSP" {
		return fmt.Errorf("unauthorized identity: %s", clientMSPID)
	}

	// parse payload JSON into struct
	var ce CollectionEvent
	if err := json.Unmarshal([]byte(payload), &ce); err != nil {
		return fmt.Errorf("invalid payload: %v", err)
	}

	// basic checks
	if ce.ID == "" || ce.Species == "" {
		return fmt.Errorf("missing required fields")
	}

	// bounding-box check for species zone (MVP - simple arrays; production: polygon service)
	// For demo: permissible bounding box for Ashwagandha
	minLat, maxLat := 19.00, 21.00
	minLon, maxLon := 73.50, 75.50
	if ce.Lat < minLat || ce.Lat > maxLat || ce.Lon < minLon || ce.Lon > maxLon {
		ce.Status = "REJECTED"
		ce.Notes = "outside permitted harvest zone"
	} else {
		// seasonal check (example: allowed harvest months April-August)
		ts, err := time.Parse(time.RFC3339, ce.Timestamp)
		if err != nil {
			return fmt.Errorf("invalid timestamp: %v", err)
		}
		month := ts.Month()
		if month < time.April || month > time.August {
			ce.Status = "PENDING"
			ce.Notes = "outside recommended season; manual review required"
		} else {
			ce.Status = "RECEIVED"
			ce.Notes = "accepted basic checks"
		}
	}

	// store to ledger
	ceBytes, _ := json.Marshal(ce)
	if err := ctx.GetStub().PutState(ce.ID, ceBytes); err != nil {
		return fmt.Errorf("failed to store collection event: %v", err)
	}

	// create an event for off-chain sync / notifications
	if err := ctx.GetStub().SetEvent("CollectionEventSubmitted", ceBytes); err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}
	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(err.Error())
	}
	if err := chaincode.Start(); err != nil {
		panic(err.Error())
	}
}
