---
page_title: "default_pool.use_tls"
subcategory: "Load Balancing"
description: "Upstream TLS Parameters."
xcsh_docs: {"aliases": ["default pool use tls"], "body_bytes": 7861, "body_sha256": "sha256:d400c2da1d0cb45e45dea56474da306adb3d343c9654f786fb06e2373151cbd6", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:disable_session_key_caching", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:disable_sni", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:no_mtls", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:skip_server_verification", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls_obj", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:volterra_trusted_ca"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/use_tls/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2312231121231002-2101200012100230-0123121210200133-0131212013023322-1001201113221120-1333333333201002-0001302121122331-1203321212112031", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "use_tls"], "schema_version": 1, "sections": [{"aliases": ["default session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "default_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:disable_session_key_caching", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "disable_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:disable_sni", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "disable_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["max session keys"], "anchor": "schema-default_pool--use_tls--max_session_keys", "description": "Exclusive with Number of session keys that are cached.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "max_session_keys"], "syntax": "attribute", "type": "number"}, {"aliases": ["no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:no_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["skip server verification"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:skip_server_verification", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "skip_server_verification"], "syntax": "attribute", "type": "object"}, {"aliases": ["sni"], "anchor": "schema-default_pool--use_tls--sni", "description": "Exclusive with SNI value to be used.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "sni"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "use_tls", "tls_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["use host header as sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "use_host_header_as_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "use mtls"], "anchor": "section", "description": "MTLS Client Certificate.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "use_tls", "use_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["use mtls obj"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_mtls_obj", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "use_tls", "use_mtls_obj"], "syntax": "attribute", "type": "object"}, {"aliases": ["use server verification"], "anchor": "section", "description": "Upstream TLS Validation Context.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "use_tls", "use_server_verification"], "syntax": "attribute", "type": "object"}, {"aliases": ["volterra trusted ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:use_tls:volterra_trusted_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "volterra_trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/use_tls/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Upstream TLS Parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- default_pool.use_tls

<a id="section"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

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

- [default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/default_session_key_caching/): complete subsection reference.

- [disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/disable_session_key_caching/): complete subsection reference.

- [disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/disable_sni/): complete subsection reference.

<a id="schema-default_pool--use_tls--max_session_keys"></a>

### max_session_keys property

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/no_mtls/): complete subsection reference.

- [skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/skip_server_verification/): complete subsection reference.

<a id="schema-default_pool--use_tls--sni"></a>

### sni property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/): complete subsection reference.

- [use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_host_header_as_sni/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/): complete subsection reference.

- [use_mtls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls_obj/): complete subsection reference.

- [use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_server_verification/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/volterra_trusted_ca/): complete subsection reference.

## Next pages

- [default_pool.use_tls.default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/default_session_key_caching/)
- [default_pool.use_tls.disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/disable_session_key_caching/)
- [default_pool.use_tls.disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/disable_sni/)
- [default_pool.use_tls.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/no_mtls/)
- [default_pool.use_tls.skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/skip_server_verification/)
- [default_pool.use_tls.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/tls_config/)
- [default_pool.use_tls.use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_host_header_as_sni/)
- [default_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/)
- [default_pool.use_tls.use_mtls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_mtls_obj/)
- [default_pool.use_tls.use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/use_server_verification/)
- [default_pool.use_tls.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/use_tls/volterra_trusted_ca/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
