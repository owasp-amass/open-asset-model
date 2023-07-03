# Open Asset Model Specification

## Goals

The Open Asset Model is an effort to uniformly describe assets
that belong to both organizations and individuals.

Asset specifications have traditionally focused upon the technical,
infrastructure-specific things that can be discovered on the internet.
While this represents a potentially significant portion of an organization's
assets, it is also limiting. The Open Asset Model seeks to expand on
this and cover the breadth and depth of both physical and digital assets
so that an organization can realize their full Attack Surface.

Open Asset Model defines not just the assets themselves, but also the
relationships within and across types of assets. This allows the model
to express the real-world interconnectedness that exist between assets.

## Source & Data Provenance

Open Asset Model also provides models to support [provenance][w3c-provenance]
tracking of assets.

When it comes to asset discovery, it is sometimes a bit of a mystery as to
how an asset was discovered. For this reason, we believe it is necessary to
capture the source, or [provenance](#data-provenance), of assets.

The source of the data is useful in a number of different ways:

* It can help understand the trustworthiness and/or accuracy of the data provided.
  Data sourced from less trusted sources may need to go through additional
  validation steps before it can be used.
* It can inform whether the data can be transferred with other parties,
  which is especially important in meeting modern data privacy requirements.
* It can be used to piece together the chain of events that lead to the discovery of
  a disputed asset.

## Asset Taxonomy

### Domain Assets

#### FQDN

| Property | Type | Required | Description |
| -------- | ---- | -------- | ----------- |
| `name` | string | true | Fully Qualified Domain Name |

##### Outgoing Relationships

| Relationships | Type |
| -------- | ---- |
| `a_record` | `IPAddress` |
| `aaaa_record` | `IPAddress` |
| `cname_record` | `FQDN` |
| `ns_record` | `FQDN` |
| `ptr_record` | `FQDN` |
| `mx_record` | `FQDN` |
| `srv_record` | `FQDN` |
| `srv_record` | `IPAddress` |

##### Incoming Relationships

| Relationship | Type |
| ------------ | ---- |
| `ns_record` | `FQDN` |
| `cname_record` | `FQDN` |
| `ptr_record` | `FQDN` |

### Networking Assets

#### AutonomousSystem

| Property | Type | Required | Description |
| -------- | ---- | -------- | ----------- |
| `number` | int | true | Autonomous System Number |

##### Outgoing Relationships

| Relationship | Type |
| ------------ | ---- |
| `announces` | `Netblock` |
| `managed_by` | `RIROrganization` |

##### Incoming Relationships

| Relationship | Type |
| ------------ | ---- |
| No incoming relationships | |

#### IP Address

| Property | Type | Required | Description |
| -------- | ---- | -------- | ----------- |
| `address` | object | true | Network address |
| `type` | string | true | IPv4 or IPv6 |

##### Outgoing Relationships

| Relationship | Type |
| ------------ | ---- |
| No outgoing relationships | |

##### Incoming Relationships

| Relationship | Type |
| ------------ | ---- |
| `a_record` | `FQDN` |
| `aaaa_record` | `FQDN` |
| `contains` | `Netblock` |

#### Netblock

| Property | Type | Required | Description |
| -------- | ---- | -------- | ----------- |
| `cidr` | object | true | CIDR representation of an  IP block |
| `type` | string | true | IPv4 or IPv6 |

##### Outgoing Relationships

| Relationship | Type |
| ------------ | ---- |
| `contains` | `IPAddress` |

##### Incoming Relationships

| Relationship | Type |
| ------------ | ---- |
| `announces` | `AutonomousSystem` |

#### RIROrganization

| Property | Type | Required | Description |
| -------- | ---- | -------- | ----------- |
| `name` | string | true | Organization Name |
| `rir_id` | string | false | the unique identifier of the RIR organization |
| `rir` | string | false | Regional Internet Registry |

##### Outgoing Relationships

| Relationship | Type |
| ------------ | ---- |
| No outgoing relationships |

##### Incoming Relationships

| Relationship | Type |
| ------------ | ---- |
| `managed_by` | `AutonomousSystem` |

## Data Provenance

In order to understand how assets have been discovered,
the Open Asset Model provides a specification that allows
an operator to trace the origin of how an asset was discovered

As an example, the following log event could be used to determine how `108.162.193.247`
was associated with the `owasp.org` domain.

```text
Source Asset {"id":1,"name":"owasp.org","tld":"org"} generated a [DNS A Record Request] that was serviced by [source] on [timestamp] and resulted in Asset {id:2,"address":"108.162.193.247","type":"v4"}
```

or in the negative case

```text
Source Asset [] generated [request type] on [timestamp] that had no sources that could fulfill the request
```

* `source_asset`: the id of the asset that was used to generate the
  request that found the asset.
* `asset`: the id of the asset discovered.
* `timestamp`: the timestamp in which the asset was discovered.
* `source`: - The source that the asset was derived from.
  * `name` - The name of the source

In the future, source can be extended to capture additional information
about the nature of the collection that was performed.

[w3c-provenance]:https://www.w3.org/2005/Incubator/prov/wiki/What_Is_Provenance]
