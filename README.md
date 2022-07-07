# opencga-cli 

Download the binary file
```
# Linux 
wget https://github.com/phamidko/opencga-cli/releases/download/v0.1.0-alpha/opencga_linux_amd64.tar.gz
tar -xvf opencga_linux_amd64.tar.gz 
```

Usage

```
./bin/linux/opencga get version -s uat.eglh.app.zettagenomics.com
Fetching from https://uat.eglh.app.zettagenomics.com/iva/conf/config.js
        IVA Version: v2.2.0-dev
        Cellbase Version: 5.0.1         Git Commit: eaae3a6f7b407c1eebdb1b4bfede941f4b506b30    Git Branch: release-5.0.x
        OpenCGA Version: 2.3.1-SNAPSHOT Git Commit: e6d0ec12831e9e8b5529f9bcc6a3b5516b19d6cc    Git Branch: release-2.3.x


./bin/linux/opencga get version -s http://iva.mseqdr.org
Fetching from https://iva.mseqdr.org/iva/conf/config.js
        IVA Version: v2.2.0
        Cellbase Version: 5.0.1                 Git Commit: eaae3a6f7b407c1eebdb1b4bfede941f4b506b30    Git Branch: release-5.0.x
        OpenCGA Version: 2.2.1-SNAPSHOT         Git Commit: 27cf2ae4bb95596daf839f107dac3d8fb63e6715    Git Branch: release-2.2.x


# Full path to config.js file
./bin/linux/opencga get version -s http://bioinfo.hpc.cam.ac.uk/web-apps/iva-prod/conf/config.js
Fetching from http://bioinfo.hpc.cam.ac.uk/web-apps/iva-prod/conf/config.js
        IVA Version: v2.2.0-dev
        Cellbase Version: Not Found     Git Commit: Not Found   Git Branch: Not Found
        OpenCGA Version: Not Found Git Commit: Not Found        Git Branch: Not Found

# Full path to config.js file
./bin/linux/opencga get version -s uat.eglh.app.zettagenomics.com/iva/conf/config.js
Fetching from https://uat.eglh.app.zettagenomics.com/iva/conf/config.js
        IVA Version: v2.2.0-dev
        Cellbase Version: 5.0.1         Git Commit: eaae3a6f7b407c1eebdb1b4bfede941f4b506b30    Git Branch: release-5.0.x
        OpenCGA Version: 2.3.1-SNAPSHOT Git Commit: e6d0ec12831e9e8b5529f9bcc6a3b5516b19d6cc    Git Branch: release-2.3.x
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