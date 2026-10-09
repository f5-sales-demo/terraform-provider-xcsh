---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_nat_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1072, "body_sha256": "sha256:e40aec35c78c964b4fd5edf8f348a226f8a40fce5fd565c3045e127e9b84867a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9976ae5d67503e6c32870094ecdf09b471ea80cfbc484860002722af35c9de08", "source_path": "examples/resources/xcsh_nat_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:nat_policy:example:resource", "parent_id": "xcsh-docs:resources:nat_policy:examples", "path": "documentation/resources/nat_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2302320100223323-1103033101202233-2011332102122012-2321313122020033-1303211313033103-1112322132223331-1321113120223303-0002311000301120", "registry_path": "docs/guides/resources--nat_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_nat_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["nat_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nat_policy/resource.tf`; digest `sha256:9976ae5d67503e6c32870094ecdf09b471ea80cfbc484860002722af35c9de08`.

```terraform
# NATPolicy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NATPolicy configuration
resource "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}
```
