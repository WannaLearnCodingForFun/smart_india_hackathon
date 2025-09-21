# smart_india_hackathon
Working out the workflow.


Entire workflow:
Fabric (Go chaincode) + Node/Express backend + Next.js frontend

We'll be running a local Fabric test-network, by implementing  small Go chaincode to record batch hashes/CIDs and emit events, 
build a Node backend that accepts collector data, pins files to IPFS or saves to Postgres, submits Fabric transactions (and returns txId),
and a Next.js frontend with collector forms + QR/consumer pages.


Atharva:
Learn Fabric fundamentals, chaincode lifecycle, Fabric CA basics.
Need to design chaincode data model & endorsement policies.
Implement and test Go chaincode functions: RecordBatch, RecordEvent, QueryBatch.
Implement chaincode events (SetEvent) for UI real-time updates.
Help integrate backend to Fabric Node SDK; verify txIds, event listeners, QSCC ledger lookups.

Final things to give: chaincode repo, deployment scripts, sample transactions and docs.

Kunsh — Backend + DevOps
Implement Node/Express server: REST endpoints for collection, processing, quality tests, batch creation.
Integrate fabric-network (Node SDK) to submit/evaluate transactions and to listen for chaincode events.
Implement DB (Postgres JSONB) or local IPFS pinning (Pinata) helpers.
Dockerize backend, DB, wallet, connection profile; prepare docker-compose.

Final things to give: REST API, wallet + connection-profile setup, docker-compose, runbook.

Aditya - Frontend(Nextjs)
Build collector forms (mobile-first), photo upload to IPFS endpoint or backend, geo-capture (Geolocation API) with manual fallback.
Build batch finalization UI and QR generation page. Consumer page to fetch and show provenance bundle + map (Leaflet).
Connect UI flows to backend endpoints, handle offline-save-to-localStorage and sync UI.

Deliverables: Next.js app with collector pages, consumer/batch pages, and QR pages.


Aditya along wth Maria:
Design simple, accessible forms with big touch targets, clear labels, minimal typing.
Create timeline / provenance visualization for consumer page and farmer/community profile template.
Prepare slides, screenshots and UI flow to narrate demo to judges.

Deliverables: UI mockups, final polished CSS/Tailwind styling, demo narrative.

Pranjal & Smera — Lab/IoT Simulation( Currently useless)
Implement a small Flask/FastAPI microservice that returns simulated lab test JSON for QualityTest.
Provide a simple IoT simulator script to generate temperature/humidity readings for CollectionEvent (optionally POST to backend).

Deliverables: lab-simulator endpoints, sample JSON certificates (pin to IPFS or store in DB).



Visual flow of Herbs::--

[Farmer / Collector] 
   │   (GPS-tagged harvest → CollectionEvent on blockchain)
   ▼
   
[Aggregation / Storage Center]
   │   (Batch assigned, custody handoff → ProcessingStep)
   ▼
   
[Processing Facility]
   │   (Drying, cleaning, grinding → ProcessingStep)
   ▼
   
[Laboratory]
   │   (Quality checks, lab certificate → QualityTest)
   ▼
   
[Manufacturer]
   │   (Formulation & packaging → Formulation event)
   ▼
   
[QR Code Generation]
   │   (Unique QR linked to provenance record on blockchain)
   ▼
   
[Consumer]
   │   (Scans QR → sees map, collector, lab results, compliance proofs)

I will tell you guys a story:
Step 1 — Harvest (Farmer / Collector)
A farmer in Maharashtra harvests Ashwagandha roots.
He opens the collector mobile DApp (or sends an SMS).
The app automatically captures:
GPS coordinates of the farm,
Timestamp,
His Collector ID,
Herb species (Withania somnifera),
Weight harvested (10 kg),
Photo of the roots.
A CollectionEvent transaction is created on the blockchain.

On the dashboard, we can see a new entry appear with “Ashwagandha harvested, 10 kg, GPS-verified.”

Step 2 — Aggregation (Cooperative Center)
The farmer delivers the roots to the local cooperative.
At the coop, staff weigh and tag the batch.
The system creates a ProcessingStep (Aggregation) event, linking the farmer’s harvest into Batch #ASH-0001.

Blockchain now shows a custody handoff: Farmer → Cooperative.

Step 3 — Processing (Drying & Cleaning)
The cooperative dries and cleans the roots.
Operator records: start & end time, final weight (8.5 kg after drying).
The DApp creates a ProcessingStep (Drying) event.
Blockchain shows: “Batch #ASH-0001 processed — dried and cleaned.”

Step 4 — Quality Testing (Laboratory)
A sample is sent to an accredited lab.
The lab uploads results:
Moisture content = 9.5% (pass),
Pesticide residue = below limit (pass),
DNA barcode = match (verified Ashwagandha).
Lab also attaches a certificate PDF (hash stored on-chain).
A QualityTest event is added.

Dashboard shows: “Batch #ASH-0001 passed all quality checks.”

Step 5 — Formulation (Manufacturer)

A manufacturer takes the processed batch and creates Ashwagandha tablets.
The system records:
Input batch (ASH-0001),
Manufacturer ID,
Date of formulation,
Output product batch (PROD-ASH-TAB-001).

A Formulation event is committed on blockchain.

Step 6 — QR Code Generation
The system generates a unique QR code for product batch PROD-ASH-TAB-001.
This QR is printed on each bottle of Ashwagandha tablets.
On blockchain, a QRGeneration event links the QR to the full provenance record.

Step 7 — Consumer Scan (Transparency)
A customer buys a bottle of Ashwagandha tablets at a retail store.
They scan the QR code using their phone camera.
Instantly, the consumer web portal opens and displays:
An interactive map showing where the herb was harvested,
Name of the cooperative that aggregated it,
Lab certificate details (moisture %, pesticide test, DNA authentication),
Manufacturer details and packaging date,
A green compliance badge showing sustainable and fair-trade sourcing.
The customer can verify authenticity in real time.












