---
page_title: "xcsh_dns_compliance_checks landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks landing."
---

# xcsh_dns_compliance_checks landing

<a id="canonical-0df8fdb3734c554a4a4bef589f19acd0c3a88fa45c35acb7d5923f8e5221913a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-944f4c14f076c05f5bbade701280c0d0ac4a2e1755b22670ea449e835e86a522"></a>

## xcsh_dns_compliance_checks — xcsh_dns_compliance_checks / b96ebfdb45d1 / 2

Breadcrumbs:

- xcsh_dns_compliance_checks

Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

<a id="canonical-bda72e591d82b8f20c1866bd8ad84c42d029922b294fe088dac54f45171c63ee"></a>

## Prerequisites — xcsh_dns_compliance_checks / b96ebfdb45d1 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6438411de113570233f7719c8c32064aa0b204c820b51d12fc0ae7ef4c4807c5"></a>

## Minimal configuration — xcsh_dns_compliance_checks / b96ebfdb45d1 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSComplianceChecks Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSComplianceChecks by name
data "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"
}

output "dns_compliance_checks_id" {
  value = data.xcsh_dns_compliance_checks.example.id
}
```

<a id="canonical-76944ea472dadf13248ba14c9dd9d37d8ecc92926f13e039e3fcf131beedd866"></a>

## Root configuration — xcsh_dns_compliance_checks / b96ebfdb45d1 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-827bc211d853b3902048e713f098569a6b99a9d1c1298d5401fe718da108d89d"></a>

## Next pages — xcsh_dns_compliance_checks / b96ebfdb45d1 / 6

- [Property reference](../guides/data-sources--dns_compliance_checks--reference--group-001.md#canonical-3c742fd4796a81209dac248974569ee528b2f8284b59ae8fbf6e8838bda5854f)
- [Examples](../guides/data-sources--dns_compliance_checks--examples--group-001.md#canonical-4d7fc5bbbfd3c706332c0069b37f53e2833a6ae6f0cfc87068d5fd48be9cbaab)
