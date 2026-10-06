---
page_title: "xcsh_route examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route examples."
---

# xcsh_route examples

<a id="canonical-1222313231020302-3320201100121100-1033003333200232-3230303210001313-0012333022000220-2202211312132001-3130100332121202-2002322131033003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- Examples

<a id="canonical-0332203323123232-2011130131002031-3312133310022323-2230322222320013-2000211012020111-1223033103102032-3332321013301211-2033011230310103"></a>

### Complete configurations for `xcsh_route`

- [Resource](resources--route--examples--group-001.md#canonical-1133233331113213-3131330001133202-1322003033201213-3003331311023122-3031101022221102-1021211012333310-3030033111030203-2220313212013230): valid configuration.

<a id="canonical-1133233331113213-3131330001133202-1322003033201213-3003331311023122-3031101022221102-1021211012333310-3030033111030203-2220313212013230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Examples](resources--route--examples--group-001.md#canonical-1222313231020302-3320201100121100-1033003333200232-3230303210001313-0012333022000220-2202211312132001-3130100332121202-2002322131033003)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_route/resource.tf`; digest `sha256:02c199de787def8c45388d6ef15d06881bc08c78cf9b74e0ee43c7ac9db177f3`.

```terraform
# Route Resource Example
# Manages route object in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Route configuration
resource "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}
```
