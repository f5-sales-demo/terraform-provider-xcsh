---
page_title: "xcsh_mitigated_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain landing."
---

# xcsh_mitigated_domain landing

<a id="canonical-efcfbe08e10eae8cd34e645079b3b2ddf4a333bb259bfc7f7c6325c8868a5fb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1497ace92e20a28c9ac483ec8ba247f6ec0f22955d4d7995585cbd827d45f83d"></a>

## xcsh_mitigated_domain — xcsh_mitigated_domain / f42ccf8e0132 / 2

Breadcrumbs:

- xcsh_mitigated_domain

Manages Mitigated Domain in F5 Distributed Cloud.

<a id="canonical-9a5661a4a545cc5d4e6415c78f2dcdfdb71c64ddb51a5279d14b77369c4ea150"></a>

## Prerequisites — xcsh_mitigated_domain / f42ccf8e0132 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1a8e78f81806081982001ca24b027c25488cc76258a74768984c80b60c64e0c5"></a>

## Minimal configuration — xcsh_mitigated_domain / f42ccf8e0132 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MitigatedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MitigatedDomain by name
data "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"
}

output "mitigated_domain_id" {
  value = data.xcsh_mitigated_domain.example.id
}
```

<a id="canonical-9f78d83eae3da9adac3c9b3b7e5efb5c0d3f837faa706ad0efdcdc93aef75b7e"></a>

## Root configuration — xcsh_mitigated_domain / f42ccf8e0132 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0945ee04fc55bbeb6551a04cecd52499950ca471fd8c11000d613e7a797cb129"></a>

## Next pages — xcsh_mitigated_domain / f42ccf8e0132 / 6

- [Property reference](../guides/data-sources--mitigated_domain--reference--group-001.md#canonical-219ebdfe9e83d72a62a01c844ad18ac7c83ab7629b5a58cda760c069bc6d72e2)
- [Examples](../guides/data-sources--mitigated_domain--examples--group-001.md#canonical-6c61924190f11f3700744be4ef278107ba6451dcbc41246a1e2d9701afbbda24)
