---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1266, "body_sha256": "sha256:b3bb83065c0df4c94b5f8193d7ef6ccd46deb697f20a24bf6e3ec9d1e2a58407", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:52717a495ae2d7db46b088a16cabbff85cbe56ae60a8545793e0ce939a5f736c", "source_path": "examples/resources/xcsh_bigip_http_proxy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bigip_http_proxy:example:resource", "parent_id": "xcsh-docs:resources:bigip_http_proxy:examples", "path": "documentation/resources/bigip_http_proxy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3001200003221011-3032223322001133-0000130213231133-0103130103101113-3133232033000130-3201133320001233-2302011022303010-3302111133330110", "registry_path": "docs/guides/resources--bigip_http_proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_bigip_http_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
