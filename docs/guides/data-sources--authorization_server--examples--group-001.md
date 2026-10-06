---
page_title: "xcsh_authorization_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server examples."
---

# xcsh_authorization_server examples

<a id="canonical-2033301013221130-0030022030303002-3031133301332223-3230301213310203-1301033002012133-1023011313333003-3130221223113223-2321031313000012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-3021220012103201-0130032102230113-0221231210301022-2323112200300110-3212010013210200-3300200232211023-2113102021313101-1310130133021322)
- Examples

<a id="canonical-1132211232011121-1213111101331202-3323021322221302-0011222102113031-0333220233000230-0213011333111311-3201220021202322-3032332302100010"></a>

### Complete configurations for `xcsh_authorization_server`

- [Data source](data-sources--authorization_server--examples--group-001.md#canonical-3033220020101301-1200103222123230-3302210133123003-1010001223210330-2100220321112211-3302123133112100-3031103103010010-1313301033330030): valid configuration.

<a id="canonical-3033220020101301-1200103222123230-3302210133123003-1010001223210330-2100220321112211-3302123133112100-3031103103010010-1313301033330030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_authorization_server](../data-sources/authorization_server.md#canonical-3021220012103201-0130032102230113-0221231210301022-2323112200300110-3212010013210200-3300200232211023-2113102021313101-1310130133021322)
- [Examples](data-sources--authorization_server--examples--group-001.md#canonical-2033301013221130-0030022030303002-3031133301332223-3230301213310203-1301033002012133-1023011313333003-3130221223113223-2321031313000012)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authorization_server/data-source.tf`; digest `sha256:268fe1444519cbb845d8fcb9ea236d26293e2aeafb324fad41e08e7bc571ffcb`.

```terraform
# AuthorizationServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AuthorizationServer by name
data "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"
}

output "authorization_server_id" {
  value = data.xcsh_authorization_server.example.id
}
```
