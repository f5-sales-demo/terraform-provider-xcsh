---
page_title: "xcsh_ip_prefix_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set examples."
---

# xcsh_ip_prefix_set examples

<a id="canonical-1212030001113112-0332032221030101-1222301101333121-2213321023322222-1230232013222230-0120233121333323-3130000133022003-0021221230111132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)
- Examples

<a id="canonical-0101320011230233-2323120220132013-1102320031310303-1101212033323021-2212031012033102-2211333010201023-3330020222011132-2030123310102012"></a>

### Complete configurations for `xcsh_ip_prefix_set`

- [Resource](resources--ip_prefix_set--examples--group-001.md#canonical-0333101020302003-3203020032220321-3020100100121231-2031131302133301-0230031323032111-1222101000220332-1103010120331311-3130013313131323): valid configuration.

<a id="canonical-0333101020302003-3203020032220321-3020100100121231-2031131302133301-0230031323032111-1222101000220332-1103010120331311-3130013313131323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232)
- [Examples](resources--ip_prefix_set--examples--group-001.md#canonical-1212030001113112-0332032221030101-1222301101333121-2213321023322222-1230232013222230-0120233121333323-3130000133022003-0021221230111132)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ip_prefix_set/resource.tf`; digest `sha256:00f8d9b11740968b5c7c6ae54585108409034d987dc62393dd9ee54d031d3c3a`.

```terraform
# IPPrefixSet Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IPPrefixSet configuration
resource "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}
```
