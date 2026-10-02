---
page_title: "virtual_server.address_translation"
subcategory: ""
description: "Specifies, when checked (enabled), that the system translates the address of the virtual server. When cleared (disabled), specifies that the system uses the address without translation. This option is useful when the system is load balancing devices that have the same IP address. The default is enabled."
xcsh_docs: {"aliases": ["virtual server address translation"], "body_bytes": 2749, "body_sha256": "sha256:c3a28aadeee9a51d6b855f5659d0b55e47354fba212718132129a2fc372f3575", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation:address_translation_disable", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation:address_translation_enable"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation", "parent_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "path": "documentation/data-sources/application_profiles/properties/virtual_server/address_translation/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1331202223132031-0030211010012330-0023312023321130-1130322230022331-1311131323321212-0122132311013200-3032201032200000-2221121310313101", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "address_translation"], "schema_version": 1, "sections": [{"aliases": ["address translation disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation:address_translation_disable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "address_translation", "address_translation_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["address translation enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation:address_translation_enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "address_translation", "address_translation_enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/address_translation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies, when checked (enabled), that the system translates the address of the virtual server. When cleared (disabled), specifies that the system uses the address without translation. This option is useful when the system is load balancing devices that have the same IP address. The default is enabled.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.address_translation

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- virtual_server.address_translation

<a id="section"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address.

Upstream description:

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address. The default is
enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_translation_choice": "[\"address_translation_disable\",\"address_translation_enable\"]"
}
```

## Direct properties

- [address_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/address_translation_disable/): complete subsection reference.

- [address_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/address_translation_enable/): complete subsection reference.

## Next pages

- [virtual_server.address_translation.address_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/address_translation_disable/)
- [virtual_server.address_translation.address_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/address_translation_enable/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
