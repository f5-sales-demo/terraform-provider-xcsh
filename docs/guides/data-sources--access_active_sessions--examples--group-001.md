---
page_title: "xcsh_access_active_sessions examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions examples."
---

# xcsh_access_active_sessions examples

<a id="canonical-9c6a94d6ca9e39822b716f3896258ace46b22249c601fe3eb13a110372c48134"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6c4d9c3ede9268671777ac85c9f44532cb8e3390a35f339a5d8c3b4ac465c92"></a>

## Examples — Examples / bd974f88f37f / 2

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)
- Examples

<a id="canonical-33e98b104e2df2a7dfdb0fed6532c0d038385ef95e538e813cc803ecc6267d6d"></a>

## Complete configurations — Examples / bd974f88f37f / 3

- [Data source](data-sources--access_active_sessions--examples--group-001.md#canonical-626ae2105a59393b5f1fd96ad81ab10f16a42206acfb346d59462b6ea7c5bd41): valid configuration.

<a id="canonical-ddc357e5f9e824764274920781f9233a7a1008d7b71f6d4f08cf9e9b34fae21b"></a>

## Next pages — Examples / bd974f88f37f / 4

- [Data source](data-sources--access_active_sessions--examples--group-001.md#canonical-626ae2105a59393b5f1fd96ad81ab10f16a42206acfb346d59462b6ea7c5bd41)
- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)

<a id="canonical-626ae2105a59393b5f1fd96ad81ab10f16a42206acfb346d59462b6ea7c5bd41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-998b27f95d5a048ac5e12991bb1e7a9856d97a6fb8393ebe91600f75f0217862"></a>

## Data source — Data source / a9930b0377f6 / 2

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)
- [Examples](data-sources--access_active_sessions--examples--group-001.md#canonical-9c6a94d6ca9e39822b716f3896258ace46b22249c601fe3eb13a110372c48134)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_access_active_sessions/data-source.tf`; digest `sha256:9f2701c9ce37e0c0584f554dedc7733344fe85678c9dadc11b3997db86cac357`.

```terraform
# AccessActiveSessions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_sessions" "example" {
  namespace = "example-value"
}

output "access_active_sessions_result" {
  value = data.xcsh_access_active_sessions.example
}
```

<a id="canonical-b05ee38bfa4f26abcc0673967dce317e027238f851a3e50340c056ca67b609c5"></a>

## Next pages — Data source / a9930b0377f6 / 3

- [Examples](data-sources--access_active_sessions--examples--group-001.md#canonical-9c6a94d6ca9e39822b716f3896258ace46b22249c601fe3eb13a110372c48134)
- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f)
