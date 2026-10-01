---
page_title: "ns_all"
subcategory: ""
description: "ns_all for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1275, "body_sha256": "sha256:4d6057f373a63351ad9ff7c1033de8a9872445f55d53160f64abc922c38fc2b9", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:ns_all", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:ns_all", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--ns_all.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ns_all"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/ns_all/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ns_all for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ns_all

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- ns_all

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: ns\_all, ns\_current, ns\_list\] Enable this option

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

OneOf alternatives in this subsection:

- [ns_all](resources--global_log_receiver--properties--ns_all.md#section)
- [ns_current](resources--global_log_receiver--properties--ns_current.md#section)
- [ns_list](resources--global_log_receiver--properties--ns_list.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ns_all = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
