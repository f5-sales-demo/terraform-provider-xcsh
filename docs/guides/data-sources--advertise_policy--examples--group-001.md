---
page_title: "xcsh_advertise_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy examples."
---

# xcsh_advertise_policy examples

<a id="canonical-3001020230333331-2030223330301300-0330111210101131-2021213112120012-0220120323212011-1312001132020230-2100222333330312-3132321031200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- Examples

<a id="canonical-3300002003231003-3211202130112212-0320111022023303-1013313232112231-0323231033203021-1300231003121013-0131331301002131-0223302021131030"></a>

### Complete configurations for `xcsh_advertise_policy`

- [Data source](data-sources--advertise_policy--examples--group-001.md#canonical-1233231313132130-3332030313113030-2321110130210211-3203001212202223-0033001301103321-2201031123102221-0331213233033331-1313122120323121): valid configuration.

<a id="canonical-1233231313132130-3332030313113030-2321110130210211-3203001212202223-0033001301103321-2201031123102221-0331213233033331-1313122120323121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Examples](data-sources--advertise_policy--examples--group-001.md#canonical-3001020230333331-2030223330301300-0330111210101131-2021213112120012-0220120323212011-1312001132020230-2100222333330312-3132321031200111)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_advertise_policy/data-source.tf`; digest `sha256:0faa48749e677b3e8d75bc3281fff3815a14a824fbf83c244cc4c6bf9f75dad8`.

```terraform
# AdvertisePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AdvertisePolicy by name
data "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}

output "advertise_policy_id" {
  value = data.xcsh_advertise_policy.example.id
}
```
