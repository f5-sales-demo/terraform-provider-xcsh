---
page_title: "dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms"
subcategory: ""
description: "Specifies the hash algorithms to be used."
xcsh_docs: {"aliases": ["cert", "certificate", "dynamic proxy https proxy tls params tls certificates custom hash algorithms", "existing certificates", "tls certificates"], "body_bytes": 3551, "body_sha256": "sha256:30d7673f283ff2ae0e6b9c257d390cd6f6b65177efdfa76d31868a2c57286576", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:custom_hash_algorithms", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/custom_hash_algorithms/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1333111333313230-3121323312102302-1312320210203101-1321110330113322-0323123321002010-3111101110023221-0112322321000011-1000311011002201", "registry_path": "docs/guides/resources--proxy--reference--group-003.md", "relationships": [{"anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--custom_hash_algorithms--hash_algorithms", "enforcement": "provider-schema", "group": "dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms:RequiredObjectAttributes:hash_algorithms", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:custom_hash_algorithms", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "custom_hash_algorithms"], "schema_version": 1, "sections": [{"aliases": ["hash algorithms"], "anchor": "schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--custom_hash_algorithms--hash_algorithms", "description": "Ordered list of hash algorithms to be used.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:tls_params:tls_certificates:custom_hash_algorithms", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "tls_params", "tls_certificates", "custom_hash_algorithms", "hash_algorithms"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies the hash algorithms to be used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.tls_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/)
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

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/tls_params/tls_certificates/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
