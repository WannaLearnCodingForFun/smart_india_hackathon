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

















