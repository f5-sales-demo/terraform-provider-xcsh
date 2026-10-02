---
page_title: "service.advertise_options.advertise_custom"
subcategory: "Container"
description: "Advertise this workload via loadbalancer on specific sites."
xcsh_docs: {"aliases": ["service advertise options advertise custom"], "body_bytes": 2298, "body_sha256": "sha256:f1a5becc0451b08db8457bf91a3f94e19d84e1548a5b91832b50706475892eea", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom:RequiredObjectAttributes:advertise_where,ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom:RequiredObjectAttributes:advertise_where,ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom"], "schema_version": 1, "sections": [{"aliases": ["advertise where"], "anchor": "section", "description": "Where should this load balancer be available.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_custom", "advertise_where"], "syntax": "block", "type": "object"}, {"aliases": ["ports"], "anchor": "section", "description": "Ports to advertise.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports:ConflictingListObjectAttributes:http_loadbalancer,tcp_loadbalancer", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports:ConflictingListObjectAttributes:http_loadbalancer,tcp_loadbalancer", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:tcp_loadbalancer", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_custom", "ports"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Advertise this workload via loadbalancer on specific sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- service.advertise_options.advertise_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where",
    "ports")}
```

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
advertise_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/): complete subsection reference.

- [ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/advertise_where/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
