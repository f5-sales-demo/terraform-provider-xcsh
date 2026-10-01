---
page_title: "rules.action"
subcategory: ""
description: "rules.action for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4720, "body_sha256": "sha256:19b67238657ac6acbba5543cf043f4338925c70df559fe2eac62d035730729cc", "canonical_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:properties:rules:action:allow", "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:community", "xcsh-docs:resources:bgp_routing_policy:properties:rules:action:deny"], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:action", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules", "path": "docs/guides/resources--bgp_routing_policy--properties--rules--action.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.action for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
- [Property reference](resources--bgp_routing_policy--reference.md)
- [rules](resources--bgp_routing_policy--properties--rules.md)
- rules.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action to be enforced if the BGP route matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "as_path"),
  validators.ConflictingObjectAttributes("allow",
    "community"),
  validators.ConflictingObjectAttributes("allow",
    "deny"),
  validators.ConflictingObjectAttributes("allow",
    "local_preference"),
  validators.ConflictingObjectAttributes("allow",
    "metric"),
  validators.ConflictingObjectAttributes("as_path",
    "community"),
  validators.ConflictingObjectAttributes("as_path",
    "deny"),
  validators.ConflictingObjectAttributes("as_path",
    "local_preference"),
  validators.ConflictingObjectAttributes("as_path",
    "metric"),
  validators.ConflictingObjectAttributes("community",
    "deny"),
  validators.ConflictingObjectAttributes("community",
    "local_preference"),
  validators.ConflictingObjectAttributes("community",
    "metric"),
  validators.ConflictingObjectAttributes("deny",
    "local_preference"),
  validators.ConflictingObjectAttributes("deny",
    "metric"),
  validators.ConflictingObjectAttributes("local_preference",
    "metric")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"allow\",\"as_path\",\"community\",\"deny\",\"local_preference\",\"metric\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow](resources--bgp_routing_policy--properties--rules--action--allow.md): complete subsection reference.

<a id="schema-rules--action--as_path"></a>

### as_path property

Type: `"string"`. Optional.

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

Upstream description:

Exclusive with \[allow community deny local\_preference metric\] AS-Path Prepending is generally
used to influence incoming traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [community](resources--bgp_routing_policy--properties--rules--action--community.md): complete subsection reference.

- [deny](resources--bgp_routing_policy--properties--rules--action--deny.md): complete subsection reference.

<a id="schema-rules--action--local_preference"></a>

### local_preference property

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

Upstream description:

Exclusive with \[allow as\_path community deny metric\] BGP Local Preference is generally used to
influence outgoing traffic.

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

<a id="schema-rules--action--metric"></a>

### metric property

Type: `"number"`. Optional.

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

Upstream description:

Exclusive with \[allow as\_path community deny local\_preference\] The Multi-Exit Discriminator
metric to indicate the preferred path to AS.

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

## Next pages

- [rules.action.allow](resources--bgp_routing_policy--properties--rules--action--allow.md)
- [rules.action.community](resources--bgp_routing_policy--properties--rules--action--community.md)
- [rules.action.deny](resources--bgp_routing_policy--properties--rules--action--deny.md)
- [rules](resources--bgp_routing_policy--properties--rules.md)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
