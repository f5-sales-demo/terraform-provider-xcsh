---
page_title: "xcsh_api_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery examples."
---

# xcsh_api_discovery examples

<a id="canonical-2103330003101230-1200233231321132-1332020330220230-0211002003332330-1301223301010230-1123100002232003-0133303013301232-2212101303032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- Examples

<a id="canonical-0013012333101323-0232223313123121-0122023121210320-2322223323223313-0022131112012112-3321302323031320-0213221222031122-1133303023103230"></a>

### Complete configurations for `xcsh_api_discovery`

- [Resource](resources--api_discovery--examples--group-001.md#canonical-3302331232122212-2213300133302023-0321322321131033-2002332321201000-1000100001310330-1200312122002231-0220330100030333-0303222002202130): valid configuration.

<a id="canonical-3302331232122212-2213300133302023-0321322321131033-2002332321201000-1000100001310330-1200312122002231-0220330100030333-0303222002202130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md#canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310)
- [Examples](resources--api_discovery--examples--group-001.md#canonical-2103330003101230-1200233231321132-1332020330220230-0211002003332330-1301223301010230-1123100002232003-0133303013301232-2212101303032210)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_discovery/resource.tf`; digest `sha256:ff9af826f29fe467445464e3b92dcd212f9666dfb00feabd76aab3e247a3f9e6`.

```terraform
# APIDiscovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDiscovery configuration
resource "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}
```
