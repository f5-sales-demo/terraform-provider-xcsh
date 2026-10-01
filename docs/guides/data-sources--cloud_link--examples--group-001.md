---
page_title: "xcsh_cloud_link examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link examples."
---

# xcsh_cloud_link examples

<a id="canonical-1eb9e5d5f9aa54cd66d121744031aeea05feb427e7cf46b447e7f37931955fe1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3dffbf132b3c39f3e42d2c9d09e0cfbf2747821965f28e6f6502be4cc409ff3e"></a>

## Examples — Examples / d4a481978538 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- Examples

<a id="canonical-661865c87de65841175c364a52231b03eba88d64139afc84fb471e5adcb8a587"></a>

## Complete configurations — Examples / d4a481978538 / 3

- [Data source](data-sources--cloud_link--examples--group-001.md#canonical-c0e036ba7698eb849403f77e85b2d7712732acf53bda0e9e342b6e1274fae81e): valid configuration.

<a id="canonical-621d60f025836f954f38edb1b29977947af64b89b713e8e2a3c3f151f1fb4b35"></a>

## Next pages — Examples / d4a481978538 / 4

- [Data source](data-sources--cloud_link--examples--group-001.md#canonical-c0e036ba7698eb849403f77e85b2d7712732acf53bda0e9e342b6e1274fae81e)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)

<a id="canonical-c0e036ba7698eb849403f77e85b2d7712732acf53bda0e9e342b6e1274fae81e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04e318a0e2f823e1693aa722485ca63232cb23dc0d99c91447f9edd3de9281cc"></a>

## Data source — Data source / fa1686514358 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
- [Examples](data-sources--cloud_link--examples--group-001.md#canonical-1eb9e5d5f9aa54cd66d121744031aeea05feb427e7cf46b447e7f37931955fe1)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_link/data-source.tf`; digest `sha256:a7eb053580e156c686886504937855a78a126659f9e2d4b1412c03b42464b698`.

```terraform
# CloudLink Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudLink by name
data "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}

output "cloud_link_id" {
  value = data.xcsh_cloud_link.example.id
}
```

<a id="canonical-36407fc1bcda2941ea84e9a94ad80562205284bc09fb83ec05d1b74cde855af0"></a>

## Next pages — Data source / fa1686514358 / 3

- [Examples](data-sources--cloud_link--examples--group-001.md#canonical-1eb9e5d5f9aa54cd66d121744031aeea05feb427e7cf46b447e7f37931955fe1)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013)
