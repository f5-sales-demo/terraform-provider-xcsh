---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_service_policy_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1109, "body_sha256": "sha256:e0299abfbc69afc702e8360a8addf87cc98c81207632d730ad9c81c7eae17900", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a", "source_path": "examples/data-sources/xcsh_service_policy_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:service_policy_set:example:data-source", "parent_id": "xcsh-docs:data-sources:service_policy_set:examples", "path": "documentation/data-sources/service_policy_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0330102310230122-1223203021122031-0203023200302031-0313123001221103-2231000301310011-2303023113023020-2021003211310330-1333032122312302", "registry_path": "docs/guides/data-sources--service_policy_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_service_policy_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_service_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_set/data-source.tf`; digest `sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a`.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```
