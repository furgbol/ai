# Contributing
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other method with the owners of this repository before making a change.

## Pull Request Process
First of all, you must fork this repository to you GitHub account and change the repository name to any other you want.
Then, follow this steps:
1. Ensure any install or build dependencies are removed before the end of the layer when doing a build.
2. Update the README.md with details of changes to the interface, this includes new environment variables, exposed ports, useful file locations and container parameters.

## Commits Format
You have to specify the package that you are changing and you commit must detail the changing.
Example:
`"package: This is my changing"`

# Project Patterns and Thechniques
In this project you must follow the same patterns and techniques that are addopted:
- S.O.L.I.D. pattern
- Singleton pattern
- Dependency injection technique

## Style Guide
In this project we follow the [Uber Golang Style Guide](https://github.com/uber-go/guide).

## Version Nomenclature
The nomenclature of versions of this project is its own standard: MAJOR.MINOR.

- **MAJOR**: Identifier of the main version in which the work is being developed.
- **MINOR**: Improvement identifier in the version in which the work is being developed.

### Examples:
- Version 1.0: Beginning of the implementation of the version oriented to the [VSS SDK Simulator](https://github.com/VSS-SDK/VSS-Simulator).
- Version 1.1: Enhancement to control robots in the simulator via joystick.