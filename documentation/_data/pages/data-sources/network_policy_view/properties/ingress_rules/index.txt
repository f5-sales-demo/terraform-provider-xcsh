---
page_title: "ingress_rules"
subcategory: ""
description: "ingress_rules for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 6956, "body_sha256": "sha256:c25c2d9bd3c7af85dc3f66c2661c10ea12105f8178826e5f9cdadef769f83d57", "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:adv_action", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_tcp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:all_udp_traffic", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:any", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:applications", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:inside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:ip_prefix_set", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_matcher", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:label_selector", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:metadata", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:outside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:prefix_list", "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules:protocol_port_range"], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:ingress_rules", "parent_id": "xcsh-docs:data-sources:network_policy_view:reference", "path": "documentation/data-sources/network_policy_view/properties/ingress_rules/index.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["ingress_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/ingress_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_rules for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_rules

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
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

- [adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/adv_action/): complete subsection reference.

- [all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_tcp_traffic/): complete subsection reference.

- [all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_traffic/): complete subsection reference.

- [all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_udp_traffic/): complete subsection reference.

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/any/): complete subsection reference.

- [applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/applications/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/inside_endpoints/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/): complete subsection reference.

- [label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_matcher/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/metadata/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/prefix_list/): complete subsection reference.

- [protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/protocol_port_range/): complete subsection reference.

## Next pages

- [ingress_rules.adv_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/adv_action/)
- [ingress_rules.all_tcp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_tcp_traffic/)
- [ingress_rules.all_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_traffic/)
- [ingress_rules.all_udp_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/all_udp_traffic/)
- [ingress_rules.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/any/)
- [ingress_rules.applications](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/applications/)
- [ingress_rules.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/inside_endpoints/)
- [ingress_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/ip_prefix_set/)
- [ingress_rules.label_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_matcher/)
- [ingress_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/label_selector/)
- [ingress_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/metadata/)
- [ingress_rules.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/outside_endpoints/)
- [ingress_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/prefix_list/)
- [ingress_rules.protocol_port_range](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/ingress_rules/protocol_port_range/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
