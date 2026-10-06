---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_link."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1012, "body_sha256": "sha256:16b80de7b9e3c6f7b56d56bb07ede0f2c5ff4ea8ce0a68e1bdce7fde3106be99", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00", "source_path": "examples/resources/xcsh_cloud_link/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_link:example:resource", "parent_id": "xcsh-docs:resources:cloud_link:examples", "path": "documentation/resources/cloud_link/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1010233102112003-1120030021313310-2013011023303020-3201222103232003-2211131202122210-2010333023323331-2231022231232122-3023120032022030", "registry_path": "docs/guides/resources--cloud_link--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_cloud_link.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_link/resource.tf`; digest `sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00`.

```terraform
# CloudLink Resource Example
# Manages new CloudLink with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudLink configuration
resource "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}
```
