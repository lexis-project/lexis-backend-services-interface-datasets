# lexis-backend-services-interface-datasets

<a href="https://doi.org/10.5281/zenodo.6080494"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.6080494.svg" alt="DOI"></a>

This repository contains the interface to display datasets from the LEXIS DDI in the LEXIS Portal.

## Acknowledgement

This code repository is a result / contains results of the LEXIS project. The project has received funding from the European Union’s Horizon 2020 Research and Innovation programme (2014-2020) under grant agreement No. 825532.

## Notes on compiling and installing this software.

The swagger command used in server/make_server.sh and client/make_client.sh
is not the standard swagger.
Instead, use the swagger for go, available at https://github.com/Stratoscale/swagger.
The current (25.03.2020) version does not compile (tests fail); however, the swagger docker image has been created by the time of the fail.
You will need to use the fish shell and setup the swagger alias in fish before running the scripts in server/ and client/ using source.
Example:
fish
export GOPATH=$HOME/gocode
alias swagger="docker run --rm -e GOPATH=$GOPATH -v $HOME:$HOME -w (pwd) -u (id -u):(id -g) stratoscale/swagger:v1.0.9"
swagger generate server
swagger generate client

You will need a recent enough version of go.
- go version go1.10.4 linux/amd64 (from Ubuntu 18.04.4) is known NOT to work
- go version go1.13 works.
You will also need GIT_TERMINAL_PROMPT=1 and GOPRIVATE="code.it4i.cz"

Once you regenerate the client or the server, you will need to run the scripts/modsfix.sh script on the base directory of the repository.

### Regarding Swagger:
- version: 0.13.0 commit: 8135eb6728e43b73489e80f94426e6d387809502 (stratoscale/swagger:v1.0.9) generates wrong code.
- stratoscale/swagger:v1.0.27 works
- quay.io/goswagger/swagger:latest is the recommended version as of June 2020.

- Add --template=stratoscale if you are using quai.io goswagger

### Regarding front-ends:
You want an HTTPS front-end to access this service.

### Apache
If you use Apache2:
When using an apache2 front-end for this service, you will need to be aware of some caveats:
- File upload to the DDI can have relatively large requests
- The apache2 proxy module had a bug where large requests would produce 502 instead of the correct error from the proxied service if the proxied service responded before swallowing the whole request (behaviour allowed by standard)
- The bug was fixed in apache2 2.4.40 (see https://bz.apache.org/bugzilla/show_bug.cgi?id=60330)
- The bug is present in 2.4.29-1ubuntu4.13, the latest version available in Ubuntu 18.4 repositories

Please use https://websiteforstudents.com/install-the-latest-apache2-2-4-34-on-ubuntu-16-04-17-10-18-04-lts-servers/ to upgrade apache2 to the latest version if you are affected by the bug.

See scripts/apache2 for an example setup

This issue affects all WP8 backend repositories, not only this one.
