---
page_title: "xcsh_malicious_user_mitigation examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation examples."
---

# xcsh_malicious_user_mitigation examples

<a id="canonical-2330111100130201-3001331323202112-0233331232123201-0231220331113032-0222113203313111-3200320101001133-2111100013032023-3232130123320213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- Examples

<a id="canonical-2233302212101002-2002010111312311-0000203323112232-0022200113022020-3122001113302131-0111101110203321-0032231211232031-1102320110230112"></a>

### Complete configurations for `xcsh_malicious_user_mitigation`

- [Data source](data-sources--malicious_user_mitigation--examples--group-001.md#canonical-2011013332031323-0002030011020021-2220122132030012-1011320320132210-2201121322030303-2010320300231203-0002323323202132-2323203033132212): valid configuration.

<a id="canonical-2011013332031323-0002030011020021-2220122132030012-1011320320132210-2201121322030303-2010320300231203-0002323323202132-2323203033132212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211)
- [Examples](data-sources--malicious_user_mitigation--examples--group-001.md#canonical-2330111100130201-3001331323202112-0233331232123201-0231220331113032-0222113203313111-3200320101001133-2111100013032023-3232130123320213)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_malicious_user_mitigation/data-source.tf`; digest `sha256:48191839abab419f70a338b5b737300e10e8068606e35e5a0829c82bf6fe5846`.

```terraform
# MaliciousUserMitigation Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MaliciousUserMitigation by name
data "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}

output "malicious_user_mitigation_id" {
  value = data.xcsh_malicious_user_mitigation.example.id
}
```
