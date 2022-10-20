# go-nmcli

go wrapper for command line tool `nmcli`.

## Features

The library can be used in a `nmcli` like scheme.

### General

| original command            | library path                     | implemented                   |
|-----------------------------|----------------------------------|-------------------------------|
| `nmcli general status`      | `NMCli.General.Status(...)`      | :negative_squared_cross_mark: |
| `nmcli general hostname`    | `NMCli.General.Hostname(...)`    | :heavy_check_mark:            |
| `nmcli general permissions` | `NMCli.General.Permissions(...)` | :heavy_check_mark:            |
| `nmcli general logging`     | `NMCli.General.Logging(...)`     | :negative_squared_cross_mark: |

### Networking

| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli networking ...`              | not implemented yet              | :negative_squared_cross_mark: |

### Radio

| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli radio ...`                   | not implemented yet              | :negative_squared_cross_mark: |

### Device

| original command                  | library path                 | implemented                   |
|-----------------------------------|------------------------------|-------------------------------|
| `nmcli device status`             | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device show`               | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device set`                | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device reapply`            | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device modify`             | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device disconnect`         | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device wifi list`          | `NMCli.Device.WiFiList(...)` | :heavy_check_mark:            |
| `nmcli device wifi connect`       | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device wifi hotspot`       | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device wifi rescan`        | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device wifi show-password` | not implemented yet          | :negative_squared_cross_mark: |
| `nmcli device wifi lldp`          | not implemented yet          | :negative_squared_cross_mark: |

### Agent

| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli agent ...`                   | not implemented yet              | :negative_squared_cross_mark: |

### Monitor

| original command                    | library path                     | implemented                   |
|-------------------------------------|----------------------------------|-------------------------------|
| `nmcli monitor ...`                 | not implemented yet              | :negative_squared_cross_mark: |

## Usage
