---
page_title: "user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria"
subcategory: ""
description: "Criteria for matching HTTP headers."
xcsh_docs: {"aliases": ["user defined api discovery policy discovery rules rule properties http header criteria"], "body_bytes": 5049, "body_sha256": "sha256:fac02986616c594e60f53160a8625d19f59386e18d1d6f4deae893dc123f267b", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "path": "documentation/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0222020233023033-3023331301211103-3221220323212121-2202212123211223-2023200012200111-0210120031121320-2030310021331111-2233003311001121", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria"], "schema_version": 1, "sections": [{"aliases": ["user defined api discovery policy discovery rules rule properties http header criteria field name"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--field_name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "field_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["user defined api discovery policy discovery rules rule properties http header criteria location"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--location", "description": "Specifies whether the rule criteria should be evaluated against request or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing traffic sent back to the client.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["user defined api discovery policy discovery rules rule properties http header criteria match type"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--match_type", "description": "Specifies how the value should be matched.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "match_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["user defined api discovery policy discovery rules rule properties http header criteria value"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--value", "description": "Configuration parameter for value", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Criteria for matching HTTP headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/)
- [user_defined_api_discovery_policy.discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/)
- user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http header criteria.

Additional upstream details:

Criteria for matching HTTP headers.

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

<a id="schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--field_name"></a>

### field_name property

Type: `"string"`. Computed.

HTTP Header Name. Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--location"></a>

### location property

Type: `"string"`. Computed.

\[Enum: REQUEST|RESPONSE\] Specifies whether the rule criteria should be evaluated against request
or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing
traffic sent back to the client. Possible values are \`REQUEST\`, \`RESPONSE\`. Defaults to
\`REQUEST\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "REQUEST",
  "enum": [
    "REQUEST",
    "RESPONSE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--match_type"></a>

### match_type property

Type: `"string"`. Computed.

\[Enum: EXACT\_MATCH|SUBSTRING|REGEX\] Specifies how the value should be matched. Possible values
are \`EXACT\_MATCH\`, \`SUBSTRING\`, \`REGEX\`. Defaults to \`EXACT\_MATCH\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "EXACT_MATCH",
  "enum": [
    "EXACT_MATCH",
    "SUBSTRING",
    "REGEX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--value"></a>

### value property

Type: `"string"`. Computed.

Value. Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```
