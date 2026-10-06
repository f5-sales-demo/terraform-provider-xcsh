---
page_title: "xcsh_cloud_link examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link examples."
---

# xcsh_cloud_link examples

<a id="canonical-3223203122101022-3131311230130300-2110032021232020-1231210112230010-1102111103020220-0113113113231101-1320200100312123-1100101312003100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- Examples

<a id="canonical-0120320303010120-3331032012032213-2033032001011123-0311130010312312-0220201232003230-3302131132022303-0112031100131103-2013022333322331"></a>

### Complete configurations for `xcsh_cloud_link`

- [Resource](resources--cloud_link--examples--group-001.md#canonical-1010233102112003-1120030021313310-2013011023303020-3201222103232003-2211131202122210-2010333023323331-2231022231232122-3023120032022030): valid configuration.

<a id="canonical-1010233102112003-1120030021313310-2013011023303020-3201222103232003-2211131202122210-2010333023323331-2231022231232122-3023120032022030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Examples](resources--cloud_link--examples--group-001.md#canonical-3223203122101022-3131311230130300-2110032021232020-1231210112230010-1102111103020220-0113113113231101-1320200100312123-1100101312003100)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_link/resource.tf`; digest `sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00`.

```terraform
# CloudLink Resource Example
# Manages new CloudLink with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudLink configuration
resource "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}
```
