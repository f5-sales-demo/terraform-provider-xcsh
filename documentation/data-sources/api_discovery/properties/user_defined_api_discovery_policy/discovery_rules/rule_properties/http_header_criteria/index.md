---
page_title: "user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria"
subcategory: ""
description: "Criteria for matching HTTP headers."
xcsh_docs: {"aliases": ["user defined api discovery policy discovery rules rule properties http header criteria"], "body_bytes": 5809, "body_sha256": "sha256:b1e58315fde7b9fa5935c244255f34c89b31ae9167b74e0eb73f0a5f7c5dfc38", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "path": "documentation/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0222020233023033-3023331301211103-3221220323212121-2202212123211223-2023200012200111-0210120031121320-2030310021331111-2233003311001121", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria"], "schema_version": 1, "sections": [{"aliases": ["field name"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--field_name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "field_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["location"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--location", "description": "Specifies whether the rule criteria should be evaluated against request or response Applies the rule to incoming traffic from the client. Applies the rule to outgoing traffic sent back to the client.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["match type"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--match_type", "description": "Specifies how the value should be matched.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "match_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--http_header_criteria--value", "description": "Configuration parameter for value", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Criteria for matching HTTP headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

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

Upstream description:

Human-readable name for the resource

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

Specifies whether the rule criteria should be evaluated against request or response

Applies the rule to incoming traffic from the client. Applies the rule to outgoing traffic sent back
to the client.

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

Upstream description:

Specifies how the value should be matched.

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

Upstream description:

Configuration parameter for value

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [user_defined_api_discovery_policy.discovery_rules.rule_properties](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
