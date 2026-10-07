---
page_title: "ddos_mitigation_rules.ip_prefix_list"
subcategory: "Load Balancing"
description: "List of IP Prefix strings to match against."
xcsh_docs: {"aliases": ["ddos mitigation rules ip prefix list"], "body_bytes": 2708, "body_sha256": "sha256:21b7b4c8d261d9e3b250de4b0c0c94e6b31b42f35947e520e66bb8f207525f02", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ip_prefix_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules", "path": "documentation/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/ip_prefix_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1032202033010100-3123021202033002-2031020103231302-3003123132103022-3022131311221333-2033223333113210-1001110012203101-2113133220312302", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_mitigation_rules", "ip_prefix_list"], "schema_version": 1, "sections": [{"aliases": ["ddos mitigation rules ip prefix list invert match"], "anchor": "schema-ddos_mitigation_rules--ip_prefix_list--invert_match", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ip_prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_mitigation_rules", "ip_prefix_list", "invert_match"], "syntax": "attribute", "type": "bool"}, {"aliases": ["ddos mitigation rules ip prefix list ip prefixes"], "anchor": "schema-ddos_mitigation_rules--ip_prefix_list--ip_prefixes", "description": "List of IPv4 prefix strings.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:ddos_mitigation_rules:ip_prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_mitigation_rules", "ip_prefix_list", "ip_prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/ip_prefix_list/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "List of IP Prefix strings to match against.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_mitigation_rules.ip_prefix_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [ddos_mitigation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/ddos_mitigation_rules/)
- ddos_mitigation_rules.ip_prefix_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ddos_mitigation_rules--ip_prefix_list--invert_match"></a>

### invert_match property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

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

<a id="schema-ddos_mitigation_rules--ip_prefix_list--ip_prefixes"></a>

### ip_prefixes property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
