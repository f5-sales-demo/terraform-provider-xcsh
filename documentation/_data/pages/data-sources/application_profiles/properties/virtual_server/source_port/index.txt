---
page_title: "virtual_server.source_port"
subcategory: ""
description: "Specifies whether the system preserves the source port of the connection. The default is Preserve."
xcsh_docs: {"aliases": ["virtual server source port"], "body_bytes": 2563, "body_sha256": "sha256:b69366f97844ee75dcf1de86cfcac85bba25e8927da7e400a636ccb64ce543a7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_change", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/source_port/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "source_port"], "schema_version": 1, "sections": [{"aliases": ["virtual server source port source port change"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_change", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_change"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server source port source port preserve"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_preserve"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server source port source port preserve strict"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "source_port", "source_port_preserve_strict"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/source_port/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specifies whether the system preserves the source port of the connection. The default is Preserve.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.source_port

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.source_port

<a id="section"></a>

Type: `"single"`. Computed.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

## Direct properties

- [source_port_change](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_change/): complete subsection reference.

- [source_port_preserve](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_preserve/): complete subsection reference.

- [source_port_preserve_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/): complete subsection reference.

## Next pages

- [virtual_server.source_port.source_port_change](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_change/)
- [virtual_server.source_port.source_port_preserve](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_preserve/)
- [virtual_server.source_port.source_port_preserve_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
