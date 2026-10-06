---
page_title: "xcsh_endpoint examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint examples."
---

# xcsh_endpoint examples

<a id="canonical-1121101120010011-2333021113122233-0010112232231322-0031100220232023-2001101311210332-1010331101033020-2312023010330321-0223130201200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- Examples

<a id="canonical-1311201320321212-3003121232033223-1322121112231113-0322032323011121-0000320313110031-3332120230202330-3230331102030111-2313032123032333"></a>

### Complete configurations for `xcsh_endpoint`

- [Resource](resources--endpoint--examples--group-001.md#canonical-2212202010123110-2131002232331010-0321202122203111-2021233032121011-1211231100220030-2031111213233103-0132231032232021-2200211010231232): valid configuration.

<a id="canonical-2212202010123110-2131002232331010-0321202122203111-2021233032121011-1211231100220030-2031111213233103-0132231032232021-2200211010231232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012)
- [Examples](resources--endpoint--examples--group-001.md#canonical-1121101120010011-2333021113122233-0010112232231322-0031100220232023-2001101311210332-1010331101033020-2312023010330321-0223130201200122)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_endpoint/resource.tf`; digest `sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e`.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```
