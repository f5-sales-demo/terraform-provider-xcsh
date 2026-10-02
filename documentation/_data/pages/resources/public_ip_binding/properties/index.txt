---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_public_ip_binding."
xcsh_docs: {"aliases": ["public ip binding"], "body_bytes": 2991, "body_sha256": "sha256:ca59e75c7e66e7eab02624a3a556356f9599d97206eda5daaf926b5eb581273a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:public_ip_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:public_ip_binding:reference", "parent_id": "xcsh-docs:resources:public_ip_binding:fundamentals", "path": "documentation/resources/public_ip_binding/properties/index.md", "product": "distributed-cloud", "provider_name": "public_ip_binding", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0000123303013230-2022101212032211-3303121103322023-1113210013222113-2332132123021221-0201100212311212-3322201001001113-1203020333320300", "registry_path": "docs/guides/resources--public_ip_binding--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["expected ip"], "anchor": "schema-expected_ip", "description": "Exact IPv4/IPv6 address expected in the allocated object.", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "id", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["managed bindings"], "anchor": "schema-managed_bindings", "description": "Last applied binding retained to reject foreign changes at teardown.", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["managed_bindings"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Existing public IP object name.", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Existing public IP namespace, usually shared.", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["original bindings"], "anchor": "schema-original_bindings", "description": "Original virtual-site bindings restored on delete; kept in Terraform state.", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["original_bindings"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual site"], "anchor": "schema-virtual_site", "description": "Desired REGIONAL_EDGE virtual site. Refreshed from the actual binding so drift is repairable.", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_site"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual site namespace"], "anchor": "schema-virtual_site_namespace", "description": "Namespace of the regional virtual site.", "document_id": "xcsh-docs:resources:public_ip_binding:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_site_namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/public_ip_binding/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_public_ip_binding.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_public_ip_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/)
- Property reference

## Direct properties

<a id="schema-expected_ip"></a>

### expected_ip property

Type: `"string"`. Required.

Exact IPv4/IPv6 address expected in the allocated object.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

<a id="schema-managed_bindings"></a>

### managed_bindings property

Type: `"string"`. Computed.

Last applied binding retained to reject foreign changes at teardown.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Existing public IP object name.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Existing public IP namespace, usually shared.

<a id="schema-original_bindings"></a>

### original_bindings property

Type: `"string"`. Computed.

Original virtual-site bindings restored on delete; kept in Terraform state.

<a id="schema-virtual_site"></a>

### virtual_site property

Type: `"string"`. Required.

Desired REGIONAL\_EDGE virtual site. Refreshed from the actual binding so drift is repairable.

<a id="schema-virtual_site_namespace"></a>

### virtual_site_namespace property

Type: `"string"`. Required.

Namespace of the regional virtual site.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `expected_ip` | [expected_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-expected_ip) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-id) |
| `managed_bindings` | [managed_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-managed_bindings) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-namespace) |
| `original_bindings` | [original_bindings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-original_bindings) |
| `virtual_site` | [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-virtual_site) |
| `virtual_site_namespace` | [virtual_site_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/#schema-virtual_site_namespace) |

## Next pages

- [xcsh_public_ip_binding](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/)
