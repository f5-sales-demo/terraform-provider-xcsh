---
page_title: "routes.match"
subcategory: ""
description: "routes.match for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 2827, "body_sha256": "sha256:b14d82c21f908b28f65fcf06b6dfc3a449cf18b1774137152c562e7734bce099", "canonical_id": "xcsh-docs:data-sources:route:properties:routes:match", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:match:headers", "xcsh-docs:data-sources:route:properties:routes:match:incoming_port", "xcsh-docs:data-sources:route:properties:routes:match:path", "xcsh-docs:data-sources:route:properties:routes:match:query_params"], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:match", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "docs/guides/data-sources--route--properties--routes--match.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.match for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.match

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Property reference](data-sources--route--reference.md)
- [routes](data-sources--route--properties--routes.md)
- routes.match

<a id="section"></a>

Type: `"list"`. Computed.

Match. Route match condition.

Upstream description:

Route match condition.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Direct properties

- [headers](data-sources--route--properties--routes--match--headers.md): complete subsection reference.

<a id="schema-routes--match--http_method"></a>

### http_method property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--route--properties--routes--match--incoming_port.md): complete subsection reference.

- [path](data-sources--route--properties--routes--match--path.md): complete subsection reference.

- [query_params](data-sources--route--properties--routes--match--query_params.md): complete subsection reference.

## Next pages

- [routes.match.headers](data-sources--route--properties--routes--match--headers.md)
- [routes.match.incoming_port](data-sources--route--properties--routes--match--incoming_port.md)
- [routes.match.path](data-sources--route--properties--routes--match--path.md)
- [routes.match.query_params](data-sources--route--properties--routes--match--query_params.md)
- [routes](data-sources--route--properties--routes.md)
- [xcsh_route](../data-sources/route.md)
