---
page_title: "routes.simple_route.advanced_options.specific_hash_policy.hash_policy"
subcategory: "Load Balancing"
description: "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request."
xcsh_docs: {"aliases": ["routes simple route advanced options specific hash policy hash policy"], "body_bytes": 5938, "body_sha256": "sha256:fe3ffa830b1692d294bab24288232669808cc986d1e13f51be8b5dc01e3d304f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2333120110011121-1313332102102032-3201233223100103-3211320101012332-3032321222123131-2001321033010020-3333231101303312-3110032111130110", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-026.md", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--header_name", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--header_name", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--source_ip", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--source_ip", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy"], "schema_version": 1, "sections": [{"aliases": ["cookie"], "anchor": "section", "description": "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:add_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:add_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:ignore_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:ConflictingObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--cookie--name", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy:cookie", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy", "cookie"], "syntax": "block", "type": "object"}, {"aliases": ["header name"], "anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--header_name", "description": "Exclusive with The name or key of the request header that will be used to obtain the hash key.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy", "header_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["source ip"], "anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--source_ip", "description": "Exclusive with Hash based on source IP address.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy", "source_ip"], "syntax": "attribute", "type": "bool"}, {"aliases": ["terminal"], "anchor": "schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--terminal", "description": "Specify if its a terminal policy.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy", "hash_policy", "terminal"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.specific_hash_policy.hash_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [routes.simple_route.advanced_options.specific_hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Upstream description:

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "header_name"),
  validators.ConflictingListObjectAttributes("cookie",
    "source_ip"),
  validators.ConflictingListObjectAttributes("header_name",
    "source_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
hash_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/cookie/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--header_name"></a>

### header_name property

Type: `"string"`. Optional.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Upstream description:

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--source_ip"></a>

### source_ip property

Type: `"bool"`. Optional.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Upstream description:

Exclusive with \[cookie header\_name\] Hash based on source IP address.

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

<a id="schema-routes--simple_route--advanced_options--specific_hash_policy--hash_policy--terminal"></a>

### terminal property

Type: `"bool"`. Optional.

Terminal. Specify if its a terminal policy.

Upstream description:

Specify if its a terminal policy.

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

## Next pages

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/hash_policy/cookie/)
- [routes.simple_route.advanced_options.specific_hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
