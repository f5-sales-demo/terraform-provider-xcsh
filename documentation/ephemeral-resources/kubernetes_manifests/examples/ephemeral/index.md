---
page_title: "Ephemeral"
subcategory: ""
description: "Ephemeral for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": ["ephemeral"], "body_bytes": 974, "body_sha256": "sha256:9124ba91b04491bd8b2f7a377d72106619a8b362ea966b7dcf6eea525b213f61", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e7611464c5404b77a4b6c720f8e02924f4419518332607f8e54efc9f31aad54e", "source_path": "examples/ephemeral-resources/xcsh_kubernetes_manifests/ephemeral.tf", "validation": "terraform validate"}, "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:example:ephemeral", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:examples", "path": "documentation/ephemeral-resources/kubernetes_manifests/examples/ephemeral/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-1222000303323313-0032220310022301-0211213320021023-0202310230003111-1122012310302201-0133201000110312-1123110111231202-3312023110332013", "registry_path": "docs/guides/ephemeral-resources--kubernetes_manifests--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["ephemeral"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/examples/ephemeral/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Ephemeral for xcsh_kubernetes_manifests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Ephemeral

Breadcrumbs:

- [xcsh_kubernetes_manifests](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/examples/)
- Ephemeral

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/ephemeral-resources/xcsh_kubernetes_manifests/ephemeral.tf`; digest `sha256:e7611464c5404b77a4b6c720f8e02924f4419518332607f8e54efc9f31aad54e`.

```terraform
# KubernetesManifests EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_kubernetes_manifests" "example" {
  site = "example-value"
}
```
