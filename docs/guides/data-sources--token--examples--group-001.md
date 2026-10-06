---
page_title: "xcsh_token examples"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token examples."
---

# xcsh_token examples

<a id="canonical-1020321111222301-2033123121331132-3331211233301231-1202000032310201-3201211232323022-2130002310311310-3103122312103310-0031001213131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_token](../data-sources/token.md#canonical-2000130020133110-1113012332320202-2223121002312130-2001210311231111-3323332201023323-2311120322030210-1131312201010133-0322033200102202)
- Examples

<a id="canonical-3303011322100331-3200101201023302-1011033320200301-3011003231301101-1212212311323231-3330001330302332-0330213301210233-2011100312122002"></a>

### Complete configurations for `xcsh_token`

- [Data source](data-sources--token--examples--group-001.md#canonical-0102202302032300-2131201001103020-1213132120331221-0031010123200100-0120100213311021-1021133231120011-0320202323201223-2323111231111021): valid configuration.

<a id="canonical-0102202302032300-2131201001103020-1213132120331221-0031010123200100-0120100213311021-1021133231120011-0320202323201223-2323111231111021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_token](../data-sources/token.md#canonical-2000130020133110-1113012332320202-2223121002312130-2001210311231111-3323332201023323-2311120322030210-1131312201010133-0322033200102202)
- [Examples](data-sources--token--examples--group-001.md#canonical-1020321111222301-2033123121331132-3331211233301231-1202000032310201-3201211232323022-2130002310311310-3103122312103310-0031001213131320)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_token/data-source.tf`; digest `sha256:2be846025946a447e16fd9c1be64b48387d966481fa919712ec2c2311d266aee`.

```terraform
# Token Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Token by name
data "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
}

output "token_id" {
  value = data.xcsh_token.example.id
}
```
