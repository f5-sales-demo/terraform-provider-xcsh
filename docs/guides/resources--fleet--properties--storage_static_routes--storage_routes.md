---
page_title: "storage_static_routes.storage_routes"
subcategory: ""
description: "storage_static_routes.storage_routes for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 4078, "body_sha256": "sha256:75ecc3a1639bccced59e99a87a33ae8cffea2983794669ce70a1046a4dba4383", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:labels", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes", "path": "docs/guides/resources--fleet--properties--storage_static_routes--storage_routes.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_static_routes](resources--fleet--properties--storage_static_routes.md)
- storage_static_routes.storage_routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of storage static routes.

Upstream description:

List of storage static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("subnets")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_static_routes--storage_routes--attrs"></a>

### attrs property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--fleet--properties--storage_static_routes--storage_routes--labels.md): complete subsection reference.

- [nexthop](resources--fleet--properties--storage_static_routes--storage_routes--nexthop.md): complete subsection reference.

- [subnets](resources--fleet--properties--storage_static_routes--storage_routes--subnets.md): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.labels](resources--fleet--properties--storage_static_routes--storage_routes--labels.md)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--properties--storage_static_routes--storage_routes--nexthop.md)
- [storage_static_routes.storage_routes.subnets](resources--fleet--properties--storage_static_routes--storage_routes--subnets.md)
- [storage_static_routes](resources--fleet--properties--storage_static_routes.md)
- [xcsh_fleet](../resources/fleet.md)
