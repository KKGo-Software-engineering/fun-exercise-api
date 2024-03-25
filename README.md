# Fun Exercise: API Design and DevOps

# Prerequisites

- [Go](https://go.dev/doc/install)
	- Mac: `brew install go` [Homebrew](https://brew.sh/) or visit [Go](https://go.dev/doc/install)
	- Windows: `choco install go` [Chocolatey](https://chocolatey.org/install)
	- Linux: `sudo apt-get install golang-go`
    - Verify installation: `go version`
- [Docker](https://docs.docker.com/get-docker/)
- [Swag](https://github.com/swaggo/swag) - Generate Swagger Documentation
  - Install swag: `https://github.com/swaggo/swag`

# Getting Started
1. Clone the repository
2. Open fun-exercise-api-design-and-devops in your favorite Editor
3. Run the following command to start the server
	```bash
	go run main.go
	```
4. Open your browser and navigate to [http://localhost:1323/api/v1/wallets](http://localhost:1323/api/v1/wallets)
5. You should see a list of wallets
6. View Swagger documentation at [http://localhost:1323/swagger/index.html](http://localhost:1323/swagger/index.html)
7. You should see the Swagger documentation for the API

<img src="./swagger.png" alt="Swagger Documentation" />

## Table of Contents
- [Challenge 0: Starter Code - Display a list of wallets](#challenge-0-display-a-list-of-wallets)
- [Challenge 1: API - Using environment variables](#challenge-1-api---using-environment-variables)
- [Challenge 2: API - Write Unit Test for /ap/v1/wallets](#challenge-2-api---write-unit-test-for-apv1wallets)
- Challenge 3: API - Using Query Parameters
- Challenge 4: API - Using Path Parameters
- Challenge 5: API - Using Request Body to Create a Wallet
- Challenge 6: API - Using Request Body to Update a Wallet
- Challenge 7: API - Using Request Body to Delete a Wallet
- Challenge 8: API - Using Request Body to Transfer Funds
- Challenge 9.0: DevOps - Dockerize the App - Single Stage Dockerfile
- Challenge 9.5: DevOps - Dockerize the App - Multi-Stage Dockerfile
- Challenge 10: DevOps - Design a CI for running static code analysis and tests

### Challenge 0: Display a list of wallets

### Challenge 1: API - Using environment variables

### Challenge 2: API - Write Unit Test for /ap/v1/wallets
