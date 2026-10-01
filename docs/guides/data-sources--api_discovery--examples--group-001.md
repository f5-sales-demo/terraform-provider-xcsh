---
page_title: "xcsh_api_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery examples."
---

# xcsh_api_discovery examples

<a id="canonical-f19e1d4dd41f135517981f417fb0513dc02fa573b8a6f2bd5378f259eec47ed7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-adac1f2eeaef4acea1e36ca7d55fc2a0613a672736c0d4f40a3998592cb8e274"></a>

## Examples — Examples / d634ab542e09 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- Examples

<a id="canonical-f8fbae90179789d2506fffef389ff1b7e7cfddb6b0f161b9245e3f06665759ff"></a>

## Complete configurations — Examples / d634ab542e09 / 3

- [Data source](data-sources--api_discovery--examples--group-001.md#canonical-481235f2a2ba9e5188f81b8ae41f470d58e35af0c67c58a01e4bec4fb89261b4): valid configuration.

<a id="canonical-e63af8451727e53fe52cb5592b07ca85ea5ed788bd7d64a1c2e7e8c1f2a8308d"></a>

## Next pages — Examples / d634ab542e09 / 4

- [Data source](data-sources--api_discovery--examples--group-001.md#canonical-481235f2a2ba9e5188f81b8ae41f470d58e35af0c67c58a01e4bec4fb89261b4)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)

<a id="canonical-481235f2a2ba9e5188f81b8ae41f470d58e35af0c67c58a01e4bec4fb89261b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcf586ed2c08ba829c89d3d49ab95b74b4fb931e284766c3328940518f5cee41"></a>

## Data source — Data source / 0f3aee40eff8 / 2

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
- [Examples](data-sources--api_discovery--examples--group-001.md#canonical-f19e1d4dd41f135517981f417fb0513dc02fa573b8a6f2bd5378f259eec47ed7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_discovery/data-source.tf`; digest `sha256:f37cea7bc8746642eef8cfd7a0d2f33975184479df2f41465a9bdda95de57bb2`.

```terraform
# APIDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDiscovery by name
data "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}

output "api_discovery_id" {
  value = data.xcsh_api_discovery.example.id
}
```

<a id="canonical-dbf4d3824ab8b7972dd6a8fedf4c4795195bbc6443811cd93c84347e0b11ee09"></a>

## Next pages — Data source / 0f3aee40eff8 / 3

- [Examples](data-sources--api_discovery--examples--group-001.md#canonical-f19e1d4dd41f135517981f417fb0513dc02fa573b8a6f2bd5378f259eec47ed7)
- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8)
