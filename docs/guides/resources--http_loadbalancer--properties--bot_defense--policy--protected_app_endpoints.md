---
page_title: "bot_defense.policy.protected_app_endpoints"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 9819, "body_sha256": "sha256:41122877695c48e738befa12206ca5d898f735e8ecbc4865a2a44f53c15dd270", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:allow_good_bots", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:domain", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:headers", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:metadata", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigate_good_bots", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mobile", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:path", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:query_params", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:undefined_flow_label", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web_mobile"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.protected_app_endpoints

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- bot_defense.policy.protected_app_endpoints

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("allow_good_bots",
    "mitigate_good_bots"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("flow_label",
    "undefined_flow_label"),
  validators.ConflictingListObjectAttributes("mobile",
    "web"),
  validators.ConflictingListObjectAttributes("mobile",
    "web_mobile"),
  validators.ConflictingListObjectAttributes("web",
    "web_mobile")}
```

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

Terraform syntax:

```terraform
protected_app_endpoints {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow_good_bots](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--allow_good_bots.md): complete subsection reference.

- [any_domain](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--any_domain.md): complete subsection reference.

- [domain](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--domain.md): complete subsection reference.

- [flow_label](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md): complete subsection reference.

- [headers](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--headers.md): complete subsection reference.

<a id="schema-bot_defense--policy--protected_app_endpoints--http_methods"></a>

### http_methods property

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 5),
}
```

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

- [metadata](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--metadata.md): complete subsection reference.

- [mitigate_good_bots](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigate_good_bots.md): complete subsection reference.

- [mitigation](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md): complete subsection reference.

- [mobile](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mobile.md): complete subsection reference.

- [path](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--path.md): complete subsection reference.

<a id="schema-bot_defense--policy--protected_app_endpoints--protocol"></a>

### protocol property

Type: `"string"`. Optional.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOTH",
    "HTTP",
    "HTTPS"),
}
```

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

- [query_params](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--query_params.md): complete subsection reference.

- [undefined_flow_label](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--undefined_flow_label.md): complete subsection reference.

- [web](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--web.md): complete subsection reference.

- [web_mobile](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--web_mobile.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--allow_good_bots.md)
- [bot_defense.policy.protected_app_endpoints.any_domain](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--any_domain.md)
- [bot_defense.policy.protected_app_endpoints.domain](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--domain.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--headers.md)
- [bot_defense.policy.protected_app_endpoints.metadata](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--metadata.md)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigate_good_bots.md)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md)
- [bot_defense.policy.protected_app_endpoints.mobile](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mobile.md)
- [bot_defense.policy.protected_app_endpoints.path](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--path.md)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--query_params.md)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--undefined_flow_label.md)
- [bot_defense.policy.protected_app_endpoints.web](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--web.md)
- [bot_defense.policy.protected_app_endpoints.web_mobile](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--web_mobile.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
