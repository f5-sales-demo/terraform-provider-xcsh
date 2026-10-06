---
page_title: "xcsh_network_interface"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface."
---

# xcsh_network_interface

<a id="canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_interface

Reads network device interface configuration in the system namespace.

<a id="canonical-2021211123103203-0201223321222303-2332232022101223-1013030100032301-2211121010333233-1210230030302113-3123031001023112-1112002331033210"></a>

### Prerequisites for `xcsh_network_interface`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1002310032300210-2333110333023232-2012332210001020-0310323121011213-3101333301000122-1112323300032132-0313132310012012-0012133323023332"></a>

### Minimal configuration for `xcsh_network_interface`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkInterface Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkInterface by name
data "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}

output "network_interface_id" {
  value = data.xcsh_network_interface.example.id
}
```

<a id="canonical-1011210301133323-1320123200002220-2302301332001320-0101203301300222-3100203332121100-3133231233001112-0002011200303030-1003222323322003"></a>

### Root configuration for `xcsh_network_interface`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2301220122330301-2022213011021302-3320220103133003-3300010320220003-3022021302122033-3332003201220101-1132323300112302-0001210323222230"></a>

### Explore this collection for `xcsh_network_interface`

- [Property reference](../guides/data-sources--network_interface--reference--group-001.md#canonical-1302132230310323-3010120302021230-0003303202100320-0132133203003123-0323321101130000-0302213333112232-2221001130100010-1233331100132210)
- [Examples](../guides/data-sources--network_interface--examples--group-001.md#canonical-2133221333220133-3003232131233011-0323131130131010-0032031132202203-3203132322000010-3200121132023132-3333031110011010-0230013230113030)
