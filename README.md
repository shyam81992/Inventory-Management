# Inventory-Management


Dependencies

Golang 1.13 or higher

Docker for Mac/Windows

github.com/shyam81992/Inventory-Management-job // Which is used to notify the shipments.

Steps to run the project 

    1. Run go mod tidy to install the dependencies.
    2. Open a new terminal and go to the folder postgres 
    3. Run the command docker-compose up (To start the postgres db)
    4. Open a new terminal and go to the folder rabbitmq and run the command docker-compose up (To start the Rabbitmq)
    5. clone the github.com/shyam81992/Inventory-Management-job in a separate folder and run the command 
        docker-compose up
    6. open a new terminal and got to the Inventory-Management project root folder and run the command 
        docker-compose up
    

Inventory-Management-Job
Listens to the shipment notifications. (Kinesis or kafka, sqs, rabbitmq)

 
