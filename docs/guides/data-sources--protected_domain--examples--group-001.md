---
page_title: "xcsh_protected_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain examples."
---

# xcsh_protected_domain examples

<a id="canonical-3030223231010012-3322000022222331-0212010302300020-3010113221023011-3121313131011111-0203010312301030-0122222022013100-1201023111230012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-0332111003332231-3231333320302220-0013131312310303-3021021201222321-0302322122212110-0022133313313221-1010303303020331-2102022101100033)
- Examples

<a id="canonical-0011012202203210-0100313031013000-3223132133310301-1120220332330313-3201311000020101-0011320121332032-2003211122011011-2221231210002112"></a>

### Complete configurations for `xcsh_protected_domain`

- [Data source](data-sources--protected_domain--examples--group-001.md#canonical-0331333123202210-0111123103133310-3212211032231312-0332100021230212-0312010101102121-1210000212010123-0332123323203200-2022111221211102): valid configuration.

<a id="canonical-0331333123202210-0111123103133310-3212211032231312-0332100021230212-0312010101102121-1210000212010123-0332123323203200-2022111221211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-0332111003332231-3231333320302220-0013131312310303-3021021201222321-0302322122212110-0022133313313221-1010303303020331-2102022101100033)
- [Examples](data-sources--protected_domain--examples--group-001.md#canonical-3030223231010012-3322000022222331-0212010302300020-3010113221023011-3121313131011111-0203010312301030-0122222022013100-1201023111230012)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protected_domain/data-source.tf`; digest `sha256:54de5a860ec2aa40dfd3370e44fe450b1b713886394cacf655ec2b67b6b0edd5`.

```terraform
# ProtectedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedDomain by name
data "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"
}

output "protected_domain_id" {
  value = data.xcsh_protected_domain.example.id
}
```
