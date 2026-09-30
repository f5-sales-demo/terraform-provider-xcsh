---
page_title: "new_relic_receiver"
subcategory: ""
description: "new_relic_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1684, "body_sha256": "sha256:1f2565ee3c926161a0f7ad4fe1570c54c903dc937865fe1e72318055e1389bc9", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--new_relic_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["new_relic_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/new_relic_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "new_relic_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# new_relic_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- new_relic_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for new relic receiver.

Upstream description:

Configuration for NewRelic endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("eu",
    "us")}
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
  "x-ves-oneof-field-endpoint_choice": "[\"eu\",\"us\"]"
}
```

Terraform syntax:

```terraform
new_relic_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_key](resources--global_log_receiver--properties--new_relic_receiver--api_key.md): complete subsection reference.

- [eu](resources--global_log_receiver--properties--new_relic_receiver--eu.md): complete subsection reference.

- [us](resources--global_log_receiver--properties--new_relic_receiver--us.md): complete subsection reference.

## Next pages

- [new_relic_receiver.api_key](resources--global_log_receiver--properties--new_relic_receiver--api_key.md)
- [new_relic_receiver.eu](resources--global_log_receiver--properties--new_relic_receiver--eu.md)
- [new_relic_receiver.us](resources--global_log_receiver--properties--new_relic_receiver--us.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
