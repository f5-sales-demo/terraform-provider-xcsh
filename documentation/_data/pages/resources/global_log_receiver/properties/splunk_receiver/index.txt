---
page_title: "splunk_receiver"
subcategory: ""
description: "splunk_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4117, "body_sha256": "sha256:37cf3b0528c1bc961c7e516259d6ce17860c0260952bc2395946172637c66239", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:no_tls", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:splunk_hec_token", "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:use_tls"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/splunk_receiver/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["splunk_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "splunk_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# splunk_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
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

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/compression/): complete subsection reference.

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

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/no_tls/): complete subsection reference.

- [splunk_hec_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/): complete subsection reference.

## Next pages

- [splunk_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/)
- [splunk_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/compression/)
- [splunk_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/no_tls/)
- [splunk_receiver.splunk_hec_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/splunk_hec_token/)
- [splunk_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
