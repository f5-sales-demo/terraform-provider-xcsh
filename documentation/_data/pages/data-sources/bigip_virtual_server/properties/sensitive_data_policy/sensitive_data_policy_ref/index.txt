---
page_title: "sensitive_data_policy.sensitive_data_policy_ref"
subcategory: ""
description: "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["sensitive data policy sensitive data policy ref"], "body_bytes": 1942, "body_sha256": "sha256:d193bec0597a6d448393fe1fd18b9c4ce94d0be2c4e4891c094d0ca8be5ed44d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy", "path": "documentation/data-sources/bigip_virtual_server/properties/sensitive_data_policy/sensitive_data_policy_ref/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0011332301230200-2000313010130022-2032000233023113-0313210132201203-1200203231223131-3322100113233312-1130101100103221-0320320313323332", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref"], "schema_version": 1, "sections": [{"aliases": ["sensitive data policy sensitive data policy ref name"], "anchor": "schema-sensitive_data_policy--sensitive_data_policy_ref--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["sensitive data policy sensitive data policy ref namespace"], "anchor": "schema-sensitive_data_policy--sensitive_data_policy_ref--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["sensitive data policy sensitive data policy ref tenant"], "anchor": "schema-sensitive_data_policy--sensitive_data_policy_ref--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/sensitive_data_policy/sensitive_data_policy_ref/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
