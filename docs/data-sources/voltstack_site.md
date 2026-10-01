---
page_title: "xcsh_voltstack_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site landing."
---

# xcsh_voltstack_site landing

<a id="canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8899a8f6cddd3781acc8d2c0b2d701a1aaa121b1ca6e6f99828a43930823bef"></a>

## xcsh_voltstack_site — xcsh_voltstack_site / b511266dffab / 2

Breadcrumbs:

- xcsh_voltstack_site

Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing
sites.

<a id="canonical-d27563f4d081ee1d5e629a47360cd60c776a50938d6ce22569d79ef6520a800a"></a>

## Prerequisites — xcsh_voltstack_site / b511266dffab / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4d9659e9b03ce7ab3b734f1de01dc4d28e5247174e6067c377742ae1021b6906"></a>

## Minimal configuration — xcsh_voltstack_site / b511266dffab / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VoltstackSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VoltstackSite by name
data "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"
}

output "voltstack_site_id" {
  value = data.xcsh_voltstack_site.example.id
}
```

<a id="canonical-5e136c983d2f6c037e703cc7975a1c71e0922d2fcdad2de18c6e49c590b9f99e"></a>

## Root configuration — xcsh_voltstack_site / b511266dffab / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3959adaeec7f57d7a218ff8250f62b50d5e34b905bb48bbd45a0ff8b25802031"></a>

## Next pages — xcsh_voltstack_site / b511266dffab / 6

- [Property reference](../guides/data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [Examples](../guides/data-sources--voltstack_site--examples--group-001.md#canonical-731327ff03d150c7c8d69f459d6f85905909d41ecbd50e308ed1feb174cffb5e)
