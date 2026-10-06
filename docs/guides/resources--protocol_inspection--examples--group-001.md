---
page_title: "xcsh_protocol_inspection examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection examples."
---

# xcsh_protocol_inspection examples

<a id="canonical-3122220013221013-1023121113312331-1202030321022111-3032013321232312-3210321300200303-3121131203021231-1020101200200202-3102211232331033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- Examples

<a id="canonical-0303102010010103-3001310331323321-0123212303101200-1330020311223003-3133103211303031-2332202201311023-2221122313300321-0202000033010232"></a>

### Complete configurations for `xcsh_protocol_inspection`

- [Resource](resources--protocol_inspection--examples--group-001.md#canonical-0131201103312101-2302203022021031-0112010133131322-0222112131332302-2122303022102221-3133020111213221-3000100102231301-3131212011310100): valid configuration.

<a id="canonical-0131201103312101-2302203022021031-0112010133131322-0222112131332302-2122303022102221-3133020111213221-3000100102231301-3131212011310100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102)
- [Examples](resources--protocol_inspection--examples--group-001.md#canonical-3122220013221013-1023121113312331-1202030321022111-3032013321232312-3210321300200303-3121131203021231-1020101200200202-3102211232331033)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protocol_inspection/resource.tf`; digest `sha256:0edd9213402a7bd99b02d8b9c2f3a4ee63eb0305d083b63f69f415300d8acfa9`.

```terraform
# ProtocolInspection Resource Example
# Manages Protocol Inspection Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolInspection configuration
resource "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}
```
