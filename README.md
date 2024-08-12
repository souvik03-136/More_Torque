---
# More Torque Taxi Service API

Welcome to the API documentation for More Torque, a taxi service company providing cars to various organizations. This document outlines the key entities and endpoints available in our system.

## Key Entities

- **Vehicle**: Cars provided for taxi services.
- **Org (Organization)**: Companies or organizations that have contracts with More Torque. Organizations can have parent-child relationships and specific policies.

## API Endpoints

### Vehicles

- **GET /vehicles/decode/:vin**
  - **Description**: Decode a VIN (Vehicle Identification Number) to fetch vehicle details.
  - **Details**: Calls NHTSA’s DecodeVin endpoint. Limits to 5 requests per minute.
  - **Response**: Manufacturer, model, year.

- **POST /vehicles**
  - **Description**: Add a new vehicle to the system.
  - **Request Body**:
    ```json
    {
      "vin": "xxxxxxxx",
      "org": "yyyyyy"
    }
    ```
  - **Validations**: VIN must be a valid 17-digit alphanumeric string. Org must exist.
  - **Response**: 
    - Success: 201 status code with vehicle details.
    - Error: 400 status code with error message.

- **GET /vehicles/:vin**
  - **Description**: Fetch details of a vehicle by VIN.
  - **Validations**: VIN must be valid and present in the system.
  - **Response**:
    - Success: 200 status code with vehicle details.
    - Error: 400 or 404 status code with an error message.

### Organizations

- **POST /orgs**
  - **Description**: Create a new organization.
  - **Request Body**:
    ```json
    {
      "name": "Org1",
      "account": "acc1",
      "website": "www.org1.com",
      "fuelReimbursementPolicy": "policy1",
      "speedLimitPolicy": "policy2"
    }
    ```
  - **Default Values**: `fuelReimbursementPolicy` defaults to `1000` if not provided.
  - **Response**:
    - Success: 201 status code with organization details.
    - Error: 400 status code with error message.

- **PATCH /orgs**
  - **Description**: Update an existing organization's details.
  - **Request Body**:
    ```json
    {
      "account": "acc1",
      "website": "www.org1.com",
      "fuelReimbursementPolicy": "policy1",
      "speedLimitPolicy": "policy2"
    }
    ```
  - **Inheritance Rules**:
    - **Fuel Reimbursement Policy**: Inherited from parent org; cannot be updated at child level if inherited.
    - **Speed Limit Policy**: Inherited from parent but can be overridden by child orgs.
  - **Response**:
    - Success: 200 status code with updated organization details.
    - Error: 400 status code with error message.

- **GET /orgs**
  - **Description**: Retrieve details of all organizations, including inherited policies.
  - **Response**:
    - Success: 200 status code with a list of organizations.
    - Error: 400 status code with error message.
  - **Bonus**: Pagination support for large number of child orgs.

---
