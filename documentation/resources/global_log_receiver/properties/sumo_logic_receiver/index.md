---
page_title: "sumo_logic_receiver"
subcategory: ""
description: "Configuration for SumoLogic endpoint."
xcsh_docs: {"aliases": ["sumo logic receiver"], "body_bytes": 1455, "body_sha256": "sha256:7946430204224efd1115462586892ffb937f2ff0fc1f9259eed9a225bc5c125e", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/sumo_logic_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sumo_logic_receiver"], "schema_version": 1, "sections": [{"aliases": ["sumo logic receiver url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "sumo_logic_receiver.url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "sumo_logic_receiver.url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:clear_secret_info", "type": "conflicts"}], "schema_path": ["sumo_logic_receiver", "url"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/sumo_logic_receiver/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration for SumoLogic endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sumo_logic_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
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

- [url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/): complete subsection reference.

## Next pages

- [sumo_logic_receiver.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
