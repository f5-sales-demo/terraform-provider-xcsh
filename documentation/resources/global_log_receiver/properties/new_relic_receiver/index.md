---
page_title: "new_relic_receiver"
subcategory: ""
description: "new_relic_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2291, "body_sha256": "sha256:26231cb54ab9e237a4c5534782c4072f532a02b4105f0e309f8318fbf0bcbc70", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/new_relic_receiver/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["new_relic_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/new_relic_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "new_relic_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# new_relic_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
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

- [api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/): complete subsection reference.

- [eu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/eu/): complete subsection reference.

- [us](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/us/): complete subsection reference.

## Next pages

- [new_relic_receiver.api_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/api_key/)
- [new_relic_receiver.eu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/eu/)
- [new_relic_receiver.us](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/us/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
