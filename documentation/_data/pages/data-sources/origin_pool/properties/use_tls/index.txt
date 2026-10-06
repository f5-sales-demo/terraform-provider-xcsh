---
page_title: "use_tls"
subcategory: "Load Balancing"
description: "Upstream TLS Parameters."
xcsh_docs: {"aliases": ["use tls"], "body_bytes": 4880, "body_sha256": "sha256:758e17658e2880653f2ccfd17a86bd21f4ed96302f3c6df591e4385c55e6ea4f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:use_tls:default_session_key_caching", "xcsh-docs:data-sources:origin_pool:properties:use_tls:disable_session_key_caching", "xcsh-docs:data-sources:origin_pool:properties:use_tls:disable_sni", "xcsh-docs:data-sources:origin_pool:properties:use_tls:no_mtls", "xcsh-docs:data-sources:origin_pool:properties:use_tls:skip_server_verification", "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_host_header_as_sni", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls_obj", "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_server_verification", "xcsh-docs:data-sources:origin_pool:properties:use_tls:volterra_trusted_ca"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "documentation/data-sources/origin_pool/properties/use_tls/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_tls"], "schema_version": 1, "sections": [{"aliases": ["use tls default session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:default_session_key_caching", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "default_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls disable session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:disable_session_key_caching", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "disable_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls disable sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:disable_sni", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "disable_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls max session keys"], "anchor": "schema-use_tls--max_session_keys", "description": "Exclusive with Number of session keys that are cached.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "max_session_keys"], "syntax": "attribute", "type": "number"}, {"aliases": ["use tls no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls skip server verification"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:skip_server_verification", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "skip_server_verification"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls sni"], "anchor": "schema-use_tls--sni", "description": "Exclusive with SNI value to be used.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "sni"], "syntax": "attribute", "type": "string"}, {"aliases": ["use tls tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["use_tls", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls use host header as sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_host_header_as_sni", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "use_host_header_as_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls use mtls"], "anchor": "section", "description": "MTLS Client Certificate.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["use_tls", "use_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls use mtls obj"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["use_tls", "use_mtls_obj"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls use server verification"], "anchor": "section", "description": "Upstream TLS Validation Context.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_server_verification", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["use_tls", "use_server_verification"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls volterra trusted ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:volterra_trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["use_tls", "volterra_trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Upstream TLS Parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- use_tls

<a id="section"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

## Direct properties

- [default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/default_session_key_caching/): complete subsection reference.

- [disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/disable_session_key_caching/): complete subsection reference.

- [disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/disable_sni/): complete subsection reference.

<a id="schema-use_tls--max_session_keys"></a>

### max_session_keys property

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/no_mtls/): complete subsection reference.

- [skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/skip_server_verification/): complete subsection reference.

<a id="schema-use_tls--sni"></a>

### sni property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/tls_config/): complete subsection reference.

- [use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_host_header_as_sni/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls/): complete subsection reference.

- [use_mtls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_mtls_obj/): complete subsection reference.

- [use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/use_server_verification/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/use_tls/volterra_trusted_ca/): complete subsection reference.
