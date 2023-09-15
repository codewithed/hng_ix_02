## Local Setup
### Prerequisites
- Docker
- Postgres docker image
### Steps
- Open the project directory
- Setup environment variables in accordance with the .env.example file provided
- Run "make postgres"
- Run "make createdb"
- Run "make migrateup"
  
## API Docs
### Assumptions/Limitations made in development
- Each name stored in the database is unique to prevent potential errors when fetching a resource
- Only the age of a person can be updated to prevent changing it to name already taken

### Accepted HTTP Verbs
* To create a person resource: POST
* To read a person resource: GET
* To update a person resource: PUT/PATCH
* To delete a person resource: Delete

### API request guide
- POST request:
  * Accepts only JSON format
  * Url example: https://localhost:3000/api
  * Request body example:
    {
      "name": "Edmund Debrah",
      "age": 20
    }
  * Response example:
    {
      "id": 2,
      "name": "Edmund Debrah",
      "age": 20
    }
    
- GET request:
    * Url example: https://localhost:3000/api/name
    * Response example: {
      "id": 2,
      "name": "Edmund Debrah",
      "age": 20}
      
- PUT/PATCH request:
    * Accepts only JSON format
    * Url example: https://localhost:3000/api/name
    * Request body example:
      {
        "age": 19
      }
    * Response example:
      {
        "id": 2,
        "name": "Edmund Debrah",
        "age": 19
      }
      
- DELETE request:
    * Url example: https://localhost:3000/api/name
    * Response example: {
  "message": "deleted person: Edmund Debrah"
}
     
## Database Model
https://dbdiagram.io/d/64fe37df02bd1c4a5e4cbe40
