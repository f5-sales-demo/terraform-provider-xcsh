---
page_title: "xcsh_advertise_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy examples."
---

# xcsh_advertise_policy examples

<a id="canonical-c122cffd8cafcc703c56445d899d66062863b9857605e22c90abff36dee4d815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0083b43e589c5a63854a2f347dee5ad3bb4f8c970b436471df7109d2bc8974c"></a>

## Examples — Examples / e0f4d879a2da / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- Examples

<a id="canonical-790b6beae4a371f0b15e4b67565605c8c2589e08ad9b496173d56bbb3435de9b"></a>

## Complete configurations — Examples / e0f4d879a2da / 3

- [Data source](data-sources--advertise_policy--examples--group-001.md#canonical-6fb7779cfe3375ccb951c925e30668ab0f0714f9a135b4a93d9ef3fd77698ed9): valid configuration.

<a id="canonical-599ee367d7027a7e0d43510f4d5b28e600171a436a20b4f7d66eecf25bd935bc"></a>

## Next pages — Examples / e0f4d879a2da / 4

- [Data source](data-sources--advertise_policy--examples--group-001.md#canonical-6fb7779cfe3375ccb951c925e30668ab0f0714f9a135b4a93d9ef3fd77698ed9)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)

<a id="canonical-6fb7779cfe3375ccb951c925e30668ab0f0714f9a135b4a93d9ef3fd77698ed9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c3adb3f5ba9f7c551004fa452d233ba1a4ee3b6cea209949992b1f6dacf198a"></a>

## Data source — Data source / 109a88d76619 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
- [Examples](data-sources--advertise_policy--examples--group-001.md#canonical-c122cffd8cafcc703c56445d899d66062863b9857605e22c90abff36dee4d815)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_advertise_policy/data-source.tf`; digest `sha256:0faa48749e677b3e8d75bc3281fff3815a14a824fbf83c244cc4c6bf9f75dad8`.

```terraform
# AdvertisePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AdvertisePolicy by name
data "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}

output "advertise_policy_id" {
  value = data.xcsh_advertise_policy.example.id
}
```

<a id="canonical-4e7ab7a67819a836b8e6290e4265731fbca9f70a0a80ab6fbd44abe9b37c0888"></a>

## Next pages — Data source / 109a88d76619 / 3

- [Examples](data-sources--advertise_policy--examples--group-001.md#canonical-c122cffd8cafcc703c56445d899d66062863b9857605e22c90abff36dee4d815)
- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a)
