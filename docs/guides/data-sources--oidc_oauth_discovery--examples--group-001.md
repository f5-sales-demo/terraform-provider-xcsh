---
page_title: "xcsh_oidc_oauth_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_oidc_oauth_discovery examples."
---

# xcsh_oidc_oauth_discovery examples

<a id="canonical-3101302130120123-0223201112000011-3300033320013211-0000131221321003-1320311131121223-0001003132213223-1313133333323201-0121321211022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-1212212133022200-1202310302030030-1123103212311111-3102001300103310-0102100100032330-3310033310100102-1213021002223010-3132011230221122)
- Examples

<a id="canonical-0301200330121320-0030033313010120-0213202231031023-1232113120032310-1023123112120302-3331023121010210-0021032233332303-0200100113200130"></a>

### Complete configurations for `xcsh_oidc_oauth_discovery`

- [Data source](data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-0023303322201221-1101130221112332-1003310322131303-1310210330333031-0122200230000202-0020020100013300-3230230112031031-0303232301311133): valid configuration.

<a id="canonical-0023303322201221-1101130221112332-1003310322131303-1310210330333031-0122200230000202-0020020100013300-3230230112031031-0303232301311133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_oidc_oauth_discovery](../data-sources/oidc_oauth_discovery.md#canonical-1212212133022200-1202310302030030-1123103212311111-3102001300103310-0102100100032330-3310033310100102-1213021002223010-3132011230221122)
- [Examples](data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-3101302130120123-0223201112000011-3300033320013211-0000131221321003-1320311131121223-0001003132213223-1313133333323201-0121321211022220)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_oidc_oauth_discovery/data-source.tf`; digest `sha256:fc43fdc88285748c686ea5a405a7c988251f6e6a9d61c4635ba5b40db4b85e6a`.

```terraform
# OIDCOauthDiscovery DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_oidc_oauth_discovery" "example" {
  namespace = "example-value"
}

output "oidc_oauth_discovery_result" {
  value = data.xcsh_oidc_oauth_discovery.example
}
```
