---
page_title: "xcsh_alert_template"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template."
---

# xcsh_alert_template

<a id="canonical-2113333112122122-3132222003103103-2101331030303103-1200332320133001-1230313330112233-3120111103131333-0332111212113231-2112301302213113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_alert_template

Manages Domain to protect in F5 Distributed Cloud.

<a id="canonical-3100212213213023-3323201310113123-1311322310310332-3110123300133011-2233132020212231-0000310012130301-2102121301322213-1031213330102202"></a>

### Prerequisites for `xcsh_alert_template`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2210313232200230-1310121233131033-1301023102330122-2111022312323111-1110130320113331-3133031302002212-0210022001333022-0320011233200201"></a>

### Minimal configuration for `xcsh_alert_template`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertTemplate Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertTemplate configuration
resource "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"

  alert_message         = "example-value"
  alert_message_details = "example-value"
  alert_name            = "example-value"
}
```

<a id="canonical-1320022201311120-2103212331302230-1012313100132312-0102323221011102-2000222312031133-3210133000211030-1220210333231303-2212231103112033"></a>

### Root configuration for `xcsh_alert_template`

Required root properties: `alert_message`, `alert_message_details`, `alert_name`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3301123132332131-3130131133000212-1220032313333300-2323213021312003-1001322203123012-3300330222203302-2120310220312111-2022313132131330"></a>

### Explore this collection for `xcsh_alert_template`

- [Property reference](../guides/resources--alert_template--reference--group-001.md#canonical-1133313323321313-1021331100132113-3110202222120000-0021012302223200-2222322030331020-2031011132012122-0322111010203330-2101232233133233)
- [Examples](../guides/resources--alert_template--examples--group-001.md#canonical-1332302023131123-1023300023220300-1231133301313001-2112102020003301-2112101230321132-2020110303233322-1213031323133120-1132031130303123)
- [Import](../guides/resources--alert_template--lifecycle--group-001.md#canonical-3001333320310202-2203222002333303-2303323303201132-2133133110032202-2020303200003301-0133222330010223-0221021330312032-1201103212102003)
- [Timeouts](../guides/resources--alert_template--lifecycle--group-001.md#canonical-0121121000100230-0213020220321103-2223331330323002-2212302133200311-3122012031200320-1130202033001333-3001220023303021-0331310231320300)
