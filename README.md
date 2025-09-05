# Blockchains Corporativos - Hyperledger Fabric

This repository contains chaincode examples and customizations developed during the course:

[Blockchains Corporativos - Hyperledger Fabric (GoLedger partnership)](https://esr.rnp.br/cursos/blockchains-corporativos-hyperledger-fabric-parceria-oficial-goledger-seg36/)

## Examples and Customizations

- Created a new asset type: **Car**
  - Properties:
    - `make` (string, IsKey)
    - `model` (string, IsKey)
    - `colour` (string, Required, Writers: org1 & org2 only)
    - `owner` (reference to Person)
- Updated the `owner` property in Car to reference the Person asset.
- Created a custom datatype: **bookRating**
  - Type: number
  - Validation: value must be >= 1.0 and <= 10.0
- Added the `bookRating` property to the Book asset.
- Updated chaincode to version 0.3, sequence 3.
- All changes and code updates were made by Adalto dos Reis Junior (adrj).

---

For details, see the chaincode source files and commit history.
