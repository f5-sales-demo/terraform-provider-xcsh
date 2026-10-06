---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1129, "body_sha256": "sha256:b20833a77e570608c00e52e547a071f427f8ced46d7c7b358c33090933e82ceb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b98a932b2c1a0b5e2b9c608b25b6b3f74aea2df606bdee9675da50e166ce11d", "source_path": "examples/data-sources/xcsh_bigip_virtual_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bigip_virtual_server:example:data-source", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:examples", "path": "documentation/data-sources/bigip_virtual_server/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-2301101033301122-1331212031100101-2131221013222020-1311301011213222-1011211133313023-0022122010200311-3222111102330012-0122213121111022", "registry_path": "docs/guides/data-sources--bigip_virtual_server--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_bigip_virtual_server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bigip_virtual_server/data-source.tf`; digest `sha256:4b98a932b2c1a0b5e2b9c608b25b6b3f74aea2df606bdee9675da50e166ce11d`.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```
