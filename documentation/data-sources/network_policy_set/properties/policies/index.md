---
page_title: "policies"
subcategory: ""
description: "policies for xcsh_network_policy_set."
xcsh_docs: {"aliases": [], "body_bytes": 1940, "body_sha256": "sha256:60164f1240470dafa2d0e60ea0475b922278c8e8832bd993fc2d7403483cd296", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_set:properties:policies", "parent_id": "xcsh-docs:data-sources:network_policy_set:reference", "path": "documentation/data-sources/network_policy_set/properties/policies/index.md", "provider_name": "network_policy_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_set/properties/policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policies for xcsh_network_policy_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policies

Breadcrumbs:

- [xcsh_network_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/)
- policies

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of references to the network policy that make up this Network policy set.

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/properties/)
- [xcsh_network_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/)
