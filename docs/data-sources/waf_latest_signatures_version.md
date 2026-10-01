---
page_title: "xcsh_waf_latest_signatures_version landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_latest_signatures_version landing."
---

# xcsh_waf_latest_signatures_version landing

<a id="canonical-e82588039d40a0bd7b98f558c6e993e8c040b7d365e0b5c93ac34a46905a3b8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b0445f50084d56275c3b967bcb6ae92871304fdad2eebd6355194670fe8e7f1"></a>

## xcsh_waf_latest_signatures_version — xcsh_waf_latest_signatures_version / 7da701f8e43f / 2

Breadcrumbs:

- xcsh_waf_latest_signatures_version

Resource retrieval operation.

<a id="canonical-a172cab524ee2ca9b37706bd54d46e9a6be484875e1c1dadddbaca2025047976"></a>

## Prerequisites — xcsh_waf_latest_signatures_version / 7da701f8e43f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-76bb597cbe0b37f70a71031add99d63d85b48dcf22ce199bbdd45a9394ed7844"></a>

## Minimal configuration — xcsh_waf_latest_signatures_version / 7da701f8e43f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFLatestSignaturesVersion DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_latest_signatures_version" "example" {
}

output "waf_latest_signatures_version_result" {
  value = data.xcsh_waf_latest_signatures_version.example
}
```

<a id="canonical-203fa207447ed34782540a3b8065b52dab654ccd211f489f2b71334afb18182a"></a>

## Root configuration — xcsh_waf_latest_signatures_version / 7da701f8e43f / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-a6c2f0c22ca0c37f12cf22de5ae0cfa436840face46dc7e335106875d5b7289c"></a>

## Next pages — xcsh_waf_latest_signatures_version / 7da701f8e43f / 6

- [Property reference](../guides/data-sources--waf_latest_signatures_version--reference--group-001.md#canonical-44a7577ee22cd73e61d1d8938343de86038f422d583f20ca2c2c2bb7c07b94d7)
- [Examples](../guides/data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-9dfaa27c323ff0ee73075c8652c72ac3d5ff126248800053866038ab4e0686d4)
