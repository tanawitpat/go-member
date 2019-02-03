# Membership API

| Method Name | HTTP Method | Path | Summary |
| :--- | :---: | :--- | :--- |
| CreateAccount | `POST` | /customer | Create customer's profile |
| UpdateAccount | `PUT` | /customer | Update customer's profile |
| InquiryAccount | `GET` | /customer/<customer_id> | Inquiry customer's profile by customer_id |


## CreateAccount

### Description
Create customer's profile.

### URL
> POST /customer

### Parameters
| Parameters | Description | Values | Remark |
| --- | --- | --- | --- |
| **first_name** | First name | string | Required |
| **last_name** | Last name  | string | Required |
| **mobile_number** | Mobile number  | string | Required |
| **email** | Email | string | Required |
| **address** | Address | Address object | Required |

#### Address object
| Parameters | Description | Values | Remark |
| --- | --- | --- | --- |
| **street_address** | Street address | string | Required |
| **subdistrict** | Subdistrict  | string | Required |
| **district** | District  | string | Required |
| **province** | Province | string | Required |
| **postal_code** | Postal code | string | Required |

### Output
Request status and customer ID

### Sample Output
```
> /customer
with header
{
    "Content-Type": "application/json",
    "request_id": "00001"
}
with body
{
    "first_name": "Tanawit",
    "last_name": "Pattanaveerangkoon",
    "mobile_number": "+66890001111",
    "email": "abc@gmail.com",
    "address": {
        "street_address": "100/100 Yotha Rd.",
        "subdistinct": "Talad Noi",
        "distinct": "Samphanthawong",
        "province": "Bangkok",
        "postal_code": "10100"
    }
}

> Success response
return value
{
    "status": "SUCCEEDED",
    "customer_id": "1",
    "account_status": "ACTIVE"
}

> Fail response
{
    "status": "FAILED",
    "customer_id": "",
    "account_status": "",
    "error": {
        "name": "BAD_REQUEST",
        "details": [
            {
                "field": "email",
                "issue": "This email has been used"
            },
            {
                "field": "mobile_number",
                "issue": "This mobile number has been used"
            }
        ]
    }
}
```

## UpdateAccount

### Description
Update customer's profile.

### URL
> PUT /customer

### Parameters
| Parameters | Description | Values | Remark |
| --- | --- | --- | --- |
| **customer_id** | Customer ID | string | Required |
| **first_name** | First name | string | Optional |
| **last_name** | Last name  | string | Optional |
| **mobile_number** | Mobile number  | string | Optional |
| **email** | Email | string | Optional |
| **address** | Address | Address object | Optional |

#### Address object
| Parameters | Description | Values | Remark |
| --- | --- | --- | --- |
| **street_address** | Street address | string | Optional |
| **subdistrict** | Subdistrict  | string | Optional |
| **district** | District  | string | Optional |
| **province** | Province | string | Optional |
| **postal_code** | Postal code | string | Optional |

### Output
Request status

### Sample Output
```
> /customer
with header
{
    "Content-Type": "application/json",
    "request_id": "00002"
}
with body
{
    "customer_id": "1",
    "mobile_number": "+66890001112",
    "email": "def@gmail.com"
}

> Success response
return value
{
    "status": "SUCCEEDED",
}

> Fail response
{
    "status": "FAILED",
    "error": {
        "name": "BAD_REQUEST",
        "details": [
            {
                "field": "email",
                "issue": "This email has been used"
            },
            {
                "field": "mobile_number",
                "issue": "This mobile number has been used"
            }
        ]
    }
}
```

## InquiryAccount

### Description
Inquiry customer's profile.

### URL
> GET /customer/<customer_id>

### Parameters
| Parameters | Description | Values | Remark |
| --- | --- | --- | --- |
| **customer_id** | Customer ID | string | Required |

### Output
Request status

### Sample Output
```
> /customer/1

> Success response
return value
{
    "status": "SUCCEEDED",
    "first_name": "Tanawit",
    "last_name": "Pattanaveerangkoon",
    "email": "def@gmail.com",
    "mobile_number": "+66890001112",
    "address": {
        "street_address": "100/100 Yotha Rd.",
        "subdistinct": "Talad Noi",
        "distinct": "Samphanthawong",
        "province": "Bangkok",
        "postal_code": "10100"
    }
}

> Fail response
{
    "status": "FAILED",
    "first_name": "",
    "last_name": "",
    "email": "",
    "mobile_number": "",
    "address": {
        "street_address": "",
        "subdistinct": "",
        "distinct": "",
        "province": "",
        "postal_code": ""
    }
    "error": {
        "name": "ACCOUNT_NOT_FOUND",
        "details": [
            {
                "field": "customer_id",
                "issue": "customer_id does not exist"
            }
        ]
    }
}
```