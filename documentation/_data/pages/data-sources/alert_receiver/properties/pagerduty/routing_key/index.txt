---
page_title: "pagerduty.routing_key"
subcategory: ""
description: "pagerduty.routing_key for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1842, "body_sha256": "sha256:f52d39ef3a8ef4758fb6aead1bb1a54ebe319e8b9109c3f0ca547d8581e6f073", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:pagerduty:routing_key:blindfold_secret_info", "xcsh-docs:data-sources:alert_receiver:properties:pagerduty:routing_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:pagerduty:routing_key", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:pagerduty", "path": "documentation/data-sources/alert_receiver/properties/pagerduty/routing_key/index.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["pagerduty", "routing_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/pagerduty/routing_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "pagerduty.routing_key for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# pagerduty.routing_key

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/)
- [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/pagerduty/)
- pagerduty.routing_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/pagerduty/routing_key/clear_secret_info/): complete subsection reference.

## Next pages

- [pagerduty.routing_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/pagerduty/routing_key/blindfold_secret_info/)
- [pagerduty.routing_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/pagerduty/routing_key/clear_secret_info/)
- [pagerduty](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/properties/pagerduty/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_receiver/)
