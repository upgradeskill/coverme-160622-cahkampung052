# Bootcamp Golang Majoo

Basic starting point for a new golang project.

## Creating a new project

 - Rename `cmd/golang-starter` folder
 - Rename `APP_NAME` in `Makefile`
 - Rename `go.mod` with proper username and repository names
 - Update imports accordingly

## Usage

- `make build`
- `make tests`
- `make run`
- `make clean`

## Rest API Documentation

In this service there are 4 endpoints that can retrieve, input, update, and delete product data

#### A. GET All Product

- **Endpoint:** /products
- **Method:** GET
- **Example Request:** http://localhost:8000/products
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success get all product",
    "data": [
        {
            "code": "001",
            "name": "Susu Sapi",
            "price": 10000
        },
        {
            "code": "002",
            "name": "Susu Kambing",
            "price": 10000
        },
        {
            "code": "003",
            "name": "Susu Kuda",
            "price": 10000
        }
    ]
}
```
#### B. GET Spesific Product

- **Endpoint:** /products/{:product_code}
- **Method:** GET
- **Example Request:** http://localhost:8000/products/002
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success get single product",
    "data": {
        "code": "002",
        "name": "Susu Kambing",
        "price": 10000
    }
}
```
- **Example Error Response (Status Code: 400**)
```json
{
    "message": "product not found",
    "data": {}
}
```
#### C. Create New Product

- **Endpoint:** /products
- **Method:** POST
- **Content-Type:** application/json
- **Example Request:** http://localhost:8000/products
- **Example Raw Body:**
```json
{
    "code": "005",
    "name": "susu murni",
    "price": 2000
}
```
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success create new product",
    "data": {
        "code": "005",
        "name": "susu murni",
        "price": 2000
    }
}
```
- **Example Error Response (Status Code: 400**)
```json
{
    "message": "product code already exist",
    "data": {}
}
```
#### D. Update Existing Product

- **Endpoint:** /products
- **Method:** PUT
- **Content-Type:** application/json
- **Example Request:** http://localhost:8000/products
- **Example Raw Body:**
```json
{
    "code": "002",
    "name": "Bakso Bakar",
    "price": 2000
}
```
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success to update product",
    "data": {
        "code": "005",
        "name": "Bakso Bakar",
        "price": 2000
    }
}
```
- **Example Error Response (Status Code: 400**)
```json
{
    "message": "product code is not in our database",
    "data": {}
}
```
#### E. Delete Existing Product

- **Endpoint:** /products/{:product_code}
- **Method:** DELETE
- **Content-Type:** application/json
- **Example Request:** http://localhost:8000/products/005
- **Example Success Response (Status Code: 200**)
```json
{
    "message": "success to delete product",
    "data": "005"
}
```
- **Example Error Response (Status Code: 400**)
```json
{
    "message": "product code is not in our database",
    "data": {}
}
```