# go-nmcli

go wrapper for command line tool `nmcli`.

## Features

The library can be used in a `nmcli` like scheme.

### General
| original command                                  | library path                        | implemented                   |
|---------------------------------------------------|-------------------------------------|-------------------------------|
| `nmcli general status`                            | `NMCli.General.Status(...)`         | :negative_squared_cross_mark: |
| `nmcli general hostname`                          | `NMCli.General.Hostname(...)`       | :heavy_check_mark:            |
| `nmcli general hostname <hostname>`               | `NMCli.General.ChangeHostname(...)` | :heavy_check_mark:            |
| `nmcli general permissions`                       | `NMCli.General.Permissions(...)`    | :heavy_check_mark:            |
| `nmcli general logging`                           | `NMCli.General.Logging(...)`        | :negative_squared_cross_mark: |
| `nmcli general logging <log level> <log domains>` | `NMCli.General.ChangeLogging(...)`  | :negative_squared_cross_mark: |

### Networking
| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli networking ...`              | not implemented yet              | :negative_squared_cross_mark: |

### Radio
| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli radio ...`                   | not implemented yet              | :negative_squared_cross_mark: |

### Device
| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli device ...`                  | not implemented yet              | :negative_squared_cross_mark: |

### Agent
| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli agent ...`                   | not implemented yet              | :negative_squared_cross_mark: |

### Monitor
| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli monitor ...`                 | not implemented yet              | :negative_squared_cross_mark: |

## Usage
