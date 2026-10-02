---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_tunnel."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1149, "body_sha256": "sha256:4ad2373f1100278d74c10cca1b4d6f7d742c1780e5355a8f1964c309d2f03cf1", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d8686dd99ea2784f4ce424a8e06432fe8502eabcb7549c1f57d3a841fff62174", "source_path": "examples/resources/xcsh_tunnel/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tunnel:example:resource", "parent_id": "xcsh-docs:resources:tunnel:examples", "path": "documentation/resources/tunnel/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3120322120112022-0023131222020213-3313233331303131-1021012313112101-0200300131221330-0222133001202000-3310231301110112-3323213303332023", "registry_path": "docs/guides/resources--tunnel--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tunnel/resource.tf`; digest `sha256:d8686dd99ea2784f4ce424a8e06432fe8502eabcb7549c1f57d3a841fff62174`.

```terraform
# Tunnel Resource Example
# Manages tunnel in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tunnel configuration
resource "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/examples/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
