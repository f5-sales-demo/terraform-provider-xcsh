---
page_title: "azure_event_hubs_receiver.connection_string"
subcategory: ""
description: "azure_event_hubs_receiver.connection_string for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1684, "body_sha256": "sha256:c8048131b31104b107a17741a7b9d54db83066967852af41e2436127682b31c1", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:blindfold_secret_info", "xcsh-docs:data-sources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_event_hubs_receiver:connection_string", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_event_hubs_receiver", "path": "docs/guides/data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_event_hubs_receiver", "connection_string"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/azure_event_hubs_receiver/connection_string/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_event_hubs_receiver.connection_string for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure_event_hubs_receiver.connection_string

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--properties--azure_event_hubs_receiver.md)
- azure_event_hubs_receiver.connection_string

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--clear_secret_info.md): complete subsection reference.

## Next pages

- [azure_event_hubs_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--blindfold_secret_info.md)
- [azure_event_hubs_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--properties--azure_event_hubs_receiver--connection_string--clear_secret_info.md)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--properties--azure_event_hubs_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
