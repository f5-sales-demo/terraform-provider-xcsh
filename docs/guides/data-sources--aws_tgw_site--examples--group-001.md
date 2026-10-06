---
page_title: "xcsh_aws_tgw_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site examples."
---

# xcsh_aws_tgw_site examples

<a id="canonical-3133121303000003-2123223103111022-1103223113300032-0031121220130230-2202132310101300-1032021013320110-0330233311313032-0312022011210233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- Examples

<a id="canonical-0130013233033133-3331323111131221-3000000002132211-2112120213310123-2301211031322333-0333332213321220-3102320100010010-3222200023220013"></a>

### Complete configurations for `xcsh_aws_tgw_site`

- [Data source](data-sources--aws_tgw_site--examples--group-001.md#canonical-2010230202333132-1032132311330303-0132211003312131-1312212020202102-2123320123201032-2003002020012310-0311221100231331-0323201131232302): valid configuration.

<a id="canonical-2010230202333132-1032132311330303-0132211003312131-1312212020202102-2123320123201032-2003002020012310-0311221100231331-0323201131232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Examples](data-sources--aws_tgw_site--examples--group-001.md#canonical-3133121303000003-2123223103111022-1103223113300032-0031121220130230-2202132310101300-1032021013320110-0330233311313032-0312022011210233)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_aws_tgw_site/data-source.tf`; digest `sha256:aef0f2c4d7539fcdc86d3bd364e4b9a22f4422d4d0595364ea8d26b8fc2a6cf7`.

```terraform
# AWSTGWSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSTGWSite by name
data "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}

output "aws_tgw_site_id" {
  value = data.xcsh_aws_tgw_site.example.id
}
```
