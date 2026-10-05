---
page_title: "xcsh_artifact_registry_token"
subcategory: ""
description: "Authentication credential for access control."
xcsh_docs: {"aliases": ["artifact registry token"], "body_bytes": 1354, "body_sha256": "sha256:bd31a9c9f89b5a30de03dc24cf0a82521e8bbba4a126cedce7a57969b5230fd9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:ephemeral-resources:artifact_registry_token:reference", "xcsh-docs:ephemeral-resources:artifact_registry_token:examples", "xcsh-docs:ephemeral-resources:artifact_registry_token:lifecycle"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:artifact_registry_token:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:artifact_registry_token:fundamentals", "parent_id": "xcsh-docs:ephemeral-resources:xcsh:navigation", "path": "documentation/ephemeral-resources/artifact_registry_token/index.md", "product": "distributed-cloud", "provider_name": "artifact_registry_token", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-1323301000023223-3012100013003333-1031230331233031-2332031031030022-0033001320123233-0303101023323201-3331031322333331-2330103332212301", "registry_path": "docs/ephemeral-resources/artifact_registry_token.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/artifact_registry_token/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Authentication credential for access control.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/artifact_registry_token/lifecycle/)
