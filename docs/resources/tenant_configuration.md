---
page_title: "xcsh_tenant_configuration"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration."
---

# xcsh_tenant_configuration

<a id="canonical-1332002001332233-1210102233002221-2012331111221332-0021003130222001-2121012132303322-1111132333301312-0120220102202112-3223020331300310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_tenant_configuration

Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration
specification. configuration.

<a id="canonical-1120211231212100-0203032332033110-2203031132211323-1123102232211012-3030111210200013-0030331230221022-0313223233200012-0001220332231300"></a>

### Prerequisites for `xcsh_tenant_configuration`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1230311130030320-3023130020013232-2110020310322133-0122213100112222-2313022200200112-0310302210031312-3102130123323301-1301233121230112"></a>

### Minimal configuration for `xcsh_tenant_configuration`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1203013311202330-0313102303301212-2320102112310301-2221101001010000-1102032232233301-3210310103122123-1312233102332333-2300003330210323"></a>

### Root configuration for `xcsh_tenant_configuration`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2302320102221331-3223111222203033-3213303102330013-2331311231201111-3321233103131211-2221131033102201-2301113021320200-3110120100300001"></a>

### Explore this collection for `xcsh_tenant_configuration`

- [Property reference](../guides/resources--tenant_configuration--reference--group-001.md#canonical-0112010312211210-2330131301100232-0100013221012023-0100320210300013-2333113013332121-3020301131332100-2310323331122100-1302302023032202)
- [Examples](../guides/resources--tenant_configuration--examples--group-001.md#canonical-3301102230300102-1113221230011001-0313201222210020-0100103213313001-3101010331113231-1023101022310003-1011322233030322-2330023033203213)
- [Import](../guides/resources--tenant_configuration--lifecycle--group-001.md#canonical-2111222301210110-0322113322233301-0220023023110102-0313333133103330-2123200121120022-0013200202301033-2001310232311332-2322212221312132)
- [Timeouts](../guides/resources--tenant_configuration--lifecycle--group-001.md#canonical-3000203121022221-0322312110331130-3221022112320110-1202021120303011-0112300123222302-2120333232302311-1000223012223200-3301202000233023)
