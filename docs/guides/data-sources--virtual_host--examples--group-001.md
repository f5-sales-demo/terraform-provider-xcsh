---
page_title: "xcsh_virtual_host examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host examples."
---

# xcsh_virtual_host examples

<a id="canonical-0100001331112310-1230100011301031-2230200220313033-0002101033032003-3322002302312013-3120010230021103-0322303302110022-2321033300201112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- Examples

<a id="canonical-0312032313002213-1221032112111211-2301111102001221-1201101112301320-0010002101303200-3202012020001002-3113032331313212-0103210312232200"></a>

### Complete configurations for `xcsh_virtual_host`

- [Data source](data-sources--virtual_host--examples--group-001.md#canonical-3310330001110212-1232132132312200-0232230121322300-3003102003212202-3022012012310032-2323033202222022-2133023021331231-1003112013212303): valid configuration.

<a id="canonical-3310330001110212-1232132132312200-0232230121322300-3003102003212202-3022012012310032-2323033202222022-2133023021331231-1003112013212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220)
- [Examples](data-sources--virtual_host--examples--group-001.md#canonical-0100001331112310-1230100011301031-2230200220313033-0002101033032003-3322002302312013-3120010230021103-0322303302110022-2321033300201112)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_host/data-source.tf`; digest `sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72`.

```terraform
# VirtualHost Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualHost by name
data "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}

output "virtual_host_id" {
  value = data.xcsh_virtual_host.example.id
}
```
