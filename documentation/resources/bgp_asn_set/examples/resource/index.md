---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bgp_asn_set."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1295, "body_sha256": "sha256:eda4792ec39a4fb98013a64e40fdf0af7e473b2dd8542cc2b0347559a7e3e05b", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_asn_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c2bb9b5fa3b8555e63a99053456d35fa40e857a94daa94de302421615af065a", "source_path": "examples/resources/xcsh_bgp_asn_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bgp_asn_set:example:resource", "parent_id": "xcsh-docs:resources:bgp_asn_set:examples", "path": "documentation/resources/bgp_asn_set/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bgp_asn_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1220303231033322-1300223213110123-0311320313311132-0000220100013200-3111113020032331-3200100303032121-2202012123302023-0020113113101020", "registry_path": "docs/guides/resources--bgp_asn_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_asn_set/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_bgp_asn_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_asn_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_bgp_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp_asn_set/resource.tf`; digest `sha256:3c2bb9b5fa3b8555e63a99053456d35fa40e857a94daa94de302421615af065a`.

```terraform
# BGPAsnSet Resource Example
# Manages bgp_asn_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPAsnSet configuration
resource "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"

  as_numbers = [1]
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/examples/)
- [xcsh_bgp_asn_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_asn_set/)
