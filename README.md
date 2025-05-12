# Minecraft Java Server Helper

A Go CLI utility to simplify the setup and management of Minecraft Java Edition servers.

## Features

- Download Minecraft server files (vanilla, Paper, Spigot, etc.)
- Install and configure mod loaders (Fabric, Forge)
- Download and manage mods from Modrinth
- Generate server configuration files
- Manage server startup and shutdown
- Track server performance
- Backup management

## Installation

### Prerequisites

- Go 1.19+
- Java 17+ (for running Minecraft servers)

### Installing from source

```bash
git clone https://github.com/yourusername/minecraft-java-server-helper.git
cd minecraft-java-server-helper
go install
```

## Usage

```bash
# Create a new vanilla server
mcjsh new vanilla --version 1.20.1 --name my-server

# Create a new server with Fabric mod loader
mcjsh new fabric --version 1.20.1 --loader-version 0.14.21 --name fabric-server

# Download mods from Modrinth
mcjsh mods add lithium sodium

# Start the server
mcjsh start --server my-server --ram 4G

# List all servers
mcjsh list

# Backup a server
mcjsh backup --server my-server
```

## Configuration

The CLI uses a configuration file located at `~/.config/mcjsh/config.yaml` by default. You can specify a different configuration file using the `--config` flag.

Example configuration:

```yaml
servers_directory: "~/minecraft-servers"
java_path: "/usr/bin/java"
default_ram: "2G"
```

## License

MIT
