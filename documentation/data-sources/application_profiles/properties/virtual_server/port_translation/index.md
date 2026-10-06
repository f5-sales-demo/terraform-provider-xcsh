---
page_title: "virtual_server.port_translation"
subcategory: ""
description: "Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any"
xcsh_docs: {"aliases": ["virtual server port translation"], "body_bytes": 1713, "body_sha256": "sha256:0c8932a31c70942a7807429307390edf87530de6fd225f22c5a811e7b1bf3769", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation:port_translation_enable"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/port_translation/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "port_translation"], "schema_version": 1, "sections": [{"aliases": ["virtual server port translation port translation disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation:port_translation_disable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "port_translation", "port_translation_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server port translation port translation enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation:port_translation_enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "port_translation", "port_translation_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/port_translation/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.port_translation

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.port_translation

<a id="section"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

## Direct properties

- [port_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/): complete subsection reference.

- [port_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/port_translation/port_translation_enable/): complete subsection reference.
