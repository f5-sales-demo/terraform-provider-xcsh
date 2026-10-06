---
page_title: "xcsh_bot_infrastructure"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure."
---

# xcsh_bot_infrastructure

<a id="canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bot_infrastructure

Manages Bot Infrastructure in F5 Distributed Cloud.

<a id="canonical-3320111303211220-3033031231110113-2233130003313120-0031133212200021-0031321033322003-2223230110220301-0312323212032322-0013110330110231"></a>

### Prerequisites for `xcsh_bot_infrastructure`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1310222133233120-0202001312321121-0300333033002132-3123303133003112-2113112232323031-1210323331300011-1203322323310322-1330331320322000"></a>

### Minimal configuration for `xcsh_bot_infrastructure`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotInfrastructure Resource Example
# Manages Bot Infrastructure in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotInfrastructure configuration
resource "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}
```

<a id="canonical-0132130032002310-0310022320210032-2320221213212102-3302122231101130-0300330101220013-2011201020321212-3302222323111003-3223323020210202"></a>

### Root configuration for `xcsh_bot_infrastructure`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1210132031031312-3312030230031322-2022101233123233-0320331303202120-0103210023012322-2310012222101032-3123131220223303-3032211032330313"></a>

### Explore this collection for `xcsh_bot_infrastructure`

- [Property reference](../guides/resources--bot_infrastructure--reference--group-001.md#canonical-3211120122133210-0122110022200232-1211012113212211-2130003230230030-1201022221000022-0012221313230312-0000102002302203-0201002321021101)
- [Examples](../guides/resources--bot_infrastructure--examples--group-001.md#canonical-0312311210210122-3303012200112332-3333133000021021-0301113321021223-3322130112200312-3222123113000213-3101322023012213-3012120222310020)
- [Import](../guides/resources--bot_infrastructure--lifecycle--group-001.md#canonical-0302021010123112-0231123212233312-2101121221122221-2031200220212303-1010313133303101-1123222202230212-1111332032201122-3202323100213112)
- [Timeouts](../guides/resources--bot_infrastructure--lifecycle--group-001.md#canonical-3103223333120321-3132023210300330-0321023230231322-1121122010023310-1230211112310312-3002133211302010-0221202203130132-3210311113201003)
