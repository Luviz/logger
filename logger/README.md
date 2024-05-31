# simple-go-tools logger

This is a simple logger, I'm writing this to learn Go.

## Log struct

| name      | type         | description                 |
| --------- | ------------ | --------------------------- |
| ID        | number       | PK                          |
| timeStamp | number       |                             |
| source    | str          | source of the log           |
| message   | str          |                             |
| value     | number/float | used in metrics default (1) |
| tags      | JSON         | any key pair                |
| logLevel  | enum         |                             |

### logLevel

1. trace
2. debug
3. info
4. waring
5. error


## Features


### TODO

- [ ] Color full dev output
- [ ] Write to CSV
- [ ] Write to DB
- [ ] tags
- [ ] support for metrics
- [ ] Configs
  - [ ] modify log template
  - [ ] connect to DB