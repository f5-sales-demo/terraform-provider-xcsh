---
page_title: "xcsh_authorization_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server examples."
---

# xcsh_authorization_server examples

<a id="canonical-8fc47a5c0c28ccc2cd7f1fabecc67d23713c219f4b177fc3dca6b5ebb9377006"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e96e15967551f62fb27aa7205a925cd3fa2f02c2717f575e1a098bacefb2404"></a>

## Examples — Examples / 19136e2f190e / 2

Breadcrumbs:

- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-c9a064e11c392b1729b64c4abb5a0c14e6107920f082e94b97489dd17471f27a)
- Examples

<a id="canonical-be69f033c43fda4e177e91d13448283bc33c74e2cc5479bf7df6877409d7c645"></a>

## Complete configurations — Examples / 19136e2f190e / 3

- [Data source](data-sources--authorization_server--examples--group-001.md#canonical-cfa08471604ea6ecf291f6c34406b93c90a395a5f26df590cd4d310477c4ff0c): valid configuration.

<a id="canonical-88b23721e8faa81aad2bfd3f084174f20bad5920287ad9b8734eb7993eb4b6c9"></a>

## Next pages — Examples / 19136e2f190e / 4

- [Data source](data-sources--authorization_server--examples--group-001.md#canonical-cfa08471604ea6ecf291f6c34406b93c90a395a5f26df590cd4d310477c4ff0c)
- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-c9a064e11c392b1729b64c4abb5a0c14e6107920f082e94b97489dd17471f27a)

<a id="canonical-cfa08471604ea6ecf291f6c34406b93c90a395a5f26df590cd4d310477c4ff0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cca5808cd96c6905553888ac4f23868333fe610c82ef42c3b0669d0e6551a633"></a>

## Data source — Data source / 09d8bf426335 / 2

Breadcrumbs:

- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-c9a064e11c392b1729b64c4abb5a0c14e6107920f082e94b97489dd17471f27a)
- [Examples](data-sources--authorization_server--examples--group-001.md#canonical-8fc47a5c0c28ccc2cd7f1fabecc67d23713c219f4b177fc3dca6b5ebb9377006)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authorization_server/data-source.tf`; digest `sha256:268fe1444519cbb845d8fcb9ea236d26293e2aeafb324fad41e08e7bc571ffcb`.

```terraform
# AuthorizationServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AuthorizationServer by name
data "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"
}

output "authorization_server_id" {
  value = data.xcsh_authorization_server.example.id
}
```

<a id="canonical-afbfc25928e7dfd90fceec518017f6fd1f09c3cdbda072019d298d5bc13f8195"></a>

## Next pages — Data source / 09d8bf426335 / 3

- [Examples](data-sources--authorization_server--examples--group-001.md#canonical-8fc47a5c0c28ccc2cd7f1fabecc67d23713c219f4b177fc3dca6b5ebb9377006)
- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-c9a064e11c392b1729b64c4abb5a0c14e6107920f082e94b97489dd17471f27a)
