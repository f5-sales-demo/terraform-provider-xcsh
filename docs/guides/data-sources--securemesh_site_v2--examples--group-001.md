---
page_title: "xcsh_securemesh_site_v2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 examples."
---

# xcsh_securemesh_site_v2 examples

<a id="canonical-dd1bd31a95e67d7324b79547dcf212c3ece23b79f1d8be24343b53effe8e03f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c769d286827ba9cab884fc3c96f0fc44592976adfd8f443864184e17caf1c97"></a>

## Examples — Examples / 92649a6673f6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- Examples

<a id="canonical-4571b44df2172fe5da6763562e4ca4e7b25f3dc508c5a7294db0828e85b56611"></a>

## Complete configurations — Examples / 92649a6673f6 / 3

- [Data source](data-sources--securemesh_site_v2--examples--group-001.md#canonical-b86558c1218d0b277fd419df3f98a6792d81e5b8f437233cf94392b9aa756770): valid configuration.

<a id="canonical-ba9f0f84ffe5d0edb5f23dd81f2a72d0504662c6fca341bb6eabb6ba7d0095a5"></a>

## Next pages — Examples / 92649a6673f6 / 4

- [Data source](data-sources--securemesh_site_v2--examples--group-001.md#canonical-b86558c1218d0b277fd419df3f98a6792d81e5b8f437233cf94392b9aa756770)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b86558c1218d0b277fd419df3f98a6792d81e5b8f437233cf94392b9aa756770"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30ad3581414d2c64933f9ad7f13a73f7e8b2c6d3702ab4249475cf4c80aaf2cd"></a>

## Data source — Data source / 6770182bd8b5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Examples](data-sources--securemesh_site_v2--examples--group-001.md#canonical-dd1bd31a95e67d7324b79547dcf212c3ece23b79f1d8be24343b53effe8e03f7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_securemesh_site_v2/data-source.tf`; digest `sha256:2abaac1742f086ab988ef94e8d836fe6f538c549072b6b6dd846213e69e59239`.

```terraform
# SecuremeshSiteV2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSiteV2 by name
data "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.xcsh_securemesh_site_v2.example.id
}
```

<a id="canonical-404cde1a90aacdcbe216b12c03a235c1cee9fcf96e5c1e258f35bacdc3ced08a"></a>

## Next pages — Data source / 6770182bd8b5 / 3

- [Examples](data-sources--securemesh_site_v2--examples--group-001.md#canonical-dd1bd31a95e67d7324b79547dcf212c3ece23b79f1d8be24343b53effe8e03f7)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
