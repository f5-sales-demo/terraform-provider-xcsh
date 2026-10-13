---
page_title: "rules.action.dynamic_pool.elastic_ips.refs"
subcategory: ""
description: "Reference to one or more cloud elastic IP objects."
xcsh_docs: {"aliases": ["rules action dynamic pool elastic ips refs"], "body_bytes": 2434, "body_sha256": "sha256:c332896054823723559b0632dd9195112e0d3e911e255da27b897b5545af5ee6", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips", "path": "documentation/data-sources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/refs/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "data-sources", "registry_anchor": "canonical-2303111331203033-3231212211032103-2111010111112110-1332223322100321-3333132030011312-0020323002121230-1011323022323230-1231332121333111", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic pool elastic ips refs kind"], "anchor": "schema-rules--action--dynamic_pool--elastic_ips--refs--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. 'route').", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules action dynamic pool elastic ips refs name"], "anchor": "schema-rules--action--dynamic_pool--elastic_ips--refs--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules action dynamic pool elastic ips refs namespace"], "anchor": "schema-rules--action--dynamic_pool--elastic_ips--refs--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules action dynamic pool elastic ips refs tenant"], "anchor": "schema-rules--action--dynamic_pool--elastic_ips--refs--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["rules action dynamic pool elastic ips refs uid"], "anchor": "schema-rules--action--dynamic_pool--elastic_ips--refs--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/refs/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Reference to one or more cloud elastic IP objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["nat_policyCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic_pool.elastic_ips.refs

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/)
- [rules.action.dynamic_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic_pool/)
- [rules.action.dynamic_pool.elastic_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/)
- rules.action.dynamic_pool.elastic_ips.refs

<a id="section"></a>

Type: `"list"`. Computed.

Reference to one or more cloud elastic IP objects.

## Direct properties

<a id="schema-rules--action--dynamic_pool--elastic_ips--refs--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-rules--action--dynamic_pool--elastic_ips--refs--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-rules--action--dynamic_pool--elastic_ips--refs--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-rules--action--dynamic_pool--elastic_ips--refs--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-rules--action--dynamic_pool--elastic_ips--refs--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.
