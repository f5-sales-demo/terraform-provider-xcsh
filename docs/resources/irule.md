---
page_title: "xcsh_irule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule landing."
---

# xcsh_irule landing

<a id="canonical-2011333112331031-0020123312211213-0023313313132301-1232121102103020-3220220031303222-0310302031122301-3220110033323311-1131002100310203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320010000030200-3300022102102110-0231220113313223-0033213332030202-1333212312002133-1102312131020021-1302111302003130-2211220100111211"></a>

## xcsh_irule — xcsh_irule / 110212133031 / 2

Breadcrumbs:

- xcsh_irule

Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed
Cloud.

<a id="canonical-2332300221112321-2122320333020203-2311330003031320-2103133313213323-2023021330101012-2002303300123330-2333030213013213-1303002010000020"></a>

## Prerequisites — xcsh_irule / 110212133031 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2320231320310213-1002223020333231-1102333322021130-0113121123023201-3010103121033333-1012110212223300-2210200202030313-3001013330101231"></a>

## Minimal configuration — xcsh_irule / 110212133031 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```

<a id="canonical-0113231323023022-3200121102033220-2120112031032323-0112131030312110-3323032203001110-1131013020223102-2313331232120213-3203022211200020"></a>

## Root configuration — xcsh_irule / 110212133031 / 5

Required root properties: `description_spec`, `irule`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2023033233223302-2220011323233122-0212031323002221-3113203121020003-1132320301120331-0100221313113030-2131330300231002-0302132011022330"></a>

## Next pages — xcsh_irule / 110212133031 / 6

- [Property reference](../guides/resources--irule--reference--group-001.md#canonical-1022201003222120-3100001332021300-1113323102020013-3212101111010201-3331213101333303-0013033103232213-3130321213220101-2011220222210011)
- [Examples](../guides/resources--irule--examples--group-001.md#canonical-0102133322030003-2000111231113322-2322313303303010-3032112202330001-1002023033312232-0231210120003301-2003333032030221-0002311200311231)
- [Import](../guides/resources--irule--lifecycle--group-001.md#canonical-2132111112131131-0130130333322031-1020232023103000-3123320110302133-1230221132223132-2031103010012131-1211210302013202-0210302130121001)
- [Timeouts](../guides/resources--irule--lifecycle--group-001.md#canonical-2033311110233021-0130130223221222-2313021031310232-2303312110331201-3220110113020300-1000310323323302-0020311302130302-0001310123111112)
