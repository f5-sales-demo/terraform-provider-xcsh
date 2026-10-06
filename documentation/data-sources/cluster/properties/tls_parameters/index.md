---
page_title: "tls_parameters"
subcategory: ""
description: "TLS configuration for upstream connections."
xcsh_docs: {"aliases": ["tls parameters"], "body_bytes": 3899, "body_sha256": "sha256:644b8fd902ea5c202116f9466441a375ec3f265b138232a50e0e07e152d419d4", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": ["xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params", "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params", "xcsh-docs:data-sources:cluster:properties:tls_parameters:default_session_key_caching", "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_session_key_caching", "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_sni", "xcsh-docs:data-sources:cluster:properties:tls_parameters:use_host_header_as_sni"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "documentation/data-sources/cluster/properties/tls_parameters/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031", "registry_path": "docs/guides/data-sources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "sections": [{"aliases": ["tls parameters cert params"], "anchor": "section", "description": "Certificate Parameters for authentication, TLS ciphers, and trust store.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:cert_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "cert_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters common params"], "anchor": "section", "description": "Information of different aspects for TLS authentication related to ciphers, certificates and trust store.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters default session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:default_session_key_caching", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "default_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters disable session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_session_key_caching", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "disable_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters disable sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:disable_sni", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "disable_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters max session keys"], "anchor": "schema-tls_parameters--max_session_keys", "description": "Exclusive with Number of session keys that are cached.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "max_session_keys"], "syntax": "attribute", "type": "number"}, {"aliases": ["tls parameters sni"], "anchor": "schema-tls_parameters--sni", "description": "Exclusive with SNI value to be used.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "sni"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls parameters use host header as sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:use_host_header_as_sni", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "use_host_header_as_sni"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "TLS configuration for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- tls_parameters

<a id="section"></a>

Type: `"single"`. Computed.

TLS configuration for upstream connections.

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
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]",
  "x-ves-oneof-field-tls_params_choice": "[\"cert_params\",\"common_params\"]"
}
```

## Direct properties

- [cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/cert_params/): complete subsection reference.

- [common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/): complete subsection reference.

- [default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/default_session_key_caching/): complete subsection reference.

- [disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/disable_session_key_caching/): complete subsection reference.

- [disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/disable_sni/): complete subsection reference.

<a id="schema-tls_parameters--max_session_keys"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="schema-tls_parameters--sni"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/use_host_header_as_sni/): complete subsection reference.
