---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 1299, "body_sha256": "sha256:230476ba7de896d1887c94207f46f56e4e63ca13024e596712d4e442a1e63f89", "child_ids": [], "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f4976402381de0085eaa3ddefc5b729e1a883d869e4792bc2f8f960a6f9b6252", "source_path": "examples/resources/xcsh_app_api_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_api_group:example:resource", "parent_id": "xcsh-docs:resources:app_api_group:examples", "path": "documentation/resources/app_api_group/examples/resource/index.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_api_group/resource.tf`; digest `sha256:f4976402381de0085eaa3ddefc5b729e1a883d869e4792bc2f8f960a6f9b6252`.

```terraform
# AppAPIGroup Resource Example
# Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppAPIGroup configuration
resource "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/examples/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
