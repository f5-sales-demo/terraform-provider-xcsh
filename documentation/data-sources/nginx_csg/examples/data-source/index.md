---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_csg."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1240, "body_sha256": "sha256:7207f03bab9caabcf9c8a40b4c292d692451ce663bc5494f360c84e454cd5208", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c9552491d1b728d4d6e84172cf532690c2cae1b286159e0a14dfd5d2f7845eb3", "source_path": "examples/data-sources/xcsh_nginx_csg/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_csg:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_csg:examples", "path": "documentation/data-sources/nginx_csg/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2203211221210023-1213132030332300-3003201110311231-1031011303202200-2310313320110123-2301111023003123-2220011200302100-0113211322232220", "registry_path": "docs/guides/data-sources--nginx_csg--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_nginx_csg.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/examples/)
- [xcsh_nginx_csg](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/)
