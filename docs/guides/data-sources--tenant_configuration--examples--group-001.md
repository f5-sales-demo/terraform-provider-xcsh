---
page_title: "xcsh_tenant_configuration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration examples."
---

# xcsh_tenant_configuration examples

<a id="canonical-3233212102312300-1110101220100003-2133212022310023-0100210120032012-3222033131310000-3321000113321031-3333220321102300-3232122211111221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md#canonical-1230121131133133-0211313013212213-1101132331100023-2203020001230102-0301100132002332-3002232111300122-2302311100323013-2003322300110201)
- Examples

<a id="canonical-0322313323202022-1312221000013302-0100032232212210-1311101332311103-2010130202310003-1102202302313002-0012201310301002-2322322033112221"></a>

### Complete configurations for `xcsh_tenant_configuration`

- [Data source](data-sources--tenant_configuration--examples--group-001.md#canonical-2231303120332320-0132300223331320-1323001201212331-3123023223222100-3323110013200100-2202223200203323-2010113102333232-1023130203323132): valid configuration.

<a id="canonical-2231303120332320-0132300223331320-1323001201212331-3123023223222100-3323110013200100-2202223200203323-2010113102333232-1023130203323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_tenant_configuration](../data-sources/tenant_configuration.md#canonical-1230121131133133-0211313013212213-1101132331100023-2203020001230102-0301100132002332-3002232111300122-2302311100323013-2003322300110201)
- [Examples](data-sources--tenant_configuration--examples--group-001.md#canonical-3233212102312300-1110101220100003-2133212022310023-0100210120032012-3222033131310000-3321000113321031-3333220321102300-3232122211111221)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tenant_configuration/data-source.tf`; digest `sha256:82e6c50873cbaa19717f3b339ef9fae150aab727845b88eb85553bf86a8cc317`.

```terraform
# TenantConfiguration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TenantConfiguration by name
data "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}

output "tenant_configuration_id" {
  value = data.xcsh_tenant_configuration.example.id
}
```
