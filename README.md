# GearX
A shopping website centered around computer, smartphone, and technology in general.

## Dependencies
- go 1.25 or later

## How to run this website from source
1. Clone this repository and install Go dependencies:
```bash
git clone https://github.com/islabre1426/GearX.git && cd GearX
go mod tidy
```

2. To start development server, execute:
```bash
go tool air
```

3. To build this website to single binary, execute:
```bash
go build
```

The result will be `gearx.exe` or `gearx`, depending on the Operating System.

To run the binary, execute:
```bash
.\gearx.exe
```

on Windows or
```bash
./gearx
```

on other OSes.