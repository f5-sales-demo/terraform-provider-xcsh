---
page_title: "xcsh_public_ip examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_public_ip examples."
---

# xcsh_public_ip examples

<a id="canonical-66ff80522379979087e1b46893f2f83bfdb92352b04293cec0b8fa51c4d25c3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cffb17325e1ea8675b6ae7a8d683b16d6a584e3d720d99fe7d928249c01fe47"></a>

## Examples — Examples / 7ea288ef5b7e / 2

Breadcrumbs:

- [xcsh_public_ip](../data-sources/public_ip.md#canonical-4d58cc44b5d0834da51e718a5f01b7b87916d0e39e4050de470be23a21d6e26a)
- Examples

<a id="canonical-bf4a83d07241683ef587d7e57ce8c3eeb01eee889f830156ed69e94436d287fb"></a>

## Complete configurations — Examples / 7ea288ef5b7e / 3

- [Data source](data-sources--public_ip--examples--group-001.md#canonical-aa06a0e79a06b89607978efc9a336705488bc9789acf848f8a69583507cda690): valid configuration.

<a id="canonical-c8c3f0ba0655b7e79bcff5e8b3a60b8cad40776511b3f34b7fffdbf1a3202b99"></a>

## Next pages — Examples / 7ea288ef5b7e / 4

- [Data source](data-sources--public_ip--examples--group-001.md#canonical-aa06a0e79a06b89607978efc9a336705488bc9789acf848f8a69583507cda690)
- [xcsh_public_ip](../data-sources/public_ip.md#canonical-4d58cc44b5d0834da51e718a5f01b7b87916d0e39e4050de470be23a21d6e26a)

<a id="canonical-aa06a0e79a06b89607978efc9a336705488bc9789acf848f8a69583507cda690"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11aaa0eeb51f97a1446f9703a24c6398d3dc02af8200eb5287c92737ace1fb18"></a>

## Data source — Data source / e0b9aac24963 / 2

Breadcrumbs:

- [xcsh_public_ip](../data-sources/public_ip.md#canonical-4d58cc44b5d0834da51e718a5f01b7b87916d0e39e4050de470be23a21d6e26a)
- [Examples](data-sources--public_ip--examples--group-001.md#canonical-66ff80522379979087e1b46893f2f83bfdb92352b04293cec0b8fa51c4d25c3f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_public_ip/data-source.tf`; digest `sha256:3249669279b5070f3336fdc2ea70d7926ac4af4c76c377a1b4478e268c3c5d4d`.

```terraform
# PublicIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PublicIP by name
data "xcsh_public_ip" "example" {
  name      = "example-public-ip"
  namespace = "staging"
}

output "public_ip_id" {
  value = data.xcsh_public_ip.example.id
}
```

<a id="canonical-ecdedca6a51cc1a434cd231788815aa72f581d490327b7a367d1a80a9a16882f"></a>

## Next pages — Data source / e0b9aac24963 / 3

- [Examples](data-sources--public_ip--examples--group-001.md#canonical-66ff80522379979087e1b46893f2f83bfdb92352b04293cec0b8fa51c4d25c3f)
- [xcsh_public_ip](../data-sources/public_ip.md#canonical-4d58cc44b5d0834da51e718a5f01b7b87916d0e39e4050de470be23a21d6e26a)
