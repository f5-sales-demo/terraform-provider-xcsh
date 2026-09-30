---
page_title: "splunk_receiver"
subcategory: ""
description: "splunk_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3409, "body_sha256": "sha256:aff5e8d874dd62f1569f63dda571dc380ab03b1b665b0b84c45bcbc53f1f8d18", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:no_tls", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:splunk_hec_token", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--splunk_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["splunk_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "splunk_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# splunk_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- splunk_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration for Splunk HEC Logs endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("endpoint"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
splunk_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [batch](resources--global_log_receiver--properties--splunk_receiver--batch.md): complete subsection reference.

- [compression](resources--global_log_receiver--properties--splunk_receiver--compression.md): complete subsection reference.

<a id="schema-splunk_receiver--endpoint"></a>

### endpoint property

Type: `"string"`. Optional.

Splunk HEC Logs Endpoint. Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Upstream description:

Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

- [no_tls](resources--global_log_receiver--properties--splunk_receiver--no_tls.md): complete subsection reference.

- [splunk_hec_token](resources--global_log_receiver--properties--splunk_receiver--splunk_hec_token.md): complete subsection reference.

- [use_tls](resources--global_log_receiver--properties--splunk_receiver--use_tls.md): complete subsection reference.

## Next pages

- [splunk_receiver.batch](resources--global_log_receiver--properties--splunk_receiver--batch.md)
- [splunk_receiver.compression](resources--global_log_receiver--properties--splunk_receiver--compression.md)
- [splunk_receiver.no_tls](resources--global_log_receiver--properties--splunk_receiver--no_tls.md)
- [splunk_receiver.splunk_hec_token](resources--global_log_receiver--properties--splunk_receiver--splunk_hec_token.md)
- [splunk_receiver.use_tls](resources--global_log_receiver--properties--splunk_receiver--use_tls.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
