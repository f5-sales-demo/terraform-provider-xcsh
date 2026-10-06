---
page_title: "xcsh_virtual_network"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network."
---

# xcsh_virtual_network

<a id="canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_virtual_network

Reads Virtual Network information from F5 Distributed Cloud.

<a id="canonical-1322013023321033-3322330303012001-0321213301222101-2023102222121030-0330333120310332-0222023231020001-1030320220123303-1031100323320110"></a>

### Prerequisites for `xcsh_virtual_network`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `network_connector`.

- network_connector: Connect to external networks

<a id="canonical-2210102003131101-3001211222312221-2333233311023002-0211011233112230-0131331113313100-0322302020221333-2312011312031223-3231230033203020"></a>

### Minimal configuration for `xcsh_virtual_network`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualNetwork Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualNetwork by name
data "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}

output "virtual_network_id" {
  value = data.xcsh_virtual_network.example.id
}
```

<a id="canonical-0302131101101011-1213232110323000-3000003203030301-1302301032201230-1320301112332003-3211120032301301-2112031110003012-0130123211323112"></a>

### Root configuration for `xcsh_virtual_network`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1331232103203330-1013100330123301-0110320020221310-2210022102221302-0132322332100331-2103120220220300-1123320032323002-0232303333112021"></a>

### Explore this collection for `xcsh_virtual_network`

- [Property reference](../guides/data-sources--virtual_network--reference--group-001.md#canonical-2130101021323011-3002330133321312-3110103001302103-2133320321012310-3010320330020032-1032031113023113-3330002130320310-0200303030321113)
- [Examples](../guides/data-sources--virtual_network--examples--group-001.md#canonical-2222201120130020-1313132202212131-3101301212113322-0001333200121000-2022312231231133-2330203232100230-0333321232313201-2322110020333313)
