---
page_title: "enable_api_discovery.custom_api_auth_discovery.api_discovery_ref"
subcategory: ""
description: "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["enable api discovery custom api auth discovery api discovery ref"], "body_bytes": 2303, "body_sha256": "sha256:18f0fad64ec28b7499d21597eaf63cf5dced192144d175a015faab3002c8369a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:custom_api_auth_discovery", "path": "documentation/data-sources/third_party_application/properties/enable_api_discovery/custom_api_auth_discovery/api_discovery_ref/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3201130232000103-0333101033310320-0221010031131322-0022102031100023-1202012123220212-0001211132123011-2231101201330033-2233331313320021", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery custom api auth discovery api discovery ref name"], "anchor": "schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable api discovery custom api auth discovery api discovery ref namespace"], "anchor": "schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable api discovery custom api auth discovery api discovery ref tenant"], "anchor": "schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:custom_api_auth_discovery:api_discovery_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery", "api_discovery_ref", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/custom_api_auth_discovery/api_discovery_ref/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/)
- [enable_api_discovery.custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/custom_api_auth_discovery/)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="section"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

## Direct properties

<a id="schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="schema-enable_api_discovery--custom_api_auth_discovery--api_discovery_ref--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

## Next pages

- [enable_api_discovery.custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/custom_api_auth_discovery/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
