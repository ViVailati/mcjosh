# TODO List for Minecraft Java Server Helper

## Core Structure
- [ ] Set up Go project structure
- [ ] Create command-line framework using [Cobra](https://github.com/spf13/cobra) or [Urfave CLI](https://github.com/urfave/cli)
- [ ] Implement configuration management with [Viper](https://github.com/spf13/viper)
- [ ] Design core data structures and interfaces
- [ ] Set up logging system

## Server Management Features
- [ ] Implement server version detection from Mojang API
- [ ] Create download functionality for vanilla server jars
- [ ] Add support for different server types (Paper, Spigot, Purpur, etc.)
- [ ] Implement server creation wizard with configurations
- [ ] Create eula.txt auto-generation
- [ ] Handle server properties configuration

## Mod Loader Support
- [ ] Implement Fabric API integration
- [ ] Implement Forge API integration
- [ ] Add mod loader version detection and compatibility checking
- [ ] Create mod loader installation pipeline

## Mod Management
- [ ] Implement [Modrinth API](https://docs.modrinth.com/api-spec/) integration
- [ ] Add mod search functionality 
- [ ] Create mod installation system
- [ ] Implement mod dependency resolution
- [ ] Add mod update checking
- [ ] Create mod configuration management

## Server Operations
- [ ] Implement server start/stop/restart commands
- [ ] Create server status monitoring
- [ ] Add memory usage tracking
- [ ] Implement backup and restore functionality
- [ ] Add world management features
- [ ] Create plugin/mod hot reloading

## Testing
- [ ] Write unit tests for core functionality
- [ ] Create integration tests for server management
- [ ] Set up CI/CD pipeline

## Documentation
- [ ] Create detailed usage documentation
- [ ] Add examples for common use cases
- [ ] Generate command reference documentation

## Distribution
- [ ] Set up build system for multiple platforms
- [ ] Create installation scripts
- [ ] Implement auto-update functionality
- [ ] Set up release workflow
