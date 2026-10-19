---
page_title: "items.get_spec.infra.sw_info"
subcategory: ""
description: "SWInfo holds information about sw version."
xcsh_docs: {"aliases": ["items get spec infra sw info"], "body_bytes": 1090, "body_sha256": "sha256:44728612b39d31138a313b60d1023235c2c59b87ff2f19208dfa9103e127c0f4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:sw_info", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/sw_info/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1100211221221200-0331232303300203-2020031013330012-1003231120120103-3022211333212220-1110031022102333-2312113312002102-1323111332201303", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "sw_info"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra sw info sw version"], "anchor": "schema-items--get_spec--infra--sw_info--sw_version", "description": "SW Version. SW Version in the site.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:sw_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "sw_info", "sw_version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/sw_info/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "SWInfo holds information about sw version.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": [], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.sw_info

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- items.get_spec.infra.sw_info

<a id="section"></a>

Type: `"single"`. Computed.

SWInfo holds information about sw version.

## Direct properties

<a id="schema-items--get_spec--infra--sw_info--sw_version"></a>

### sw_version property

Type: `"string"`. Computed.

SW Version. SW Version in the site.
