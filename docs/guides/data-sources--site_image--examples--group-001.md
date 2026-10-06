---
page_title: "xcsh_site_image examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_image examples."
---

# xcsh_site_image examples

<a id="canonical-1012123311312022-0323311010031330-1232003023023331-0010122132312101-1231331122323201-0003202220330030-2222032332303213-3312032213033003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_image](../data-sources/site_image.md#canonical-3131222302011312-0030101332231322-0021111121002033-3033310200203330-1301310123023111-0121222100300001-3102203221303023-0331320320202102)
- Examples

<a id="canonical-0032100221311302-3122232011111003-3320013000231012-3123212332130323-1222222212002023-2322232011002331-2010130212310013-2301332222220021"></a>

### Complete configurations for `xcsh_site_image`

- [Data source](data-sources--site_image--examples--group-001.md#canonical-1132221030332110-3321221020101213-1111200220022233-2000303203110102-2022010032000233-3111110123220010-1230231003203232-0130102100010330): valid configuration.

<a id="canonical-1132221030332110-3321221020101213-1111200220022233-2000303203110102-2022010032000233-3111110123220010-1230231003203232-0130102100010330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site_image](../data-sources/site_image.md#canonical-3131222302011312-0030101332231322-0021111121002033-3033310200203330-1301310123023111-0121222100300001-3102203221303023-0331320320202102)
- [Examples](data-sources--site_image--examples--group-001.md#canonical-1012123311312022-0323311010031330-1232003023023331-0010122132312101-1231331122323201-0003202220330030-2222032332303213-3312032213033003)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_image/data-source.tf`; digest `sha256:f749f6969ede8da33793449ff8e0b7f889a9d0fc592de94c4d8d386d5af43af8`.

```terraform
# SiteImage DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "example" {
  site_name = "example-value"
}

output "site_image_result" {
  value     = data.xcsh_site_image.example
  sensitive = true
}
```
