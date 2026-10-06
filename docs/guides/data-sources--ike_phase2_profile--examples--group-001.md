---
page_title: "xcsh_ike_phase2_profile examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile examples."
---

# xcsh_ike_phase2_profile examples

<a id="canonical-1121112020300021-3233110313120222-0322131211210022-1311222232320300-1301122133123101-2131331322210303-3200100030131231-1301120232011312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- Examples

<a id="canonical-1320011223000313-0320002322030220-1210313232103000-0222020000032022-2023100311112122-3220130112021101-1233303332111221-2322011210032222"></a>

### Complete configurations for `xcsh_ike_phase2_profile`

- [Data source](data-sources--ike_phase2_profile--examples--group-001.md#canonical-1023103312121113-2113111213301030-0033301220223123-2000012223311202-3113120302223233-3221333220320200-1322332312300122-1102303301323031): valid configuration.

<a id="canonical-1023103312121113-2113111213301030-0033301220223123-2000012223311202-3113120302223233-3221333220320200-1322332312300122-1102303301323031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130)
- [Examples](data-sources--ike_phase2_profile--examples--group-001.md#canonical-1121112020300021-3233110313120222-0322131211210022-1311222232320300-1301122133123101-2131331322210303-3200100030131231-1301120232011312)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike_phase2_profile/data-source.tf`; digest `sha256:a5a7c67365793986b4f75d4d921baeb441a3835387aa16c35286802039150551`.

```terraform
# IKEPhase2Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase2Profile by name
data "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"
}

output "ike_phase2_profile_id" {
  value = data.xcsh_ike_phase2_profile.example.id
}
```
