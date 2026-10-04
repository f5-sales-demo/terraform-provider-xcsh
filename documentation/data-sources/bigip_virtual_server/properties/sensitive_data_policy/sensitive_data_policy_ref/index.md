---
page_title: "sensitive_data_policy.sensitive_data_policy_ref"
subcategory: ""
description: "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["sensitive data policy sensitive data policy ref"], "body_bytes": 1942, "body_sha256": "sha256:d193bec0597a6d448393fe1fd18b9c4ce94d0be2c4e4891c094d0ca8be5ed44d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy", "path": "documentation/data-sources/bigip_virtual_server/properties/sensitive_data_policy/sensitive_data_policy_ref/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0011332301230200-2000313010130022-2032000233023113-0313210132201203-1200203231223131-3322100113233312-1130101100103221-0320320313323332", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref"], "schema_version": 1, "sections": [{"aliases": ["sensitive data policy sensitive data policy ref name"], "anchor": "schema-sensitive_data_policy--sensitive_data_policy_ref--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["sensitive data policy sensitive data policy ref namespace"], "anchor": "schema-sensitive_data_policy--sensitive_data_policy_ref--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["sensitive data policy sensitive data policy ref tenant"], "anchor": "schema-sensitive_data_policy--sensitive_data_policy_ref--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/sensitive_data_policy/sensitive_data_policy_ref/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_policy.sensitive_data_policy_ref

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/sensitive_data_policy/)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-sensitive_data_policy--sensitive_data_policy_ref--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-sensitive_data_policy--sensitive_data_policy_ref--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-sensitive_data_policy--sensitive_data_policy_ref--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

## Next pages

- [sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/sensitive_data_policy/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
