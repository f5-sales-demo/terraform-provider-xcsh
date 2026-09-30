---
page_title: "datadog_receiver"
subcategory: ""
description: "datadog_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4269, "body_sha256": "sha256:006f10665fe06c58e6e354ac9284085d65cd0e4c6d24f7a1f28728002f3b00f4", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:datadog_api_key", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:no_tls", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:use_tls"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/datadog_receiver/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["datadog_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/datadog_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "datadog_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# datadog_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- datadog_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Datadog Configuration. Configuration for Datadog endpoint.

Upstream description:

Configuration for Datadog endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"endpoint\",\"site\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

## Direct properties

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/): complete subsection reference.

- [datadog_api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/datadog_api_key/): complete subsection reference.

<a id="schema-datadog_receiver--endpoint"></a>

### endpoint property

Type: `"string"`. Computed.

Exclusive with \[site\] Datadog Endpoint,.

Upstream description:

Exclusive with \[site\] Datadog Endpoint,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/no_tls/): complete subsection reference.

<a id="schema-datadog_receiver--site"></a>

### site property

Type: `"string"`. Computed.

Exclusive with \[endpoint\] Datadog Site,.

Upstream description:

Exclusive with \[endpoint\] Datadog Site,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/): complete subsection reference.

## Next pages

- [datadog_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/batch/)
- [datadog_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/)
- [datadog_receiver.datadog_api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/datadog_api_key/)
- [datadog_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/no_tls/)
- [datadog_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
