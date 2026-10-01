---
page_title: "new_relic_receiver"
subcategory: ""
description: "new_relic_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1532, "body_sha256": "sha256:7c37098041dba18f2b1ca07a6ea18aedc5abde29f672f8da5b8c4720e138620b", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver:api_key", "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver:eu", "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver:us"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:new_relic_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "docs/guides/data-sources--global_log_receiver--properties--new_relic_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["new_relic_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/new_relic_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "new_relic_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# new_relic_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- new_relic_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for new relic receiver.

Upstream description:

Configuration for NewRelic endpoint.

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

## Direct properties

- [api_key](data-sources--global_log_receiver--properties--new_relic_receiver--api_key.md): complete subsection reference.

- [eu](data-sources--global_log_receiver--properties--new_relic_receiver--eu.md): complete subsection reference.

- [us](data-sources--global_log_receiver--properties--new_relic_receiver--us.md): complete subsection reference.

## Next pages

- [new_relic_receiver.api_key](data-sources--global_log_receiver--properties--new_relic_receiver--api_key.md)
- [new_relic_receiver.eu](data-sources--global_log_receiver--properties--new_relic_receiver--eu.md)
- [new_relic_receiver.us](data-sources--global_log_receiver--properties--new_relic_receiver--us.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
