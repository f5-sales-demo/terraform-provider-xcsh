---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_network_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1300, "body_sha256": "sha256:e4ca4234430eafdf28f5530f5cb4b063b0c1d728afcd8fd9fbc87f2cb0765687", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4d106a33f6bcd90b712f1a25c8242900f48156f66650342f586b89690227f887", "source_path": "examples/resources/xcsh_network_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_policy:example:resource", "parent_id": "xcsh-docs:resources:network_policy:examples", "path": "documentation/resources/network_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0213011323220312-3331100322223123-1210233332003233-0302322131301111-1210212201300100-3023300103002122-2311220212102332-2313013332003121", "registry_path": "docs/guides/resources--network_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Resource for xcsh_network_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["network_policyCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy/resource.tf`; digest `sha256:4d106a33f6bcd90b712f1a25c8242900f48156f66650342f586b89690227f887`.

```terraform
# NetworkPolicy Resource Example
# Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicy configuration
resource "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/examples/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
