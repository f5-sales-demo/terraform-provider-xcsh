---
page_title: "xcsh_dns_compliance_checks landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks landing."
---

# xcsh_dns_compliance_checks landing

<a id="canonical-b8f0d3ec4e785ef6fabb81b92471aa91dd33fcda47717d8aa7b007104462e924"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8f9bb1c60d5b493400b522b2afa60491012627cc6c321706ca5c943df301fda"></a>

## xcsh_dns_compliance_checks — xcsh_dns_compliance_checks / ec8e5e56b845 / 2

Breadcrumbs:

- xcsh_dns_compliance_checks

Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

<a id="canonical-371c9346578d419351e2a7e31b11020b6c75b5560c5d7e439c3a3f049431c500"></a>

## Prerequisites — xcsh_dns_compliance_checks / ec8e5e56b845 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a21ef4cab2c9aadf28ead3eabcedcb348ee06bf59d1cf75a1ec8250e67fe5a18"></a>

## Minimal configuration — xcsh_dns_compliance_checks / ec8e5e56b845 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1ab5fa5bf768f6bd84ab05165bc4e2179398c861d954ae8e02e47cbe1421535a"></a>

## Root configuration — xcsh_dns_compliance_checks / ec8e5e56b845 / 5

Required root properties: `domain_denylist`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-7b3353f6ac1681d37efd819ecc77d58d71b3923cfc6066499f07e7576a9b7b88"></a>

## Next pages — xcsh_dns_compliance_checks / ec8e5e56b845 / 6

- [Property reference](../guides/resources--dns_compliance_checks--reference--group-001.md#canonical-eee1b4d710d52db719c841deb6e007376a78ef3ab29c2e2a6fc1e96dd6936c22)
- [Examples](../guides/resources--dns_compliance_checks--examples--group-001.md#canonical-eb430301335b9fce6c29bb4e206953f38e52742ac4f0703a58f37fdb51b3d161)
- [Import](../guides/resources--dns_compliance_checks--lifecycle--group-001.md#canonical-3e8c45883cd1a3948b530d5d66eb509661ed3222c493c43b4a23e42b133d76b1)
- [Timeouts](../guides/resources--dns_compliance_checks--lifecycle--group-001.md#canonical-b6648163dc46cf44b575809fa12f03e825820f2f5bcaf9ac5820452ffa5bfc13)
