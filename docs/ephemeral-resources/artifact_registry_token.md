---
page_title: "xcsh_artifact_registry_token"
subcategory: ""
description: "xcsh_artifact_registry_token for xcsh_artifact_registry_token."
xcsh_docs: {"aliases": [], "body_bytes": 1227, "body_sha256": "sha256:ce14176b63fc9552205766285c13bd98d56b2bab93ce051eee4f15242bd9e9a8", "canonical_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:fundamentals", "child_ids": ["xcsh-docs:ephemeral-resources:artifact_registry_token:reference", "xcsh-docs:ephemeral-resources:artifact_registry_token:examples", "xcsh-docs:ephemeral-resources:artifact_registry_token:lifecycle"], "collection_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:artifact_registry_token:fundamentals", "parent_id": null, "path": "docs/ephemeral-resources/artifact_registry_token.md", "provider_name": "artifact_registry_token", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "ephemeral-resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/artifact_registry_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_artifact_registry_token for xcsh_artifact_registry_token.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_artifact_registry_token

Breadcrumbs:

- xcsh_artifact_registry_token

Authentication credential for access control.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/ephemeral-resources--artifact_registry_token--reference.md)
- [Examples](../guides/ephemeral-resources--artifact_registry_token--examples.md)
- [Lifecycle](../guides/ephemeral-resources--artifact_registry_token--lifecycle.md)
