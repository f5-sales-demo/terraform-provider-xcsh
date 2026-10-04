---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cloud_region."
xcsh_docs: {"aliases": ["cloud region"], "body_bytes": 3556, "body_sha256": "sha256:0b80a385c64e2e35085537ef05f28642930b49def992b78a926bc53248186ed5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:cloud_region:properties:default_policy_group", "xcsh-docs:data-sources:cloud_region:properties:policy_group"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_region:reference", "parent_id": "xcsh-docs:data-sources:cloud_region:fundamentals", "path": "documentation/data-sources/cloud_region/properties/index.md", "product": "distributed-cloud", "provider_name": "cloud_region", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1110122010012032-2131203101013123-0321120130133102-2023302312022110-2201320321102001-2021000331031320-3100331122100231-2101010311313120", "registry_path": "docs/guides/data-sources--cloud_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:cloud_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["default policy group"], "anchor": "section", "description": "Configuration parameter for default policy group.", "document_id": "xcsh-docs:data-sources:cloud_region:properties:default_policy_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_policy_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:cloud_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:cloud_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:cloud_region:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the CloudRegion to look up.", "document_id": "xcsh-docs:data-sources:cloud_region:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the CloudRegion.", "document_id": "xcsh-docs:data-sources:cloud_region:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["policy group"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cloud_region:properties:policy_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policy_group"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_region/properties/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Property reference for xcsh_cloud_region.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cloud_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [default_policy_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/default_policy_group/): complete subsection reference.

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

Name of the CloudRegion to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the CloudRegion.

- [policy_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/policy_group/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/#schema-annotations) |
| `default_policy_group` | [default_policy_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/default_policy_group/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/#schema-namespace) |
| `policy_group` | [policy_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/policy_group/#section) |
| `policy_group.name` | [policy_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/policy_group/#schema-policy_group--name) |
| `policy_group.namespace` | [policy_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/policy_group/#schema-policy_group--namespace) |
| `policy_group.tenant` | [policy_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/policy_group/#schema-policy_group--tenant) |

## Next pages

- [default_policy_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/default_policy_group/)
- [policy_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/properties/policy_group/)
- [xcsh_cloud_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/)
