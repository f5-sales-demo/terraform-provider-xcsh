---
page_title: "opsgenie.api_key"
subcategory: ""
description: "opsgenie.api_key for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1345, "body_sha256": "sha256:ca6e9a212d1c3d6a131add02a42473b21399c9c544f9bc4692bfb57d36d8d413", "canonical_id": "xcsh-docs:data-sources:alert_receiver:properties:opsgenie:api_key", "child_ids": ["xcsh-docs:data-sources:alert_receiver:properties:opsgenie:api_key:blindfold_secret_info", "xcsh-docs:data-sources:alert_receiver:properties:opsgenie:api_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_receiver:properties:opsgenie:api_key", "parent_id": "xcsh-docs:data-sources:alert_receiver:properties:opsgenie", "path": "docs/guides/data-sources--alert_receiver--properties--opsgenie--api_key.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["opsgenie", "api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_receiver/properties/opsgenie/api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "opsgenie.api_key for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# opsgenie.api_key

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
- [Property reference](data-sources--alert_receiver--reference.md)
- [opsgenie](data-sources--alert_receiver--properties--opsgenie.md)
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

- [blindfold_secret_info](data-sources--alert_receiver--properties--opsgenie--api_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--alert_receiver--properties--opsgenie--api_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [opsgenie.api_key.blindfold_secret_info](data-sources--alert_receiver--properties--opsgenie--api_key--blindfold_secret_info.md)
- [opsgenie.api_key.clear_secret_info](data-sources--alert_receiver--properties--opsgenie--api_key--clear_secret_info.md)
- [opsgenie](data-sources--alert_receiver--properties--opsgenie.md)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md)
