---
page_title: "disable_ssh_access"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable ssh access"], "body_bytes": 1593, "body_sha256": "sha256:433ec1644c034e5d14a11b0ada4ce4d721c1ee6ad48f8cb6b27c72587a5180af", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:disable_ssh_access", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "documentation/resources/nfv_service/properties/disable_ssh_access/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1121230113200220-3301032102311323-1313313310300012-0131301302121233-3323233302031130-1303122223201233-3332210000021300-3200231130103032", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_ssh_access"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/disable_ssh_access/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_ssh_access

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- disable_ssh_access

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ssh\_access, enabled\_ssh\_access; Default: disable\_ssh\_access\] Configuration
parameter for disable ssh access.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [disable_ssh_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/disable_ssh_access/#section)
- [enabled_ssh_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ssh_access = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
