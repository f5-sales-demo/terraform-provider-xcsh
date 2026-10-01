---
page_title: "sensitive_data_policy.sensitive_data_policy_ref"
subcategory: ""
description: "sensitive_data_policy.sensitive_data_policy_ref for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 1942, "body_sha256": "sha256:d193bec0597a6d448393fe1fd18b9c4ce94d0be2c4e4891c094d0ca8be5ed44d", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy:sensitive_data_policy_ref", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:sensitive_data_policy", "path": "documentation/data-sources/bigip_virtual_server/properties/sensitive_data_policy/sensitive_data_policy_ref/index.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["sensitive_data_policy", "sensitive_data_policy_ref"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/sensitive_data_policy/sensitive_data_policy_ref/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sensitive_data_policy.sensitive_data_policy_ref for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
