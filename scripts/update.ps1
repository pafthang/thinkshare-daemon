Remove-Item ./secret/config.json
git pull
git checkout master
git submodule update --init --recursive

# build GO 
go build  -o daemon.exe ./cmd/daemon/
go build  -o cli.exe ./cmd/cli/

Set-Location .\hub
go build -o hub.exe  ./cmd/server/
Set-Location ../

# build .NET
Set-Location .\hid
dotnet build . --output "bin" --self-contained true --runtime win-x64
Set-Location ..
