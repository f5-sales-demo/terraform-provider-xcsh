---
page_title: "xcsh_nginx_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_server examples."
---

# xcsh_nginx_server examples

<a id="canonical-2300333101313203-3301002102123333-2223020201133011-3223003103223111-2010103032122130-3001221310203300-2300030133223022-2102202303301131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-2001223000223133-3323332323212021-3322301333223013-0302011011112310-1200100232112222-1312233312332302-3320223103231232-2123102210132102)
- Examples

<a id="canonical-0322133033221202-0132331000003130-2001011102100131-2320230121333101-1020322322221123-2303311203110100-1130110123010103-2232031310111023"></a>

### Complete configurations for `xcsh_nginx_server`

- [Data source](data-sources--nginx_server--examples--group-001.md#canonical-3222330301303300-0232330221032232-3300301221130110-1220023211310012-3300131212330201-2103330102213200-1301120232101131-2312031002211310): valid configuration.

<a id="canonical-3222330301303300-0232330221032232-3300301221130110-1220023211310012-3300131212330201-2103330102213200-1301120232101131-2312031002211310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-2001223000223133-3323332323212021-3322301333223013-0302011011112310-1200100232112222-1312233312332302-3320223103231232-2123102210132102)
- [Examples](data-sources--nginx_server--examples--group-001.md#canonical-2300333101313203-3301002102123333-2223020201133011-3223003103223111-2010103032122130-3001221310203300-2300030133223022-2102202303301131)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_server/data-source.tf`; digest `sha256:255e0e3089bba8c6998465d72659b209b4c83e33fec880e96e0bccd476d39e6e`.

```terraform
# NginxServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServer by name
data "xcsh_nginx_server" "example" {
  name      = "example-nginx-server"
  namespace = "staging"
}

output "nginx_server_id" {
  value = data.xcsh_nginx_server.example.id
}
```
