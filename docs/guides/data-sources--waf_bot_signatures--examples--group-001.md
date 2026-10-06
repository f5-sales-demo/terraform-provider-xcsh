---
page_title: "xcsh_waf_bot_signatures examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_bot_signatures examples."
---

# xcsh_waf_bot_signatures examples

<a id="canonical-0221322122210100-2300022313321102-2033010210311101-2201221100231320-3020012201200310-1222222312200201-3033013013221222-3211032233120313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md#canonical-1232013323203331-2030333000300300-3031302123320310-3302000013332010-1200120322302012-2021223130131323-2112213021213111-3103303123120210)
- Examples

<a id="canonical-3003320313132330-0111030111032012-2111203222133022-3020313010130330-1101032312132211-3112003210201030-1222221320112231-1032023311111231"></a>

### Complete configurations for `xcsh_waf_bot_signatures`

- [Data source](data-sources--waf_bot_signatures--examples--group-001.md#canonical-1203312111202322-1332133333021203-1220133213132101-3020210131322101-2112213132033113-3211312311133100-2013320301000210-0222020121330322): valid configuration.

<a id="canonical-1203312111202322-1332133333021203-1220133213132101-3020210131322101-2112213132033113-3211312311133100-2013320301000210-0222020121330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md#canonical-1232013323203331-2030333000300300-3031302123320310-3302000013332010-1200120322302012-2021223130131323-2112213021213111-3103303123120210)
- [Examples](data-sources--waf_bot_signatures--examples--group-001.md#canonical-0221322122210100-2300022313321102-2033010210311101-2201221100231320-3020012201200310-1222222312200201-3033013013221222-3211032233120313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_bot_signatures/data-source.tf`; digest `sha256:0cea795e7e9cddaecc9f49e84c9c13efa8ac69b7e21b4ee8402220f9f8ec4b3f`.

```terraform
# WAFBotSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_bot_signatures" "example" {
}

output "waf_bot_signatures_result" {
  value = data.xcsh_waf_bot_signatures.example
}
```
