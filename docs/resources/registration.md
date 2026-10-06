---
page_title: "xcsh_registration"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration."
---

# xcsh_registration

<a id="canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_registration

Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this
message, never used by users. configuration.

<a id="canonical-2121012300032101-0123100221332033-2031033132300021-1300002011303033-1103310332120221-2033203133310212-1020102330112001-0130331213100122"></a>

### Prerequisites for `xcsh_registration`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1333131320120022-1232132220020010-0020223310010320-3303022113233001-0223313102213321-2112011112023123-3103002302202333-2212102112123211"></a>

### Minimal configuration for `xcsh_registration`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
```

<a id="canonical-0131111323030122-1312321031013113-1112030211130133-1133221211222103-0101322301213010-3111123002111130-0331213020301221-1102023230332200"></a>

### Root configuration for `xcsh_registration`

Required root properties: `name`, `namespace`, `token`. Full root flags and choices appear in the property reference.

<a id="canonical-3201130333131210-0121001000320303-2200200202202202-1013020120211323-3021233022332023-3113333021200010-1330333023211030-0122323320210020"></a>

### Explore this collection for `xcsh_registration`

- [Property reference](../guides/resources--registration--reference--group-001.md#canonical-0023010312231331-2303222311312233-3333323230101330-0112112211222220-1231111332020310-2321221110010130-2133212213030231-1131332203133003)
- [Examples](../guides/resources--registration--examples--group-001.md#canonical-1011213302031020-3021233102330101-3101202322303030-0012313312202121-1232211201031213-0231311233333121-1000210233301011-3232132023213203)
- [Import](../guides/resources--registration--lifecycle--group-001.md#canonical-0003213230202100-1010100003100000-3213030201322202-1222011212300110-1030000212113132-2232120010122000-2210001300013102-0211313022120221)
- [Timeouts](../guides/resources--registration--lifecycle--group-001.md#canonical-3032332112220010-3223313323330323-3012001001103003-3002300023122021-1132212113120220-2113011212211230-3212202011120002-0202220201321133)
