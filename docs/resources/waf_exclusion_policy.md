---
page_title: "xcsh_waf_exclusion_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy."
---

# xcsh_waf_exclusion_policy

<a id="canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_waf_exclusion_policy

Manages WAF exclusion policy in F5 Distributed Cloud.

<a id="canonical-2032003030002133-2302113220103201-0320333301312203-1330032011222012-2100233101211121-3122233320233113-0232303312320110-0202223220021332"></a>

### Prerequisites for `xcsh_waf_exclusion_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3331302031130011-0122311132210133-3022120020023333-2232031011003120-1133001222100230-0232023212023023-1331210112300222-1100301220030133"></a>

### Minimal configuration for `xcsh_waf_exclusion_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFExclusionPolicy Resource Example
# Manages WAF exclusion policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WAFExclusionPolicy configuration
resource "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}
```

<a id="canonical-1310123232113213-3333313221320123-3032330232311310-2203322003122130-2322211122333302-3230212111323233-3000021231301330-2101310211123111"></a>

### Root configuration for `xcsh_waf_exclusion_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3231301133100001-0331132202112320-2032030321133310-1112312113131131-2001133111220332-1311210333011133-0121310200101313-2301122000313031"></a>

### Explore this collection for `xcsh_waf_exclusion_policy`

- [Property reference](../guides/resources--waf_exclusion_policy--reference--group-001.md#canonical-3201012220022020-0302300332132202-1210300000101133-0132320021112000-0303020230312320-2033302130202231-1302212000313220-0123110313003311)
- [Examples](../guides/resources--waf_exclusion_policy--examples--group-001.md#canonical-0110122321102222-3111203320123211-3222003212013303-2333230331313031-2003120230233113-3123231103322202-1020122311021101-3011330131323000)
- [Import](../guides/resources--waf_exclusion_policy--lifecycle--group-001.md#canonical-3012022233000222-3231102202010320-1203311221201203-3303310332312323-2121112333030221-0311201302002230-0031000013221133-3132302210232322)
- [Timeouts](../guides/resources--waf_exclusion_policy--lifecycle--group-001.md#canonical-2020132113200123-3130311000012232-0300331331310301-1332232010012020-3013300331120223-3210132203031030-1121111221312203-2010101011133110)
