# Flyweight Pattern in Go

## Overview
This repository demonstrates the application of the Flyweight design pattern in Go. The project highlights how to minimize memory usage by sharing as much data as possible with other similar objects, showcasing flexibility and best practices in design patterns and unit testing.

## Pattern Description
The Flyweight pattern is used to minimize memory usage by sharing data among similar objects. This is achieved by storing common data externally and passing it to the objects that need it. In this project, we have implemented the Flyweight pattern using a `FlyweightFactory` to manage and reuse flyweight objects.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the creation and usage of flyweights.
- **pkg/flyweight/**: Houses the Flyweight implementation.
  - **flyweight.go**: Defines the `Flyweight` interface.
  - **concrete_flyweight.go**: Implements the `ConcreteFlyweight`.
  - **flyweight_factory.go**: Implements the `FlyweightFactory` to create and manage flyweight objects.
  - **concrete_flyweight_test.go**: Unit tests for `ConcreteFlyweight`.
  - **flyweight_factory_test.go**: Unit tests for `FlyweightFactory`.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).


### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Flyweight_Pattern.git
cd Go_Flyweight_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp