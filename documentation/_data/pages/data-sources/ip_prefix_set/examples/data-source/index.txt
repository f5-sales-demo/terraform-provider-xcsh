---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1059, "body_sha256": "sha256:4e52d61775fe34daf0892a8e6dc97c7c50bfc82411ebf4baa2f5cd4d60ca2dd4", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e21d61ff8d692544ea192645c0aec9c60f1857a80969caf8df2d970c69dd061b", "source_path": "examples/data-sources/xcsh_ip_prefix_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ip_prefix_set:example:data-source", "parent_id": "xcsh-docs:data-sources:ip_prefix_set:examples", "path": "documentation/data-sources/ip_prefix_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0122032310030323-2130201101200121-3322131133110023-2131031230011120-3230303011320302-2110201030222033-1313000232102002-2200031021020000", "registry_path": "docs/guides/data-sources--ip_prefix_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_ip_prefix_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ip_prefix_set/data-source.tf`; digest `sha256:e21d61ff8d692544ea192645c0aec9c60f1857a80969caf8df2d970c69dd061b`.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```
