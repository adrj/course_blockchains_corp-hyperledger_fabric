# Blockchains Corporativos - Hyperledger Fabric

This repository contains chaincode examples and customizations developed during the course:

[Blockchains Corporativos - Hyperledger Fabric (GoLedger partnership)](https://esr.rnp.br/cursos/blockchains-corporativos-hyperledger-fabric-parceria-oficial-goledger-seg36/)


## Examples and Customizations

- Created a new asset type: **Car**
  - Properties:
    - `id` (string, IsKey, uses transaction ID)
    - `make` (string, Required)  <!-- Atualizado: agora obrigatório -->
    - `model` (string, Required)
    - `colour` (string, Required, Writers: org1 & org2 only)
    - `owner` (reference to Person, Required)
    - `dateTransfered` (datetime)
- Updated the `owner` property in Car to reference the Person asset.
- Added **make** as a required property to Car asset for full compatibility with the `registerCar` transaction.
- Created a custom datatype: **bookRating**
  - Type: number
  - Validation: value must be >= 1.0 and <= 10.0
- Added the `bookRating` property to the Book asset.
- Added a custom transaction (`registerCar`) to register a new Car asset using the transaction ID as the car's primary key. The transaction receives make, model, colour, owner, and dateTransfered as parameters, and automatically sets the car's ID to the transaction ID (deterministic and available to all peers).
- Updated chaincode to version 0.3, sequence 3.
- All changes and code updates were made by Adalto dos Reis Junior (adrj).

---

For details, see the chaincode source files and commit history.
