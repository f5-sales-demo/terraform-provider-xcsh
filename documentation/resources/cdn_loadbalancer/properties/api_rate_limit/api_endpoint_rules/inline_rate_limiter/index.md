---
page_title: "api_rate_limit.api_endpoint_rules.inline_rate_limiter"
subcategory: "Load Balancing"
description: "Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the required rate_limiter_choice when no stored rate-limiter object is used."
xcsh_docs: {"aliases": ["api rate limit api endpoint rules inline rate limiter"], "body_bytes": 5320, "body_sha256": "sha256:d827499f324b89859ab99617a3f1f7dd8e63c17abf3945019f6119c088d5a869", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:ref_user_id", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:use_http_lb_user_id"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "path": "documentation/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0301122111301232-2013220232300313-3011332300103100-3023223212103132-0232010111232230-0311320233220000-3301331203330121-2320222322032211", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules.inline_rate_limiter:ConflictingObjectAttributes:ref_user_id,use_http_lb_user_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:ref_user_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules.inline_rate_limiter:ConflictingObjectAttributes:ref_user_id,use_http_lb_user_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:use_http_lb_user_id", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--threshold", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules.inline_rate_limiter:RequiredObjectAttributes:threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter"], "schema_version": 1, "sections": [{"aliases": ["api rate limit api endpoint rules inline rate limiter ref user id"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:ref_user_id", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--ref_user_id--name", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:ref_user_id", "type": "requires"}], "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter", "ref_user_id"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit api endpoint rules inline rate limiter threshold"], "anchor": "schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--threshold", "description": "The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter", "threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["api rate limit api endpoint rules inline rate limiter unit"], "anchor": "schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--unit", "description": "Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter", "unit"], "syntax": "attribute", "type": "string"}, {"aliases": ["api rate limit api endpoint rules inline rate limiter use http lb user id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:use_http_lb_user_id", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter", "use_http_lb_user_id"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the required rate_limiter_choice when no stored rate-limiter object is used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.inline_rate_limiter

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inline rate limiter.

Upstream description:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref_user_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/ref_user_id/): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--threshold"></a>

### threshold property

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--unit"></a>

### unit property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/use_http_lb_user_id/): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/ref_user_id/)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/use_http_lb_user_id/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
