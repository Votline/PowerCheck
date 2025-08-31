# PowerChecker - Minimalist Battery Monitor

PowerChecker is a lightweight battery monitoring application built with OpenGL and GLFW.

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
- **[Go-gl](https://github.com/go-gl/gl)** - graphics rendering
- **[Go-gl/glfw](https://github.com/go-gl/glfw)** (v3.3) - window management

## Installation & Building

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/Votline/PowerChecker
    cd PowerChecker
    ```

2.  **Download dependencies**:
    ```bash
    go mod download
    ```

3.  **Build the application**:
    ```bash
    go build -o powercheck
    ```

4.  **Run**:
    ```bash
    ./powercheck
    ```

## Licenses
This project is licensed under [MIT](LICENSE).

The full license texts are available in the [licenses directory](licenses/)
