---
page_title: "Ephemeral"
subcategory: ""
description: "Ephemeral for xcsh_artifact_registry_token."
xcsh_docs: {"aliases": ["ephemeral"], "body_bytes": 996, "body_sha256": "sha256:82a4ac3e034345a906f9803bbc3b0e8ed6987a0187a9d2e4a4a00cd5363df2bb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f91e2c9b88b3ed39e41f1e34edb4d5addb5289872f1ac18a74ddd0ff3f1aec6f", "source_path": "examples/ephemeral-resources/xcsh_artifact_registry_token/ephemeral.tf", "validation": "terraform validate"}, "id": "xcsh-docs:ephemeral-resources:artifact_registry_token:example:ephemeral", "parent_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:examples", "path": "documentation/ephemeral-resources/artifact_registry_token/examples/ephemeral/index.md", "product": "distributed-cloud", "provider_name": "artifact_registry_token", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-0331000100300321-1112003032322321-3321000221031110-0311121332232201-2031031331021200-3333230213330010-2031301221220022-0123102213131101", "registry_path": "docs/guides/ephemeral-resources--artifact_registry_token--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["ephemeral"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/artifact_registry_token/examples/ephemeral/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Ephemeral for xcsh_artifact_registry_token.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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
