---
page_title: "xcsh_certificate landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate landing."
---

# xcsh_certificate landing

<a id="canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0feb913bb4c34238e7d4480278ba2b62ccb1653f52b74897a0678fc36d8c540"></a>

## xcsh_certificate — xcsh_certificate / 2055c7ec1dd2 / 2

Breadcrumbs:

- xcsh_certificate

Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.

<a id="canonical-9406614af30287ac737757d5fddfc3fd8f58ceb8ba60ed41687fa9ea8f899454"></a>

## Prerequisites — xcsh_certificate / 2055c7ec1dd2 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-64019377b1059b3536e34a8697168237ed9987bda0e9e7b95d99101149f2465e"></a>

## Minimal configuration — xcsh_certificate / 2055c7ec1dd2 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```

<a id="canonical-6955c3664fbef5fc38e3bc5fa960facb4f1485af5196540843eaa9556cc0aa47"></a>

## Root configuration — xcsh_certificate / 2055c7ec1dd2 / 5

Required root properties: `certificate_url`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e088d62478a9c7000a553076580d468c3ce5a0b9e76007c26c9fd86bde13764a"></a>

## Next pages — xcsh_certificate / 2055c7ec1dd2 / 6

- [Property reference](../guides/resources--certificate--reference--group-001.md#canonical-6e191e632e125ee6cf3e18d3e730166767b3b520ab1f5c6d147e575ffbede87c)
- [Examples](../guides/resources--certificate--examples--group-001.md#canonical-fadf3f9d1d214e0a699b1607ad361706f2c97596ce1ac016dffdb3f418d41ffe)
- [Import](../guides/resources--certificate--lifecycle--group-001.md#canonical-51bcd7cddc02f66ed213b762d1cc278d0616a168948b32327351b508f94ab0a4)
- [Timeouts](../guides/resources--certificate--lifecycle--group-001.md#canonical-090bdaa9a218102a1c0c4356f0385466b05f5294138c2b24943240fb8d709bb5)
