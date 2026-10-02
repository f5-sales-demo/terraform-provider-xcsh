---
page_title: "bot_defense.policy.protected_app_endpoints"
subcategory: "Load Balancing"
description: "List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints"], "body_bytes": 10627, "body_sha256": "sha256:f710f8d4e1f4f2fdda09aa4b290b86cf017cb9a0fe9c2182d2534611dcd949c8", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:allow_good_bots", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:domain", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigate_good_bots", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mobile", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:path", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:query_params", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:undefined_flow_label", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web_mobile"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0213211022001102-2021221103131121-0023212122010020-3303121222223020-1030103200223133-2211210302322303-3103111330101122-3111000301230032", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints"], "schema_version": 1, "sections": [{"aliases": ["allow good bots"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:allow_good_bots", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "allow_good_bots"], "syntax": "attribute", "type": "object"}, {"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["flow label"], "anchor": "section", "description": "Bot Defense Flow Label Category allows to associate traffic with selected category.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label"], "syntax": "attribute", "type": "object"}, {"aliases": ["headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["http methods"], "anchor": "schema-bot_defense--policy--protected_app_endpoints--http_methods", "description": "List of HTTP methods.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "http_methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigate good bots"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigate_good_bots", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigate_good_bots"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigation"], "anchor": "section", "description": "Modify Bot Defense behavior for a matching request.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation"], "syntax": "attribute", "type": "object"}, {"aliases": ["mobile"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mobile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mobile"], "syntax": "attribute", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol"], "anchor": "schema-bot_defense--policy--protected_app_endpoints--protocol", "description": "SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["query params"], "anchor": "section", "description": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:query_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "query_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["undefined flow label"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:undefined_flow_label", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "undefined_flow_label"], "syntax": "attribute", "type": "object"}, {"aliases": ["web"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "web"], "syntax": "attribute", "type": "object"}, {"aliases": ["web mobile"], "anchor": "section", "description": "Web and Mobile traffic type.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web_mobile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "web_mobile"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/)
- bot_defense.policy.protected_app_endpoints

<a id="section"></a>

Type: `"list"`. Computed.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [allow_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/allow_good_bots/): complete subsection reference.

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/domain/): complete subsection reference.

- [flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/): complete subsection reference.

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/headers/): complete subsection reference.

<a id="schema-bot_defense--policy--protected_app_endpoints--http_methods"></a>

### http_methods property

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/metadata/): complete subsection reference.

- [mitigate_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigate_good_bots/): complete subsection reference.

- [mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/): complete subsection reference.

- [mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mobile/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/path/): complete subsection reference.

<a id="schema-bot_defense--policy--protected_app_endpoints--protocol"></a>

### protocol property

Type: `"string"`. Computed.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/query_params/): complete subsection reference.

- [undefined_flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/undefined_flow_label/): complete subsection reference.

- [web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/web/): complete subsection reference.

- [web_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/web_mobile/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/allow_good_bots/)
- [bot_defense.policy.protected_app_endpoints.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/any_domain/)
- [bot_defense.policy.protected_app_endpoints.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/domain/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [bot_defense.policy.protected_app_endpoints.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/headers/)
- [bot_defense.policy.protected_app_endpoints.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/metadata/)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigate_good_bots/)
- [bot_defense.policy.protected_app_endpoints.mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/)
- [bot_defense.policy.protected_app_endpoints.mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mobile/)
- [bot_defense.policy.protected_app_endpoints.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/path/)
- [bot_defense.policy.protected_app_endpoints.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/query_params/)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/undefined_flow_label/)
- [bot_defense.policy.protected_app_endpoints.web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/web/)
- [bot_defense.policy.protected_app_endpoints.web_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/web_mobile/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
