---
page_title: "xcsh_api_testing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing examples."
---

# xcsh_api_testing examples

<a id="canonical-2332231100103211-2032311033113031-3002312201302303-1303231131130130-0310212022223232-1201123210232101-0230031123302302-0012222311112303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- Examples

<a id="canonical-3001230030000203-3031123133320203-2322202122231023-0211003330033003-3201202101132033-3223310210230323-1031102010103100-0331230202331121"></a>

### Complete configurations for `xcsh_api_testing`

- [Data source](data-sources--api_testing--examples--group-001.md#canonical-3330001211231130-1030221022230232-2223323322000123-0133313130201001-0132022222302222-2221123013220202-0231300013303320-0121231221110120): valid configuration.

<a id="canonical-3330001211231130-1030221022230232-2223323322000123-0133313130201001-0132022222302222-2221123013220202-0231300013303320-0121231221110120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Examples](data-sources--api_testing--examples--group-001.md#canonical-2332231100103211-2032311033113031-3002312201302303-1303231131130130-0310212022223232-1201123210232101-0230031123302302-0012222311112303)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_testing/data-source.tf`; digest `sha256:509ae39ca1a6610bcc317fbc3a5f0883f2fcba290d264181d678caabac1f11fd`.

```terraform
# APITesting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APITesting by name
data "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}

output "api_testing_id" {
  value = data.xcsh_api_testing.example.id
}
```
