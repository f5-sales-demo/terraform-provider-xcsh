---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_addon_service."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1292, "body_sha256": "sha256:67b20a7a65fbedfb7f384d457bd9a3fe1fc8fdec25c70626d021770c1a55808a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:addon_service:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:809191858b0cd3ac6042ff132999dd4d2b0f0773a32efcde70b3a2eff7c8fa98", "source_path": "examples/data-sources/xcsh_addon_service/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:addon_service:example:data-source", "parent_id": "xcsh-docs:data-sources:addon_service:examples", "path": "documentation/data-sources/addon_service/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "addon_service", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2131011112100033-2103211030113023-0312103132202320-3213200310222301-1103122223130032-0301321223022231-0011221023322310-3333233023013213", "registry_path": "docs/guides/data-sources--addon_service--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_addon_service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_addon_service/data-source.tf`; digest `sha256:809191858b0cd3ac6042ff132999dd4d2b0f0773a32efcde70b3a2eff7c8fa98`.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/examples/)
- [xcsh_addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/)
