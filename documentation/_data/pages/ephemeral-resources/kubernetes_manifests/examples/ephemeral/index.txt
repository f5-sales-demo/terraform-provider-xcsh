---
page_title: "Ephemeral"
subcategory: ""
description: "Ephemeral for xcsh_kubernetes_manifests."
xcsh_docs: {"aliases": ["ephemeral"], "body_bytes": 1240, "body_sha256": "sha256:3de2fe482f71a88fe05964c4860c6a8360312fd8147895c5fbe28c65dd0c0ac5", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e7611464c5404b77a4b6c720f8e02924f4419518332607f8e54efc9f31aad54e", "source_path": "examples/ephemeral-resources/xcsh_kubernetes_manifests/ephemeral.tf", "validation": "terraform validate"}, "id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:example:ephemeral", "parent_id": "xcsh-docs:ephemeral-resources:kubernetes_manifests:examples", "path": "documentation/ephemeral-resources/kubernetes_manifests/examples/ephemeral/index.md", "product": "distributed-cloud", "provider_name": "kubernetes_manifests", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "ephemeral-resources", "registry_anchor": "canonical-1222000303323313-0032220310022301-0211213320021023-0202310230003111-1122012310302201-0133201000110312-1123110111231202-3312023110332013", "registry_path": "docs/guides/ephemeral-resources--kubernetes_manifests--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["ephemeral"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/ephemeral-resources/kubernetes_manifests/examples/ephemeral/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Ephemeral for xcsh_kubernetes_manifests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
