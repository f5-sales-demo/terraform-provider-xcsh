---
page_title: "dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms"
subcategory: ""
description: "dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3051, "body_sha256": "sha256:f0b89e46daa3a19eb700faa791ad48cf44ad400fef259d6d87f7be8dedd8303f", "canonical_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:custom_hash_algorithms", "child_ids": [], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:custom_hash_algorithms", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "path": "docs/guides/resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--custom_hash_algorithms.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "custom_hash_algorithms"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [dynamic_proxy](resources--proxy--properties--dynamic_proxy.md)
- [dynamic_proxy.https_proxy](resources--proxy--properties--dynamic_proxy--https_proxy.md)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms

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

<a id="schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--custom_hash_algorithms--hash_algorithms"></a>

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

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md)
- [xcsh_proxy](../resources/proxy.md)
