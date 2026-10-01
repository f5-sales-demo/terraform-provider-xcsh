---
page_title: "xcsh_endpoint landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint landing."
---

# xcsh_endpoint landing

<a id="canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ad06553f9086bc3cd252209aa9546a8d7facf048002ee4ea986047a3194af82"></a>

## xcsh_endpoint — xcsh_endpoint / 04a601900c90 / 2

Breadcrumbs:

- xcsh_endpoint

Manages endpoint will create the object in the storage backend for namespace metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-e8f5296bb6a5736e2abc2e4a438481118e244499a73c9ed32f38d0bdb65bab9e"></a>

## Prerequisites — xcsh_endpoint / 04a601900c90 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-6d9eab1fd74fdef76f87c1195fed14da7e4c75a51c30d685431a8664ede8a47c"></a>

## Minimal configuration — xcsh_endpoint / 04a601900c90 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Endpoint Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Endpoint by name
data "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}

output "endpoint_id" {
  value = data.xcsh_endpoint.example.id
}
```

<a id="canonical-5d34ae8fdd0c9581e03902775a3e7efc1b288c747e61d48ebc214905202ab1e3"></a>

## Root configuration — xcsh_endpoint / 04a601900c90 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-fccaa217639e878a49364d0022943d95af53d8260f987f0495e92f5a40dbbc5c"></a>

## Next pages — xcsh_endpoint / 04a601900c90 / 6

- [Property reference](../guides/data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [Examples](../guides/data-sources--endpoint--examples--group-001.md#canonical-b10c566b27db1e8b6a31ee834ea7dbdeee24513c193a9ead9e8c97bbdddcba7f)
