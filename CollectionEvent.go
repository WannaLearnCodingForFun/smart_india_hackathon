package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CollectionEvent struct 
type CollectionEvent struct {
	ID         string   `json:"id"`         // Unique Event ID (e.g., "COL123")
	Herb       string   `json:"herb"`       // Herb ka naam (e.g., Ashwagandha)
	GPS        string   `json:"gps"`        // GPS Coordinates jahan se herb collect hua
	Collector  string   `json:"collector"`  // Collector/Farmer ka ID ya naam
	Timestamp  string   `json:"timestamp"`  // Event ka timestamp
	Quality    string   `json:"quality"`    // Initial quality notes (e.g., Moisture: 12%)
	Certified  bool     `json:"certified"`  // Kya event sustainability rules ke hisaab se valid hai?
	Docs       []string `json:"docs"`       // Additional docs ka array (lab reports, photo links, etc.)
	PreviousTx string   `json:"previousTx"` // Last blockchain Tx ka reference (audit trail strengthen karne ke liye)
}

// SmartContract define karta hai chaincode structure
type SmartContract struct {
	contractapi.Contract
}

//  Helper: CollectionEventExists -> check karega ki ID ledger me already hai ya nahi
func (s *SmartContract) CollectionEventExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	eventJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("ledger read failed: %v", err)
	}
	return eventJSON != nil, nil
}

// Create: AddCollectionEvent -> naye collection event ko ledger me add karega
func (s *SmartContract) AddCollectionEvent(ctx contractapi.TransactionContextInterface,
	id string, herb string, gps string, collector string, quality string, docs []string) error {

	// Step 1: Pehle check karo ID already exist to nahi karti
	exists, err := s.CollectionEventExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("collection event %s already exists", id)
	}

	// Step 2: Timestamp automatic generate karo agar user ne nahi diya
	timestamp := time.Now().UTC().Format(time.RFC3339)

	// Step 3: Sustainability check -> (demo rule: Ashwagandha ok, banned herb not ok)
	certified := true
	if herb == "BannedHerb" {
		certified = false
	}

	// Step 4: Previous Tx ka ID le lo (audit trail ke liye)
	previousTx := ctx.GetStub().GetTxID()

	// Step 5: Event object banao
	event := CollectionEvent{
		ID:         id,
		Herb:       herb,
		GPS:        gps,
		Collector:  collector,
		Timestamp:  timestamp,
		Quality:    quality,
		Certified:  certified,
		Docs:       docs,
		PreviousTx: previousTx,
	}

	// Step 6: Marshal into JSON
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %v", err)
	}

	// Step 7: Ledger me store karenge
	return ctx.GetStub().PutState(id, eventJSON)
}

🔹 Read: ReadCollectionEvent -> ek event ko ledger se fetch karega
func (s *SmartContract) ReadCollectionEvent(ctx contractapi.TransactionContextInterface, id string) (*CollectionEvent, error) {
	eventJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("ledger read error: %v", err)
	}
	if eventJSON == nil {
		return nil, fmt.Errorf("collection event %s does not exist", id)
	}

	var event CollectionEvent
	err = json.Unmarshal(eventJSON, &event)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %v", err)
	}

	return &event, nil
}

// Update: UpdateCollectionQuality -> agar naya quality test aaya to update kar sakte ho
func (s *SmartContract) UpdateCollectionQuality(ctx contractapi.TransactionContextInterface, id string, newQuality string) error {
	event, err := s.ReadCollectionEvent(ctx, id)
	if err != nil {
		return err
	}

	// Purana data rakho aur sirf quality update karo
	event.Quality = newQuality
	event.PreviousTx = ctx.GetStub().GetTxID()

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, eventJSON)
}

//  Query: GetAllCollectionEvents -> saare events ek sath retrieve karne ka function
func (s *SmartContract) GetAllCollectionEvents(ctx contractapi.TransactionContextInterface) ([]*CollectionEvent, error) {
	query := `{"selector":{"id":{"$regex":".*"}}}` // CouchDB query -> saare docs with any id

	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var events []*CollectionEvent
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var event CollectionEvent
		err = json.Unmarshal(queryResponse.Value, &event)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}
