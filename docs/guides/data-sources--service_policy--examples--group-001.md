---
page_title: "xcsh_service_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy examples."
---

# xcsh_service_policy examples

<a id="canonical-b43809c2072d0e215aefdfe6e7ca74d503f696f88f9d9f4abacdb2e701bde2c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85c41a22dca16ad6503422ad53bf176d1bf13029913cd55504fa0f2116651a4c"></a>

## Examples — Examples / 054e076385b8 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- Examples

<a id="canonical-30b45fb8f700aaa212a052189834c95741d00bc89b46a42abd5f62a65e9d4305"></a>

## Complete configurations — Examples / 054e076385b8 / 3

- [Data source](data-sources--service_policy--examples--group-001.md#canonical-df25f45f7d3ec2e5e23ae83d8bc09ef11cb7d573b1963cd79d5b081d7d236bc7): valid configuration.

<a id="canonical-4acc36f91257d911de20dbd89b0b782505020ec2b3c87f4736dd409c8d374d6e"></a>

## Next pages — Examples / 054e076385b8 / 4

- [Data source](data-sources--service_policy--examples--group-001.md#canonical-df25f45f7d3ec2e5e23ae83d8bc09ef11cb7d573b1963cd79d5b081d7d236bc7)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)

<a id="canonical-df25f45f7d3ec2e5e23ae83d8bc09ef11cb7d573b1963cd79d5b081d7d236bc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72b3380d60a534b820871b9664c8e4ace6486c3f419fc01c4db74dee84b2428b"></a>

## Data source — Data source / 642cb0283fd3 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
- [Examples](data-sources--service_policy--examples--group-001.md#canonical-b43809c2072d0e215aefdfe6e7ca74d503f696f88f9d9f4abacdb2e701bde2c0)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy/data-source.tf`; digest `sha256:9189dd7384a926dd448e27affcd3a13f4fb5015e7cdfa0bb42501f00fe48f58d`.

```terraform
# ServicePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicy by name
data "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}

output "service_policy_id" {
  value = data.xcsh_service_policy.example.id
}
```

<a id="canonical-a9ed37915a82abe31a1dbe82d422980ef7851a088f8d80636f309684076a8ac1"></a>

## Next pages — Data source / 642cb0283fd3 / 3

- [Examples](data-sources--service_policy--examples--group-001.md#canonical-b43809c2072d0e215aefdfe6e7ca74d503f696f88f9d9f4abacdb2e701bde2c0)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3)
