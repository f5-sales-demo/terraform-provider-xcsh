---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1246, "body_sha256": "sha256:34c90c1272c98eb66f47703cf5ed59aba04162e3d23e720781fb833de4f012ad", "child_ids": [], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:45ebeb1276612423e72f2b472cb13dd3d370a975c3a55fde67381031e1b5970b", "source_path": "examples/resources/xcsh_voltstack_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:voltstack_site:example:resource", "parent_id": "xcsh-docs:resources:voltstack_site:examples", "path": "documentation/resources/voltstack_site/examples/resource/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_voltstack_site/resource.tf`; digest `sha256:45ebeb1276612423e72f2b472cb13dd3d370a975c3a55fde67381031e1b5970b`.

```terraform
# VoltstackSite Resource Example
# Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VoltstackSite configuration
resource "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/examples/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
