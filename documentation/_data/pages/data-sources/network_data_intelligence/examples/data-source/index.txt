---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_data_intelligence."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1393, "body_sha256": "sha256:24e8020091bc4f9206b60fcc55bb7255d93a4e15e7d661948069c317d4da739a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_data_intelligence:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:529d85fb608129abcf053afa1a4c7516767c8567633cbc06efe8d845d135d233", "source_path": "examples/data-sources/xcsh_network_data_intelligence/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_data_intelligence:example:data-source", "parent_id": "xcsh-docs:data-sources:network_data_intelligence:examples", "path": "documentation/data-sources/network_data_intelligence/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_data_intelligence", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2310133133233022-2132112222002311-0131322010333011-3133102122000021-1223111121232020-3223130120320103-2312031313322200-1322132211030230", "registry_path": "docs/guides/data-sources--network_data_intelligence--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_data_intelligence/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_network_data_intelligence.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_data_intelligence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_data_intelligence/data-source.tf`; digest `sha256:529d85fb608129abcf053afa1a4c7516767c8567633cbc06efe8d845d135d233`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_data_intelligence" "us" {
  regions = ["us"]
}

output "data_intelligence_https_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_data_intelligence.us.cidr_blocks
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/examples/)
- [xcsh_network_data_intelligence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/)
