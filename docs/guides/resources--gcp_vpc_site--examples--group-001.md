---
page_title: "xcsh_gcp_vpc_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site examples."
---

# xcsh_gcp_vpc_site examples

<a id="canonical-9c3b5148b7bb2f407b76bff289c776f5e6ffb922e40edef4bd02be0a5ddabf62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62dc1d8f5c947e27fcf2bf8036eb57f570e502dd9eb0b9eacd86b0d040b9f082"></a>

## Examples — Examples / 41c604b5f2b3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- Examples

<a id="canonical-caa81a4497b980a5e035eaade35beff972e97299d254975e399b41a1049cb5c1"></a>

## Complete configurations — Examples / 41c604b5f2b3 / 3

- [Resource](resources--gcp_vpc_site--examples--group-001.md#canonical-07e372d6ca0d3821065d1505e2d18a208229992cdc49c510553d4ca9f2d10914): valid configuration.

<a id="canonical-2747d4bb43edf64ab2f3596d60765eedf44b67d01a91d8bee0f06631d2156e26"></a>

## Next pages — Examples / 41c604b5f2b3 / 4

- [Resource](resources--gcp_vpc_site--examples--group-001.md#canonical-07e372d6ca0d3821065d1505e2d18a208229992cdc49c510553d4ca9f2d10914)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-07e372d6ca0d3821065d1505e2d18a208229992cdc49c510553d4ca9f2d10914"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-572e01167f75f29a7a815f69342f9f71a39214146682224782171a1e8a89d385"></a>

## Resource — Resource / 52fd74af9513 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Examples](resources--gcp_vpc_site--examples--group-001.md#canonical-9c3b5148b7bb2f407b76bff289c776f5e6ffb922e40edef4bd02be0a5ddabf62)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_gcp_vpc_site/resource.tf`; digest `sha256:f7e954e266a46dcc4a155a5544354f116719772d491a99e466ea77159bc0413b`.

```terraform
# GCPVPCSite Resource Example
# Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GCPVPCSite configuration
resource "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"

  gcp_region    = "example-value"
  instance_type = "example-value"
  ssh_key       = "example-value"
}
```

<a id="canonical-be6804586f909cd54bd6f1ce34dcd402e345f0a08b0095a5d5866a7b87bfe514"></a>

## Next pages — Resource / 52fd74af9513 / 3

- [Examples](resources--gcp_vpc_site--examples--group-001.md#canonical-9c3b5148b7bb2f407b76bff289c776f5e6ffb922e40edef4bd02be0a5ddabf62)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
