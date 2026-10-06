---
page_title: "xcsh_nginx_service_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery examples."
---

# xcsh_nginx_service_discovery examples

<a id="canonical-1312122322200101-0103213331231131-3012332303232320-3111010023031103-1130020230030100-0101103120101222-1112023222022322-2320032102130001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- Examples

<a id="canonical-3202301033111301-2000232013311320-0030010231102011-1103001331232133-0303330132030222-2300112022231313-1213211033130320-2132113212302221"></a>

### Complete configurations for `xcsh_nginx_service_discovery`

- [Data source](data-sources--nginx_service_discovery--examples--group-001.md#canonical-0310300123231333-2012120212321021-0230032122031112-2211112131002020-0102221212302113-0320210320121210-2212101313231121-1113102022120021): valid configuration.

<a id="canonical-0310300123231333-2012120212321021-0230032122031112-2211112131002020-0102221212302113-0320210320121210-2212101313231121-1113102022120021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311)
- [Examples](data-sources--nginx_service_discovery--examples--group-001.md#canonical-1312122322200101-0103213331231131-3012332303232320-3111010023031103-1130020230030100-0101103120101222-1112023222022322-2320032102130001)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_service_discovery/data-source.tf`; digest `sha256:c46b908211689c5888bfaaf70a487544236d409e407922fa907fe86ca0588fa6`.

```terraform
# NginxServiceDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServiceDiscovery by name
data "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}

output "nginx_service_discovery_id" {
  value = data.xcsh_nginx_service_discovery.example.id
}
```
