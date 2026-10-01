---
page_title: "xcsh_aws_tgw_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site landing."
---

# xcsh_aws_tgw_site landing

<a id="canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-641512fc64389f68e04ee4acae4584531b8beb8f38394f062c5fc2b121f9a965"></a>

## xcsh_aws_tgw_site — xcsh_aws_tgw_site / f9cf8aee7c09 / 2

Breadcrumbs:

- xcsh_aws_tgw_site

Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS
Transit Gateway.

<a id="canonical-b8680b8b3a425ea3ac23fdd274deb03248f2834e6a1d1e0fa218fdb71b1e539c"></a>

## Prerequisites — xcsh_aws_tgw_site / f9cf8aee7c09 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f454855bed3e0e0b86a452d6b36caa9692d5c5791accaf83f338220e1f00a99f"></a>

## Minimal configuration — xcsh_aws_tgw_site / f9cf8aee7c09 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AWSTGWSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSTGWSite by name
data "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}

output "aws_tgw_site_id" {
  value = data.xcsh_aws_tgw_site.example.id
}
```

<a id="canonical-8a992d0d53ad62f50c456557b7a1908904a641b333c175ade2d57d5df8e83494"></a>

## Root configuration — xcsh_aws_tgw_site / f9cf8aee7c09 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-74ce9432a7cb228f4495c55b8f7013b5c30470a8c8d754fa9c1f32747f7c67ff"></a>

## Next pages — xcsh_aws_tgw_site / f9cf8aee7c09 / 6

- [Property reference](../guides/data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [Examples](../guides/data-sources--aws_tgw_site--examples--group-001.md#canonical-df6730039bad354a53ad7c0e0d66872ca27b44704e247e143cbf5dce3628592f)
