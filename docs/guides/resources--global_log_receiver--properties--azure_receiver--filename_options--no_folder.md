---
page_title: "azure_receiver.filename_options.no_folder"
subcategory: ""
description: "azure_receiver.filename_options.no_folder for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1215, "body_sha256": "sha256:78530e0c3a79ea7962d32ed59ddf31d6105665310f9d66b5b69fe46ab421f9a8", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:no_folder", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "path": "docs/guides/resources--global_log_receiver--properties--azure_receiver--filename_options--no_folder.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure_receiver", "filename_options", "no_folder"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/filename_options/no_folder/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure_receiver.filename_options.no_folder for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.filename_options.no_folder

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [azure_receiver](resources--global_log_receiver--properties--azure_receiver.md)
- [azure_receiver.filename_options](resources--global_log_receiver--properties--azure_receiver--filename_options.md)
- azure_receiver.filename_options.no_folder

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
no_folder = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [azure_receiver.filename_options](resources--global_log_receiver--properties--azure_receiver--filename_options.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
