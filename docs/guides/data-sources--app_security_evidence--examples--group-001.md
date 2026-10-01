---
page_title: "xcsh_app_security_evidence examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_security_evidence examples."
---

# xcsh_app_security_evidence examples

<a id="canonical-55f5d1c3e7d73954b71f9df8f5d4b0cb8083e568ca0b2d0e5ad0365d3df75be6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab18c6e2a64a02a82139e1b590be977daeaa9c9d2a0cc132a1baf20a940b28c3"></a>

## Examples — Examples / 6fd1f24ed3f1 / 2

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
- Examples

<a id="canonical-ef5a938991afe1e878229d09a1bd15e1d1fa668d92379b862372553002edb818"></a>

## Complete configurations — Examples / 6fd1f24ed3f1 / 3

- [Data source](data-sources--app_security_evidence--examples--group-001.md#canonical-49683661702847a5664ad5349a0f79b046d1f59757fb90b6965bdf54a97d905a): valid configuration.

<a id="canonical-509f5ef16cffaa785065359a9bb33d8b7b2cf727ac1098a2efebe3a2a926c89c"></a>

## Next pages — Examples / 6fd1f24ed3f1 / 4

- [Data source](data-sources--app_security_evidence--examples--group-001.md#canonical-49683661702847a5664ad5349a0f79b046d1f59757fb90b6965bdf54a97d905a)
- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)

<a id="canonical-49683661702847a5664ad5349a0f79b046d1f59757fb90b6965bdf54a97d905a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34773f97e9d391739504531a97dd61a7bdeccef4a94147e86bbc32e7ec5a7ff2"></a>

## Data source — Data source / c048d32447aa / 2

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
- [Examples](data-sources--app_security_evidence--examples--group-001.md#canonical-55f5d1c3e7d73954b71f9df8f5d4b0cb8083e568ca0b2d0e5ad0365d3df75be6)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_security_evidence/data-source.tf`; digest `sha256:a712a6cc48d2b8d3894590d0570f0c9aa69c910e418b124dca6ff57b24d59d2a`.

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

<a id="canonical-aca4905263de4f5c31742b9f0e7f166f5cbb690303e34348f276373fcdaa20a1"></a>

## Next pages — Data source / c048d32447aa / 3

- [Examples](data-sources--app_security_evidence--examples--group-001.md#canonical-55f5d1c3e7d73954b71f9df8f5d4b0cb8083e568ca0b2d0e5ad0365d3df75be6)
- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-d3c80b619742ede9129532b0a717cd17f1a275a4761a59fe318edc65c1f1d29b)
