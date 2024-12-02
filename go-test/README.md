# Workflow Overview  ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

## Overview
This documentation describes the CI pipeline for the Go service, including its structure and configuration for continuous integration and deployment using GitHub Actions.

## CI Pipeline Workflow

The CI pipeline includes the following steps:
1. **Test**: Runs tests and ensures the code is functional.
2. **Build**: Compiles the application and uploads build artifacts.
3. **Docker Build & Push**: Builds the Docker image and pushes it to Docker Hub.
4. **K8s Manifest Update**: Updates Kubernetes deployment files and pushes changes.

## Pipeline Configuration

### Workflow Definition
The pipeline is triggered by the following events:
- **Push**: Triggered on push to `main` branch.
- **Pull Request**: Triggered on pull request targeting `dev` branch.

### Jobs

#### 1. Test Job
- **Purpose**: Runs tests to ensure the quality of the code.
- **Steps**:
  - Checkout the code.
  - Set up Go environment.
  - Install dependencies (`go mod tidy`, `go get`).
  - Run tests (`go test`).

#### 2. Build Job
- **Purpose**: Builds the application and uploads build artifacts.
- **Steps**:
  - Checkout the code.
  - Build the Go application (`go build`).
  - Upload the build artifact.

#### 3. Docker Build & Push Job
- **Purpose**: Builds and pushes the Docker image to Docker Hub.
- **Steps**:
  - Checkout the code.
  - Set up Docker Buildx.
  - Login to Docker Hub using credentials stored in GitHub Secrets.
  - Read version from the `VERSION` file.
  - Build the Docker image and tag it.
  - Push the Docker image to Docker Hub (only on `main` branch).

#### 4. Kubernetes Manifest Update Job
- **Purpose**: Updates the Kubernetes manifest with the new image version.
- **Steps**:
  - Checkout the code.
  - Read version from the `VERSION` file.
  - Clone the Kubernetes repository and update the image tag in the manifest.
  - Commit and push changes if the manifest was updated.
  - Create a new branch for versioned release and push the branch.


## Development Workflow
- **Main Branch**: Use `main` for Prod.
- **Dev Branch**: Use `dev` for Pre-Prod.
- **Feature Branches**: Use `feature/*` for new features.
- **Bugfix Branches**: Use `bugfix/*` for fixes.
- **Hotfix Branches**: Use `hotfix/*` for urgent fixes.
