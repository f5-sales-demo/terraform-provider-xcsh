---
page_title: "slack.url"
subcategory: ""
description: "slack.url for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1649, "body_sha256": "sha256:7ad0eae01308e64fae317d96618fd1e2a694d8a7ff4bc237bfed44a6de016251", "canonical_id": "xcsh-docs:resources:alert_receiver:properties:slack:url", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:slack:url:blindfold_secret_info", "xcsh-docs:resources:alert_receiver:properties:slack:url:clear_secret_info"], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:slack:url", "parent_id": "xcsh-docs:resources:alert_receiver:properties:slack", "path": "docs/guides/resources--alert_receiver--properties--slack--url.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["slack", "url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/slack/url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "slack.url for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slack.url

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Property reference](resources--alert_receiver--reference.md)
- [slack](resources--alert_receiver--properties--slack.md)
- slack.url

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
url {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--alert_receiver--properties--slack--url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--alert_receiver--properties--slack--url--clear_secret_info.md): complete subsection reference.

## Next pages

- [slack.url.blindfold_secret_info](resources--alert_receiver--properties--slack--url--blindfold_secret_info.md)
- [slack.url.clear_secret_info](resources--alert_receiver--properties--slack--url--clear_secret_info.md)
- [slack](resources--alert_receiver--properties--slack.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
