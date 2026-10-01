---
page_title: "xcsh_alert_policy landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy landing."
---

# xcsh_alert_policy landing

<a id="canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd338094b9e94a0b2ab9636e17c7eb111432bd772e0779243531add16452c5fe"></a>

## xcsh_alert_policy — xcsh_alert_policy / c6512f8f708a / 2

Breadcrumbs:

- xcsh_alert_policy

Manages new Alert Policy Object in F5 Distributed Cloud.

<a id="canonical-38a719da6872014d8074895c9036cdeda542ebd85ab5186c083eae5479137ca4"></a>

## Prerequisites — xcsh_alert_policy / c6512f8f708a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-ee9b7f6060e918b98c72e4633166f8ad84ae6269e47becfb46c738c58366e09d"></a>

## Minimal configuration — xcsh_alert_policy / c6512f8f708a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertPolicy Resource Example
# Manages new Alert Policy Object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertPolicy configuration
resource "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}
```

<a id="canonical-c80f551687393a18d8157af2ce12ce116bb6306528567d28f43a120468c1c73f"></a>

## Root configuration — xcsh_alert_policy / c6512f8f708a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1a8fff6c6a9d64b1f33676b789b3a63821250e5f6e5391c3591a677ede9452c1"></a>

## Next pages — xcsh_alert_policy / c6512f8f708a / 6

- [Property reference](../guides/resources--alert_policy--reference--group-001.md#canonical-37b4c8e808a6bc083c227d8d83de17312825613498b6d0a118b9fbe582ca685a)
- [Examples](../guides/resources--alert_policy--examples--group-001.md#canonical-23b7b4b5d648e60bcae0210a5ebcbb48e237cd6477c8164b3b0eadf6dfa7eafb)
- [Import](../guides/resources--alert_policy--lifecycle--group-001.md#canonical-093da74da01765ba279cfb8a6a3715d6e91781b9586437c83d6df39e058c6c71)
- [Timeouts](../guides/resources--alert_policy--lifecycle--group-001.md#canonical-adaf574656cbbba01a95093539cdb32e7e64bf699aec5f576cea2f2cf1e5e05a)
