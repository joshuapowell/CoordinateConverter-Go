# Coordinate Converter for Go

A simple coordinate converter, intended to be a relatively simple first 
program to write in the Go language.

## Usage

### Conversion

#### North Coordinate
To convert a degrees-minutes-seconds or DMS coordinate of 30° 15' 50" N to
it's decimal degree equivalent.
```
ConvertDmsToDd(30, 15, 50, North)

    returns 30.26388888888889
```

#### East Coordinate
To convert a degrees-minutes-seconds or DMS coordinate of 57° 45' 02" E to
it's decimal degree equivalent.
```
ConvertDmsToDd(57, 45, 02, East)

    returns 57.75055555555556
```

#### South Coordinate
To convert a degrees-minutes-seconds or DMS coordinate of 63° 19' 15" S to
it's decimal degree equivalent.
```
ConvertDmsToDd(63, 19, 15, South)

    returns -63.32083333333334
```

#### West Coordinate
To convert a degrees-minutes-seconds or DMS coordinate of 57° 45' 02" W to
it's decimal degree equivalent.
```
ConvertDmsToDd(57, 45, 02, West)

    returns -57.75055555555556
```

### Testing
```
go test
```

