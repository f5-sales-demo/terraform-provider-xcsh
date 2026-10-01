---
page_title: "xcsh_waf_bot_signatures landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_bot_signatures landing."
---

# xcsh_waf_bot_signatures landing

<a id="canonical-6e1fb8fd8cfc0c30cdc9be34f2007f846063ac8689adc77b969c99d5d3cdb624"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aa2bd65f97a11f441bcbf9dc5562ce72c13b745ce4ccb393ae42ba5c55cb757"></a>

## xcsh_waf_bot_signatures — xcsh_waf_bot_signatures / ffeaa193b80a / 2

Breadcrumbs:

- xcsh_waf_bot_signatures

Bot detection and defense configuration.

<a id="canonical-604ec14226b451d312406968b15d9ec1c1e25f3fe00faefa1496b66b3effc0a9"></a>

## Prerequisites — xcsh_waf_bot_signatures / ffeaa193b80a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-fd39878392ddc749346e958676ccdc4754c82c675c977b3267cbb0c101660bbb"></a>

## Minimal configuration — xcsh_waf_bot_signatures / ffeaa193b80a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFBotSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_bot_signatures" "example" {
}

output "waf_bot_signatures_result" {
  value = data.xcsh_waf_bot_signatures.example
}
```

<a id="canonical-d886fc961924754ea497ec3e8d5edeb1a9b5924d3e3f3b6fce1981da778573b2"></a>

## Root configuration — xcsh_waf_bot_signatures / ffeaa193b80a / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-8be32c1b83824e61b1ce43d95758ab99649926c1a9d054624114e86c54327a88"></a>

## Next pages — xcsh_waf_bot_signatures / ffeaa193b80a / 6

- [Property reference](../guides/data-sources--waf_bot_signatures--reference--group-001.md#canonical-22cb542c67b8286056ff79f058923875f4a5f2939ec4b6862cfc1f2817ad6a73)
- [Examples](../guides/data-sources--waf_bot_signatures--examples--group-001.md#canonical-29e9a910b02b7e528f124d51a1a50b78c81a18346aab6821cf1c7a6ae53af637)
