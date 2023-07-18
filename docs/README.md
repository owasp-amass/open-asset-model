# open-asset-model

[![GoDoc](https://pkg.go.dev/badge/github.com/owasp-amass/open-asset-model/?utm_source=godoc)](https://pkg.go.dev/github.com/owasp-amass/open-asset-model)
[![License](https://img.shields.io/badge/license-apache%202-blue)](https://www.apache.org/licenses/LICENSE-2.0)
[![Contribute Yes](https://img.shields.io/badge/contribute-yes-brightgreen.svg)](./CONTRIBUTING.md)

[![Chat on Discord](https://img.shields.io/discord/433729817918308352.svg?logo=discord)](https://discord.gg/HNePVyX3cp)
[![Follow on Twitter](https://img.shields.io/twitter/follow/owaspamass.svg?logo=twitter)](https://twitter.com/owaspamass)

Asset definitions for an organization's external attack surface

# Context

It all started at a panel back in 2022. Jeff Foley, the amass project lead, was asked the question “What is an attack surface - is it simply the public infrastructure assets”. Having written an asset collection tool designed specifically for this purpose, you might expect him to say “that’s exactly right”. But he didn’t.

Amass has more or less commoditized the IT asset collection process and made it freely available across numerous platforms. As it has grown in popularity, the team spent a lot of time thinking about how it could better support the needs of its users. For this reason, we imagined the Amass Ecosystem, which includes this project - the Open Asset Model. This is a community-driven effort to uniformly describe assets that belong to both organizations and individuals.

Asset specifications have traditionally focused upon the technical, infrastructure-specific things that can be discovered on the internet. While this represents a potentially significant portion of an organization's assets, it is also limiting. The Open Asset Model seeks to expand on this and cover the breadth and depth of both physical and digital assets so that an organization can realize their full Attack Surface.

Open Asset Model defines not just the assets themselves, but also the relationships within and across types of assets. This allows the model to express the real-world interconnectedness that exist between assets. For details, refer to the [Taxonomy](docs/taxonomy.md) documentation.

## Goals

- To provide a transport specification that enables organizations
  to exchange their asset inventory both internally and externally
- To have the aforementioned specification represent the breadth and
  depth of assets that belong to both organizations and individuals.
- To drive awareness that the attack surface is much larger than
  your public infrastructure.
- To enable a community driven approach to maintaining and improving
  a model that every security organization’s asset inventory should encompass.

## Future Plans

- Evaluate and improve the initial model - Domains, IP Addresses, Autonomous Systems, Netblocks, and Regional Internet Registry Orgs.
- Support additional IT assets that were not in the initial model
  - Certificates
  - Tech Stack
- Extend the model to represent a more expansive view of what the community is calling External Attack Surface
  - Organizations & Enterprises
  - Mobile Applications
  - Social Applications
  - Cloud Vendors and 3rd Parties
  - VIPs, Executives and Key Personnel
  - Physical Locations
- Make Open Asset model available for other programming languages.

## Contributing

Open Asset Model is only as good as the community that's backing it.
If the aforementioned goals resonate with you, we'd love to have your help.
This could be a bug or suggestion you drop in an issue, dropping a message in Discord,
or tweaking a line (or character) or two in the project as a pull request.

Our [CONTRIBUTING.md](CONTRIBUTING.md) document contains details on how to get started.

## Assets

Supported asset types:

| Asset | Type definition |
|:------|:-----------|
| Fully Qualified Domain Name | `FQDN` |
| Autonomous System | `AutonomousSystem` |
| Regional Internet Registry Organization | `RIROrganization` |
| IP Address | `IPAddress` |
| Netblock | `Netblock` |

## Documentation

The documentation can be found in the Go packages repository: [open-asset-model](https://pkg.go.dev/github.com/owasp-amass/open-asset-model), there you can find the reference and descriptions to the types.

## Contributing

We are always happy to get new contributors on board! Please check [CONTRIBUTING.md](CONTRIBUTING.md) to learn how to
contribute to our codebase, and join our [Discord Server](https://discord.gg/HNePVyX3cp) to discuss current project goals.

## Licensing

This program is free software: you can redistribute it and/or modify it under the terms of the [Apache license](LICENSE).
