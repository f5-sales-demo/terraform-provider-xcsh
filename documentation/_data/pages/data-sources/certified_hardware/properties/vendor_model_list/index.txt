---
page_title: "vendor_model_list"
subcategory: ""
description: "List of supported hardware vendor and model for this certified hardware."
xcsh_docs: {"aliases": ["vendor model list"], "body_bytes": 1159, "body_sha256": "sha256:1a48e1d8203ccd2e553e263f4519f055f171356228a14fca9fc074314068deeb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:vendor_model_list", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/vendor_model_list/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1210110121010203-3220102133233012-0133010210100132-0022020322110100-1301123111120033-0220003210120331-3020311203330112-0103022201331031", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vendor_model_list"], "schema_version": 1, "sections": [{"aliases": ["vendor model list model"], "anchor": "schema-vendor_model_list--model", "description": "Hw Model or instance type from cloud provider like number of interfaces, vCPUs, memory.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:vendor_model_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vendor_model_list", "model"], "syntax": "attribute", "type": "string"}, {"aliases": ["vendor model list vendor"], "anchor": "schema-vendor_model_list--vendor", "description": "Vendor could be F5 Distributed Cloud, Dell, Cloud provider like AWS or Azure.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:vendor_model_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vendor_model_list", "vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/vendor_model_list/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of supported hardware vendor and model for this certified hardware.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vendor_model_list

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- vendor_model_list

<a id="section"></a>

Type: `"list"`. Computed.

List of supported hardware vendor and model for this certified hardware.

## Direct properties

<a id="schema-vendor_model_list--model"></a>

### model property

Type: `"string"`. Computed.

Hw Model or instance type from cloud provider like number of interfaces, vCPUs, memory.

<a id="schema-vendor_model_list--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor could be F5 Distributed Cloud, Dell, Cloud provider like AWS or Azure.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
