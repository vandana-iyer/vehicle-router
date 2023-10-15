# Vehicle Router Project

## Introduction

The `Vehicle Router` project is an application designed to effectively schedule and optimize load assignments to drivers. The project encompasses proximity-based strategies for efficient vehicle routing.

## Prerequisites

- Go (latest version recommended)
- An input file with load data for testing purposes

## Running the Project

1. **Navigate to the Project's Root Directory**:
    ```bash
    cd path_to_directory/vehicle-router
    ```

2. **Run the Application**:
   Provide the path to your load input file when running the project.
    ```bash
    go run cmd/main.go <input_file>
    ```
   The input file needs to be in the below format
      
   ```bash
   loadNumber pickup dropoff 
   1 (-50.1,80.0) (90.1,12.2) 
   2 (-24.5,-19.2) (98.5,1,8) 
   3 (0.3,8.9) (40.9,55.0)
   4 (5.3,-61.1) (77.8,-5.4)
   ```
   



## Testing the Project

All tests in the module can be run by using the following command from the root directory:

```bash
go test ./...