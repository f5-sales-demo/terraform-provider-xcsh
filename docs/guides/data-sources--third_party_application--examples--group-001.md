---
page_title: "xcsh_third_party_application examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_third_party_application examples."
---

# xcsh_third_party_application examples

<a id="canonical-0120101222002302-3233320311331323-0310120130031313-1113121231111320-1332122011122230-1002311331331021-2313102221223323-3132113203212103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md#canonical-2320222231123022-0323332222212132-3300113320120123-0222011302331223-3222001221131320-2123313213312301-2022200110110030-3312032011311320)
- Examples

<a id="canonical-2113330331202202-2110212000302020-3220331221130013-2202203120221211-3332012312230130-1102213321022302-1213110310103322-1123232131311113"></a>

### Complete configurations for `xcsh_third_party_application`

- [Data source](data-sources--third_party_application--examples--group-001.md#canonical-0200101300201233-2010323323200232-1302213031311002-0003000231113013-2132310012222121-2002232120011103-0022033012313231-1122120330222102): valid configuration.

<a id="canonical-0200101300201233-2010323323200232-1302213031311002-0003000231113013-2132310012222121-2002232120011103-0022033012313231-1122120330222102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md#canonical-2320222231123022-0323332222212132-3300113320120123-0222011302331223-3222001221131320-2123313213312301-2022200110110030-3312032011311320)
- [Examples](data-sources--third_party_application--examples--group-001.md#canonical-0120101222002302-3233320311331323-0310120130031313-1113121231111320-1332122011122230-1002311331331021-2313102221223323-3132113203212103)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_third_party_application/data-source.tf`; digest `sha256:e3e758943d32635744f0c37b6cc645f6b182639afd8477f0071a1bfaa5279e53`.

```terraform
# ThirdPartyApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ThirdPartyApplication by name
data "xcsh_third_party_application" "example" {
  name      = "example-third-party-application"
  namespace = "staging"
}

output "third_party_application_id" {
  value = data.xcsh_third_party_application.example.id
}
```
