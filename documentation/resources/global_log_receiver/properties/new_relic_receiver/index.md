---
page_title: "new_relic_receiver"
subcategory: ""
description: "Configuration for NewRelic endpoint."
xcsh_docs: {"aliases": ["new relic receiver"], "body_bytes": 1622, "body_sha256": "sha256:37a9371c6211dcbee857baefea6d645387c2564a8a92ae841327e274079ee66d", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/new_relic_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver:ConflictingObjectAttributes:eu,us", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver:ConflictingObjectAttributes:eu,us", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["new_relic_receiver"], "schema_version": 1, "sections": [{"aliases": ["new relic receiver api key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver.api_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver.api_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["new_relic_receiver", "api_key"], "syntax": "block", "type": "object"}, {"aliases": ["new relic receiver eu"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["new_relic_receiver", "eu"], "syntax": "attribute", "type": "object"}, {"aliases": ["new relic receiver us"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["new_relic_receiver", "us"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/new_relic_receiver/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration for NewRelic endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

Additional upstream details:

Configuration for NewRelic endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
