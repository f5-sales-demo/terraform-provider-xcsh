---
page_title: "Resource"
subcategory: "Networking"
description: "Resource for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 1049, "body_sha256": "sha256:7d779f6ec8ed40e373444e76468f5628c3e30d02d215a58ac04fa16c08c971c3", "canonical_id": "xcsh-docs:resources:endpoint:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e", "source_path": "examples/resources/xcsh_endpoint/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:endpoint:example:resource", "parent_id": "xcsh-docs:resources:endpoint:examples", "path": "docs/guides/resources--endpoint--example--resource.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md)
- [Examples](resources--endpoint--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_endpoint/resource.tf`; digest `sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e`.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--endpoint--examples.md)
- [xcsh_endpoint](../resources/endpoint.md)
