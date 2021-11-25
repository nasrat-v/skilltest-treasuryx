# skilltest-treasuryx
Initial repo for treasuryX skill test

# Build
go build

# Start
./skilltest-treasuryx

# .Env
A .env file is provided. Feel free to modify values.

## Approach
I choose to create a Web server with gin-gonic library for it's efficiency and it's facility to use
The code is splinted into differents packages. Each package concentrates a precise part of the project.
The  project is managed by a ServiceManager class. This class allows to initialize in the right order the different services. Such as database, router, server and controller.
The server manages the sending of several requests. You just have to remember to update the bank's response file with the payments IDs and Status.

### Controller
The controller contains only one route and therefore only one handler.
The authentication is done via a middleware allowing not to reach the handler if the request is not authenticated.
The body validation is done directly in the json binding thanks to the "validator" library.
The only one route is "/api/payment".

### Database
The architecture of the database is as follows:
We have two tables, one "Account" and one "Payment".
For each payment request we will register the bank account.
Then we will link the payments and the accounts through their IDs.
This will allow us to trace the creditor and debtor informations for each new payment.

### Payment
Each payment is unique and must have a unique ID to not be refused.
For each payment, a file containing the payment ID + "_payment.xml" is created. For example for the ID "JXJ984XXXZ" we have the file "JXJ984XXXZ_payment.xml"

### Bank
After each payment a goroutine is launched and will wait for the answer from the bank.
That is, it will not finish until the file exists or until the payment id is inserted in the response file.
Once found, the status is updated in the database.
You can make several request and update the bank response file after. The status will be updated.

### Limits
If the user spams the server and the bank does not update the response file, the server will overload because of the goroutines that do not terminate