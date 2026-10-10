---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_csg."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1021, "body_sha256": "sha256:e755df93c5955456abd4d525f83a4292c67fb54c067ef423a6369ff8dc62b349", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c9552491d1b728d4d6e84172cf532690c2cae1b286159e0a14dfd5d2f7845eb3", "source_path": "examples/data-sources/xcsh_nginx_csg/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_csg:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_csg:examples", "path": "documentation/data-sources/nginx_csg/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2203211221210023-1213132030332300-3003201110311231-1031011303202200-2310313320110123-2301111023003123-2220011200302100-0113211322232220", "registry_path": "docs/guides/data-sources--nginx_csg--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_nginx_csg.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nginx_csg](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_csg/data-source.tf`; digest `sha256:c9552491d1b728d4d6e84172cf532690c2cae1b286159e0a14dfd5d2f7845eb3`.

```terraform
# NginxCsg Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxCsg by name
data "xcsh_nginx_csg" "example" {
  name      = "example-nginx-csg"
  namespace = "staging"
}

output "nginx_csg_id" {
  value = data.xcsh_nginx_csg.example.id
}
```
