---
page_title: "xcsh_kubernetes_manifests"
subcategory: ""
description: "Kubernetes workload configuration."
xcsh_docs: {"aliases": ["kubernetes manifests"], "body_bytes": 1313, "body_sha256": "sha256:d69751e4878b1b0b6575a012a19bc1079d7e589cb0dfa2ebe54346a55d124815", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "xcsh-docs:ephemeral-resources:kubernetes_manifests:examples", "xcsh-docs:ephemeral-resources:kubernetes_manifests:lifecycle"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:fundamentals", "parent_id": "xcsh-docs:ephemeral-resources:xcsh:navigation", "path": "documentation/ephemeral-resources/kubernetes_manifests/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-3020030201123201-1030130333132003-2002102001021330-1320220001031302-0001300203331022-2100320312131223-1211321010102033-2012323132111020", "registry_path": "docs/ephemeral-resources/kubernetes_manifests.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Kubernetes workload configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/lifecycle/)
