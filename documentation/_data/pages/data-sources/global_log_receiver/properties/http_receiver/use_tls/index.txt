---
page_title: "http_receiver.use_tls"
subcategory: ""
description: "TLS Parameters for client connection to the endpoint."
xcsh_docs: {"aliases": ["http receiver use tls"], "body_bytes": 5417, "body_sha256": "sha256:a42b43ac0439178555654de2e32182a543f444c05e889472adb7e1778a11d0e2", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_certificate", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_hostname", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_certificate", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_hostname", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:no_ca"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver", "path": "documentation/data-sources/global_log_receiver/properties/http_receiver/use_tls/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_receiver", "use_tls"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "disable verify certificate", "existing certificates", "tls certificates"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "disable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_hostname", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "disable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "enable verify certificate", "existing certificates", "tls certificates"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_certificate", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "enable_verify_certificate"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable verify hostname"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_hostname", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "enable_verify_hostname"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtls disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "mtls_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtls enable"], "anchor": "section", "description": "MTLS Client config allows configuration of mTLS client OPTIONS.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http_receiver", "use_tls", "mtls_enable"], "syntax": "attribute", "type": "object"}, {"aliases": ["no ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:no_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "no_ca"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "trusted ca url"], "anchor": "schema-http_receiver--use_tls--trusted_ca_url", "description": "Exclusive with The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_receiver", "use_tls", "trusted_ca_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/http_receiver/use_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "TLS Parameters for client connection to the endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.use_tls

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/)
- http_receiver.use_tls

<a id="section"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

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

## Direct properties

- [disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/disable_verify_certificate/): complete subsection reference.

- [disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/disable_verify_hostname/): complete subsection reference.

- [enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/enable_verify_certificate/): complete subsection reference.

- [enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/enable_verify_hostname/): complete subsection reference.

- [mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_disabled/): complete subsection reference.

- [mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/): complete subsection reference.

- [no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/no_ca/): complete subsection reference.

<a id="schema-http_receiver--use_tls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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

- [http_receiver.use_tls.disable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/disable_verify_certificate/)
- [http_receiver.use_tls.disable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/disable_verify_hostname/)
- [http_receiver.use_tls.enable_verify_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/enable_verify_certificate/)
- [http_receiver.use_tls.enable_verify_hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/enable_verify_hostname/)
- [http_receiver.use_tls.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_disabled/)
- [http_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/)
- [http_receiver.use_tls.no_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/use_tls/no_ca/)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
