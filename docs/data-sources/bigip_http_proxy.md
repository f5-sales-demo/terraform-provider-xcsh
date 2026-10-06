---
page_title: "xcsh_bigip_http_proxy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy."
---

# xcsh_bigip_http_proxy

<a id="canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bigip_http_proxy

Reads BIG-IP HTTP Proxy information from F5 Distributed Cloud.

<a id="canonical-1002110102000201-1310333003123112-1003130310233222-0201301112012332-3030331230131100-0112122013001233-1231212032313321-1121202110331013"></a>

### Prerequisites for `xcsh_bigip_http_proxy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0230113033031323-3200003121233223-0212002011301001-3200331030222233-3320132332211230-3010010331102302-0103233120131231-3223312223212101"></a>

### Minimal configuration for `xcsh_bigip_http_proxy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0132223010100331-0213121200131003-1113131020002211-2330123010033021-1203103031010030-0013110112010003-1132112232333012-2201103110033132"></a>

### Root configuration for `xcsh_bigip_http_proxy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2333230030012330-1230323013113102-2311100233200003-2213303120200002-1223301120101031-3202132310131201-2100102200112033-0010302203131321"></a>

### Explore this collection for `xcsh_bigip_http_proxy`

- [Property reference](../guides/data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [Examples](../guides/data-sources--bigip_http_proxy--examples--group-001.md#canonical-1212221012032032-0310133313321202-0020322303033021-3211313313332203-0231022200011013-1020122033132320-0303011021123030-2121103122210030)
