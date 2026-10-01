---
page_title: "xcsh_dns_compliance_checks examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks examples."
---

# xcsh_dns_compliance_checks examples

<a id="canonical-4d7fc5bbbfd3c706332c0069b37f53e2833a6ae6f0cfc87068d5fd48be9cbaab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae22e99e7967b72db357dd7439a3bfed8d8f7297bc8a6d9381ed855fcfe76c79"></a>

## Examples — Examples / 4f27cdafa700 / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0df8fdb3734c554a4a4bef589f19acd0c3a88fa45c35acb7d5923f8e5221913a)
- Examples

<a id="canonical-05595c0af51d64be216438cecfa11843a35616105bdeb1131ac336de287d2573"></a>

## Complete configurations — Examples / 4f27cdafa700 / 3

- [Data source](data-sources--dns_compliance_checks--examples--group-001.md#canonical-ebc426dc526f8ad740edd9f08873d12ed6286295e3e735d956099d97bde0a7a5): valid configuration.

<a id="canonical-977434a930aca1013f6ec95e38ece5ab783a0534e3c44c167774a13782260200"></a>

## Next pages — Examples / 4f27cdafa700 / 4

- [Data source](data-sources--dns_compliance_checks--examples--group-001.md#canonical-ebc426dc526f8ad740edd9f08873d12ed6286295e3e735d956099d97bde0a7a5)
- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0df8fdb3734c554a4a4bef589f19acd0c3a88fa45c35acb7d5923f8e5221913a)

<a id="canonical-ebc426dc526f8ad740edd9f08873d12ed6286295e3e735d956099d97bde0a7a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3bf168fd105b2dbf671a60c2993f54a76d70760ce8e18e2ea45e462c5b67abc"></a>

## Data source — Data source / b89021a41550 / 2

Breadcrumbs:

- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0df8fdb3734c554a4a4bef589f19acd0c3a88fa45c35acb7d5923f8e5221913a)
- [Examples](data-sources--dns_compliance_checks--examples--group-001.md#canonical-4d7fc5bbbfd3c706332c0069b37f53e2833a6ae6f0cfc87068d5fd48be9cbaab)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_compliance_checks/data-source.tf`; digest `sha256:494f38d0fa830fb2beff3f30865ce3a3a079a872bc62acf125c470d1e3df90d4`.

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

<a id="canonical-25738cdbbcd3fd328f7c942040a854cfdef56698868baed12a1f106c9f5ac7a4"></a>

## Next pages — Data source / b89021a41550 / 3

- [Examples](data-sources--dns_compliance_checks--examples--group-001.md#canonical-4d7fc5bbbfd3c706332c0069b37f53e2833a6ae6f0cfc87068d5fd48be9cbaab)
- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0df8fdb3734c554a4a4bef589f19acd0c3a88fa45c35acb7d5923f8e5221913a)
