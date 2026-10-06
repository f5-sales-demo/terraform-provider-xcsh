---
page_title: "xcsh_nginx_service_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery examples."
---

# xcsh_nginx_service_discovery examples

<a id="canonical-3001202301033101-0031000202120112-0333302030330103-1200011310233021-2031103233001231-1131103122222132-2013233222111311-0012112302231222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- Examples

<a id="canonical-1103201012221133-2210301002022122-2120013330120213-3230102310222022-1223001103120222-1310330330300110-1133100313022032-2333132221211333"></a>

### Complete configurations for `xcsh_nginx_service_discovery`

- [Resource](resources--nginx_service_discovery--examples--group-001.md#canonical-0203300202021231-3011031103221111-0023202232111213-2230220131210232-1301122131321333-1303010122113202-2311331103011311-2122021221313311): valid configuration.

<a id="canonical-0203300202021231-3011031103221111-0023202232111213-2230220131210232-1301122131321333-1303010122113202-2311331103011311-2122021221313311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_nginx_service_discovery](../resources/nginx_service_discovery.md#canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023)
- [Examples](resources--nginx_service_discovery--examples--group-001.md#canonical-3001202301033101-0031000202120112-0333302030330103-1200011310233021-2031103233001231-1131103122222132-2013233222111311-0012112302231222)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nginx_service_discovery/resource.tf`; digest `sha256:dd43ba5ab8c8a702c201e30638f18d6c0eafcf265b5180bf65bbcaa03289e5e6`.

```terraform
# NginxServiceDiscovery Resource Example
# Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NginxServiceDiscovery configuration
resource "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}
```
