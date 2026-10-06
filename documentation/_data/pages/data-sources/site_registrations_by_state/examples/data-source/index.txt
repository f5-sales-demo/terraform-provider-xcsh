---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1086, "body_sha256": "sha256:643ce2378e844b90310dc9853ec4e27ba40be146410c075c0bd1421804962d02", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:87878119dc98d70aee1fa7c3e546432844de313704e34499df4b6dfdc776ce5a", "source_path": "examples/data-sources/xcsh_site_registrations_by_state/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_registrations_by_state:example:data-source", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:examples", "path": "documentation/data-sources/site_registrations_by_state/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2122023233332330-2113130130003332-1301333213013322-3220013132032113-0310120123300322-0211102000110123-3123233230233132-1121301001023212", "registry_path": "docs/guides/data-sources--site_registrations_by_state--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_site_registrations_by_state.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations_by_state/data-source.tf`; digest `sha256:87878119dc98d70aee1fa7c3e546432844de313704e34499df4b6dfdc776ce5a`.

```terraform
# SiteRegistrationsByState DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_state" "example" {
  state = "NOTSET"
}

output "site_registrations_by_state_result" {
  value = data.xcsh_site_registrations_by_state.example
}
```
