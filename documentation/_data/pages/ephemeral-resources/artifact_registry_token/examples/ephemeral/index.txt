---
page_title: "Ephemeral"
subcategory: ""
description: "Ephemeral for xcsh_artifact_registry_token."
xcsh_docs: {"aliases": ["ephemeral"], "body_bytes": 1271, "body_sha256": "sha256:ceaa9a96e51177777816c92a406a50411fe7d6959f8d2eeb53029709b6a8c539", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f91e2c9b88b3ed39e41f1e34edb4d5addb5289872f1ac18a74ddd0ff3f1aec6f", "source_path": "examples/ephemeral-resources/xcsh_artifact_registry_token/ephemeral.tf", "validation": "terraform validate"}, "id": "xcsh-docs:ephemeral-resources:artifact_registry_token:example:ephemeral", "parent_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:examples", "path": "documentation/ephemeral-resources/artifact_registry_token/examples/ephemeral/index.md", "product": "distributed-cloud", "provider_name": "artifact_registry_token", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-0331000100300321-1112003032322321-3321000221031110-0311121332232201-2031031331021200-3333230213330010-2031301221220022-0123102213131101", "registry_path": "docs/guides/ephemeral-resources--artifact_registry_token--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["ephemeral"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/artifact_registry_token/examples/ephemeral/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Ephemeral for xcsh_artifact_registry_token.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Ephemeral

Breadcrumbs:

- [xcsh_artifact_registry_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/examples/)
- Ephemeral

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/ephemeral-resources/xcsh_artifact_registry_token/ephemeral.tf`; digest `sha256:f91e2c9b88b3ed39e41f1e34edb4d5addb5289872f1ac18a74ddd0ff3f1aec6f`.

```terraform
# ArtifactRegistryToken EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_artifact_registry_token" "example" {
  namespace = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/examples/)
- [xcsh_artifact_registry_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/)
