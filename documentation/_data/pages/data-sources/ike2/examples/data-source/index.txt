---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike2."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1177, "body_sha256": "sha256:51efd7f20a2cd2de985082b7053951949bf3fe95742d537331862a8b899704b9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:db87600cf054124d776876313565955caf46f37f46627a7f890a5aaf280c673b", "source_path": "examples/data-sources/xcsh_ike2/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike2:example:data-source", "parent_id": "xcsh-docs:data-sources:ike2:examples", "path": "documentation/data-sources/ike2/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0032223302130332-1133100133332220-3111102102021312-3033312130120213-3121021201332330-2333213110213301-0312310200330000-3020013312200213", "registry_path": "docs/guides/data-sources--ike2--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike2/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_ike2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike2/data-source.tf`; digest `sha256:db87600cf054124d776876313565955caf46f37f46627a7f890a5aaf280c673b`.

```terraform
# Ike2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike2 by name
data "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}

output "ike2_id" {
  value = data.xcsh_ike2.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/examples/)
- [xcsh_ike2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/)
