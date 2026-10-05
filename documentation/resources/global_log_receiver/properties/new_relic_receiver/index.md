---
page_title: "new_relic_receiver"
subcategory: ""
description: "Configuration for NewRelic endpoint."
xcsh_docs: {"aliases": ["new relic receiver"], "body_bytes": 2291, "body_sha256": "sha256:26231cb54ab9e237a4c5534782c4072f532a02b4105f0e309f8318fbf0bcbc70", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/new_relic_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver:ConflictingObjectAttributes:eu,us", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver:ConflictingObjectAttributes:eu,us", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["new_relic_receiver"], "schema_version": 1, "sections": [{"aliases": ["new relic receiver api key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver.api_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "new_relic_receiver.api_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:api_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["new_relic_receiver", "api_key"], "syntax": "block", "type": "object"}, {"aliases": ["new relic receiver eu"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:eu", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["new_relic_receiver", "eu"], "syntax": "attribute", "type": "object"}, {"aliases": ["new relic receiver us"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:new_relic_receiver:us", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["new_relic_receiver", "us"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/new_relic_receiver/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration for NewRelic endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
