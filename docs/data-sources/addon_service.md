---
page_title: "xcsh_addon_service landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service landing."
---

# xcsh_addon_service landing

<a id="canonical-bc6915e40d9e500717e2d5968bc07fc54c844773be24d2f94cca4e8c5e83fe99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ac0bef858ce745b00fd2c70059264c55b6c5bdb8327b416390a520ebbc7360f"></a>

## xcsh_addon_service — xcsh_addon_service / 39d83e413826 / 2

Breadcrumbs:

- xcsh_addon_service

Retrieves information about an F5 Distributed Cloud Addon Service.

Addon services are system-managed resources that provide additional functionality such as Bot
Defense, Client Side Defense, and other security features. This data source allows you to query
addon service details including tier requirements and activation type.

~&gt; \*\*Note:\*\* Addon services cannot be created or modified via Terraform. To activate or
subscribe to an addon service, please use the F5 Distributed Cloud Console or contact your account
team.

<a id="canonical-70eed51f3e0ab04500e53561f4646d8abaccc45b4a313f0dc2e68d87efa6b0c4"></a>

## Prerequisites — xcsh_addon_service / 39d83e413826 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-421772c7bb92add6407be3f5374625f2d2ba551a7d9ef1637e4c848c4a9539e2"></a>

## Minimal configuration — xcsh_addon_service / 39d83e413826 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```

<a id="canonical-d3aca3543b71b589ad1004fa17972ab50b70f856c068ffb39b497cc7f6503c23"></a>

## Root configuration — xcsh_addon_service / 39d83e413826 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0b2d28cce7179d8018474cae0f67b169a62dddc91db1857cd70062c283f4c120"></a>

## Next pages — xcsh_addon_service / 39d83e413826 / 6

- [Property reference](../guides/data-sources--addon_service--reference--group-001.md#canonical-11d603208fbd7d6adbc72e5aae21645bfb00d74593d5c996a45e0a5dfc81a055)
- [Examples](../guides/data-sources--addon_service--examples--group-001.md#canonical-a0399fbaae69f1178cd71f7f0650d72facb92ceb4977c1fad637c238fe1acb8f)
