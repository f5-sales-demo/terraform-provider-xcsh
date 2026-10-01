---
page_title: "access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms"
subcategory: ""
description: "access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 3330, "body_sha256": "sha256:65c26eac2e16be385249ca766bdbf973191084b114a44a2284b82b2ee102899d", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "child_ids": [], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "path": "docs/guides/resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "custom_hash_algorithms"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.tls_config](resources--secret_management_access--properties--access_info--tls_config.md)
- [access_info.tls_config.common_params](resources--secret_management_access--properties--access_info--tls_config--common_params.md)
- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md)
- access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms"></a>

### hash_algorithms property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [access_info.tls_config.common_params.tls_certificates](resources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
