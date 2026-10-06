---
page_title: "xcsh_healthcheck"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck."
---

# xcsh_healthcheck

<a id="canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

<a id="canonical-1310201202033131-1133333300311323-1003232123301102-3012012002302130-0120310130122330-2300023101111113-3110013111131303-0000321033312031"></a>

### Prerequisites for `xcsh_healthcheck`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-0120221012033333-3112021030231232-1220020302100103-0133111213222002-3022120111011331-2210111302101233-3323310123003122-3212300231200100"></a>

### Minimal configuration for `xcsh_healthcheck`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Healthcheck Resource Example
# Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Healthcheck configuration
resource "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"

  healthy_threshold   = 1
  interval            = 1
  timeout             = 1
  unhealthy_threshold = 1
}
```

<a id="canonical-0312011233210200-3200210233222303-3111011203302203-2203013123211300-0110030021033023-2103112233002233-3121033213110112-0030031332123322"></a>

### Root configuration for `xcsh_healthcheck`

Required root properties: `healthy_threshold`, `interval`, `name`, `namespace`, `timeout`, `unhealthy_threshold`. Full root flags and choices appear in the property reference.

<a id="canonical-2103302202233002-1330003330010132-1020013121102113-2012313323020101-2102130003323131-0222330102202312-0131310101000123-0001231013203023"></a>

### Explore this collection for `xcsh_healthcheck`

- [Property reference](../guides/resources--healthcheck--reference--group-001.md#canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213)
- [Examples](../guides/resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- [Import](../guides/resources--healthcheck--lifecycle--group-001.md#canonical-2200022120320020-3330113001002213-0131113001023023-2302333211211011-2120001023103030-1231103003310122-0201010303231030-3100212123301230)
- [Timeouts](../guides/resources--healthcheck--lifecycle--group-001.md#canonical-0013121231322312-2303030031222132-1200202220031030-3202003300322133-1332120130020322-0223120011231312-1113030020333231-2212002323303110)
