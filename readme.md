Introduction

This Postman collection contains a set of API requests to interact with the MNC Bank API.
Getting started

To get started, make sure you have Postman installed on your local machine. You can import this collection into Postman to get all the API endpoints ready for testing.
Usage
Customer Endpoints

    Register: Use this endpoint to register new customers. POST Method. URL: http://localhost:8080/Register
    Login: Use this endpoint to log in as a customer. Method: POST. URL: [URL here]
    Deleting a Subscriber: Use this endpoint to delete a customer. GET method. URL: [URL here]

Merchant Endpoints

    Adding a Merchant: Use this endpoint to add a new merchant. POST method. URL: http://localhost:8080/merchant
    Get All Merchants: Use this endpoint to get all merchants. Method: GET. URL: [URL here]
    Get By Id: Use this endpoint to get a specific merchant by ID. Method GET. URL: [URL here]
    Delete Merchant: Use this endpoint to delete a merchant. DELETE method. URL: http://localhost:8080/merchant/1
    Updating a Merchant: Use this endpoint to update a merchant. Method: PUT. URL: http://localhost:8080/merchant

Transaction Endpoints

    Add Transaction: Use this endpoint to add a new transaction. POST method. URL: http://localhost:8080/Transaksi
    Get All Transactions: Use this endpoint to get all transactions. Method: GET. URL: http://localhost:8080/Transaksi
    Get Transactions Based on Id: Use this endpoint to get a specific transaction by ID. Method: GET. URL: http://localhost:8080/Transaksi

Exit

    Exit: Use this endpoint to exit. POST method. URL: http://localhost:8080/logout

Note

Replace [URL Here] with the actual URL for a particular endpoint