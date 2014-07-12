
all:
	go build -o elementary-fixes main.go LaunchpadAPI.go Config.go

format:
	go fmt main.go LaunchpadAPI.go Config.go
