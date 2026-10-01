---
page_title: "opsgenie.api_key"
subcategory: ""
description: "opsgenie.api_key for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2174, "body_sha256": "sha256:cd685f44ef9a729af2ea5e0304ab6450bf6b537450c9beb010d64f0e7181467f", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:opsgenie:api_key:blindfold_secret_info", "xcsh-docs:resources:alert_receiver:properties:opsgenie:api_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:opsgenie:api_key", "parent_id": "xcsh-docs:resources:alert_receiver:properties:opsgenie", "path": "documentation/resources/alert_receiver/properties/opsgenie/api_key/index.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["opsgenie", "api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/opsgenie/api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "opsgenie.api_key for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# opsgenie.api_key

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/)
- opsgenie.api_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
api_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/clear_secret_info/): complete subsection reference.

## Next pages

- [opsgenie.api_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/blindfold_secret_info/)
- [opsgenie.api_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/api_key/clear_secret_info/)
- [opsgenie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/opsgenie/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
