---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1299, "body_sha256": "sha256:515614b291f358006ce200040bed012b29af9ecb178b32b04445c3972297527e", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ip_prefix_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:00f8d9b11740968b5c7c6ae54585108409034d987dc62393dd9ee54d031d3c3a", "source_path": "examples/resources/xcsh_ip_prefix_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ip_prefix_set:example:resource", "parent_id": "xcsh-docs:resources:ip_prefix_set:examples", "path": "documentation/resources/ip_prefix_set/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0333101020302003-3203020032220321-3020100100121231-2031131302133301-0230031323032111-1222101000220332-1103010120331311-3130013313131323", "registry_path": "docs/guides/resources--ip_prefix_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ip_prefix_set/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_ip_prefix_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ip_prefix_set/resource.tf`; digest `sha256:00f8d9b11740968b5c7c6ae54585108409034d987dc62393dd9ee54d031d3c3a`.

```terraform
# IPPrefixSet Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IPPrefixSet configuration
resource "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/examples/)
- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/)
