---
page_title: "xcsh_site_registrations examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations examples."
---

# xcsh_site_registrations examples

<a id="canonical-3032020103011321-0332312120220310-2110113322301312-1231201123320220-2303332131331323-2132310013310121-1131233223031310-1313331210103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-2320011012012203-1211121021012323-3333302232010031-0232010323320310-2122323133012221-2310311330012133-3232132103233312-1122133001030203)
- Examples

<a id="canonical-3103312133303111-2333112001213222-3113101302013111-3213313003102133-0011030101212301-2211200131223102-0213110030033020-3201230030221301"></a>

### Complete configurations for `xcsh_site_registrations`

- [Data source](data-sources--site_registrations--examples--group-001.md#canonical-1032132320201220-3321013310231312-0221133002123132-0200103001330002-1330032302230000-0313003301013012-3312011020312120-1021032233230303): valid configuration.

<a id="canonical-1032132320201220-3321013310231312-0221133002123132-0200103001330002-1330032302230000-0313003301013012-3312011020312120-1021032233230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-2320011012012203-1211121021012323-3333302232010031-0232010323320310-2122323133012221-2310311330012133-3232132103233312-1122133001030203)
- [Examples](data-sources--site_registrations--examples--group-001.md#canonical-3032020103011321-0332312120220310-2110113322301312-1231201123320220-2303332131331323-2132310013310121-1131233223031310-1313331210103313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations/data-source.tf`; digest `sha256:382726257bf8ed644e4ba85a525a0636702e8ea87e1b309384adf95a431060c9`.

```terraform
# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
```
