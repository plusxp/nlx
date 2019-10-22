---
id: environments
title: Environments
---

## Environments

NLX is build on a modern infrastructure based on OpenStack and Kubernetes, with a development street that automates most of the build, test and deployment tasks. The project and code repository are available on [Gitlab](https://gitlab.com/commonground/nlx/nlx). NLX developers have their own CI/CD flow with several private and public environments. The following Environments are publicly available for different audiences:

* `Acceptance (Acc)`; for testing new features and formal acceptance - available for team members and stakeholders only.
* `Demo`; for testing the general setup and functionality - available for all interested parties.
* `Preproduction(Preprod)`; for testing the interaction of NLX with developers own software - available for all developers after acquiring a certificate.
* `Production`; for production functionality only - available to all users and developers after testing the interaction on preproduction.

To make use of the NLX preproduction or production environment you will need **verified and signed certificates**. Read the instructions on the page ["further-reading/production"](../production) for obtaining a certificate.

Besides the public environments there are private development and test environments for the NLX team to develop new features:

* `Development`; private environments for development - available for team members only.
* `Test`; environment for testing new features - available for team members and stakeholders.

All new features are developed in separate [Gitlab branches](https://gitlab.com/commonground/nlx/nlx/branches), after the initial tests the branches are merged into the master branch and deployed to the feature test and acceptance environments using the [CI/CD pipelines of Gitlab](https://gitlab.com/commonground/nlx/nlx/pipelines). The pipeline for NLX requires manual code review and approval followed bij automated build, unit test, integration tests, packaging and deployment. The link to the feature test environment is made available in the details of the [merge requests on Gitlab](https://gitlab.com/commonground/nlx/nlx/merge_requests).

After approval by Product Owner (eventually after consulting with stakeholders), the version of the master is tagged and this identical version is deployed to preproduction, production and demo.
All components like the directory, certportal, docs, demo and insight are available on all environments. The components are accessible using the following url construction:
  `<componentname.environment.product.topleveldomain>`

Leaving the environment empty will redirect to production. For example:

|URL   | Environment | Component  |
|:-----|:----------:|:-------------|
|[insight.demo.nlx.io](https://insight.demo.nlx.io/) |`Demo` |Insight application for viewing log entries|
|[docs.preprod.nlx.io](https://docs.preprod.nlx.io)|`Preproduction`| Technical documentation of NLX|
|[certportal.acc.nlx.io](https://certportal.acc.nlx.io/) |`Acceptance`| NLX certificate portal for this environment|
|[directory.acc.nlx.io](https://directory.acc.nlx.io/) |`Acceptance`|Directory with all available api’s for this environment  |
|[demo.nlx.io](https://demo.nlx.io/) |`Production`| Demo application (applying for a parking permit in Haarlem)|
|[nlx.io](https://nlx.io/) |`Production`| NLX home page |

*Note: All non production CA's (Certificate Authority's) are unsafe, for development and demonstration purpose only. Preproduction & production require a restricted CA configuration and related certificates.*

An overview of the environments and the deployment flow is illustrated in the following model:

<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" width="631px" height="441px" viewBox="-0.5 -0.5 631 441" content="&lt;mxfile modified=&quot;2019-05-08T14:01:11.658Z&quot; host=&quot;www.draw.io&quot; agent=&quot;Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36&quot; version=&quot;10.6.7&quot; etag=&quot;yTOTiUo_z62A_tE7FTsV&quot; type=&quot;github&quot;&gt;&lt;diagram id=&quot;EKVPhLETGR-r4NFndtJj&quot;&gt;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&lt;/diagram&gt;&lt;/mxfile&gt;"><defs/><g><rect x="0" y="0" width="630" height="440" rx="35.2" ry="35.2" fill="none" stroke="#38aeda" pointer-events="none"/><g transform="translate(251.5,7.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="126" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(71, 78, 87); line-height: 1.2; vertical-align: top; width: 126px; white-space: nowrap; overflow-wrap: normal; font-weight: bold; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit">NLX Environments view<br /></font></font></font></font></div></div></foreignObject><text x="63" y="12" fill="#474E57" text-anchor="middle" font-size="12px" font-family="source sans pro" font-weight="bold">[Not supported by viewer]</text></switch></g><rect x="20" y="155" width="590" height="265" rx="21.2" ry="21.2" fill="none" stroke="#474e57" stroke-dasharray="3 3" pointer-events="none"/><g transform="translate(567.5,162.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="39" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(71, 78, 87); line-height: 1.2; vertical-align: top; width: 40px; white-space: nowrap; overflow-wrap: normal; text-align: right;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit">Public   </font></font></font></font></div></div></foreignObject><text x="20" y="12" fill="#474E57" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><rect x="20" y="30" width="590" height="110" rx="16.5" ry="16.5" fill="none" stroke="#474e57" stroke-dasharray="3 3" pointer-events="none"/><g transform="translate(563.5,37.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="43" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(71, 78, 87); line-height: 1.2; vertical-align: top; width: 44px; white-space: nowrap; overflow-wrap: normal; text-align: right;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit">Private . </font></font></font></font></div></div></foreignObject><text x="22" y="12" fill="#474E57" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><rect x="30" y="40" width="440" height="90" rx="13.5" ry="13.5" fill-opacity="0.25" fill="#38aeda" stroke="none" pointer-events="none"/><g transform="translate(346.5,46.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="121" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(71, 78, 87); line-height: 1.2; vertical-align: top; width: 122px; white-space: nowrap; overflow-wrap: normal; font-weight: bold; text-align: right;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><span style="font-weight: normal">Development branches  </span><br /></font></font></font></font></div></div></foreignObject><text x="61" y="12" fill="#474E57" text-anchor="middle" font-size="12px" font-family="source sans pro" font-weight="bold">[Not supported by viewer]</text></switch></g><rect x="50" y="65" width="100" height="40" rx="6" ry="6" fill="#ffffff" stroke="#38aeda" pointer-events="none"/><g transform="translate(65.5,78.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="68" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(56, 174, 218); line-height: 1.2; vertical-align: top; width: 68px; white-space: nowrap; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><div style="text-align: center"><span>Development</span></div></div></div></foreignObject><text x="34" y="12" fill="#38AEDA" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><rect x="200" y="65" width="100" height="40" rx="6" ry="6" fill="#ffffff" stroke="#38aeda" pointer-events="none"/><g transform="translate(239.5,78.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="20" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(56, 174, 218); line-height: 1.2; vertical-align: top; width: 22px; white-space: nowrap; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><div style="text-align: center">Test</div></div></div></foreignObject><text x="10" y="12" fill="#38AEDA" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><path d="M 150 85 Q 170 80 193.72 83.95" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 198.9 84.82 L 191.42 87.12 L 193.72 83.95 L 192.57 80.21 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><rect x="340" y="67.5" width="100" height="35" fill="#ffbc2c" stroke="#ffffff" pointer-events="none"/><path d="M 350 67.5 L 350 102.5 M 430 67.5 L 430 102.5" fill="none" stroke="#ffffff" stroke-miterlimit="10" pointer-events="none"/><g transform="translate(353.5,79.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="72" height="10" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 10px; font-family: Helvetica; color: rgb(255, 255, 255); line-height: 1.2; vertical-align: top; width: 72px; white-space: nowrap; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;">Test CA (Public)</div></div></foreignObject><text x="36" y="10" fill="#FFFFFF" text-anchor="middle" font-size="10px" font-family="Helvetica">Test CA (Public)</text></switch></g><path d="M 340 85 L 302.24 85" fill="none" stroke="#000000" stroke-miterlimit="10" stroke-dasharray="3 3" pointer-events="none"/><path d="M 308.12 81.5 L 301.12 85 L 308.12 88.5" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><rect x="100" y="165" width="420" height="85" rx="12.75" ry="12.75" fill-opacity="0.25" fill="#38aeda" stroke="none" pointer-events="none"/><g transform="translate(428.5,171.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="89" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(71, 78, 87); line-height: 1.2; vertical-align: top; width: 90px; white-space: nowrap; overflow-wrap: normal; font-weight: bold; text-align: right;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><span style="font-weight: normal">Testing branches  </span><br /></font></font></font></font></div></div></foreignObject><text x="45" y="12" fill="#474E57" text-anchor="middle" font-size="12px" font-family="source sans pro" font-weight="bold">[Not supported by viewer]</text></switch></g><rect x="247" y="190" width="100" height="40" rx="6" ry="6" fill="#ffffff" stroke="#38aeda" pointer-events="none"/><g transform="translate(268.5,203.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="56" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(56, 174, 218); line-height: 1.2; vertical-align: top; width: 58px; white-space: nowrap; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><div style="text-align: center">Acceptance</div></div></div></foreignObject><text x="28" y="12" fill="#38AEDA" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><path d="M 224 210 L 240.63 210" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 245.88 210 L 238.88 213.5 L 240.63 210 L 238.88 206.5 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><rect x="387" y="192.5" width="100" height="35" fill="#ffbc2c" stroke="#ffffff" pointer-events="none"/><path d="M 397 192.5 L 397 227.5 M 477 192.5 L 477 227.5" fill="none" stroke="#ffffff" stroke-miterlimit="10" pointer-events="none"/><g transform="translate(401.5,204.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="70" height="10" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 10px; font-family: Helvetica; color: rgb(255, 255, 255); line-height: 1.2; vertical-align: top; width: 72px; white-space: nowrap; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;">Acc CA (Public)</div></div></foreignObject><text x="35" y="10" fill="#FFFFFF" text-anchor="middle" font-size="10px" font-family="Helvetica">Acc CA (Public)</text></switch></g><path d="M 387 210 L 349.24 210" fill="none" stroke="#000000" stroke-miterlimit="10" stroke-dasharray="3 3" pointer-events="none"/><path d="M 355.12 206.5 L 348.12 210 L 355.12 213.5" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 169.78 240 L 157 202 L 190.5 180 L 224 202 L 211.22 240 Z" fill="#e0e4ea" stroke="#ffffff" stroke-miterlimit="10" pointer-events="none"/><g transform="translate(158.5,198.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="63" height="22" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 10px; font-family: Helvetica; color: rgb(231, 93, 102); line-height: 1.2; vertical-align: top; width: 63px; white-space: normal; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;">Acceptance PO</div></div></foreignObject><text x="32" y="16" fill="#E75D66" text-anchor="middle" font-size="10px" font-family="Helvetica">Acceptance PO</text></switch></g><path d="M 250 105 Q 290 150 295.9 183.73" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 296.81 188.9 L 292.15 182.61 L 295.9 183.73 L 299.05 181.4 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><rect x="130" y="260" width="470" height="150" rx="22.5" ry="22.5" fill-opacity="0.25" fill="#38aeda" stroke="none" pointer-events="none"/><g transform="translate(505.5,266.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="92" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(71, 78, 87); line-height: 1.2; vertical-align: top; width: 93px; white-space: nowrap; overflow-wrap: normal; font-weight: bold; text-align: right;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><font style="vertical-align: inherit"><span style="font-weight: normal">Release branches  </span><br /></font></font></font></font></div></div></foreignObject><text x="46" y="12" fill="#474E57" text-anchor="middle" font-size="12px" font-family="source sans pro" font-weight="bold">[Not supported by viewer]</text></switch></g><rect x="300" y="306" width="100" height="40" rx="6" ry="6" fill="#ffffff" stroke="#38aeda" pointer-events="none"/><g transform="translate(304.5,319.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="90" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(56, 174, 218); line-height: 1.2; vertical-align: top; width: 90px; white-space: nowrap; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><div style="text-align: center">Production (Prod)</div></div></div></foreignObject><text x="45" y="12" fill="#38AEDA" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><rect x="300" y="366" width="100" height="35" fill="#ffbc2c" stroke="#ffffff" pointer-events="none"/><path d="M 310 366 L 310 401 M 390 366 L 390 401" fill="none" stroke="#ffffff" stroke-miterlimit="10" pointer-events="none"/><g transform="translate(311.5,372.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="76" height="22" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 10px; font-family: Helvetica; color: rgb(255, 255, 255); line-height: 1.2; vertical-align: top; width: 76px; white-space: normal; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;">Production CA (Restricted)</div></div></foreignObject><text x="38" y="16" fill="#FFFFFF" text-anchor="middle" font-size="10px" font-family="Helvetica">Production CA (Restricted)</text></switch></g><path d="M 350 366 L 350 348.24" fill="none" stroke="#000000" stroke-miterlimit="10" stroke-dasharray="3 3" pointer-events="none"/><path d="M 353.5 354.12 L 350 347.12 L 346.5 354.12" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><rect x="180" y="306" width="100" height="40" rx="6" ry="6" fill="#ffffff" stroke="#38aeda" pointer-events="none"/><g transform="translate(181.5,312.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="96" height="26" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(56, 174, 218); line-height: 1.2; vertical-align: top; width: 96px; white-space: normal; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><div style="text-align: center">PreProduction (PreProd)</div></div></div></foreignObject><text x="48" y="19" fill="#38AEDA" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><rect x="180" y="366" width="100" height="35" fill="#ffbc2c" stroke="#ffffff" pointer-events="none"/><path d="M 190 366 L 190 401 M 270 366 L 270 401" fill="none" stroke="#ffffff" stroke-miterlimit="10" pointer-events="none"/><g transform="translate(191.5,372.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="76" height="22" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 10px; font-family: Helvetica; color: rgb(255, 255, 255); line-height: 1.2; vertical-align: top; width: 76px; white-space: normal; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;">PreProduction CA (Restricted)</div></div></foreignObject><text x="38" y="16" fill="#FFFFFF" text-anchor="middle" font-size="10px" font-family="Helvetica">PreProduction CA (Restricted)</text></switch></g><path d="M 230 366 L 230 348.24" fill="none" stroke="#000000" stroke-miterlimit="10" stroke-dasharray="3 3" pointer-events="none"/><path d="M 233.5 354.12 L 230 347.12 L 226.5 354.12" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><rect x="420" y="306" width="100" height="40" rx="6" ry="6" fill="#ffffff" stroke="#38aeda" pointer-events="none"/><g transform="translate(454.5,319.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="30" height="12" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 12px; font-family: &quot;source sans pro&quot;; color: rgb(56, 174, 218); line-height: 1.2; vertical-align: top; width: 30px; white-space: nowrap; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;"><div style="text-align: center">Demo</div></div></div></foreignObject><text x="15" y="12" fill="#38AEDA" text-anchor="middle" font-size="12px" font-family="source sans pro">[Not supported by viewer]</text></switch></g><rect x="420" y="366" width="100" height="35" fill="#ffbc2c" stroke="#ffffff" pointer-events="none"/><path d="M 430 366 L 430 401 M 510 366 L 510 401" fill="none" stroke="#ffffff" stroke-miterlimit="10" pointer-events="none"/><g transform="translate(431.5,372.5)"><switch><foreignObject style="overflow:visible;" pointer-events="all" width="76" height="22" requiredFeatures="http://www.w3.org/TR/SVG11/feature#Extensibility"><div xmlns="http://www.w3.org/1999/xhtml" style="display: inline-block; font-size: 10px; font-family: Helvetica; color: rgb(255, 255, 255); line-height: 1.2; vertical-align: top; width: 76px; white-space: normal; overflow-wrap: normal; text-align: center;"><div xmlns="http://www.w3.org/1999/xhtml" style="display:inline-block;text-align:inherit;text-decoration:inherit;">Demo CA (Public)</div></div></foreignObject><text x="38" y="16" fill="#FFFFFF" text-anchor="middle" font-size="10px" font-family="Helvetica">Demo CA (Public)</text></switch></g><path d="M 470 366 L 470 348.24" fill="none" stroke="#000000" stroke-miterlimit="10" stroke-dasharray="3 3" pointer-events="none"/><path d="M 473.5 354.12 L 470 347.12 L 466.5 354.12" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 297 230 Q 350 265 350 299.63" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 350 304.88 L 346.5 297.88 L 350 299.63 L 353.5 297.88 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 322 230 Q 440 255 466.77 300.51" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 469.43 305.04 L 462.87 300.78 L 466.77 300.51 L 468.9 297.23 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 272 230 Q 220 265 228.49 299.81" fill="none" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/><path d="M 229.74 304.91 L 224.68 298.94 L 228.49 299.81 L 231.48 297.28 Z" fill="#000000" stroke="#000000" stroke-miterlimit="10" pointer-events="none"/></g></svg>