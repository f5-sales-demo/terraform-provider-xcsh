---
page_title: "sumo_logic_receiver.url"
subcategory: ""
description: "sumo_logic_receiver.url for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1872, "body_sha256": "sha256:247d881fb2f59d1f3c09ae07ea219d7f4df6ee4dc1b9cdfb393154adb809b1cf", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:clear_secret_info"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver", "path": "docs/guides/resources--global_log_receiver--properties--sumo_logic_receiver--url.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sumo_logic_receiver", "url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/sumo_logic_receiver/url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sumo_logic_receiver.url for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sumo_logic_receiver.url

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [sumo_logic_receiver](resources--global_log_receiver--properties--sumo_logic_receiver.md)
- sumo_logic_receiver.url

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

- [blindfold_secret_info](resources--global_log_receiver--properties--sumo_logic_receiver--url--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--properties--sumo_logic_receiver--url--clear_secret_info.md): complete subsection reference.

## Next pages

- [sumo_logic_receiver.url.blindfold_secret_info](resources--global_log_receiver--properties--sumo_logic_receiver--url--blindfold_secret_info.md)
- [sumo_logic_receiver.url.clear_secret_info](resources--global_log_receiver--properties--sumo_logic_receiver--url--clear_secret_info.md)
- [sumo_logic_receiver](resources--global_log_receiver--properties--sumo_logic_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
