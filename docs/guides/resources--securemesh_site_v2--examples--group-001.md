---
page_title: "xcsh_securemesh_site_v2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 examples."
---

# xcsh_securemesh_site_v2 examples

<a id="canonical-73c06db8e151ec7d688d6a0f99293e7afad26d835f5b001586581f6dcf25331f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2fb926f1be5c8e64b9ba3589ed716c4d91a4858876c74af8a1bbacf15f4f144"></a>

## Examples — Examples / 1a663e635fb2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- Examples

<a id="canonical-bc9e883179f8d2ec02133acf898f44129df9e22fb6c2b0ad971671ea2b427158"></a>

## Complete configurations — Examples / 1a663e635fb2 / 3

- [Resource](resources--securemesh_site_v2--examples--group-001.md#canonical-0bfee1e897b09d1bd8c6007acecb283a1f36b68df33a51f122a7ea6ad7e2e69b): valid configuration.

<a id="canonical-82336f5d05ae35ecaa968f3f6378fcf9bfa2081e6d87bff9f23f3a5c54ee4776"></a>

## Next pages — Examples / 1a663e635fb2 / 4

- [Resource](resources--securemesh_site_v2--examples--group-001.md#canonical-0bfee1e897b09d1bd8c6007acecb283a1f36b68df33a51f122a7ea6ad7e2e69b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0bfee1e897b09d1bd8c6007acecb283a1f36b68df33a51f122a7ea6ad7e2e69b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c73bcfcbe092c7428ab52b722b2f70cd4634d85df9a76f368d098eb1c6e5f527"></a>

## Resource — Resource / 9037641a1bb5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Examples](resources--securemesh_site_v2--examples--group-001.md#canonical-73c06db8e151ec7d688d6a0f99293e7afad26d835f5b001586581f6dcf25331f)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_securemesh_site_v2/resource.tf`; digest `sha256:dd55f175aeb63b8fcdedf3ae50d54522d7f389806ca40efe077ad7621a9cf6e2`.

```terraform
# SecuremeshSiteV2 Resource Example
# Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites with security and networking controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSiteV2 configuration
resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}
```

<a id="canonical-e14c6e5014c148d24c734150359002481449cf4edcba01b8bb005227f17418cc"></a>

## Next pages — Resource / 9037641a1bb5 / 3

- [Examples](resources--securemesh_site_v2--examples--group-001.md#canonical-73c06db8e151ec7d688d6a0f99293e7afad26d835f5b001586581f6dcf25331f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
