# Go-Ini

This implementation is based on these articles:
- [WikiPedia](https://en.wikipedia.org/wiki/INI_file) 
- [Apache](http://commons.apache.org/proper/commons-configuration/apidocs/org/apache/commons/configuration2/INIConfiguration.html)

## Features

- Section
- Global Properties
- Read from file
- Multiple seperators [":", "="]
- Comments above line and in-line (#) (;)
- Section Nesting
- Configuration options for features with multiple implementations (like sub-sections)

## Supported value types

- String
- Bool
- Int (0, 8, 16, 32, 64)
- Uint (0, 8, 16, 32, 64)

## Planned Features

- Write to file
- Hot reload/write

## Planned Supported value types

- Array (Types above)
- Multi-line String / Array
