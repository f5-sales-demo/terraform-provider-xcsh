---
page_title: "Ephemeral"
subcategory: ""
description: "Ephemeral for xcsh_artifact_registry_token."
xcsh_docs: {"aliases": [], "body_bytes": 1172, "body_sha256": "sha256:86baa5b79d5184126976ad7d49d6ef04d7d5540f85ce503276d6504bcf2efdc6", "child_ids": [], "collection_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f91e2c9b88b3ed39e41f1e34edb4d5addb5289872f1ac18a74ddd0ff3f1aec6f", "source_path": "examples/ephemeral-resources/xcsh_artifact_registry_token/ephemeral.tf", "validation": "terraform validate"}, "id": "xcsh-docs:ephemeral-resources:artifact_registry_token:example:ephemeral", "parent_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:examples", "path": "documentation/ephemeral-resources/artifact_registry_token/examples/ephemeral/index.md", "provider_name": "artifact_registry_token", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "ephemeral-resources", "role": "example", "schema_path": ["ephemeral"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/artifact_registry_token/examples/ephemeral/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ephemeral for xcsh_artifact_registry_token.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
