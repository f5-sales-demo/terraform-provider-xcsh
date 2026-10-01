---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_address_allocator."
xcsh_docs: {"aliases": [], "body_bytes": 1392, "body_sha256": "sha256:e7824cedcd89d96c38e373a69373a3e2783bdd6d70f7a8fa151f75e05d4afec8", "child_ids": [], "collection_id": "xcsh-docs:resources:address_allocator:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0d0c2e8aebf1264b1b9bd98f9f1564545390bccd08c3989c8f49aed2a844faa2", "source_path": "examples/resources/xcsh_address_allocator/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:address_allocator:example:resource", "parent_id": "xcsh-docs:resources:address_allocator:examples", "path": "documentation/resources/address_allocator/examples/resource/index.md", "provider_name": "address_allocator", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/address_allocator/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_address_allocator.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["address_allocatorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/examples/)
- [xcsh_address_allocator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/address_allocator/)
