---
page_title: "rules.match.community"
subcategory: ""
description: "rules.match.community for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2752, "body_sha256": "sha256:de1d51924f2beb71dce53b031677981c8c527d5db7dfad8fd734d46bc98f77bc", "canonical_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match:community", "parent_id": "xcsh-docs:resources:bgp_routing_policy:properties:rules:match", "path": "docs/guides/resources--bgp_routing_policy--properties--rules--match--community.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "match", "community"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/properties/rules/match/community/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.match.community for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.match.community

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
- [Property reference](resources--bgp_routing_policy--reference.md)
- [rules](resources--bgp_routing_policy--properties--rules.md)
- [rules.match](resources--bgp_routing_policy--properties--rules--match.md)
- rules.match.community

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

BGP Community list. List of BGP communities.

Upstream description:

List of BGP communities.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("community")}
```

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

Terraform syntax:

```terraform
community {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rules--match--community--community"></a>

### community property

Type: `["list", "string"]`. Optional.

Unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being
value.

Upstream description:

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rules.match](resources--bgp_routing_policy--properties--rules--match.md)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
