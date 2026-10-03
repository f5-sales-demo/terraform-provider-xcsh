---
page_title: "Resource"
subcategory: "Networking"
description: "Resource for xcsh_endpoint."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1255, "body_sha256": "sha256:aeb877811ca020cf65b6df565da9362875c0e17fec3b2f9e78d516e54969f652", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e", "source_path": "examples/resources/xcsh_endpoint/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:endpoint:example:resource", "parent_id": "xcsh-docs:resources:endpoint:examples", "path": "documentation/resources/endpoint/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2212202010123110-2131002232331010-0321202122203111-2021233032121011-1211231100220030-2031111213233103-0132231032232021-2200211010231232", "registry_path": "docs/guides/resources--endpoint--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/examples/resource/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Resource for xcsh_endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_endpoint/resource.tf`; digest `sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e`.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/examples/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
