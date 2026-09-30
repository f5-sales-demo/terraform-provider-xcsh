---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1167, "body_sha256": "sha256:1c291701c5e05f3dead608ca25d3f0e2bf1c7727a6b450d8ca37d0ae24f64613", "child_ids": [], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:52717a495ae2d7db46b088a16cabbff85cbe56ae60a8545793e0ce939a5f736c", "source_path": "examples/resources/xcsh_bigip_http_proxy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bigip_http_proxy:example:resource", "parent_id": "xcsh-docs:resources:bigip_http_proxy:examples", "path": "documentation/resources/bigip_http_proxy/examples/resource/index.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bigip_http_proxy/resource.tf`; digest `sha256:52717a495ae2d7db46b088a16cabbff85cbe56ae60a8545793e0ce939a5f736c`.

```terraform
# BigIPHTTPProxy Resource Example
# Manages BIG-IP HTTP Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BigIPHTTPProxy configuration
resource "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/examples/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
