# Reader Package

The `reader` package offers utilities to read and process load data from files.

## Features

- **LoadsFileReader Interface**: A generalized interface that specifies the method to read loads from files.

- **loadsTextFileReader Implementation**: A concrete implementation of the `LoadsFileReader` interface that reads loads data specifically from text files.

- **Factory Function - NewLoadsFileReader**: This is a factory function that instantiates and returns a new instance of `loadsTextFileReader`, allowing for easy creation of new file reader objects

## How it Works

The `reader` package provides a `loadsTextFileReader` type which is responsible for reading the load data from a given text file. The text file is expected to have a specific header format, represented by `loadNumber pickup dropoff`. Each subsequent line within the file should describe a load with its respective load number, pickup location, and dropoff location.

Coordinates should be provided in the format `(latitude,longitude)`. If the file format deviates from this expected format, errors will arise during the reading process.

### Sample Text File Format

```
loadNumber pickup dropoff
1 (40.7128,74.0060) (34.0522,118.2437)
2 (51.5074,0.1278) (48.8566,2.3522)
...
```


