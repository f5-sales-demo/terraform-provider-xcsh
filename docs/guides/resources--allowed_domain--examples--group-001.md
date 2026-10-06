---
page_title: "xcsh_allowed_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain examples."
---

# xcsh_allowed_domain examples

<a id="canonical-0100103011021332-2023322133032001-2333102002332112-0211100100330332-3023211020232111-2132102123233230-3312033032113322-3030130003223213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300)
- Examples

<a id="canonical-0333222013111223-1311312112312123-2122000322201232-1333121011110001-0232301230101010-1100323331333102-2031330011123110-3313322033313233"></a>

### Complete configurations for `xcsh_allowed_domain`

- [Resource](resources--allowed_domain--examples--group-001.md#canonical-3131212011321233-2111320300320200-2323102123033320-2000223303201032-2010330322312102-3301211022032010-2232031213111110-2320112201230211): valid configuration.

<a id="canonical-3131212011321233-2111320300320200-2323102123033320-2000223303201032-2010330322312102-3301211022032010-2232031213111110-2320112201230211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300)
- [Examples](resources--allowed_domain--examples--group-001.md#canonical-0100103011021332-2023322133032001-2333102002332112-0211100100330332-3023211020232111-2132102123233230-3312033032113322-3030130003223213)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_allowed_domain/resource.tf`; digest `sha256:311173eaa23e9e9afb5856fa5a592866eb29afd044fe81b255b02bc483f4948e`.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```
