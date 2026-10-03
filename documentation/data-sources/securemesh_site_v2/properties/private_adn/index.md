---
page_title: "private_adn"
subcategory: ""
description: "X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support."
xcsh_docs: {"aliases": ["private adn"], "body_bytes": 2368, "body_sha256": "sha256:73c89f77d47d8d0170b311de9b005db921c34f9135f442c9d9e27fe38c1ec449", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:private_adn", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/private_adn/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1202313021223030-0121131323321332-2312202322112231-3212210331300302-0303320322301001-1121333100211113-3321310112012112-3133330323023222", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["private_adn"], "schema_version": 1, "sections": [{"aliases": ["private adn private adn"], "anchor": "schema-private_adn--private_adn", "description": "Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:private_adn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_adn", "private_adn"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/private_adn/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_adn

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- private_adn

<a id="section"></a>

Type: `"single"`. Computed.

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Upstream description:

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

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

## Direct properties

<a id="schema-private_adn--private_adn"></a>

### private_adn property

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
