---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_proxy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1193, "body_sha256": "sha256:94bfa0d5c625b045ef0e7211d0990d20387a03551a0821a94ccbdc0f37b64049", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8d3ef9470714fab754f0f7631e40f5f93a8f5acfb3c6def25dff74cba83c2ef7", "source_path": "examples/resources/xcsh_proxy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:proxy:example:resource", "parent_id": "xcsh-docs:resources:proxy:examples", "path": "documentation/resources/proxy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1131121133130333-1031102020110210-0003123302322333-0011132131332003-1313211113301313-1021103110320002-3333022302000102-2301322011030023", "registry_path": "docs/guides/resources--proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_proxy/resource.tf`; digest `sha256:8d3ef9470714fab754f0f7631e40f5f93a8f5acfb3c6def25dff74cba83c2ef7`.

```terraform
# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Proxy configuration
resource "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/examples/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
