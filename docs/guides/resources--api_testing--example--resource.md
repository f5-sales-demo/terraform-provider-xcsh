---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 916, "body_sha256": "sha256:96ac93693cd112f4157158560e4f10c24286da1d2dd626a746089ac89d8fed7b", "canonical_id": "xcsh-docs:resources:api_testing:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c1227be913c1052c8bc552ce9ebe6740f44bc74215b29a565808e69bd8cea2fb", "source_path": "examples/resources/xcsh_api_testing/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_testing:example:resource", "parent_id": "xcsh-docs:resources:api_testing:examples", "path": "docs/guides/resources--api_testing--example--resource.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md)
- [Examples](resources--api_testing--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_testing/resource.tf`; digest `sha256:c1227be913c1052c8bc552ce9ebe6740f44bc74215b29a565808e69bd8cea2fb`.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--api_testing--examples.md)
- [xcsh_api_testing](../resources/api_testing.md)
