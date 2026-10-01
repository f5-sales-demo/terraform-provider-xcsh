---
page_title: "xcsh_cloud_link landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link landing."
---

# xcsh_cloud_link landing

<a id="canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20bbe42eabefef7ece8fa0a3faeb3dfff8e1d23385010321b4f1b468cc0be68c"></a>

## xcsh_cloud_link — xcsh_cloud_link / ee0db775d2c5 / 2

Breadcrumbs:

- xcsh_cloud_link

Manages new CloudLink with configured parameters in F5 Distributed Cloud.

<a id="canonical-6f27d4432d0d710ba82738eddbbe92137bf938f8d2dd6aa6ae1d6b921ae40bfd"></a>

## Prerequisites — xcsh_cloud_link / ee0db775d2c5 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4c7bd50eeed9332a131e0e2757961b40d35f9ac81cf950c921061c1fabd2a110"></a>

## Minimal configuration — xcsh_cloud_link / ee0db775d2c5 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudLink Resource Example
# Manages new CloudLink with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudLink configuration
resource "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}
```

<a id="canonical-55891ead6e2408ee62c2172de7e3ecc59f2cc9ff03827daaf6302e31ef3727bf"></a>

## Root configuration — xcsh_cloud_link / ee0db775d2c5 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-bf2f396d5fef16d97d0b7d3230fcd820f1665b445e605537a4dbebe0f1d67e52"></a>

## Next pages — xcsh_cloud_link / ee0db775d2c5 / 6

- [Property reference](../guides/resources--cloud_link--reference--group-001.md#canonical-f9f572eb4dd4355ccf1d4d4a87719ef48acf695e515e7a8097f8aa7e0a96fa35)
- [Examples](../guides/resources--cloud_link--examples--group-001.md#canonical-eb8da44addd6c73094389b886d916b0452553228175d7b5178810d9b504760d0)
- [Import](../guides/resources--cloud_link--lifecycle--group-001.md#canonical-b02ea34a0e6b77f2b0ced2f011a4c3d6ab6e632360b9e31a06291237057db8c5)
- [Timeouts](../guides/resources--cloud_link--lifecycle--group-001.md#canonical-ff85d75f630f7b81df6185e1842da42825109e5acdd7594baaa694f83d4bfee8)
