---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_proxy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 983, "body_sha256": "sha256:edb7ba9925a9686ecd56d40df50606cb43c7c25756a5f381f9c2d63c34420540", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d2bc93690268dd4557ac75b114b1e4c362c8f8d843db7cc8d8dd1623ef5dbfe6", "source_path": "examples/data-sources/xcsh_proxy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:proxy:example:data-source", "parent_id": "xcsh-docs:data-sources:proxy:examples", "path": "documentation/data-sources/proxy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "data-sources", "registry_anchor": "canonical-3011233233011100-1322002211231232-3311121212313310-1330022232310010-1231033133323023-2233122330313330-3203030201132033-1021101202221100", "registry_path": "docs/guides/data-sources--proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/examples/data-source/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Data source for xcsh_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["proxyCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_proxy/data-source.tf`; digest `sha256:d2bc93690268dd4557ac75b114b1e4c362c8f8d843db7cc8d8dd1623ef5dbfe6`.

```terraform
# Proxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Proxy by name
data "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}

output "proxy_id" {
  value = data.xcsh_proxy.example.id
}
```
