---
page_title: "xcsh_kubernetes_manifests"
subcategory: ""
description: "xcsh_kubernetes_manifests for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": [], "body_bytes": 1087, "body_sha256": "sha256:6cb2ac4baa83e1f2d426309259c0e7471b8c11efe358fd7be2d53c3c3fadf50d", "canonical_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:fundamentals", "child_ids": ["xcsh-docs:ephemeral-resources:kubernetes_manifests:reference", "xcsh-docs:ephemeral-resources:kubernetes_manifests:examples", "xcsh-docs:ephemeral-resources:kubernetes_manifests:lifecycle"], "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:fundamentals", "parent_id": null, "path": "docs/ephemeral-resources/kubernetes_manifests.md", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "ephemeral-resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_kubernetes_manifests for xcsh_kubernetes_manifests.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](../guides/ephemeral-resources--kubernetes_manifests--reference.md)
- [Examples](../guides/ephemeral-resources--kubernetes_manifests--examples.md)
- [Lifecycle](../guides/ephemeral-resources--kubernetes_manifests--lifecycle.md)
