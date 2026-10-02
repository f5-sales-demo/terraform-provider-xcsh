---
page_title: "service.advertise_options"
subcategory: "Container"
description: "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers."
xcsh_docs: {"aliases": ["service advertise options"], "body_bytes": 3429, "body_sha256": "sha256:bf6a57147c8cc1d8c17960571f28f4501afc1cf45afda9226f0bd1b8061d3ddb", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "xcsh-docs:resources:workload:properties:service:advertise_options:do_not_advertise"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options", "parent_id": "xcsh-docs:resources:workload:properties:service", "path": "documentation/resources/workload/properties/service/advertise_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130", "registry_path": "docs/guides/resources--workload--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_in_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_custom,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_in_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_custom,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,advertise_on_public", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_on_public,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_custom,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:do_not_advertise", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_in_cluster,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:do_not_advertise", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options:ConflictingObjectAttributes:advertise_on_public,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:do_not_advertise", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options"], "schema_version": 1, "sections": [{"aliases": ["advertise custom"], "anchor": "section", "description": "Advertise this workload via loadbalancer on specific sites.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom:RequiredObjectAttributes:advertise_where,ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:advertise_where", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom:RequiredObjectAttributes:advertise_where,ports", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports", "type": "requires"}], "schema_path": ["service", "advertise_options", "advertise_custom"], "syntax": "block", "type": "object"}, {"aliases": ["advertise in cluster"], "anchor": "section", "description": "Advertise the workload locally in-cluster.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_in_cluster:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_in_cluster:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_in_cluster:port", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_in_cluster"], "syntax": "block", "type": "object"}, {"aliases": ["advertise on public"], "anchor": "section", "description": "Advertise this workload via loadbalancer on Internet with default VIP.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_on_public:ConflictingObjectAttributes:multi_ports,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_on_public"], "syntax": "block", "type": "object"}, {"aliases": ["do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:do_not_advertise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "do_not_advertise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- service.advertise_options

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

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/): complete subsection reference.

- [advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/): complete subsection reference.

- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/): complete subsection reference.

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/do_not_advertise/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_in_cluster/)
- [service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_on_public/)
- [service.advertise_options.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/do_not_advertise/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
