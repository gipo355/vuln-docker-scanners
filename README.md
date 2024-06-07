# Cli vulnerability docker scanners

This is a docker cli wrapper for various tools to be used in github actions and other CI/CD pipelines

check <https://github.com/gipo355/vuln-docker-scanners-namp-action> for example usages

## nmap

- comes ready with vulners and vulscan scripts
- generate sarif reports to be uploaded to github security tab

## load test: TODO

- uses wrk to generate load tests

### small lib for github

to be separated
sdk available on <https://github.com/octokit/go-sdk> too

# Usage

- requires `--network=host` to run nmap in docker to access the host network

- needs a volume mounted to $workdir to extract reports

- github will pass all environment variables to the docker container, mount the volumes and set the workdir
  _(this is why we don't set workdir by default, to provide compatibility with github actions)_

- if using this container directly as a github action, github won't pass `--network=host` flag, so you need to run the action
  as nodejs and run this container inside the nodejs action with `exec` or `spawn` to pass the `--network=host` flag

```
set workdir with `--workdir` flag
by default it will emit them in $workdir/$report-dir/$report-name/$report-name.ext
changing the workdir will change the location of the reports
example:
`docker run --network=host --workdir=/app --volume .:/app gipo355/vuln-docker-scanners nmap --vulner=true --vulscan=true --target=localhost --port=80 --generate-reports=true --generate-sarif=true`
```

# Notes

we could use the nmap alpine container directly actually, but we need to install the vulners and vulscan scripts.
and we would still need to parse the xml output to generate sarif reports in a separate step.

This lib aims to be an extensible collection of tools to be used in CI/CD pipelines, specifically made for github.

# references and libs

<https://github.com/vmware-tanzu/sonobuoy/blob/main/cmd/sonobuoy/app/delete.go#L38-L58>

nmap xml formatter already available
<https://github.com/vdjagilev/nmap-formatter/wiki/Use-as-a-library>
<https://github.com/vdjagilev/nmap-formatter>

<https://gist.github.com/PSJoshi/1ddb53b42a1b099355df9eac86ced222>

# CREDITS

thanks to @vdjagilev for the nmap formatter lib to parse nmap xml output

Made with cobra, go
