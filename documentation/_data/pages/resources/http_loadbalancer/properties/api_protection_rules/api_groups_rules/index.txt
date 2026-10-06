---
page_title: "api_protection_rules.api_groups_rules"
subcategory: "Load Balancing"
description: "This category includes rules per API group or Server URL. For API groups, refer to API Definition which includes API groups derived from uploaded swaggers."
xcsh_docs: {"aliases": ["api protection rules api groups rules"], "body_bytes": 6499, "body_sha256": "sha256:f6c288459f5eb109d721b2609f838d8f8325b33ccaf1753984b5dbc93d2a2193", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "schema-api_protection_rules--api_groups_rules--specific_domain", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:any_domain", "type": "conflicts"}, {"anchor": "schema-api_protection_rules--api_groups_rules--base_path", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules:RequiredListObjectAttributes:base_path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules"], "schema_version": 1, "sections": [{"aliases": ["api protection rules api groups rules action"], "anchor": "section", "description": "The action to take if the input request matches the rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:deny", "type": "conflicts"}], "schema_path": ["api_protection_rules", "api_groups_rules", "action"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api groups rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:any_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["api protection rules api groups rules api group"], "anchor": "schema-api_protection_rules--api_groups_rules--api_group", "description": "API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with \"x-F5 Distributed Cloud-API-group\" extensions", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "api_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["api protection rules api groups rules base path"], "anchor": "schema-api_protection_rules--api_groups_rules--base_path", "description": "Prefix of the request path. For example: /v1.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "base_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["api protection rules api groups rules client matcher"], "anchor": "section", "description": "Client conditions for matching a rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_client,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:client_selector,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:any_client,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_threat_category_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.client_matcher:ConflictingObjectAttributes:client_selector,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_threat_category_list", "type": "conflicts"}], "schema_path": ["api_protection_rules", "api_groups_rules", "client_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api groups rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_protection_rules--api_groups_rules--metadata--name", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:metadata", "type": "requires"}], "schema_path": ["api_protection_rules", "api_groups_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api groups rules request matcher"], "anchor": "section", "description": "Request conditions for matching a rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["api protection rules api groups rules specific domain"], "anchor": "schema-api_protection_rules--api_groups_rules--specific_domain", "description": "Exclusive with The rule will apply for a specific domain. For example: api.example.com.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "specific_domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This category includes rules per API group or Server URL. For API groups, refer to API Definition which includes API groups derived from uploaded swaggers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- api_protection_rules.api_groups_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

This category includes rules per API group or Server URL. For API groups, refer to API Definition
which includes API groups derived from uploaded swaggers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain")}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
api_groups_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/): complete subsection reference.

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/any_domain/): complete subsection reference.

<a id="schema-api_protection_rules--api_groups_rules--api_group"></a>

### api_group property

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Additional upstream details:

Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-api_protection_rules--api_groups_rules--base_path"></a>

### base_path property

Type: `"string"`. Optional.

Base Path. Prefix of the request path. For example: /v1.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/metadata/): complete subsection reference.

- [request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/): complete subsection reference.

<a id="schema-api_protection_rules--api_groups_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
