---
page_title: "xcsh_service_policy"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy."
---

# xcsh_service_policy

<a id="canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_service_policy

Manages service\_policy creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-3100022012222301-0202331022231111-0223122212030130-3001221021310002-1230031221131332-3012031100020030-2203312021111200-2200112032322030"></a>

### Prerequisites for `xcsh_service_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-1200323013023133-0200112012122112-2203110330201113-1222101330021100-2311303310330121-2313101022020101-2100232201222131-1233331303201022"></a>

### Minimal configuration for `xcsh_service_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicy Resource Example
# Manages service_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicy configuration
resource "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}
```

<a id="canonical-1313110032110233-2102320321020311-1011202212333201-1201331012023022-0000322121320233-1201223310131201-1231023023130023-0332122011121032"></a>

### Root configuration for `xcsh_service_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2031323010331003-1020122222213303-3313133110011302-1303120121323211-1132213033331232-0330320311322103-1332331212321321-2012320110000312"></a>

### Explore this collection for `xcsh_service_policy`

- [Property reference](../guides/resources--service_policy--reference--group-001.md#canonical-3233121223111310-1123021201111211-2020111333013010-3123332230330123-3220120121331001-3202000010100310-1011313102310030-1312202213103330)
- [Examples](../guides/resources--service_policy--examples--group-001.md#canonical-1220321203332321-0223210020200002-3001003033013233-2101111202111230-3201121300222213-2000333221301112-3223311102130320-3311232312310321)
- [Import](../guides/resources--service_policy--lifecycle--group-001.md#canonical-3001210010320130-2032113022230020-3233101011020330-3120033130332021-2132221202012233-0010012202130213-0110100100330330-1301210030303010)
- [Timeouts](../guides/resources--service_policy--lifecycle--group-001.md#canonical-1020100312322001-0022102320010322-0313030322230202-1311032033323020-0020123001102001-2112020233021011-2222201222200103-3312223323002112)
