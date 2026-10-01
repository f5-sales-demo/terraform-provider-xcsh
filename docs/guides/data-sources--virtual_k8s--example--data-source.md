---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_virtual_k8s."
xcsh_docs: {"aliases": [], "body_bytes": 1060, "body_sha256": "sha256:03223de471fa8bbaf32afd21ab358d96aecf712ba556c51bb2c31f73efca6ac4", "canonical_id": "xcsh-docs:data-sources:virtual_k8s:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:85c7894f9985a87a12ec00003bcdb34e2727cf22575e06597a586ec7fcc0019d", "source_path": "examples/data-sources/xcsh_virtual_k8s/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_k8s:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_k8s:examples", "path": "docs/guides/data-sources--virtual_k8s--example--data-source.md", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_virtual_k8s.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md)
- [Examples](data-sources--virtual_k8s--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_k8s/data-source.tf`; digest `sha256:85c7894f9985a87a12ec00003bcdb34e2727cf22575e06597a586ec7fcc0019d`.

```terraform
# VirtualK8S Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualK8S by name
data "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}

output "virtual_k8s_id" {
  value = data.xcsh_virtual_k8s.example.id
}
```

## Next pages

- [Examples](data-sources--virtual_k8s--examples.md)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md)
