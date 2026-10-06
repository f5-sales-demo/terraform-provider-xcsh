---
page_title: "xcsh_alert_gen_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy."
---

# xcsh_alert_gen_policy

<a id="canonical-2201222232120330-1303130313333202-2020210123203131-3130311000211300-3201232201020110-2210203113113010-2232121111101221-3030222001111001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_alert_gen_policy

Manages Alert Generation Policy in F5 Distributed Cloud.

<a id="canonical-0321133211320201-1013333010200110-3313123303133033-3231312013003320-2220011131333313-0223301121100121-2102311111303231-0203310332122032"></a>

### Prerequisites for `xcsh_alert_gen_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3213000232212203-2033121300022301-1111322202313320-1333301000013022-3221223322223232-3130001101011012-3323323101022220-1030221102212322"></a>

### Minimal configuration for `xcsh_alert_gen_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertGenPolicy Resource Example
# Manages Alert Generation Policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertGenPolicy configuration
resource "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}
```

<a id="canonical-0030202313111330-1110120301231000-0022331120002313-3022000111131102-1201201100310122-0003202033022022-2211310113112302-0322332233313230"></a>

### Root configuration for `xcsh_alert_gen_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0332130133330132-1001303302301013-1111032222321223-1300313321332131-1322003302300003-3001221032130030-1132201310310200-1031100311230032"></a>

### Explore this collection for `xcsh_alert_gen_policy`

- [Property reference](../guides/resources--alert_gen_policy--reference--group-001.md#canonical-0113323010121111-3310031311010033-0120221122230330-1032002302130100-3011003320313103-0332032231313220-3233111110000200-2012111221210303)
- [Examples](../guides/resources--alert_gen_policy--examples--group-001.md#canonical-1333311220130202-1032210111110011-0220332203212331-1012033311301010-2031321003013300-0312030101013302-0002322233120201-1110333102211333)
- [Import](../guides/resources--alert_gen_policy--lifecycle--group-001.md#canonical-1103101021010221-2231010233003223-1320220111303111-0312011221230000-2132221132030030-1203131200111010-1303223102220113-0221200020321012)
- [Timeouts](../guides/resources--alert_gen_policy--lifecycle--group-001.md#canonical-2203233101331112-3132232220332331-3132000013020003-3311232002232333-2133322212003010-3230223012222220-2112113223211330-3330202203113113)
