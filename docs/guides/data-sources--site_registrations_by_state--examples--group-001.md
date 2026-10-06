---
page_title: "xcsh_site_registrations_by_state examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_state examples."
---

# xcsh_site_registrations_by_state examples

<a id="canonical-3303201010310230-2003223220001303-1101321302202101-2030220322223133-1333332012121232-2102231003300003-1210020331110123-0320113132311303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- Examples

<a id="canonical-1313100301111211-0233132113310233-3111302103121301-3223001000311331-0311313332012330-3110123232210121-2002322111011101-1213103010130233"></a>

### Complete configurations for `xcsh_site_registrations_by_state`

- [Data source](data-sources--site_registrations_by_state--examples--group-001.md#canonical-2122023233332330-2113130130003332-1301333213013322-3220013132032113-0310120123300322-0211102000110123-3123233230233132-1121301001023212): valid configuration.

<a id="canonical-2122023233332330-2113130130003332-1301333213013322-3220013132032113-0310120123300322-0211102000110123-3123233230233132-1121301001023212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303)
- [Examples](data-sources--site_registrations_by_state--examples--group-001.md#canonical-3303201010310230-2003223220001303-1101321302202101-2030220322223133-1333332012121232-2102231003300003-1210020331110123-0320113132311303)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations_by_state/data-source.tf`; digest `sha256:87878119dc98d70aee1fa7c3e546432844de313704e34499df4b6dfdc776ce5a`.

```terraform
# SiteRegistrationsByState DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_state" "example" {
  state = "NOTSET"
}

output "site_registrations_by_state_result" {
  value = data.xcsh_site_registrations_by_state.example
}
```
