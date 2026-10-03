---
page_title: "snat_pool.snat_pool"
subcategory: "Networking"
description: "List of IPv4 prefixes that represent an endpoint."
xcsh_docs: {"aliases": ["snat pool snat pool"], "body_bytes": 2309, "body_sha256": "sha256:ea60ab04bd80c4d4a2e900da00906d3a43f0960ae3944d93718fc3c65e277e05", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:snat_pool:snat_pool", "parent_id": "xcsh-docs:resources:endpoint:properties:snat_pool", "path": "documentation/resources/endpoint/properties/snat_pool/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3111213201100002-0211301331321122-0020202102011002-3313121001230300-1210221230210033-0002002022223133-2012000010023123-2222031331101001", "registry_path": "docs/guides/resources--endpoint--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["snat_pool", "snat_pool"], "schema_version": 1, "sections": [{"aliases": ["snat pool snat pool prefixes"], "anchor": "schema-snat_pool--snat_pool--prefixes", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:endpoint:properties:snat_pool:snat_pool", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["snat_pool", "snat_pool", "prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/snat_pool/snat_pool/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IPv4 prefixes that represent an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["endpointCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# snat_pool.snat_pool

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/)
- snat_pool.snat_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-snat_pool--snat_pool--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
