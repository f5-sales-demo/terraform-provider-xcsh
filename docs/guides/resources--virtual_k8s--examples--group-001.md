---
page_title: "xcsh_virtual_k8s examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s examples."
---

# xcsh_virtual_k8s examples

<a id="canonical-2321313332101022-1203232122231303-0133012333202120-0131022023203202-1233113033322022-0320211011333031-0111110210103001-1010200302220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- Examples

<a id="canonical-0330313021231321-0100333212200202-1322123012123310-1333022132110311-2200123331012203-0103131212003313-3213013122011112-3200312131332323"></a>

### Complete configurations for `xcsh_virtual_k8s`

- [Resource](resources--virtual_k8s--examples--group-001.md#canonical-3221033020111102-1333301300323130-0233033032222101-0230102233210103-2221333132030102-0203011022023031-0133112130001103-1023002111311130): valid configuration.

<a id="canonical-3221033020111102-1333301300323130-0233033032222101-0230102233210103-2221333132030102-0203011022023031-0133112130001103-1023002111311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- [Examples](resources--virtual_k8s--examples--group-001.md#canonical-2321313332101022-1203232122231303-0133012333202120-0131022023203202-1233113033322022-0320211011333031-0111110210103001-1010200302220301)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_k8s/resource.tf`; digest `sha256:a5065caf12c5cd73a6083207375f30b6129bce86fd82163c8845c9e0ab44c400`.

```terraform
# VirtualK8S Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualK8S configuration
resource "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}
```
