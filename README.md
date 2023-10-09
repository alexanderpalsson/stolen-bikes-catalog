## How to use

### Run application
`Make run` will start the server on port 8080.

### API specification
Found in ./api directory

### Implementation
Although the requirements initially called for a SQL database, a simple NoSQL MongoDB database was used due to my limited experience with SQL.

The implementation follows a straightforward REST handler structure, comprising two primary entities: "officers" and "reports." The project layout is modular, with each entity residing in separate packages, each equipped with its own storage and handler packages.

**Officer API**: This API allows for the creation, updating, listing, and deletion of officers. Whenever a new officer is created, the system will attempt to assign them to the oldest open case.

**Report API**: This API supports the creation, listing, and resolution of stolen bike reports. Upon report creation, the system will endeavor to assign an officer to the case. When a case is marked as resolved, the system will assign an officer to the oldest pending case if one exists.

It's important to note that due to time constraints, several planned features remain unimplemented. These pending tasks include:

TODO:

* Lack of Testing: The codebase currently lacks testing, and it should ideally include both unit and integration tests.
* Absence of Authentication and Access Control: At present, anyone can resolve bike cases. Implementing a simple JWT-token-based authentication system would enhance security and access control.