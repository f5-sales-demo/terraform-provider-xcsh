---
page_title: "xcsh_protected_application landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application landing."
---

# xcsh_protected_application landing

<a id="canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ff8bed7989d496b6f2dcb9902f1a03079f99873af07cf24284158939f583c2c"></a>

## xcsh_protected_application — xcsh_protected_application / ddfbc919938e / 2

Breadcrumbs:

- xcsh_protected_application

Manages applications protected by Bot Defense in F5 Distributed Cloud.

<a id="canonical-1e4db40dc72a4d89d7dea23f56027a2b6da0046d3a19c30e11eea17179e041ee"></a>

## Prerequisites — xcsh_protected_application / ddfbc919938e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a1ec774af219804547445c3f5e3cae26c3e501ee0ea991eb91e2060c13b39846"></a>

## Minimal configuration — xcsh_protected_application / ddfbc919938e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedApplication by name
data "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}

output "protected_application_id" {
  value = data.xcsh_protected_application.example.id
}
```

<a id="canonical-2eeeb418b4edbf059668f309d6e839adbd28ed201131fcabf0fb454054e351b6"></a>

## Root configuration — xcsh_protected_application / ddfbc919938e / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-f14beaf16fcd9e99f55bd5fbe78fec178ed42efeb3691015785a2879216d3700"></a>

## Next pages — xcsh_protected_application / ddfbc919938e / 6

- [Property reference](../guides/data-sources--protected_application--reference--group-001.md#canonical-01fe0beb3a9a157152e3c31967c36fef103d6c17bd215e4c7ebc5bb5565d1d92)
- [Examples](../guides/data-sources--protected_application--examples--group-001.md#canonical-a13d8b7dc315447b82100ed2c0d9253b9dd6100f89b018dc7c88b70babf866d2)
