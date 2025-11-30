# Chapter 1: Setup
- `go mod init github.com/dbrucknr/go-design-patterns`
- `mkdir cmd`
- `mkdir cmd/web` - this is a convention in Go for projects that serve web pages
    - If you create an API it goes in the `cmd/api` directory.
- `touch cmd/web/main.go` - Here we will create a simple web server using the net/http package.
- Adding a routing package: [Chi](https://github.com/go-chi/chi)
  - Install command: `go get -u github.com/go-chi/chi/v5`
- Lets add a routes.go file to define our routes
- We will also take advantage of another Chi package to handle middleware.
- Run: go mod tidy to eliminate warning
- Lets begin to think about templates
- mkdir templates
- touch templates/base.layout.gohtml
- add a `render.go` file to handle template rendering
- These changes will warrant command line params for using template caching
- Lets add in support for static assets
- `mkdir static`
- I added in a large image of puppies that's approximately 1MB in size.
- I'm going to try and compress the image to reduce its size using: https://pixelied.com/convert.
  - Convert from jpg to webp
  - Set the quality to 20

# Chapter 2: Factory Pattern + Abstract Factory Pattern
- Factory Pattern: Create an instance of an object with sensible default values.
- Abstract Factory Pattern: Creates families of related or dependent objects without relying on their concrete classes.

- Lets setup some models / types in the project.
- I added a toolkit from the course: `go get -u github.com/tsawler/toolbox` for parsing JSON.
- Visit: http://localhost:4000/test-patterns to test the factories and the handlers that serve them.

# Chapter 3: Repository Pattern
- Repository Pattern: Provides an abstraction layer between the application's business logic and the data source.
  - Allows us to change databases with ease.
  - Makes writing unit tests much simpler
- We will use MariaDB (A MySQL database)
- We need a driver to connect to MariaDB
  - https://github.com/go-sql-driver/mysql
  - `go get -u github.com/go-sql-driver/mysql`
- We need to launch a Database, which will be done using Docker.
  - `touch compose.yml`
mkdir sql
touch sql/breeders_mysql.sql
  - I copied a data dump from the course.
Run: `docker compose up -d`
Add a db.go file to handle database connection initialization. Use it in the main function.
Then run: `go run ./cmd/web/` to verify connection worked (after adding db.go)

How this works in the app:

In /models/models.go we have a Factory method called `New`
- It accepts a type. If a non nil value is passed, the app will use a standard database repository.
- However, if a nil value is passed, we set the package var to a TestRepository, which will satisfy all Repository interface requirements, but won't require the test context to use a real database connection.

We satisfy the Repository interface for the test context in the file: `/models/dogs_testDB.go`

in `/cmd/web/setup_test.go` we pass a nil value to set the package level variable repo to a TestRepository instance.

Then, in handlers_test.go our test is setup to use the TestRepository instance, which enables to bypass a real database and check the response based on the app level context required to access the handler.

# Chapter 4: Singleton Pattern
- Allows us to restrict the instantiation of a class to one / singular instance.
- The pattern is useful when exactly one object is needed to coordinate actions across a system.
