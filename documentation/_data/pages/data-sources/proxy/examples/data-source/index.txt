---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_proxy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1190, "body_sha256": "sha256:3dff74ffbb9213d9d6dca35331d14d6de92c7c84e2c90dd6627445f81ab2a75c", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d2bc93690268dd4557ac75b114b1e4c362c8f8d843db7cc8d8dd1623ef5dbfe6", "source_path": "examples/data-sources/xcsh_proxy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:proxy:example:data-source", "parent_id": "xcsh-docs:data-sources:proxy:examples", "path": "documentation/data-sources/proxy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3011233233011100-1322002211231232-3311121212313310-1330022232310010-1231033133323023-2233122330313330-3203030201132033-1021101202221100", "registry_path": "docs/guides/data-sources--proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/examples/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
