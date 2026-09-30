---
page_title: "policies"
subcategory: ""
description: "policies for xcsh_service_policy_set."
xcsh_docs: {"aliases": [], "body_bytes": 1809, "body_sha256": "sha256:a949586e8acfd486f676f0c49c7b32933c334a7f0a365033f4b1c82600cf7c71", "child_ids": [], "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_set:properties:policies", "parent_id": "xcsh-docs:data-sources:service_policy_set:reference", "path": "documentation/data-sources/service_policy_set/properties/policies/index.md", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/properties/policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policies for xcsh_service_policy_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policies

Breadcrumbs:

- [xcsh_service_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/properties/)
- policies

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of references to service\_policy objects.

## Direct properties

<a id="schema-policies--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="schema-policies--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-policies--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-policies--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="schema-policies--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/properties/)
- [xcsh_service_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/)
