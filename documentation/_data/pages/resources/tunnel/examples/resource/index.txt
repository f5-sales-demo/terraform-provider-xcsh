---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_tunnel."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 945, "body_sha256": "sha256:8fe5ed0f5b8d836b3173c5169668498165b9e3040908e9566eac8e4db100b3ba", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d8686dd99ea2784f4ce424a8e06432fe8502eabcb7549c1f57d3a841fff62174", "source_path": "examples/resources/xcsh_tunnel/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tunnel:example:resource", "parent_id": "xcsh-docs:resources:tunnel:examples", "path": "documentation/resources/tunnel/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3120322120112022-0023131222020213-3313233331303131-1021012313112101-0200300131221330-0222133001202000-3310231301110112-3323213303332023", "registry_path": "docs/guides/resources--tunnel--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
