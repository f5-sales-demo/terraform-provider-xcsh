---
page_title: "xcsh_ike_phase2_profile examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile examples."
---

# xcsh_ike_phase2_profile examples

<a id="canonical-3103222302032203-1330320330012010-0213312233122220-1021012121002330-3303103000313230-1233321202101112-0230001312120213-0030313130000110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- Examples

<a id="canonical-3333011332102102-0021323022322233-1121021211211322-2230030023000223-1003002212020200-1200031213321312-2131202113132130-2331232001300202"></a>

### Complete configurations for `xcsh_ike_phase2_profile`

- [Resource](resources--ike_phase2_profile--examples--group-001.md#canonical-1220231231012030-2030203220222323-3133212110332020-0000003000321201-3132330323300003-2332102320201201-2210233330213220-2101030123033212): valid configuration.

<a id="canonical-1220231231012030-2030203220222323-3133212110332020-0000003000321201-3132330323300003-2332102320201201-2210233330213220-2101030123033212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231)
- [Examples](resources--ike_phase2_profile--examples--group-001.md#canonical-3103222302032203-1330320330012010-0213312233122220-1021012121002330-3303103000313230-1233321202101112-0230001312120213-0030313130000110)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike_phase2_profile/resource.tf`; digest `sha256:80e0e1a4beab03b8fb465c2b172d0e27a9d0b25515c600ec2364548afb3d3ffd`.

```terraform
# IKEPhase2Profile Resource Example
# Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase2Profile configuration
resource "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  encryption_algos     = ["example-value"]
}
```
