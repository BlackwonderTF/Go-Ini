# Go-Ini

Mostly based on [this](https://en.wikipedia.org/wiki/INI_file) WikiPedia article.

## Features

None yet

## Supported value types

None yet

## Planned Features

- Section
- Section Nesting
  ```ini
  [Section]
  key=value
    [SubSection1]
    key=value
    [SubSection2]
    key=value
  ```
- Comments (# ;)
- Global Properties (http://commons.apache.org/proper/commons-configuration/apidocs/org/apache/commons/configuration2/INIConfiguration.html)

  ```ini
  globalkey=value

  [section]
  sctionkey=value
  ```

- Multi-line
- Read from file
- Write to file

## Planned Supported value types

- String
- Bool
- Int (0-64)
- Uint (0-64)
- Array (Types above)
