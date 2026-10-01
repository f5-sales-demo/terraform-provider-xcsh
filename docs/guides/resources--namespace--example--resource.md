---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_namespace."
xcsh_docs: {"aliases": [], "body_bytes": 962, "body_sha256": "sha256:e55c5c4897b9f8d11233fb15046352a34f3280d69f799feb407dee7f946c29b2", "canonical_id": "xcsh-docs:resources:namespace:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:namespace:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d0fc0c13acd246e58f1623298d7e64f7fe481e01837ebb788df9a83a8e965566", "source_path": "examples/resources/xcsh_namespace/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:namespace:example:resource", "parent_id": "xcsh-docs:resources:namespace:examples", "path": "docs/guides/resources--namespace--example--resource.md", "provider_name": "namespace", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/namespace/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_namespace.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["namespaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md)
- [Examples](resources--namespace--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_namespace/resource.tf`; digest `sha256:d0fc0c13acd246e58f1623298d7e64f7fe481e01837ebb788df9a83a8e965566`.

```terraform
# Namespace Resource Example
# Manages new namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Namespace configuration
resource "xcsh_namespace" "example" {
  name      = "example-namespace"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--namespace--examples.md)
- [xcsh_namespace](../resources/namespace.md)
