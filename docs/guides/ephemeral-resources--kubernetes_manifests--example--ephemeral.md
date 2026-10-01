---
page_title: "Ephemeral"
subcategory: ""
description: "Ephemeral for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": [], "body_bytes": 1034, "body_sha256": "sha256:59f082ad516aa72f0a9853614d1b31c35bad26247ea4784f0571b58dc3eab737", "canonical_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:example:ephemeral", "child_ids": [], "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e7611464c5404b77a4b6c720f8e02924f4419518332607f8e54efc9f31aad54e", "source_path": "examples/ephemeral-resources/xcsh_kubernetes_manifests/ephemeral.tf", "validation": "terraform validate"}, "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:example:ephemeral", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:examples", "path": "docs/guides/ephemeral-resources--kubernetes_manifests--example--ephemeral.md", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "ephemeral-resources", "publishing_destination": "registry", "role": "example", "schema_path": ["ephemeral"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/examples/ephemeral/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ephemeral for xcsh_kubernetes_manifests.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Ephemeral

Breadcrumbs:

- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md)
- [Examples](ephemeral-resources--kubernetes_manifests--examples.md)
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

## Next pages

- [Examples](ephemeral-resources--kubernetes_manifests--examples.md)
- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md)
