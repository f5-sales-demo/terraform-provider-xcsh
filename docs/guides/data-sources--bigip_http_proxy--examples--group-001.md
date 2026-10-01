---
page_title: "xcsh_bigip_http_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy examples."
---

# xcsh_bigip_http_proxy examples

<a id="canonical-66a4638e347f7e6208eb33c9e5df7fa32d2a01474868f7b8331496cc994da90c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d83cb601c3ea3322dfbf3f3bd39b8ad1a708d8beb3eaaca495dd9e11c06a8b75"></a>

## Examples — Examples / 936ece71b1b7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- Examples

<a id="canonical-1684b736706e157eeb17cc154c94c5d10fa0a5e65131bb10df1fb80296fb2934"></a>

## Complete configurations — Examples / 936ece71b1b7 / 3

- [Data source](data-sources--bigip_http_proxy--examples--group-001.md#canonical-af79b65f2bbf67653435184b3d775d44c3508db164f6f5cfdfa2d1f6e2feec2f): valid configuration.

<a id="canonical-f4c05bba2a96111d7e889a73f17ac9033581aa57d40ab3095602f94388b4bc02"></a>

## Next pages — Examples / 936ece71b1b7 / 4

- [Data source](data-sources--bigip_http_proxy--examples--group-001.md#canonical-af79b65f2bbf67653435184b3d775d44c3508db164f6f5cfdfa2d1f6e2feec2f)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)

<a id="canonical-af79b65f2bbf67653435184b3d775d44c3508db164f6f5cfdfa2d1f6e2feec2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-582f01e015df1326a129f872ec0254e3463c4f058ef4563595b363c353aa4e2f"></a>

## Data source — Data source / f40b13f80710 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
- [Examples](data-sources--bigip_http_proxy--examples--group-001.md#canonical-66a4638e347f7e6208eb33c9e5df7fa32d2a01474868f7b8331496cc994da90c)
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

<a id="canonical-66da097b4f614c60e8dbd9ca06cb34f1c845e046830da8c1e6168a472a1becd9"></a>

## Next pages — Data source / f40b13f80710 / 3

- [Examples](data-sources--bigip_http_proxy--examples--group-001.md#canonical-66a4638e347f7e6208eb33c9e5df7fa32d2a01474868f7b8331496cc994da90c)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0bf4ebbc404da6557682625d19cc0a58e283072897707ad19978a417dbde5416)
