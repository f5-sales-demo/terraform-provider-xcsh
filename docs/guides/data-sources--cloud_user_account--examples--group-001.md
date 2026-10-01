---
page_title: "xcsh_cloud_user_account examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account examples."
---

# xcsh_cloud_user_account examples

<a id="canonical-5b28d2a6e4eaaa25e915ccb803d71dba9a9e9a9309e122177c605a960bd2d6f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caed836922959e8cbe51ed11172c2f303b7db950e345b023acf76dab6df7761e"></a>

## Examples — Examples / d094a83e5aed / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- Examples

<a id="canonical-6b8e3f055e2c9f56de5bff5b2d35f69f51bb73477c6415800332304a7ea2248d"></a>

## Complete configurations — Examples / d094a83e5aed / 3

- [Data source](data-sources--cloud_user_account--examples--group-001.md#canonical-381b4d0baefb41c30470efa033c91ab382b6e20575db0b9fa1f055fef93144ea): valid configuration.

<a id="canonical-76d360066abe66977edbea0033311f900c192f6268add9e0ff286962ec76ed25"></a>

## Next pages — Examples / d094a83e5aed / 4

- [Data source](data-sources--cloud_user_account--examples--group-001.md#canonical-381b4d0baefb41c30470efa033c91ab382b6e20575db0b9fa1f055fef93144ea)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)

<a id="canonical-381b4d0baefb41c30470efa033c91ab382b6e20575db0b9fa1f055fef93144ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5226a7ec671acc301f6d154da715f59465afd5f3a7dd759b53ca234f9e3a5ea2"></a>

## Data source — Data source / aa1adb1d2849 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
- [Examples](data-sources--cloud_user_account--examples--group-001.md#canonical-5b28d2a6e4eaaa25e915ccb803d71dba9a9e9a9309e122177c605a960bd2d6f7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_user_account/data-source.tf`; digest `sha256:4bd66f34c8654fcf2e99224154cac4393ee6bb933699a10ecdb50c1d8d402e99`.

```terraform
# CloudUserAccount Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudUserAccount by name
data "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}

output "cloud_user_account_id" {
  value = data.xcsh_cloud_user_account.example.id
}
```

<a id="canonical-1511853c4d2f7f9e3e450103842094f86872a0fe765b6aead6d11865b6985010"></a>

## Next pages — Data source / aa1adb1d2849 / 3

- [Examples](data-sources--cloud_user_account--examples--group-001.md#canonical-5b28d2a6e4eaaa25e915ccb803d71dba9a9e9a9309e122177c605a960bd2d6f7)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489)
