---
page_title: "xcsh_service_policy_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_set landing."
---

# xcsh_service_policy_set landing

<a id="canonical-52eaa9a13b0184a72441569b070040914137b030f2dff5fcdbb25301ba8d99b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6f6095be66cff84e31a2718330511fed6a8a1c3cff920aaed7b7001563376e9"></a>

## xcsh_service_policy_set — xcsh_service_policy_set / 798ccf734010 / 2

Breadcrumbs:

- xcsh_service_policy_set

Manages a Service Policy Set resource in F5 Distributed Cloud for get service\_policy\_set reads a
given object from storage backend for metadata.namespace. configuration. (read-only data source)

<a id="canonical-db72365b066949e95c5998b8e95262361a43b6b4d0695fa5f859481b9f10600e"></a>

## Prerequisites — xcsh_service_policy_set / 798ccf734010 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e4e10a708192c49808c4fb6df2315ff7afee780224247607ef95cade792cacd9"></a>

## Minimal configuration — xcsh_service_policy_set / 798ccf734010 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```

<a id="canonical-53ae71654bb7dfddb41c27e536923e0c3a92fa5487d4b8e52dbcd7c3ce8d9273"></a>

## Root configuration — xcsh_service_policy_set / 798ccf734010 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-54bab6e7644442c751f3990c4c745ba0fbb126a519e84016f403df71d602d5f0"></a>

## Next pages — xcsh_service_policy_set / 798ccf734010 / 6

- [Property reference](../guides/data-sources--service_policy_set--reference--group-001.md#canonical-d53562a654aeea2fda5c0f68cc93c43674dbd4ca72481940eaa91fb9ed762863)
- [Examples](../guides/data-sources--service_policy_set--examples--group-001.md#canonical-64a0cc77566504934e629646dfc340565e8343844ecc4a17ecf571158b2ff726)
