---
page_title: "virtual_server.connection_rate_limit_mode"
subcategory: ""
description: "Configuration parameter for connection rate limit mode."
xcsh_docs: {"aliases": ["virtual server connection rate limit mode"], "body_bytes": 4928, "body_sha256": "sha256:506e0eaf0fbed81a3cf5a5f79edce7e51e921eb155028f37faa26bfe8c99751a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_destination_address"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode"], "schema_version": 1, "sections": [{"aliases": ["per destination address"], "anchor": "section", "description": "Destination Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_destination_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_destination_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["per source address"], "anchor": "section", "description": "Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_source_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["per source destination address"], "anchor": "section", "description": "Destination and Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_source_destination_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_source_destination_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["per virtual server"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["per virtual server destination address"], "anchor": "section", "description": "Destination Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_destination_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_destination_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["per virtual server source address"], "anchor": "section", "description": "Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_source_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["per virtual server source destination address"], "anchor": "section", "description": "Destination and Source Address Mask.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_destination_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_source_destination_address"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for connection rate limit mode.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [virtual_server.connection_rate_limit_mode.per_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_destination_address/)
- [virtual_server.connection_rate_limit_mode.per_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_address/)
- [virtual_server.connection_rate_limit_mode.per_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_destination_address/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
