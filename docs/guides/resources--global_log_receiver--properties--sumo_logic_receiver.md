---
page_title: "sumo_logic_receiver"
subcategory: ""
description: "sumo_logic_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1147, "body_sha256": "sha256:c4bcc0f620aba6459af69d1cd2e469119a22b2d1dbec7fefc619f5585e9f5629", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--sumo_logic_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sumo_logic_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/sumo_logic_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sumo_logic_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sumo_logic_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- sumo_logic_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sumo logic receiver.

Upstream description:

Configuration for SumoLogic endpoint.

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
sumo_logic_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [url](resources--global_log_receiver--properties--sumo_logic_receiver--url.md): complete subsection reference.

## Next pages

- [sumo_logic_receiver.url](resources--global_log_receiver--properties--sumo_logic_receiver--url.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
