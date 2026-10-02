---
page_title: "splunk_receiver"
subcategory: ""
description: "Configuration for Splunk HEC Logs endpoint."
xcsh_docs: {"aliases": ["splunk receiver"], "body_bytes": 3777, "body_sha256": "sha256:b725929ff2dd1f0012432a17e66ca41670620c10f716df798b0d51d6d6875e33", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:no_tls", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:splunk_hec_token", "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/splunk_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["splunk_receiver"], "schema_version": 1, "sections": [{"aliases": ["batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["splunk_receiver", "batch"], "syntax": "attribute", "type": "object"}, {"aliases": ["compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:compression", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["splunk_receiver", "compression"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint"], "anchor": "schema-splunk_receiver--endpoint", "description": "Splunk HEC Logs Endpoint, (Note: must not contain `/services/collector`)", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "endpoint"], "syntax": "attribute", "type": "string"}, {"aliases": ["no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:no_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["splunk_receiver", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["splunk hec token"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:splunk_hec_token", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["splunk_receiver", "splunk_hec_token"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls"], "anchor": "section", "description": "TLS Parameters for client connection to the endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:splunk_receiver:use_tls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["splunk_receiver", "use_tls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/splunk_receiver/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for Splunk HEC Logs endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# splunk_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- splunk_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for Splunk HEC Logs endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

## Direct properties

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/compression/): complete subsection reference.

<a id="schema-splunk_receiver--endpoint"></a>

### endpoint property

Type: `"string"`. Computed.

Splunk HEC Logs Endpoint. Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Upstream description:

Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/no_tls/): complete subsection reference.

- [splunk_hec_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/): complete subsection reference.

## Next pages

- [splunk_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/batch/)
- [splunk_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/compression/)
- [splunk_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/no_tls/)
- [splunk_receiver.splunk_hec_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/)
- [splunk_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
