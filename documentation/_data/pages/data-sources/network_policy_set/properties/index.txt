---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_policy_set."
xcsh_docs: {"aliases": ["network policy set"], "body_bytes": 3175, "body_sha256": "sha256:b5c12dee48df6c6f46c16ec7199fdf7b13fb4ff7707ff318b21051905f540e2b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_policy_set:properties:policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_set:reference", "parent_id": "xcsh-docs:data-sources:network_policy_set:fundamentals", "path": "documentation/data-sources/network_policy_set/properties/index.md", "product": "distributed-cloud", "provider_name": "network_policy_set", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3202333230333132-0113330223032322-2031313232020132-3230321223333031-0123100330113200-2131003230131301-1122203022233100-1003132220303022", "registry_path": "docs/guides/data-sources--network_policy_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the NetworkPolicySet to look up.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the NetworkPolicySet.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["policies"], "anchor": "section", "description": "Ordered list of references to the network policy that make up this Network policy set.", "document_id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_set/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_network_policy_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the NetworkPolicySet to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the NetworkPolicySet.

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/#schema-namespace) |
| `policies` | [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/#section) |
| `policies.kind` | [policies.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/#schema-policies--kind) |
| `policies.name` | [policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/#schema-policies--name) |
| `policies.namespace` | [policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/#schema-policies--namespace) |
| `policies.tenant` | [policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/#schema-policies--tenant) |
| `policies.uid` | [policies.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/#schema-policies--uid) |
