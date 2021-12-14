# Go Gin Example With MongoDB

## Requisites
1. Go 1.17+
2. mongoDB 3.6+
3. gin 1.7+

## Installation
You can install and run the project with these steps.

#### 1. install dependencies
install the dependencies from go mod with
```$xslt
    go mod install
```

#### 2. configuration
add existing config.yml with cp from example :
```$xslt
    cp config.yml.example config.yml
```
and then fill the necessary data, example
```$xslt
ServerPort: ":3000"
DatabaseHost: "localhost"
DatabasePort: "27017"
DatabaseUser: "user"
DatabasePassword: "password"
DatabaseName: "loans"
```
or you can just use the example config.yml, or you can just straight run the server 
without set up any of config, but it will use the default config.

#### 4. run the server
go to this project root folder
```$xslt
    cd ROOT-FOLDER-PROJECT/
```
and then
```$xslt
    go run main.go
```
## Features
1. Open API (swagger), you can access it on (http://localhost:3000/swagger/index.html)
2. using MongoDB
3. validation by struct
4. using config with viper (https://github.com/spf13/viper)
5. using gin HTTP framework (https://github.com/gin-gonic/gin)

## API Documentations
### 1. Create Loan Process
#### Request Body
create loan process with this API endpoint and with this request body and set Content-Type to application/json
```json
{
    "full_name": "John Doe",
    "gender": "L",
    "ktp_number": "7878888888888",
    "image_of_ktp": "image.jpg",
    "image_of_selfie": "image.jpg",
    "date_of_birth": "1993-05-03",
    "address": "Jl. Cipanas ",
    "address_province": "Jawa Barat",
    "phone_number": "+62898293929",
    "email": "example@mail.com",
    "nationality": "INDONESIA",
    "loan_amount": "1500000",
    "tenor": "6"
}
```
#### Response 400 (Bad Request)
this is the response for Bad Request For KTP Number Validation
```json
{
    "meta": {
        "code": 400,
        "message": "Bad Request",
        "error": "Can't Process With the KTP Number (Is Already being Used)"
    }
}
```
this is the response for Bad Request For Gender
```json
{
    "meta": {
      "code": 400,
      "message": "Bad Request",
      "error": "Gender Is Not Valid, the suggesting value is (L or P)"
    }
}
```
this is the response for Bad Request For Age (From date_of_birth value)
```json
{
    "meta": {
      "code": 400,
      "message": "Bad Request",
      "error": "Age Is Not Valid, the suggesting value format is (2006-01-02)"
    }
}
```
this is the response for Bad Request For Age (From date_of_birth value)
```json
{
    "meta": {
      "code": 400,
      "message": "Bad Request",
      "error": "Age Is Not Valid, you must be at least 17 years old or not older than 80 years old"
    }
}
```
this is the response for Bad Request For Amount Loan
```json
{
    "meta": {
      "code": 400,
      "message": "Bad Request",
      "error": "Amount Loan Is Not Valid, the suggesting value is (1000000 - 10000000)"
    }
}
```
this is the response for Bad Request For Province
```json
{
    "meta": {
      "code": 400,
      "message": "Bad Request",
      "error": "Province Is Not Valid, the suggesting value is (DKI JAKARTA, JAWA BARAT, JAWA TIMUR OR SUMATERA UTARA)"
    }
}
```
this is the response for Bad Request For Tenor
```json
{
    "meta": {
      "code": 400,
      "message": "Bad Request",
      "error": "Tenor Is Not Valid, the suggesting value is (3, 6, 9, 12 or 24)"
    }
}
```

#### Response 200 (OK)
this is the response for Success Create Loan Process With Status Accepted and process the installment
```json
{
  "meta": {
    "code": 200,
    "message": "Operation Loan Process Is Successfully Executed",
    "error": null
  },
  "data": {
    "id": 9,
    "full_name": "wwewe",
    "gender": "L",
    "ktp_number": "567867867888888",
    "image_of_ktp": "image.jpg",
    "image_of_selfie": "image.jpg",
    "date_of_birth": "1993-05-03",
    "address": "Jl. Cipanas ",
    "address_province": "JAWA BARAT",
    "phone_number": "+62898293929",
    "email": "example@mail.com",
    "nationality": "INDONESIA",
    "loan_amount": "1500000",
    "status": "Accepted",
    "tenor": "6",
    "created_at": "2021-12-11T09:57:48.089427Z",
    "installment": {
      "id": 3,
      "loan_id": 9,
      "all_paid_off": false,
      "tenor_remaining": "6",
      "total_tenor": "6",
      "installment_amount": "252500.00",
      "total_installment_amount": "1500000",
      "created_at": "2021-12-11T09:57:48.091599Z"
    }
  }
}
```

## License
This package is open-sourced software licensed under the MIT license.