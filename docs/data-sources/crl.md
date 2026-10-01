---
page_title: "xcsh_crl landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl landing."
---

# xcsh_crl landing

<a id="canonical-73f9576382ba605e0c52a90b24f3c47e46f2fa2fe6a252eff356e5dcb2c0e943"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d89d7c2d774657e79ac4af62b549d98af52ebb26db46cb7150e41b2047a9484"></a>

## xcsh_crl — xcsh_crl / bfc660bf2205 / 2

Breadcrumbs:

- xcsh_crl

Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.

<a id="canonical-5ecd1cac4994654c1889fc6c1b2ccc156e083aa23130fa07a11bdb28cc93dbf1"></a>

## Prerequisites — xcsh_crl / bfc660bf2205 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d32d560ed2b1c794b42381040365e6c04b7d7f73213ce3ba510603e429faf7b9"></a>

## Minimal configuration — xcsh_crl / bfc660bf2205 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```

<a id="canonical-9541233ccefa5a2db26b09ac1584e7d5da5e2d7420c4f492cf4e1ef7971a1442"></a>

## Root configuration — xcsh_crl / bfc660bf2205 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b399bab3aa1924c1d1e60bc603528b96b2f5d05a7ea9c65ab21afced33e0d018"></a>

## Next pages — xcsh_crl / bfc660bf2205 / 6

- [Property reference](../guides/data-sources--crl--reference--group-001.md#canonical-c2c016728d6cc79099c35ae246c02f69a282f25fbac91acb6b380a38ba11c747)
- [Examples](../guides/data-sources--crl--examples--group-001.md#canonical-b3f67ac2969dbe28fd8ecd6d912bf4b2163c31694d52b729046d2acbf3da56ec)
