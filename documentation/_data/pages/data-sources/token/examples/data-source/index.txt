---
page_title: "Data source"
subcategory: "Identity"
description: "Data source for xcsh_token."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1189, "body_sha256": "sha256:fb15736b3648a787218770d43042c1f57046470d03dfb1963079d2008b3ab59b", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:token:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2be846025946a447e16fd9c1be64b48387d966481fa919712ec2c2311d266aee", "source_path": "examples/data-sources/xcsh_token/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:token:example:data-source", "parent_id": "xcsh-docs:data-sources:token:examples", "path": "documentation/data-sources/token/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0102202302032300-2131201001103020-1213132120331221-0031010123200100-0120100213311021-1021133231120011-0320202323201223-2323111231111021", "registry_path": "docs/guides/data-sources--token--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/token/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_token.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tokenCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_token/data-source.tf`; digest `sha256:2be846025946a447e16fd9c1be64b48387d966481fa919712ec2c2311d266aee`.

```terraform
# Token Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Token by name
data "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
}

output "token_id" {
  value = data.xcsh_token.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/examples/)
- [xcsh_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/)
