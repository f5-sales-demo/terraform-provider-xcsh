---
page_title: "xcsh_kubernetes_manifests"
subcategory: ""
description: "Kubernetes workload configuration."
xcsh_docs: {"aliases": ["kubernetes manifests"], "body_bytes": 1326, "body_sha256": "sha256:e06693ae7a8859caa9f4520d9d9c797fd014556041140cf9890e21b5d2781618", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "xcsh-docs:ephemeral-resources:kubernetes_manifests:examples", "xcsh-docs:ephemeral-resources:kubernetes_manifests:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:fundamentals", "parent_id": "xcsh-docs:ephemeral-resources:xcsh:navigation", "path": "documentation/ephemeral-resources/kubernetes_manifests/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-3020030201123201-1030130333132003-2002102001021330-1320220001031302-0001300203331022-2100320312131223-1211321010102033-2012323132111020", "registry_path": "docs/ephemeral-resources/kubernetes_manifests.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Kubernetes workload configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_kubernetes_manifests

Breadcrumbs:

- xcsh_kubernetes_manifests

Kubernetes workload configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `site`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/lifecycle/)
