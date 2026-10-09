---
page_title: "routes.route_destination.mirror_policy.percent"
subcategory: ""
description: "Fraction used where sampling percentages are needed. Example sampled requests."
xcsh_docs: {"aliases": ["routes route destination mirror policy percent"], "body_bytes": 3168, "body_sha256": "sha256:289dfa28fc855fb26e8d7b348f4226c441a0bbe499dc15d5a545eb459426f625", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "path": "documentation/resources/route/properties/routes/route_destination/mirror_policy/percent/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3110023323023222-0101122211022033-2121203321103121-3100133012031000-2322233300021221-3200131321221020-3221323010022221-0311103301212101", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [{"anchor": "schema-routes--route_destination--mirror_policy--percent--numerator", "enforcement": "provider-schema", "group": "routes.route_destination.mirror_policy.percent:RequiredObjectAttributes:numerator", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "mirror_policy", "percent"], "schema_version": 1, "sections": [{"aliases": ["routes route destination mirror policy percent denominator"], "anchor": "schema-routes--route_destination--mirror_policy--percent--denominator", "description": "Denominator used in fraction where sampling percentages are needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use million as denominator.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["HUNDRED", "MILLION", "TEN_THOUSAND"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "mirror_policy", "percent", "denominator"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination mirror policy percent numerator"], "anchor": "schema-routes--route_destination--mirror_policy--percent--numerator", "description": "Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:percent", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "mirror_policy", "percent", "numerator"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/mirror_policy/percent/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Fraction used where sampling percentages are needed. Example sampled requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["routeCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.mirror_policy.percent

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [routes.route_destination.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/)
- routes.route_destination.mirror_policy.percent

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Fraction used where sampling percentages are needed. Example sampled requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("numerator")}
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
percent {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--route_destination--mirror_policy--percent--denominator"></a>

### denominator property

Type: `"string"`. Optional.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HUNDRED","MILLION","TEN_THOUSAND"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("HUNDRED",
    "TEN_THOUSAND",
    "MILLION"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HUNDRED",
  "enum": [
    "HUNDRED",
    "TEN_THOUSAND",
    "MILLION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--route_destination--mirror_policy--percent--numerator"></a>

### numerator property

Type: `"number"`. Optional.

Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
