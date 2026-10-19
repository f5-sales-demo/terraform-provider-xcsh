---
page_title: "rules.action.dynamic_pool.pools"
subcategory: ""
description: "List of IPv4 prefixes that represent an endpoint."
xcsh_docs: {"aliases": ["rules action dynamic pool pools"], "body_bytes": 1084, "body_sha256": "sha256:049be81095864f05809d52c11cb502a4448c7e3e7b7e6ece6d5b63d096e97969", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:pools", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool", "path": "documentation/data-sources/nat_policy/properties/rules/action/dynamic_pool/pools/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1112003013333113-0132301012130010-1213131333321323-3232323211301230-1133113232120212-1202230232131223-1101020321033333-1112120230103023", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic_pool", "pools"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic pool pools prefixes"], "anchor": "schema-rules--action--dynamic_pool--pools--prefixes", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:pools", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "pools", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/dynamic_pool/pools/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "List of IPv4 prefixes that represent an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["nat_policyCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic_pool.pools

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/)
- [rules.action.dynamic_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic_pool/)
- rules.action.dynamic_pool.pools

<a id="section"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

## Direct properties

<a id="schema-rules--action--dynamic_pool--pools--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.
