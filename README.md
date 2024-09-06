# My Go Template
An opinionated template for building modular monoliths or microservices

## Features
- Chi Router for HTTP based endpoints
- Zerolog for logging capabilities
- Koanf for configuration, supports files and env vars
- Makefile with the most common tasks

## Getting Started
- Clone the repository or download the code
- Change the module name in the go.mod file

## Folder Structure
- cmd/main.go - main entry point
- internal - application specific business logic
- pkg - common or shared packages that might be used in multiple modules
- resources - application specific resources, such as config files, databases, etc.
