---
page_title: "egress_rules.ip_prefix_set"
subcategory: ""
description: "egress_rules.ip_prefix_set for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1173, "body_sha256": "sha256:66797fda645b42b4ee83747141f727688016f7377095a7ff89cae4005eb13204", "canonical_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set", "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set:ref"], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set", "parent_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "path": "docs/guides/data-sources--network_policy_view--properties--egress_rules--ip_prefix_set.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/egress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_rules.ip_prefix_set for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- [egress_rules](data-sources--network_policy_view--properties--egress_rules.md)
- egress_rules.ip_prefix_set

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

- [ref](data-sources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [egress_rules.ip_prefix_set.ref](data-sources--network_policy_view--properties--egress_rules--ip_prefix_set--ref.md)
- [egress_rules](data-sources--network_policy_view--properties--egress_rules.md)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
