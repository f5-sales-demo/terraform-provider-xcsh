---
page_title: "xcsh_tenant_configuration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration examples."
---

# xcsh_tenant_configuration examples

<a id="canonical-3301102230300102-1113221230011001-0313201222210020-0100103213313001-3101010331113231-1023101022310003-1011322233030322-2330023033203213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- Examples

<a id="canonical-0312002033000203-1231202003322102-1300023213010303-0120033100101130-2023030111123203-3202130230222013-3212112210001012-3000213332300333"></a>

### Complete configurations for `xcsh_tenant_configuration`

- [Resource](resources--tenant_configuration--examples--group-001.md#canonical-0133311010123103-1113111120010223-2203223032033002-3120223101003000-2230220203303012-0210313323231230-3012203203332021-3022120230310103): valid configuration.

<a id="canonical-0133311010123103-1113111120010223-2203223032033002-3120223101003000-2230220203303012-0210313323231230-3012203203332021-3022120230310103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310)
- [Examples](resources--tenant_configuration--examples--group-001.md#canonical-3301102230300102-1113221230011001-0313201222210020-0100103213313001-3101010331113231-1023101022310003-1011322233030322-2330023033203213)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tenant_configuration/resource.tf`; digest `sha256:649c10c3cc2ff997128da4475c6c82fbfad1826ae6d86d2d3e3332f24b98496b`.

```terraform
# TenantConfiguration Resource Example
# Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TenantConfiguration configuration
resource "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}
```
