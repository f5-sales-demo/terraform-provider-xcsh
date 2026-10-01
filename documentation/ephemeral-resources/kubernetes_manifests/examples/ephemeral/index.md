---
page_title: "Ephemeral"
subcategory: ""
description: "Ephemeral for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": [], "body_bytes": 1240, "body_sha256": "sha256:3de2fe482f71a88fe05964c4860c6a8360312fd8147895c5fbe28c65dd0c0ac5", "child_ids": [], "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e7611464c5404b77a4b6c720f8e02924f4419518332607f8e54efc9f31aad54e", "source_path": "examples/ephemeral-resources/xcsh_kubernetes_manifests/ephemeral.tf", "validation": "terraform validate"}, "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:example:ephemeral", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:examples", "path": "documentation/ephemeral-resources/kubernetes_manifests/examples/ephemeral/index.md", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "ephemeral-resources", "role": "example", "schema_path": ["ephemeral"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/examples/ephemeral/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ephemeral for xcsh_kubernetes_manifests.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/examples/)
- [xcsh_kubernetes_manifests](https://f5-sales-demo.github.io/terraform-provider-xcsh/ephemeral-resources/kubernetes_manifests/)
