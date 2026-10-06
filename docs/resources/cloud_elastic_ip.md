---
page_title: "xcsh_cloud_elastic_ip"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip."
---

# xcsh_cloud_elastic_ip

<a id="canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_cloud_elastic_ip

Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5
Distributed Cloud.

<a id="canonical-3222011013031322-3111321033032113-2023233221213002-3110023120300233-1221112012220202-1200323112332003-0303022120022011-3320312202201301"></a>

### Prerequisites for `xcsh_cloud_elastic_ip`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2031200032231101-3331222223020333-0132310210232112-3320112301331330-3330020102300012-0321200121331230-2200302033002011-0301312330232221"></a>

### Minimal configuration for `xcsh_cloud_elastic_ip`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudElasticIP Resource Example
# Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudElasticIP configuration
resource "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"

  item_count = 1
}
```

<a id="canonical-3100213133133022-2110110103103232-3333331231101303-1331010310111233-1101321202300230-2301020220202010-2202203112332312-1123331133033321"></a>

### Root configuration for `xcsh_cloud_elastic_ip`

Required root properties: `item_count`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2101123130001331-0301331221011031-1322013012110113-0300330332210223-2123301201022120-2101111331110301-2023203303211303-0130331030133211"></a>

### Explore this collection for `xcsh_cloud_elastic_ip`

- [Property reference](../guides/resources--cloud_elastic_ip--reference--group-001.md#canonical-3000332101332101-3033231112032333-3203132203302322-1221122003220232-1112221212310323-1330110330201000-0101323211230023-3313311003132210)
- [Examples](../guides/resources--cloud_elastic_ip--examples--group-001.md#canonical-0131101220123332-1213332212210201-1213012320222330-1201113223023221-2313030213232111-1003202233132311-1231101021130003-2221323011023332)
- [Import](../guides/resources--cloud_elastic_ip--lifecycle--group-001.md#canonical-1100222313213321-3102211332332302-0121312102002123-2212121031030030-2231021010322103-0032031232020320-3310221223102311-2222322011311333)
- [Timeouts](../guides/resources--cloud_elastic_ip--lifecycle--group-001.md#canonical-2313202201103131-1101223201312300-2201132311211020-1313201233302113-2113323233223000-2013223000211120-3103310331200201-0333030201213301)
