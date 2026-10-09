---
page_title: "api_protection_rules.api_endpoint_rules"
subcategory: "Load Balancing"
description: "This category defines specific rules per API endpoints. If request matches any of these rules, skipping second category rules."
xcsh_docs: {"aliases": ["api protection rules api endpoint rules"], "body_bytes": 4771, "body_sha256": "sha256:e821ff579a13c07a05430ebae5a1dedafcbfe8f20412dc9fb7489423e53e8715", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:api_endpoint_method", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules"], "schema_version": 1, "sections": [{"aliases": ["api protection rules api endpoint rules action"], "anchor": "section", "description": "The action to take if the input request matches the rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "action"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api endpoint rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:any_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["api protection rules api endpoint rules api endpoint method", "succeeded", "success", "successful"], "anchor": "section", "description": "A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:api_endpoint_method", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "api_endpoint_method"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api endpoint rules api endpoint path"], "anchor": "schema-api_protection_rules--api_endpoint_rules--api_endpoint_path", "description": "The endpoint (path) of the request.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "api_endpoint_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["api protection rules api endpoint rules client matcher"], "anchor": "section", "description": "Client conditions for matching a rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:client_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "client_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api endpoint rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api endpoint rules request matcher"], "anchor": "section", "description": "Request conditions for matching a rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api endpoint rules specific domain"], "anchor": "schema-api_protection_rules--api_endpoint_rules--specific_domain", "description": "Exclusive with The rule will apply for a specific domain. For example: api.example.com.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "specific_domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This category defines specific rules per API endpoints. If request matches any of these rules, skipping second category rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- api_protection_rules.api_endpoint_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

This category defines specific rules per API endpoints. If request matches any of these rules,
skipping second category rules.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_endpoint_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/action/): complete subsection reference.

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/any_domain/): complete subsection reference.

- [api_endpoint_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/api_endpoint_method/): complete subsection reference.

<a id="schema-api_protection_rules--api_endpoint_rules--api_endpoint_path"></a>

### api_endpoint_path property

Type: `"string"`. Optional.

API Endpoint. The endpoint (path) of the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/client_matcher/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/metadata/): complete subsection reference.

- [request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/): complete subsection reference.

<a id="schema-api_protection_rules--api_endpoint_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```
