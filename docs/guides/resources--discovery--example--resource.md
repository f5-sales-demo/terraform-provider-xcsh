---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1075, "body_sha256": "sha256:49a854a606567bfc6bd04d38aa52b5e3e1b26185f4be697bdd49da049d7becea", "canonical_id": "xcsh-docs:resources:discovery:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c", "source_path": "examples/resources/xcsh_discovery/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:discovery:example:resource", "parent_id": "xcsh-docs:resources:discovery:examples", "path": "docs/guides/resources--discovery--example--resource.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Examples](resources--discovery--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_discovery/resource.tf`; digest `sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c`.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--discovery--examples.md)
- [xcsh_discovery](../resources/discovery.md)
