---
page_title: "xcsh_crl examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl examples."
---

# xcsh_crl examples

<a id="canonical-2303331213223002-2112213123320220-3331203230311231-2101022333102302-0112033003011221-1031110223130221-0010123102223023-3303312211123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-1303332111131203-2002232212001132-0030110222210023-0210330330101332-1012330233220233-3212220211023233-3303111232113130-2302300032211003)
- Examples

<a id="canonical-2113022310100123-0130023121011030-1220123332313232-1232230302323013-2033323113231232-0020101002000021-0303212023311302-1303221200033202"></a>

### Complete configurations for `xcsh_crl`

- [Data source](data-sources--crl--examples--group-001.md#canonical-3020121233030311-3122033322301333-2312211200310011-2031300000230123-2113303100023233-1310201220030010-2102030210123201-0312131312201322): valid configuration.

<a id="canonical-3020121233030311-3122033322301333-2312211200310011-2031300000230123-2113303100023233-1310201220030010-2102030210123201-0312131312201322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_crl](../data-sources/crl.md#canonical-1303332111131203-2002232212001132-0030110222210023-0210330330101332-1012330233220233-3212220211023233-3303111232113130-2302300032211003)
- [Examples](data-sources--crl--examples--group-001.md#canonical-2303331213223002-2112213123320220-3331203230311231-2101022333102302-0112033003011221-1031110223130221-0010123102223023-3303312211123230)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_crl/data-source.tf`; digest `sha256:7250b8110197f6a15add9c7d7bb8ba0193edd8faf34e61501e055b382c218d8a`.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```
