---
page_title: "egress_rules"
subcategory: ""
description: "egress_rules for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 5404, "body_sha256": "sha256:33458904dcef30f0dac92032165cb8a29e302b6a8a3a90b900d1fa0ad188e8e5", "canonical_id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:egress_rules:adv_action", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_tcp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_traffic", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:all_udp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:any", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:applications", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:inside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:ip_prefix_set", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:label_matcher", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:label_selector", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:metadata", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:outside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:prefix_list", "xcsh-docs:data-sources:network_policy_view:properties:egress_rules:protocol_port_range"], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:egress_rules", "parent_id": "xcsh-docs:data-sources:network_policy_view:reference", "path": "docs/guides/data-sources--network_policy_view--properties--egress_rules.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/egress_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_rules for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- egress_rules

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections from policy endpoints.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

<a id="schema-egress_rules--action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](data-sources--network_policy_view--properties--egress_rules--adv_action.md): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy_view--properties--egress_rules--all_tcp_traffic.md): complete subsection reference.

- [all_traffic](data-sources--network_policy_view--properties--egress_rules--all_traffic.md): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy_view--properties--egress_rules--all_udp_traffic.md): complete subsection reference.

- [any](data-sources--network_policy_view--properties--egress_rules--any.md): complete subsection reference.

- [applications](data-sources--network_policy_view--properties--egress_rules--applications.md): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--properties--egress_rules--inside_endpoints.md): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy_view--properties--egress_rules--ip_prefix_set.md): complete subsection reference.

- [label_matcher](data-sources--network_policy_view--properties--egress_rules--label_matcher.md): complete subsection reference.

- [label_selector](data-sources--network_policy_view--properties--egress_rules--label_selector.md): complete subsection reference.

- [metadata](data-sources--network_policy_view--properties--egress_rules--metadata.md): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--properties--egress_rules--outside_endpoints.md): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--properties--egress_rules--prefix_list.md): complete subsection reference.

- [protocol_port_range](data-sources--network_policy_view--properties--egress_rules--protocol_port_range.md): complete subsection reference.

## Next pages

- [egress_rules.adv_action](data-sources--network_policy_view--properties--egress_rules--adv_action.md)
- [egress_rules.all_tcp_traffic](data-sources--network_policy_view--properties--egress_rules--all_tcp_traffic.md)
- [egress_rules.all_traffic](data-sources--network_policy_view--properties--egress_rules--all_traffic.md)
- [egress_rules.all_udp_traffic](data-sources--network_policy_view--properties--egress_rules--all_udp_traffic.md)
- [egress_rules.any](data-sources--network_policy_view--properties--egress_rules--any.md)
- [egress_rules.applications](data-sources--network_policy_view--properties--egress_rules--applications.md)
- [egress_rules.inside_endpoints](data-sources--network_policy_view--properties--egress_rules--inside_endpoints.md)
- [egress_rules.ip_prefix_set](data-sources--network_policy_view--properties--egress_rules--ip_prefix_set.md)
- [egress_rules.label_matcher](data-sources--network_policy_view--properties--egress_rules--label_matcher.md)
- [egress_rules.label_selector](data-sources--network_policy_view--properties--egress_rules--label_selector.md)
- [egress_rules.metadata](data-sources--network_policy_view--properties--egress_rules--metadata.md)
- [egress_rules.outside_endpoints](data-sources--network_policy_view--properties--egress_rules--outside_endpoints.md)
- [egress_rules.prefix_list](data-sources--network_policy_view--properties--egress_rules--prefix_list.md)
- [egress_rules.protocol_port_range](data-sources--network_policy_view--properties--egress_rules--protocol_port_range.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
