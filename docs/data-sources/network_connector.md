---
page_title: "xcsh_network_connector"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector."
---

# xcsh_network_connector

<a id="canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_connector

Reads a Network Connector created in the system namespace.

<a id="canonical-1131312001220100-2223030031200130-0333203330122012-0033312322100131-2112032102313011-0011303210221030-2121020200003301-2002021231010331"></a>

### Prerequisites for `xcsh_network_connector`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_network`.

- virtual_network: Network to connect

<a id="canonical-2203023111012333-0121220122021001-0323330001100131-1033021113111223-2300331323000032-1011321112222232-2221020332013333-0032131203301323"></a>

### Minimal configuration for `xcsh_network_connector`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkConnector by name
data "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}

output "network_connector_id" {
  value = data.xcsh_network_connector.example.id
}
```

<a id="canonical-1231010230000023-2211132030203303-2100323202320003-1013121221331012-1111002201210313-3303023012123033-0300333320202101-0002231310213020"></a>

### Root configuration for `xcsh_network_connector`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2233020110231222-1122122133003232-1133110111303301-2001213122102133-1321021033321011-3321301212122212-3000303200332002-0120332303020013"></a>

### Explore this collection for `xcsh_network_connector`

- [Property reference](../guides/data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [Examples](../guides/data-sources--network_connector--examples--group-001.md#canonical-3111123311200233-3221113313323111-1030301032222222-2233132330113212-2012130321031310-0022330033310111-0021302301033202-2220033120232122)
