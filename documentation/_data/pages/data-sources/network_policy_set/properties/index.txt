---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_policy_set."
xcsh_docs: {"aliases": ["network policy set"], "body_bytes": 3432, "body_sha256": "sha256:e8260d3ef180442eae032951adf40e65eb7e20e94c2bfe586b17de7256ce69e0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_policy_set:properties:policies"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_set:reference", "parent_id": "xcsh-docs:data-sources:network_policy_set:fundamentals", "path": "documentation/data-sources/network_policy_set/properties/index.md", "product": "distributed-cloud", "provider_name": "network_policy_set", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3202333230333132-0113330223032322-2031313232020132-3230321223333031-0123100330113200-2131003230131301-1122203022233100-1003132220303022", "registry_path": "docs/guides/data-sources--network_policy_set--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the NetworkPolicySet to look up.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the NetworkPolicySet.", "document_id": "xcsh-docs:data-sources:network_policy_set:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["policies"], "anchor": "section", "description": "Ordered list of references to the network policy that make up this Network policy set.", "document_id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_set/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_network_policy_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/policies/)
- [xcsh_network_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/)
