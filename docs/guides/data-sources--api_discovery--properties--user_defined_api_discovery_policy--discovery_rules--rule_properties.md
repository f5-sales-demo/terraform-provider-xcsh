---
page_title: "user_defined_api_discovery_policy.discovery_rules.rule_properties"
subcategory: ""
description: "user_defined_api_discovery_policy.discovery_rules.rule_properties for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 3914, "body_sha256": "sha256:dd5c3baa6f13e2b536eb2f1c3e60e52804921bf88aabad0a272719b48c74373b", "canonical_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:inclusion"], "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "path": "docs/guides/data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "user_defined_api_discovery_policy.discovery_rules.rule_properties for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# user_defined_api_discovery_policy.discovery_rules.rule_properties

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md)
- [Property reference](data-sources--api_discovery--reference.md)
- [user_defined_api_discovery_policy](data-sources--api_discovery--properties--user_defined_api_discovery_policy.md)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md)
- user_defined_api_discovery_policy.discovery_rules.rule_properties

<a id="section"></a>

Type: `"single"`. Computed.

Determines whether matching endpoints are included in API Discovery or excluded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-criteria": "[\"http_header_criteria\",\"pattern\"]",
  "x-ves-oneof-field-rule_type_choice": "[\"exclusion\",\"inclusion\"]"
}
```

## Direct properties

- [exclusion](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--exclusion.md): complete subsection reference.

- [http_header_criteria](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria.md): complete subsection reference.

- [inclusion](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--inclusion.md): complete subsection reference.

<a id="schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--pattern"></a>

### pattern property

Type: `"string"`. Computed.

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Upstream description:

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

## Next pages

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--exclusion.md)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria.md)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules--rule_properties--inclusion.md)
- [user_defined_api_discovery_policy.discovery_rules](data-sources--api_discovery--properties--user_defined_api_discovery_policy--discovery_rules.md)
- [xcsh_api_discovery](../data-sources/api_discovery.md)
