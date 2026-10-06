---
page_title: "xcsh_nfv_service examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service examples."
---

# xcsh_nfv_service examples

<a id="canonical-3323120001301022-1300222011110323-3230233332130232-2013313320213001-1331011000332003-3011300111102130-2323011313120301-3001310223031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- Examples

<a id="canonical-1113000113310321-0300220030303230-0203231330303331-2311101021130332-2223313332321133-1000330020210111-0121002002021311-1101023302020122"></a>

### Complete configurations for `xcsh_nfv_service`

- [Resource](resources--nfv_service--examples--group-001.md#canonical-1213221123222201-0030223100121310-2210301222223030-2232002320212033-0031131003022211-2120300230211333-1333011030220133-3133220011303213): valid configuration.

<a id="canonical-1213221123222201-0030223100121310-2210301222223030-2232002320212033-0031131003022211-2120300230211333-1333011030220133-3133220011303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Examples](resources--nfv_service--examples--group-001.md#canonical-3323120001301022-1300222011110323-3230233332130232-2013313320213001-1331011000332003-3011300111102130-2323011313120301-3001310223031303)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nfv_service/resource.tf`; digest `sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9`.

```terraform
# NfvService Resource Example
# Manages new NFV service with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NfvService configuration
resource "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}
```
