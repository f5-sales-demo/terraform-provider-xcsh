---
page_title: "xcsh_certificate_chain examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain examples."
---

# xcsh_certificate_chain examples

<a id="canonical-ae11625a3f3cabc566632a4055fe0cbec86e2eadd028c0ed14ee98a590bd533c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62774503f47021bcddf1cbd58c5f3b30a85ceec6ce91144f78f03f5f1a881229"></a>

## Examples — Examples / aa07abe3666b / 2

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)
- Examples

<a id="canonical-37798d9a128f9d234f405048b05afe9682cd87ca588935b819b7e5d4547ed55b"></a>

## Complete configurations — Examples / aa07abe3666b / 3

- [Resource](resources--certificate_chain--examples--group-001.md#canonical-b42df567f26e092d50b038aaa78a56662980049695d6088d078a03e469d97e29): valid configuration.

<a id="canonical-576fc18cc92214f959db58d42ad59c6b467a4c35ba4eaec350eabdcfc9696260"></a>

## Next pages — Examples / aa07abe3666b / 4

- [Resource](resources--certificate_chain--examples--group-001.md#canonical-b42df567f26e092d50b038aaa78a56662980049695d6088d078a03e469d97e29)
- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)

<a id="canonical-b42df567f26e092d50b038aaa78a56662980049695d6088d078a03e469d97e29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34e6ef59cf56a25d3ee2303c2256cff17f71def7a8fbe999bc189d3ebeb9ea98"></a>

## Resource — Resource / 981e295b67b3 / 2

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)
- [Examples](resources--certificate_chain--examples--group-001.md#canonical-ae11625a3f3cabc566632a4055fe0cbec86e2eadd028c0ed14ee98a590bd533c)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate_chain/resource.tf`; digest `sha256:b97518ab6ca31864cd9125be1aa4c371e7fb2f6d12ef9376188442d135d4d743`.

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

<a id="canonical-1eff5d5696b14b0dd42625fd8ac6e0211e26c37a714882dd1766449a725b4056"></a>

## Next pages — Resource / 981e295b67b3 / 3

- [Examples](resources--certificate_chain--examples--group-001.md#canonical-ae11625a3f3cabc566632a4055fe0cbec86e2eadd028c0ed14ee98a590bd533c)
- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)
