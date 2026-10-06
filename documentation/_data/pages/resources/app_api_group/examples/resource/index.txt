---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_api_group."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1074, "body_sha256": "sha256:dd59c4f7c02290aed2f89754c87a0e0771112dd5abdaff6bdf3a34d079b0f42b", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f4976402381de0085eaa3ddefc5b729e1a883d869e4792bc2f8f960a6f9b6252", "source_path": "examples/resources/xcsh_app_api_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_api_group:example:resource", "parent_id": "xcsh-docs:resources:app_api_group:examples", "path": "documentation/resources/app_api_group/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3301200002210131-1021312213020221-2231032301122223-2332121010230332-0102102332223110-2001121313222212-3303121133031212-2012230103103100", "registry_path": "docs/guides/resources--app_api_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_app_api_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_api_group/resource.tf`; digest `sha256:f4976402381de0085eaa3ddefc5b729e1a883d869e4792bc2f8f960a6f9b6252`.

```terraform
# AppAPIGroup Resource Example
# Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppAPIGroup configuration
resource "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}
```
