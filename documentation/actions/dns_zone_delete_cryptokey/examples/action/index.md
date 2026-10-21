---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_dns_zone_delete_cryptokey."
xcsh_docs: {"aliases": ["action"], "body_bytes": 933, "body_sha256": "sha256:b8215cf05111d9116a6b8081d4ecf2592d9c9af97bf62548cfaf5d4d812fd00c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:41efc29a59313c81b918dfd8dc2bfaca42cc84009a5aa5450b2c6881b14ad7fe", "source_path": "examples/actions/xcsh_dns_zone_delete_cryptokey/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:dns_zone_delete_cryptokey:example:action", "parent_id": "xcsh-docs:actions:dns_zone_delete_cryptokey:examples", "path": "documentation/actions/dns_zone_delete_cryptokey/examples/action/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_delete_cryptokey", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "actions", "registry_anchor": "canonical-0103002301110103-1303330031001123-0201200102011303-1321300201133303-0010201131221020-2001101020103120-3232303103223130-0011331032321000", "registry_path": "docs/guides/actions--dns_zone_delete_cryptokey--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_delete_cryptokey/examples/action/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Action for xcsh_dns_zone_delete_cryptokey.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": [], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_dns_zone_delete_cryptokey](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_delete_cryptokey/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_dns_zone_delete_cryptokey/action.tf`; digest `sha256:41efc29a59313c81b918dfd8dc2bfaca42cc84009a5aa5450b2c6881b14ad7fe`.

```terraform
# DNSZoneDeleteCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_delete_cryptokey" "example" {
  config {
  }
}
```
