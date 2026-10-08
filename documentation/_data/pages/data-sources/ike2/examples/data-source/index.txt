---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike2."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 973, "body_sha256": "sha256:1af09c5374f005e9827b5c35410731079dbed48c6190d940d12f013e8578beb3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:db87600cf054124d776876313565955caf46f37f46627a7f890a5aaf280c673b", "source_path": "examples/data-sources/xcsh_ike2/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike2:example:data-source", "parent_id": "xcsh-docs:data-sources:ike2:examples", "path": "documentation/data-sources/ike2/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0032223302130332-1133100133332220-3111102102021312-3033312130120213-3121021201332330-2333213110213301-0312310200330000-3020013312200213", "registry_path": "docs/guides/data-sources--ike2--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike2/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_ike2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["ike2CreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
