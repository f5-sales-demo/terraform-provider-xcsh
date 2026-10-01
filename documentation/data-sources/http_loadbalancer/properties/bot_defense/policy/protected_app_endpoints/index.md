---
page_title: "bot_defense.policy.protected_app_endpoints"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 10627, "body_sha256": "sha256:f710f8d4e1f4f2fdda09aa4b290b86cf017cb9a0fe9c2182d2534611dcd949c8", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:allow_good_bots", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:domain", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigate_good_bots", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mobile", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:path", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:query_params", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:undefined_flow_label", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web_mobile"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
