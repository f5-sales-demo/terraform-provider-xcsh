---
page_title: "tls_parameters"
subcategory: ""
description: "TLS configuration for upstream connections."
xcsh_docs: {"aliases": ["tls parameters"], "body_bytes": 5068, "body_sha256": "sha256:cdfd0bc338242c656a980bfdc6ffa8d1581011e28945a7368c973e3cf081abb1", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": ["xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "xcsh-docs:resources:cluster:properties:tls_parameters:common_params", "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters", "parent_id": "xcsh-docs:resources:cluster:reference", "path": "documentation/resources/cluster/properties/tls_parameters/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [{"anchor": "schema-tls_parameters--max_session_keys", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--max_session_keys", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--sni", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--sni", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters"], "schema_version": 1, "sections": [{"aliases": ["tls parameters cert params"], "anchor": "section", "description": "Certificate Parameters for authentication, TLS ciphers, and trust store.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,tls_validation_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,tls_validation_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:tls_validation_params,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:tls_validation_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:volterra_trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:ConflictingObjectAttributes:tls_validation_params,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:volterra_trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.cert_params:RequiredObjectAttributes:certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:certificates", "type": "requires"}], "schema_path": ["tls_parameters", "cert_params"], "syntax": "block", "type": "object"}, {"aliases": ["tls parameters common params"], "anchor": "section", "description": "Information of different aspects for TLS authentication related to ciphers, certificates and trust store.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params"], "syntax": "block", "type": "object"}, {"aliases": ["tls parameters default session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "default_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters disable session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "disable_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters disable sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "disable_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls parameters max session keys"], "anchor": "schema-tls_parameters--max_session_keys", "description": "Exclusive with Number of session keys that are cached.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "max_session_keys"], "syntax": "attribute", "type": "number"}, {"aliases": ["tls parameters sni"], "anchor": "schema-tls_parameters--sni", "description": "Exclusive with SNI value to be used.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "sni"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls parameters use host header as sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_parameters", "use_host_header_as_sni"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "TLS configuration for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["clusterCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
EnumExtractionComplete: false
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/use_host_header_as_sni/): complete subsection reference.
