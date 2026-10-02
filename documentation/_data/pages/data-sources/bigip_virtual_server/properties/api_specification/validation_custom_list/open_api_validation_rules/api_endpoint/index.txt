---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.api_endpoint"
subcategory: ""
description: "API Endpoint. This defines API endpoint."
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules api endpoint"], "body_bytes": 2215, "body_sha256": "sha256:1b0ca2d4854286ad7eceb3451ff8fd90806574ac4449680e274f4dcb9d40ef7d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules", "path": "documentation/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3212033202223010-1022002030122221-2021031203333123-1200213123322120-2102111110300020-1303213300022330-0003021012331022-3122302331203232", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint"], "schema_version": 1, "sections": [{"aliases": ["methods"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--methods", "description": "Methods. Methods to be matched. Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`. Defaults to `ANY`.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["path"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--path", "description": "Path. Path to be matched.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "API Endpoint. This defines API endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="section"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

## Direct properties

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--methods"></a>

### methods property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path to be matched.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
