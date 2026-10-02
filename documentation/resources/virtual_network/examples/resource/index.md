---
page_title: "Resource"
subcategory: "Networking"
description: "Resource for xcsh_virtual_network."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1276, "body_sha256": "sha256:076d32e149153efdfd68a9935f5970939ae70513bd7bdd306b99f347b232ede2", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473", "source_path": "examples/resources/xcsh_virtual_network/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:virtual_network:example:resource", "parent_id": "xcsh-docs:resources:virtual_network:examples", "path": "documentation/resources/virtual_network/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1322123120312230-3132110232130120-0302230100113022-0331311030201303-3002210133220221-3332123213020313-3120112033301002-2020030121000232", "registry_path": "docs/guides/resources--virtual_network--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_virtual_network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_network/resource.tf`; digest `sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473`.

```terraform
# VirtualNetwork Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/examples/)
- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
