---
page_title: "xcsh_application_profiles examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles examples."
---

# xcsh_application_profiles examples

<a id="canonical-1033011113233320-1301333330113200-2100233022211032-0313022133301123-2122221113330300-0232313020112220-3023003000001232-2122011022223110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- Examples

<a id="canonical-3211322330023313-1010003101101123-1333132003322123-1101320303021332-2120220133200221-0210002132132333-2120333220222332-3033232122031132"></a>

### Complete configurations for `xcsh_application_profiles`

- [Data source](data-sources--application_profiles--examples--group-001.md#canonical-0332101230303031-2111330303202023-1003220102010333-1302010131120302-3320322131011313-3010302133013302-3030302001113320-0213010310210231): valid configuration.

<a id="canonical-0332101230303031-2111330303202023-1003220102010333-1302010131120302-3320322131011313-3010302133013302-3030302001113320-0213010310210231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Examples](data-sources--application_profiles--examples--group-001.md#canonical-1033011113233320-1301333330113200-2100233022211032-0313022133301123-2122221113330300-0232313020112220-3023003000001232-2122011022223110)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_application_profiles/data-source.tf`; digest `sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc`.

```terraform
# ApplicationProfiles Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ApplicationProfiles by name
data "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}

output "application_profiles_id" {
  value = data.xcsh_application_profiles.example.id
}
```
