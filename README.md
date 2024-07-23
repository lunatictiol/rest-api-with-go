## README

### Event Management API

**Overview**
This Go REST API provides endpoints for managing events. Users can register, log in, create events, and view event details. The API utilizes the Gin framework for efficient routing and JWT authentication for secure user sessions.


**API Endpoints**
* **Authentication**
  * `/signup`: Register a new user
  * `/login`: Authenticate a user and receive a JWT token
* **Events**
  * `/events`: Get a list of all events
  * `/events/:id`: Get details of a specific event
  * `/events`: Create a new event (requires authentication)
  * `/events/:id`: Update an event (requires authentication and ownership)
  * `/events/:id`: Delete an event (requires authentication and ownership)

**Authentication**
The API uses JWT authentication. Include the JWT token in the `Authorization` header with the prefix `Authorization` for protected endpoints.

**Data Structure**
* **User:**
  * ID 
  * Email
  * Password (hashed)
* **Event:**
  * ID 
  * Title
  * Description
  * date_time
  * Location
  * user ID (foreign key to User)

**Additional Notes**
* Error handling: The API should return appropriate error codes and messages.
* Input validation: Validate user input to prevent invalid data.
* Security: Implement security best practices, including input sanitization, password hashing, and protection against common vulnerabilities.
