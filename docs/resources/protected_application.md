---
page_title: "xcsh_protected_application"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application."
---

# xcsh_protected_application

<a id="canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_protected_application

Manages applications protected by Bot Defense in F5 Distributed Cloud.

<a id="canonical-2213321303013320-1212112001103103-2231212110213102-3301322012313300-3021031032032331-1233330211233000-3001210320031213-2021131301102123"></a>

### Prerequisites for `xcsh_protected_application`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3032110132130212-0311323122203122-0200220203123301-2112021121301101-3012031033010333-1020202202223202-1203222200002233-3110111110030212"></a>

### Minimal configuration for `xcsh_protected_application`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedApplication Resource Example
# Manages applications protected by Bot Defense in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedApplication configuration
resource "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}
```

<a id="canonical-1201030033033201-3030332013302033-3113210300200310-1030013332001223-3331110322220002-2302222223133230-0201132000031203-2023112000111032"></a>

### Root configuration for `xcsh_protected_application`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3012210120022120-1320131101013112-0303032021032110-1201132333022121-3023313230323011-2202111230200220-1011213323103301-1212120313132113"></a>

### Explore this collection for `xcsh_protected_application`

- [Property reference](../guides/resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Examples](../guides/resources--protected_application--examples--group-001.md#canonical-0110110122211031-1102303221021232-0000022301112032-0310000330222201-3210220230020020-2312030223300300-2231200033133123-0312331122232112)
- [Import](../guides/resources--protected_application--lifecycle--group-001.md#canonical-0112302200331003-2330231312231032-1202220233011001-2023010321003301-0312222323010120-2000220100001223-0223021100133221-2331002221210112)
- [Timeouts](../guides/resources--protected_application--lifecycle--group-001.md#canonical-2200213232111023-0231000313103130-1301320103311030-2213110320101110-2300002211003322-3130013301323220-3123031221222331-2332330312012311)
