---
page_title: "tls_parameters"
subcategory: ""
description: "TLS configuration for upstream connections."
xcsh_docs: {"aliases": ["tls parameters"], "body_bytes": 6409, "body_sha256": "sha256:2a5fc10d46733a5c406f164b5fddcc218ccc72ba22844b56ecdda6e8efa87a9c", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": ["xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "xcsh-docs:resources:cluster:properties:tls_parameters:common_params", "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/tls_parameters/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [{"anchor": "schema-tls_parameters--max_session_keys", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--max_session_keys", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--sni", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--sni", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "sections": [{"aliases": ["cert params"], "anchor": "section", "description": "Certificate Parameters for authentication, TLS ciphers, and trust store.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,tls_validation_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,tls_validation_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:tls_validation_params,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:volterra_trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:tls_validation_params,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:volterra_trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:RequiredObjectAttributes:certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:certificates", "type": "requires"}], "schema_path": ["tls_parameters", "cert_params"], "syntax": "block", "type": "object"}, {"aliases": ["common params"], "anchor": "section", "description": "Information of different aspects for TLS authentication related to ciphers, certificates and trust store.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params"], "syntax": "block", "type": "object"}, {"aliases": ["default session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "default_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "disable_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "disable_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["max session keys"], "anchor": "schema-tls_parameters--max_session_keys", "description": "Exclusive with Number of session keys that are cached.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "max_session_keys"], "syntax": "attribute", "type": "number"}, {"aliases": ["sni"], "anchor": "schema-tls_parameters--sni", "description": "Exclusive with SNI value to be used.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "sni"], "syntax": "attribute", "type": "string"}, {"aliases": ["use host header as sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "use_host_header_as_sni"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "TLS configuration for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- tls_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for upstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cert_params",
    "common_params"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni")}
```

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

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/): complete subsection reference.

- [common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/): complete subsection reference.

- [default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/default_session_key_caching/): complete subsection reference.

- [disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/disable_session_key_caching/): complete subsection reference.

- [disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/disable_sni/): complete subsection reference.

<a id="schema-tls_parameters--max_session_keys"></a>

### max_session_keys property

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

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

<a id="schema-tls_parameters--sni"></a>

### sni property

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/use_host_header_as_sni/): complete subsection reference.

## Next pages

- [tls_parameters.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/)
- [tls_parameters.default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/default_session_key_caching/)
- [tls_parameters.disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/disable_session_key_caching/)
- [tls_parameters.disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/disable_sni/)
- [tls_parameters.use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/use_host_header_as_sni/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
