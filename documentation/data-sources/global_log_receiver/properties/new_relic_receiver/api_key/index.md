---
page_title: "new_relic_receiver.api_key"
subcategory: ""
description: "new_relic_receiver.api_key for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1973, "body_sha256": "sha256:cfd5dcd427f29ff8d20a01f9045475ef9fd488b2ca2b1633e151bab813d44d46", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver:api_key:blindfold_secret_info", "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver:api_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver:api_key", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver", "path": "documentation/data-sources/global_log_receiver/properties/new_relic_receiver/api_key/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["new_relic_receiver", "api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/new_relic_receiver/api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "new_relic_receiver.api_key for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# new_relic_receiver.api_key

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [new_relic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/new_relic_receiver/)
- new_relic_receiver.api_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/new_relic_receiver/api_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/new_relic_receiver/api_key/clear_secret_info/): complete subsection reference.

## Next pages

- [new_relic_receiver.api_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/new_relic_receiver/api_key/blindfold_secret_info/)
- [new_relic_receiver.api_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/new_relic_receiver/api_key/clear_secret_info/)
- [new_relic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/new_relic_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
