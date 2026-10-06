---
page_title: "xcsh_access_active_session examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session examples."
---

# xcsh_access_active_session examples

<a id="canonical-0133223210300032-1021113231320322-2030300200002213-0122010310110213-2022102020233203-1100022112100102-0200030100333233-3121121301222132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-1000030313032121-0103232201301200-2011330102102012-2203212012120203-2301010031212003-1303323312131012-1222202312111300-1032323113103023)
- Examples

<a id="canonical-0202210312111222-2312133223023210-0103001103332113-2212103313230203-2002213111210111-0110020202313110-3312202022200212-1020202113100031"></a>

### Complete configurations for `xcsh_access_active_session`

- [Data source](data-sources--access_active_session--examples--group-001.md#canonical-1122210201123331-0211211312033112-1200212103201123-0130011330021332-0220333011310002-3213212230011310-1131320222212210-1221321100012001): valid configuration.

<a id="canonical-1122210201123331-0211211312033112-1200212103201123-0130011330021332-0220333011310002-3213212230011310-1131320222212210-1221321100012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-1000030313032121-0103232201301200-2011330102102012-2203212012120203-2301010031212003-1303323312131012-1222202312111300-1032323113103023)
- [Examples](data-sources--access_active_session--examples--group-001.md#canonical-0133223210300032-1021113231320322-2030300200002213-0122010310110213-2022102020233203-1100022112100102-0200030100333233-3121121301222132)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_access_active_session/data-source.tf`; digest `sha256:d3411f6908febc7c8a9bf8b482b2e026429a95e16a75b366bc2a33e4f32326f7`.

```terraform
# AccessActiveSession DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_session" "example" {
  id        = "example-value"
  namespace = "example-value"
}

output "access_active_session_result" {
  value = data.xcsh_access_active_session.example
}
```
