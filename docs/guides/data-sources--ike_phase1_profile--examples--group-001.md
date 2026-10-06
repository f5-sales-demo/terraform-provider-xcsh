---
page_title: "xcsh_ike_phase1_profile examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile examples."
---

# xcsh_ike_phase1_profile examples

<a id="canonical-3131103321231102-0203133013231322-3031033002100130-2320212022331101-0123010211103000-3322032303232021-1321012003023000-2103110133313222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- Examples

<a id="canonical-0220011233033213-1303332210231222-1111222312121222-3131303103303121-3202330332220010-0030020213000013-1133310301311223-1302333001320212"></a>

### Complete configurations for `xcsh_ike_phase1_profile`

- [Data source](data-sources--ike_phase1_profile--examples--group-001.md#canonical-3323210120231233-1120213301210313-1322301121031313-2031222010201011-0012332201233030-1121233113122111-3111220230221103-2211022201310130): valid configuration.

<a id="canonical-3323210120231233-1120213301210313-1322301121031313-2031222010201011-0012332201233030-1121233113122111-3111220230221103-2211022201310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- [Examples](data-sources--ike_phase1_profile--examples--group-001.md#canonical-3131103321231102-0203133013231322-3031033002100130-2320212022331101-0123010211103000-3322032303232021-1321012003023000-2103110133313222)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike_phase1_profile/data-source.tf`; digest `sha256:a22af56dca6a3b92413f9588bd02d5a8a0c4c469599c9023c9feb46580028ab5`.

```terraform
# IKEPhase1Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase1Profile by name
data "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"
}

output "ike_phase1_profile_id" {
  value = data.xcsh_ike_phase1_profile.example.id
}
```
