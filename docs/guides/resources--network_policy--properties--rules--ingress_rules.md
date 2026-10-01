---
page_title: "rules.ingress_rules"
subcategory: "Security"
description: "rules.ingress_rules for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 8005, "body_sha256": "sha256:f473b0ed156f7db4c6780ff8ab03d961b21ace1708bdaec19a9e46087016c829", "canonical_id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:ingress_rules:adv_action", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_tcp_traffic", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_traffic", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:all_udp_traffic", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:any", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:applications", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:inside_endpoints", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:ip_prefix_set", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_matcher", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:label_selector", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:metadata", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:outside_endpoints", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:prefix_list", "xcsh-docs:resources:network_policy:properties:rules:ingress_rules:protocol_port_range"], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:ingress_rules", "parent_id": "xcsh-docs:resources:network_policy:properties:rules", "path": "docs/guides/resources--network_policy--properties--rules--ingress_rules.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "ingress_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/ingress_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.ingress_rules for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ingress_rules

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
- [rules](resources--network_policy--properties--rules.md)
- rules.ingress_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections to policy endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
```

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

Terraform syntax:

```terraform
ingress_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--ingress_rules--action"></a>

### action property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

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

- [adv_action](resources--network_policy--properties--rules--ingress_rules--adv_action.md): complete subsection reference.

- [all_tcp_traffic](resources--network_policy--properties--rules--ingress_rules--all_tcp_traffic.md): complete subsection reference.

- [all_traffic](resources--network_policy--properties--rules--ingress_rules--all_traffic.md): complete subsection reference.

- [all_udp_traffic](resources--network_policy--properties--rules--ingress_rules--all_udp_traffic.md): complete subsection reference.

- [any](resources--network_policy--properties--rules--ingress_rules--any.md): complete subsection reference.

- [applications](resources--network_policy--properties--rules--ingress_rules--applications.md): complete subsection reference.

- [inside_endpoints](resources--network_policy--properties--rules--ingress_rules--inside_endpoints.md): complete subsection reference.

- [ip_prefix_set](resources--network_policy--properties--rules--ingress_rules--ip_prefix_set.md): complete subsection reference.

- [label_matcher](resources--network_policy--properties--rules--ingress_rules--label_matcher.md): complete subsection reference.

- [label_selector](resources--network_policy--properties--rules--ingress_rules--label_selector.md): complete subsection reference.

- [metadata](resources--network_policy--properties--rules--ingress_rules--metadata.md): complete subsection reference.

- [outside_endpoints](resources--network_policy--properties--rules--ingress_rules--outside_endpoints.md): complete subsection reference.

- [prefix_list](resources--network_policy--properties--rules--ingress_rules--prefix_list.md): complete subsection reference.

- [protocol_port_range](resources--network_policy--properties--rules--ingress_rules--protocol_port_range.md): complete subsection reference.

## Next pages

- [rules.ingress_rules.adv_action](resources--network_policy--properties--rules--ingress_rules--adv_action.md)
- [rules.ingress_rules.all_tcp_traffic](resources--network_policy--properties--rules--ingress_rules--all_tcp_traffic.md)
- [rules.ingress_rules.all_traffic](resources--network_policy--properties--rules--ingress_rules--all_traffic.md)
- [rules.ingress_rules.all_udp_traffic](resources--network_policy--properties--rules--ingress_rules--all_udp_traffic.md)
- [rules.ingress_rules.any](resources--network_policy--properties--rules--ingress_rules--any.md)
- [rules.ingress_rules.applications](resources--network_policy--properties--rules--ingress_rules--applications.md)
- [rules.ingress_rules.inside_endpoints](resources--network_policy--properties--rules--ingress_rules--inside_endpoints.md)
- [rules.ingress_rules.ip_prefix_set](resources--network_policy--properties--rules--ingress_rules--ip_prefix_set.md)
- [rules.ingress_rules.label_matcher](resources--network_policy--properties--rules--ingress_rules--label_matcher.md)
- [rules.ingress_rules.label_selector](resources--network_policy--properties--rules--ingress_rules--label_selector.md)
- [rules.ingress_rules.metadata](resources--network_policy--properties--rules--ingress_rules--metadata.md)
- [rules.ingress_rules.outside_endpoints](resources--network_policy--properties--rules--ingress_rules--outside_endpoints.md)
- [rules.ingress_rules.prefix_list](resources--network_policy--properties--rules--ingress_rules--prefix_list.md)
- [rules.ingress_rules.protocol_port_range](resources--network_policy--properties--rules--ingress_rules--protocol_port_range.md)
- [rules](resources--network_policy--properties--rules.md)
- [xcsh_network_policy](../resources/network_policy.md)
