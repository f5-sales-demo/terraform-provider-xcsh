---
page_title: "palo_alto_fw_service.service_nodes"
subcategory: ""
description: "Configuration parameter for service nodes."
xcsh_docs: {"aliases": ["palo alto fw service service nodes"], "body_bytes": 1459, "body_sha256": "sha256:989ba21fe788da7cbf10ad6d31d4caa235e75f58165df185f4f78671a2fe565a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2320101032110032-1132301213320202-3223120220002213-1110110111210103-3332303021330201-3211121030111103-3232223021200202-2320023112312222", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes"], "schema_version": 1, "sections": [{"aliases": ["nodes"], "anchor": "section", "description": "Configuration parameter for nodes", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["palo_alto_fw_service", "service_nodes", "nodes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for service nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.service_nodes

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- palo_alto_fw_service.service_nodes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for service nodes.

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

## Direct properties

- [nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.service_nodes.nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/service_nodes/nodes/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
