---
page_title: "https_management.advertise_on_internet"
subcategory: ""
description: "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available."
xcsh_docs: {"aliases": ["https management advertise on internet"], "body_bytes": 1904, "body_sha256": "sha256:6de525a8473716782e6f7f6b33766b211b90b45c6cb8f3e57d756d54df0a3c13", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_internet:public_ip"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_internet", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management", "path": "documentation/resources/nfv_service/properties/https_management/advertise_on_internet/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0120200112111003-3331321000320321-3001301020110212-2233312213000030-3121001203022212-2312200333022211-2120220200002130-3331131211211023", "registry_path": "docs/guides/resources--nfv_service--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_internet"], "schema_version": 1, "sections": [{"aliases": ["public ip"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_internet:public_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-https_management--advertise_on_internet--public_ip--name", "enforcement": "provider-schema", "group": "https_management.advertise_on_internet.public_ip:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_internet:public_ip", "type": "requires"}], "schema_path": ["https_management", "advertise_on_internet", "public_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_internet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_internet

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/)
- https_management.advertise_on_internet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

Terraform syntax:

```terraform
advertise_on_internet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_internet/public_ip/): complete subsection reference.

## Next pages

- [https_management.advertise_on_internet.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_internet/public_ip/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
