# About

### Structure
Basically this application can be divided into three parts: __main application__, __services__ and __data stores__.

  - __main application__ The main application uses `gorilla mux` to handle http requests. It calls different `MakeEndpoint` methods defined in __services__ to create corresponding handler and handle incoming requests.

  - __services__ Service part is the middle layer between outside requests and the inner database services. It basically defines handlers that decode raw http requests:  
    - call __data store__ functions to do operations on the database
    - deal with errors
    - encode output for http response


  - __data stores__ The data stores uses SQL queries to communicate with database directly. It receives params passed from __services__ layers and abstract out data to form standard SQL queries. It also handles errors and wrap query results with formatted models for outside use.

# Get started

### Application setup
To Get started with the project, you need to have the following done first:

  1. Database setup
    - Install [PostgreSQL](https://www.postgresql.org/).

    - Create database based on scripts provided in `utils/sql/`. Please create two databases, one for normal development, one for testing (recommended).

  2. Setup golang environment.
    - Download golang from [here](https://golang.org/dl/).

  3. Application setup
    - Clone the project and copy it under `$GOPATH/`. Usually `$GOPATH` is defined when you install golang. By default it is `/Users/<your-user-name>/go` for macOS. If you want to change your workspace, you might need to set `$GOPATH` environment variable. See [here](https://github.com/golang/go/wiki/SettingGOPATH).

    - Ensure dependency: use `brew` to install `dep`:
      ```
      $ brew install dep
      $ brew upgrade dep
      ```

      Then under root directory, run

      ```
      $ dep init
      $ dep ensure
      $ dep ensure -update
      ```

    - Setup development config: please proceed to root directory and run
      ```
      $ ./setup
      ```
      and please adjust your settings accordingly.

    - Run your application! Please run the following command under your application's root directory
      ```
      $ go run main.go
      ```

### Test your code
Please remember to test your code after development. testing files can be found inside `tests/`. You can run
```
$ go test ./...
```
under root directory to test your code.

### Code formatting
In order to have better coding experience, please follow the following coding format rules (will be adding along the way).

  - __naming__ Please use CamelCase for all names. For public variable/constants/structs/methods, please capitalize the first letter.
  - __packages__ Please name your packages all in lower cases, no underlines.

Note: you before pushing your code please run
```
$ go fmt ./...
```
This will format your code.
