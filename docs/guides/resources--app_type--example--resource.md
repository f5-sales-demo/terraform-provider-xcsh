---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_type."
xcsh_docs: {"aliases": [], "body_bytes": 1030, "body_sha256": "sha256:0fd999f124d67edd1c3823546ca073d38ab0f4968a0f684ad5688334538b8ba7", "canonical_id": "xcsh-docs:resources:app_type:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9988ab36f101ce764d3f26103a86f6c75f3dcea74d272aeebe958208da89b71d", "source_path": "examples/resources/xcsh_app_type/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_type:example:resource", "parent_id": "xcsh-docs:resources:app_type:examples", "path": "docs/guides/resources--app_type--example--resource.md", "provider_name": "app_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_app_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md)
- [Examples](resources--app_type--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_type/resource.tf`; digest `sha256:9988ab36f101ce764d3f26103a86f6c75f3dcea74d272aeebe958208da89b71d`.

```terraform
# AppType Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppType configuration
resource "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--app_type--examples.md)
- [xcsh_app_type](../resources/app_type.md)
