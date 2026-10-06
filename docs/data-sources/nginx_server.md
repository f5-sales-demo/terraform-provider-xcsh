---
page_title: "xcsh_nginx_server"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_server."
---

# xcsh_nginx_server

<a id="canonical-2001223000223133-3323332323212021-3322301333223013-0302011011112310-1200100232112222-1312233312332302-3320223103231232-2123102210132102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_nginx_server

Reads an NGINX server block configuration.

<a id="canonical-3013003001232002-3022311302001302-1323332000023002-3022123030011313-0030122323202221-1320230321313002-1213200033022323-2121102123023133"></a>

### Prerequisites for `xcsh_nginx_server`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2220110033130101-3023112000313010-1312000123210021-1103211123010012-2032110031001231-2322313230130210-3320210220332202-2213322131313122"></a>

### Minimal configuration for `xcsh_nginx_server`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3223010222010133-1221001221311213-1033122300012232-0333013031031303-2001202012023231-1322321120331033-2011000013121133-1321301330111001"></a>

### Root configuration for `xcsh_nginx_server`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2012230211222320-3123310002332202-1312003210221220-2132022201230113-2110221322201103-2303102132003132-1100312031002102-2103001031002100"></a>

### Explore this collection for `xcsh_nginx_server`

- [Property reference](../guides/data-sources--nginx_server--reference--group-001.md#canonical-1023323233213133-0302133111112113-0100222132220233-3300220100330002-1122012022213202-1333110201013021-3010020312201133-3123122020002010)
- [Examples](../guides/data-sources--nginx_server--examples--group-001.md#canonical-2300333101313203-3301002102123333-2223020201133011-3223003103223111-2010103032122130-3001221310203300-2300030133223022-2102202303301131)
