---
page_title: "routes.match"
subcategory: ""
description: "routes.match for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 3064, "body_sha256": "sha256:a58507fe803e4954c4c86a1ad6fb35696e23d3921f3d47509251813baf4437da", "canonical_id": "xcsh-docs:resources:route:properties:routes:match", "child_ids": ["xcsh-docs:resources:route:properties:routes:match:headers", "xcsh-docs:resources:route:properties:routes:match:incoming_port", "xcsh-docs:resources:route:properties:routes:match:path", "xcsh-docs:resources:route:properties:routes:match:query_params"], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:match", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "docs/guides/resources--route--properties--routes--match.md", "provider_name": "route", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.match for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.match

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Property reference](resources--route--reference.md)
- [routes](resources--route--properties--routes.md)
- routes.match

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
match {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](resources--route--properties--routes--match--headers.md): complete subsection reference.

<a id="schema-routes--match--http_method"></a>

### http_method property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

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

- [incoming_port](resources--route--properties--routes--match--incoming_port.md): complete subsection reference.

- [path](resources--route--properties--routes--match--path.md): complete subsection reference.

- [query_params](resources--route--properties--routes--match--query_params.md): complete subsection reference.

## Next pages

- [routes.match.headers](resources--route--properties--routes--match--headers.md)
- [routes.match.incoming_port](resources--route--properties--routes--match--incoming_port.md)
- [routes.match.path](resources--route--properties--routes--match--path.md)
- [routes.match.query_params](resources--route--properties--routes--match--query_params.md)
- [routes](resources--route--properties--routes.md)
- [xcsh_route](../resources/route.md)
