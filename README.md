# Cli vulnerability docker scanners

this is a docker cli wrapper for nmap to be used in github actions and other CI/CD pipelines

check <https://github.com/gipo355/vuln-docker-scanners-namp-action> for example usages

## nmap

comes ready with vulners and vulscan scripts

generate sarif reports to be uploaded to github security tab

## load test

# todo

- add licenses
- fork vulner repo
- copy vulner scripts in continer
- add more nmap options
- create auto release and docker publish on tag
- generate sarif report
- add cobra and viper
- split github utils in its own lib
- understand golang versioning forlib
- split smiattack cli with nmap in its own repo for action using the github lib and utils

- add load test

- make a list of all go libs and tools

# splits

cli nmap repo

- provides commands like scan, list, etc
- provides writeToFile, writeToSarif, output to console

utils repo

- provides github utils like getRepo, getIssue, etc

# note

requires `--network=host` to run nmap in docker

also needs a volume mounted to $workdir to extract reports

set workdir with `--workdir` flag

_note we don't set workdir by default to provide compatibility with github actions_

by default it will emit them in $workdir/$report-dir/$report-name/$report-name.ext

changing the workdir will change the location of the reports

example:

we encapsulate the nmap command to be able to extend this cli with more programs later on

`docker run --network=host --workdir=/app --volume .:/app gipo355/vuln-docker-scanners nmap --vulner --vulscan --target=localhost --port=80 --generate-reports --generate-sarif`

# references and libs

<https://github.com/vmware-tanzu/sonobuoy/blob/main/cmd/sonobuoy/app/delete.go#L38-L58>

nmap formatter already available
<https://github.com/vdjagilev/nmap-formatter/wiki/Use-as-a-library>
<https://github.com/vdjagilev/nmap-formatter>

<https://gist.github.com/PSJoshi/1ddb53b42a1b099355df9eac86ced222>

# CREDITS

thank you @vdjagilev for the nmap formatter lib
