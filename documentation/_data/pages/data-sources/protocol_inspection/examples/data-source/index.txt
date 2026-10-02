---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protocol_inspection."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1370, "body_sha256": "sha256:73ed7b020f376d8347f4099a9b9da1ca1e7ecd1976f9db1f401635af361f72fe", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8e357f2e82e2e660e9b2335aec0a0bf6fb5efafdd6746da23a0a982fc6492237", "source_path": "examples/data-sources/xcsh_protocol_inspection/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protocol_inspection:example:data-source", "parent_id": "xcsh-docs:data-sources:protocol_inspection:examples", "path": "documentation/data-sources/protocol_inspection/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2321033120121020-3101302133330101-2131210100112013-1223322021203030-3211213013202023-3112213231210012-2132113032320231-1203012130102321", "registry_path": "docs/guides/data-sources--protocol_inspection--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_protocol_inspection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protocol_inspection/data-source.tf`; digest `sha256:8e357f2e82e2e660e9b2335aec0a0bf6fb5efafdd6746da23a0a982fc6492237`.

```terraform
# ProtocolInspection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolInspection by name
data "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}

output "protocol_inspection_id" {
  value = data.xcsh_protocol_inspection.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/examples/)
- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
