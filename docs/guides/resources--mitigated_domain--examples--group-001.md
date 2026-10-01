---
page_title: "xcsh_mitigated_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain examples."
---

# xcsh_mitigated_domain examples

<a id="canonical-f8cc4ca693157d4c768911d71d3f9378ab6faa7e3fb24d78c4c693a8ced45d35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecfe512d98073e4c2eb546936be350bdfd35e2fc81d59ce0d9571f93c86c42e8"></a>

## Examples — Examples / 8f57957434c5 / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)
- Examples

<a id="canonical-ff5524d1e4d0a7d997f1bd5225f01797a64cdba351a43f3b899f1553908a2809"></a>

## Complete configurations — Examples / 8f57957434c5 / 3

- [Resource](resources--mitigated_domain--examples--group-001.md#canonical-d5aba5b4869d6ddc89926d24fe238509fd425da33b51ccef0c59e71090856039): valid configuration.

<a id="canonical-c0497fcb82a8c5c890e435c72d51fc641680da22e69b62c41d6bb8da74d0f31c"></a>

## Next pages — Examples / 8f57957434c5 / 4

- [Resource](resources--mitigated_domain--examples--group-001.md#canonical-d5aba5b4869d6ddc89926d24fe238509fd425da33b51ccef0c59e71090856039)
- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)

<a id="canonical-d5aba5b4869d6ddc89926d24fe238509fd425da33b51ccef0c59e71090856039"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dcb9e6159f39d42a308f65343e4598d1f1d1c5fb4ddf2b82d3f8e2bfa9b4fa5"></a>

## Resource — Resource / 518ed0d184db / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)
- [Examples](resources--mitigated_domain--examples--group-001.md#canonical-f8cc4ca693157d4c768911d71d3f9378ab6faa7e3fb24d78c4c693a8ced45d35)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_mitigated_domain/resource.tf`; digest `sha256:d859b6b9d3a8006d5c27b6cad66087687ae1b57c2b118afe91822b84dab15960`.

```terraform
# MitigatedDomain Resource Example
# Manages Mitigated Domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MitigatedDomain configuration
resource "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"

  mitigated_domain = "example-value"
}
```

<a id="canonical-78815d4902ccd9ae66e9a44de2dfae6741cb18b4930af83f77b77aef6e17fad7"></a>

## Next pages — Resource / 518ed0d184db / 3

- [Examples](resources--mitigated_domain--examples--group-001.md#canonical-f8cc4ca693157d4c768911d71d3f9378ab6faa7e3fb24d78c4c693a8ced45d35)
- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)
