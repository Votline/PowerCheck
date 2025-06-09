# PowerCheck - Minimalist Battery Monitor

PowerCheck is a lightweight battery monitoring application built with OpenGL and GLFW.

## Description

The application displays current battery level.

### Key Features:
- **Battery percentage**: Accurate reading from `/sys/class/power_supply/BAT0/capacity`
- **Instant actions**: One-click shutdown or suspend
- **Lightweight**: Minimal resource consumption
- **Silent mode**: Run with `--smode=true` for background operation
- **SD** - Immediate shutdown
- **SS** - Suspend system

## Technologies
- **Go** (1.24.3) - primary language
- **OpenGL** - graphics rendering
- **GLFW** (v3.3) - window management

## Installation
```bash
git clone https://github.com/Votline/PowerCheck && cd PowerCheck 
```
```bash
go mod download
```

```bash
go build
```
