---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 996, "body_sha256": "sha256:874ac34733a1726041b0691bb48a0c423412a15679534b9627682fc742398615", "canonical_id": "xcsh-docs:resources:api_discovery:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ff9af826f29fe467445464e3b92dcd212f9666dfb00feabd76aab3e247a3f9e6", "source_path": "examples/resources/xcsh_api_discovery/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_discovery:example:resource", "parent_id": "xcsh-docs:resources:api_discovery:examples", "path": "docs/guides/resources--api_discovery--example--resource.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_api_discovery](../resources/api_discovery.md)
- [Examples](resources--api_discovery--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_discovery/resource.tf`; digest `sha256:ff9af826f29fe467445464e3b92dcd212f9666dfb00feabd76aab3e247a3f9e6`.

```terraform
# APIDiscovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDiscovery configuration
resource "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--api_discovery--examples.md)
- [xcsh_api_discovery](../resources/api_discovery.md)
