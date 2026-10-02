---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bgp_asn_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1264, "body_sha256": "sha256:ff31f94ada2d5bea36123c8f05770bff905590caf5a0d299a0845f53fff38273", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_asn_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:101ddc7bff960e99df5f4c481b2d71baf9042bf80e41813366a6137bff577b85", "source_path": "examples/data-sources/xcsh_bgp_asn_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bgp_asn_set:example:data-source", "parent_id": "xcsh-docs:data-sources:bgp_asn_set:examples", "path": "documentation/data-sources/bgp_asn_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bgp_asn_set", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0133022332032332-3102110222003332-3302133133030221-1021010001001311-2013002012320103-0132011011131131-3131232011011312-1203110023113022", "registry_path": "docs/guides/data-sources--bgp_asn_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_asn_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bgp_asn_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_asn_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bgp_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp_asn_set/data-source.tf`; digest `sha256:101ddc7bff960e99df5f4c481b2d71baf9042bf80e41813366a6137bff577b85`.

```terraform
# BGPAsnSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPAsnSet by name
data "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"
}

output "bgp_asn_set_id" {
  value = data.xcsh_bgp_asn_set.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/examples/)
- [xcsh_bgp_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/)
