---
page_title: "opsgenie.api_key"
subcategory: ""
description: "opsgenie.api_key for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1897, "body_sha256": "sha256:37592605c9a573d1a8d99ce3263e05c14bfc5eb24ef09057a9879a0fd5a9e14f", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:opsgenie:api_key:blindfold_secret_info", "xcsh-docs:data-sources:alert_receiver:properties:opsgenie:api_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:opsgenie:api_key", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:opsgenie", "path": "documentation/data-sources/alert_receiver/properties/opsgenie/api_key/index.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["opsgenie", "api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/opsgenie/api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "opsgenie.api_key for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# opsgenie.api_key

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/opsgenie/)
- opsgenie.api_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/opsgenie/api_key/clear_secret_info/): complete subsection reference.

## Next pages

- [opsgenie.api_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/)
- [opsgenie.api_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/opsgenie/api_key/clear_secret_info/)
- [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/opsgenie/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
