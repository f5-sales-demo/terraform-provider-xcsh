---
page_title: "xcsh_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site examples."
---

# xcsh_site examples

<a id="canonical-1131302300230133-1313031231101112-3002302221232002-2310311022331312-3010032300202333-3211222122213210-3303330020123231-2121302003130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-3103121123130212-2311101033212231-0302320211230210-2003212033020122-0110200112201132-1223021030222031-3113123130220130-2002110333013120)
- Examples

<a id="canonical-1122213003221310-0303201022101132-2022021121323033-1311030113023111-3023313102010010-3300230302122020-3130233203120303-3333031010132010"></a>

### Complete configurations for `xcsh_site`

- [Data source](data-sources--site--examples--group-001.md#canonical-3112312311123332-3110001102030110-2000232122320221-3102311002112103-0110201010112003-3110010033013022-3130233000000112-3122130202232102): valid configuration.

<a id="canonical-3112312311123332-3110001102030110-2000232122320221-3102311002112103-0110201010112003-3110010033013022-3130233000000112-3122130202232102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-3103121123130212-2311101033212231-0302320211230210-2003212033020122-0110200112201132-1223021030222031-3113123130220130-2002110333013120)
- [Examples](data-sources--site--examples--group-001.md#canonical-1131302300230133-1313031231101112-3002302221232002-2310311022331312-3010032300202333-3211222122213210-3303330020123231-2121302003130122)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site/data-source.tf`; digest `sha256:3e384d82af36bf37fe3ff1c335c8824183b28374dc3c5fb9a56016a0094aa348`.

```terraform
# Site Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Site by name
data "xcsh_site" "example" {
  name      = "example-site"
  namespace = "staging"
}

output "site_id" {
  value = data.xcsh_site.example.id
}
```
