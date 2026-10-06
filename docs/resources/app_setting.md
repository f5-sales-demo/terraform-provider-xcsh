---
page_title: "xcsh_app_setting"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting."
---

# xcsh_app_setting

<a id="canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_app_setting

Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

<a id="canonical-3002313012022020-3212323203232233-3220330023230123-2001103213301211-3132230321331211-1010301203003223-3211020310221033-1132330320332233"></a>

### Prerequisites for `xcsh_app_setting`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3301303222122113-0220112113020010-2223300101031101-1221303331213101-1110033111330301-3301211133202120-1200121020010020-3121110230012200"></a>

### Minimal configuration for `xcsh_app_setting`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSetting Resource Example
# Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppSetting configuration
resource "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}
```

<a id="canonical-2223202221123113-2020111120130333-2330202323030032-3210000322001330-0022231322333331-1023203002003023-0310131012323030-2120000221302300"></a>

### Root configuration for `xcsh_app_setting`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1131112303022121-0203111310130323-3121013210221213-0122223222103120-3021301123211232-2233332223200312-2200001311103311-1220020201223132"></a>

### Explore this collection for `xcsh_app_setting`

- [Property reference](../guides/resources--app_setting--reference--group-001.md#canonical-2122310201222303-0130033213230301-0221313313103320-2312231332122202-2212322133022101-3100330100032011-2330302131301110-0022232333200010)
- [Examples](../guides/resources--app_setting--examples--group-001.md#canonical-3320122113230000-3221120120232111-0331321100100320-1223112322232312-0212003023000333-3013212020121300-1022310012331022-1323330321021011)
- [Import](../guides/resources--app_setting--lifecycle--group-001.md#canonical-0020100003210303-3312011210033212-3313003111212202-2010110121101100-0032231120133233-0030303031013133-3011203100112232-2003002330113223)
- [Timeouts](../guides/resources--app_setting--lifecycle--group-001.md#canonical-3320223010201200-0221132032121332-2120300123022030-1232002232301201-3303200231131300-1300031003122321-0210313211200023-3003133112012003)
