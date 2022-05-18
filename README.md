# opencga-cli 

Usage

```
./bin/opencga get version -s iva.mseqdr.org
Fetching from https://iva.mseqdr.org/iva/conf/config.js
        IVA Version: v2.2.0
        Cellbase Version: 5.0.1                 Git Commit: eaae3a6f7b407c1eebdb1b4bfede941f4b506b30    Git Branch: release-5.0.x
        OpenCGA Version: 2.2.1-SNAPSHOT         Git Commit: 27cf2ae4bb95596daf839f107dac3d8fb63e6715    Git Branch: release-2.2.x
```

Build
```
git clone git@github.com/phamidko/opencga-cli
cd opencga-cli
go mod init github.com/phamidko/opencga-cli
go get -u github.com/spf13/cobra@v1.4.0
go get -u github.com/mvdan/xurls
make build
```


Help

```
./bin/opencga -h
opencga cli

Usage:
  opencga [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get type of resource to query
  help        Help about any command

Flags:
  -h, --help   help for opencga

Use "opencga [command] --help" for more information about a command.
```


Optional
```
go install github.com/spf13/cobra-cli@latest
```