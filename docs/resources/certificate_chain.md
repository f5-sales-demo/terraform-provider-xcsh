---
page_title: "xcsh_certificate_chain landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain landing."
---

# xcsh_certificate_chain landing

<a id="canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75aae910d1a5c2e5d7f1ce23f82c84e0295b6e3a57207988360abb091a5967c4"></a>

## xcsh_certificate_chain — xcsh_certificate_chain / 12aec357a4b5 / 2

Breadcrumbs:

- xcsh_certificate_chain

Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for
TLS.

<a id="canonical-b7c689d62c44fdc3e54b649b42e25c70241b625418d5f1a9038bef048df22ca9"></a>

## Prerequisites — xcsh_certificate_chain / 12aec357a4b5 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-c17a9060605ce5b044938a359a7636cfa2e6f96f99f52f2927679384e5aa2b04"></a>

## Minimal configuration — xcsh_certificate_chain / 12aec357a4b5 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CertificateChain Resource Example
# Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for TLS.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CertificateChain configuration
resource "xcsh_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "staging"

  certificate_url = "example-value"
}
```

<a id="canonical-8753d578f90be59c7d4cd2bf4fae55cb14fe8728415dbeb7d07d76186052c86f"></a>

## Root configuration — xcsh_certificate_chain / 12aec357a4b5 / 5

Required root properties: `certificate_url`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-7ca9944a9c76f1fafee2f901ba9a929211644a604d464cbb2b1cfb081061c154"></a>

## Next pages — xcsh_certificate_chain / 12aec357a4b5 / 6

- [Property reference](../guides/resources--certificate_chain--reference--group-001.md#canonical-c976c618810c3d923d3d6cdad9e66893de920464905dd953c308630b99c54079)
- [Examples](../guides/resources--certificate_chain--examples--group-001.md#canonical-ae11625a3f3cabc566632a4055fe0cbec86e2eadd028c0ed14ee98a590bd533c)
- [Import](../guides/resources--certificate_chain--lifecycle--group-001.md#canonical-7d09b404d603ce47b91ca76b13ee58644070bdaec24e088be28e2a4b918f5495)
- [Timeouts](../guides/resources--certificate_chain--lifecycle--group-001.md#canonical-5b376bc5d39d1210832247605041b8f4eb18e3ffe1ff4b0ec092a1f0256cadbf)
