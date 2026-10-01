---
page_title: "webhook.url"
subcategory: ""
description: "webhook.url for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1400, "body_sha256": "sha256:c0a401e6003ff32891c2036862639e5df1a232cb218177f9ef770cfb467c5917", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:url", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:webhook:url:blindfold_secret_info", "xcsh-docs:data-sources:alert_receiver:properties:webhook:url:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:webhook:url", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:webhook", "path": "docs/guides/data-sources--alert_receiver--properties--webhook--url.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["webhook", "url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/webhook/url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "webhook.url for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.url

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
- webhook.url

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

- [blindfold_secret_info](data-sources--alert_receiver--properties--webhook--url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--properties--webhook--url--clear_secret_info.md): complete subsection reference.

## Next pages

- [webhook.url.blindfold_secret_info](data-sources--alert_receiver--properties--webhook--url--blindfold_secret_info.md)
- [webhook.url.clear_secret_info](data-sources--alert_receiver--properties--webhook--url--clear_secret_info.md)
- [webhook](data-sources--alert_receiver--properties--webhook.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
