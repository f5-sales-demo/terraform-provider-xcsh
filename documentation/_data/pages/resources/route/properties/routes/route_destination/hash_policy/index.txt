---
page_title: "routes.route_destination.hash_policy"
subcategory: ""
description: "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request."
xcsh_docs: {"aliases": ["routes route destination hash policy"], "body_bytes": 4871, "body_sha256": "sha256:eceefe2119b09381960329e9cad1b9181e854e90829e0adcf47e800c9d658e73", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "parent_id": "xcsh-docs:resources:route:properties:routes:route_destination", "path": "documentation/resources/route/properties/routes/route_destination/hash_policy/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123", "registry_path": "docs/guides/resources--route--reference--group-002.md", "relationships": [{"anchor": "schema-routes--route_destination--hash_policy--header_name", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--hash_policy--header_name", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--hash_policy--source_ip", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--hash_policy--source_ip", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination", "hash_policy"], "schema_version": 1, "sections": [{"aliases": ["cookie"], "anchor": "section", "description": "Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from the client in its response to the client, based on the endpoint the request gets sent to. The client then", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:add_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:add_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:ignore_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:ignore_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:ConflictingObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie:samesite_strict", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--hash_policy--cookie--name", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy.cookie:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "type": "requires"}], "schema_path": ["routes", "route_destination", "hash_policy", "cookie"], "syntax": "block", "type": "object"}, {"aliases": ["header name"], "anchor": "schema-routes--route_destination--hash_policy--header_name", "description": "Exclusive with The name or key of the request header that will be used to obtain the hash key.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "header_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["source ip"], "anchor": "schema-routes--route_destination--hash_policy--source_ip", "description": "Exclusive with Hash based on source IP address.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "source_ip"], "syntax": "attribute", "type": "bool"}, {"aliases": ["terminal"], "anchor": "schema-routes--route_destination--hash_policy--terminal", "description": "Specify if its a terminal policy.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "hash_policy", "terminal"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/hash_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination.hash_policy

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- routes.route_destination.hash_policy

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
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

- [cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/): complete subsection reference.

<a id="schema-routes--route_destination--hash_policy--header_name"></a>

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

<a id="schema-routes--route_destination--hash_policy--source_ip"></a>

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

<a id="schema-routes--route_destination--hash_policy--terminal"></a>

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

- [routes.route_destination.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
