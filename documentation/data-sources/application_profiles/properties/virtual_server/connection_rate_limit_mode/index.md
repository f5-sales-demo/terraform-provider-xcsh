---
page_title: "virtual_server.connection_rate_limit_mode"
subcategory: ""
description: "Configuration parameter for connection rate limit mode."
xcsh_docs: {"aliases": ["virtual server connection rate limit mode"], "body_bytes": 2892, "body_sha256": "sha256:12560e9064dbeb5b676276eb680484f8c05de56544280cbab14ccfd3105667a5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_destination_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode"], "schema_version": 1, "sections": [{"aliases": ["virtual server connection rate limit mode per destination address"], "anchor": "section", "description": "Destination Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_destination_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_destination_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server connection rate limit mode per source address"], "anchor": "section", "description": "Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_source_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server connection rate limit mode per source destination address"], "anchor": "section", "description": "Destination and Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_source_destination_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server connection rate limit mode per virtual server"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server connection rate limit mode per virtual server destination address"], "anchor": "section", "description": "Destination Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_destination_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server connection rate limit mode per virtual server source address"], "anchor": "section", "description": "Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_source_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server connection rate limit mode per virtual server source destination address"], "anchor": "section", "description": "Destination and Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_destination_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_source_destination_address"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration parameter for connection rate limit mode.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["application_profilesCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.connection_rate_limit_mode

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.connection_rate_limit_mode

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for connection rate limit mode.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connection_rate_limit_mode_choice": "[\"per_destination_address\",\"per_source_address\",\"per_source_destination_address\",\"per_virtual_server\",\"per_virtual_server_destination_address\",\"per_virtual_server_source_address\",\"per_virtual_server_source_destination_address\"]"
}
```

## Direct properties

- [per_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_destination_address/): complete subsection reference.

- [per_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_address/): complete subsection reference.

- [per_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/): complete subsection reference.

- [per_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server/): complete subsection reference.

- [per_virtual_server_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/): complete subsection reference.

- [per_virtual_server_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/): complete subsection reference.

- [per_virtual_server_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_destination_address/): complete subsection reference.
