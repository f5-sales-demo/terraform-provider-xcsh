---
page_title: "xcsh_waf_latest_signatures_version examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_latest_signatures_version examples."
---

# xcsh_waf_latest_signatures_version examples

<a id="canonical-2131332222021330-0302033333003232-1303001311302012-1102301302223003-3111333301021202-1020200000001103-2012120003202223-1032001220123110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_waf_latest_signatures_version](../data-sources/waf_latest_signatures_version.md#canonical-3220021120200003-2131100022002331-1323212033111120-3012322121033220-3000100023133103-1211320023113021-0322300310221012-2100112203232030)
- Examples

<a id="canonical-2310112200331022-3313021023330120-0231212100211102-2231220213111301-0301331231221103-0222330123032323-0010222103021212-2012301231132011"></a>

### Complete configurations for `xcsh_waf_latest_signatures_version`

- [Data source](data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-3322302323021011-1001303221012110-1002132321030331-0110222012202212-3312233122102101-2222201132221212-1031100103330133-0330300331023101): valid configuration.

<a id="canonical-3322302323021011-1001303221012110-1002132321030331-0110222012202212-3312233122102101-2222201132221212-1031100103330133-0330300331023101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_waf_latest_signatures_version](../data-sources/waf_latest_signatures_version.md#canonical-3220021120200003-2131100022002331-1323212033111120-3012322121033220-3000100023133103-1211320023113021-0322300310221012-2100112203232030)
- [Examples](data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-2131332222021330-0302033333003232-1303001311302012-1102301302223003-3111333301021202-1020200000001103-2012120003202223-1032001220123110)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_latest_signatures_version/data-source.tf`; digest `sha256:6a87350aef0ac879fd92c4cae37ff6595f45e1a414503ebb0313bae760a73e29`.

```terraform
# WAFLatestSignaturesVersion DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_latest_signatures_version" "example" {
}

output "waf_latest_signatures_version_result" {
  value = data.xcsh_waf_latest_signatures_version.example
}
```
