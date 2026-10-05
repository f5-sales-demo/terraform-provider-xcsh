---
page_title: "stateful_service.advertise_options"
subcategory: "Container"
description: "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers."
xcsh_docs: {"aliases": ["stateful service advertise options"], "body_bytes": 3591, "body_sha256": "sha256:f9c8805996f9db99317c032d6c9dcacbca1df32def2d81f19e3b05d0945e1871", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service", "path": "documentation/resources/workload/properties/stateful_service/advertise_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323", "registry_path": "docs/guides/resources--workload--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_in_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_custom,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_in_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_on_public,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_custom,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options:ConflictingObjectAttributes:advertise_on_public,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise custom"], "anchor": "section", "description": "Advertise this workload via loadbalancer on specific sites.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom:RequiredObjectAttributes:advertise_where,ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:advertise_where", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_custom:RequiredObjectAttributes:advertise_where,ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports", "type": "requires"}], "schema_path": ["stateful_service", "advertise_options", "advertise_custom"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise in cluster"], "anchor": "section", "description": "Advertise the workload locally in-cluster.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_in_cluster:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_in_cluster:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:port", "type": "conflicts"}], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public"], "anchor": "section", "description": "Advertise this workload via loadbalancer on Internet with default VIP.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "stateful_service.advertise_options.advertise_on_public:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port", "type": "conflicts"}], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public"], "syntax": "block", "type": "object"}, {"aliases": ["stateful service advertise options do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:do_not_advertise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "do_not_advertise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- stateful_service.advertise_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_in_cluster"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/): complete subsection reference.

- [advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/): complete subsection reference.

- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/): complete subsection reference.

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/do_not_advertise/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- [stateful_service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/advertise_options/do_not_advertise/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
