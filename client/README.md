# Go API client for openapi

Customers ...

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/moov-io/customers](https://github.com/moov-io/customers)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8087*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CustomersApi* | [**AddCustomerAddress**](docs/CustomersApi.md#addcustomeraddress) | **Post** /customers/{customerID}/address | Add an Address onto an existing Customer record
*CustomersApi* | [**CreateCustomer**](docs/CustomersApi.md#createcustomer) | **Post** /customers | Create a new customer
*CustomersApi* | [**GetCustomer**](docs/CustomersApi.md#getcustomer) | **Get** /customers/{customerID} | Retrieves a Customer object associated with the customer ID.
*CustomersApi* | [**GetCustomerDocumentContents**](docs/CustomersApi.md#getcustomerdocumentcontents) | **Get** /customers/{customerID}/documents/{documentID} | Retrieve the referenced document
*CustomersApi* | [**GetCustomerDocuments**](docs/CustomersApi.md#getcustomerdocuments) | **Get** /customers/{customerID}/documents | Get documents for a customer
*CustomersApi* | [**Ping**](docs/CustomersApi.md#ping) | **Get** /ping | Ping the Customers service to check if running
*CustomersApi* | [**ReplaceCustomerMetadata**](docs/CustomersApi.md#replacecustomermetadata) | **Put** /customers/{customerID}/metadata | Replace the metadata object for a customer. Metadata is a map of unique keys associated to values to act as foreign key relationships or arbitrary data associated to a Customer.
*CustomersApi* | [**UpdateCustomerStatus**](docs/CustomersApi.md#updatecustomerstatus) | **Put** /customers/{customerID}/status | Update the status for a customer, which can only be updated by authenticated users with permissions.
*CustomersApi* | [**UploadCustomerDocument**](docs/CustomersApi.md#uploadcustomerdocument) | **Post** /customers/{customerID}/documents | Upload a document for the given customer.


## Documentation For Models

 - [Address](docs/Address.md)
 - [CreateAddress](docs/CreateAddress.md)
 - [CreateCustomer](docs/CreateCustomer.md)
 - [CreatePhone](docs/CreatePhone.md)
 - [Customer](docs/Customer.md)
 - [CustomerMetadata](docs/CustomerMetadata.md)
 - [Document](docs/Document.md)
 - [Error](docs/Error.md)
 - [Phone](docs/Phone.md)
 - [UpdateCustomerStatus](docs/UpdateCustomerStatus.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Author



