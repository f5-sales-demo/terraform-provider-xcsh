---
page_title: "xcsh_artifact_registry_token"
subcategory: ""
description: "Authentication credential for access control."
xcsh_docs: {"aliases": ["artifact registry token"], "body_bytes": 1367, "body_sha256": "sha256:76a1d147a17315b93a69e5d043ecdbb2abd2026431dec15fe16a2dd6e63a1a52", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:ephemeral-resources:artifact_registry_token:reference", "xcsh-docs:ephemeral-resources:artifact_registry_token:examples", "xcsh-docs:ephemeral-resources:artifact_registry_token:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:artifact_registry_token:fundamentals", "parent_id": "xcsh-docs:ephemeral-resources:xcsh:navigation", "path": "documentation/ephemeral-resources/artifact_registry_token/index.md", "product": "distributed-cloud", "provider_name": "artifact_registry_token", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-1323301000023223-3012100013003333-1031230331233031-2332031031030022-0033001320123233-0303101023323201-3331031322333331-2330103332212301", "registry_path": "docs/ephemeral-resources/artifact_registry_token.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/artifact_registry_token/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Authentication credential for access control.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/lifecycle/)
