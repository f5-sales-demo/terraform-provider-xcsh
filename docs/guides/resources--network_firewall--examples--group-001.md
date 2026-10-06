---
page_title: "xcsh_network_firewall examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall examples."
---

# xcsh_network_firewall examples

<a id="canonical-1031032103323200-0122100200112332-1320330020112301-3331011310010120-3320010130103103-2232331310312302-0023100313031112-2201332200222201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- Examples

<a id="canonical-1203232331110131-0112111011232320-2131201301232002-3203011022100030-2103022222113222-0323323312203302-2202120313123312-2310221122112323"></a>

### Complete configurations for `xcsh_network_firewall`

- [Resource](resources--network_firewall--examples--group-001.md#canonical-3131002303311013-0301222221021220-1112323130123222-3021021322310303-3202210313132233-3033332133000102-2203100030222302-0100331332220220): valid configuration.

<a id="canonical-3131002303311013-0301222221021220-1112323130123222-3021021322310303-3202210313132233-3033332133000102-2203100030222302-0100331332220220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md#canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333)
- [Examples](resources--network_firewall--examples--group-001.md#canonical-1031032103323200-0122100200112332-1320330020112301-3331011310010120-3320010130103103-2232331310312302-0023100313031112-2201332200222201)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_firewall/resource.tf`; digest `sha256:7eeae23d1d7b0c0adbae62b0e59aafa91676ffc023ed6aa05396e73c69627d84`.

```terraform
# NetworkFirewall Resource Example
# Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkFirewall configuration
resource "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}
```
