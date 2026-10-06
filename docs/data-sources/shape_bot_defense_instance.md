---
page_title: "xcsh_shape_bot_defense_instance"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_shape_bot_defense_instance."
---

# xcsh_shape_bot_defense_instance

<a id="canonical-3000123112021030-0302103133011033-2210121132011122-2132220211231033-3233312310101331-1222032010313211-1333211301300023-0231320200212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_shape_bot_defense_instance

Reads Shape Bot Defense Instance information from F5 Distributed Cloud.

<a id="canonical-0120130012012012-2232101011020132-2023220312101312-1020223211310020-1112013202323201-1030032322003220-0303321202031132-0110002100221033"></a>

### Prerequisites for `xcsh_shape_bot_defense_instance`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2222212113123311-0110021302320332-1222100331101002-3033233200133210-0222033113312012-1133202222123222-3323323303020233-1122200223231232"></a>

### Minimal configuration for `xcsh_shape_bot_defense_instance`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ShapeBotDefenseInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ShapeBotDefenseInstance by name
data "xcsh_shape_bot_defense_instance" "example" {
  name      = "example-shape-bot-defense-instance"
  namespace = "staging"
}

output "shape_bot_defense_instance_id" {
  value = data.xcsh_shape_bot_defense_instance.example.id
}
```

<a id="canonical-3133023123113030-1333010202201022-3122111111230202-2102021210130323-3231322322310031-2222022211210220-1002111031310200-0002310213102003"></a>

### Root configuration for `xcsh_shape_bot_defense_instance`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1322230211201132-2000022113112132-2200210120222123-3122220201201031-3222201123301021-3112212322230100-3033213310222300-3113032032001311"></a>

### Explore this collection for `xcsh_shape_bot_defense_instance`

- [Property reference](../guides/data-sources--shape_bot_defense_instance--reference--group-001.md#canonical-2202330303322033-2030023020022023-0313320010100303-0032302302312101-0010011310310312-2332312202210132-2003200330000113-3210132332301223)
- [Examples](../guides/data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-1100231133123201-2223300212332101-0310101230300323-0221011123303230-1201321212303220-0102322330223331-1111031331202111-0000011131120220)
