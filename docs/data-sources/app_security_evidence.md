---
page_title: "xcsh_app_security_evidence landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_security_evidence landing."
---

# xcsh_app_security_evidence landing

<a id="canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2318546f473b4e398b02f5b17e6ca45271cce146bb715fd7246b10820b8317c2"></a>

## xcsh_app_security_evidence — xcsh_app_security_evidence / 384fc9dc2e1e / 2

Breadcrumbs:

- xcsh_app_security_evidence

Resource creation operation.

<a id="canonical-c89e6b403d3c784b53b5c1df66e4587c295a435070686ea1f6955b04219711f6"></a>

## Prerequisites — xcsh_app_security_evidence / 384fc9dc2e1e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-03939f5274871377a83117d6b1b5f1fce17e96c7d4a30764e64bccb4cd3310eb"></a>

## Minimal configuration — xcsh_app_security_evidence / 384fc9dc2e1e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSecurityEvidence DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_app_security_evidence" "example" {
  namespace = "example-value"
}

output "app_security_evidence_result" {
  value = data.xcsh_app_security_evidence.example
}
```

<a id="canonical-5397d8454ff32c27ed54847d6d1c91f4b2a91972d58959abfcb25492ff201e80"></a>

## Root configuration — xcsh_app_security_evidence / 384fc9dc2e1e / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-043879669a98584a6829bbec401626a0ec308d8e414870b6e43eb06f42642bbf"></a>

## Next pages — xcsh_app_security_evidence / 384fc9dc2e1e / 6

- [Property reference](../guides/data-sources--app_security_evidence--reference--group-001.md#canonical-160f5b6ebef4ab35b1510dc3be35d8747129089705ae396e4d0fa10d3481c008)
- [Examples](../guides/data-sources--app_security_evidence--examples--group-001.md#canonical-55f5d1c3e7d73954b71f9df8f5d4b0cb8083e568ca0b2d0e5ad0365d3df75be6)
