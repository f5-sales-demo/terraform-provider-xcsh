---
page_title: "ingress_rules"
subcategory: ""
description: "ingress_rules for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 5447, "body_sha256": "sha256:cb58d360e6211bc1ee88b18329facaf85b97313ae646521801bc52f0bf8a43e2", "canonical_id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:adv_action", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_tcp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_udp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:any", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:inside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:ip_prefix_set", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_matcher", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_selector", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:metadata", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:outside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:prefix_list", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:protocol_port_range"], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "parent_id": "xcsh-docs:data-sources:network_policy_view:reference", "path": "docs/guides/data-sources--network_policy_view--properties--ingress_rules.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/ingress_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_rules for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- ingress_rules

<a id="section"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections to policy endpoints.

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

<a id="schema-ingress_rules--action"></a>

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

- [adv_action](data-sources--network_policy_view--properties--ingress_rules--adv_action.md): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy_view--properties--ingress_rules--all_tcp_traffic.md): complete subsection reference.

- [all_traffic](data-sources--network_policy_view--properties--ingress_rules--all_traffic.md): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy_view--properties--ingress_rules--all_udp_traffic.md): complete subsection reference.

- [any](data-sources--network_policy_view--properties--ingress_rules--any.md): complete subsection reference.

- [applications](data-sources--network_policy_view--properties--ingress_rules--applications.md): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--properties--ingress_rules--inside_endpoints.md): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy_view--properties--ingress_rules--ip_prefix_set.md): complete subsection reference.

- [label_matcher](data-sources--network_policy_view--properties--ingress_rules--label_matcher.md): complete subsection reference.

- [label_selector](data-sources--network_policy_view--properties--ingress_rules--label_selector.md): complete subsection reference.

- [metadata](data-sources--network_policy_view--properties--ingress_rules--metadata.md): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--properties--ingress_rules--outside_endpoints.md): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--properties--ingress_rules--prefix_list.md): complete subsection reference.

- [protocol_port_range](data-sources--network_policy_view--properties--ingress_rules--protocol_port_range.md): complete subsection reference.

## Next pages

- [ingress_rules.adv_action](data-sources--network_policy_view--properties--ingress_rules--adv_action.md)
- [ingress_rules.all_tcp_traffic](data-sources--network_policy_view--properties--ingress_rules--all_tcp_traffic.md)
- [ingress_rules.all_traffic](data-sources--network_policy_view--properties--ingress_rules--all_traffic.md)
- [ingress_rules.all_udp_traffic](data-sources--network_policy_view--properties--ingress_rules--all_udp_traffic.md)
- [ingress_rules.any](data-sources--network_policy_view--properties--ingress_rules--any.md)
- [ingress_rules.applications](data-sources--network_policy_view--properties--ingress_rules--applications.md)
- [ingress_rules.inside_endpoints](data-sources--network_policy_view--properties--ingress_rules--inside_endpoints.md)
- [ingress_rules.ip_prefix_set](data-sources--network_policy_view--properties--ingress_rules--ip_prefix_set.md)
- [ingress_rules.label_matcher](data-sources--network_policy_view--properties--ingress_rules--label_matcher.md)
- [ingress_rules.label_selector](data-sources--network_policy_view--properties--ingress_rules--label_selector.md)
- [ingress_rules.metadata](data-sources--network_policy_view--properties--ingress_rules--metadata.md)
- [ingress_rules.outside_endpoints](data-sources--network_policy_view--properties--ingress_rules--outside_endpoints.md)
- [ingress_rules.prefix_list](data-sources--network_policy_view--properties--ingress_rules--prefix_list.md)
- [ingress_rules.protocol_port_range](data-sources--network_policy_view--properties--ingress_rules--protocol_port_range.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
