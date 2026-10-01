---
page_title: "insecure_registry_list"
subcategory: ""
description: "insecure_registry_list for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 3151, "body_sha256": "sha256:b9cee096365350ca933ab228ed276c32e4aa6c779bc6846ec1e2223f324c13b8", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:insecure_registry_list", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:insecure_registry_list", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--insecure_registry_list.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["insecure_registry_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/insecure_registry_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "insecure_registry_list for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# insecure_registry_list

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- insecure_registry_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: insecure\_registry\_list, no\_insecure\_registries; Default: no\_insecure\_registries\]
Docker Insecure Registry List. List of docker insecure registries.

Upstream description:

List of docker insecure registries.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("insecure_registries")}
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

OneOf alternatives in this subsection:

- [insecure_registry_list](resources--k8s_cluster--properties--insecure_registry_list.md#section)
- [no_insecure_registries](resources--k8s_cluster--properties--no_insecure_registries.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
insecure_registry_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-insecure_registry_list--insecure_registries"></a>

### insecure_registries property

Type: `["list", "string"]`. Optional.

List of docker insecure registries in format 'example.com:5000'.

Upstream description:

List of docker insecure registries in format "example.com:5000"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
