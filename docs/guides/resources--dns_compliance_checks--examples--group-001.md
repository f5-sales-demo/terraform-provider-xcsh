---
page_title: "xcsh_dns_compliance_checks examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks examples."
---

# xcsh_dns_compliance_checks examples

<a id="canonical-eb430301335b9fce6c29bb4e206953f38e52742ac4f0703a58f37fdb51b3d161"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8fc5bab2fdf7219a05521688958f91ca7964b1c459683d0b83e291b2a552eaa"></a>

## Examples — Examples / 80b633462cb3 / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)
- Examples

<a id="canonical-44f2515689f3149a9aff99821ef105f3f1646a925c3234cb1b1c9119691c57fa"></a>

## Complete configurations — Examples / 80b633462cb3 / 3

- [Resource](resources--dns_compliance_checks--examples--group-001.md#canonical-3403bd104b63c43a3f40329c6d3def5602cc0bc34c82a3aff846436e8abbb331): valid configuration.

<a id="canonical-597bfecf30f4ad6fdf404720845f6c5f380a0494b37417ebdfbec1a55b0a119e"></a>

## Next pages — Examples / 80b633462cb3 / 4

- [Resource](resources--dns_compliance_checks--examples--group-001.md#canonical-3403bd104b63c43a3f40329c6d3def5602cc0bc34c82a3aff846436e8abbb331)
- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)

<a id="canonical-3403bd104b63c43a3f40329c6d3def5602cc0bc34c82a3aff846436e8abbb331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86d59c653f900b7319fd21e2da903341120133576a8251dcc6d4a761a755346a"></a>

## Resource — Resource / ff285222ed19 / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)
- [Examples](resources--dns_compliance_checks--examples--group-001.md#canonical-eb430301335b9fce6c29bb4e206953f38e52742ac4f0703a58f37fdb51b3d161)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_compliance_checks/resource.tf`; digest `sha256:4b55554787f8c66271c8f849e8f82effc0413dcd79554cefe77c85a2e413eba1`.

```terraform
# DNSComplianceChecks Resource Example
# Manages DNS Compliance Checks Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSComplianceChecks configuration
resource "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"

  domain_denylist = ["example-value"]
}
```

<a id="canonical-d1b4e4c3b0f29979742a8797a0072e89fb05731e006b006847b44974a8f289ae"></a>

## Next pages — Resource / ff285222ed19 / 3

- [Examples](resources--dns_compliance_checks--examples--group-001.md#canonical-eb430301335b9fce6c29bb4e206953f38e52742ac4f0703a58f37fdb51b3d161)
- [xcsh_dns_compliance_checks](../resources/dns_compliance_checks.md#canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924)
