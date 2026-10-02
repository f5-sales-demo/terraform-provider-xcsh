---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1290, "body_sha256": "sha256:4af244722638dbcaf5cbc0f961bcf1dff968e02854647edfbe287f46bc53bc65", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e21d61ff8d692544ea192645c0aec9c60f1857a80969caf8df2d970c69dd061b", "source_path": "examples/data-sources/xcsh_ip_prefix_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ip_prefix_set:example:data-source", "parent_id": "xcsh-docs:data-sources:ip_prefix_set:examples", "path": "documentation/data-sources/ip_prefix_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0122032310030323-2130201101200121-3322131133110023-2131031230011120-3230303011320302-2110201030222033-1313000232102002-2200031021020000", "registry_path": "docs/guides/data-sources--ip_prefix_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_ip_prefix_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/examples/)
- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/)
