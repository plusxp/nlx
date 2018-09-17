<h1><img alt="NLX" src="logo.png" width="200"></h1>

[![pipeline status](https://gitlab.com/commonground/nlx/badges/master/pipeline.svg)](https://gitlab.com/commonground/nlx/commits/master)  [![coverage report](https://gitlab.com/commonground/nlx/badges/master/coverage.svg)](https://gitlab.com/commonground/nlx/commits/master)  [![Repo Status](https://img.shields.io/badge/status-proof%20of%20concept-lightgrey.svg?longCache=true)](https://docs.nlx.io/introduction/product-vision/)

NLX is an open source inter-organisational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

This repository contains all of the components to the current **Proof of Concept** of the [NLX Product Vision](https://docs.nlx.io/introduction/product-vision/). Do **not** use this code in production.

## Developing on NLX

Please find the latest documentation for using NLX on [docs.nlx.io](https://docs.nlx.io). This is a good place to start if you would like to develop an application or service that uses or provides API access over NLX.

## Questions and contributions

Read more on how to ask questions, file bugs and contribute code and documentation in [`CONTRIBUTING.md`](CONTRIBUTING.md).

## Building and running an NLX network locally

The NLX project consists of multiple components that together make up the entire NLX platform. Some components run as centralized NLX services, others run on-premise at organizations. All components are maintained in a single repository. This means that a developer has all the tools and code to build and test the complete NLX platform in a single repository. It simplifies version and dependency management and allows changes that affect multiple components to be combined in a single feature branch and merge-request.

If you want to develop locally, or run your own NLX network, you will likely want to start all the components.

### Requirements

Make sure you have installed the following tools:

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [helm](https://docs.helm.sh/using_helm/)
- [skaffold](https://github.com/GoogleContainerTools/skaffold#installation)

For autocompletion and local development tasks, it's also recommended to install the following:

- [go](https://golang.org/doc/install)

Also you will need to have [configured a `GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH) with `${GOPATH}/bin` added to your `PATH`.
After you set the `GOPATH`, be sure to reopen your terminal/shell to be sure the environment variables have been set correctly.

### Cloning

```bash
mkdir -p $GOPATH/src/go.nlx.io
cd $GOPATH/src/go.nlx.io
git clone https://gitlab.com/commonground/nlx
cd nlx
```

If you wish to contribute, fork the project and set the push origin to your fork.

```bash
git remote set-url --push origin git@gitlab.com:<YOUR-GITLAB-USERNAME>/nlx.git
```

### Running complete stack in kubernetes/minikube

Setup minikube on your local development machine. For developers, it's advised to setup minikube with 4 cores, 8GB RAM and 100+G storage.

e.g.: `minikube start --vm-driver=kvm2 --cpus 4 --memory 8192 --disk-size=100G`

Read the [minikube README](https://github.com/kubernetes/minikube) for more information.

Once minikube is running, initialize helm by running `helm init`.

Next, install the following dependencies:

- `traefik` for web and rest-api requests.
- `nginx-ingress` for grpc and mutual-tls connections. Latest version is currently(2018-09-06) broken, so needs `--version 0.17.1`
- `postgres` for directory-db and txlog-db.

```bash
helm install stable/traefik --name traefik --namespace traefik --values helm/traefik-values.yaml
helm install stable/nginx-ingress --version 0.17.1 --name nginx-ingress --namespace=nginx-ingress --values helm/nginx-ingress-values.yaml
helm install stable/postgresql --name postgresql --namespace=postgresql --values helm/postgresql-values.yaml
```

When these components are running, you can start all the NLX components by executing:

```bash
MINIKUBE_IP=$(minikube ip) skaffold dev
```

Finally, add the minikube hostnames to your machine's `/etc/hosts` file so you can reach the services from your browser.

```bash
echo "$(minikube ip)                 traefik.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)            docs.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)      certportal.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)       directory.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)   directory-api.dev.nlx.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)           txlog.dev.rdw.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)        irma-api.dev.rdw.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)         insight.dev.rdw.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)           txlog.dev.brp.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)        irma-api.dev.brp.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)         insight.dev.brp.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)       txlog.dev.denhaag.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)    irma-api.dev.denhaag.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)     insight.dev.denhaag.minikube" | sudo tee -a /etc/hosts
echo "$(minikube ip)      outway.dev.denhaag.minikube" | sudo tee -a /etc/hosts
```

You may now test the following sites:

- https://traefik.minikube:30443              A webinterface showing the status of the traefik ingress controller.
- http://docs.dev.nlx.minikube:30080          The NLX docs
- http://certportal.dev.nlx.minikube:30080    The NLX certportal
- http://directory.dev.nlx.minikube:30080     The NLX directory
- http://txlog.dev.rdw.minikube:30080/        Transactionlogs for the RDW example organization
- http://txlog.dev.brp.minikube:30080/        Transactionlogs for the BRP example organization
- http://txlog.dev.denhaag.minikube:30080/    Transactionlogs for the Den-Haag example organization
- http://outway.dev.denhaag.minikube:30080/   Outway in the Den-Haag example organization

To test a full request through outway>inway, use the PostmanEcho service through the exampleorg outway: `curl http://outway.dev.exampleorg.minikube:30080/DemoProviderOrganization/PostmanEcho/get?foo1=bar1&foo2=bar2`

Note the ports; `30080` and `30443` are routed via traefik (TLS handled by traefik), whereas `:443` and `:80` are used by nginx-ingress, which does "tcp-proxying" with ssl passthrough so the mutual TLS can be handled by inway/outway/directory/etc.

If you want to connect over IP instead of using a hostname, the ingress controller cannot route the request properly. Therefore you must setup a port-forward directly to the application you want to expose. This is useful, for example, when testing IRMA using a phone on the same WiFi network as your host machine.

```bash
kubectl --namespace nlx-rdw-dev port-forward deployment/irma-api-server 2222:8080
socat tcp-listen:3333,fork tcp:127.0.0.1:2222
```

You can now let your phone connect to the IRMA api server of RDW on `your.host.machine.ip:3333`

## Troubleshooting

If you are running into other issues, please [Post an Issue on GitLab](https://gitlab.com/commonground/nlx/issues).

## Deploying and releasing

**NOTE:** Automated releases are currently not available.

The [CI system of GitLab](https://gitlab.com/commonground/nlx/pipelines) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, GitLab builds and deploys it to the test and staging environments.

## License

Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)
