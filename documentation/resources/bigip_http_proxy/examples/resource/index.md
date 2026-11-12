---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1032, "body_sha256": "sha256:2c04871b942c01fd224cbbff9ab42945e19f5833021623c725a9ba915e05d0db", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:52717a495ae2d7db46b088a16cabbff85cbe56ae60a8545793e0ce939a5f736c", "source_path": "examples/resources/xcsh_bigip_http_proxy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bigip_http_proxy:example:resource", "parent_id": "xcsh-docs:resources:bigip_http_proxy:examples", "path": "documentation/resources/bigip_http_proxy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-3001200003221011-3032223322001133-0000130213231133-0103130103101113-3133232033000130-3201133320001233-2302011022303010-3302111133330110", "registry_path": "docs/guides/resources--bigip_http_proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/examples/resource/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Resource for xcsh_bigip_http_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
