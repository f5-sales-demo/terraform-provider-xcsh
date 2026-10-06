---
page_title: "xcsh_virtual_host examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host examples."
---

# xcsh_virtual_host examples

<a id="canonical-1210132032103233-2030210232320032-1132323223302032-0101302113222301-0232320020302021-2032322012000331-2220323303330101-1302002211023223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- Examples

<a id="canonical-0232201301230223-0022123312010202-2113111022131102-3202122100213132-3001213031233333-0311002300012332-3300113130102001-0010233031013122"></a>

### Complete configurations for `xcsh_virtual_host`

- [Resource](resources--virtual_host--examples--group-001.md#canonical-0021210100211230-0220022122312333-1311203322311022-1321013213023310-1022123303321020-3333121131313222-3103010321300233-3323002200013011): valid configuration.

<a id="canonical-0021210100211230-0220022122312333-1311203322311022-1321013213023310-1022123303321020-3333121131313222-3103010321300233-3323002200013011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Examples](resources--virtual_host--examples--group-001.md#canonical-1210132032103233-2030210232320032-1132323223302032-0101302113222301-0232320020302021-2032322012000331-2220323303330101-1302002211023223)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_host/resource.tf`; digest `sha256:7132f3f3fb88713f102679821dabd8f6cf4e11c9765e5f1be76eb1b01ecae4a0`.

```terraform
# VirtualHost Resource Example
# Manages virtual host in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualHost configuration
resource "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}
```
