---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_address_allocator."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1155, "body_sha256": "sha256:94e3034c332c4108a942c28ea42234889cce8940ad4410f2d4bfb261548a8226", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:address_allocator:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0d0c2e8aebf1264b1b9bd98f9f1564545390bccd08c3989c8f49aed2a844faa2", "source_path": "examples/resources/xcsh_address_allocator/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:address_allocator:example:resource", "parent_id": "xcsh-docs:resources:address_allocator:examples", "path": "documentation/resources/address_allocator/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "address_allocator", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2200201023213310-0330310001123331-2033001310110202-2023100221002013-1313323312211120-0002122100113112-1300221101123201-0230210303031011", "registry_path": "docs/guides/resources--address_allocator--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/address_allocator/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_address_allocator.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_address_allocator/resource.tf`; digest `sha256:0d0c2e8aebf1264b1b9bd98f9f1564545390bccd08c3989c8f49aed2a844faa2`.

```terraform
# AddressAllocator Resource Example
# Manages Address Allocator will create an address allocator object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AddressAllocator configuration
resource "xcsh_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "staging"

  address_pool = ["example-value"]
}
```
