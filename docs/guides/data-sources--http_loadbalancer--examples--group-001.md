---
page_title: "xcsh_http_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer examples."
---

# xcsh_http_loadbalancer examples

<a id="canonical-2012103322223222-2130303310322103-3112301200231103-1003312232213332-1311123122123311-3322111313213332-3221000203301021-1012200122123110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- Examples

<a id="canonical-1321100300232023-1002313033030320-0131102210010211-3223113311322001-3212313103333020-1321032311133211-2032210200120102-3102203131200330"></a>

### Complete configurations for `xcsh_http_loadbalancer`

- [Data source](data-sources--http_loadbalancer--examples--group-001.md#canonical-3130300133300312-1122213020302223-1201233323112022-2232022310002133-0210131131322231-2010030002222320-2320131131002120-3230202332313130): valid configuration.

<a id="canonical-3130300133300312-1122213020302223-1201233323112022-2232022310002133-0210131131322231-2010030002222320-2320131131002120-3230202332313130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Examples](data-sources--http_loadbalancer--examples--group-001.md#canonical-2012103322223222-2130303310322103-3112301200231103-1003312232213332-1311123122123311-3322111313213332-3221000203301021-1012200122123110)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_http_loadbalancer/data-source.tf`; digest `sha256:f7d3fa4bb60b803988578a72b5dac4a76f304bac2925ae4fdfe075100bed37ad`.

```terraform
# HTTPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing HTTPLoadBalancer by name
data "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"
}

output "http_loadbalancer_id" {
  value = data.xcsh_http_loadbalancer.example.id
}
```
