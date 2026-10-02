---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cminstance."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1255, "body_sha256": "sha256:eacef67085860bda96847e14efd433dfc2f5ab01039589930819108b7fb0ab17", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:833c4b202ec741b02e9ef290bb6512a37d3156dd48242949d7ad995d973193c7", "source_path": "examples/data-sources/xcsh_cminstance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cminstance:example:data-source", "parent_id": "xcsh-docs:data-sources:cminstance:examples", "path": "documentation/data-sources/cminstance/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1303001130110232-3311203011112302-0331320312021002-2222332203313102-1300223231331331-2122011210121101-2121012030322331-2030110303322011", "registry_path": "docs/guides/data-sources--cminstance--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_cminstance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cminstanceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cminstance/data-source.tf`; digest `sha256:833c4b202ec741b02e9ef290bb6512a37d3156dd48242949d7ad995d973193c7`.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/examples/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
