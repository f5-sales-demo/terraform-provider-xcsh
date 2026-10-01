---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1315, "body_sha256": "sha256:109bf4b6a69c5661072726ca84b9571783e558299a7530b03dec4949f46ae7d0", "child_ids": [], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f", "source_path": "examples/resources/xcsh_application_profiles/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:application_profiles:example:resource", "parent_id": "xcsh-docs:resources:application_profiles:examples", "path": "documentation/resources/application_profiles/examples/resource/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_application_profiles/resource.tf`; digest `sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f`.

```terraform
# ApplicationProfiles Resource Example
# Manages Application Profiles in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ApplicationProfiles configuration
resource "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/examples/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
