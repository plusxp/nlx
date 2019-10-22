---
id: production
title: Production environment
---


## Introduction

To make use of the NLX production environment, you will need verified and signed certificates. This page elaborates on how the production environment is setup and how to obtain certificates.

## Preprod vs prod

Both preprod and prod run the same version of NLX components. Both environments have a closed CA PKI. Which means you'll need to obtain certificates through a manual verification process.

The preprod environment is meant for testing your production setup, before it actually moves to production. The preprod environment has a closed CA, but should not be used with real-life data.

The prod environment is where the actual production processes communicate. Services present real data and all logs are kept for accountability. The prod environment should never be used for testing.

Note that the preprod and prod environments are not stable yet. There will be breaking changes and required upgrades as we improve on features like TLS, protocols, discovery, proxying, etc.

## Obtaining a certificate

If you require a certificate for preprod and/or prod, please send an email to support@nlx.io. We'll help you getting started with the manual verification process from there.