# Blockchains Corporativos - Hyperledger Fabric

This repository contains chaincode examples and customizations developed during the course:

[Blockchains Corporativos - Hyperledger Fabric (GoLedger partnership)](https://esr.rnp.br/cursos/blockchains-corporativos-hyperledger-fabric-parceria-oficial-goledger-seg36/)


## Examples and Customizations

### Asset Types
- **Car**
  - Properties:
    - `id` (string, IsKey, uses transaction ID)
    - `make` (string, Required)
    - `model` (string, Required)
    - `colour` (string, Required, Writers: org1 & org2 only)
    - `owner` (reference to Person, Required)
    - `dateTransfered` (datetime)
    - `lastTransferTxId` (string) - ID of the transaction that last transferred this car
- **Library**
  - Comprehensive library management asset
- **Person**
  - User/person asset with CPF validation
- **Secret**
  - Private data asset for confidential information
- **Custom Dynamic Asset Types**
  - Support for creating new asset types dynamically

### Data Types
- **BookRating**: Number validation (1.0 <= value <= 10.0)
- **BookType**: Enumeration for book categories
- **CPF**: Brazilian tax ID with validation

### Transactions
- **registerCar**: Register a new Car asset using the transaction ID as the car's primary key
- **transferCar**: Transfer car ownership between parties
- **getCarsByOwner**: Query all cars owned by a specific person, ordered by transfer date (descending)
- **createNewLibrary**: Create a new library with comprehensive metadata
- **getBooksByAuthor**: Query books by author name
- **getNumberOfBooksFromLibrary**: Get the total count of books in a library
- **updateBookTenant**: Update book tenant information

### Event Types
- **createLibraryLog**: Logging events for library creation
- Enhanced transaction event logging

### Features Added
- Collections configuration for private data management
- Comprehensive test coverage for all new functionality
- Enhanced Car asset with transaction tracking capabilities
- Refactored registerCar transaction with improved structure and event logging
- **getCarsByOwner transaction**: New query functionality to retrieve cars by owner with date-based sorting
- **CouchDB indexes**: Created optimized indexes for car queries (owner-dateTransfered-index.json, dateTransfered-index.json)
- **Enhanced data validation**: Improved CPF validation and asset key handling
- **Comprehensive testing data**: Created test data with valid Brazilian CPFs for functional validation
- Vendor dependencies and build configurations included
- Updated chaincode to latest version with sequence management
- All changes and code updates were made by Adalto dos Reis Junior (adrj)

---

For details, see the chaincode source files and commit history.
