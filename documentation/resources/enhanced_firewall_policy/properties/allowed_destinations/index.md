---
page_title: "allowed_destinations"
subcategory: ""
description: "allowed_destinations for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2358, "body_sha256": "sha256:57fae24d47c09e1cada921fa42cde2b7fd040fad025b93272862ddabb1e44806", "child_ids": [], "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:allowed_destinations", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:reference", "path": "documentation/resources/enhanced_firewall_policy/properties/allowed_destinations/index.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["allowed_destinations"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/allowed_destinations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allowed_destinations for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allowed_destinations

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- allowed_destinations

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
allowed_destinations {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-allowed_destinations--prefix"></a>

### prefix property

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
