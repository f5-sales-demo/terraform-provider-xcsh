---
page_title: "xcsh_token examples"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token examples."
---

# xcsh_token examples

<a id="canonical-1231001331010212-1313303001022020-1211202200302331-3202013321101030-2030333303033133-2200310331003311-3031023232120123-0203113200222102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020)
- Examples

<a id="canonical-2213222333200020-2012211301322011-0203202231203022-1112023120200213-0222302213213320-1133213110203122-0222123002111213-2312233311032130"></a>

### Complete configurations for `xcsh_token`

- [Resource](resources--token--examples--group-001.md#canonical-2101330310331202-0332121202233333-3312120320221103-2222113332330330-0203101000211121-0112012123211223-2321302231333312-3120102111310332): valid configuration.

<a id="canonical-2101330310331202-0332121202233333-3312120320221103-2222113332330330-0203101000211121-0112012123211223-2321302231333312-3120102111310332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020)
- [Examples](resources--token--examples--group-001.md#canonical-1231001331010212-1313303001022020-1211202200302331-3202013321101030-2030333303033133-2200310331003311-3031023232120123-0203113200222102)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_token/resource.tf`; digest `sha256:e34df2fe171a3a579b8dd3181a5ec00f7f49f2449740ae1edcf8cdaec5c57bb9`.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```
