---
id: introduction
title: Introduction
---

**NLX** is an open source peer-to-peer system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API ecosystem with many organizations.

NLX is built as a core component of the [Common Ground vision](https://github.com/VNG-Realisatie/common-ground), which aims to convert or replace the current monolithic information systems of Dutch municipalities (and wider government) with a state of the art, API-first software landscape, fulfilling societal demand for automation, transparency and privacy.

Important business benefits of NLX are:

* Lower integration costs because of standardised API integration layer
* Better data quality because data is used at the data source
* Better AVG compliance because end users get insight into data usage
* Better logging and auditing of data usage
* More agile software systems because NLX facilitates component based software systems

To reduce integration costs and simplify building component based software systems, everyone should be (technically) able to use API's in other organizations as easy as in their own, while core data objects should only be manipulated by the one administratively responsible and used by all others. An additional advantage is that public data can easily be made available to everyone. NLX makes it possible to easily and securely connect to any API in the ecosystem.

**NOTE**: We are currently running NLX software in production and looking for organizations to participate.
[participating organizations](https://directory.nlx.io/) contact: support@nlx.io

## High level business overview

Starting with NLX implies starting with API’s. More specific it's starting with
generic API’s instead of one-off integration solutions. NLX is useful for both
organizations that want to consume API's and organizations that want to
provider API's to other consuming organizations.

NLX is developed mainly by VNG Realisatie, but anyone can contribute on Gitlab.
To start using NLX, it is not necessary to contribute to NLX development, since
NLX can be installed 'as is'. A good way to start is by reading the
documentation and having your system engineers or suppliers follow the
[Try NLX guide](../try-nlx/setup-your-environment.md) to learn the concepts
of NLX and set up the first test implementation. Some knowledge of IT, the
internet, API's and security is required. The NLX team is available through
Gitlab for support to developers.

Setting up the first Outway for testing is done in a matter of a few hours,
setting up an Inway requires having a basic API and a publicly resolvable URL.
Configuring the inway is done in a few hours, setting up the API might take
more time depending on your infrastructure. Launching a service on a public
available network (the internet) requires more than just bringing the NLX and
API software to production. A project should address all functional and non
functional requirements and business context before launch.

Consuming API’s is easier than providing an API. A good place to start is to
consume an already available API. Such as for example the [BAG
API](https://zakelijk.kadaster.nl/-/bag-api) of the Dutch Kadaster or the [KVK
API](https://developers.kvk.nl/) of the Dutch Chamber of Commerce.

Providing API’s trough NLX is quite straightforward when the API’s are already
available. Suitable API’s to start with as a provider are two examples of API’s
that Amsterdam has made available. Consider providing an API for monuments or
garbage bin’s by reusing the already developed API specifications: [Monumenten
API](https://api.data.amsterdam.nl/monumenten) or [Afval
API](https://api.data.amsterdam.nl/afval/).

## High level technical overview

Within the NLX integration landscape there are two main concepts:
**applications** and **services**. Applications are only available within
organizations and usually provide users a **client** interface to query and
mutate data. Services provide **API's** to applications and are accessable
within and across organizations.

NLX provides a developer friendly way to use standardised resources between
organizations. It provides a gateway for querying services on the ecosystem as
well as an  gateway to offer services to the ecosystem. NLX provides support
for HTTP/1.x services like REST/JSON and SOAP/XML. HTTP/2 support (gRPC) has 
been added.

NLX provides two different types of gateways: the **inway** and **outway**.
Through an inway an organization can provide services to the NLX ecosystem and
through an outway an organization can query services on the NLX ecosystem. The
gateways are usually deployed centrally within the organization although it is
possible for one organization to deploy multiple instances of inway and outway
on different locations.

<svg class="high-level-overview-illustration" viewBox="0 0 665 139" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><defs><rect id="a" x="0" y="0" width="300" height="88" rx="5"/><mask id="c" maskContentUnits="userSpaceOnUse" maskUnits="objectBoundingBox" x="0" y="0" width="300" height="88" fill="#fff"><use xlink:href="#a"/></mask><rect id="b" x="365" y="0" width="300" height="88" rx="5"/><mask id="d" maskContentUnits="userSpaceOnUse" maskUnits="objectBoundingBox" x="0" y="0" width="300" height="88" fill="#fff"><use xlink:href="#b"/></mask></defs><g fill="none" fill-rule="evenodd"><path d="M123.068 46l8.672 4.878-.98 1.744-12-6.75-1.55-.872 1.55-.872 12-6.75.98 1.744L123.068 44H168.932l-8.672-4.878.98-1.744 12 6.75 1.55.872-1.55.872-12 6.75-.98-1.744L168.932 46H123.068z" fill="#9B9B9B" fill-rule="nonzero"/><use stroke="#777" mask="url(#c)" stroke-width="4" stroke-dasharray="5,3" xlink:href="#a"/><g transform="translate(25 23)"><rect stroke="#7096FF" stroke-width="2" fill="#FFF" x="1" y="1" width="88" height="42" rx="5"/><text font-family="Muli" font-size="16" font-weight="500" fill="#517FFF"><tspan x="25.165" y="26">Client</tspan></text></g><use stroke="#777" mask="url(#d)" stroke-width="4" stroke-dasharray="5,3" xlink:href="#b"/><g transform="translate(550 20)"><rect stroke="#7096FF" stroke-width="2" fill="#FFF" x="1" y="1" width="88" height="42" rx="5"/><text font-family="Muli" font-size="16" font-weight="500" fill="#517FFF"><tspan x="34.061" y="26">API</tspan></text></g><path d="M270 45c22.395-4 43.062-6 62-6 18.938 0 39.605 2 62 6" stroke="#FFF" stroke-width="5"/><path d="M277.192 45.469l9.371 3.346-.673 1.884-12.966-4.63-1.674-.599 1.38-1.12 10.691-8.676 1.26 1.553-7.726 6.27 7.992-1.366C302.377 39.377 318.762 38 334 38s31.622 1.377 49.166 4.133l7.979 1.364-7.726-6.27 1.26-1.553 10.69 8.676 1.38 1.12-1.673.598-12.966 4.63-.673-1.883 9.37-3.346-7.964-1.363C365.41 41.37 349.13 40 334 40c-15.13 0-31.411 1.369-48.83 4.104l-7.978 1.365z" fill="#9B9B9B" fill-rule="nonzero"/><g transform="translate(177 23)"><rect fill="#7096FF" width="90" height="44" rx="5"/><text font-family="Muli" font-size="16" font-weight="500" fill="#FFF"><tspan x="18.605" y="26">Outway</tspan></text></g><g transform="translate(399 22)"><rect fill="#7096FF" width="90" height="44" rx="5"/><text font-family="Muli" font-size="16" font-weight="500" fill="#FFF"><tspan x="24.597" y="26">Inway</tspan></text></g><text font-family="Muli" font-size="16" font-weight="500" fill="#2C2C2C"><tspan x="64.092" y="135">Requesting application</tspan></text><text font-family="Muli" font-size="16" font-weight="500" fill="#2C2C2C"><tspan x="442.096" y="135">Providing service</tspan></text><path d="M497.068 45l8.672 4.878-.98 1.744-12-6.75-1.55-.872 1.55-.872 12-6.75.98 1.744L497.068 43H542.932l-8.672-4.878.98-1.744 12 6.75 1.55.872-1.55.872-12 6.75-.98-1.744L542.932 45H497.068z" fill="#9B9B9B" fill-rule="nonzero"/></g></svg>

Here you can see a full request-response flow on NLX. An application performs a
request on the outway within the same organization. The outway routes the
request to the inway of the organization providing the service. The inway
routes the request to the service. The service responds to the request and this
response is routed through the NLX landscape, to the outway and arrives at the
application.