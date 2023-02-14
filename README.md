# GoCompany

Welcome to the GoCompany wiki!

GoCompany is web application for Company PO Management.

Backend consists of:
1. REST Middleware (gorilla/mux)
2. Postgres DB (gorm)
3. Payload validator using validator10
4. Kafka event publishing mechanism.
5. Test Suite 
6. Docket Setup


**Pre-Requisites:**

1. Go 1.16 or newer 
2. Postgres DB(with role user/user)
3. Kafka producer with topic: go-company
4. Docker


**Bootstrap infrastructure and run application**
1. .env setup
Copy .env.example into .env
2. start server
in terminal/cmd with this app as the working directory run the following:
docker compose build
docker compose up -d



After starting server RESTful API server will be running at http://localhost:3000. It provides the following endpoints:

Company related APIs:
* GET /company/?name=company1         returns a detailed company object (takes query param)
* POST /company                       creates a new company (need to pass request body payload)
* PATCH /company                      updates an existing company( need to pass request body payload)
* DELETE /company/?name=company1      deletes a company (takes query param)

Sample payload for create:
    {
        "name"  : "com1",
        "amount" : 113,
        "registered": true  ,
        "type" : "NonProfit",
        "Description": "company1 India"
    }


Sample Outputs:

1. Create API(with token):

curl --location --request POST '127.0.0.1:8083/company/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "com1",
    "amount": 113,
    "registered": true,
    "type": "NonProfit",
    "Description": "company1 India"
}'

2. Get Company:
curl --location --request GET 'http://localhost:3000/company/?name=com1'

3. Delete Company:
curl --location --request DELETE 'http://localhost:3000/company/?name=com1'


****KAFKA published message:****
1. {"EventType":**"company_created"**,"Payload":{"ID":"0cc444a4-d6d1-4793-95a0-387c911f400f","Name":"com1","Description":"company1 India","Amount":113,"Registered":true,"Type":"NonProfit"},"Time":"2023-02-11T09:53:50.838038+04:00"}

2. "EventType":"**company_deleted",**"Payload":{"ID":"","Name":"com1","Description":"","Amount":0,"Registered":false,"Type":""},"Time":"2023-02-11T10:14:42.063397+04:00"}








