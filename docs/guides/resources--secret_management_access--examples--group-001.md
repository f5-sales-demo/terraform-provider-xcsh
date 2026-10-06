---
page_title: "xcsh_secret_management_access examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access examples."
---

# xcsh_secret_management_access examples

<a id="canonical-2111202301110310-0313132001022331-3123020103123130-2301210312131331-1013011323103223-2100201310303012-1333233323020312-3022111021121303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- Examples

<a id="canonical-1301022100202002-3232233200000032-1121331010301321-2233211100223211-3103202123221010-3221021131222213-0031313211233303-0003301211023300"></a>

### Complete configurations for `xcsh_secret_management_access`

- [Resource](resources--secret_management_access--examples--group-001.md#canonical-0303230202110201-0203211233311100-3302130311223000-2223101332121131-3312313012232330-1131211023302311-0201003013011023-2331121003023322): valid configuration.

<a id="canonical-0303230202110201-0203211233311100-3302130311223000-2223101332121131-3312313012232330-1131211023302311-0201003013011023-2331121003023322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Examples](resources--secret_management_access--examples--group-001.md#canonical-2111202301110310-0313132001022331-3123020103123130-2301210312131331-1013011323103223-2100201310303012-1333233323020312-3022111021121303)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_secret_management_access/resource.tf`; digest `sha256:49f920527939852314c086a0e57abc992c317f684e7845daaf75a5df25aa0d81`.

```terraform
# SecretManagementAccess Resource Example
# Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecretManagementAccess configuration
resource "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"

  provider_name = "example-value"
}
```
