---
page_title: "xcsh_cdn_cache_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule examples."
---

# xcsh_cdn_cache_rule examples

<a id="canonical-1123232111331111-0131132300333101-3320021113101203-1311110200002032-3130303103303300-0112302233013033-1211133123333011-3211232123213233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- Examples

<a id="canonical-0130210212230211-3111233113111330-0013113123331330-0222120212023022-0210321313232020-2021331210320010-1003020201022210-3030130113132330"></a>

### Complete configurations for `xcsh_cdn_cache_rule`

- [Data source](data-sources--cdn_cache_rule--examples--group-001.md#canonical-3131132023323330-2231001231222123-1230022132103031-3223031123010101-3033002301133022-0212323320223202-3122311210222302-1013111031332123): valid configuration.

<a id="canonical-3131132023323330-2231001231222123-1230022132103031-3223031123010101-3033002301133022-0212323320223202-3122311210222302-1013111031332123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Examples](data-sources--cdn_cache_rule--examples--group-001.md#canonical-1123232111331111-0131132300333101-3320021113101203-1311110200002032-3130303103303300-0112302233013033-1211133123333011-3211232123213233)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_cache_rule/data-source.tf`; digest `sha256:fc686ceb8d4fdca2d4b78094b5b3751ef65d5c692bd05f9089c3561d05468ad9`.

```terraform
# CDNCacheRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNCacheRule by name
data "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}

output "cdn_cache_rule_id" {
  value = data.xcsh_cdn_cache_rule.example.id
}
```
