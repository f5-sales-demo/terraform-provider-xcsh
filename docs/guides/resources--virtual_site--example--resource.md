---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_virtual_site."
xcsh_docs: {"aliases": [], "body_bytes": 1042, "body_sha256": "sha256:2274f5d4793e4f1f7fa2332467d7cdce1edd6a00c69780b910661c4c71a1fd6d", "canonical_id": "xcsh-docs:resources:virtual_site:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d4f2ae5a53db544456cc0750c3ea7b5a8e1e23a8aad03279e065c9ce6ac7f872", "source_path": "examples/resources/xcsh_virtual_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:virtual_site:example:resource", "parent_id": "xcsh-docs:resources:virtual_site:examples", "path": "docs/guides/resources--virtual_site--example--resource.md", "provider_name": "virtual_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_virtual_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md)
- [Examples](resources--virtual_site--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_site/resource.tf`; digest `sha256:d4f2ae5a53db544456cc0750c3ea7b5a8e1e23a8aad03279e065c9ce6ac7f872`.

```terraform
# VirtualSite Resource Example
# Manages virtual site object in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualSite configuration
resource "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--virtual_site--examples.md)
- [xcsh_virtual_site](../resources/virtual_site.md)
