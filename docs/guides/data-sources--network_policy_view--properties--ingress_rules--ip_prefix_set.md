---
page_title: "ingress_rules.ip_prefix_set"
subcategory: ""
description: "ingress_rules.ip_prefix_set for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1182, "body_sha256": "sha256:a9636eee13abcb09bb0b8549a4ee83f7154cccd184ae88972cee961eadd54a84", "canonical_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:ip_prefix_set", "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:ip_prefix_set", "parent_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "path": "docs/guides/data-sources--network_policy_view--properties--ingress_rules--ip_prefix_set.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_rules.ip_prefix_set for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- [ingress_rules](data-sources--network_policy_view--properties--ingress_rules.md)
- ingress_rules.ip_prefix_set

<a id="section"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [ref](data-sources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [ingress_rules.ip_prefix_set.ref](data-sources--network_policy_view--properties--ingress_rules--ip_prefix_set--ref.md)
- [ingress_rules](data-sources--network_policy_view--properties--ingress_rules.md)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
