---
page_title: "xcsh_certificate"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate."
---

# xcsh_certificate

<a id="canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_certificate

Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.

<a id="canonical-2300333223210103-2323103003100203-2032133110102000-0213202322022312-0230302301121103-3311022313102021-1322001213203330-0312312030111000"></a>

### Prerequisites for `xcsh_certificate`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-2110001212011022-3303000220132230-1303131311133111-3331313330033331-2033112030322320-2322120032311001-1220133322213222-2033202121101110"></a>

### Minimal configuration for `xcsh_certificate`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```

<a id="canonical-1210000121031313-2301001121230311-0312320310222012-2113011220020313-3231212120132331-2200322132132321-1131212101000101-1021330210121132"></a>

### Root configuration for `xcsh_certificate`

Required root properties: `certificate_url`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1221111130031212-1033233233113330-0320320323301133-2221120033223023-1033011020112233-1101211211100020-1003322222211111-1230300022221013"></a>

### Explore this collection for `xcsh_certificate`

- [Property reference](../guides/resources--certificate--reference--group-001.md#canonical-1232012101321203-0232010211323212-3033033201203103-3213030001121213-1213230323110200-2223013311301231-0110133211131133-3323323132201330)
- [Examples](../guides/resources--certificate--examples--group-001.md#canonical-3322313303332131-0131020110320022-1221212301120013-2231031201130012-3302302113112112-3032012230000112-3133333123033310-0120311001333332)
- [Import](../guides/resources--certificate--lifecycle--group-001.md#canonical-1101233031133031-3130000233121232-3102010323131202-3101303002132031-0012011222011220-2110202303020302-1303110123110020-3321102223002210)
- [Timeouts](../guides/resources--certificate--lifecycle--group-001.md#canonical-0021002331222221-2202012001000222-0130003010031112-3300032011101212-2300113311022110-0103203002230210-2110030210003323-2031130021232311)
