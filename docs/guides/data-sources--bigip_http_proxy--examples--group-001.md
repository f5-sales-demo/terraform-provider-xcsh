---
page_title: "xcsh_bigip_http_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy examples."
---

# xcsh_bigip_http_proxy examples

<a id="canonical-1212221012032032-0310133313321202-0020322303033021-3211313313332203-0231022200011013-1020122033132320-0303011021123030-2121103122210030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- Examples

<a id="canonical-3120033023120001-3003322203030202-3133233303330323-3103212320223101-2213002031202332-2303322222302210-2111313121320101-3000122220231311"></a>

### Complete configurations for `xcsh_bigip_http_proxy`

- [Data source](data-sources--bigip_http_proxy--examples--group-001.md#canonical-2233132123121133-0223233312131211-0310031101201023-0331131311311010-3003110020312301-1210331233113033-3133220231013312-3202333232300233): valid configuration.

<a id="canonical-2233132123121133-0223233312131211-0310031101201023-0331131311311010-3003110020312301-1210331233113033-3133220231013312-3202333232300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Examples](data-sources--bigip_http_proxy--examples--group-001.md#canonical-1212221012032032-0310133313321202-0020322303033021-3211313313332203-0231022200011013-1020122033132320-0303011021123030-2121103122210030)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bigip_http_proxy/data-source.tf`; digest `sha256:436b4e693b05709f4b55969967c5b9c433a210d85502a116c164d87cdf218449`.

```terraform
# BigIPHTTPProxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPHTTPProxy by name
data "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}

output "bigip_http_proxy_id" {
  value = data.xcsh_bigip_http_proxy.example.id
}
```
