---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_voltstack_site."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1345, "body_sha256": "sha256:43c9feb84b6c575a21b67fceb4e82e1a35f72e053ba65df9bbaaa528aaf408ec", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:45ebeb1276612423e72f2b472cb13dd3d370a975c3a55fde67381031e1b5970b", "source_path": "examples/resources/xcsh_voltstack_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:voltstack_site:example:resource", "parent_id": "xcsh-docs:resources:voltstack_site:examples", "path": "documentation/resources/voltstack_site/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1020000130123312-3102130130230323-0321020002112101-3223231102313300-0220213103003221-2033233100300223-2233210001213310-3202031130121302", "registry_path": "docs/guides/resources--voltstack_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_voltstack_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
