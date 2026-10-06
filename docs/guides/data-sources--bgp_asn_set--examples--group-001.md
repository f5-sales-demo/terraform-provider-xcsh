---
page_title: "xcsh_bgp_asn_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set examples."
---

# xcsh_bgp_asn_set examples

<a id="canonical-0121101003311303-3302213232300110-0000312323132103-3123212202002220-2302122013101120-3233013032232201-3111302230331100-1001302210011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-0101132211310203-2233232230121033-0201030033302303-2112033202323033-1200310311200222-2223131321000021-3211120032230011-1102013013320333)
- Examples

<a id="canonical-3312223030222330-1022223001023200-3100212113131210-1002230230100120-3233233133233111-1202002321232130-0100231020121020-3310312123322233"></a>

### Complete configurations for `xcsh_bgp_asn_set`

- [Data source](data-sources--bgp_asn_set--examples--group-001.md#canonical-0133022332032332-3102110222003332-3302133133030221-1021010001001311-2013002012320103-0132011011131131-3131232011011312-1203110023113022): valid configuration.

<a id="canonical-0133022332032332-3102110222003332-3302133133030221-1021010001001311-2013002012320103-0132011011131131-3131232011011312-1203110023113022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-0101132211310203-2233232230121033-0201030033302303-2112033202323033-1200310311200222-2223131321000021-3211120032230011-1102013013320333)
- [Examples](data-sources--bgp_asn_set--examples--group-001.md#canonical-0121101003311303-3302213232300110-0000312323132103-3123212202002220-2302122013101120-3233013032232201-3111302230331100-1001302210011031)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp_asn_set/data-source.tf`; digest `sha256:101ddc7bff960e99df5f4c481b2d71baf9042bf80e41813366a6137bff577b85`.

```terraform
# BGPAsnSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPAsnSet by name
data "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"
}

output "bgp_asn_set_id" {
  value = data.xcsh_bgp_asn_set.example.id
}
```
