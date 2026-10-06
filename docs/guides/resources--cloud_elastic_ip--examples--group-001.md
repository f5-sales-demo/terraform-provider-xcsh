---
page_title: "xcsh_cloud_elastic_ip examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip examples."
---

# xcsh_cloud_elastic_ip examples

<a id="canonical-0131101220123332-1213332212210201-1213012320222330-1201113223023221-2313030213232111-1003202233132311-1231101021130003-2221323011023332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)
- Examples

<a id="canonical-3210002211031111-2213020201120230-3312012000023002-3032131031310033-3101032121012020-3330100330013123-2222210112321103-3333020311232132"></a>

### Complete configurations for `xcsh_cloud_elastic_ip`

- [Resource](resources--cloud_elastic_ip--examples--group-001.md#canonical-2303013103321313-0312200010132320-2031231330230232-2013133233303102-0010021011001023-2022030310033122-2213032031103011-1301301200210022): valid configuration.

<a id="canonical-2303013103321313-0312200010132320-2031231330230232-2013133233303102-0010021011001023-2022030310033122-2213032031103011-1301301200210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)
- [Examples](resources--cloud_elastic_ip--examples--group-001.md#canonical-0131101220123332-1213332212210201-1213012320222330-1201113223023221-2313030213232111-1003202233132311-1231101021130003-2221323011023332)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_elastic_ip/resource.tf`; digest `sha256:e8be04df02915d54850afcd099a22da22eee85dc886b468e9bf65ca321094b18`.

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
