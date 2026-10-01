---
page_title: "xcsh_certificate_chain landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain landing."
---

# xcsh_certificate_chain landing

<a id="canonical-ce3d2afc41637e97fe322a943a0e26b7ffd06d1052a2a27eea5dfed15f9134f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7beea005ac7df0d00e02a8ce3d69d85f62fb295177dc6658827fb508470ca810"></a>

## xcsh_certificate_chain — xcsh_certificate_chain / 825752767c51 / 2

Breadcrumbs:

- xcsh_certificate_chain

Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for
TLS.

<a id="canonical-55418a360d7d54c1ccd91bba62726f927b7650fa99ab69ce5f23dc09f2d9a5db"></a>

## Prerequisites — xcsh_certificate_chain / 825752767c51 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-14558076b79cf78abe24c489d4f7bd565462bb0d68434c6ab6250ea341ce6f08"></a>

## Minimal configuration — xcsh_certificate_chain / 825752767c51 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CertificateChain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertificateChain by name
data "xcsh_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "staging"
}

output "certificate_chain_id" {
  value = data.xcsh_certificate_chain.example.id
}
```

<a id="canonical-6088a247a03064d8b630f132a78323f132950a4133721a15b0129ff1d26c85f5"></a>

## Root configuration — xcsh_certificate_chain / 825752767c51 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-6198c970660795bf6d7d40b241021f95c7a3d06c74190ab53473cc37c7aea681"></a>

## Next pages — xcsh_certificate_chain / 825752767c51 / 6

- [Property reference](../guides/data-sources--certificate_chain--reference--group-001.md#canonical-d9436ba22d8301c6a0f02ff16a4d8396e3f87a177d590cb144774d33852ac5b8)
- [Examples](../guides/data-sources--certificate_chain--examples--group-001.md#canonical-f4afd07874bc514a268ecf295767366e70decfe868e724cc487b9b5f454ea748)
