---
page_title: "xcsh_application_profiles examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles examples."
---

# xcsh_application_profiles examples

<a id="canonical-0121211302010133-1133122330330220-2313210032111110-0221132112232120-2100033100002131-2020331200001231-2002232202120301-1003301312102122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- Examples

<a id="canonical-1021210213330301-1232221302223212-3031213313022311-1313001231031032-3100303031102112-3322033100103031-1302200030102321-3332331020120330"></a>

### Complete configurations for `xcsh_application_profiles`

- [Resource](resources--application_profiles--examples--group-001.md#canonical-0320321111331220-0331023100121133-3000212131200102-2302232203333031-3221123230331223-2102011000000010-1100112021001032-2031132130301202): valid configuration.

<a id="canonical-0320321111331220-0331023100121133-3000212131200102-2302232203333031-3221123230331223-2102011000000010-1100112021001032-2031132130301202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Examples](resources--application_profiles--examples--group-001.md#canonical-0121211302010133-1133122330330220-2313210032111110-0221132112232120-2100033100002131-2020331200001231-2002232202120301-1003301312102122)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_application_profiles/resource.tf`; digest `sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f`.

```terraform
# ApplicationProfiles Resource Example
# Manages Application Profiles in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ApplicationProfiles configuration
resource "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}
```
