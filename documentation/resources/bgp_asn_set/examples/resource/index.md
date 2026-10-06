---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bgp_asn_set."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1076, "body_sha256": "sha256:b16b0226f7a3ed530f538ae2dca88f1b208d2b4e7e62692a188882c9f0bc822b", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_asn_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c2bb9b5fa3b8555e63a99053456d35fa40e857a94daa94de302421615af065a", "source_path": "examples/resources/xcsh_bgp_asn_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bgp_asn_set:example:resource", "parent_id": "xcsh-docs:resources:bgp_asn_set:examples", "path": "documentation/resources/bgp_asn_set/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bgp_asn_set", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1220303231033322-1300223213110123-0311320313311132-0000220100013200-3111113020032331-3200100303032121-2202012123302023-0020113113101020", "registry_path": "docs/guides/resources--bgp_asn_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_asn_set/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_bgp_asn_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bgp_asn_setCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
