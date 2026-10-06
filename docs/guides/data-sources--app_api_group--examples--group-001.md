---
page_title: "xcsh_app_api_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group examples."
---

# xcsh_app_api_group examples

<a id="canonical-0302100133100312-2121120001033302-0022331022132312-3310031221011031-0332100330131023-1202303321310212-2322110002012103-3331001030301032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- Examples

<a id="canonical-0101130023012221-2200331210232131-2332222312033203-0112111310010202-0313131013100322-1002313331332312-0113121233200230-0311231111331201"></a>

### Complete configurations for `xcsh_app_api_group`

- [Data source](data-sources--app_api_group--examples--group-001.md#canonical-3313023333231021-0222020033001332-3103101120320231-3021031120132230-1111122322111110-3301023103232020-0230000020312220-0031200222103313): valid configuration.

<a id="canonical-3313023333231021-0222020033001332-3103101120320231-3021031120132230-1111122322111110-3301023103232020-0230000020312220-0031200222103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010)
- [Examples](data-sources--app_api_group--examples--group-001.md#canonical-0302100133100312-2121120001033302-0022331022132312-3310031221011031-0332100330131023-1202303321310212-2322110002012103-3331001030301032)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_api_group/data-source.tf`; digest `sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312`.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```
