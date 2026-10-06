---
page_title: "xcsh_authorization_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server examples."
---

# xcsh_authorization_server examples

<a id="canonical-0123232121303011-2013012103001211-1021110113133132-3012133011121032-2011112332012332-0013012201121022-2123332313132320-2212100302331330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-3123011103301332-2020213322103110-2220312130201323-0120021121102210-3010302120211132-3100321223002321-1020323131322210-3122032302302331)
- Examples

<a id="canonical-2202331103211020-3100300001213111-0120332020033101-0223103102001023-3213100330002223-0312221021312302-1131333103022220-0030322112210132"></a>

### Complete configurations for `xcsh_authorization_server`

- [Resource](resources--authorization_server--examples--group-001.md#canonical-1110123012323002-2000131221012030-2202032322113010-0301211100213311-0132010130001223-3203220221023213-3131223222131113-1301133220321013): valid configuration.

<a id="canonical-1110123012323002-2000131221012030-2202032322113010-0301211100213311-0132010130001223-3203220221023213-3131223222131113-1301133220321013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-3123011103301332-2020213322103110-2220312130201323-0120021121102210-3010302120211132-3100321223002321-1020323131322210-3122032302302331)
- [Examples](resources--authorization_server--examples--group-001.md#canonical-0123232121303011-2013012103001211-1021110113133132-3012133011121032-2011112332012332-0013012201121022-2123332313132320-2212100302331330)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authorization_server/resource.tf`; digest `sha256:bbb05ccd40cd71dbeef4a9b6ed31c1cec2684fa158b0357f99c63ef6b31e6af5`.

```terraform
# AuthorizationServer Resource Example
# Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AuthorizationServer configuration
resource "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"

  jwks_uri = "example-value"
}
```
