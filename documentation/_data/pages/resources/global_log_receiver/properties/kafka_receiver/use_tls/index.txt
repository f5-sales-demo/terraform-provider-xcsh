---
page_title: "kafka_receiver.use_tls"
subcategory: ""
description: "TLS Parameters for client connection to the endpoint."
xcsh_docs: {"aliases": ["kafka receiver use tls"], "body_bytes": 6096, "body_sha256": "sha256:3e984270115854d240342f250e3e948b8ce4d4e934e7a76441be9b92e6f6916c", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_certificate", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_hostname", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_certificate", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_hostname", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_disabled", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:no_ca"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "path": "documentation/resources/global_log_receiver/properties/kafka_receiver/use_tls/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [{"anchor": "schema-kafka_receiver--use_tls--trusted_ca_url", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:no_ca", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_receiver", "use_tls"], "schema_version": 1, "sections": [{"aliases": ["kafka receiver use tls disable verify certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "disable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver use tls disable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_hostname", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "disable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver use tls enable verify certificate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_certificate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "enable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver use tls enable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_hostname", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "enable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver use tls mtls disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "mtls_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver use tls mtls enable"], "anchor": "section", "description": "MTLS Client config allows configuration of mTLS client OPTIONS.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "mtls_enable"], "syntax": "block", "type": "object"}, {"aliases": ["kafka receiver use tls no ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:no_ca", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "no_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["kafka receiver use tls trusted ca url"], "anchor": "schema-kafka_receiver--use_tls--trusted_ca_url", "description": "Exclusive with The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "use_tls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/use_tls/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "TLS Parameters for client connection to the endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.use_tls

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- kafka_receiver.use_tls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/disable_verify_certificate/): complete subsection reference.

- [disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/disable_verify_hostname/): complete subsection reference.

- [enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_certificate/): complete subsection reference.

- [enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_hostname/): complete subsection reference.

- [mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_disabled/): complete subsection reference.

- [mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/): complete subsection reference.

- [no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/no_ca/): complete subsection reference.

<a id="schema-kafka_receiver--use_tls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

## Next pages

- [kafka_receiver.use_tls.disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/disable_verify_certificate/)
- [kafka_receiver.use_tls.disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/disable_verify_hostname/)
- [kafka_receiver.use_tls.enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_certificate/)
- [kafka_receiver.use_tls.enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/enable_verify_hostname/)
- [kafka_receiver.use_tls.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_disabled/)
- [kafka_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/)
- [kafka_receiver.use_tls.no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/no_ca/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
