---
page_title: "xcsh_forwarding_class"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class."
---

# xcsh_forwarding_class

<a id="canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_forwarding_class

Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users
in system namespace. configuration.

<a id="canonical-3013123331103333-1231201230223330-0321113230130220-1323311021011303-0303211131330320-2211212233003100-0300031101022233-2002130220122102"></a>

### Prerequisites for `xcsh_forwarding_class`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0211012220210220-2201210122201010-2123232210120201-0333223032130201-3021220123312001-1032210130032201-2131022322300001-1023113023120301"></a>

### Minimal configuration for `xcsh_forwarding_class`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardingClass Resource Example
# Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardingClass configuration
resource "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}
```

<a id="canonical-1220020231322321-1221102123023130-2300102111100012-1110133031033231-2332112323030201-1303213102302120-3223013311112033-2333120120321323"></a>

### Root configuration for `xcsh_forwarding_class`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0203303123012011-2122003120330102-1310022023000302-2002322110111211-2020203010212230-1312233201100001-3232110200010323-1010320323113003"></a>

### Explore this collection for `xcsh_forwarding_class`

- [Property reference](../guides/resources--forwarding_class--reference--group-001.md#canonical-0102111112020232-1131131013230311-1223320113131321-1013330221212303-3022123312213323-3323312223011303-2202303002230110-1331212022130030)
- [Examples](../guides/resources--forwarding_class--examples--group-001.md#canonical-1121111133103312-1312010112003212-2002312232032321-0022321033300030-1300301001320300-2220313031122222-2232333122231221-0111320013023222)
- [Import](../guides/resources--forwarding_class--lifecycle--group-001.md#canonical-1211101330330220-2322333121023332-2313121232330110-3122301132213330-3200103033112131-2321133230121200-1133230222020303-0100110310320131)
- [Timeouts](../guides/resources--forwarding_class--lifecycle--group-001.md#canonical-3311003020230302-2133230220213031-1223113232012000-1102320013201012-1313011131333203-0022010333301233-3112131212121010-0032011233222323)
